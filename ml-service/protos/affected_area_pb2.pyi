from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class AffectedAreaRequest(_message.Message):
    __slots__ = ["image"]
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    image: bytes
    def __init__(self, image: _Optional[bytes] = ...) -> None: ...

class AffectedAreaResponse(_message.Message):
    __slots__ = ["area"]
    AREA_FIELD_NUMBER: _ClassVar[int]
    area: int
    def __init__(self, area: _Optional[int] = ...) -> None: ...
