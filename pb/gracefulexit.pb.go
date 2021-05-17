// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gracefulexit.proto

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

type TransferFailed_Error int32

const (
	TransferFailed_NOT_FOUND                TransferFailed_Error = 0
	TransferFailed_STORAGE_NODE_UNAVAILABLE TransferFailed_Error = 1
	TransferFailed_HASH_VERIFICATION        TransferFailed_Error = 2
	TransferFailed_UNKNOWN                  TransferFailed_Error = 10
)

var TransferFailed_Error_name = map[int32]string{
	0:  "NOT_FOUND",
	1:  "STORAGE_NODE_UNAVAILABLE",
	2:  "HASH_VERIFICATION",
	10: "UNKNOWN",
}

var TransferFailed_Error_value = map[string]int32{
	"NOT_FOUND":                0,
	"STORAGE_NODE_UNAVAILABLE": 1,
	"HASH_VERIFICATION":        2,
	"UNKNOWN":                  10,
}

func (x TransferFailed_Error) String() string {
	return proto.EnumName(TransferFailed_Error_name, int32(x))
}

func (TransferFailed_Error) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{3, 0}
}

type ExitFailed_Reason int32

const (
	ExitFailed_VERIFICATION_FAILED                 ExitFailed_Reason = 0
	ExitFailed_INACTIVE_TIMEFRAME_EXCEEDED         ExitFailed_Reason = 1
	ExitFailed_OVERALL_FAILURE_PERCENTAGE_EXCEEDED ExitFailed_Reason = 2
)

var ExitFailed_Reason_name = map[int32]string{
	0: "VERIFICATION_FAILED",
	1: "INACTIVE_TIMEFRAME_EXCEEDED",
	2: "OVERALL_FAILURE_PERCENTAGE_EXCEEDED",
}

var ExitFailed_Reason_value = map[string]int32{
	"VERIFICATION_FAILED":                 0,
	"INACTIVE_TIMEFRAME_EXCEEDED":         1,
	"OVERALL_FAILURE_PERCENTAGE_EXCEEDED": 2,
}

func (x ExitFailed_Reason) String() string {
	return proto.EnumName(ExitFailed_Reason_name, int32(x))
}

func (ExitFailed_Reason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{9, 0}
}

type GracefulExitFeasibilityRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GracefulExitFeasibilityRequest) Reset()         { *m = GracefulExitFeasibilityRequest{} }
func (m *GracefulExitFeasibilityRequest) String() string { return proto.CompactTextString(m) }
func (*GracefulExitFeasibilityRequest) ProtoMessage()    {}
func (*GracefulExitFeasibilityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{0}
}
func (m *GracefulExitFeasibilityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GracefulExitFeasibilityRequest.Unmarshal(m, b)
}
func (m *GracefulExitFeasibilityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GracefulExitFeasibilityRequest.Marshal(b, m, deterministic)
}
func (m *GracefulExitFeasibilityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GracefulExitFeasibilityRequest.Merge(m, src)
}
func (m *GracefulExitFeasibilityRequest) XXX_Size() int {
	return xxx_messageInfo_GracefulExitFeasibilityRequest.Size(m)
}
func (m *GracefulExitFeasibilityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GracefulExitFeasibilityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GracefulExitFeasibilityRequest proto.InternalMessageInfo

type GracefulExitFeasibilityResponse struct {
	JoinedAt             time.Time `protobuf:"bytes,1,opt,name=joined_at,json=joinedAt,proto3,stdtime" json:"joined_at"`
	MonthsRequired       int32     `protobuf:"varint,2,opt,name=months_required,json=monthsRequired,proto3" json:"months_required,omitempty"`
	IsAllowed            bool      `protobuf:"varint,3,opt,name=is_allowed,json=isAllowed,proto3" json:"is_allowed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GracefulExitFeasibilityResponse) Reset()         { *m = GracefulExitFeasibilityResponse{} }
func (m *GracefulExitFeasibilityResponse) String() string { return proto.CompactTextString(m) }
func (*GracefulExitFeasibilityResponse) ProtoMessage()    {}
func (*GracefulExitFeasibilityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{1}
}
func (m *GracefulExitFeasibilityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GracefulExitFeasibilityResponse.Unmarshal(m, b)
}
func (m *GracefulExitFeasibilityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GracefulExitFeasibilityResponse.Marshal(b, m, deterministic)
}
func (m *GracefulExitFeasibilityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GracefulExitFeasibilityResponse.Merge(m, src)
}
func (m *GracefulExitFeasibilityResponse) XXX_Size() int {
	return xxx_messageInfo_GracefulExitFeasibilityResponse.Size(m)
}
func (m *GracefulExitFeasibilityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GracefulExitFeasibilityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GracefulExitFeasibilityResponse proto.InternalMessageInfo

func (m *GracefulExitFeasibilityResponse) GetJoinedAt() time.Time {
	if m != nil {
		return m.JoinedAt
	}
	return time.Time{}
}

func (m *GracefulExitFeasibilityResponse) GetMonthsRequired() int32 {
	if m != nil {
		return m.MonthsRequired
	}
	return 0
}

func (m *GracefulExitFeasibilityResponse) GetIsAllowed() bool {
	if m != nil {
		return m.IsAllowed
	}
	return false
}

type TransferSucceeded struct {
	OriginalOrderLimit   *OrderLimit `protobuf:"bytes,1,opt,name=original_order_limit,json=originalOrderLimit,proto3" json:"original_order_limit,omitempty"`
	OriginalPieceHash    *PieceHash  `protobuf:"bytes,2,opt,name=original_piece_hash,json=originalPieceHash,proto3" json:"original_piece_hash,omitempty"`
	ReplacementPieceHash *PieceHash  `protobuf:"bytes,3,opt,name=replacement_piece_hash,json=replacementPieceHash,proto3" json:"replacement_piece_hash,omitempty"`
	OriginalPieceId      PieceID     `protobuf:"bytes,4,opt,name=original_piece_id,json=originalPieceId,proto3,customtype=PieceID" json:"original_piece_id"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TransferSucceeded) Reset()         { *m = TransferSucceeded{} }
func (m *TransferSucceeded) String() string { return proto.CompactTextString(m) }
func (*TransferSucceeded) ProtoMessage()    {}
func (*TransferSucceeded) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{2}
}
func (m *TransferSucceeded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferSucceeded.Unmarshal(m, b)
}
func (m *TransferSucceeded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferSucceeded.Marshal(b, m, deterministic)
}
func (m *TransferSucceeded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferSucceeded.Merge(m, src)
}
func (m *TransferSucceeded) XXX_Size() int {
	return xxx_messageInfo_TransferSucceeded.Size(m)
}
func (m *TransferSucceeded) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferSucceeded.DiscardUnknown(m)
}

var xxx_messageInfo_TransferSucceeded proto.InternalMessageInfo

func (m *TransferSucceeded) GetOriginalOrderLimit() *OrderLimit {
	if m != nil {
		return m.OriginalOrderLimit
	}
	return nil
}

func (m *TransferSucceeded) GetOriginalPieceHash() *PieceHash {
	if m != nil {
		return m.OriginalPieceHash
	}
	return nil
}

func (m *TransferSucceeded) GetReplacementPieceHash() *PieceHash {
	if m != nil {
		return m.ReplacementPieceHash
	}
	return nil
}

type TransferFailed struct {
	OriginalPieceId      PieceID              `protobuf:"bytes,1,opt,name=original_piece_id,json=originalPieceId,proto3,customtype=PieceID" json:"original_piece_id"`
	Error                TransferFailed_Error `protobuf:"varint,2,opt,name=error,proto3,enum=gracefulexit.TransferFailed_Error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TransferFailed) Reset()         { *m = TransferFailed{} }
func (m *TransferFailed) String() string { return proto.CompactTextString(m) }
func (*TransferFailed) ProtoMessage()    {}
func (*TransferFailed) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{3}
}
func (m *TransferFailed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferFailed.Unmarshal(m, b)
}
func (m *TransferFailed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferFailed.Marshal(b, m, deterministic)
}
func (m *TransferFailed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferFailed.Merge(m, src)
}
func (m *TransferFailed) XXX_Size() int {
	return xxx_messageInfo_TransferFailed.Size(m)
}
func (m *TransferFailed) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferFailed.DiscardUnknown(m)
}

var xxx_messageInfo_TransferFailed proto.InternalMessageInfo

func (m *TransferFailed) GetError() TransferFailed_Error {
	if m != nil {
		return m.Error
	}
	return TransferFailed_NOT_FOUND
}

type StorageNodeMessage struct {
	// Types that are valid to be assigned to Message:
	//	*StorageNodeMessage_Succeeded
	//	*StorageNodeMessage_Failed
	Message              isStorageNodeMessage_Message `protobuf_oneof:"Message"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *StorageNodeMessage) Reset()         { *m = StorageNodeMessage{} }
func (m *StorageNodeMessage) String() string { return proto.CompactTextString(m) }
func (*StorageNodeMessage) ProtoMessage()    {}
func (*StorageNodeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{4}
}
func (m *StorageNodeMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageNodeMessage.Unmarshal(m, b)
}
func (m *StorageNodeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageNodeMessage.Marshal(b, m, deterministic)
}
func (m *StorageNodeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageNodeMessage.Merge(m, src)
}
func (m *StorageNodeMessage) XXX_Size() int {
	return xxx_messageInfo_StorageNodeMessage.Size(m)
}
func (m *StorageNodeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageNodeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StorageNodeMessage proto.InternalMessageInfo

type isStorageNodeMessage_Message interface {
	isStorageNodeMessage_Message()
}

type StorageNodeMessage_Succeeded struct {
	Succeeded *TransferSucceeded `protobuf:"bytes,1,opt,name=succeeded,proto3,oneof"`
}
type StorageNodeMessage_Failed struct {
	Failed *TransferFailed `protobuf:"bytes,2,opt,name=failed,proto3,oneof"`
}

func (*StorageNodeMessage_Succeeded) isStorageNodeMessage_Message() {}
func (*StorageNodeMessage_Failed) isStorageNodeMessage_Message()    {}

func (m *StorageNodeMessage) GetMessage() isStorageNodeMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *StorageNodeMessage) GetSucceeded() *TransferSucceeded {
	if x, ok := m.GetMessage().(*StorageNodeMessage_Succeeded); ok {
		return x.Succeeded
	}
	return nil
}

func (m *StorageNodeMessage) GetFailed() *TransferFailed {
	if x, ok := m.GetMessage().(*StorageNodeMessage_Failed); ok {
		return x.Failed
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StorageNodeMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StorageNodeMessage_OneofMarshaler, _StorageNodeMessage_OneofUnmarshaler, _StorageNodeMessage_OneofSizer, []interface{}{
		(*StorageNodeMessage_Succeeded)(nil),
		(*StorageNodeMessage_Failed)(nil),
	}
}

func _StorageNodeMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StorageNodeMessage)
	// Message
	switch x := m.Message.(type) {
	case *StorageNodeMessage_Succeeded:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Succeeded); err != nil {
			return err
		}
	case *StorageNodeMessage_Failed:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Failed); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StorageNodeMessage.Message has unexpected type %T", x)
	}
	return nil
}

func _StorageNodeMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StorageNodeMessage)
	switch tag {
	case 1: // Message.succeeded
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransferSucceeded)
		err := b.DecodeMessage(msg)
		m.Message = &StorageNodeMessage_Succeeded{msg}
		return true, err
	case 2: // Message.failed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransferFailed)
		err := b.DecodeMessage(msg)
		m.Message = &StorageNodeMessage_Failed{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StorageNodeMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StorageNodeMessage)
	// Message
	switch x := m.Message.(type) {
	case *StorageNodeMessage_Succeeded:
		s := proto.Size(x.Succeeded)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StorageNodeMessage_Failed:
		s := proto.Size(x.Failed)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NotReady struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotReady) Reset()         { *m = NotReady{} }
func (m *NotReady) String() string { return proto.CompactTextString(m) }
func (*NotReady) ProtoMessage()    {}
func (*NotReady) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{5}
}
func (m *NotReady) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotReady.Unmarshal(m, b)
}
func (m *NotReady) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotReady.Marshal(b, m, deterministic)
}
func (m *NotReady) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotReady.Merge(m, src)
}
func (m *NotReady) XXX_Size() int {
	return xxx_messageInfo_NotReady.Size(m)
}
func (m *NotReady) XXX_DiscardUnknown() {
	xxx_messageInfo_NotReady.DiscardUnknown(m)
}

var xxx_messageInfo_NotReady proto.InternalMessageInfo

type TransferPiece struct {
	OriginalPieceId PieceID         `protobuf:"bytes,1,opt,name=original_piece_id,json=originalPieceId,proto3,customtype=PieceID" json:"original_piece_id"`
	PrivateKey      PiecePrivateKey `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3,customtype=PiecePrivateKey" json:"private_key"`
	// addressed_order_limit contains the new piece id.
	AddressedOrderLimit  *AddressedOrderLimit `protobuf:"bytes,3,opt,name=addressed_order_limit,json=addressedOrderLimit,proto3" json:"addressed_order_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TransferPiece) Reset()         { *m = TransferPiece{} }
func (m *TransferPiece) String() string { return proto.CompactTextString(m) }
func (*TransferPiece) ProtoMessage()    {}
func (*TransferPiece) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{6}
}
func (m *TransferPiece) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferPiece.Unmarshal(m, b)
}
func (m *TransferPiece) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferPiece.Marshal(b, m, deterministic)
}
func (m *TransferPiece) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferPiece.Merge(m, src)
}
func (m *TransferPiece) XXX_Size() int {
	return xxx_messageInfo_TransferPiece.Size(m)
}
func (m *TransferPiece) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferPiece.DiscardUnknown(m)
}

var xxx_messageInfo_TransferPiece proto.InternalMessageInfo

func (m *TransferPiece) GetAddressedOrderLimit() *AddressedOrderLimit {
	if m != nil {
		return m.AddressedOrderLimit
	}
	return nil
}

type DeletePiece struct {
	OriginalPieceId      PieceID  `protobuf:"bytes,1,opt,name=original_piece_id,json=originalPieceId,proto3,customtype=PieceID" json:"original_piece_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePiece) Reset()         { *m = DeletePiece{} }
func (m *DeletePiece) String() string { return proto.CompactTextString(m) }
func (*DeletePiece) ProtoMessage()    {}
func (*DeletePiece) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{7}
}
func (m *DeletePiece) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePiece.Unmarshal(m, b)
}
func (m *DeletePiece) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePiece.Marshal(b, m, deterministic)
}
func (m *DeletePiece) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePiece.Merge(m, src)
}
func (m *DeletePiece) XXX_Size() int {
	return xxx_messageInfo_DeletePiece.Size(m)
}
func (m *DeletePiece) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePiece.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePiece proto.InternalMessageInfo

type ExitCompleted struct {
	// when everything is completed
	ExitCompleteSignature []byte `protobuf:"bytes,1,opt,name=exit_complete_signature,json=exitCompleteSignature,proto3" json:"exit_complete_signature,omitempty"`
	// satellite who issued this exit completed
	SatelliteId NodeID `protobuf:"bytes,2,opt,name=satellite_id,json=satelliteId,proto3,customtype=NodeID" json:"satellite_id"`
	// storage node this exit completed was issued to
	NodeId NodeID `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	// timestamp when the exit completed
	Completed            time.Time `protobuf:"bytes,4,opt,name=completed,proto3,stdtime" json:"completed"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ExitCompleted) Reset()         { *m = ExitCompleted{} }
func (m *ExitCompleted) String() string { return proto.CompactTextString(m) }
func (*ExitCompleted) ProtoMessage()    {}
func (*ExitCompleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{8}
}
func (m *ExitCompleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitCompleted.Unmarshal(m, b)
}
func (m *ExitCompleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitCompleted.Marshal(b, m, deterministic)
}
func (m *ExitCompleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitCompleted.Merge(m, src)
}
func (m *ExitCompleted) XXX_Size() int {
	return xxx_messageInfo_ExitCompleted.Size(m)
}
func (m *ExitCompleted) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitCompleted.DiscardUnknown(m)
}

var xxx_messageInfo_ExitCompleted proto.InternalMessageInfo

func (m *ExitCompleted) GetExitCompleteSignature() []byte {
	if m != nil {
		return m.ExitCompleteSignature
	}
	return nil
}

func (m *ExitCompleted) GetCompleted() time.Time {
	if m != nil {
		return m.Completed
	}
	return time.Time{}
}

type ExitFailed struct {
	// on failure
	ExitFailureSignature []byte            `protobuf:"bytes,1,opt,name=exit_failure_signature,json=exitFailureSignature,proto3" json:"exit_failure_signature,omitempty"`
	Reason               ExitFailed_Reason `protobuf:"varint,2,opt,name=reason,proto3,enum=gracefulexit.ExitFailed_Reason" json:"reason,omitempty"`
	// satellite who issued this exit failed
	SatelliteId NodeID `protobuf:"bytes,3,opt,name=satellite_id,json=satelliteId,proto3,customtype=NodeID" json:"satellite_id"`
	// storage node this exit failed was issued to
	NodeId NodeID `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	// timestamp when the exit failed
	Failed               time.Time `protobuf:"bytes,5,opt,name=failed,proto3,stdtime" json:"failed"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ExitFailed) Reset()         { *m = ExitFailed{} }
func (m *ExitFailed) String() string { return proto.CompactTextString(m) }
func (*ExitFailed) ProtoMessage()    {}
func (*ExitFailed) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{9}
}
func (m *ExitFailed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitFailed.Unmarshal(m, b)
}
func (m *ExitFailed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitFailed.Marshal(b, m, deterministic)
}
func (m *ExitFailed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitFailed.Merge(m, src)
}
func (m *ExitFailed) XXX_Size() int {
	return xxx_messageInfo_ExitFailed.Size(m)
}
func (m *ExitFailed) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitFailed.DiscardUnknown(m)
}

var xxx_messageInfo_ExitFailed proto.InternalMessageInfo

func (m *ExitFailed) GetExitFailureSignature() []byte {
	if m != nil {
		return m.ExitFailureSignature
	}
	return nil
}

func (m *ExitFailed) GetReason() ExitFailed_Reason {
	if m != nil {
		return m.Reason
	}
	return ExitFailed_VERIFICATION_FAILED
}

func (m *ExitFailed) GetFailed() time.Time {
	if m != nil {
		return m.Failed
	}
	return time.Time{}
}

type SatelliteMessage struct {
	// Types that are valid to be assigned to Message:
	//	*SatelliteMessage_NotReady
	//	*SatelliteMessage_TransferPiece
	//	*SatelliteMessage_DeletePiece
	//	*SatelliteMessage_ExitCompleted
	//	*SatelliteMessage_ExitFailed
	Message              isSatelliteMessage_Message `protobuf_oneof:"Message"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SatelliteMessage) Reset()         { *m = SatelliteMessage{} }
func (m *SatelliteMessage) String() string { return proto.CompactTextString(m) }
func (*SatelliteMessage) ProtoMessage()    {}
func (*SatelliteMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f0acbf2ce5fa631, []int{10}
}
func (m *SatelliteMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SatelliteMessage.Unmarshal(m, b)
}
func (m *SatelliteMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SatelliteMessage.Marshal(b, m, deterministic)
}
func (m *SatelliteMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SatelliteMessage.Merge(m, src)
}
func (m *SatelliteMessage) XXX_Size() int {
	return xxx_messageInfo_SatelliteMessage.Size(m)
}
func (m *SatelliteMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SatelliteMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SatelliteMessage proto.InternalMessageInfo

type isSatelliteMessage_Message interface {
	isSatelliteMessage_Message()
}

type SatelliteMessage_NotReady struct {
	NotReady *NotReady `protobuf:"bytes,1,opt,name=not_ready,json=notReady,proto3,oneof"`
}
type SatelliteMessage_TransferPiece struct {
	TransferPiece *TransferPiece `protobuf:"bytes,2,opt,name=transfer_piece,json=transferPiece,proto3,oneof"`
}
type SatelliteMessage_DeletePiece struct {
	DeletePiece *DeletePiece `protobuf:"bytes,3,opt,name=delete_piece,json=deletePiece,proto3,oneof"`
}
type SatelliteMessage_ExitCompleted struct {
	ExitCompleted *ExitCompleted `protobuf:"bytes,4,opt,name=exit_completed,json=exitCompleted,proto3,oneof"`
}
type SatelliteMessage_ExitFailed struct {
	ExitFailed *ExitFailed `protobuf:"bytes,5,opt,name=exit_failed,json=exitFailed,proto3,oneof"`
}

func (*SatelliteMessage_NotReady) isSatelliteMessage_Message()      {}
func (*SatelliteMessage_TransferPiece) isSatelliteMessage_Message() {}
func (*SatelliteMessage_DeletePiece) isSatelliteMessage_Message()   {}
func (*SatelliteMessage_ExitCompleted) isSatelliteMessage_Message() {}
func (*SatelliteMessage_ExitFailed) isSatelliteMessage_Message()    {}

func (m *SatelliteMessage) GetMessage() isSatelliteMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *SatelliteMessage) GetNotReady() *NotReady {
	if x, ok := m.GetMessage().(*SatelliteMessage_NotReady); ok {
		return x.NotReady
	}
	return nil
}

func (m *SatelliteMessage) GetTransferPiece() *TransferPiece {
	if x, ok := m.GetMessage().(*SatelliteMessage_TransferPiece); ok {
		return x.TransferPiece
	}
	return nil
}

func (m *SatelliteMessage) GetDeletePiece() *DeletePiece {
	if x, ok := m.GetMessage().(*SatelliteMessage_DeletePiece); ok {
		return x.DeletePiece
	}
	return nil
}

func (m *SatelliteMessage) GetExitCompleted() *ExitCompleted {
	if x, ok := m.GetMessage().(*SatelliteMessage_ExitCompleted); ok {
		return x.ExitCompleted
	}
	return nil
}

func (m *SatelliteMessage) GetExitFailed() *ExitFailed {
	if x, ok := m.GetMessage().(*SatelliteMessage_ExitFailed); ok {
		return x.ExitFailed
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SatelliteMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SatelliteMessage_OneofMarshaler, _SatelliteMessage_OneofUnmarshaler, _SatelliteMessage_OneofSizer, []interface{}{
		(*SatelliteMessage_NotReady)(nil),
		(*SatelliteMessage_TransferPiece)(nil),
		(*SatelliteMessage_DeletePiece)(nil),
		(*SatelliteMessage_ExitCompleted)(nil),
		(*SatelliteMessage_ExitFailed)(nil),
	}
}

func _SatelliteMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SatelliteMessage)
	// Message
	switch x := m.Message.(type) {
	case *SatelliteMessage_NotReady:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NotReady); err != nil {
			return err
		}
	case *SatelliteMessage_TransferPiece:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TransferPiece); err != nil {
			return err
		}
	case *SatelliteMessage_DeletePiece:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DeletePiece); err != nil {
			return err
		}
	case *SatelliteMessage_ExitCompleted:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExitCompleted); err != nil {
			return err
		}
	case *SatelliteMessage_ExitFailed:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExitFailed); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SatelliteMessage.Message has unexpected type %T", x)
	}
	return nil
}

func _SatelliteMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SatelliteMessage)
	switch tag {
	case 1: // Message.not_ready
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NotReady)
		err := b.DecodeMessage(msg)
		m.Message = &SatelliteMessage_NotReady{msg}
		return true, err
	case 2: // Message.transfer_piece
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransferPiece)
		err := b.DecodeMessage(msg)
		m.Message = &SatelliteMessage_TransferPiece{msg}
		return true, err
	case 3: // Message.delete_piece
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeletePiece)
		err := b.DecodeMessage(msg)
		m.Message = &SatelliteMessage_DeletePiece{msg}
		return true, err
	case 4: // Message.exit_completed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ExitCompleted)
		err := b.DecodeMessage(msg)
		m.Message = &SatelliteMessage_ExitCompleted{msg}
		return true, err
	case 5: // Message.exit_failed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ExitFailed)
		err := b.DecodeMessage(msg)
		m.Message = &SatelliteMessage_ExitFailed{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SatelliteMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SatelliteMessage)
	// Message
	switch x := m.Message.(type) {
	case *SatelliteMessage_NotReady:
		s := proto.Size(x.NotReady)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SatelliteMessage_TransferPiece:
		s := proto.Size(x.TransferPiece)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SatelliteMessage_DeletePiece:
		s := proto.Size(x.DeletePiece)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SatelliteMessage_ExitCompleted:
		s := proto.Size(x.ExitCompleted)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SatelliteMessage_ExitFailed:
		s := proto.Size(x.ExitFailed)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterEnum("gracefulexit.TransferFailed_Error", TransferFailed_Error_name, TransferFailed_Error_value)
	proto.RegisterEnum("gracefulexit.ExitFailed_Reason", ExitFailed_Reason_name, ExitFailed_Reason_value)
	proto.RegisterType((*GracefulExitFeasibilityRequest)(nil), "gracefulexit.GracefulExitFeasibilityRequest")
	proto.RegisterType((*GracefulExitFeasibilityResponse)(nil), "gracefulexit.GracefulExitFeasibilityResponse")
	proto.RegisterType((*TransferSucceeded)(nil), "gracefulexit.TransferSucceeded")
	proto.RegisterType((*TransferFailed)(nil), "gracefulexit.TransferFailed")
	proto.RegisterType((*StorageNodeMessage)(nil), "gracefulexit.StorageNodeMessage")
	proto.RegisterType((*NotReady)(nil), "gracefulexit.NotReady")
	proto.RegisterType((*TransferPiece)(nil), "gracefulexit.TransferPiece")
	proto.RegisterType((*DeletePiece)(nil), "gracefulexit.DeletePiece")
	proto.RegisterType((*ExitCompleted)(nil), "gracefulexit.ExitCompleted")
	proto.RegisterType((*ExitFailed)(nil), "gracefulexit.ExitFailed")
	proto.RegisterType((*SatelliteMessage)(nil), "gracefulexit.SatelliteMessage")
}

func init() { proto.RegisterFile("gracefulexit.proto", fileDescriptor_8f0acbf2ce5fa631) }

var fileDescriptor_8f0acbf2ce5fa631 = []byte{
	// 1073 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcd, 0x52, 0x1b, 0x47,
	0x10, 0xd6, 0x82, 0x11, 0xa8, 0x25, 0x84, 0x18, 0xfe, 0x14, 0xb0, 0x2d, 0xd5, 0xe6, 0x00, 0x87,
	0x44, 0x24, 0xe4, 0xb7, 0xca, 0xa9, 0xa4, 0x56, 0x68, 0x85, 0x36, 0x16, 0x2b, 0x32, 0x12, 0xc4,
	0x95, 0xcb, 0xd6, 0xa2, 0x6d, 0xc4, 0xda, 0xab, 0x1d, 0x79, 0x66, 0xe4, 0x98, 0x6b, 0x9e, 0x20,
	0xc7, 0x9c, 0x73, 0xce, 0x83, 0xe4, 0x15, 0x92, 0xaa, 0x38, 0xc9, 0x29, 0xaf, 0x91, 0xda, 0x5f,
	0x24, 0xfe, 0x0a, 0x97, 0x4f, 0x68, 0x7b, 0xbe, 0x6e, 0xbe, 0xee, 0xe9, 0xef, 0xab, 0x01, 0x32,
	0xe0, 0x76, 0x1f, 0xcf, 0xc6, 0x1e, 0xbe, 0x76, 0x65, 0x6d, 0xc4, 0x99, 0x64, 0xa4, 0x30, 0x19,
	0xdb, 0x84, 0x01, 0x1b, 0xb0, 0xe8, 0x64, 0xb3, 0x32, 0x60, 0x6c, 0xe0, 0xe1, 0x6e, 0xf8, 0x75,
	0x3a, 0x3e, 0xdb, 0x95, 0xee, 0x10, 0x85, 0xb4, 0x87, 0xa3, 0x18, 0x50, 0x1c, 0xa2, 0xb4, 0x5d,
	0xff, 0x2c, 0x49, 0x28, 0x30, 0xee, 0x20, 0x17, 0xd1, 0x97, 0x5a, 0x85, 0xc7, 0x07, 0x71, 0x69,
	0xfd, 0xb5, 0x2b, 0x9b, 0x68, 0x0b, 0xf7, 0xd4, 0xf5, 0x5c, 0x79, 0x41, 0xf1, 0xe5, 0x18, 0x85,
	0x54, 0x7f, 0x53, 0xa0, 0x72, 0x2b, 0x44, 0x8c, 0x98, 0x2f, 0x90, 0x68, 0x90, 0x7b, 0xce, 0x5c,
	0x1f, 0x1d, 0xcb, 0x96, 0x65, 0xa5, 0xaa, 0xec, 0xe4, 0xf7, 0x36, 0x6b, 0x11, 0xb1, 0x5a, 0x42,
	0xac, 0xd6, 0x4b, 0x88, 0xd5, 0x17, 0x7e, 0x7f, 0x53, 0xc9, 0xfc, 0xfc, 0x77, 0x45, 0xa1, 0x0b,
	0x51, 0x9a, 0x26, 0xc9, 0x36, 0x2c, 0x0d, 0x99, 0x2f, 0xcf, 0x85, 0xc5, 0xf1, 0xe5, 0xd8, 0xe5,
	0xe8, 0x94, 0x67, 0xaa, 0xca, 0xce, 0x1c, 0x2d, 0x46, 0x61, 0x1a, 0x47, 0xc9, 0x23, 0x00, 0x57,
	0x58, 0xb6, 0xe7, 0xb1, 0x1f, 0xd1, 0x29, 0xcf, 0x56, 0x95, 0x9d, 0x05, 0x9a, 0x73, 0x85, 0x16,
	0x05, 0xd4, 0x5f, 0x67, 0x60, 0xb9, 0xc7, 0x6d, 0x5f, 0x9c, 0x21, 0xef, 0x8e, 0xfb, 0x7d, 0x44,
	0x07, 0x1d, 0xd2, 0x80, 0x55, 0xc6, 0xdd, 0x81, 0xeb, 0xdb, 0x9e, 0x15, 0xf6, 0x6f, 0x79, 0xee,
	0xd0, 0x4d, 0xb8, 0x92, 0x5a, 0x3c, 0x93, 0x4e, 0xf0, 0xa7, 0x1d, 0x9c, 0x50, 0x92, 0xe0, 0x2f,
	0x63, 0x44, 0x83, 0x95, 0xb4, 0xca, 0xc8, 0xc5, 0x3e, 0x5a, 0xe7, 0xb6, 0x38, 0x0f, 0x79, 0xe6,
	0xf7, 0x96, 0x93, 0x22, 0x47, 0xc1, 0x49, 0xcb, 0x16, 0xe7, 0x74, 0x39, 0x41, 0xa7, 0x21, 0x72,
	0x00, 0xeb, 0x1c, 0x47, 0x9e, 0xdd, 0xc7, 0x21, 0xfa, 0x72, 0xb2, 0xca, 0xec, 0x6d, 0x55, 0x56,
	0x27, 0x12, 0x2e, 0x0b, 0x3d, 0x81, 0xe5, 0x2b, 0x5c, 0x5c, 0xa7, 0xfc, 0xa0, 0xaa, 0xec, 0x14,
	0xea, 0x4b, 0xc1, 0x78, 0xff, 0x7c, 0x53, 0x99, 0x0f, 0xd1, 0x46, 0x83, 0x2e, 0x4d, 0xf1, 0x30,
	0x1c, 0xf5, 0x5f, 0x05, 0x8a, 0xc9, 0x90, 0x9a, 0xb6, 0xeb, 0xa1, 0x73, 0x73, 0x3d, 0xe5, 0x7e,
	0xf5, 0xc8, 0x97, 0x30, 0x87, 0x9c, 0x33, 0x1e, 0x8e, 0xa2, 0xb8, 0xa7, 0xd6, 0xa6, 0x56, 0x78,
	0xfa, 0x3f, 0xd5, 0xf4, 0x00, 0x49, 0xa3, 0x04, 0xf5, 0x19, 0xcc, 0x85, 0xdf, 0x64, 0x11, 0x72,
	0x66, 0xa7, 0x67, 0x35, 0x3b, 0xc7, 0x66, 0xa3, 0x94, 0x21, 0x0f, 0xa1, 0xdc, 0xed, 0x75, 0xa8,
	0x76, 0xa0, 0x5b, 0x66, 0xa7, 0xa1, 0x5b, 0xc7, 0xa6, 0x76, 0xa2, 0x19, 0x6d, 0xad, 0xde, 0xd6,
	0x4b, 0x0a, 0x59, 0x83, 0xe5, 0x96, 0xd6, 0x6d, 0x59, 0x27, 0x3a, 0x35, 0x9a, 0xc6, 0xbe, 0xd6,
	0x33, 0x3a, 0x66, 0x69, 0x86, 0xe4, 0x61, 0xfe, 0xd8, 0x7c, 0x6a, 0x76, 0xbe, 0x37, 0x4b, 0xa0,
	0xfe, 0xa2, 0x00, 0xe9, 0x4a, 0xc6, 0xed, 0x01, 0x9a, 0xcc, 0xc1, 0x43, 0x14, 0xc2, 0x1e, 0x20,
	0xf9, 0x06, 0x72, 0x22, 0x59, 0x8b, 0xf8, 0xfa, 0x2b, 0x37, 0xd3, 0x4d, 0xb7, 0xa7, 0x95, 0xa1,
	0x97, 0x39, 0xe4, 0x73, 0xc8, 0x9e, 0x85, 0x8d, 0xc4, 0xf7, 0xfe, 0xf0, 0xae, 0x66, 0x5b, 0x19,
	0x1a, 0xa3, 0xeb, 0x39, 0x98, 0x8f, 0x39, 0xa8, 0x00, 0x0b, 0x26, 0x93, 0x14, 0x6d, 0xe7, 0x42,
	0xfd, 0x43, 0x81, 0xc5, 0x24, 0x27, 0x1c, 0xe7, 0xbb, 0xde, 0x44, 0x7e, 0xc4, 0xdd, 0x57, 0xb6,
	0x44, 0xeb, 0x05, 0x5e, 0x84, 0x14, 0x0b, 0xf5, 0x8d, 0x38, 0x6d, 0x29, 0x44, 0x1d, 0x45, 0xe7,
	0x4f, 0xf1, 0x82, 0xc2, 0x28, 0xfd, 0x4d, 0xbe, 0x83, 0x35, 0xdb, 0x71, 0x38, 0x0a, 0x81, 0xce,
	0x94, 0x46, 0xa2, 0xc5, 0x7c, 0x54, 0x4b, 0x7d, 0x44, 0x4b, 0x60, 0x13, 0x72, 0x59, 0xb1, 0xaf,
	0x07, 0xd5, 0x6f, 0x21, 0xdf, 0x40, 0x0f, 0x25, 0xbe, 0x7b, 0x63, 0xea, 0x7f, 0x0a, 0x2c, 0x06,
	0xf6, 0xb3, 0xcf, 0x86, 0xa3, 0xa0, 0x64, 0x70, 0x11, 0x1b, 0xc1, 0xc4, 0xad, 0x7e, 0x1c, 0xb1,
	0x84, 0x3b, 0xf0, 0x6d, 0x39, 0xe6, 0x18, 0x15, 0xa5, 0x6b, 0x38, 0x81, 0xef, 0x26, 0x87, 0xe4,
	0x63, 0x28, 0x08, 0x5b, 0xa2, 0xe7, 0xb9, 0x32, 0x64, 0x10, 0xcd, 0xa8, 0x18, 0x33, 0xc8, 0x06,
	0xcb, 0x62, 0x34, 0x68, 0x3e, 0xc5, 0x18, 0x0e, 0xd9, 0x86, 0x79, 0x9f, 0x39, 0x21, 0x7a, 0xf6,
	0x46, 0x74, 0x36, 0x38, 0x36, 0x1c, 0x52, 0x87, 0x5c, 0x42, 0x27, 0x52, 0xe3, 0x7d, 0x8d, 0xf0,
	0x32, 0x4d, 0xfd, 0x69, 0x16, 0x20, 0x34, 0xda, 0x48, 0x98, 0x9f, 0xc2, 0x7a, 0xd8, 0x66, 0xb0,
	0x46, 0x63, 0x7e, 0xbd, 0xcb, 0x55, 0x8c, 0xb1, 0x63, 0x3e, 0xd1, 0xe4, 0x17, 0x90, 0xe5, 0x68,
	0x0b, 0xe6, 0xc7, 0x92, 0xbc, 0xb2, 0xe3, 0x97, 0xf5, 0x6b, 0x34, 0x84, 0xd1, 0x18, 0x7e, 0x6d,
	0x3a, 0xb3, 0x6f, 0x35, 0x9d, 0x07, 0x77, 0x4e, 0xe7, 0xab, 0x54, 0x3a, 0x73, 0x6f, 0x31, 0x9a,
	0x38, 0x47, 0x7d, 0x01, 0xd9, 0x88, 0x2b, 0xd9, 0x80, 0x95, 0x49, 0xe5, 0x5b, 0x4d, 0xcd, 0x68,
	0xeb, 0x81, 0x6b, 0x54, 0x60, 0xcb, 0x30, 0xb5, 0xfd, 0x9e, 0x71, 0xa2, 0x5b, 0x3d, 0xe3, 0x50,
	0x6f, 0x52, 0xed, 0x50, 0xb7, 0xf4, 0x67, 0xfb, 0xba, 0xde, 0xd0, 0x1b, 0x25, 0x85, 0x6c, 0xc3,
	0xfb, 0x9d, 0x13, 0x9d, 0x6a, 0xed, 0x76, 0x98, 0x74, 0x4c, 0x75, 0xeb, 0x48, 0xa7, 0xfb, 0xba,
	0xd9, 0x0b, 0x9c, 0x26, 0x05, 0xce, 0xa8, 0x7f, 0xcd, 0x40, 0xa9, 0x9b, 0xf4, 0x98, 0x78, 0xc7,
	0x67, 0x90, 0xf3, 0x99, 0xb4, 0x78, 0x20, 0xdc, 0xd8, 0x3b, 0xd6, 0xa7, 0xe7, 0x9a, 0xc8, 0xba,
	0x95, 0xa1, 0x0b, 0x7e, 0xfc, 0x9b, 0x34, 0xa0, 0x28, 0x63, 0x85, 0x47, 0x7b, 0x1f, 0x3b, 0xc7,
	0xd6, 0xcd, 0xce, 0x11, 0x79, 0x7c, 0x86, 0x2e, 0xca, 0x29, 0x5b, 0xf8, 0x1a, 0x0a, 0x4e, 0x28,
	0xa6, 0xb8, 0x46, 0x24, 0xcb, 0xf7, 0xa6, 0x6b, 0x4c, 0xc8, 0xad, 0x95, 0xa1, 0x79, 0x67, 0x42,
	0x7d, 0x0d, 0x28, 0x4e, 0xc9, 0x25, 0xd9, 0xcf, 0xad, 0xeb, 0x9b, 0x91, 0x6a, 0x2c, 0x60, 0x81,
	0x53, 0xa2, 0x7b, 0x02, 0xf9, 0x74, 0x1b, 0xd3, 0x7b, 0x2c, 0xdf, 0xb6, 0x5c, 0xad, 0x0c, 0x05,
	0x4c, 0xbf, 0x26, 0x2c, 0x70, 0xef, 0x1f, 0x05, 0xd6, 0xd2, 0xf9, 0x4e, 0x3e, 0x2f, 0x48, 0x07,
	0xe6, 0x8f, 0x38, 0xeb, 0xa3, 0x10, 0xa4, 0x3a, 0x5d, 0xf7, 0xba, 0x9b, 0x6f, 0x3e, 0xbe, 0x82,
	0xb8, 0x72, 0x63, 0x3b, 0xca, 0x47, 0x0a, 0x79, 0x05, 0x1b, 0xb7, 0xbc, 0x5f, 0xc8, 0x07, 0xd3,
	0xe9, 0x77, 0xbf, 0x84, 0x36, 0x3f, 0xbc, 0x27, 0x3a, 0x7a, 0x14, 0xd5, 0x57, 0x7f, 0x20, 0x42,
	0x32, 0xfe, 0xbc, 0xe6, 0xb2, 0xdd, 0x3e, 0x1b, 0x0e, 0x99, 0xbf, 0x3b, 0x3a, 0x3d, 0xcd, 0x86,
	0xbb, 0xfe, 0xc9, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x41, 0x4e, 0x8a, 0xe6, 0x09, 0x00,
	0x00,
}
