// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/finality/v1/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the module.
type Params struct {
	// max_active_finality_providers is the maximum number of active finality providers in the BTC staking protocol
	MaxActiveFinalityProviders uint32 `protobuf:"varint,1,opt,name=max_active_finality_providers,json=maxActiveFinalityProviders,proto3" json:"max_active_finality_providers,omitempty"`
	// signed_blocks_window defines the size of the sliding window for tracking finality provider liveness
	SignedBlocksWindow int64 `protobuf:"varint,2,opt,name=signed_blocks_window,json=signedBlocksWindow,proto3" json:"signed_blocks_window,omitempty"`
	// finality_sig_timeout defines how much time (in terms of blocks) finality providers have to cast a finality
	// vote before being judged as missing their voting turn on the given block
	FinalitySigTimeout int64 `protobuf:"varint,3,opt,name=finality_sig_timeout,json=finalitySigTimeout,proto3" json:"finality_sig_timeout,omitempty"`
	// min_signed_per_window defines the minimum number of blocks that a finality provider is required to sign
	// within the sliding window to avoid being jailed
	MinSignedPerWindow cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=min_signed_per_window,json=minSignedPerWindow,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"min_signed_per_window"`
	// min_pub_rand is the minimum number of public randomness each
	// message should commit
	MinPubRand uint64 `protobuf:"varint,5,opt,name=min_pub_rand,json=minPubRand,proto3" json:"min_pub_rand,omitempty"`
	// jail_duration is the minimum period of time that a finality provider remains jailed
	JailDuration time.Duration `protobuf:"bytes,6,opt,name=jail_duration,json=jailDuration,proto3,stdduration" json:"jail_duration"`
	// finality_activation_height is the babylon block height which the finality module will
	// start to accept finality voting and the minimum allowed value for the public randomness
	// commit start height.
	FinalityActivationHeight uint64 `protobuf:"varint,7,opt,name=finality_activation_height,json=finalityActivationHeight,proto3" json:"finality_activation_height,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_25539c9a61c72ee9, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxActiveFinalityProviders() uint32 {
	if m != nil {
		return m.MaxActiveFinalityProviders
	}
	return 0
}

func (m *Params) GetSignedBlocksWindow() int64 {
	if m != nil {
		return m.SignedBlocksWindow
	}
	return 0
}

func (m *Params) GetFinalitySigTimeout() int64 {
	if m != nil {
		return m.FinalitySigTimeout
	}
	return 0
}

func (m *Params) GetMinPubRand() uint64 {
	if m != nil {
		return m.MinPubRand
	}
	return 0
}

func (m *Params) GetJailDuration() time.Duration {
	if m != nil {
		return m.JailDuration
	}
	return 0
}

func (m *Params) GetFinalityActivationHeight() uint64 {
	if m != nil {
		return m.FinalityActivationHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "babylon.finality.v1.Params")
}

func init() { proto.RegisterFile("babylon/finality/v1/params.proto", fileDescriptor_25539c9a61c72ee9) }

var fileDescriptor_25539c9a61c72ee9 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x92, 0x31, 0x8f, 0xd3, 0x3e,
	0x18, 0xc6, 0xeb, 0x7f, 0xfb, 0x2f, 0x52, 0x68, 0x07, 0xc2, 0x21, 0xe5, 0x8a, 0x48, 0x23, 0xa6,
	0x0a, 0xe9, 0x12, 0xee, 0x90, 0x18, 0x10, 0x4b, 0xab, 0x0a, 0x31, 0x1c, 0x52, 0xd5, 0x43, 0x42,
	0x62, 0xb1, 0x9c, 0xc4, 0xe7, 0xbe, 0x5c, 0x6c, 0x57, 0xb1, 0xd3, 0x6b, 0xbf, 0x05, 0xe3, 0x8d,
	0x8c, 0x8c, 0x0c, 0x7c, 0x88, 0xdb, 0x38, 0x31, 0x21, 0x86, 0x03, 0xb5, 0x03, 0x5f, 0x03, 0xc5,
	0x8e, 0xcb, 0x12, 0xe5, 0xf5, 0xf3, 0x7b, 0xfd, 0xbc, 0x79, 0xf2, 0x7a, 0x51, 0x4a, 0xd2, 0x4d,
	0x21, 0x45, 0x72, 0x0e, 0x82, 0x14, 0xa0, 0x37, 0xc9, 0xea, 0x38, 0x59, 0x92, 0x92, 0x70, 0x15,
	0x2f, 0x4b, 0xa9, 0xa5, 0x7f, 0xbf, 0x21, 0x62, 0x47, 0xc4, 0xab, 0xe3, 0xc1, 0x01, 0x93, 0x4c,
	0x1a, 0x3d, 0xa9, 0xdf, 0x2c, 0x3a, 0xb8, 0x47, 0x38, 0x08, 0x99, 0x98, 0x67, 0x73, 0x74, 0x98,
	0x49, 0xc5, 0xa5, 0xc2, 0x96, 0xb5, 0x45, 0x23, 0x85, 0x4c, 0x4a, 0x56, 0xd0, 0xc4, 0x54, 0x69,
	0x75, 0x9e, 0xe4, 0x55, 0x49, 0x34, 0x48, 0x61, 0xf5, 0xc7, 0xdf, 0xda, 0x5e, 0x77, 0x66, 0x26,
	0xf1, 0xc7, 0xde, 0x23, 0x4e, 0xd6, 0x98, 0x64, 0x1a, 0x56, 0x14, 0xbb, 0x41, 0xea, 0x4b, 0x57,
	0x90, 0xd3, 0x52, 0x05, 0x28, 0x42, 0xa3, 0xfe, 0x7c, 0xc0, 0xc9, 0x7a, 0x6c, 0x98, 0x57, 0x0d,
	0x32, 0x73, 0x84, 0xff, 0xd4, 0x3b, 0x50, 0xc0, 0x04, 0xcd, 0x71, 0x5a, 0xc8, 0xec, 0x42, 0xe1,
	0x4b, 0x10, 0xb9, 0xbc, 0x0c, 0xfe, 0x8b, 0xd0, 0xa8, 0x3d, 0xf7, 0xad, 0x36, 0x31, 0xd2, 0x3b,
	0xa3, 0xd4, 0x1d, 0x7b, 0x27, 0x05, 0x0c, 0x6b, 0xe0, 0x54, 0x56, 0x3a, 0x68, 0xdb, 0x0e, 0xa7,
	0x9d, 0x01, 0x7b, 0x6b, 0x15, 0x1f, 0xbc, 0x07, 0x1c, 0x04, 0x6e, 0x7c, 0x96, 0xb4, 0x74, 0x26,
	0x9d, 0x08, 0x8d, 0x7a, 0x93, 0xe7, 0xd7, 0xb7, 0xc3, 0xd6, 0xcf, 0xdb, 0xe1, 0x43, 0x1b, 0x83,
	0xca, 0x2f, 0x62, 0x90, 0x09, 0x27, 0x7a, 0x11, 0x9f, 0x52, 0x46, 0xb2, 0xcd, 0x94, 0x66, 0xdf,
	0xbf, 0x1e, 0x79, 0x4d, 0x4a, 0x53, 0x9a, 0x7d, 0xfe, 0xf3, 0xe5, 0x09, 0x9a, 0xfb, 0x1c, 0xc4,
	0x99, 0xb9, 0x73, 0x46, 0xcb, 0x66, 0xb8, 0xc8, 0xeb, 0xd5, 0x56, 0xcb, 0x2a, 0xc5, 0x25, 0x11,
	0x79, 0xf0, 0x7f, 0x84, 0x46, 0x9d, 0xb9, 0xc7, 0x41, 0xcc, 0xaa, 0x74, 0x4e, 0x44, 0xee, 0xbf,
	0xf1, 0xfa, 0x1f, 0x08, 0x14, 0xd8, 0xa5, 0x1a, 0x74, 0x23, 0x34, 0xba, 0x7b, 0x72, 0x18, 0xdb,
	0xd8, 0x63, 0x17, 0x7b, 0x3c, 0x6d, 0x80, 0x49, 0xbf, 0x9e, 0xef, 0xea, 0xd7, 0x10, 0x59, 0xdb,
	0x5e, 0xdd, 0xee, 0x44, 0xff, 0xa5, 0x37, 0xd8, 0xa7, 0x61, 0xfe, 0x83, 0x39, 0xc6, 0x0b, 0x0a,
	0x6c, 0xa1, 0x83, 0x3b, 0xc6, 0x3e, 0x70, 0xc4, 0x78, 0x0f, 0xbc, 0x36, 0xfa, 0x8b, 0xce, 0xd5,
	0xa7, 0x61, 0x6b, 0x72, 0x7a, 0xbd, 0x0d, 0xd1, 0xcd, 0x36, 0x44, 0xbf, 0xb7, 0x21, 0xfa, 0xb8,
	0x0b, 0x5b, 0x37, 0xbb, 0xb0, 0xf5, 0x63, 0x17, 0xb6, 0xde, 0x9f, 0x30, 0xd0, 0x8b, 0x2a, 0x8d,
	0x33, 0xc9, 0x93, 0x66, 0xdf, 0x0a, 0x92, 0xaa, 0x23, 0x90, 0xae, 0x4c, 0xd6, 0xff, 0x56, 0x54,
	0x6f, 0x96, 0x54, 0xa5, 0x5d, 0xf3, 0x05, 0xcf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x27,
	0x77, 0xaf, 0xc3, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FinalityActivationHeight != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.FinalityActivationHeight))
		i--
		dAtA[i] = 0x38
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.JailDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.JailDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	if m.MinPubRand != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinPubRand))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.MinSignedPerWindow.Size()
		i -= size
		if _, err := m.MinSignedPerWindow.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.FinalitySigTimeout != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.FinalitySigTimeout))
		i--
		dAtA[i] = 0x18
	}
	if m.SignedBlocksWindow != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SignedBlocksWindow))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxActiveFinalityProviders != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxActiveFinalityProviders))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxActiveFinalityProviders != 0 {
		n += 1 + sovParams(uint64(m.MaxActiveFinalityProviders))
	}
	if m.SignedBlocksWindow != 0 {
		n += 1 + sovParams(uint64(m.SignedBlocksWindow))
	}
	if m.FinalitySigTimeout != 0 {
		n += 1 + sovParams(uint64(m.FinalitySigTimeout))
	}
	l = m.MinSignedPerWindow.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MinPubRand != 0 {
		n += 1 + sovParams(uint64(m.MinPubRand))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.JailDuration)
	n += 1 + l + sovParams(uint64(l))
	if m.FinalityActivationHeight != 0 {
		n += 1 + sovParams(uint64(m.FinalityActivationHeight))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxActiveFinalityProviders", wireType)
			}
			m.MaxActiveFinalityProviders = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxActiveFinalityProviders |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedBlocksWindow", wireType)
			}
			m.SignedBlocksWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignedBlocksWindow |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalitySigTimeout", wireType)
			}
			m.FinalitySigTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinalitySigTimeout |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSignedPerWindow", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinSignedPerWindow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinPubRand", wireType)
			}
			m.MinPubRand = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinPubRand |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.JailDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalityActivationHeight", wireType)
			}
			m.FinalityActivationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinalityActivationHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
