from dataclasses import dataclass
from typing import Optional


@dataclass
class IQAModelParams:
    num_parts: int
    num_top_parts: int
    quality_threshold: float


@dataclass
class TextSummParams:
    nlp_sum_type: str
    limit_sentence: int


@dataclass
class DiarisationParams:
    model_size: str
    num_speakers: int
    use_api: Optional[bool]
