// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: contact.proto

package pb

import (
	fmt "fmt"
	math "math"
	time "time"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CheckInRequest struct {
	Address              string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Version              *NodeVersion  `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Capacity             *NodeCapacity `protobuf:"bytes,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Operator             *NodeOperator `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CheckInRequest) Reset()         { *m = CheckInRequest{} }
func (m *CheckInRequest) String() string { return proto.CompactTextString(m) }
func (*CheckInRequest) ProtoMessage()    {}
func (*CheckInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{0}
}
func (m *CheckInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckInRequest.Unmarshal(m, b)
}
func (m *CheckInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckInRequest.Marshal(b, m, deterministic)
}
func (m *CheckInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckInRequest.Merge(m, src)
}
func (m *CheckInRequest) XXX_Size() int {
	return xxx_messageInfo_CheckInRequest.Size(m)
}
func (m *CheckInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckInRequest proto.InternalMessageInfo

func (m *CheckInRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CheckInRequest) GetVersion() *NodeVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *CheckInRequest) GetCapacity() *NodeCapacity {
	if m != nil {
		return m.Capacity
	}
	return nil
}

func (m *CheckInRequest) GetOperator() *NodeOperator {
	if m != nil {
		return m.Operator
	}
	return nil
}

type CheckInResponse struct {
	PingNodeSuccess      bool     `protobuf:"varint,1,opt,name=ping_node_success,json=pingNodeSuccess,proto3" json:"ping_node_success,omitempty"`
	PingErrorMessage     string   `protobuf:"bytes,2,opt,name=ping_error_message,json=pingErrorMessage,proto3" json:"ping_error_message,omitempty"`
	PingNodeSuccessQuic  bool     `protobuf:"varint,3,opt,name=ping_node_success_quic,json=pingNodeSuccessQuic,proto3" json:"ping_node_success_quic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckInResponse) Reset()         { *m = CheckInResponse{} }
func (m *CheckInResponse) String() string { return proto.CompactTextString(m) }
func (*CheckInResponse) ProtoMessage()    {}
func (*CheckInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{1}
}
func (m *CheckInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckInResponse.Unmarshal(m, b)
}
func (m *CheckInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckInResponse.Marshal(b, m, deterministic)
}
func (m *CheckInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckInResponse.Merge(m, src)
}
func (m *CheckInResponse) XXX_Size() int {
	return xxx_messageInfo_CheckInResponse.Size(m)
}
func (m *CheckInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckInResponse proto.InternalMessageInfo

func (m *CheckInResponse) GetPingNodeSuccess() bool {
	if m != nil {
		return m.PingNodeSuccess
	}
	return false
}

func (m *CheckInResponse) GetPingErrorMessage() string {
	if m != nil {
		return m.PingErrorMessage
	}
	return ""
}

func (m *CheckInResponse) GetPingNodeSuccessQuic() bool {
	if m != nil {
		return m.PingNodeSuccessQuic
	}
	return false
}

type GetTimeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTimeRequest) Reset()         { *m = GetTimeRequest{} }
func (m *GetTimeRequest) String() string { return proto.CompactTextString(m) }
func (*GetTimeRequest) ProtoMessage()    {}
func (*GetTimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{2}
}
func (m *GetTimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTimeRequest.Unmarshal(m, b)
}
func (m *GetTimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTimeRequest.Marshal(b, m, deterministic)
}
func (m *GetTimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTimeRequest.Merge(m, src)
}
func (m *GetTimeRequest) XXX_Size() int {
	return xxx_messageInfo_GetTimeRequest.Size(m)
}
func (m *GetTimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTimeRequest proto.InternalMessageInfo

type GetTimeResponse struct {
	Timestamp            time.Time `protobuf:"bytes,1,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetTimeResponse) Reset()         { *m = GetTimeResponse{} }
func (m *GetTimeResponse) String() string { return proto.CompactTextString(m) }
func (*GetTimeResponse) ProtoMessage()    {}
func (*GetTimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{3}
}
func (m *GetTimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTimeResponse.Unmarshal(m, b)
}
func (m *GetTimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTimeResponse.Marshal(b, m, deterministic)
}
func (m *GetTimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTimeResponse.Merge(m, src)
}
func (m *GetTimeResponse) XXX_Size() int {
	return xxx_messageInfo_GetTimeResponse.Size(m)
}
func (m *GetTimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTimeResponse proto.InternalMessageInfo

func (m *GetTimeResponse) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

type ContactPingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactPingRequest) Reset()         { *m = ContactPingRequest{} }
func (m *ContactPingRequest) String() string { return proto.CompactTextString(m) }
func (*ContactPingRequest) ProtoMessage()    {}
func (*ContactPingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{4}
}
func (m *ContactPingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactPingRequest.Unmarshal(m, b)
}
func (m *ContactPingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactPingRequest.Marshal(b, m, deterministic)
}
func (m *ContactPingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactPingRequest.Merge(m, src)
}
func (m *ContactPingRequest) XXX_Size() int {
	return xxx_messageInfo_ContactPingRequest.Size(m)
}
func (m *ContactPingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactPingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContactPingRequest proto.InternalMessageInfo

type ContactPingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactPingResponse) Reset()         { *m = ContactPingResponse{} }
func (m *ContactPingResponse) String() string { return proto.CompactTextString(m) }
func (*ContactPingResponse) ProtoMessage()    {}
func (*ContactPingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{5}
}
func (m *ContactPingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactPingResponse.Unmarshal(m, b)
}
func (m *ContactPingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactPingResponse.Marshal(b, m, deterministic)
}
func (m *ContactPingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactPingResponse.Merge(m, src)
}
func (m *ContactPingResponse) XXX_Size() int {
	return xxx_messageInfo_ContactPingResponse.Size(m)
}
func (m *ContactPingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactPingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContactPingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CheckInRequest)(nil), "contact.CheckInRequest")
	proto.RegisterType((*CheckInResponse)(nil), "contact.CheckInResponse")
	proto.RegisterType((*GetTimeRequest)(nil), "contact.GetTimeRequest")
	proto.RegisterType((*GetTimeResponse)(nil), "contact.GetTimeResponse")
	proto.RegisterType((*ContactPingRequest)(nil), "contact.ContactPingRequest")
	proto.RegisterType((*ContactPingResponse)(nil), "contact.ContactPingResponse")
}

func init() { proto.RegisterFile("contact.proto", fileDescriptor_a5036fff2565fb15) }

var fileDescriptor_a5036fff2565fb15 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4b, 0x6e, 0xd4, 0x40,
	0x10, 0xc5, 0x10, 0x61, 0x4f, 0x45, 0x64, 0x92, 0x4e, 0x00, 0xcb, 0x20, 0x4d, 0xe4, 0x55, 0x04,
	0xc8, 0x23, 0x4d, 0xb6, 0x59, 0xcd, 0x28, 0x42, 0x2c, 0x80, 0xa1, 0x09, 0x2c, 0xd8, 0x8c, 0x3c,
	0xed, 0xc2, 0x34, 0x60, 0x97, 0xd3, 0xdd, 0x46, 0x62, 0xcb, 0x09, 0x38, 0x02, 0xb7, 0xe0, 0x0a,
	0x9c, 0x02, 0xae, 0x82, 0xdc, 0x6e, 0x7b, 0x7e, 0xd9, 0xb9, 0xdf, 0x7b, 0xf5, 0x5c, 0xf5, 0xaa,
	0xe0, 0x9e, 0xa0, 0xd2, 0xa4, 0xc2, 0x24, 0x95, 0x22, 0x43, 0xcc, 0x77, 0xcf, 0x08, 0x72, 0xca,
	0xa9, 0x05, 0xa3, 0x51, 0x4e, 0x94, 0x7f, 0xc5, 0xb1, 0x7d, 0x2d, 0xeb, 0x8f, 0x63, 0x23, 0x0b,
	0xd4, 0x26, 0x2d, 0x2a, 0x27, 0x80, 0x92, 0x32, 0x6c, 0xbf, 0xe3, 0xdf, 0x1e, 0x1c, 0xcc, 0x3e,
	0xa1, 0xf8, 0xf2, 0xa2, 0xe4, 0x78, 0x5d, 0xa3, 0x36, 0x2c, 0x04, 0x3f, 0xcd, 0x32, 0x85, 0x5a,
	0x87, 0xde, 0xa9, 0x77, 0x36, 0xe0, 0xdd, 0x93, 0x3d, 0x05, 0xff, 0x1b, 0x2a, 0x2d, 0xa9, 0x0c,
	0x6f, 0x9f, 0x7a, 0x67, 0xfb, 0x93, 0xa3, 0xc4, 0x5a, 0xbd, 0xa2, 0x0c, 0xdf, 0xb7, 0x04, 0xef,
	0x14, 0x2c, 0x81, 0x40, 0xa4, 0x55, 0x2a, 0xa4, 0xf9, 0x1e, 0xde, 0xb1, 0x6a, 0xb6, 0x52, 0xcf,
	0x1c, 0xc3, 0x7b, 0x4d, 0xa3, 0xa7, 0x0a, 0x55, 0x6a, 0x48, 0x85, 0x7b, 0xdb, 0xfa, 0xd7, 0x8e,
	0xe1, 0xbd, 0x26, 0xfe, 0xe5, 0xc1, 0xb0, 0xef, 0x5c, 0x57, 0x54, 0x6a, 0x64, 0x4f, 0xe0, 0xa8,
	0x92, 0x65, 0xbe, 0x68, 0xea, 0x16, 0xba, 0x16, 0xa2, 0x1b, 0x22, 0xe0, 0xc3, 0x86, 0x68, 0xac,
	0xde, 0xb6, 0x30, 0x7b, 0x06, 0xcc, 0x6a, 0x51, 0x29, 0x52, 0x8b, 0x02, 0xb5, 0x4e, 0x73, 0xb4,
	0x73, 0x0d, 0xf8, 0x61, 0xc3, 0x5c, 0x36, 0xc4, 0xcb, 0x16, 0x67, 0xe7, 0xf0, 0x60, 0xc7, 0x79,
	0x71, 0x5d, 0x4b, 0x61, 0x67, 0x0b, 0xf8, 0xf1, 0x96, 0xfd, 0x9b, 0x5a, 0x8a, 0xf8, 0x10, 0x0e,
	0x9e, 0xa3, 0xb9, 0x92, 0x05, 0xba, 0x6c, 0xe3, 0x77, 0x30, 0xec, 0x11, 0xd7, 0xf3, 0x14, 0x06,
	0xfd, 0x82, 0x6c, 0xaf, 0xfb, 0x93, 0x28, 0x69, 0x57, 0x98, 0x74, 0x2b, 0x4c, 0xae, 0x3a, 0xc5,
	0x34, 0xf8, 0xf3, 0x77, 0x74, 0xeb, 0xe7, 0xbf, 0x91, 0xc7, 0x57, 0x65, 0xf1, 0x09, 0xb0, 0x59,
	0x7b, 0x09, 0x73, 0x59, 0xe6, 0xdd, 0xcf, 0xee, 0xc3, 0xf1, 0x06, 0xda, 0xfe, 0x70, 0x32, 0x07,
	0xdf, 0xc1, 0xec, 0x12, 0x82, 0xb9, 0xeb, 0x9b, 0x3d, 0x4a, 0xba, 0xdb, 0xda, 0xb5, 0x8a, 0x1e,
	0xdf, 0x4c, 0x3a, 0xc7, 0x1f, 0x1e, 0xec, 0x59, 0x8f, 0x0b, 0xf0, 0xdd, 0x4a, 0xd8, 0xc3, 0x55,
	0xc5, 0xc6, 0x79, 0x45, 0xe1, 0x2e, 0xe1, 0x92, 0xb8, 0x00, 0xdf, 0x85, 0xb3, 0x56, 0xbd, 0x19,
	0xe0, 0x5a, 0xf5, 0x56, 0x8e, 0xd3, 0x93, 0x0f, 0x4c, 0x1b, 0x52, 0x9f, 0x13, 0x49, 0x63, 0x41,
	0x45, 0x41, 0xe5, 0xb8, 0x5a, 0x2e, 0xef, 0xda, 0x08, 0xcf, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xde, 0xed, 0x19, 0x48, 0x39, 0x03, 0x00, 0x00,
}
