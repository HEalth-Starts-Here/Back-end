from dataclasses import dataclass
from typing import Optional


@dataclass
class DiarisationParams:
    model_size: str
    num_speakers: int
    device: str
    use_api: Optional[bool]
