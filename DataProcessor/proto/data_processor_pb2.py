# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: data_processor.proto
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
    'data_processor.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from proto import common_pb2 as common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14\x64\x61ta_processor.proto\x12\rDataProcessor\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x0c\x63ommon.proto\"^\n\x04\x44\x61ta\x12\x11\n\tsource_id\x18\x01 \x01(\x05\x12\x11\n\tdata_json\x18\x02 \x01(\t\x12\x30\n\x0c\x61rrival_time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp2G\n\rDataProcessor\x12\x36\n\x0bProcessData\x12\x13.DataProcessor.Data\x1a\x0e.common.Status\"\x00(\x01\x42\x15Z\x13/pkg/internal/protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'data_processor_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\023/pkg/internal/proto'
  _globals['_DATA']._serialized_start=86
  _globals['_DATA']._serialized_end=180
  _globals['_DATAPROCESSOR']._serialized_start=182
  _globals['_DATAPROCESSOR']._serialized_end=253
# @@protoc_insertion_point(module_scope)
