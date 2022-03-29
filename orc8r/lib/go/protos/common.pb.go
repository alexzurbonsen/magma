//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.10.0
// source: orc8r/protos/common.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// --------------------------------------------------------------------------
// Logging levels
// --------------------------------------------------------------------------
type LogLevel int32

const (
	LogLevel_DEBUG   LogLevel = 0
	LogLevel_INFO    LogLevel = 1
	LogLevel_WARNING LogLevel = 2
	LogLevel_ERROR   LogLevel = 3
	LogLevel_FATAL   LogLevel = 4
)

// Enum value maps for LogLevel.
var (
	LogLevel_name = map[int32]string{
		0: "DEBUG",
		1: "INFO",
		2: "WARNING",
		3: "ERROR",
		4: "FATAL",
	}
	LogLevel_value = map[string]int32{
		"DEBUG":   0,
		"INFO":    1,
		"WARNING": 2,
		"ERROR":   3,
		"FATAL":   4,
	}
)

func (x LogLevel) Enum() *LogLevel {
	p := new(LogLevel)
	*p = x
	return p
}

func (x LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_orc8r_protos_common_proto_enumTypes[0].Descriptor()
}

func (LogLevel) Type() protoreflect.EnumType {
	return &file_orc8r_protos_common_proto_enumTypes[0]
}

func (x LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogLevel.Descriptor instead.
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_orc8r_protos_common_proto_rawDescGZIP(), []int{0}
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_common_proto_rawDescGZIP(), []int{0}
}

// -------------------------------------------------------------------------------
// Bytes is a special message type used to marshal & unmarshal unknown types as is
// -------------------------------------------------------------------------------
type Bytes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []byte `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Bytes) Reset() {
	*x = Bytes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bytes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bytes) ProtoMessage() {}

func (x *Bytes) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bytes.ProtoReflect.Descriptor instead.
func (*Bytes) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_common_proto_rawDescGZIP(), []int{1}
}

func (x *Bytes) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

// --------------------------------------------------------------------------
// NetworkID uniquely identifies the network
// --------------------------------------------------------------------------
type NetworkID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NetworkID) Reset() {
	*x = NetworkID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkID) ProtoMessage() {}

func (x *NetworkID) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkID.ProtoReflect.Descriptor instead.
func (*NetworkID) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_common_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// --------------------------------------------------------------------------
// IDList is a generic definition of an array of IDs (network, gateway, etc.)
// --------------------------------------------------------------------------
type IDList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IDList) Reset() {
	*x = IDList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDList) ProtoMessage() {}

func (x *IDList) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDList.ProtoReflect.Descriptor instead.
func (*IDList) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_common_proto_rawDescGZIP(), []int{3}
}

func (x *IDList) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_orc8r_protos_common_proto protoreflect.FileDescriptor

var file_orc8r_protos_common_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64,
	0x22, 0x19, 0x0a, 0x05, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x1b, 0x0a, 0x09, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x06, 0x49, 0x44, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x2a, 0x42, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49,
	0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x09, 0x0a,
	0x05, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x04, 0x42, 0x1b, 0x5a, 0x19, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2f, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orc8r_protos_common_proto_rawDescOnce sync.Once
	file_orc8r_protos_common_proto_rawDescData = file_orc8r_protos_common_proto_rawDesc
)

func file_orc8r_protos_common_proto_rawDescGZIP() []byte {
	file_orc8r_protos_common_proto_rawDescOnce.Do(func() {
		file_orc8r_protos_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_orc8r_protos_common_proto_rawDescData)
	})
	return file_orc8r_protos_common_proto_rawDescData
}

var file_orc8r_protos_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_orc8r_protos_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_orc8r_protos_common_proto_goTypes = []interface{}{
	(LogLevel)(0),     // 0: magma.orc8r.LogLevel
	(*Void)(nil),      // 1: magma.orc8r.Void
	(*Bytes)(nil),     // 2: magma.orc8r.Bytes
	(*NetworkID)(nil), // 3: magma.orc8r.NetworkID
	(*IDList)(nil),    // 4: magma.orc8r.IDList
}
var file_orc8r_protos_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orc8r_protos_common_proto_init() }
func file_orc8r_protos_common_proto_init() {
	if File_orc8r_protos_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orc8r_protos_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orc8r_protos_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bytes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orc8r_protos_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orc8r_protos_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orc8r_protos_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orc8r_protos_common_proto_goTypes,
		DependencyIndexes: file_orc8r_protos_common_proto_depIdxs,
		EnumInfos:         file_orc8r_protos_common_proto_enumTypes,
		MessageInfos:      file_orc8r_protos_common_proto_msgTypes,
	}.Build()
	File_orc8r_protos_common_proto = out.File
	file_orc8r_protos_common_proto_rawDesc = nil
	file_orc8r_protos_common_proto_goTypes = nil
	file_orc8r_protos_common_proto_depIdxs = nil
}
