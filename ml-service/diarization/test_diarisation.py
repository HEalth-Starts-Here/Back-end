import io
import grpc

from api.diarisation_pb2 import DiarisationRequest
from api.diarisation_pb2_grpc import DiarisationStub


if __name__ == "__main__":
    channel = grpc.insecure_channel("127.0.0.1:50051")
    client = DiarisationStub(channel)

    path = "/mnt/e/Downloads/X3gh52xDth8_48.mp3"
    with open(path, "rb") as f:
        audio = io.BytesIO(f.read())

    request = DiarisationRequest(audio=audio.getvalue())
    response = client.transcribeAudio(request)
    print(response.text)
