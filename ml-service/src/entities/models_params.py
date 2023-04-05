from dataclasses import dataclass


@dataclass
class IQAModelParams:
    num_parts: int
    num_top_parts: int
    quality_threshold: float


@dataclass
class TextSummParams:
    nlp_sum_type: str
    limit_sentence: int