// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/claim/v1/claim.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Action defines the list of available actions to claim the airdrop tokens.
type Action int32

const (
	// UNSPECIFIED defines an invalid action.
	ActionUnspecified Action = 0
	// VOTE defines a proposal vote.
	ActionVote Action = 1
	// DELEGATE defines an staking delegation.
	ActionDelegate Action = 2
	// EVM defines an EVM transaction.
	ActionEVM Action = 3
	// IBC Transfer defines a fungible token transfer transaction via IBC.
	ActionIBCTransfer Action = 4
)

var Action_name = map[int32]string{
	0: "ACTION_UNSPECIFIED",
	1: "ACTION_VOTE",
	2: "ACTION_DELEGATE",
	3: "ACTION_EVM",
	4: "ACTION_IBC_TRANSFER",
}

var Action_value = map[string]int32{
	"ACTION_UNSPECIFIED":  0,
	"ACTION_VOTE":         1,
	"ACTION_DELEGATE":     2,
	"ACTION_EVM":          3,
	"ACTION_IBC_TRANSFER": 4,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_873e678807803576, []int{0}
}

// ActionCompleted marks defines if a given action is completed for the user
type ActionCompleted struct {
	// action enum
	Action Action `protobuf:"varint,1,opt,name=action,proto3,enum=evmos.claim.v1.Action" json:"action,omitempty"`
	// true if the action has been completed
	Completed bool `protobuf:"varint,2,opt,name=completed,proto3" json:"completed,omitempty"`
}

func (m *ActionCompleted) Reset()         { *m = ActionCompleted{} }
func (m *ActionCompleted) String() string { return proto.CompactTextString(m) }
func (*ActionCompleted) ProtoMessage()    {}
func (*ActionCompleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_873e678807803576, []int{0}
}
func (m *ActionCompleted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActionCompleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActionCompleted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActionCompleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionCompleted.Merge(m, src)
}
func (m *ActionCompleted) XXX_Size() int {
	return m.Size()
}
func (m *ActionCompleted) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionCompleted.DiscardUnknown(m)
}

var xxx_messageInfo_ActionCompleted proto.InternalMessageInfo

func (m *ActionCompleted) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return ActionUnspecified
}

func (m *ActionCompleted) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

// ClaimRecordAddress is the metadata of claim data per address
type ClaimRecordAddress struct {
	// bech32 or hex address of claim user
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// total initial claimable amount for the user
	InitialClaimableAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=initial_claimable_amount,json=initialClaimableAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"initial_claimable_amount"`
	// slice of the available actions completed
	ActionsCompleted []bool `protobuf:"varint,3,rep,packed,name=actions_completed,json=actionsCompleted,proto3" json:"actions_completed,omitempty"`
}

func (m *ClaimRecordAddress) Reset()         { *m = ClaimRecordAddress{} }
func (m *ClaimRecordAddress) String() string { return proto.CompactTextString(m) }
func (*ClaimRecordAddress) ProtoMessage()    {}
func (*ClaimRecordAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_873e678807803576, []int{1}
}
func (m *ClaimRecordAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecordAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecordAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecordAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecordAddress.Merge(m, src)
}
func (m *ClaimRecordAddress) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecordAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecordAddress.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecordAddress proto.InternalMessageInfo

func (m *ClaimRecordAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClaimRecordAddress) GetActionsCompleted() []bool {
	if m != nil {
		return m.ActionsCompleted
	}
	return nil
}

// ClaimRecord defines the initial claimable airdrop amount and the list of
// completed actions to claim the tokens.
type ClaimRecord struct {
	// total initial claimable amount for the user
	InitialClaimableAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=initial_claimable_amount,json=initialClaimableAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"initial_claimable_amount"`
	// slice of the available actions completed
	ActionsCompleted []bool `protobuf:"varint,2,rep,packed,name=actions_completed,json=actionsCompleted,proto3" json:"actions_completed,omitempty"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_873e678807803576, []int{2}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func (m *ClaimRecord) GetActionsCompleted() []bool {
	if m != nil {
		return m.ActionsCompleted
	}
	return nil
}

func init() {
	proto.RegisterEnum("evmos.claim.v1.Action", Action_name, Action_value)
	proto.RegisterType((*ActionCompleted)(nil), "evmos.claim.v1.ActionCompleted")
	proto.RegisterType((*ClaimRecordAddress)(nil), "evmos.claim.v1.ClaimRecordAddress")
	proto.RegisterType((*ClaimRecord)(nil), "evmos.claim.v1.ClaimRecord")
}

func init() { proto.RegisterFile("evmos/claim/v1/claim.proto", fileDescriptor_873e678807803576) }

var fileDescriptor_873e678807803576 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xbd, 0x49, 0x95, 0xaf, 0xd9, 0xea, 0x4b, 0xdd, 0x05, 0xaa, 0xc8, 0x02, 0xc7, 0xe4,
	0x00, 0x11, 0xa8, 0xb6, 0x02, 0x2f, 0x80, 0xe3, 0xb8, 0xc8, 0x12, 0x4d, 0x91, 0x9b, 0xe6, 0xc0,
	0xc5, 0xda, 0xd8, 0xdb, 0x64, 0xd5, 0xc4, 0x1b, 0x79, 0xb7, 0x11, 0xbc, 0x01, 0xca, 0x89, 0x17,
	0xc8, 0x09, 0x5e, 0x05, 0xa9, 0xc7, 0x9e, 0x10, 0xe2, 0x50, 0xa1, 0xe4, 0x45, 0x50, 0xbc, 0x9b,
	0x34, 0x48, 0x88, 0x1b, 0x27, 0x8f, 0xe7, 0xff, 0x9b, 0xd9, 0xff, 0x8c, 0xbd, 0xd0, 0x20, 0xd3,
	0x31, 0xe3, 0x4e, 0x3c, 0xc2, 0x74, 0xec, 0x4c, 0x9b, 0x32, 0xb0, 0x27, 0x19, 0x13, 0x0c, 0x55,
	0x72, 0xcd, 0x96, 0xa9, 0x69, 0xd3, 0xb8, 0x3f, 0x60, 0x03, 0x96, 0x4b, 0xce, 0x2a, 0x92, 0x94,
	0x61, 0xc6, 0x8c, 0xaf, 0x5a, 0xf4, 0x31, 0x27, 0xce, 0xb4, 0xd9, 0x27, 0x02, 0x37, 0x9d, 0x98,
	0xd1, 0x54, 0xe9, 0x8f, 0x37, 0x7a, 0x7a, 0xb9, 0xd1, 0x07, 0x24, 0x25, 0x9c, 0x72, 0x89, 0xd4,
	0x23, 0xb8, 0xef, 0xc6, 0x82, 0xb2, 0xd4, 0x63, 0xe3, 0xc9, 0x88, 0x08, 0x92, 0x20, 0x1b, 0x96,
	0x70, 0x9e, 0xaa, 0x02, 0x0b, 0x34, 0x2a, 0x2f, 0x0e, 0xed, 0xdf, 0xcd, 0xd8, 0xb2, 0x20, 0x54,
	0x14, 0x7a, 0x08, 0xcb, 0xf1, 0xba, 0xb8, 0x5a, 0xb0, 0x40, 0x63, 0x37, 0xbc, 0x4b, 0xd4, 0xbf,
	0x02, 0x88, 0xbc, 0x55, 0x65, 0x48, 0x62, 0x96, 0x25, 0x6e, 0x92, 0x64, 0x84, 0x73, 0x54, 0x85,
	0xff, 0x61, 0x19, 0xe6, 0xa7, 0x94, 0xc3, 0xf5, 0x2b, 0x1a, 0xc2, 0x2a, 0x4d, 0xa9, 0xa0, 0x78,
	0x14, 0xe5, 0x27, 0xe2, 0xfe, 0x88, 0x44, 0x78, 0xcc, 0xae, 0x52, 0x91, 0x77, 0x2f, 0xb7, 0xec,
	0xeb, 0xdb, 0x9a, 0xf6, 0xe3, 0xb6, 0xf6, 0x64, 0x40, 0xc5, 0xf0, 0xaa, 0x6f, 0xc7, 0x6c, 0xec,
	0xa8, 0x49, 0xe5, 0xe3, 0x88, 0x27, 0x97, 0x8e, 0xf8, 0x30, 0x21, 0xdc, 0x0e, 0x52, 0x11, 0x1e,
	0xaa, 0x7e, 0xde, 0xba, 0x9d, 0x9b, 0x77, 0x43, 0xcf, 0xe1, 0x81, 0x1c, 0x81, 0x47, 0x77, 0x03,
	0x14, 0xad, 0x62, 0x63, 0x37, 0xd4, 0x95, 0xb0, 0xd9, 0x4a, 0xfd, 0x0b, 0x80, 0x7b, 0x5b, 0x73,
	0xfc, 0xd5, 0x26, 0xf8, 0xf7, 0x36, 0x0b, 0x7f, 0xb6, 0xf9, 0xec, 0x1b, 0x80, 0x25, 0xf9, 0x7d,
	0xd0, 0x11, 0x44, 0xae, 0xd7, 0x0d, 0x4e, 0x3b, 0xd1, 0x79, 0xe7, 0xec, 0xad, 0xef, 0x05, 0xc7,
	0x81, 0xdf, 0xd6, 0x35, 0xe3, 0xc1, 0x6c, 0x6e, 0x1d, 0x48, 0xe6, 0x3c, 0xe5, 0x13, 0x12, 0xd3,
	0x0b, 0x4a, 0x12, 0x54, 0x83, 0x7b, 0x0a, 0xef, 0x9d, 0x76, 0x7d, 0x1d, 0x18, 0x95, 0xd9, 0xdc,
	0x82, 0x92, 0xeb, 0x31, 0x41, 0xd0, 0x53, 0xb8, 0xaf, 0x80, 0xb6, 0xff, 0xc6, 0x7f, 0xed, 0x76,
	0x7d, 0xbd, 0x60, 0xa0, 0xd9, 0xdc, 0xaa, 0x48, 0xa8, 0x4d, 0x46, 0x64, 0x80, 0x05, 0x41, 0x8f,
	0x20, 0x54, 0xa0, 0xdf, 0x3b, 0xd1, 0x8b, 0xc6, 0xff, 0xb3, 0xb9, 0x55, 0x96, 0x8c, 0xdf, 0x3b,
	0x41, 0x36, 0xbc, 0xa7, 0xe4, 0xa0, 0xe5, 0x45, 0xdd, 0xd0, 0xed, 0x9c, 0x1d, 0xfb, 0xa1, 0xbe,
	0xb3, 0x6d, 0x2c, 0x68, 0x79, 0xdd, 0x0c, 0xa7, 0xfc, 0x82, 0x64, 0xc6, 0xce, 0xc7, 0xcf, 0xa6,
	0xd6, 0x7a, 0x75, 0xbd, 0x30, 0xc1, 0xcd, 0xc2, 0x04, 0x3f, 0x17, 0x26, 0xf8, 0xb4, 0x34, 0xb5,
	0x9b, 0xa5, 0xa9, 0x7d, 0x5f, 0x9a, 0xda, 0xbb, 0xed, 0xfd, 0x8a, 0x21, 0xce, 0x38, 0xe5, 0x8e,
	0xbc, 0x5a, 0xef, 0xd5, 0xe5, 0xca, 0x77, 0xdc, 0x2f, 0xe5, 0x7f, 0xfc, 0xcb, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xe9, 0xf9, 0xdb, 0x08, 0x78, 0x03, 0x00, 0x00,
}

func (m *ActionCompleted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActionCompleted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActionCompleted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Completed {
		i--
		if m.Completed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Action != 0 {
		i = encodeVarintClaim(dAtA, i, uint64(m.Action))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ClaimRecordAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecordAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecordAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionsCompleted) > 0 {
		for iNdEx := len(m.ActionsCompleted) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionsCompleted[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintClaim(dAtA, i, uint64(len(m.ActionsCompleted)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.InitialClaimableAmount.Size()
		i -= size
		if _, err := m.InitialClaimableAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClaim(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionsCompleted) > 0 {
		for iNdEx := len(m.ActionsCompleted) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionsCompleted[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintClaim(dAtA, i, uint64(len(m.ActionsCompleted)))
		i--
		dAtA[i] = 0x12
	}
	{
		size := m.InitialClaimableAmount.Size()
		i -= size
		if _, err := m.InitialClaimableAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClaim(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintClaim(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ActionCompleted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Action != 0 {
		n += 1 + sovClaim(uint64(m.Action))
	}
	if m.Completed {
		n += 2
	}
	return n
}

func (m *ClaimRecordAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	l = m.InitialClaimableAmount.Size()
	n += 1 + l + sovClaim(uint64(l))
	if len(m.ActionsCompleted) > 0 {
		n += 1 + sovClaim(uint64(len(m.ActionsCompleted))) + len(m.ActionsCompleted)*1
	}
	return n
}

func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.InitialClaimableAmount.Size()
	n += 1 + l + sovClaim(uint64(l))
	if len(m.ActionsCompleted) > 0 {
		n += 1 + sovClaim(uint64(len(m.ActionsCompleted))) + len(m.ActionsCompleted)*1
	}
	return n
}

func sovClaim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaim(x uint64) (n int) {
	return sovClaim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ActionCompleted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ActionCompleted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActionCompleted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= Action(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Completed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Completed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClaimRecordAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClaimRecordAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecordAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaimableAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialClaimableAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ActionsCompleted = append(m.ActionsCompleted, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionsCompleted) == 0 {
					m.ActionsCompleted = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ActionsCompleted = append(m.ActionsCompleted, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionsCompleted", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaimableAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialClaimableAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ActionsCompleted = append(m.ActionsCompleted, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionsCompleted) == 0 {
					m.ActionsCompleted = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ActionsCompleted = append(m.ActionsCompleted, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionsCompleted", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipClaim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaim
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthClaim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaim = fmt.Errorf("proto: unexpected end of group")
)