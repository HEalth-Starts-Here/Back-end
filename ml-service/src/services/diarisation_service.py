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

# from pyannote.audio.pipelines.speaker_verification import PretrainedSpeakerEmbedding

# from pyannote.audio import Audio
# from pyannote.core import Segment

# from sklearn.cluster import AgglomerativeClustering

from pydub import AudioSegment

from protos.diarisation_pb2 import DiarisationResponse

from protos.diarisation_pb2_grpc import DiarisationServicer
from src.entities.logger import setup_default_logger
from src.entities.models_params import DiarisationParams


logger = setup_default_logger("diarisation_logs", sys.stdout)


class DiarisationService(DiarisationServicer):
    def __init__(self, params: DiarisationParams) -> None:
        self.params = params

        # device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
        # self.embedding_model = PretrainedSpeakerEmbedding(
        #     "speechbrain/spkrec-ecapa-voxceleb", device=device
        # )

    def transcribeAudio(self, request, context):
        if self.params.use_api:
            self._transcribe_with_api(request.audio)
        else:
            self._transcribe(request.audio)

    def _transcribe(self, audio):
        pass
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

        # return DiarisationResponse(text=f.getvalue())

    def _transcribe_with_api(self, audio):
        # Открываем файл test.wav в бинарном режиме и считываем его содержимое
        # with open("test.wav", "rb") as f:
        #     wav_data = f.read()

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
                    2,
                ]
            },
        ).json()

        data = response["data"]
        logger.info(f"API response size: {len(data)}")
        return DiarisationResponse(text=data)
