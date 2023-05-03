import logging
import sys
import io
import numpy as np
import whisper
import datetime
import librosa
import torch
import requests
import base64
import grpc

from time import time
from pyannote.audio.pipelines.speaker_verification import PretrainedSpeakerEmbedding
from pyannote.audio import Audio
from pyannote.core import Segment

from sklearn.cluster import AgglomerativeClustering

from pydub import AudioSegment

from api.diarisation_pb2 import DiarisationResponse
from api.diarisation_pb2_grpc import DiarisationServicer
from entities.models_params import DiarisationParams


logger = logging.getLogger(__name__)
handler = logging.StreamHandler(sys.stdout)
logger.setLevel(logging.DEBUG)
logger.addHandler(handler)


def get_time(secs):
    return datetime.timedelta(seconds=round(secs))


class DiarisationService(DiarisationServicer):
    def __init__(self, params: DiarisationParams) -> None:
        self.params = params
        device = torch.device(self.params.device)

        self.embedding_model = PretrainedSpeakerEmbedding(
            "speechbrain/spkrec-ecapa-voxceleb", device=device
        )
        if not self.params.use_api:
            logger.info(f"Load {self.params.model_size} model")
            self.model = whisper.load_model(
                self.params.model_size, download_root="models/", device=device
            )
        logger.info("Service setup completed")

    def transcribeAudio(self, request, context):
        if self.params.use_api:
            text = self._transcribe_with_api(request.audio, context)
        else:
            text = self._transcribe(request.audio)
        return DiarisationResponse(text=text)

    def _transcribe(self, audio):
        data = io.BytesIO(audio)

        sound = AudioSegment.from_file(data)
        duration = sound.duration_seconds
        audio_wave, _ = librosa.load(sound.export(format="wav"), sr=16000)

        logger.info("Transcribing audio...")
        start_time = time()
        result = self.model.transcribe(audio_wave)
        segments = result["segments"]

        logger.debug(f"Segments count: {len(segments)}")
        if len(segments) > 1:
            audio = Audio()

            def segment_embedding(segment):
                start = segment["start"]
                # Whisper overshoots the end timestamp in the last segment
                end = min(duration, segment["end"])
                clip = Segment(start, end)
                waveform, _ = audio.crop(sound.export(format="wav"), clip)
                return self.embedding_model(waveform[None])

            embeddings = np.zeros(shape=(len(segments), 192))
            for i, segment in enumerate(segments):
                embeddings[i] = segment_embedding(segment)
            embeddings = np.nan_to_num(embeddings)

            clustering = AgglomerativeClustering(self.params.num_speakers).fit(embeddings)
            labels = clustering.labels_
            for i in range(len(segments)):
                segments[i]["speaker"] = "SPEAKER " + str(labels[i] + 1)
        elif len(segments) == 0:
            return "No text in audio"
        else:
            segments[0]["speaker"] = "SPEAKER 1"

        f = io.StringIO()
        for i, segment in enumerate(segments):
            if i == 0 or segments[i - 1]["speaker"] != segment["speaker"]:
                f.write(
                    "\n" + segment["speaker"] + " " + str(get_time(segment["start"])) + "\n"
                )
            f.write(segment["text"][1:] + " ")

        end_time = time()
        logger.debug(f"Total transcribing time: {get_time(end_time - start_time)}")

        return f.getvalue()

    def _transcribe_with_api(self, audio, context):
        # Кодируем содержимое файла в base64
        base64_data = base64.b64encode(audio).decode()

        logger.info("Waiting response from API...")
        # Отправляем запрос API с использованием закодированной строки
        start_time = time()
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
            timeout=3600
        )
        end_time = time()
        logger.info(f"Total processing time: {get_time(end_time - start_time)}")

        logger.info(f"Response status code: {response.status_code}")
        if not response.ok:
            context.abort(grpc.StatusCode.UNAVAILABLE, "HuggingFace API unavailable")

        data = response.json()["data"]
        logger.info(f"API response size: {len(data[0])}")
        return data[0]
