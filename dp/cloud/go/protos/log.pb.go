//
//Copyright 2022 The Magma Authors.
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
// 	protoc-gen-go v1.28.0
// 	protoc        v3.10.0
// source: dp/protos/log.proto

package protos

import (
	context "context"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ListLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId  string      `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	Filter     *LogFilter  `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	Pagination *Pagination `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListLogsRequest) Reset() {
	*x = ListLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dp_protos_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLogsRequest) ProtoMessage() {}

func (x *ListLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dp_protos_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLogsRequest.ProtoReflect.Descriptor instead.
func (*ListLogsRequest) Descriptor() ([]byte, []int) {
	return file_dp_protos_log_proto_rawDescGZIP(), []int{0}
}

func (x *ListLogsRequest) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *ListLogsRequest) GetFilter() *LogFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ListLogsRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type LogFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From                string               `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                  string               `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Name                string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SerialNumber        string               `protobuf:"bytes,4,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	FccId               string               `protobuf:"bytes,5,opt,name=fcc_id,json=fccId,proto3" json:"fcc_id,omitempty"`
	ResponseCode        *wrappers.Int64Value `protobuf:"bytes,6,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`
	BeginTimestampMilli *wrappers.Int64Value `protobuf:"bytes,7,opt,name=begin_timestamp_milli,json=beginTimestampMilli,proto3" json:"begin_timestamp_milli,omitempty"`
	EndTimestampMilli   *wrappers.Int64Value `protobuf:"bytes,8,opt,name=end_timestamp_milli,json=endTimestampMilli,proto3" json:"end_timestamp_milli,omitempty"`
}

func (x *LogFilter) Reset() {
	*x = LogFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dp_protos_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogFilter) ProtoMessage() {}

func (x *LogFilter) ProtoReflect() protoreflect.Message {
	mi := &file_dp_protos_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogFilter.ProtoReflect.Descriptor instead.
func (*LogFilter) Descriptor() ([]byte, []int) {
	return file_dp_protos_log_proto_rawDescGZIP(), []int{1}
}

func (x *LogFilter) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *LogFilter) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *LogFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LogFilter) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *LogFilter) GetFccId() string {
	if x != nil {
		return x.FccId
	}
	return ""
}

func (x *LogFilter) GetResponseCode() *wrappers.Int64Value {
	if x != nil {
		return x.ResponseCode
	}
	return nil
}

func (x *LogFilter) GetBeginTimestampMilli() *wrappers.Int64Value {
	if x != nil {
		return x.BeginTimestampMilli
	}
	return nil
}

func (x *LogFilter) GetEndTimestampMilli() *wrappers.Int64Value {
	if x != nil {
		return x.EndTimestampMilli
	}
	return nil
}

type ListLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*Log `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *ListLogsResponse) Reset() {
	*x = ListLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dp_protos_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLogsResponse) ProtoMessage() {}

func (x *ListLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dp_protos_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLogsResponse.ProtoReflect.Descriptor instead.
func (*ListLogsResponse) Descriptor() ([]byte, []int) {
	return file_dp_protos_log_proto_rawDescGZIP(), []int{2}
}

func (x *ListLogsResponse) GetLogs() []*Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From           string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To             string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Message        string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	SerialNumber   string `protobuf:"bytes,5,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	FccId          string `protobuf:"bytes,6,opt,name=fcc_id,json=fccId,proto3" json:"fcc_id,omitempty"`
	TimestampMilli int64  `protobuf:"varint,7,opt,name=timestamp_milli,json=timestampMilli,proto3" json:"timestamp_milli,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dp_protos_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_dp_protos_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_dp_protos_log_proto_rawDescGZIP(), []int{3}
}

func (x *Log) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Log) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Log) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Log) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Log) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *Log) GetFccId() string {
	if x != nil {
		return x.FccId
	}
	return ""
}

func (x *Log) GetTimestampMilli() int64 {
	if x != nil {
		return x.TimestampMilli
	}
	return 0
}

var File_dp_protos_log_proto protoreflect.FileDescriptor

var file_dp_protos_log_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x64, 0x70, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x64, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x2b,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x64, 0x70, 0x2e, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x64, 0x70, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xdf, 0x02, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x06,
	0x66, 0x63, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x63,
	0x63, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x4f, 0x0a, 0x15, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x13, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x12, 0x4b, 0x0a, 0x13, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x11, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x69,
	0x6c, 0x6c, 0x69, 0x22, 0x35, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x64, 0x70,
	0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x03, 0x4c,
	0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x66, 0x63, 0x63,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x63, 0x63, 0x49, 0x64,
	0x12, 0x27, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x69,
	0x6c, 0x6c, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x32, 0x51, 0x0a, 0x0a, 0x4c, 0x6f, 0x67,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4c,
	0x6f, 0x67, 0x73, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x64, 0x70, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x64, 0x70, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18,
	0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x64, 0x70, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dp_protos_log_proto_rawDescOnce sync.Once
	file_dp_protos_log_proto_rawDescData = file_dp_protos_log_proto_rawDesc
)

func file_dp_protos_log_proto_rawDescGZIP() []byte {
	file_dp_protos_log_proto_rawDescOnce.Do(func() {
		file_dp_protos_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_dp_protos_log_proto_rawDescData)
	})
	return file_dp_protos_log_proto_rawDescData
}

var file_dp_protos_log_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dp_protos_log_proto_goTypes = []interface{}{
	(*ListLogsRequest)(nil),     // 0: magma.dp.ListLogsRequest
	(*LogFilter)(nil),           // 1: magma.dp.LogFilter
	(*ListLogsResponse)(nil),    // 2: magma.dp.ListLogsResponse
	(*Log)(nil),                 // 3: magma.dp.Log
	(*Pagination)(nil),          // 4: magma.dp.Pagination
	(*wrappers.Int64Value)(nil), // 5: google.protobuf.Int64Value
}
var file_dp_protos_log_proto_depIdxs = []int32{
	1, // 0: magma.dp.ListLogsRequest.filter:type_name -> magma.dp.LogFilter
	4, // 1: magma.dp.ListLogsRequest.pagination:type_name -> magma.dp.Pagination
	5, // 2: magma.dp.LogFilter.response_code:type_name -> google.protobuf.Int64Value
	5, // 3: magma.dp.LogFilter.begin_timestamp_milli:type_name -> google.protobuf.Int64Value
	5, // 4: magma.dp.LogFilter.end_timestamp_milli:type_name -> google.protobuf.Int64Value
	3, // 5: magma.dp.ListLogsResponse.logs:type_name -> magma.dp.Log
	0, // 6: magma.dp.LogFetcher.ListLogs:input_type -> magma.dp.ListLogsRequest
	2, // 7: magma.dp.LogFetcher.ListLogs:output_type -> magma.dp.ListLogsResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_dp_protos_log_proto_init() }
func file_dp_protos_log_proto_init() {
	if File_dp_protos_log_proto != nil {
		return
	}
	file_dp_protos_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dp_protos_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLogsRequest); i {
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
		file_dp_protos_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogFilter); i {
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
		file_dp_protos_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLogsResponse); i {
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
		file_dp_protos_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_dp_protos_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dp_protos_log_proto_goTypes,
		DependencyIndexes: file_dp_protos_log_proto_depIdxs,
		MessageInfos:      file_dp_protos_log_proto_msgTypes,
	}.Build()
	File_dp_protos_log_proto = out.File
	file_dp_protos_log_proto_rawDesc = nil
	file_dp_protos_log_proto_goTypes = nil
	file_dp_protos_log_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogFetcherClient is the client API for LogFetcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogFetcherClient interface {
	ListLogs(ctx context.Context, in *ListLogsRequest, opts ...grpc.CallOption) (*ListLogsResponse, error)
}

type logFetcherClient struct {
	cc grpc.ClientConnInterface
}

func NewLogFetcherClient(cc grpc.ClientConnInterface) LogFetcherClient {
	return &logFetcherClient{cc}
}

func (c *logFetcherClient) ListLogs(ctx context.Context, in *ListLogsRequest, opts ...grpc.CallOption) (*ListLogsResponse, error) {
	out := new(ListLogsResponse)
	err := c.cc.Invoke(ctx, "/magma.dp.LogFetcher/ListLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogFetcherServer is the server API for LogFetcher service.
type LogFetcherServer interface {
	ListLogs(context.Context, *ListLogsRequest) (*ListLogsResponse, error)
}

// UnimplementedLogFetcherServer can be embedded to have forward compatible implementations.
type UnimplementedLogFetcherServer struct {
}

func (*UnimplementedLogFetcherServer) ListLogs(context.Context, *ListLogsRequest) (*ListLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLogs not implemented")
}

func RegisterLogFetcherServer(s *grpc.Server, srv LogFetcherServer) {
	s.RegisterService(&_LogFetcher_serviceDesc, srv)
}

func _LogFetcher_ListLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogFetcherServer).ListLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.dp.LogFetcher/ListLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogFetcherServer).ListLogs(ctx, req.(*ListLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogFetcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.dp.LogFetcher",
	HandlerType: (*LogFetcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLogs",
			Handler:    _LogFetcher_ListLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dp/protos/log.proto",
}
