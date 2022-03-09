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
// source: orc8r/protos/redis.proto

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

type RedisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerializedMsg []byte `protobuf:"bytes,1,opt,name=serialized_msg,json=serializedMsg,proto3" json:"serialized_msg,omitempty"`
	Version       uint64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	IsGarbage     bool   `protobuf:"varint,3,opt,name=is_garbage,json=isGarbage,proto3" json:"is_garbage,omitempty"`
}

func (x *RedisState) Reset() {
	*x = RedisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_redis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisState) ProtoMessage() {}

func (x *RedisState) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_redis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisState.ProtoReflect.Descriptor instead.
func (*RedisState) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_redis_proto_rawDescGZIP(), []int{0}
}

func (x *RedisState) GetSerializedMsg() []byte {
	if x != nil {
		return x.SerializedMsg
	}
	return nil
}

func (x *RedisState) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *RedisState) GetIsGarbage() bool {
	if x != nil {
		return x.IsGarbage
	}
	return false
}

var File_orc8r_protos_redis_proto protoreflect.FileDescriptor

var file_orc8r_protos_redis_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72,
	0x65, 0x64, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x22, 0x6c, 0x0a, 0x0a, 0x52, 0x65, 0x64, 0x69, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x67, 0x61, 0x72,
	0x62, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x47, 0x61,
	0x72, 0x62, 0x61, 0x67, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6f,
	0x72, 0x63, 0x38, 0x72, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orc8r_protos_redis_proto_rawDescOnce sync.Once
	file_orc8r_protos_redis_proto_rawDescData = file_orc8r_protos_redis_proto_rawDesc
)

func file_orc8r_protos_redis_proto_rawDescGZIP() []byte {
	file_orc8r_protos_redis_proto_rawDescOnce.Do(func() {
		file_orc8r_protos_redis_proto_rawDescData = protoimpl.X.CompressGZIP(file_orc8r_protos_redis_proto_rawDescData)
	})
	return file_orc8r_protos_redis_proto_rawDescData
}

var file_orc8r_protos_redis_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_orc8r_protos_redis_proto_goTypes = []interface{}{
	(*RedisState)(nil), // 0: magma.orc8r.RedisState
}
var file_orc8r_protos_redis_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orc8r_protos_redis_proto_init() }
func file_orc8r_protos_redis_proto_init() {
	if File_orc8r_protos_redis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orc8r_protos_redis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisState); i {
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
			RawDescriptor: file_orc8r_protos_redis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orc8r_protos_redis_proto_goTypes,
		DependencyIndexes: file_orc8r_protos_redis_proto_depIdxs,
		MessageInfos:      file_orc8r_protos_redis_proto_msgTypes,
	}.Build()
	File_orc8r_protos_redis_proto = out.File
	file_orc8r_protos_redis_proto_rawDesc = nil
	file_orc8r_protos_redis_proto_goTypes = nil
	file_orc8r_protos_redis_proto_depIdxs = nil
}
