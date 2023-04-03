import io
import grpc

from PIL import Image

from protos.iqa_pb2 import IQARequest
from protos.iqa_pb2_grpc import IQAStub


if __name__ == "__main__":
    channel = grpc.insecure_channel("localhost:50051")
    client = IQAStub(channel)

    img = Image.open("skin-lesions2.jpg")
    img_bytes = io.BytesIO()
    img.save(img_bytes, format=img.format)

    request = IQARequest(image=img_bytes.getvalue())
    response = client.calculateQuality(request)
    print(f"Image quality: {response.quality}")
