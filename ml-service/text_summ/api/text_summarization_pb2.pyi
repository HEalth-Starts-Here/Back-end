from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class TextSummRequest(_message.Message):
    __slots__ = ["text"]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    text: str
    def __init__(self, text: _Optional[str] = ...) -> None: ...

class TextSummResponse(_message.Message):
    __slots__ = ["summarized_text"]
    SUMMARIZED_TEXT_FIELD_NUMBER: _ClassVar[int]
    summarized_text: str
    def __init__(self, summarized_text: _Optional[str] = ...) -> None: ...
