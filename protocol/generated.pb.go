// Copyright 2017 Sourced Technologies SL
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy
// of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gopkg.in/bblfsh/sdk.v1/protocol/generated.proto

/*
	Package protocol is a generated protocol buffer package.

	It is generated from these files:
		gopkg.in/bblfsh/sdk.v1/protocol/generated.proto

	It has these top-level messages:
		ParseRequest
		ParseResponse
		VersionRequest
		VersionResponse
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import gopkg_in_bblfsh_sdk_v1_uast "gopkg.in/bblfsh/sdk.v1/uast"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Encoding is the encoding used for the content string. Currently only
// UTF-8 or Base64 encodings are supported. You should use UTF-8 if you can
// and Base64 as a fallback.
var Encoding_name = map[int32]string{
	0: "UTF8",
	1: "BASE64",
}
var Encoding_value = map[string]int32{
	"UTF8":   0,
	"BASE64": 1,
}

func (Encoding) EnumDescriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

// Status is the status of a response.
var Status_name = map[int32]string{
	0: "OK",
	1: "ERROR",
	2: "FATAL",
}
var Status_value = map[string]int32{
	"OK":    0,
	"ERROR": 1,
	"FATAL": 2,
}

func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *ParseRequest) Reset()                    { *m = ParseRequest{} }
func (m *ParseRequest) String() string            { return proto.CompactTextString(m) }
func (*ParseRequest) ProtoMessage()               {}
func (*ParseRequest) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *ParseResponse) Reset()                    { *m = ParseResponse{} }
func (m *ParseResponse) String() string            { return proto.CompactTextString(m) }
func (*ParseResponse) ProtoMessage()               {}
func (*ParseResponse) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *VersionRequest) Reset()                    { *m = VersionRequest{} }
func (m *VersionRequest) String() string            { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()               {}
func (*VersionRequest) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *VersionResponse) Reset()                    { *m = VersionResponse{} }
func (m *VersionResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()               {}
func (*VersionResponse) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func init() {
	proto.RegisterType((*ParseRequest)(nil), "gopkg.in.bblfsh.sdk.v1.protocol.ParseRequest")
	proto.RegisterType((*ParseResponse)(nil), "gopkg.in.bblfsh.sdk.v1.protocol.ParseResponse")
	proto.RegisterType((*VersionRequest)(nil), "gopkg.in.bblfsh.sdk.v1.protocol.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "gopkg.in.bblfsh.sdk.v1.protocol.VersionResponse")
	proto.RegisterEnum("gopkg.in.bblfsh.sdk.v1.protocol.Encoding", Encoding_name, Encoding_value)
	proto.RegisterEnum("gopkg.in.bblfsh.sdk.v1.protocol.Status", Status_name, Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ProtocolService service

type ProtocolServiceClient interface {
	// Parse uses DefaultParser to process the given parsing request to get the UAST.
	Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error)
	// Version uses DefaultVersioner to process the given version request to get the version.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type protocolServiceClient struct {
	cc *grpc.ClientConn
}

func NewProtocolServiceClient(cc *grpc.ClientConn) ProtocolServiceClient {
	return &protocolServiceClient{cc}
}

func (c *protocolServiceClient) Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error) {
	out := new(ParseResponse)
	err := grpc.Invoke(ctx, "/gopkg.in.bblfsh.sdk.v1.protocol.ProtocolService/Parse", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc.Invoke(ctx, "/gopkg.in.bblfsh.sdk.v1.protocol.ProtocolService/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProtocolService service

type ProtocolServiceServer interface {
	// Parse uses DefaultParser to process the given parsing request to get the UAST.
	Parse(context.Context, *ParseRequest) (*ParseResponse, error)
	// Version uses DefaultVersioner to process the given version request to get the version.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

func RegisterProtocolServiceServer(s *grpc.Server, srv ProtocolServiceServer) {
	s.RegisterService(&_ProtocolService_serviceDesc, srv)
}

func _ProtocolService_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServiceServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gopkg.in.bblfsh.sdk.v1.protocol.ProtocolService/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServiceServer).Parse(ctx, req.(*ParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtocolService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gopkg.in.bblfsh.sdk.v1.protocol.ProtocolService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProtocolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gopkg.in.bblfsh.sdk.v1.protocol.ProtocolService",
	HandlerType: (*ProtocolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Parse",
			Handler:    _ProtocolService_Parse_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _ProtocolService_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gopkg.in/bblfsh/sdk.v1/protocol/generated.proto",
}

func (m *ParseRequest) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParseRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Filename) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Filename)))
		i += copy(dAtA[i:], m.Filename)
	}
	if len(m.Language) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Language)))
		i += copy(dAtA[i:], m.Language)
	}
	if len(m.Content) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Content)))
		i += copy(dAtA[i:], m.Content)
	}
	if m.Encoding != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Encoding))
	}
	return i, nil
}

func (m *ParseResponse) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParseResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.Status))
	}
	if len(m.Errors) > 0 {
		for _, s := range m.Errors {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.UAST != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.UAST.ProtoSize()))
		n1, err := m.UAST.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *VersionRequest) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *VersionResponse) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(len(m.Version)))
		i += copy(dAtA[i:], m.Version)
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ParseRequest) ProtoSize() (n int) {
	var l int
	_ = l
	l = len(m.Filename)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Language)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Encoding != 0 {
		n += 1 + sovGenerated(uint64(m.Encoding))
	}
	return n
}

func (m *ParseResponse) ProtoSize() (n int) {
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovGenerated(uint64(m.Status))
	}
	if len(m.Errors) > 0 {
		for _, s := range m.Errors {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if m.UAST != nil {
		l = m.UAST.ProtoSize()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *VersionRequest) ProtoSize() (n int) {
	var l int
	_ = l
	return n
}

func (m *VersionResponse) ProtoSize() (n int) {
	var l int
	_ = l
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ParseRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ParseRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParseRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filename", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filename = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Language", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Language = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encoding", wireType)
			}
			m.Encoding = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Encoding |= (Encoding(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *ParseResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ParseResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParseResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (Status(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Errors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Errors = append(m.Errors, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UAST", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UAST == nil {
				m.UAST = &gopkg_in_bblfsh_sdk_v1_uast.Node{}
			}
			if err := m.UAST.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *VersionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VersionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *VersionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VersionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("gopkg.in/bblfsh/sdk.v1/protocol/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xbd, 0x69, 0xea, 0x3a, 0xf3, 0xf5, 0x4b, 0xa3, 0x15, 0x42, 0xd6, 0x1e, 0x1c, 0xd3,
	0x0b, 0xa1, 0xa8, 0x6b, 0x08, 0xa5, 0x42, 0x5c, 0xaa, 0x44, 0x24, 0x17, 0x10, 0xa9, 0x9c, 0x94,
	0x03, 0x37, 0x27, 0xd9, 0xb8, 0x56, 0x5c, 0x6f, 0xf0, 0xda, 0x79, 0x86, 0x2a, 0x4f, 0xc0, 0x25,
	0x52, 0x11, 0x3d, 0x70, 0x45, 0xbc, 0x00, 0x47, 0x8e, 0x3c, 0x01, 0x42, 0xe9, 0x81, 0x2b, 0x8f,
	0x80, 0xbc, 0xb6, 0x43, 0x8a, 0x54, 0xa5, 0xb7, 0xfd, 0xcf, 0xcc, 0x6f, 0x77, 0xe6, 0x3f, 0x0b,
	0x96, 0xcb, 0x27, 0x63, 0x97, 0x7a, 0x81, 0xd5, 0xef, 0xfb, 0x23, 0x71, 0x6a, 0x89, 0xe1, 0x98,
	0x4e, 0x1f, 0x5b, 0x93, 0x90, 0x47, 0x7c, 0xc0, 0x7d, 0xcb, 0x65, 0x01, 0x0b, 0x9d, 0x88, 0x0d,
	0xa9, 0x0c, 0xe1, 0x6a, 0x0e, 0xd0, 0x14, 0xa0, 0x29, 0x40, 0x73, 0x80, 0xec, 0xbb, 0x5e, 0x74,
	0x1a, 0xf7, 0xe9, 0x80, 0x9f, 0x59, 0x2e, 0x77, 0x79, 0x7a, 0x55, 0x3f, 0x1e, 0x49, 0x25, 0x85,
	0x3c, 0xa5, 0x04, 0x79, 0x78, 0x43, 0x03, 0xb1, 0x23, 0xa2, 0x7f, 0x1f, 0xdf, 0xfd, 0x8c, 0x60,
	0xfb, 0xd8, 0x09, 0x05, 0xb3, 0xd9, 0xbb, 0x98, 0x89, 0x08, 0x13, 0xd0, 0x46, 0x9e, 0xcf, 0x02,
	0xe7, 0x8c, 0xe9, 0xc8, 0x44, 0xb5, 0x92, 0xbd, 0xd4, 0x49, 0xce, 0x77, 0x02, 0x37, 0x76, 0x5c,
	0xa6, 0x17, 0xd2, 0x5c, 0xae, 0xb1, 0x0e, 0x5b, 0x03, 0x1e, 0x44, 0x2c, 0x88, 0xf4, 0x0d, 0x99,
	0xca, 0x25, 0x6e, 0x81, 0xc6, 0x82, 0x01, 0x1f, 0x7a, 0x81, 0xab, 0x17, 0x4d, 0x54, 0x2b, 0xd7,
	0x1f, 0xd0, 0x35, 0x23, 0xd3, 0x56, 0x06, 0xd8, 0x4b, 0xf4, 0xb9, 0x76, 0x7e, 0x51, 0x55, 0x7e,
	0x7f, 0xa8, 0x2a, 0xbb, 0x5f, 0x10, 0xfc, 0x9f, 0xf5, 0x2c, 0x26, 0x3c, 0x10, 0x0c, 0x1f, 0x81,
	0x2a, 0x22, 0x27, 0x8a, 0x85, 0x6c, 0xb9, 0x5c, 0xbf, 0xbf, 0xf6, 0x81, 0xae, 0x2c, 0xb7, 0x33,
	0x0c, 0xdf, 0x05, 0x95, 0x85, 0x21, 0x0f, 0x85, 0x5e, 0x30, 0x37, 0x6a, 0x25, 0x3b, 0x53, 0xf8,
	0x08, 0x8a, 0x89, 0x6d, 0x72, 0xa4, 0xff, 0xea, 0xf7, 0x6e, 0xba, 0x36, 0xa9, 0xa1, 0xaf, 0xf9,
	0x90, 0x35, 0xb5, 0xc5, 0x8f, 0x6a, 0xf1, 0xa4, 0xd1, 0xed, 0xd9, 0x12, 0x5c, 0xe9, 0x9a, 0x40,
	0xf9, 0x0d, 0x0b, 0x85, 0xc7, 0x83, 0xcc, 0xea, 0x95, 0xdc, 0x53, 0xd8, 0x59, 0xe6, 0xb2, 0x91,
	0x74, 0xd8, 0x9a, 0xa6, 0xa1, 0x6c, 0x0d, 0xb9, 0xfc, 0x8b, 0xed, 0xbd, 0x00, 0x2d, 0x37, 0x0a,
	0x63, 0x28, 0x9e, 0xf4, 0xda, 0xcf, 0x2a, 0x0a, 0xd1, 0x66, 0x73, 0x53, 0x9e, 0x93, 0xa9, 0x9a,
	0x8d, 0x6e, 0xeb, 0xf0, 0xa0, 0x82, 0x08, 0xcc, 0xe6, 0xa6, 0xda, 0x74, 0x04, 0x3b, 0x3c, 0x20,
	0xdb, 0xe7, 0x1f, 0x0d, 0xe5, 0xd3, 0xa5, 0xa1, 0x7c, 0xbd, 0x34, 0x94, 0x3d, 0x1b, 0xd4, 0xd4,
	0x0d, 0x5c, 0x86, 0x42, 0xe7, 0x65, 0x45, 0x21, 0xea, 0x6c, 0x6e, 0x16, 0x3a, 0x63, 0x7c, 0x07,
	0x36, 0x5b, 0xb6, 0xdd, 0xb1, 0x2b, 0x88, 0x94, 0x66, 0x73, 0x73, 0xb3, 0x95, 0x98, 0x92, 0x44,
	0xdb, 0x8d, 0x5e, 0xe3, 0x55, 0xa5, 0x90, 0x46, 0xdb, 0x4e, 0xe4, 0xf8, 0xd7, 0xef, 0xac, 0xff,
	0x42, 0xb0, 0x73, 0x9c, 0x79, 0xdd, 0x65, 0xe1, 0xd4, 0x1b, 0x30, 0x3c, 0x82, 0x4d, 0xb9, 0x35,
	0xbc, 0xbf, 0x76, 0x3b, 0xab, 0x3f, 0x92, 0xd0, 0xdb, 0x96, 0x67, 0xce, 0xf9, 0xb0, 0x95, 0x99,
	0x89, 0xad, 0xb5, 0xe8, 0xf5, 0x95, 0x90, 0x47, 0xb7, 0x07, 0xd2, 0xd7, 0x9a, 0xc6, 0xb7, 0x85,
	0x81, 0xbe, 0x2f, 0x0c, 0xf4, 0x73, 0x61, 0x28, 0xef, 0xaf, 0x0c, 0xe5, 0xe2, 0xca, 0x40, 0x6f,
	0xb5, 0xbc, 0xbe, 0xaf, 0xca, 0xd3, 0x93, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x75, 0x36,
	0x87, 0x17, 0x04, 0x00, 0x00,
}
