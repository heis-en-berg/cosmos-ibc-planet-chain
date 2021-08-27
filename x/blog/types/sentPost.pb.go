// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blog/sentPost.proto

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

type SentPost struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	PostID  string `protobuf:"bytes,3,opt,name=postID,proto3" json:"postID,omitempty"`
	Title   string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Chain   string `protobuf:"bytes,5,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (m *SentPost) Reset()         { *m = SentPost{} }
func (m *SentPost) String() string { return proto.CompactTextString(m) }
func (*SentPost) ProtoMessage()    {}
func (*SentPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bd27a1b5f48ba3f, []int{0}
}
func (m *SentPost) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SentPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SentPost.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SentPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SentPost.Merge(m, src)
}
func (m *SentPost) XXX_Size() int {
	return m.Size()
}
func (m *SentPost) XXX_DiscardUnknown() {
	xxx_messageInfo_SentPost.DiscardUnknown(m)
}

var xxx_messageInfo_SentPost proto.InternalMessageInfo

func (m *SentPost) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SentPost) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SentPost) GetPostID() string {
	if m != nil {
		return m.PostID
	}
	return ""
}

func (m *SentPost) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SentPost) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func init() {
	proto.RegisterType((*SentPost)(nil), "heisenberg.planet.blog.SentPost")
}

func init() { proto.RegisterFile("blog/sentPost.proto", fileDescriptor_8bd27a1b5f48ba3f) }

var fileDescriptor_8bd27a1b5f48ba3f = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe3, 0xd0, 0x16, 0xf0, 0xc0, 0x60, 0xaa, 0xca, 0x62, 0xb0, 0x2a, 0xa6, 0x4a, 0xa8,
	0xf1, 0xc0, 0x1b, 0x20, 0x18, 0xd8, 0x50, 0xd9, 0xd8, 0x92, 0xf4, 0xe4, 0x58, 0x0a, 0xbe, 0xc8,
	0x3e, 0xa4, 0xf2, 0x16, 0x3c, 0x16, 0x63, 0x47, 0x46, 0x94, 0xbc, 0x08, 0xb2, 0x5d, 0xb6, 0xfb,
	0x7e, 0x7d, 0x37, 0x7c, 0xfc, 0xba, 0xe9, 0xd1, 0xe8, 0x00, 0x8e, 0x5e, 0x30, 0x50, 0x35, 0x78,
	0x24, 0x14, 0xab, 0x0e, 0x6c, 0x00, 0xd7, 0x80, 0x37, 0xd5, 0xd0, 0xd7, 0x0e, 0xa8, 0x8a, 0xda,
	0xcd, 0xd2, 0xa0, 0xc1, 0xa4, 0xe8, 0x78, 0x65, 0xfb, 0xf6, 0xc0, 0x2f, 0x5e, 0x4f, 0xff, 0x42,
	0xf2, 0xf3, 0xd6, 0x43, 0x4d, 0xe8, 0x25, 0x5b, 0xb3, 0xcd, 0xe5, 0xee, 0x1f, 0xc5, 0x15, 0x2f,
	0xed, 0x5e, 0x96, 0x6b, 0xb6, 0x99, 0xed, 0x4a, 0xbb, 0x17, 0x2b, 0xbe, 0x18, 0x30, 0xd0, 0xf3,
	0xa3, 0x3c, 0x4b, 0xe2, 0x89, 0xc4, 0x92, 0xcf, 0xc9, 0x52, 0x0f, 0x72, 0x96, 0xe6, 0x0c, 0x71,
	0x6d, 0xbb, 0xda, 0x3a, 0x39, 0xcf, 0x6b, 0x82, 0x87, 0xa7, 0xef, 0x51, 0xb1, 0xe3, 0xa8, 0xd8,
	0xef, 0xa8, 0xd8, 0xd7, 0xa4, 0x8a, 0xe3, 0xa4, 0x8a, 0x9f, 0x49, 0x15, 0x6f, 0x77, 0xc6, 0x52,
	0xf7, 0xd1, 0x54, 0x2d, 0xbe, 0xeb, 0x18, 0xb3, 0x05, 0xb7, 0x8d, 0x39, 0x3a, 0xe7, 0xe8, 0x83,
	0x4e, 0xdd, 0xf4, 0x39, 0x40, 0x68, 0x16, 0xa9, 0xe3, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x28,
	0x4a, 0xc3, 0x59, 0x0c, 0x01, 0x00, 0x00,
}

func (m *SentPost) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SentPost) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SentPost) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintSentPost(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintSentPost(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PostID) > 0 {
		i -= len(m.PostID)
		copy(dAtA[i:], m.PostID)
		i = encodeVarintSentPost(dAtA, i, uint64(len(m.PostID)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintSentPost(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSentPost(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSentPost(dAtA []byte, offset int, v uint64) int {
	offset -= sovSentPost(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SentPost) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSentPost(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovSentPost(uint64(m.Id))
	}
	l = len(m.PostID)
	if l > 0 {
		n += 1 + l + sovSentPost(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovSentPost(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovSentPost(uint64(l))
	}
	return n
}

func sovSentPost(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSentPost(x uint64) (n int) {
	return sovSentPost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SentPost) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSentPost
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
			return fmt.Errorf("proto: SentPost: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SentPost: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSentPost
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
				return ErrInvalidLengthSentPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSentPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSentPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSentPost
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
				return ErrInvalidLengthSentPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSentPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSentPost
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
				return ErrInvalidLengthSentPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSentPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSentPost
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
				return ErrInvalidLengthSentPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSentPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSentPost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSentPost
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
func skipSentPost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSentPost
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
					return 0, ErrIntOverflowSentPost
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
					return 0, ErrIntOverflowSentPost
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
				return 0, ErrInvalidLengthSentPost
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSentPost
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSentPost
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSentPost        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSentPost          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSentPost = fmt.Errorf("proto: unexpected end of group")
)
