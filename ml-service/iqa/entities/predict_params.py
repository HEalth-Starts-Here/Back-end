from dataclasses import dataclass
from typing import Optional

from entities.models_params import IQAModelParams


@dataclass
class DownloadParams:
    file_url: str
    path_to_save: str


@dataclass
class PredictParams:
    model_path: str
    model_params: IQAModelParams
    download_params: Optional[DownloadParams]
