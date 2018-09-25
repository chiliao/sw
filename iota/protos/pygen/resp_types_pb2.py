# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: resp_types.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from github.com.gogo.protobuf.gogoproto import gogo_pb2 as github_dot_com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2
from github.com.pensando.sw.venice.utils.apigen.annotations import includes_pb2 as github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_includes__pb2
try:
  google_dot_api_dot_annotations__pb2 = github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_includes__pb2.google_dot_api_dot_annotations__pb2
except AttributeError:
  google_dot_api_dot_annotations__pb2 = github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_includes__pb2.google.api.annotations_pb2
try:
  github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_venice__pb2 = github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_includes__pb2.github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_venice__pb2
except AttributeError:
  github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_venice__pb2 = github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_includes__pb2.github.com.pensando.sw.venice.utils.apigen.annotations.venice_pb2
try:
  github_dot_com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2 = github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_includes__pb2.github_dot_com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2
except AttributeError:
  github_dot_com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2 = github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_includes__pb2.github.com.gogo.protobuf.gogoproto.gogo_pb2
import common_pb2 as common__pb2
try:
  github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_includes__pb2 = common__pb2.github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_includes__pb2
except AttributeError:
  github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_includes__pb2 = common__pb2.github.com.pensando.sw.venice.utils.apigen.annotations.includes_pb2
try:
  google_dot_api_dot_annotations__pb2 = common__pb2.google_dot_api_dot_annotations__pb2
except AttributeError:
  google_dot_api_dot_annotations__pb2 = common__pb2.google.api.annotations_pb2
try:
  github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_venice__pb2 = common__pb2.github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_venice__pb2
except AttributeError:
  github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_venice__pb2 = common__pb2.github.com.pensando.sw.venice.utils.apigen.annotations.venice_pb2
try:
  github_dot_com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2 = common__pb2.github_dot_com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2
except AttributeError:
  github_dot_com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2 = common__pb2.github.com.gogo.protobuf.gogoproto.gogo_pb2

from github.com.pensando.sw.venice.utils.apigen.annotations.includes_pb2 import *

DESCRIPTOR = _descriptor.FileDescriptor(
  name='resp_types.proto',
  package='iotamodel',
  syntax='proto3',
  serialized_pb=_b('\n\x10resp_types.proto\x12\tiotamodel\x1a\x1cgoogle/api/annotations.proto\x1a-github.com/gogo/protobuf/gogoproto/gogo.proto\x1a\x45github.com/pensando/sw/venice/utils/apigen/annotations/includes.proto\x1a\x0c\x63ommon.proto\"\xd2\x01\n\x0fIotaAPIResponse\x12X\n\tAPIStatus\x18\x01 \x01(\tBE\xaa\x86\x19)StrEnum(IotaAPIResponse.APIResponseTypee)\xea\xde\x1f\x14\x61pi-status,omitempty\"e\n\x0f\x41PIResponseType\x12\x11\n\rAPI_STATUS_OK\x10\x00\x12\x13\n\x0f\x41PI_BAD_REQUEST\x10\x01\x12\x14\n\x10\x41PI_SERVER_ERROR\x10\x02\x12\x14\n\x10\x41PI_AUTH_FAILURE\x10\x03\"T\n\x15\x43lusterHealthResponse\x12;\n\x06Status\x18\x01 \x03(\x0b\x32\x15.iotamodel.NodeStatusB\x14\xea\xde\x1f\x10status,omitempty\"\xfa\x01\n\nNodeStatus\x12Q\n\x13HealthCheckResponse\x18\x01 \x01(\x0b\x32\x1a.iotamodel.IotaAPIResponseB\x18\xea\xde\x1f\x14\x61pi-status,omitempty\x12J\n\x06Health\x18\x02 \x01(\tB:\xaa\x86\x19\"StrEnum(NodeStatus.HealthCodeType)\xea\xde\x1f\x10health,omitempty\"M\n\x0eHealthCodeType\x12\r\n\tHEALTH_OK\x10\x00\x12\x0f\n\x0bNAPLES_DOWN\x10\x01\x12\r\n\tNODE_DOWN\x10\x02\x12\x0c\n\x08\x41PP_DOWN\x10\x03\"|\n\x16InstantiateAppResponse\x12\x62\n\x16InstantiateAppResponse\x18\x01 \x01(\x0b\x32\x1a.iotamodel.IotaAPIResponseB&\xea\xde\x1f\"instantiate-app-response,omitempty\"~\n\x17InstantiateTopoResponse\x12\x63\n\x16InstantiateTopoReponse\x18\x01 \x01(\x0b\x32\x1a.iotamodel.IotaAPIResponseB\'\xea\xde\x1f#instantiate-topo-response,omitempty\"f\n\x0f\x41\x64\x64NodeResponse\x12S\n\x0e\x41\x64\x64NodeReponse\x18\x01 \x01(\x0b\x32\x1a.iotamodel.IotaAPIResponseB\x1f\xea\xde\x1f\x1b\x61\x64\x64-node-response,omitempty\"\xb2\x01\n\x17GeneratedConfigResponse\x12U\n\x17GeneratedConfigResponse\x18\x01 \x01(\x0b\x32\x1a.iotamodel.IotaAPIResponseB\x18\xea\xde\x1f\x14\x61pi-status,omitemtpy\x12@\n\nConfigInfo\x18\x02 \x01(\x0b\x32\x15.iotamodel.ConfigInfoB\x15\xea\xde\x1f\x11\x63onfigs,omitemtpy\"f\n\x12\x43onfigPushResponse\x12P\n\x12\x43onfigPushResponse\x18\x01 \x01(\x0b\x32\x1a.iotamodel.IotaAPIResponseB\x18\xea\xde\x1f\x14\x61pi-status,omitemtpy\"\xe2\x01\n\x12TriggerAppResponse\x12G\n\tAPIStatus\x18\x01 \x01(\x0b\x32\x1a.iotamodel.IotaAPIResponseB\x18\xea\xde\x1f\x14\x61pi-status,omitemtpy\x12+\n\tAppHandle\x18\x02 \x01(\tB\x18\xea\xde\x1f\x14\x61pp-handle,omitemtpy\x12+\n\tAppStdOut\x18\x03 \x01(\tB\x18\xea\xde\x1f\x14\x61pp-stdout,omitemtpy\x12)\n\x08\x45xitCode\x18\x04 \x01(\rB\x17\xea\xde\x1f\x13\x65xit-code,omitemtpyP\x02\x62\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,github_dot_com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2.DESCRIPTOR,github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_includes__pb2.DESCRIPTOR,common__pb2.DESCRIPTOR,],
  public_dependencies=[github_dot_com_dot_pensando_dot_sw_dot_venice_dot_utils_dot_apigen_dot_annotations_dot_includes__pb2.DESCRIPTOR,])



_IOTAAPIRESPONSE_APIRESPONSETYPE = _descriptor.EnumDescriptor(
  name='APIResponseType',
  full_name='iotamodel.IotaAPIResponse.APIResponseType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='API_STATUS_OK', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='API_BAD_REQUEST', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='API_SERVER_ERROR', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='API_AUTH_FAILURE', index=3, number=3,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=303,
  serialized_end=404,
)
_sym_db.RegisterEnumDescriptor(_IOTAAPIRESPONSE_APIRESPONSETYPE)

_NODESTATUS_HEALTHCODETYPE = _descriptor.EnumDescriptor(
  name='HealthCodeType',
  full_name='iotamodel.NodeStatus.HealthCodeType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='HEALTH_OK', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NAPLES_DOWN', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NODE_DOWN', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='APP_DOWN', index=3, number=3,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=666,
  serialized_end=743,
)
_sym_db.RegisterEnumDescriptor(_NODESTATUS_HEALTHCODETYPE)


_IOTAAPIRESPONSE = _descriptor.Descriptor(
  name='IotaAPIResponse',
  full_name='iotamodel.IotaAPIResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='APIStatus', full_name='iotamodel.IotaAPIResponse.APIStatus', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\252\206\031)StrEnum(IotaAPIResponse.APIResponseTypee)\352\336\037\024api-status,omitempty'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _IOTAAPIRESPONSE_APIRESPONSETYPE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=194,
  serialized_end=404,
)


_CLUSTERHEALTHRESPONSE = _descriptor.Descriptor(
  name='ClusterHealthResponse',
  full_name='iotamodel.ClusterHealthResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Status', full_name='iotamodel.ClusterHealthResponse.Status', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\020status,omitempty'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=406,
  serialized_end=490,
)


_NODESTATUS = _descriptor.Descriptor(
  name='NodeStatus',
  full_name='iotamodel.NodeStatus',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='HealthCheckResponse', full_name='iotamodel.NodeStatus.HealthCheckResponse', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\024api-status,omitempty'))),
    _descriptor.FieldDescriptor(
      name='Health', full_name='iotamodel.NodeStatus.Health', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\252\206\031\"StrEnum(NodeStatus.HealthCodeType)\352\336\037\020health,omitempty'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _NODESTATUS_HEALTHCODETYPE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=493,
  serialized_end=743,
)


_INSTANTIATEAPPRESPONSE = _descriptor.Descriptor(
  name='InstantiateAppResponse',
  full_name='iotamodel.InstantiateAppResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='InstantiateAppResponse', full_name='iotamodel.InstantiateAppResponse.InstantiateAppResponse', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\"instantiate-app-response,omitempty'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=745,
  serialized_end=869,
)


_INSTANTIATETOPORESPONSE = _descriptor.Descriptor(
  name='InstantiateTopoResponse',
  full_name='iotamodel.InstantiateTopoResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='InstantiateTopoReponse', full_name='iotamodel.InstantiateTopoResponse.InstantiateTopoReponse', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037#instantiate-topo-response,omitempty'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=871,
  serialized_end=997,
)


_ADDNODERESPONSE = _descriptor.Descriptor(
  name='AddNodeResponse',
  full_name='iotamodel.AddNodeResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='AddNodeReponse', full_name='iotamodel.AddNodeResponse.AddNodeReponse', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\033add-node-response,omitempty'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=999,
  serialized_end=1101,
)


_GENERATEDCONFIGRESPONSE = _descriptor.Descriptor(
  name='GeneratedConfigResponse',
  full_name='iotamodel.GeneratedConfigResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='GeneratedConfigResponse', full_name='iotamodel.GeneratedConfigResponse.GeneratedConfigResponse', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\024api-status,omitemtpy'))),
    _descriptor.FieldDescriptor(
      name='ConfigInfo', full_name='iotamodel.GeneratedConfigResponse.ConfigInfo', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\021configs,omitemtpy'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1104,
  serialized_end=1282,
)


_CONFIGPUSHRESPONSE = _descriptor.Descriptor(
  name='ConfigPushResponse',
  full_name='iotamodel.ConfigPushResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ConfigPushResponse', full_name='iotamodel.ConfigPushResponse.ConfigPushResponse', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\024api-status,omitemtpy'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1284,
  serialized_end=1386,
)


_TRIGGERAPPRESPONSE = _descriptor.Descriptor(
  name='TriggerAppResponse',
  full_name='iotamodel.TriggerAppResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='APIStatus', full_name='iotamodel.TriggerAppResponse.APIStatus', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\024api-status,omitemtpy'))),
    _descriptor.FieldDescriptor(
      name='AppHandle', full_name='iotamodel.TriggerAppResponse.AppHandle', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\024app-handle,omitemtpy'))),
    _descriptor.FieldDescriptor(
      name='AppStdOut', full_name='iotamodel.TriggerAppResponse.AppStdOut', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\024app-stdout,omitemtpy'))),
    _descriptor.FieldDescriptor(
      name='ExitCode', full_name='iotamodel.TriggerAppResponse.ExitCode', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\023exit-code,omitemtpy'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1389,
  serialized_end=1615,
)

_IOTAAPIRESPONSE_APIRESPONSETYPE.containing_type = _IOTAAPIRESPONSE
_CLUSTERHEALTHRESPONSE.fields_by_name['Status'].message_type = _NODESTATUS
_NODESTATUS.fields_by_name['HealthCheckResponse'].message_type = _IOTAAPIRESPONSE
_NODESTATUS_HEALTHCODETYPE.containing_type = _NODESTATUS
_INSTANTIATEAPPRESPONSE.fields_by_name['InstantiateAppResponse'].message_type = _IOTAAPIRESPONSE
_INSTANTIATETOPORESPONSE.fields_by_name['InstantiateTopoReponse'].message_type = _IOTAAPIRESPONSE
_ADDNODERESPONSE.fields_by_name['AddNodeReponse'].message_type = _IOTAAPIRESPONSE
_GENERATEDCONFIGRESPONSE.fields_by_name['GeneratedConfigResponse'].message_type = _IOTAAPIRESPONSE
_GENERATEDCONFIGRESPONSE.fields_by_name['ConfigInfo'].message_type = common__pb2._CONFIGINFO
_CONFIGPUSHRESPONSE.fields_by_name['ConfigPushResponse'].message_type = _IOTAAPIRESPONSE
_TRIGGERAPPRESPONSE.fields_by_name['APIStatus'].message_type = _IOTAAPIRESPONSE
DESCRIPTOR.message_types_by_name['IotaAPIResponse'] = _IOTAAPIRESPONSE
DESCRIPTOR.message_types_by_name['ClusterHealthResponse'] = _CLUSTERHEALTHRESPONSE
DESCRIPTOR.message_types_by_name['NodeStatus'] = _NODESTATUS
DESCRIPTOR.message_types_by_name['InstantiateAppResponse'] = _INSTANTIATEAPPRESPONSE
DESCRIPTOR.message_types_by_name['InstantiateTopoResponse'] = _INSTANTIATETOPORESPONSE
DESCRIPTOR.message_types_by_name['AddNodeResponse'] = _ADDNODERESPONSE
DESCRIPTOR.message_types_by_name['GeneratedConfigResponse'] = _GENERATEDCONFIGRESPONSE
DESCRIPTOR.message_types_by_name['ConfigPushResponse'] = _CONFIGPUSHRESPONSE
DESCRIPTOR.message_types_by_name['TriggerAppResponse'] = _TRIGGERAPPRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

IotaAPIResponse = _reflection.GeneratedProtocolMessageType('IotaAPIResponse', (_message.Message,), dict(
  DESCRIPTOR = _IOTAAPIRESPONSE,
  __module__ = 'resp_types_pb2'
  # @@protoc_insertion_point(class_scope:iotamodel.IotaAPIResponse)
  ))
_sym_db.RegisterMessage(IotaAPIResponse)

ClusterHealthResponse = _reflection.GeneratedProtocolMessageType('ClusterHealthResponse', (_message.Message,), dict(
  DESCRIPTOR = _CLUSTERHEALTHRESPONSE,
  __module__ = 'resp_types_pb2'
  # @@protoc_insertion_point(class_scope:iotamodel.ClusterHealthResponse)
  ))
_sym_db.RegisterMessage(ClusterHealthResponse)

NodeStatus = _reflection.GeneratedProtocolMessageType('NodeStatus', (_message.Message,), dict(
  DESCRIPTOR = _NODESTATUS,
  __module__ = 'resp_types_pb2'
  # @@protoc_insertion_point(class_scope:iotamodel.NodeStatus)
  ))
_sym_db.RegisterMessage(NodeStatus)

InstantiateAppResponse = _reflection.GeneratedProtocolMessageType('InstantiateAppResponse', (_message.Message,), dict(
  DESCRIPTOR = _INSTANTIATEAPPRESPONSE,
  __module__ = 'resp_types_pb2'
  # @@protoc_insertion_point(class_scope:iotamodel.InstantiateAppResponse)
  ))
_sym_db.RegisterMessage(InstantiateAppResponse)

InstantiateTopoResponse = _reflection.GeneratedProtocolMessageType('InstantiateTopoResponse', (_message.Message,), dict(
  DESCRIPTOR = _INSTANTIATETOPORESPONSE,
  __module__ = 'resp_types_pb2'
  # @@protoc_insertion_point(class_scope:iotamodel.InstantiateTopoResponse)
  ))
_sym_db.RegisterMessage(InstantiateTopoResponse)

AddNodeResponse = _reflection.GeneratedProtocolMessageType('AddNodeResponse', (_message.Message,), dict(
  DESCRIPTOR = _ADDNODERESPONSE,
  __module__ = 'resp_types_pb2'
  # @@protoc_insertion_point(class_scope:iotamodel.AddNodeResponse)
  ))
_sym_db.RegisterMessage(AddNodeResponse)

GeneratedConfigResponse = _reflection.GeneratedProtocolMessageType('GeneratedConfigResponse', (_message.Message,), dict(
  DESCRIPTOR = _GENERATEDCONFIGRESPONSE,
  __module__ = 'resp_types_pb2'
  # @@protoc_insertion_point(class_scope:iotamodel.GeneratedConfigResponse)
  ))
_sym_db.RegisterMessage(GeneratedConfigResponse)

ConfigPushResponse = _reflection.GeneratedProtocolMessageType('ConfigPushResponse', (_message.Message,), dict(
  DESCRIPTOR = _CONFIGPUSHRESPONSE,
  __module__ = 'resp_types_pb2'
  # @@protoc_insertion_point(class_scope:iotamodel.ConfigPushResponse)
  ))
_sym_db.RegisterMessage(ConfigPushResponse)

TriggerAppResponse = _reflection.GeneratedProtocolMessageType('TriggerAppResponse', (_message.Message,), dict(
  DESCRIPTOR = _TRIGGERAPPRESPONSE,
  __module__ = 'resp_types_pb2'
  # @@protoc_insertion_point(class_scope:iotamodel.TriggerAppResponse)
  ))
_sym_db.RegisterMessage(TriggerAppResponse)


_IOTAAPIRESPONSE.fields_by_name['APIStatus'].has_options = True
_IOTAAPIRESPONSE.fields_by_name['APIStatus']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\252\206\031)StrEnum(IotaAPIResponse.APIResponseTypee)\352\336\037\024api-status,omitempty'))
_CLUSTERHEALTHRESPONSE.fields_by_name['Status'].has_options = True
_CLUSTERHEALTHRESPONSE.fields_by_name['Status']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\020status,omitempty'))
_NODESTATUS.fields_by_name['HealthCheckResponse'].has_options = True
_NODESTATUS.fields_by_name['HealthCheckResponse']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\024api-status,omitempty'))
_NODESTATUS.fields_by_name['Health'].has_options = True
_NODESTATUS.fields_by_name['Health']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\252\206\031\"StrEnum(NodeStatus.HealthCodeType)\352\336\037\020health,omitempty'))
_INSTANTIATEAPPRESPONSE.fields_by_name['InstantiateAppResponse'].has_options = True
_INSTANTIATEAPPRESPONSE.fields_by_name['InstantiateAppResponse']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\"instantiate-app-response,omitempty'))
_INSTANTIATETOPORESPONSE.fields_by_name['InstantiateTopoReponse'].has_options = True
_INSTANTIATETOPORESPONSE.fields_by_name['InstantiateTopoReponse']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037#instantiate-topo-response,omitempty'))
_ADDNODERESPONSE.fields_by_name['AddNodeReponse'].has_options = True
_ADDNODERESPONSE.fields_by_name['AddNodeReponse']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\033add-node-response,omitempty'))
_GENERATEDCONFIGRESPONSE.fields_by_name['GeneratedConfigResponse'].has_options = True
_GENERATEDCONFIGRESPONSE.fields_by_name['GeneratedConfigResponse']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\024api-status,omitemtpy'))
_GENERATEDCONFIGRESPONSE.fields_by_name['ConfigInfo'].has_options = True
_GENERATEDCONFIGRESPONSE.fields_by_name['ConfigInfo']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\021configs,omitemtpy'))
_CONFIGPUSHRESPONSE.fields_by_name['ConfigPushResponse'].has_options = True
_CONFIGPUSHRESPONSE.fields_by_name['ConfigPushResponse']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\024api-status,omitemtpy'))
_TRIGGERAPPRESPONSE.fields_by_name['APIStatus'].has_options = True
_TRIGGERAPPRESPONSE.fields_by_name['APIStatus']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\024api-status,omitemtpy'))
_TRIGGERAPPRESPONSE.fields_by_name['AppHandle'].has_options = True
_TRIGGERAPPRESPONSE.fields_by_name['AppHandle']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\024app-handle,omitemtpy'))
_TRIGGERAPPRESPONSE.fields_by_name['AppStdOut'].has_options = True
_TRIGGERAPPRESPONSE.fields_by_name['AppStdOut']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\024app-stdout,omitemtpy'))
_TRIGGERAPPRESPONSE.fields_by_name['ExitCode'].has_options = True
_TRIGGERAPPRESPONSE.fields_by_name['ExitCode']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\352\336\037\023exit-code,omitemtpy'))
# @@protoc_insertion_point(module_scope)
