# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: model_orchestrator.proto
# Protobuf Python Version: 5.28.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
    1,
    '',
    'model_orchestrator.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from proto import common_pb2 as common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18model_orchestrator.proto\x12\x11ModelOrchestrator\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x0c\x63ommon.proto\"y\n\rProcessedData\x12\x11\n\tsource_id\x18\x01 \x01(\x05\x12\x10\n\x08model_id\x18\x02 \x01(\x05\x12\x11\n\tdata_json\x18\x03 \x01(\t\x12\x30\n\x0c\x61rrival_time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp2U\n\x11ModelOrchestrator\x12@\n\x08SendData\x12 .ModelOrchestrator.ProcessedData\x1a\x0e.common.Status\"\x00(\x01\x42\x15Z\x13/pkg/internal/protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'model_orchestrator_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\023/pkg/internal/proto'
  _globals['_PROCESSEDDATA']._serialized_start=94
  _globals['_PROCESSEDDATA']._serialized_end=215
  _globals['_MODELORCHESTRATOR']._serialized_start=217
  _globals['_MODELORCHESTRATOR']._serialized_end=302
# @@protoc_insertion_point(module_scope)