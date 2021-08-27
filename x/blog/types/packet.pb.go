// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blog/packet.proto

package types

import (
	fmt "fmt"
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

type BlogPacketData struct {
	// Types that are valid to be assigned to Packet:
	//	*BlogPacketData_NoData
	//	*BlogPacketData_IbcPostPacket
	Packet isBlogPacketData_Packet `protobuf_oneof:"packet"`
}

func (m *BlogPacketData) Reset()         { *m = BlogPacketData{} }
func (m *BlogPacketData) String() string { return proto.CompactTextString(m) }
func (*BlogPacketData) ProtoMessage()    {}
func (*BlogPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcaa85416e29107a, []int{0}
}
func (m *BlogPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlogPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlogPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlogPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogPacketData.Merge(m, src)
}
func (m *BlogPacketData) XXX_Size() int {
	return m.Size()
}
func (m *BlogPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_BlogPacketData proto.InternalMessageInfo

type isBlogPacketData_Packet interface {
	isBlogPacketData_Packet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type BlogPacketData_NoData struct {
	NoData *NoData `protobuf:"bytes,1,opt,name=noData,proto3,oneof" json:"noData,omitempty"`
}
type BlogPacketData_IbcPostPacket struct {
	IbcPostPacket *IbcPostPacketData `protobuf:"bytes,2,opt,name=ibcPostPacket,proto3,oneof" json:"ibcPostPacket,omitempty"`
}

func (*BlogPacketData_NoData) isBlogPacketData_Packet()        {}
func (*BlogPacketData_IbcPostPacket) isBlogPacketData_Packet() {}

func (m *BlogPacketData) GetPacket() isBlogPacketData_Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *BlogPacketData) GetNoData() *NoData {
	if x, ok := m.GetPacket().(*BlogPacketData_NoData); ok {
		return x.NoData
	}
	return nil
}

func (m *BlogPacketData) GetIbcPostPacket() *IbcPostPacketData {
	if x, ok := m.GetPacket().(*BlogPacketData_IbcPostPacket); ok {
		return x.IbcPostPacket
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BlogPacketData) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BlogPacketData_NoData)(nil),
		(*BlogPacketData_IbcPostPacket)(nil),
	}
}

type NoData struct {
}

func (m *NoData) Reset()         { *m = NoData{} }
func (m *NoData) String() string { return proto.CompactTextString(m) }
func (*NoData) ProtoMessage()    {}
func (*NoData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcaa85416e29107a, []int{1}
}
func (m *NoData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NoData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NoData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NoData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoData.Merge(m, src)
}
func (m *NoData) XXX_Size() int {
	return m.Size()
}
func (m *NoData) XXX_DiscardUnknown() {
	xxx_messageInfo_NoData.DiscardUnknown(m)
}

var xxx_messageInfo_NoData proto.InternalMessageInfo

// this line is used by starport scaffolding # ibc/packet/proto/message
// IbcPostPacketData defines a struct for the packet payload
type IbcPostPacketData struct {
	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *IbcPostPacketData) Reset()         { *m = IbcPostPacketData{} }
func (m *IbcPostPacketData) String() string { return proto.CompactTextString(m) }
func (*IbcPostPacketData) ProtoMessage()    {}
func (*IbcPostPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcaa85416e29107a, []int{2}
}
func (m *IbcPostPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IbcPostPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IbcPostPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IbcPostPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IbcPostPacketData.Merge(m, src)
}
func (m *IbcPostPacketData) XXX_Size() int {
	return m.Size()
}
func (m *IbcPostPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_IbcPostPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_IbcPostPacketData proto.InternalMessageInfo

func (m *IbcPostPacketData) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *IbcPostPacketData) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *IbcPostPacketData) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

// IbcPostPacketAck defines a struct for the packet acknowledgment
type IbcPostPacketAck struct {
	PostID string `protobuf:"bytes,1,opt,name=postID,proto3" json:"postID,omitempty"`
}

func (m *IbcPostPacketAck) Reset()         { *m = IbcPostPacketAck{} }
func (m *IbcPostPacketAck) String() string { return proto.CompactTextString(m) }
func (*IbcPostPacketAck) ProtoMessage()    {}
func (*IbcPostPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcaa85416e29107a, []int{3}
}
func (m *IbcPostPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IbcPostPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IbcPostPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IbcPostPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IbcPostPacketAck.Merge(m, src)
}
func (m *IbcPostPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *IbcPostPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_IbcPostPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_IbcPostPacketAck proto.InternalMessageInfo

func (m *IbcPostPacketAck) GetPostID() string {
	if m != nil {
		return m.PostID
	}
	return ""
}

func init() {
	proto.RegisterType((*BlogPacketData)(nil), "heisenberg.planet.blog.BlogPacketData")
	proto.RegisterType((*NoData)(nil), "heisenberg.planet.blog.NoData")
	proto.RegisterType((*IbcPostPacketData)(nil), "heisenberg.planet.blog.IbcPostPacketData")
	proto.RegisterType((*IbcPostPacketAck)(nil), "heisenberg.planet.blog.IbcPostPacketAck")
}

func init() { proto.RegisterFile("blog/packet.proto", fileDescriptor_bcaa85416e29107a) }

var fileDescriptor_bcaa85416e29107a = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xb3, 0x8a, 0xb1, 0x1d, 0x51, 0xec, 0x22, 0x25, 0xa7, 0x45, 0x72, 0xf2, 0x0f, 0xdd,
	0x80, 0x5e, 0xbc, 0x1a, 0x2a, 0xd8, 0x8b, 0xd4, 0x1c, 0x05, 0x0f, 0x49, 0x58, 0xd2, 0xd0, 0xb8,
	0x13, 0x92, 0x11, 0xf4, 0x2d, 0x7c, 0x0b, 0x5f, 0xc5, 0x63, 0x8f, 0x1e, 0x25, 0x79, 0x11, 0xc9,
	0x6e, 0x04, 0x8b, 0xf6, 0xb6, 0x1f, 0xf3, 0xcd, 0x6f, 0x67, 0xe6, 0x83, 0x51, 0x52, 0x60, 0x16,
	0x94, 0x71, 0xba, 0x54, 0x24, 0xcb, 0x0a, 0x09, 0xf9, 0x78, 0xa1, 0xf2, 0x5a, 0xe9, 0x44, 0x55,
	0x99, 0x2c, 0x8b, 0x58, 0x2b, 0x92, 0x9d, 0xc9, 0x7f, 0x67, 0x70, 0x10, 0x16, 0x98, 0xcd, 0x8d,
	0x79, 0x1a, 0x53, 0xcc, 0xaf, 0xc0, 0xd5, 0xd8, 0xbd, 0x3c, 0x76, 0xcc, 0x4e, 0xf6, 0x2e, 0x84,
	0xfc, 0xbf, 0x57, 0xde, 0x19, 0xd7, 0xad, 0x13, 0xf5, 0x7e, 0x7e, 0x0f, 0xfb, 0x79, 0x92, 0xce,
	0xb1, 0x26, 0x8b, 0xf3, 0xb6, 0x0c, 0xe0, 0x74, 0x13, 0x60, 0xf6, 0xdb, 0xdc, 0xb3, 0xd6, 0x09,
	0xe1, 0x00, 0x5c, 0xbb, 0x87, 0x3f, 0x00, 0xd7, 0x7e, 0xe8, 0x3f, 0xc2, 0xe8, 0x4f, 0x27, 0x3f,
	0x82, 0x1d, 0xca, 0xa9, 0x50, 0x66, 0xe8, 0x61, 0x64, 0x05, 0xf7, 0x60, 0x37, 0x45, 0x4d, 0x4a,
	0xdb, 0x59, 0x86, 0xd1, 0x8f, 0x34, 0x95, 0x4a, 0xc5, 0x84, 0x95, 0xb7, 0xdd, 0x57, 0xac, 0xf4,
	0xcf, 0xe0, 0x70, 0x0d, 0x7f, 0x9d, 0x2e, 0xf9, 0x18, 0xdc, 0x12, 0x6b, 0x9a, 0x4d, 0x7b, 0x7c,
	0xaf, 0xc2, 0x9b, 0x8f, 0x46, 0xb0, 0x55, 0x23, 0xd8, 0x57, 0x23, 0xd8, 0x5b, 0x2b, 0x9c, 0x55,
	0x2b, 0x9c, 0xcf, 0x56, 0x38, 0x0f, 0xe7, 0x59, 0x4e, 0x8b, 0xe7, 0x44, 0xa6, 0xf8, 0x14, 0x74,
	0xeb, 0x4f, 0x94, 0x9e, 0x74, 0x07, 0x08, 0xec, 0x01, 0x82, 0x97, 0xc0, 0x84, 0x44, 0xaf, 0xa5,
	0xaa, 0x13, 0xd7, 0x84, 0x74, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xa0, 0x17, 0x69, 0xb9,
	0x01, 0x00, 0x00,
}

func (m *BlogPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlogPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlogPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Packet != nil {
		{
			size := m.Packet.Size()
			i -= size
			if _, err := m.Packet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *BlogPacketData_NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlogPacketData_NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.NoData != nil {
		{
			size, err := m.NoData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *BlogPacketData_IbcPostPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlogPacketData_IbcPostPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.IbcPostPacket != nil {
		{
			size, err := m.IbcPostPacket.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPacket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *NoData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NoData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NoData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *IbcPostPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IbcPostPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IbcPostPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IbcPostPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IbcPostPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IbcPostPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PostID) > 0 {
		i -= len(m.PostID)
		copy(dAtA[i:], m.PostID)
		i = encodeVarintPacket(dAtA, i, uint64(len(m.PostID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlogPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Packet != nil {
		n += m.Packet.Size()
	}
	return n
}

func (m *BlogPacketData_NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NoData != nil {
		l = m.NoData.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *BlogPacketData_IbcPostPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IbcPostPacket != nil {
		l = m.IbcPostPacket.Size()
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}
func (m *NoData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *IbcPostPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func (m *IbcPostPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PostID)
	if l > 0 {
		n += 1 + l + sovPacket(uint64(l))
	}
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlogPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: BlogPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlogPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &NoData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &BlogPacketData_NoData{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcPostPacket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &IbcPostPacketData{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Packet = &BlogPacketData_IbcPostPacket{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *NoData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: NoData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NoData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *IbcPostPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: IbcPostPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IbcPostPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *IbcPostPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: IbcPostPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IbcPostPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
