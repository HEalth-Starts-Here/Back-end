from dataclasses import dataclass


@dataclass
class TextSummParams:
    nlp_sum_type: str
    limit_sentence: int
