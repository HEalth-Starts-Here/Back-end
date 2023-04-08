import grpc

from protos.text_summarization_pb2 import TextSummRequest
from protos.text_summarization_pb2_grpc import TextSummStub


if __name__ == "__main__":
    channel = grpc.insecure_channel("127.0.0.1:50051")
    client = TextSummStub(channel)

    with open("tests/data/text.txt", "r") as file:
        text = file.read()

    request = TextSummRequest(text=text)
    response = client.summarizeText(request)
    print(f"Summarized text: {response.summarized_text}")
