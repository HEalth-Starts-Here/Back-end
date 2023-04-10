import logging
import pprint
import sys
import click
import grpc

from concurrent import futures

from api.text_summarization_pb2_grpc import add_TextSummServicer_to_server
from entities.service_params import read_service_params
from service.text_summ_service import TextSummService


logger = logging.getLogger(__name__)
handler = logging.StreamHandler(sys.stdout)
logger.setLevel(logging.INFO)
logger.addHandler(handler)


@click.command()
@click.argument("config_path")
def serve(config_path: str):
    params = read_service_params(config_path)

    logger.info(f"Strarting service with params:\n {pprint.pformat(params)}")

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=params.max_workers))
    add_TextSummServicer_to_server(
        TextSummService(params.model_params), server
    )

    server.add_insecure_port(params.port)
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
