import pprint
import sys
import click
import grpc

from concurrent import futures


from protos.iqa_pb2_grpc import add_IQAServicer_to_server
from src.services.iqa_service import IQAService
from src.entities.service_params import read_service_params
from src.entities.logger import setup_default_logger


logger = setup_default_logger("service_logs", sys.stdout)


@click.command()
@click.argument("config_path")
def serve(config_path: str):
    params = read_service_params(config_path)

    logger.info(f"Strarting service with params:\n {pprint.pformat(params)}")

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=params.max_workers))
    if params.service == "IQA":
        add_IQAServicer_to_server(IQAService(params.predict_params), server)
    else:
        raise ValueError(f"Unknown service: {params.service}")

    server.add_insecure_port(params.port)
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
