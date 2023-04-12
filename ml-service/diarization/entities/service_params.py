import yaml

from dataclasses import dataclass
from marshmallow_dataclass import class_schema

from entities.models_params import DiarisationParams


@dataclass
class ServiceParams:
    port: str
    model_params: DiarisationParams
    max_receive_message_length: int = 4
    max_workers: int = 10


ServiceParamsSchema = class_schema(ServiceParams)


def read_service_params(config_path: str) -> ServiceParams:
    with open(config_path, "r") as input_stream:
        schema = ServiceParamsSchema()
        return schema.load(yaml.safe_load(input_stream))
