// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/group/v1beta1/events.proto

package group

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// EventCreateGroup is an event emitted when a group is created.
type EventCreateGroup struct {
	// group_id is the unique ID of the group.
	GroupId uint64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (m *EventCreateGroup) Reset()         { *m = EventCreateGroup{} }
func (m *EventCreateGroup) String() string { return proto.CompactTextString(m) }
func (*EventCreateGroup) ProtoMessage()    {}
func (*EventCreateGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_7879e051fb126fc0, []int{0}
}
func (m *EventCreateGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateGroup.Merge(m, src)
}
func (m *EventCreateGroup) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateGroup.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateGroup proto.InternalMessageInfo

func (m *EventCreateGroup) GetGroupId() uint64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

// EventUpdateGroup is an event emitted when a group is updated.
type EventUpdateGroup struct {
	// group_id is the unique ID of the group.
	GroupId uint64 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (m *EventUpdateGroup) Reset()         { *m = EventUpdateGroup{} }
func (m *EventUpdateGroup) String() string { return proto.CompactTextString(m) }
func (*EventUpdateGroup) ProtoMessage()    {}
func (*EventUpdateGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_7879e051fb126fc0, []int{1}
}
func (m *EventUpdateGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateGroup.Merge(m, src)
}
func (m *EventUpdateGroup) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateGroup.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateGroup proto.InternalMessageInfo

func (m *EventUpdateGroup) GetGroupId() uint64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

// EventCreateGroupPolicy is an event emitted when a group policy is created.
type EventCreateGroupPolicy struct {
	// address is the account address of the group policy.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *EventCreateGroupPolicy) Reset()         { *m = EventCreateGroupPolicy{} }
func (m *EventCreateGroupPolicy) String() string { return proto.CompactTextString(m) }
func (*EventCreateGroupPolicy) ProtoMessage()    {}
func (*EventCreateGroupPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_7879e051fb126fc0, []int{2}
}
func (m *EventCreateGroupPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateGroupPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateGroupPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateGroupPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateGroupPolicy.Merge(m, src)
}
func (m *EventCreateGroupPolicy) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateGroupPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateGroupPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateGroupPolicy proto.InternalMessageInfo

func (m *EventCreateGroupPolicy) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// EventUpdateGroupPolicy is an event emitted when a group policy is updated.
type EventUpdateGroupPolicy struct {
	// address is the account address of the group policy.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *EventUpdateGroupPolicy) Reset()         { *m = EventUpdateGroupPolicy{} }
func (m *EventUpdateGroupPolicy) String() string { return proto.CompactTextString(m) }
func (*EventUpdateGroupPolicy) ProtoMessage()    {}
func (*EventUpdateGroupPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_7879e051fb126fc0, []int{3}
}
func (m *EventUpdateGroupPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateGroupPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateGroupPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateGroupPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateGroupPolicy.Merge(m, src)
}
func (m *EventUpdateGroupPolicy) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateGroupPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateGroupPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateGroupPolicy proto.InternalMessageInfo

func (m *EventUpdateGroupPolicy) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// EventSubmitProposal is an event emitted when a proposal is created.
type EventSubmitProposal struct {
	// proposal_id is the unique ID of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
}

func (m *EventSubmitProposal) Reset()         { *m = EventSubmitProposal{} }
func (m *EventSubmitProposal) String() string { return proto.CompactTextString(m) }
func (*EventSubmitProposal) ProtoMessage()    {}
func (*EventSubmitProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_7879e051fb126fc0, []int{4}
}
func (m *EventSubmitProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSubmitProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSubmitProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSubmitProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSubmitProposal.Merge(m, src)
}
func (m *EventSubmitProposal) XXX_Size() int {
	return m.Size()
}
func (m *EventSubmitProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSubmitProposal.DiscardUnknown(m)
}

var xxx_messageInfo_EventSubmitProposal proto.InternalMessageInfo

func (m *EventSubmitProposal) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

// EventWithdrawProposal is an event emitted when a proposal is withdrawn.
type EventWithdrawProposal struct {
	// proposal_id is the unique ID of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
}

func (m *EventWithdrawProposal) Reset()         { *m = EventWithdrawProposal{} }
func (m *EventWithdrawProposal) String() string { return proto.CompactTextString(m) }
func (*EventWithdrawProposal) ProtoMessage()    {}
func (*EventWithdrawProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_7879e051fb126fc0, []int{5}
}
func (m *EventWithdrawProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventWithdrawProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventWithdrawProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventWithdrawProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventWithdrawProposal.Merge(m, src)
}
func (m *EventWithdrawProposal) XXX_Size() int {
	return m.Size()
}
func (m *EventWithdrawProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_EventWithdrawProposal.DiscardUnknown(m)
}

var xxx_messageInfo_EventWithdrawProposal proto.InternalMessageInfo

func (m *EventWithdrawProposal) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

// EventVote is an event emitted when a voter votes on a proposal.
type EventVote struct {
	// proposal_id is the unique ID of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
}

func (m *EventVote) Reset()         { *m = EventVote{} }
func (m *EventVote) String() string { return proto.CompactTextString(m) }
func (*EventVote) ProtoMessage()    {}
func (*EventVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_7879e051fb126fc0, []int{6}
}
func (m *EventVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventVote.Merge(m, src)
}
func (m *EventVote) XXX_Size() int {
	return m.Size()
}
func (m *EventVote) XXX_DiscardUnknown() {
	xxx_messageInfo_EventVote.DiscardUnknown(m)
}

var xxx_messageInfo_EventVote proto.InternalMessageInfo

func (m *EventVote) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

// EventExec is an event emitted when a proposal is executed.
type EventExec struct {
	// proposal_id is the unique ID of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
}

func (m *EventExec) Reset()         { *m = EventExec{} }
func (m *EventExec) String() string { return proto.CompactTextString(m) }
func (*EventExec) ProtoMessage()    {}
func (*EventExec) Descriptor() ([]byte, []int) {
	return fileDescriptor_7879e051fb126fc0, []int{7}
}
func (m *EventExec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventExec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventExec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventExec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventExec.Merge(m, src)
}
func (m *EventExec) XXX_Size() int {
	return m.Size()
}
func (m *EventExec) XXX_DiscardUnknown() {
	xxx_messageInfo_EventExec.DiscardUnknown(m)
}

var xxx_messageInfo_EventExec proto.InternalMessageInfo

func (m *EventExec) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

func init() {
	proto.RegisterType((*EventCreateGroup)(nil), "cosmos.group.v1beta1.EventCreateGroup")
	proto.RegisterType((*EventUpdateGroup)(nil), "cosmos.group.v1beta1.EventUpdateGroup")
	proto.RegisterType((*EventCreateGroupPolicy)(nil), "cosmos.group.v1beta1.EventCreateGroupPolicy")
	proto.RegisterType((*EventUpdateGroupPolicy)(nil), "cosmos.group.v1beta1.EventUpdateGroupPolicy")
	proto.RegisterType((*EventSubmitProposal)(nil), "cosmos.group.v1beta1.EventSubmitProposal")
	proto.RegisterType((*EventWithdrawProposal)(nil), "cosmos.group.v1beta1.EventWithdrawProposal")
	proto.RegisterType((*EventVote)(nil), "cosmos.group.v1beta1.EventVote")
	proto.RegisterType((*EventExec)(nil), "cosmos.group.v1beta1.EventExec")
}

func init() { proto.RegisterFile("cosmos/group/v1beta1/events.proto", fileDescriptor_7879e051fb126fc0) }

var fileDescriptor_7879e051fb126fc0 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0x2f, 0xca, 0x2f, 0x2d, 0xd0, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34,
	0xd4, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x81,
	0x28, 0xd1, 0x03, 0x2b, 0xd1, 0x83, 0x2a, 0x91, 0x92, 0x84, 0x88, 0xc6, 0x83, 0xd5, 0xe8, 0x43,
	0x95, 0x80, 0x39, 0x4a, 0xba, 0x5c, 0x02, 0xae, 0x20, 0x03, 0x9c, 0x8b, 0x52, 0x13, 0x4b, 0x52,
	0xdd, 0x41, 0xda, 0x84, 0x24, 0xb9, 0x38, 0xc0, 0xfa, 0xe3, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x58, 0x82, 0xd8, 0xc1, 0x7c, 0xcf, 0x14, 0xb8, 0xf2, 0xd0, 0x82, 0x14, 0x62, 0x94, 0xfb,
	0x70, 0x89, 0xa1, 0x9b, 0x1e, 0x90, 0x9f, 0x93, 0x99, 0x5c, 0x29, 0x64, 0xc4, 0xc5, 0x9e, 0x98,
	0x92, 0x52, 0x94, 0x5a, 0x5c, 0x0c, 0xd6, 0xc3, 0xe9, 0x24, 0x71, 0x69, 0x8b, 0x2e, 0xcc, 0xf5,
	0x8e, 0x10, 0x99, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0xf4, 0x20, 0x98, 0x42, 0xb8, 0x69, 0x48, 0x96,
	0x53, 0x60, 0x9a, 0x19, 0x97, 0x30, 0xd8, 0xb4, 0xe0, 0xd2, 0xa4, 0xdc, 0xcc, 0x92, 0x80, 0xa2,
	0xfc, 0x82, 0xfc, 0xe2, 0xc4, 0x1c, 0x21, 0x79, 0x2e, 0xee, 0x02, 0x28, 0x1b, 0xe1, 0x21, 0x2e,
	0x98, 0x90, 0x67, 0x8a, 0x92, 0x05, 0x97, 0x28, 0x58, 0x5f, 0x78, 0x66, 0x49, 0x46, 0x4a, 0x51,
	0x62, 0x39, 0xf1, 0x3a, 0x75, 0xb8, 0x38, 0xc1, 0x3a, 0xc3, 0xf2, 0x4b, 0x52, 0x89, 0x57, 0xed,
	0x5a, 0x91, 0x9a, 0x4c, 0x50, 0xb5, 0x93, 0xdd, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb,
	0x31, 0x44, 0xa9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0x42, 0xa3, 0x1e,
	0x4a, 0xe9, 0x16, 0xa7, 0x64, 0xeb, 0x57, 0x40, 0x52, 0x53, 0x12, 0x1b, 0x38, 0x39, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x54, 0x9c, 0xae, 0x89, 0x64, 0x02, 0x00, 0x00,
}

func (m *EventCreateGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GroupId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.GroupId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GroupId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.GroupId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventCreateGroupPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateGroupPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateGroupPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateGroupPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateGroupPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateGroupPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventSubmitProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSubmitProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSubmitProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposalId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventWithdrawProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventWithdrawProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventWithdrawProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposalId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposalId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventExec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventExec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventExec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposalId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreateGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GroupId != 0 {
		n += 1 + sovEvents(uint64(m.GroupId))
	}
	return n
}

func (m *EventUpdateGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GroupId != 0 {
		n += 1 + sovEvents(uint64(m.GroupId))
	}
	return n
}

func (m *EventCreateGroupPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventUpdateGroupPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventSubmitProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovEvents(uint64(m.ProposalId))
	}
	return n
}

func (m *EventWithdrawProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovEvents(uint64(m.ProposalId))
	}
	return n
}

func (m *EventVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovEvents(uint64(m.ProposalId))
	}
	return n
}

func (m *EventExec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovEvents(uint64(m.ProposalId))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreateGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreateGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			m.GroupId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUpdateGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUpdateGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			m.GroupId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventCreateGroupPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreateGroupPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateGroupPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUpdateGroupPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUpdateGroupPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateGroupPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventSubmitProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSubmitProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSubmitProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventWithdrawProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventWithdrawProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventWithdrawProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventExec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventExec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventExec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
