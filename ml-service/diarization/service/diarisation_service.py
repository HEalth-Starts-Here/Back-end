import logging
import sys

# import io
# import numpy as np
# import whisper
# import datetime
# import wave
# import contextlib
# import torch
import requests
import base64
import grpc

# from pyannote.audio.pipelines.speaker_verification import PretrainedSpeakerEmbedding
# from pyannote.audio import Audio
# from pyannote.core import Segment

# from sklearn.cluster import AgglomerativeClustering

# from pydub import AudioSegment

from api.diarisation_pb2 import DiarisationResponse
from api.diarisation_pb2_grpc import DiarisationServicer
from entities.models_params import DiarisationParams


logger = logging.getLogger(__name__)
handler = logging.StreamHandler(sys.stdout)
logger.setLevel(logging.INFO)
logger.addHandler(handler)


class DiarisationService(DiarisationServicer):
    def __init__(self, params: DiarisationParams) -> None:
        self.params = params

        # device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
        # self.embedding_model = PretrainedSpeakerEmbedding(
        #     "speechbrain/spkrec-ecapa-voxceleb", device=device
        # )

    def transcribeAudio(self, request, context):
        if self.params.use_api:
            text = self._transcribe_with_api(request.audio, context)
        else:
            text = self._transcribe(request.audio)
        return DiarisationResponse(text=text)

    def _transcribe(self, audio):
        # data = io.BytesIO(audio)

        # sound = AudioSegment.from_file(data)
        # sound.export("temp/file.wav", format="wav")
        # path = "temp/file.wav"

        # model = whisper.load_model(self.params.model_size)

        # logger.info("Transcribing audio...")
        # result = model.transcribe(path)
        # segments = result["segments"]

        # with contextlib.closing(wave.open(path, "r")) as f:
        #     frames = f.getnframes()
        #     rate = f.getframerate()
        #     duration = frames / float(rate)

        # audio = Audio()

        # def segment_embedding(segment):
        #     start = segment["start"]
        #     # Whisper overshoots the end timestamp in the last segment
        #     end = min(duration, segment["end"])
        #     clip = Segment(start, end)
        #     waveform, sample_rate = audio.crop(path, clip)
        #     return self.embedding_model(waveform[None])

        # embeddings = np.zeros(shape=(len(segments), 192))
        # for i, segment in enumerate(segments):
        #     embeddings[i] = segment_embedding(segment)

        # embeddings = np.nan_to_num(embeddings)

        # clustering = AgglomerativeClustering(self.params.num_speakers).fit(embeddings)
        # labels = clustering.labels_
        # for i in range(len(segments)):
        #     segments[i]["speaker"] = "SPEAKER " + str(labels[i] + 1)

        # def time(secs):
        #     return datetime.timedelta(seconds=round(secs))

        # f = io.StringIO()

        # for i, segment in enumerate(segments):
        #     if i == 0 or segments[i - 1]["speaker"] != segment["speaker"]:
        #         f.write(
        #             "\n" + segment["speaker"] + " " + str(time(segment["start"])) + "\n"
        #         )
        #     f.write(segment["text"][1:] + " ")

        # f.getvalue()
        pass

    def _transcribe_with_api(self, audio, context):
        # Кодируем содержимое файла в base64
        base64_data = base64.b64encode(audio).decode()

        logger.info("Waiting response from API...")
        # Отправляем запрос API с использованием закодированной строки
        response = requests.post(
            "https://dwarkesh-whisper-speaker-recognition.hf.space/run/predict",
            json={
                "data": [
                    {
                        "name": "test.wav",
                        "data": "data:audio/wav;base64," + base64_data,
                    },
                    self.params.num_speakers,
                ]
            },
        )

        logger.info(f"Response status code: {response.status_code}")
        if not response.ok:
            context.abort(grpc.StatusCode.UNAVAILABLE, "HuggingFace API unavailable")

        data = response.json()["data"]
        logger.info(f"API response size: {len(data[0])}")
        return data[0]
