from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class IQARequest(_message.Message):
    __slots__ = ["image"]
    IMAGE_FIELD_NUMBER: _ClassVar[int]
    image: bytes
    def __init__(self, image: _Optional[bytes] = ...) -> None: ...

class IQAResponse(_message.Message):
    __slots__ = ["quality"]
    QUALITY_FIELD_NUMBER: _ClassVar[int]
    quality: bool
    def __init__(self, quality: bool = ...) -> None: ...
