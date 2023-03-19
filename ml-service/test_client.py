import io
import grpc

from PIL import Image

from protos.affected_area_pb2 import AffectedAreaRequest
from protos.affected_area_pb2_grpc import AffectedAreaStub


if __name__ == "__main__":
    channel = grpc.insecure_channel("localhost:50051")
    client = AffectedAreaStub(channel)
    img = Image.open("skin-lesions2.jpg")
    img_bytes = io.BytesIO()
    img.save(img_bytes, format=img.format)
    request = AffectedAreaRequest(image=img_bytes.getvalue())
    # print(img.tobytes())
    print(client.calculateArea(request))
