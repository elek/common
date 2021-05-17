// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: heldamount.proto

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

type GetHeldAmountRequest struct {
	Period               time.Time `protobuf:"bytes,1,opt,name=period,proto3,stdtime" json:"period"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetHeldAmountRequest) Reset()         { *m = GetHeldAmountRequest{} }
func (m *GetHeldAmountRequest) String() string { return proto.CompactTextString(m) }
func (*GetHeldAmountRequest) ProtoMessage()    {}
func (*GetHeldAmountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c5d9f6e3ee97993, []int{0}
}
func (m *GetHeldAmountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHeldAmountRequest.Unmarshal(m, b)
}
func (m *GetHeldAmountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHeldAmountRequest.Marshal(b, m, deterministic)
}
func (m *GetHeldAmountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHeldAmountRequest.Merge(m, src)
}
func (m *GetHeldAmountRequest) XXX_Size() int {
	return xxx_messageInfo_GetHeldAmountRequest.Size(m)
}
func (m *GetHeldAmountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHeldAmountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHeldAmountRequest proto.InternalMessageInfo

func (m *GetHeldAmountRequest) GetPeriod() time.Time {
	if m != nil {
		return m.Period
	}
	return time.Time{}
}

type GetHeldAmountResponse struct {
	Period               time.Time `protobuf:"bytes,1,opt,name=period,proto3,stdtime" json:"period"`
	NodeId               NodeID    `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	CreatedAt            time.Time `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	Codes                string    `protobuf:"bytes,4,opt,name=codes,proto3" json:"codes,omitempty"`
	UsageAtRest          float64   `protobuf:"fixed64,5,opt,name=usage_at_rest,json=usageAtRest,proto3" json:"usage_at_rest,omitempty"`
	UsageGet             int64     `protobuf:"varint,6,opt,name=usage_get,json=usageGet,proto3" json:"usage_get,omitempty"`
	UsagePut             int64     `protobuf:"varint,7,opt,name=usage_put,json=usagePut,proto3" json:"usage_put,omitempty"`
	UsageGetRepair       int64     `protobuf:"varint,8,opt,name=usage_get_repair,json=usageGetRepair,proto3" json:"usage_get_repair,omitempty"`
	UsagePutRepair       int64     `protobuf:"varint,9,opt,name=usage_put_repair,json=usagePutRepair,proto3" json:"usage_put_repair,omitempty"`
	UsageGetAudit        int64     `protobuf:"varint,10,opt,name=usage_get_audit,json=usageGetAudit,proto3" json:"usage_get_audit,omitempty"`
	CompAtRest           int64     `protobuf:"varint,11,opt,name=comp_at_rest,json=compAtRest,proto3" json:"comp_at_rest,omitempty"`
	CompGet              int64     `protobuf:"varint,12,opt,name=comp_get,json=compGet,proto3" json:"comp_get,omitempty"`
	CompPut              int64     `protobuf:"varint,13,opt,name=comp_put,json=compPut,proto3" json:"comp_put,omitempty"`
	CompGetRepair        int64     `protobuf:"varint,14,opt,name=comp_get_repair,json=compGetRepair,proto3" json:"comp_get_repair,omitempty"`
	CompPutRepair        int64     `protobuf:"varint,15,opt,name=comp_put_repair,json=compPutRepair,proto3" json:"comp_put_repair,omitempty"`
	CompGetAudit         int64     `protobuf:"varint,16,opt,name=comp_get_audit,json=compGetAudit,proto3" json:"comp_get_audit,omitempty"`
	SurgePercent         int64     `protobuf:"varint,17,opt,name=surge_percent,json=surgePercent,proto3" json:"surge_percent,omitempty"`
	Held                 int64     `protobuf:"varint,18,opt,name=held,proto3" json:"held,omitempty"`
	Owed                 int64     `protobuf:"varint,19,opt,name=owed,proto3" json:"owed,omitempty"`
	Disposed             int64     `protobuf:"varint,20,opt,name=disposed,proto3" json:"disposed,omitempty"`
	Paid                 int64     `protobuf:"varint,21,opt,name=paid,proto3" json:"paid,omitempty"`
	Distributed          int64     `protobuf:"varint,22,opt,name=distributed,proto3" json:"distributed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetHeldAmountResponse) Reset()         { *m = GetHeldAmountResponse{} }
func (m *GetHeldAmountResponse) String() string { return proto.CompactTextString(m) }
func (*GetHeldAmountResponse) ProtoMessage()    {}
func (*GetHeldAmountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c5d9f6e3ee97993, []int{1}
}
func (m *GetHeldAmountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHeldAmountResponse.Unmarshal(m, b)
}
func (m *GetHeldAmountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHeldAmountResponse.Marshal(b, m, deterministic)
}
func (m *GetHeldAmountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHeldAmountResponse.Merge(m, src)
}
func (m *GetHeldAmountResponse) XXX_Size() int {
	return xxx_messageInfo_GetHeldAmountResponse.Size(m)
}
func (m *GetHeldAmountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHeldAmountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHeldAmountResponse proto.InternalMessageInfo

func (m *GetHeldAmountResponse) GetPeriod() time.Time {
	if m != nil {
		return m.Period
	}
	return time.Time{}
}

func (m *GetHeldAmountResponse) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *GetHeldAmountResponse) GetCodes() string {
	if m != nil {
		return m.Codes
	}
	return ""
}

func (m *GetHeldAmountResponse) GetUsageAtRest() float64 {
	if m != nil {
		return m.UsageAtRest
	}
	return 0
}

func (m *GetHeldAmountResponse) GetUsageGet() int64 {
	if m != nil {
		return m.UsageGet
	}
	return 0
}

func (m *GetHeldAmountResponse) GetUsagePut() int64 {
	if m != nil {
		return m.UsagePut
	}
	return 0
}

func (m *GetHeldAmountResponse) GetUsageGetRepair() int64 {
	if m != nil {
		return m.UsageGetRepair
	}
	return 0
}

func (m *GetHeldAmountResponse) GetUsagePutRepair() int64 {
	if m != nil {
		return m.UsagePutRepair
	}
	return 0
}

func (m *GetHeldAmountResponse) GetUsageGetAudit() int64 {
	if m != nil {
		return m.UsageGetAudit
	}
	return 0
}

func (m *GetHeldAmountResponse) GetCompAtRest() int64 {
	if m != nil {
		return m.CompAtRest
	}
	return 0
}

func (m *GetHeldAmountResponse) GetCompGet() int64 {
	if m != nil {
		return m.CompGet
	}
	return 0
}

func (m *GetHeldAmountResponse) GetCompPut() int64 {
	if m != nil {
		return m.CompPut
	}
	return 0
}

func (m *GetHeldAmountResponse) GetCompGetRepair() int64 {
	if m != nil {
		return m.CompGetRepair
	}
	return 0
}

func (m *GetHeldAmountResponse) GetCompPutRepair() int64 {
	if m != nil {
		return m.CompPutRepair
	}
	return 0
}

func (m *GetHeldAmountResponse) GetCompGetAudit() int64 {
	if m != nil {
		return m.CompGetAudit
	}
	return 0
}

func (m *GetHeldAmountResponse) GetSurgePercent() int64 {
	if m != nil {
		return m.SurgePercent
	}
	return 0
}

func (m *GetHeldAmountResponse) GetHeld() int64 {
	if m != nil {
		return m.Held
	}
	return 0
}

func (m *GetHeldAmountResponse) GetOwed() int64 {
	if m != nil {
		return m.Owed
	}
	return 0
}

func (m *GetHeldAmountResponse) GetDisposed() int64 {
	if m != nil {
		return m.Disposed
	}
	return 0
}

func (m *GetHeldAmountResponse) GetPaid() int64 {
	if m != nil {
		return m.Paid
	}
	return 0
}

func (m *GetHeldAmountResponse) GetDistributed() int64 {
	if m != nil {
		return m.Distributed
	}
	return 0
}

type GetAllPaystubsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllPaystubsRequest) Reset()         { *m = GetAllPaystubsRequest{} }
func (m *GetAllPaystubsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllPaystubsRequest) ProtoMessage()    {}
func (*GetAllPaystubsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c5d9f6e3ee97993, []int{2}
}
func (m *GetAllPaystubsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllPaystubsRequest.Unmarshal(m, b)
}
func (m *GetAllPaystubsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllPaystubsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllPaystubsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllPaystubsRequest.Merge(m, src)
}
func (m *GetAllPaystubsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllPaystubsRequest.Size(m)
}
func (m *GetAllPaystubsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllPaystubsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllPaystubsRequest proto.InternalMessageInfo

type GetAllPaystubsResponse struct {
	Paystub              []*GetHeldAmountResponse `protobuf:"bytes,1,rep,name=paystub,proto3" json:"paystub,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GetAllPaystubsResponse) Reset()         { *m = GetAllPaystubsResponse{} }
func (m *GetAllPaystubsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllPaystubsResponse) ProtoMessage()    {}
func (*GetAllPaystubsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c5d9f6e3ee97993, []int{3}
}
func (m *GetAllPaystubsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllPaystubsResponse.Unmarshal(m, b)
}
func (m *GetAllPaystubsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllPaystubsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllPaystubsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllPaystubsResponse.Merge(m, src)
}
func (m *GetAllPaystubsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllPaystubsResponse.Size(m)
}
func (m *GetAllPaystubsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllPaystubsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllPaystubsResponse proto.InternalMessageInfo

func (m *GetAllPaystubsResponse) GetPaystub() []*GetHeldAmountResponse {
	if m != nil {
		return m.Paystub
	}
	return nil
}

type GetPaymentRequest struct {
	Period               time.Time `protobuf:"bytes,1,opt,name=period,proto3,stdtime" json:"period"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetPaymentRequest) Reset()         { *m = GetPaymentRequest{} }
func (m *GetPaymentRequest) String() string { return proto.CompactTextString(m) }
func (*GetPaymentRequest) ProtoMessage()    {}
func (*GetPaymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c5d9f6e3ee97993, []int{4}
}
func (m *GetPaymentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPaymentRequest.Unmarshal(m, b)
}
func (m *GetPaymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPaymentRequest.Marshal(b, m, deterministic)
}
func (m *GetPaymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPaymentRequest.Merge(m, src)
}
func (m *GetPaymentRequest) XXX_Size() int {
	return xxx_messageInfo_GetPaymentRequest.Size(m)
}
func (m *GetPaymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPaymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPaymentRequest proto.InternalMessageInfo

func (m *GetPaymentRequest) GetPeriod() time.Time {
	if m != nil {
		return m.Period
	}
	return time.Time{}
}

type GetPaymentResponse struct {
	NodeId               NodeID    `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	CreatedAt            time.Time `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	Period               time.Time `protobuf:"bytes,3,opt,name=period,proto3,stdtime" json:"period"`
	Amount               int64     `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Receipt              string    `protobuf:"bytes,5,opt,name=receipt,proto3" json:"receipt,omitempty"`
	Notes                string    `protobuf:"bytes,6,opt,name=notes,proto3" json:"notes,omitempty"`
	Id                   int64     `protobuf:"varint,7,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetPaymentResponse) Reset()         { *m = GetPaymentResponse{} }
func (m *GetPaymentResponse) String() string { return proto.CompactTextString(m) }
func (*GetPaymentResponse) ProtoMessage()    {}
func (*GetPaymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c5d9f6e3ee97993, []int{5}
}
func (m *GetPaymentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPaymentResponse.Unmarshal(m, b)
}
func (m *GetPaymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPaymentResponse.Marshal(b, m, deterministic)
}
func (m *GetPaymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPaymentResponse.Merge(m, src)
}
func (m *GetPaymentResponse) XXX_Size() int {
	return xxx_messageInfo_GetPaymentResponse.Size(m)
}
func (m *GetPaymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPaymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPaymentResponse proto.InternalMessageInfo

func (m *GetPaymentResponse) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *GetPaymentResponse) GetPeriod() time.Time {
	if m != nil {
		return m.Period
	}
	return time.Time{}
}

func (m *GetPaymentResponse) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *GetPaymentResponse) GetReceipt() string {
	if m != nil {
		return m.Receipt
	}
	return ""
}

func (m *GetPaymentResponse) GetNotes() string {
	if m != nil {
		return m.Notes
	}
	return ""
}

func (m *GetPaymentResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetAllPaymentsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllPaymentsRequest) Reset()         { *m = GetAllPaymentsRequest{} }
func (m *GetAllPaymentsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllPaymentsRequest) ProtoMessage()    {}
func (*GetAllPaymentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c5d9f6e3ee97993, []int{6}
}
func (m *GetAllPaymentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllPaymentsRequest.Unmarshal(m, b)
}
func (m *GetAllPaymentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllPaymentsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllPaymentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllPaymentsRequest.Merge(m, src)
}
func (m *GetAllPaymentsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllPaymentsRequest.Size(m)
}
func (m *GetAllPaymentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllPaymentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllPaymentsRequest proto.InternalMessageInfo

type GetAllPaymentsResponse struct {
	Payment              []*GetPaymentResponse `protobuf:"bytes,1,rep,name=payment,proto3" json:"payment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetAllPaymentsResponse) Reset()         { *m = GetAllPaymentsResponse{} }
func (m *GetAllPaymentsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllPaymentsResponse) ProtoMessage()    {}
func (*GetAllPaymentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c5d9f6e3ee97993, []int{7}
}
func (m *GetAllPaymentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllPaymentsResponse.Unmarshal(m, b)
}
func (m *GetAllPaymentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllPaymentsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllPaymentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllPaymentsResponse.Merge(m, src)
}
func (m *GetAllPaymentsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllPaymentsResponse.Size(m)
}
func (m *GetAllPaymentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllPaymentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllPaymentsResponse proto.InternalMessageInfo

func (m *GetAllPaymentsResponse) GetPayment() []*GetPaymentResponse {
	if m != nil {
		return m.Payment
	}
	return nil
}

func init() {
	proto.RegisterType((*GetHeldAmountRequest)(nil), "heldamount.GetHeldAmountRequest")
	proto.RegisterType((*GetHeldAmountResponse)(nil), "heldamount.GetHeldAmountResponse")
	proto.RegisterType((*GetAllPaystubsRequest)(nil), "heldamount.GetAllPaystubsRequest")
	proto.RegisterType((*GetAllPaystubsResponse)(nil), "heldamount.GetAllPaystubsResponse")
	proto.RegisterType((*GetPaymentRequest)(nil), "heldamount.GetPaymentRequest")
	proto.RegisterType((*GetPaymentResponse)(nil), "heldamount.GetPaymentResponse")
	proto.RegisterType((*GetAllPaymentsRequest)(nil), "heldamount.GetAllPaymentsRequest")
	proto.RegisterType((*GetAllPaymentsResponse)(nil), "heldamount.GetAllPaymentsResponse")
}

func init() { proto.RegisterFile("heldamount.proto", fileDescriptor_5c5d9f6e3ee97993) }

var fileDescriptor_5c5d9f6e3ee97993 = []byte{
	// 735 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xbe, 0x4e, 0xda, 0xfc, 0x9c, 0xfc, 0xb4, 0x9d, 0x9b, 0xf6, 0xce, 0x0d, 0x82, 0x98, 0x80,
	0x4a, 0x56, 0x89, 0x54, 0x36, 0x48, 0xb0, 0x49, 0x41, 0x2a, 0x15, 0x12, 0x4a, 0x4d, 0x11, 0x12,
	0x9b, 0xc8, 0xc9, 0x1c, 0x82, 0x51, 0x9c, 0x31, 0x9e, 0xb1, 0x50, 0xdf, 0x82, 0x25, 0x4b, 0x1e,
	0x87, 0x67, 0x60, 0x51, 0x5e, 0x05, 0xcd, 0x8c, 0x27, 0x76, 0x02, 0x6d, 0xd5, 0xc2, 0xce, 0xf3,
	0x9d, 0xef, 0x7c, 0x3e, 0x3f, 0xf3, 0x0d, 0x6c, 0xbf, 0xc7, 0x39, 0xf3, 0x43, 0x9e, 0x2c, 0x64,
	0x3f, 0x8a, 0xb9, 0xe4, 0x04, 0x32, 0xa4, 0x0d, 0x33, 0x3e, 0xe3, 0x06, 0x6f, 0x77, 0x66, 0x9c,
	0xcf, 0xe6, 0x38, 0xd0, 0xa7, 0x49, 0xf2, 0x6e, 0x20, 0x83, 0x10, 0x85, 0xf4, 0xc3, 0xc8, 0x10,
	0xba, 0xa7, 0xd0, 0x3a, 0x42, 0xf9, 0x1c, 0xe7, 0x6c, 0xa8, 0xb3, 0x3d, 0xfc, 0x98, 0xa0, 0x90,
	0xe4, 0x09, 0x94, 0x22, 0x8c, 0x03, 0xce, 0xa8, 0xe3, 0x3a, 0xbd, 0xda, 0x41, 0xbb, 0x6f, 0x94,
	0xfa, 0x56, 0xa9, 0x7f, 0x6a, 0x95, 0x0e, 0x2b, 0xdf, 0xce, 0x3b, 0xff, 0x7c, 0xfe, 0xd1, 0x71,
	0xbc, 0x34, 0xa7, 0xfb, 0xb5, 0x04, 0xbb, 0x6b, 0xb2, 0x22, 0xe2, 0x0b, 0x81, 0x7f, 0xa6, 0x4b,
	0x1e, 0x40, 0x79, 0xc1, 0x19, 0x8e, 0x03, 0x46, 0x0b, 0xae, 0xd3, 0xab, 0x1f, 0x36, 0x15, 0xe5,
	0xfb, 0x79, 0xa7, 0xf4, 0x92, 0x33, 0x3c, 0x7e, 0xe6, 0x95, 0x54, 0xf8, 0x98, 0x91, 0xa7, 0x00,
	0xd3, 0x18, 0x7d, 0x89, 0x6c, 0xec, 0x4b, 0x5a, 0xbc, 0xc6, 0xaf, 0xaa, 0x69, 0xde, 0x50, 0x92,
	0x16, 0x6c, 0x4e, 0x39, 0x43, 0x41, 0x37, 0x5c, 0xa7, 0x57, 0xf5, 0xcc, 0x81, 0x74, 0xa1, 0x91,
	0x08, 0x7f, 0x86, 0x63, 0x5f, 0x8e, 0x63, 0x14, 0x92, 0x6e, 0xba, 0x4e, 0xcf, 0xf1, 0x6a, 0x1a,
	0x1c, 0xaa, 0x4e, 0x25, 0xb9, 0x05, 0x55, 0xc3, 0x99, 0xa1, 0xa4, 0x25, 0xd7, 0xe9, 0x15, 0xbd,
	0x8a, 0x06, 0x8e, 0x30, 0x17, 0x8c, 0x12, 0x49, 0xcb, 0xb9, 0xe0, 0x28, 0x91, 0xa4, 0x07, 0xdb,
	0xcb, 0xcc, 0x71, 0x8c, 0x91, 0x1f, 0xc4, 0xb4, 0xa2, 0x39, 0x4d, 0x2b, 0xe0, 0x69, 0x34, 0x63,
	0x46, 0xc9, 0x92, 0x59, 0xcd, 0x31, 0x47, 0x89, 0x65, 0xee, 0xc3, 0x56, 0xa6, 0xe9, 0x27, 0x2c,
	0x90, 0x14, 0x34, 0xb1, 0x61, 0x25, 0x87, 0x0a, 0x24, 0x2e, 0xd4, 0xa7, 0x3c, 0x8c, 0x96, 0x8d,
	0xd5, 0x34, 0x09, 0x14, 0x96, 0xf6, 0xf5, 0x3f, 0x54, 0x34, 0x43, 0xb5, 0x55, 0xd7, 0xd1, 0xb2,
	0x3a, 0xab, 0xae, 0x6c, 0x48, 0x35, 0xd5, 0xc8, 0x42, 0xaa, 0xa7, 0x7d, 0xd8, 0xb2, 0x59, 0xb6,
	0xd0, 0xa6, 0xf9, 0x7f, 0x9a, 0x9c, 0xd5, 0x69, 0x25, 0x2c, 0x6f, 0x2b, 0xe3, 0x65, 0xfd, 0xdc,
	0x87, 0xe6, 0x52, 0xcf, 0xb4, 0xb3, 0xad, 0x69, 0xf5, 0x54, 0xce, 0x74, 0x73, 0x0f, 0x1a, 0x22,
	0x89, 0xd5, 0x7c, 0x30, 0x9e, 0xe2, 0x42, 0xd2, 0x1d, 0x43, 0xd2, 0xe0, 0xc8, 0x60, 0x84, 0xc0,
	0x86, 0x72, 0x0e, 0x25, 0x3a, 0xa6, 0xbf, 0x15, 0xc6, 0x3f, 0x21, 0xa3, 0xff, 0x1a, 0x4c, 0x7d,
	0x93, 0x36, 0x54, 0x58, 0x20, 0x22, 0x2e, 0x90, 0xd1, 0x96, 0x59, 0x99, 0x3d, 0x2b, 0x7e, 0xe4,
	0x07, 0x8c, 0xee, 0x1a, 0xbe, 0xfa, 0x26, 0x2e, 0xd4, 0x58, 0x20, 0x64, 0x1c, 0x4c, 0x12, 0x89,
	0x8c, 0xee, 0xe9, 0x50, 0x1e, 0xea, 0xfe, 0xa7, 0x1d, 0x32, 0x9c, 0xcf, 0x47, 0xfe, 0x99, 0x90,
	0xc9, 0x44, 0xa4, 0xce, 0xeb, 0xbe, 0x86, 0xbd, 0xf5, 0x40, 0xea, 0x9d, 0xc7, 0x50, 0x8e, 0x0c,
	0x46, 0x1d, 0xb7, 0xd8, 0xab, 0x1d, 0xdc, 0xed, 0xe7, 0x1e, 0x82, 0xdf, 0xfa, 0xcd, 0xb3, 0x19,
	0xdd, 0x13, 0xd8, 0x39, 0x42, 0x39, 0xf2, 0xcf, 0x42, 0xfc, 0x5b, 0x2e, 0xff, 0x52, 0x00, 0x92,
	0xd7, 0x4c, 0xcb, 0xcc, 0x99, 0xd4, 0xb9, 0x86, 0x49, 0x0b, 0x37, 0x33, 0x69, 0xd6, 0x42, 0xf1,
	0x06, 0x0f, 0xca, 0x1e, 0x94, 0xcc, 0xf8, 0xb4, 0xc7, 0x8b, 0x5e, 0x7a, 0x22, 0x14, 0xca, 0x31,
	0x4e, 0x31, 0x88, 0x8c, 0xbd, 0xab, 0x9e, 0x3d, 0xaa, 0x47, 0x61, 0xc1, 0x25, 0x0a, 0x6d, 0xeb,
	0xaa, 0x67, 0x0e, 0xa4, 0x09, 0x85, 0x80, 0xa5, 0x66, 0x2e, 0x04, 0xab, 0xdb, 0x55, 0xc3, 0x59,
	0x6e, 0xd7, 0xcb, 0x6d, 0x37, 0x0d, 0xa4, 0x63, 0x7b, 0xa4, 0xb7, 0xab, 0xb0, 0x74, 0xbb, 0x77,
	0xd6, 0xb6, 0xbb, 0x36, 0x67, 0xcf, 0xd2, 0x0f, 0xce, 0x0b, 0x00, 0xd9, 0xea, 0xc9, 0x09, 0x80,
	0x61, 0xbf, 0x92, 0xc9, 0x84, 0xb8, 0x97, 0xdc, 0x11, 0x5d, 0x52, 0xfb, 0xea, 0x5b, 0x44, 0xde,
	0x40, 0x73, 0xf5, 0x4e, 0x92, 0xf5, 0xa4, 0x5f, 0x2f, 0x72, 0xbb, 0x7b, 0x19, 0x25, 0x15, 0x7e,
	0x61, 0x6b, 0x55, 0x8d, 0x90, 0xdb, 0x17, 0x75, 0x6c, 0x04, 0xaf, 0x18, 0xc8, 0x4a, 0x95, 0x7a,
	0xb6, 0x17, 0x54, 0x99, 0x5f, 0xc8, 0x05, 0x55, 0xae, 0xac, 0xe6, 0xb0, 0xf5, 0x96, 0x08, 0xc9,
	0xe3, 0x0f, 0xfd, 0x80, 0x0f, 0xa6, 0x3c, 0x0c, 0xf9, 0x62, 0x10, 0x4d, 0x26, 0x25, 0x7d, 0xc3,
	0x1e, 0xfe, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x1c, 0xca, 0xd8, 0x8e, 0x07, 0x00, 0x00,
}
