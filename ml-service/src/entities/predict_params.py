from dataclasses import dataclass
from typing import Optional

from src.entities.models_params import DiarisationParams, IQAModelParams, TextSummParams


@dataclass
class DownloadParams:
    file_url: str
    path_to_save: str


@dataclass
class PredictParams:
    model_path: str
    model_params: IQAModelParams | TextSummParams | DiarisationParams
    download_params: Optional[DownloadParams]
