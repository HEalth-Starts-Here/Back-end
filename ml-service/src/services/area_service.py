import io

import protos.affected_area_pb2 as pb
import protos.affected_area_pb2_grpc as pb_grpc

from PIL import Image
from ultralytics import YOLO
from torchvision.transforms import Resize, ToPILImage


class AffectedAreaModel:
    def __init__(self, model_path: str) -> None:
        self._model = YOLO(model_path)

    def predict(self, image):
        img = Image.open(io.BytesIO(image)).convert("RGB")
        result = self._model.predict(img)
        return result


class AffectedAreaService(pb_grpc.AffectedAreaServicer):
    def __init__(self, model_path: str) -> None:
        self._model = AffectedAreaModel(model_path)

    def calculateArea(self, request, context):
        result = self._model.predict(request.image)
        if result[0].masks:
            area = result[0].masks.masks.squeeze().nonzero().shape[0]

            resizer = Resize(result[0].orig_shape, antialias=False)
            converter = ToPILImage()

            mask = resizer(result[0].masks.masks)
            mask_img = converter(mask)
            byte_mask = io.BytesIO()
            mask_img.save(byte_mask, format="JPEG")
            mask = byte_mask.getvalue()
        else:
            area = 0
            mask = None

        return pb.AffectedAreaResponse(area=area, mask=mask)
