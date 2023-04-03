import io
import os
import sys
import torch
import torchvision
import numpy as np

from PIL import Image

from protos.iqa_pb2 import IQAResponse
from protos.iqa_pb2_grpc import IQAServicer

from src.entities.predict_params import PredictParams
from src.models.iqa_models import HyperNet, TargetNet
from src.entities.logger import setup_default_logger
from src.utils import download_file


logger = setup_default_logger("iqa_logs", sys.stdout)


class IQAService(IQAServicer):
    def __init__(self, params: PredictParams) -> None:
        self.params = params
        self.device = torch.device("cuda" if torch.cuda.is_available() else "cpu")

        if params.download_params and not os.path.exists(self.params.model_path):
            download_file(self.params.download_params)

        self._model_hyper = HyperNet(16, 112, 224, 112, 56, 28, 14, 7).to(self.device)
        self._model_hyper.train(False)
        self._model_hyper.load_state_dict(
            torch.load(self.params.model_path, map_location=self.device)
        )

        self._transforms = torchvision.transforms.Compose(
            [
                torchvision.transforms.Resize((512, 384)),
                torchvision.transforms.RandomCrop(size=224),
                torchvision.transforms.ToTensor(),
                torchvision.transforms.Normalize(
                    mean=(0.485, 0.456, 0.406), std=(0.229, 0.224, 0.225)
                ),
            ]
        )

    def calculateQuality(self, request, context):
        # random crop 10 patches and calculate mean quality score
        # quality score ranges from 0-100, a higher score indicates a better quality

        pred_scores = []
        img = Image.open(io.BytesIO(request.image)).convert("RGB")
        for _ in range(self.params.num_parts):
            img_tensor = self._transforms(img).to(self.device)
            img_tensor = img_tensor.unsqueeze(0)
            weights = self._model_hyper(img_tensor)

            model_target = TargetNet(weights).to(self.device)
            for param in model_target.parameters():
                param.requires_grad = False

            pred = model_target(weights["target_in_vec"])
            pred_scores.append(float(pred.item()))
        score = np.mean(sorted(pred_scores)[-self.params.num_top_parts:])
        logger.info(f"Image score: {score}")
        return IQAResponse(quality=score > self.params.quality_threshold)
