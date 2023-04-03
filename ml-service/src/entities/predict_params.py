from dataclasses import dataclass


@dataclass
class PredictParams:
    model_path: str
    num_parts: int
    num_top_parts: int
    quality_threshold: float
