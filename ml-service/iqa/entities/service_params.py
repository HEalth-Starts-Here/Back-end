import yaml

from dataclasses import dataclass
from marshmallow_dataclass import class_schema

from entities.predict_params import PredictParams


@dataclass
class ServiceParams:
    port: str
    predict_params: PredictParams
    max_workers: int = 10


ServiceParamsSchema = class_schema(ServiceParams)


def read_service_params(config_path: str) -> ServiceParams:
    with open(config_path, "r") as input_stream:
        schema = ServiceParamsSchema()
        return schema.load(yaml.safe_load(input_stream))
