from dataclasses import dataclass
from typing import Optional


@dataclass
class DownloadParams:
    file_url: str
    path_to_save: str


@dataclass
class PredictParams:
    model_path: str
    num_parts: int
    num_top_parts: int
    quality_threshold: float
    download_params: Optional[DownloadParams]
