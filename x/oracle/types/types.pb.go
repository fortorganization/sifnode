// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sifnode/oracle/v1/types.proto

package types

import (
	fmt "fmt"
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

// StatusText is an enum used to represent the status of the prophecy
type StatusText int32

const (
	// Default value
	StatusText_STATUS_TEXT_UNSPECIFIED StatusText = 0
	// Pending status
	StatusText_STATUS_TEXT_PENDING StatusText = 1
	// Success status
	StatusText_STATUS_TEXT_SUCCESS StatusText = 2
)

var StatusText_name = map[int32]string{
	0: "STATUS_TEXT_UNSPECIFIED",
	1: "STATUS_TEXT_PENDING",
	2: "STATUS_TEXT_SUCCESS",
}

var StatusText_value = map[string]int32{
	"STATUS_TEXT_UNSPECIFIED": 0,
	"STATUS_TEXT_PENDING":     1,
	"STATUS_TEXT_SUCCESS":     2,
}

func (x StatusText) String() string {
	return proto.EnumName(StatusText_name, int32(x))
}

func (StatusText) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dac1b931484f4203, []int{0}
}

type GenesisState struct {
	AddressWhitelist map[uint32]*ValidatorWhiteList `protobuf:"bytes,1,rep,name=address_whitelist,json=addressWhitelist,proto3" json:"address_whitelist,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AdminAddress     string                         `protobuf:"bytes,2,opt,name=admin_address,json=adminAddress,proto3" json:"admin_address,omitempty"`
	Prophecies       []*Prophecy                    `protobuf:"bytes,3,rep,name=prophecies,proto3" json:"prophecies,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac1b931484f4203, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetAddressWhitelist() map[uint32]*ValidatorWhiteList {
	if m != nil {
		return m.AddressWhitelist
	}
	return nil
}

func (m *GenesisState) GetAdminAddress() string {
	if m != nil {
		return m.AdminAddress
	}
	return ""
}

func (m *GenesisState) GetProphecies() []*Prophecy {
	if m != nil {
		return m.Prophecies
	}
	return nil
}

// Claim contains an arbitrary claim with arbitrary content made by a given
// validator
type Claim struct {
	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	Content          string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *Claim) Reset()         { *m = Claim{} }
func (m *Claim) String() string { return proto.CompactTextString(m) }
func (*Claim) ProtoMessage()    {}
func (*Claim) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac1b931484f4203, []int{1}
}
func (m *Claim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Claim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Claim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Claim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Claim.Merge(m, src)
}
func (m *Claim) XXX_Size() int {
	return m.Size()
}
func (m *Claim) XXX_DiscardUnknown() {
	xxx_messageInfo_Claim.DiscardUnknown(m)
}

var xxx_messageInfo_Claim proto.InternalMessageInfo

func (m *Claim) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Claim) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *Claim) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// Prophecy is what the prophecy becomes when being saved to the database.
//  Tendermint/Amino does not support maps so we must serialize those variables
//  into bytes.
type Prophecy struct {
	Id              []byte     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status          StatusText `protobuf:"varint,2,opt,name=status,proto3,enum=sifnode.oracle.v1.StatusText" json:"status,omitempty"`
	ClaimValidators []string   `protobuf:"bytes,3,rep,name=claim_validators,json=claimValidators,proto3" json:"claim_validators,omitempty"`
}

func (m *Prophecy) Reset()         { *m = Prophecy{} }
func (m *Prophecy) String() string { return proto.CompactTextString(m) }
func (*Prophecy) ProtoMessage()    {}
func (*Prophecy) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac1b931484f4203, []int{2}
}
func (m *Prophecy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Prophecy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Prophecy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Prophecy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prophecy.Merge(m, src)
}
func (m *Prophecy) XXX_Size() int {
	return m.Size()
}
func (m *Prophecy) XXX_DiscardUnknown() {
	xxx_messageInfo_Prophecy.DiscardUnknown(m)
}

var xxx_messageInfo_Prophecy proto.InternalMessageInfo

func (m *Prophecy) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Prophecy) GetStatus() StatusText {
	if m != nil {
		return m.Status
	}
	return StatusText_STATUS_TEXT_UNSPECIFIED
}

func (m *Prophecy) GetClaimValidators() []string {
	if m != nil {
		return m.ClaimValidators
	}
	return nil
}

// ValidatorWhiteList is struct that contains validator and its voting power
type ValidatorWhiteList struct {
	WhiteList map[string]uint32 `protobuf:"bytes,1,rep,name=white_list,json=whiteList,proto3" json:"white_list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *ValidatorWhiteList) Reset()         { *m = ValidatorWhiteList{} }
func (m *ValidatorWhiteList) String() string { return proto.CompactTextString(m) }
func (*ValidatorWhiteList) ProtoMessage()    {}
func (*ValidatorWhiteList) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac1b931484f4203, []int{3}
}
func (m *ValidatorWhiteList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorWhiteList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorWhiteList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorWhiteList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorWhiteList.Merge(m, src)
}
func (m *ValidatorWhiteList) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorWhiteList) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorWhiteList.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorWhiteList proto.InternalMessageInfo

func (m *ValidatorWhiteList) GetWhiteList() map[string]uint32 {
	if m != nil {
		return m.WhiteList
	}
	return nil
}

func init() {
	proto.RegisterEnum("sifnode.oracle.v1.StatusText", StatusText_name, StatusText_value)
	proto.RegisterType((*GenesisState)(nil), "sifnode.oracle.v1.GenesisState")
	proto.RegisterMapType((map[uint32]*ValidatorWhiteList)(nil), "sifnode.oracle.v1.GenesisState.AddressWhitelistEntry")
	proto.RegisterType((*Claim)(nil), "sifnode.oracle.v1.Claim")
	proto.RegisterType((*Prophecy)(nil), "sifnode.oracle.v1.Prophecy")
	proto.RegisterType((*ValidatorWhiteList)(nil), "sifnode.oracle.v1.ValidatorWhiteList")
	proto.RegisterMapType((map[string]uint32)(nil), "sifnode.oracle.v1.ValidatorWhiteList.WhiteListEntry")
}

func init() { proto.RegisterFile("sifnode/oracle/v1/types.proto", fileDescriptor_dac1b931484f4203) }

var fileDescriptor_dac1b931484f4203 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xd1, 0x8e, 0xd2, 0x40,
	0x14, 0xa5, 0x25, 0xbb, 0xca, 0x5d, 0xc0, 0x32, 0xae, 0xd9, 0x86, 0xcd, 0x36, 0x04, 0x63, 0x82,
	0x6b, 0xd2, 0x66, 0xd1, 0x4d, 0x8c, 0xeb, 0xcb, 0x0a, 0x75, 0x43, 0x62, 0x08, 0x69, 0x41, 0x8d,
	0x26, 0x36, 0x43, 0x3b, 0x0b, 0xa3, 0xa5, 0x25, 0x9d, 0x81, 0x5d, 0x12, 0x3f, 0xc2, 0xbf, 0xf0,
	0x37, 0x7c, 0xf4, 0x71, 0x1f, 0x7d, 0x34, 0xf0, 0x23, 0x86, 0x29, 0xc5, 0xca, 0xf2, 0xe0, 0xdb,
	0xcc, 0x9c, 0x73, 0xef, 0x3d, 0x67, 0xee, 0xbd, 0x70, 0xc4, 0xe8, 0x65, 0x10, 0x7a, 0xc4, 0x08,
	0x23, 0xec, 0xfa, 0xc4, 0x98, 0x9e, 0x18, 0x7c, 0x36, 0x26, 0x4c, 0x1f, 0x47, 0x21, 0x0f, 0x51,
	0x69, 0x05, 0xeb, 0x31, 0xac, 0x4f, 0x4f, 0xca, 0xfb, 0x83, 0x70, 0x10, 0x0a, 0xd4, 0x58, 0x9e,
	0x62, 0x62, 0xf5, 0x87, 0x0c, 0xf9, 0x0b, 0x12, 0x10, 0x46, 0x99, 0xcd, 0x31, 0x27, 0xa8, 0x0f,
	0x25, 0xec, 0x79, 0x11, 0x61, 0xcc, 0xb9, 0x1a, 0x52, 0x4e, 0x7c, 0xca, 0xb8, 0x2a, 0x55, 0xb2,
	0xb5, 0xbd, 0xfa, 0xa9, 0x7e, 0x2b, 0xab, 0x9e, 0x8e, 0xd5, 0xcf, 0xe3, 0xc0, 0x77, 0x49, 0x9c,
	0x19, 0xf0, 0x68, 0x66, 0x29, 0x78, 0xe3, 0x19, 0x3d, 0x84, 0x02, 0xf6, 0x46, 0x34, 0x70, 0x56,
	0x88, 0x2a, 0x57, 0xa4, 0x5a, 0xce, 0xca, 0x8b, 0xc7, 0x55, 0x12, 0x74, 0x06, 0x30, 0x8e, 0xc2,
	0xf1, 0x90, 0xb8, 0x94, 0x30, 0x35, 0x2b, 0x14, 0x1c, 0x6e, 0x51, 0xd0, 0x89, 0x49, 0x33, 0x2b,
	0x45, 0x2f, 0x7f, 0x86, 0x07, 0x5b, 0xc5, 0x20, 0x05, 0xb2, 0x5f, 0xc8, 0x4c, 0x95, 0x2a, 0x52,
	0xad, 0x60, 0x2d, 0x8f, 0xe8, 0x0c, 0x76, 0xa6, 0xd8, 0x9f, 0x10, 0x21, 0x62, 0xaf, 0xfe, 0x68,
	0x4b, 0x89, 0xb7, 0xd8, 0xa7, 0x1e, 0xe6, 0x61, 0x24, 0x92, 0xbd, 0xa1, 0x8c, 0x5b, 0x71, 0xcc,
	0x0b, 0xf9, 0xb9, 0x54, 0xfd, 0x04, 0x3b, 0x0d, 0x1f, 0xd3, 0x11, 0x2a, 0x82, 0x4c, 0x3d, 0x91,
	0x3a, 0x67, 0xc9, 0xd4, 0x43, 0x4f, 0xa0, 0x34, 0x4d, 0x22, 0x37, 0xac, 0x2a, 0x6b, 0x20, 0xb1,
	0xab, 0xc2, 0x1d, 0x37, 0x0c, 0x38, 0x09, 0xb8, 0x9a, 0x15, 0x94, 0xe4, 0x5a, 0xfd, 0x0a, 0x77,
	0x13, 0x8f, 0xa9, 0x12, 0x79, 0x51, 0xe2, 0x14, 0x76, 0x19, 0xc7, 0x7c, 0x12, 0xe7, 0x2d, 0xd6,
	0x8f, 0xb6, 0xa8, 0xb7, 0x05, 0xa1, 0x4b, 0xae, 0xb9, 0xb5, 0x22, 0xa3, 0xc7, 0xa0, 0xb8, 0x4b,
	0xc9, 0xce, 0x5a, 0x46, 0xfc, 0xc3, 0x39, 0xeb, 0x9e, 0x78, 0x5f, 0x1b, 0x66, 0xd5, 0xef, 0x12,
	0xa0, 0xdb, 0xfe, 0x91, 0x0d, 0x20, 0xc6, 0xc3, 0x49, 0xcd, 0xc7, 0xb3, 0xff, 0xfa, 0x3a, 0x7d,
	0x7d, 0x8a, 0xc7, 0x23, 0x77, 0x95, 0xdc, 0xcb, 0x2f, 0xa1, 0xf8, 0x2f, 0x98, 0x6e, 0x57, 0x2e,
	0x6e, 0xd7, 0x7e, 0xba, 0x5d, 0x85, 0x54, 0x1f, 0x8e, 0x3f, 0x02, 0xfc, 0xb5, 0x8a, 0x0e, 0xe1,
	0xc0, 0xee, 0x9e, 0x77, 0x7b, 0xb6, 0xd3, 0x35, 0xdf, 0x77, 0x9d, 0x5e, 0xdb, 0xee, 0x98, 0x8d,
	0xd6, 0xeb, 0x96, 0xd9, 0x54, 0x32, 0xe8, 0x00, 0xee, 0xa7, 0xc1, 0x8e, 0xd9, 0x6e, 0xb6, 0xda,
	0x17, 0x8a, 0xb4, 0x09, 0xd8, 0xbd, 0x46, 0xc3, 0xb4, 0x6d, 0x45, 0x7e, 0xd5, 0xfc, 0x39, 0xd7,
	0xa4, 0x9b, 0xb9, 0x26, 0xfd, 0x9e, 0x6b, 0xd2, 0xb7, 0x85, 0x96, 0xb9, 0x59, 0x68, 0x99, 0x5f,
	0x0b, 0x2d, 0xf3, 0xe1, 0x78, 0x40, 0xf9, 0x70, 0xd2, 0xd7, 0xdd, 0x70, 0x64, 0xd8, 0xf4, 0xd2,
	0x1d, 0x62, 0x1a, 0x18, 0xc9, 0x76, 0x5e, 0x27, 0xfb, 0x29, 0x96, 0xb3, 0xbf, 0x2b, 0x96, 0xee,
	0xe9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56, 0x65, 0x70, 0xc0, 0xbe, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Prophecies) > 0 {
		for iNdEx := len(m.Prophecies) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Prophecies[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AdminAddress) > 0 {
		i -= len(m.AdminAddress)
		copy(dAtA[i:], m.AdminAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.AdminAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AddressWhitelist) > 0 {
		for k := range m.AddressWhitelist {
			v := m.AddressWhitelist[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintTypes(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintTypes(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintTypes(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Claim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Claim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Claim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Prophecy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Prophecy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Prophecy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClaimValidators) > 0 {
		for iNdEx := len(m.ClaimValidators) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ClaimValidators[iNdEx])
			copy(dAtA[i:], m.ClaimValidators[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.ClaimValidators[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Status != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorWhiteList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorWhiteList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorWhiteList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WhiteList) > 0 {
		for k := range m.WhiteList {
			v := m.WhiteList[k]
			baseI := i
			i = encodeVarintTypes(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintTypes(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintTypes(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AddressWhitelist) > 0 {
		for k, v := range m.AddressWhitelist {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovTypes(uint64(l))
			}
			mapEntrySize := 1 + sovTypes(uint64(k)) + l
			n += mapEntrySize + 1 + sovTypes(uint64(mapEntrySize))
		}
	}
	l = len(m.AdminAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Prophecies) > 0 {
		for _, e := range m.Prophecies {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Claim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Prophecy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovTypes(uint64(m.Status))
	}
	if len(m.ClaimValidators) > 0 {
		for _, s := range m.ClaimValidators {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *ValidatorWhiteList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.WhiteList) > 0 {
		for k, v := range m.WhiteList {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTypes(uint64(len(k))) + 1 + sovTypes(uint64(v))
			n += mapEntrySize + 1 + sovTypes(uint64(mapEntrySize))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressWhitelist", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AddressWhitelist == nil {
				m.AddressWhitelist = make(map[uint32]*ValidatorWhiteList)
			}
			var mapkey uint32
			var mapvalue *ValidatorWhiteList
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthTypes
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthTypes
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &ValidatorWhiteList{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTypes(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTypes
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.AddressWhitelist[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdminAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prophecies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prophecies = append(m.Prophecies, &Prophecy{})
			if err := m.Prophecies[len(m.Prophecies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Claim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Claim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Claim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Prophecy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Prophecy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Prophecy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= StatusText(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimValidators", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimValidators = append(m.ClaimValidators, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ValidatorWhiteList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ValidatorWhiteList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorWhiteList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhiteList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.WhiteList == nil {
				m.WhiteList = make(map[string]uint32)
			}
			var mapkey string
			var mapvalue uint32
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTypes
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTypes
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTypes(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTypes
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.WhiteList[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
