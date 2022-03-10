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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.10.0
// source: orc8r/protos/ctraced.proto

package protos

import (
	context "context"
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

type StartTraceRequest_TraceType int32

const (
	StartTraceRequest_ALL        StartTraceRequest_TraceType = 0
	StartTraceRequest_SUBSCRIBER StartTraceRequest_TraceType = 1
	StartTraceRequest_PROTOCOL   StartTraceRequest_TraceType = 2
	StartTraceRequest_INTERFACE  StartTraceRequest_TraceType = 3
	StartTraceRequest_CUSTOM     StartTraceRequest_TraceType = 4
)

// Enum value maps for StartTraceRequest_TraceType.
var (
	StartTraceRequest_TraceType_name = map[int32]string{
		0: "ALL",
		1: "SUBSCRIBER",
		2: "PROTOCOL",
		3: "INTERFACE",
		4: "CUSTOM",
	}
	StartTraceRequest_TraceType_value = map[string]int32{
		"ALL":        0,
		"SUBSCRIBER": 1,
		"PROTOCOL":   2,
		"INTERFACE":  3,
		"CUSTOM":     4,
	}
)

func (x StartTraceRequest_TraceType) Enum() *StartTraceRequest_TraceType {
	p := new(StartTraceRequest_TraceType)
	*p = x
	return p
}

func (x StartTraceRequest_TraceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StartTraceRequest_TraceType) Descriptor() protoreflect.EnumDescriptor {
	return file_orc8r_protos_ctraced_proto_enumTypes[0].Descriptor()
}

func (StartTraceRequest_TraceType) Type() protoreflect.EnumType {
	return &file_orc8r_protos_ctraced_proto_enumTypes[0]
}

func (x StartTraceRequest_TraceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StartTraceRequest_TraceType.Descriptor instead.
func (StartTraceRequest_TraceType) EnumDescriptor() ([]byte, []int) {
	return file_orc8r_protos_ctraced_proto_rawDescGZIP(), []int{0, 0}
}

// SPECIFIED ONLY IF trace_type is PROTOCOL
type StartTraceRequest_ProtocolName int32

const (
	StartTraceRequest_SCTP     StartTraceRequest_ProtocolName = 0
	StartTraceRequest_DIAMETER StartTraceRequest_ProtocolName = 1
)

// Enum value maps for StartTraceRequest_ProtocolName.
var (
	StartTraceRequest_ProtocolName_name = map[int32]string{
		0: "SCTP",
		1: "DIAMETER",
	}
	StartTraceRequest_ProtocolName_value = map[string]int32{
		"SCTP":     0,
		"DIAMETER": 1,
	}
)

func (x StartTraceRequest_ProtocolName) Enum() *StartTraceRequest_ProtocolName {
	p := new(StartTraceRequest_ProtocolName)
	*p = x
	return p
}

func (x StartTraceRequest_ProtocolName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StartTraceRequest_ProtocolName) Descriptor() protoreflect.EnumDescriptor {
	return file_orc8r_protos_ctraced_proto_enumTypes[1].Descriptor()
}

func (StartTraceRequest_ProtocolName) Type() protoreflect.EnumType {
	return &file_orc8r_protos_ctraced_proto_enumTypes[1]
}

func (x StartTraceRequest_ProtocolName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StartTraceRequest_ProtocolName.Descriptor instead.
func (StartTraceRequest_ProtocolName) EnumDescriptor() ([]byte, []int) {
	return file_orc8r_protos_ctraced_proto_rawDescGZIP(), []int{0, 1}
}

type StartTraceRequest_InterfaceName int32

const (
	StartTraceRequest_S1AP StartTraceRequest_InterfaceName = 0
	StartTraceRequest_GX   StartTraceRequest_InterfaceName = 1
	StartTraceRequest_GT   StartTraceRequest_InterfaceName = 2
)

// Enum value maps for StartTraceRequest_InterfaceName.
var (
	StartTraceRequest_InterfaceName_name = map[int32]string{
		0: "S1AP",
		1: "GX",
		2: "GT",
	}
	StartTraceRequest_InterfaceName_value = map[string]int32{
		"S1AP": 0,
		"GX":   1,
		"GT":   2,
	}
)

func (x StartTraceRequest_InterfaceName) Enum() *StartTraceRequest_InterfaceName {
	p := new(StartTraceRequest_InterfaceName)
	*p = x
	return p
}

func (x StartTraceRequest_InterfaceName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StartTraceRequest_InterfaceName) Descriptor() protoreflect.EnumDescriptor {
	return file_orc8r_protos_ctraced_proto_enumTypes[2].Descriptor()
}

func (StartTraceRequest_InterfaceName) Type() protoreflect.EnumType {
	return &file_orc8r_protos_ctraced_proto_enumTypes[2]
}

func (x StartTraceRequest_InterfaceName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StartTraceRequest_InterfaceName.Descriptor instead.
func (StartTraceRequest_InterfaceName) EnumDescriptor() ([]byte, []int) {
	return file_orc8r_protos_ctraced_proto_rawDescGZIP(), []int{0, 2}
}

// StartTraceRequest specifies options for a call trace started on a gateway
type StartTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId   string                      `protobuf:"bytes,8,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	TraceType StartTraceRequest_TraceType `protobuf:"varint,1,opt,name=trace_type,json=traceType,proto3,enum=magma.orc8r.StartTraceRequest_TraceType" json:"trace_type,omitempty"`
	// SPECIFIED ONLY IF trace_type is SUBSCRIBER
	Imsi     string                         `protobuf:"bytes,2,opt,name=imsi,proto3" json:"imsi,omitempty"` // Include prefix 'IMSI'
	Protocol StartTraceRequest_ProtocolName `protobuf:"varint,3,opt,name=protocol,proto3,enum=magma.orc8r.StartTraceRequest_ProtocolName" json:"protocol,omitempty"`
	// SPECIFIED IF trace_type is INTERFACE
	Interface StartTraceRequest_InterfaceName `protobuf:"varint,4,opt,name=interface,proto3,enum=magma.orc8r.StartTraceRequest_InterfaceName" json:"interface,omitempty"`
	// SPECIFIED FOR ALL
	// After the timeout runs out, the call trace stops automatically.
	Timeout uint32 `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"` // Specified in seconds
	// SPECIFIED ONLY IF trace_type is CUSTOM
	// TShark capture filters to be used for call tracing, based on BPF syntax
	// Custom traces are performed using TShark, and capture filters specified
	// are equivalent to running tshark with the -f option
	// Example:
	//    capture_filters = "tcp and port 80"
	//    Runs `tshark -f 'tcp and port 80'`
	CaptureFilters string `protobuf:"bytes,6,opt,name=capture_filters,json=captureFilters,proto3" json:"capture_filters,omitempty"`
	// SPECIFIED ONLY IF trace_type is CUSTOM
	// TShark display filters to be used for call tracing
	// Custom traces are performed using TShark, and display filters specified
	// are equivalent to running tshark with the -Y option
	// Example:
	//    display_filters = "ip.addr == 10.0.0.1"
	//    Runs `tshark -Y 'ip.addr == 10.0.0.1'`
	// NOTE:
	//    While TShark is used to capture traffic during call tracing with
	//    capture filters, when TShark is stopped, the resultant capture is
	//    passed through the display filters and saved again, so the final
	//    capture received has been processed by both the capture and display
	//    filters.
	DisplayFilters string `protobuf:"bytes,7,opt,name=display_filters,json=displayFilters,proto3" json:"display_filters,omitempty"`
}

func (x *StartTraceRequest) Reset() {
	*x = StartTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_ctraced_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTraceRequest) ProtoMessage() {}

func (x *StartTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_ctraced_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTraceRequest.ProtoReflect.Descriptor instead.
func (*StartTraceRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_ctraced_proto_rawDescGZIP(), []int{0}
}

func (x *StartTraceRequest) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *StartTraceRequest) GetTraceType() StartTraceRequest_TraceType {
	if x != nil {
		return x.TraceType
	}
	return StartTraceRequest_ALL
}

func (x *StartTraceRequest) GetImsi() string {
	if x != nil {
		return x.Imsi
	}
	return ""
}

func (x *StartTraceRequest) GetProtocol() StartTraceRequest_ProtocolName {
	if x != nil {
		return x.Protocol
	}
	return StartTraceRequest_SCTP
}

func (x *StartTraceRequest) GetInterface() StartTraceRequest_InterfaceName {
	if x != nil {
		return x.Interface
	}
	return StartTraceRequest_S1AP
}

func (x *StartTraceRequest) GetTimeout() uint32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *StartTraceRequest) GetCaptureFilters() string {
	if x != nil {
		return x.CaptureFilters
	}
	return ""
}

func (x *StartTraceRequest) GetDisplayFilters() string {
	if x != nil {
		return x.DisplayFilters
	}
	return ""
}

type StartTraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // May fail due to an existing tracing session
}

func (x *StartTraceResponse) Reset() {
	*x = StartTraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_ctraced_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTraceResponse) ProtoMessage() {}

func (x *StartTraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_ctraced_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTraceResponse.ProtoReflect.Descriptor instead.
func (*StartTraceResponse) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_ctraced_proto_rawDescGZIP(), []int{1}
}

func (x *StartTraceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type EndTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId string `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
}

func (x *EndTraceRequest) Reset() {
	*x = EndTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_ctraced_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndTraceRequest) ProtoMessage() {}

func (x *EndTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_ctraced_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndTraceRequest.ProtoReflect.Descriptor instead.
func (*EndTraceRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_ctraced_proto_rawDescGZIP(), []int{2}
}

func (x *EndTraceRequest) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

type EndTraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success      bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`                              // May fail due to no existing tracing session
	TraceContent []byte `protobuf:"bytes,2,opt,name=trace_content,json=traceContent,proto3" json:"trace_content,omitempty"` // Max size of 4MB
}

func (x *EndTraceResponse) Reset() {
	*x = EndTraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_ctraced_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndTraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndTraceResponse) ProtoMessage() {}

func (x *EndTraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_ctraced_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndTraceResponse.ProtoReflect.Descriptor instead.
func (*EndTraceResponse) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_ctraced_proto_rawDescGZIP(), []int{3}
}

func (x *EndTraceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *EndTraceResponse) GetTraceContent() []byte {
	if x != nil {
		return x.TraceContent
	}
	return nil
}

type ReportEndedTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId      string `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Success      bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	TraceContent []byte `protobuf:"bytes,3,opt,name=trace_content,json=traceContent,proto3" json:"trace_content,omitempty"`
}

func (x *ReportEndedTraceRequest) Reset() {
	*x = ReportEndedTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_ctraced_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportEndedTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportEndedTraceRequest) ProtoMessage() {}

func (x *ReportEndedTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_ctraced_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportEndedTraceRequest.ProtoReflect.Descriptor instead.
func (*ReportEndedTraceRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_ctraced_proto_rawDescGZIP(), []int{4}
}

func (x *ReportEndedTraceRequest) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *ReportEndedTraceRequest) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ReportEndedTraceRequest) GetTraceContent() []byte {
	if x != nil {
		return x.TraceContent
	}
	return nil
}

type ReportEndedTraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportEndedTraceResponse) Reset() {
	*x = ReportEndedTraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_ctraced_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportEndedTraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportEndedTraceResponse) ProtoMessage() {}

func (x *ReportEndedTraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_ctraced_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportEndedTraceResponse.ProtoReflect.Descriptor instead.
func (*ReportEndedTraceResponse) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_ctraced_proto_rawDescGZIP(), []int{5}
}

var File_orc8r_protos_ctraced_proto protoreflect.FileDescriptor

var file_orc8r_protos_ctraced_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x61,
	0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x22, 0xae, 0x04, 0x0a, 0x11, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x0a, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28,
	0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x12, 0x47, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x4a, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38,
	0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72,
	0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0x4d, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x46, 0x41, 0x43, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x43,
	0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x04, 0x22, 0x26, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x43, 0x54, 0x50, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x01, 0x22,
	0x29, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x53, 0x31, 0x41, 0x50, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x58,
	0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x54, 0x10, 0x02, 0x22, 0x2e, 0x0a, 0x12, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2c, 0x0a, 0x0f, 0x45, 0x6e,
	0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x10, 0x45, 0x6e, 0x64, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x73, 0x0a, 0x17, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb6, 0x01, 0x0a,
	0x10, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x53, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38,
	0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38,
	0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0c, 0x45, 0x6e, 0x64, 0x43, 0x61, 0x6c,
	0x6c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x1c, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f,
	0x72, 0x63, 0x38, 0x72, 0x2e, 0x45, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63,
	0x38, 0x72, 0x2e, 0x45, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x7c, 0x0a, 0x13, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x65, 0x0a, 0x14,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x12, 0x24, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63,
	0x38, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x61, 0x67,
	0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45,
	0x6e, 0x64, 0x65, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6f, 0x72, 0x63,
	0x38, 0x72, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orc8r_protos_ctraced_proto_rawDescOnce sync.Once
	file_orc8r_protos_ctraced_proto_rawDescData = file_orc8r_protos_ctraced_proto_rawDesc
)

func file_orc8r_protos_ctraced_proto_rawDescGZIP() []byte {
	file_orc8r_protos_ctraced_proto_rawDescOnce.Do(func() {
		file_orc8r_protos_ctraced_proto_rawDescData = protoimpl.X.CompressGZIP(file_orc8r_protos_ctraced_proto_rawDescData)
	})
	return file_orc8r_protos_ctraced_proto_rawDescData
}

var file_orc8r_protos_ctraced_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_orc8r_protos_ctraced_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_orc8r_protos_ctraced_proto_goTypes = []interface{}{
	(StartTraceRequest_TraceType)(0),     // 0: magma.orc8r.StartTraceRequest.TraceType
	(StartTraceRequest_ProtocolName)(0),  // 1: magma.orc8r.StartTraceRequest.ProtocolName
	(StartTraceRequest_InterfaceName)(0), // 2: magma.orc8r.StartTraceRequest.InterfaceName
	(*StartTraceRequest)(nil),            // 3: magma.orc8r.StartTraceRequest
	(*StartTraceResponse)(nil),           // 4: magma.orc8r.StartTraceResponse
	(*EndTraceRequest)(nil),              // 5: magma.orc8r.EndTraceRequest
	(*EndTraceResponse)(nil),             // 6: magma.orc8r.EndTraceResponse
	(*ReportEndedTraceRequest)(nil),      // 7: magma.orc8r.ReportEndedTraceRequest
	(*ReportEndedTraceResponse)(nil),     // 8: magma.orc8r.ReportEndedTraceResponse
}
var file_orc8r_protos_ctraced_proto_depIdxs = []int32{
	0, // 0: magma.orc8r.StartTraceRequest.trace_type:type_name -> magma.orc8r.StartTraceRequest.TraceType
	1, // 1: magma.orc8r.StartTraceRequest.protocol:type_name -> magma.orc8r.StartTraceRequest.ProtocolName
	2, // 2: magma.orc8r.StartTraceRequest.interface:type_name -> magma.orc8r.StartTraceRequest.InterfaceName
	3, // 3: magma.orc8r.CallTraceService.StartCallTrace:input_type -> magma.orc8r.StartTraceRequest
	5, // 4: magma.orc8r.CallTraceService.EndCallTrace:input_type -> magma.orc8r.EndTraceRequest
	7, // 5: magma.orc8r.CallTraceController.ReportEndedCallTrace:input_type -> magma.orc8r.ReportEndedTraceRequest
	4, // 6: magma.orc8r.CallTraceService.StartCallTrace:output_type -> magma.orc8r.StartTraceResponse
	6, // 7: magma.orc8r.CallTraceService.EndCallTrace:output_type -> magma.orc8r.EndTraceResponse
	8, // 8: magma.orc8r.CallTraceController.ReportEndedCallTrace:output_type -> magma.orc8r.ReportEndedTraceResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_orc8r_protos_ctraced_proto_init() }
func file_orc8r_protos_ctraced_proto_init() {
	if File_orc8r_protos_ctraced_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orc8r_protos_ctraced_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTraceRequest); i {
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
		file_orc8r_protos_ctraced_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTraceResponse); i {
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
		file_orc8r_protos_ctraced_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndTraceRequest); i {
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
		file_orc8r_protos_ctraced_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndTraceResponse); i {
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
		file_orc8r_protos_ctraced_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportEndedTraceRequest); i {
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
		file_orc8r_protos_ctraced_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportEndedTraceResponse); i {
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
			RawDescriptor: file_orc8r_protos_ctraced_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_orc8r_protos_ctraced_proto_goTypes,
		DependencyIndexes: file_orc8r_protos_ctraced_proto_depIdxs,
		EnumInfos:         file_orc8r_protos_ctraced_proto_enumTypes,
		MessageInfos:      file_orc8r_protos_ctraced_proto_msgTypes,
	}.Build()
	File_orc8r_protos_ctraced_proto = out.File
	file_orc8r_protos_ctraced_proto_rawDesc = nil
	file_orc8r_protos_ctraced_proto_goTypes = nil
	file_orc8r_protos_ctraced_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CallTraceServiceClient is the client API for CallTraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CallTraceServiceClient interface {
	// Start a call trace on the gateway with the specified options.
	// Only a single call trace can be active on a gateway at a time.
	//
	StartCallTrace(ctx context.Context, in *StartTraceRequest, opts ...grpc.CallOption) (*StartTraceResponse, error)
	// End the call trace that is currently active on the gateway.
	//
	EndCallTrace(ctx context.Context, in *EndTraceRequest, opts ...grpc.CallOption) (*EndTraceResponse, error)
}

type callTraceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCallTraceServiceClient(cc grpc.ClientConnInterface) CallTraceServiceClient {
	return &callTraceServiceClient{cc}
}

func (c *callTraceServiceClient) StartCallTrace(ctx context.Context, in *StartTraceRequest, opts ...grpc.CallOption) (*StartTraceResponse, error) {
	out := new(StartTraceResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.CallTraceService/StartCallTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callTraceServiceClient) EndCallTrace(ctx context.Context, in *EndTraceRequest, opts ...grpc.CallOption) (*EndTraceResponse, error) {
	out := new(EndTraceResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.CallTraceService/EndCallTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CallTraceServiceServer is the server API for CallTraceService service.
type CallTraceServiceServer interface {
	// Start a call trace on the gateway with the specified options.
	// Only a single call trace can be active on a gateway at a time.
	//
	StartCallTrace(context.Context, *StartTraceRequest) (*StartTraceResponse, error)
	// End the call trace that is currently active on the gateway.
	//
	EndCallTrace(context.Context, *EndTraceRequest) (*EndTraceResponse, error)
}

// UnimplementedCallTraceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCallTraceServiceServer struct {
}

func (*UnimplementedCallTraceServiceServer) StartCallTrace(context.Context, *StartTraceRequest) (*StartTraceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartCallTrace not implemented")
}
func (*UnimplementedCallTraceServiceServer) EndCallTrace(context.Context, *EndTraceRequest) (*EndTraceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndCallTrace not implemented")
}

func RegisterCallTraceServiceServer(s *grpc.Server, srv CallTraceServiceServer) {
	s.RegisterService(&_CallTraceService_serviceDesc, srv)
}

func _CallTraceService_StartCallTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallTraceServiceServer).StartCallTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.CallTraceService/StartCallTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallTraceServiceServer).StartCallTrace(ctx, req.(*StartTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallTraceService_EndCallTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallTraceServiceServer).EndCallTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.CallTraceService/EndCallTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallTraceServiceServer).EndCallTrace(ctx, req.(*EndTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CallTraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.CallTraceService",
	HandlerType: (*CallTraceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartCallTrace",
			Handler:    _CallTraceService_StartCallTrace_Handler,
		},
		{
			MethodName: "EndCallTrace",
			Handler:    _CallTraceService_EndCallTrace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/protos/ctraced.proto",
}

// CallTraceControllerClient is the client API for CallTraceController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CallTraceControllerClient interface {
	// Report that a call trace has ended
	//
	ReportEndedCallTrace(ctx context.Context, in *ReportEndedTraceRequest, opts ...grpc.CallOption) (*ReportEndedTraceResponse, error)
}

type callTraceControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCallTraceControllerClient(cc grpc.ClientConnInterface) CallTraceControllerClient {
	return &callTraceControllerClient{cc}
}

func (c *callTraceControllerClient) ReportEndedCallTrace(ctx context.Context, in *ReportEndedTraceRequest, opts ...grpc.CallOption) (*ReportEndedTraceResponse, error) {
	out := new(ReportEndedTraceResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.CallTraceController/ReportEndedCallTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CallTraceControllerServer is the server API for CallTraceController service.
type CallTraceControllerServer interface {
	// Report that a call trace has ended
	//
	ReportEndedCallTrace(context.Context, *ReportEndedTraceRequest) (*ReportEndedTraceResponse, error)
}

// UnimplementedCallTraceControllerServer can be embedded to have forward compatible implementations.
type UnimplementedCallTraceControllerServer struct {
}

func (*UnimplementedCallTraceControllerServer) ReportEndedCallTrace(context.Context, *ReportEndedTraceRequest) (*ReportEndedTraceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportEndedCallTrace not implemented")
}

func RegisterCallTraceControllerServer(s *grpc.Server, srv CallTraceControllerServer) {
	s.RegisterService(&_CallTraceController_serviceDesc, srv)
}

func _CallTraceController_ReportEndedCallTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportEndedTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallTraceControllerServer).ReportEndedCallTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.CallTraceController/ReportEndedCallTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallTraceControllerServer).ReportEndedCallTrace(ctx, req.(*ReportEndedTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CallTraceController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.CallTraceController",
	HandlerType: (*CallTraceControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportEndedCallTrace",
			Handler:    _CallTraceController_ReportEndedCallTrace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/protos/ctraced.proto",
}
