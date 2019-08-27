// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nats_test.proto

/*
	Package test is a generated protocol buffer package.

	It is generated from these files:
		nats_test.proto

	It has these top-level messages:
		TestARequest
		TestAResponse
		FeedDataRequest
		FeedDataResponse
*/
package test

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TestARequest struct {
	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Id    int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *TestARequest) Reset()                    { *m = TestARequest{} }
func (m *TestARequest) String() string            { return proto.CompactTextString(m) }
func (*TestARequest) ProtoMessage()               {}
func (*TestARequest) Descriptor() ([]byte, []int) { return fileDescriptorNatsTest, []int{0} }

func (m *TestARequest) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *TestARequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type TestAResponse struct {
	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	Id     int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *TestAResponse) Reset()                    { *m = TestAResponse{} }
func (m *TestAResponse) String() string            { return proto.CompactTextString(m) }
func (*TestAResponse) ProtoMessage()               {}
func (*TestAResponse) Descriptor() ([]byte, []int) { return fileDescriptorNatsTest, []int{1} }

func (m *TestAResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *TestAResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FeedDataRequest struct {
	Data int64 `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *FeedDataRequest) Reset()                    { *m = FeedDataRequest{} }
func (m *FeedDataRequest) String() string            { return proto.CompactTextString(m) }
func (*FeedDataRequest) ProtoMessage()               {}
func (*FeedDataRequest) Descriptor() ([]byte, []int) { return fileDescriptorNatsTest, []int{2} }

func (m *FeedDataRequest) GetData() int64 {
	if m != nil {
		return m.Data
	}
	return 0
}

type FeedDataResponse struct {
	Sum int64 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (m *FeedDataResponse) Reset()                    { *m = FeedDataResponse{} }
func (m *FeedDataResponse) String() string            { return proto.CompactTextString(m) }
func (*FeedDataResponse) ProtoMessage()               {}
func (*FeedDataResponse) Descriptor() ([]byte, []int) { return fileDescriptorNatsTest, []int{3} }

func (m *FeedDataResponse) GetSum() int64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func init() {
	proto.RegisterType((*TestARequest)(nil), "cdl.protonats.TestARequest")
	proto.RegisterType((*TestAResponse)(nil), "cdl.protonats.TestAResponse")
	proto.RegisterType((*FeedDataRequest)(nil), "cdl.protonats.FeedDataRequest")
	proto.RegisterType((*FeedDataResponse)(nil), "cdl.protonats.FeedDataResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	GetTestA(ctx context.Context, in *TestARequest, opts ...grpc.CallOption) (*TestAResponse, error)
	FeedData(ctx context.Context, opts ...grpc.CallOption) (TestService_FeedDataClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) GetTestA(ctx context.Context, in *TestARequest, opts ...grpc.CallOption) (*TestAResponse, error) {
	out := new(TestAResponse)
	err := grpc.Invoke(ctx, "/cdl.protonats.TestService/GetTestA", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) FeedData(ctx context.Context, opts ...grpc.CallOption) (TestService_FeedDataClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/cdl.protonats.TestService/FeedData", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceFeedDataClient{stream}
	return x, nil
}

type TestService_FeedDataClient interface {
	Send(*FeedDataRequest) error
	CloseAndRecv() (*FeedDataResponse, error)
	grpc.ClientStream
}

type testServiceFeedDataClient struct {
	grpc.ClientStream
}

func (x *testServiceFeedDataClient) Send(m *FeedDataRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceFeedDataClient) CloseAndRecv() (*FeedDataResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FeedDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TestService service

type TestServiceServer interface {
	GetTestA(context.Context, *TestARequest) (*TestAResponse, error)
	FeedData(TestService_FeedDataServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_GetTestA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).GetTestA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cdl.protonats.TestService/GetTestA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).GetTestA(ctx, req.(*TestARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_FeedData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).FeedData(&testServiceFeedDataServer{stream})
}

type TestService_FeedDataServer interface {
	SendAndClose(*FeedDataResponse) error
	Recv() (*FeedDataRequest, error)
	grpc.ServerStream
}

type testServiceFeedDataServer struct {
	grpc.ServerStream
}

func (x *testServiceFeedDataServer) SendAndClose(m *FeedDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceFeedDataServer) Recv() (*FeedDataRequest, error) {
	m := new(FeedDataRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cdl.protonats.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTestA",
			Handler:    _TestService_GetTestA_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FeedData",
			Handler:       _TestService_FeedData_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "nats_test.proto",
}

func (m *TestARequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestARequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Input) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNatsTest(dAtA, i, uint64(len(m.Input)))
		i += copy(dAtA[i:], m.Input)
	}
	if m.Id != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintNatsTest(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *TestAResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestAResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Output) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNatsTest(dAtA, i, uint64(len(m.Output)))
		i += copy(dAtA[i:], m.Output)
	}
	if m.Id != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintNatsTest(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *FeedDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeedDataRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNatsTest(dAtA, i, uint64(m.Data))
	}
	return i, nil
}

func (m *FeedDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeedDataResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Sum != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNatsTest(dAtA, i, uint64(m.Sum))
	}
	return i, nil
}

func encodeFixed64NatsTest(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32NatsTest(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintNatsTest(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TestARequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Input)
	if l > 0 {
		n += 1 + l + sovNatsTest(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovNatsTest(uint64(m.Id))
	}
	return n
}

func (m *TestAResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Output)
	if l > 0 {
		n += 1 + l + sovNatsTest(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovNatsTest(uint64(m.Id))
	}
	return n
}

func (m *FeedDataRequest) Size() (n int) {
	var l int
	_ = l
	if m.Data != 0 {
		n += 1 + sovNatsTest(uint64(m.Data))
	}
	return n
}

func (m *FeedDataResponse) Size() (n int) {
	var l int
	_ = l
	if m.Sum != 0 {
		n += 1 + sovNatsTest(uint64(m.Sum))
	}
	return n
}

func sovNatsTest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNatsTest(x uint64) (n int) {
	return sovNatsTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TestARequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNatsTest
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
			return fmt.Errorf("proto: TestARequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestARequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Input", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNatsTest
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
				return ErrInvalidLengthNatsTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Input = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNatsTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNatsTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNatsTest
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
func (m *TestAResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNatsTest
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
			return fmt.Errorf("proto: TestAResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestAResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Output", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNatsTest
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
				return ErrInvalidLengthNatsTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Output = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNatsTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNatsTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNatsTest
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
func (m *FeedDataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNatsTest
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
			return fmt.Errorf("proto: FeedDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeedDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			m.Data = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNatsTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Data |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNatsTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNatsTest
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
func (m *FeedDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNatsTest
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
			return fmt.Errorf("proto: FeedDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeedDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
			}
			m.Sum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNatsTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sum |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNatsTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNatsTest
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
func skipNatsTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNatsTest
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
					return 0, ErrIntOverflowNatsTest
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
					return 0, ErrIntOverflowNatsTest
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
				return 0, ErrInvalidLengthNatsTest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNatsTest
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
				next, err := skipNatsTest(dAtA[start:])
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
	ErrInvalidLengthNatsTest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNatsTest   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("nats_test.proto", fileDescriptorNatsTest) }

var fileDescriptorNatsTest = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4b, 0x2c, 0x29,
	0x8e, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4d, 0x4e, 0xc9,
	0x81, 0x30, 0x41, 0x32, 0x4a, 0x26, 0x5c, 0x3c, 0x21, 0xa9, 0xc5, 0x25, 0x8e, 0x41, 0xa9, 0x85,
	0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x22, 0x5c, 0xac, 0x99, 0x79, 0x05, 0xa5, 0x25, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0x93, 0x02, 0xa3,
	0x06, 0x73, 0x10, 0x53, 0x66, 0x8a, 0x92, 0x39, 0x17, 0x2f, 0x54, 0x57, 0x71, 0x41, 0x7e, 0x5e,
	0x71, 0xaa, 0x90, 0x18, 0x17, 0x5b, 0x7e, 0x69, 0x09, 0x42, 0x1f, 0x94, 0x87, 0xa1, 0x51, 0x95,
	0x8b, 0xdf, 0x2d, 0x35, 0x35, 0xc5, 0x25, 0xb1, 0x24, 0x11, 0x66, 0xa3, 0x10, 0x17, 0x4b, 0x4a,
	0x62, 0x49, 0x22, 0x58, 0x23, 0x73, 0x10, 0x98, 0xad, 0xa4, 0xc2, 0x25, 0x80, 0x50, 0x06, 0xb5,
	0x42, 0x80, 0x8b, 0xb9, 0xb8, 0x34, 0x17, 0xaa, 0x0c, 0xc4, 0x34, 0x5a, 0xce, 0xc8, 0xc5, 0x0d,
	0x72, 0x46, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x3b, 0x17, 0x87, 0x7b, 0x6a, 0x09,
	0xd8, 0x61, 0x42, 0xd2, 0x7a, 0x28, 0xfe, 0xd4, 0x43, 0xf6, 0xa4, 0x94, 0x0c, 0x76, 0x49, 0x88,
	0x45, 0x4a, 0x0c, 0x42, 0xfe, 0x5c, 0x1c, 0x30, 0xeb, 0x85, 0xe4, 0xd0, 0xd4, 0xa2, 0x39, 0x5f,
	0x4a, 0x1e, 0xa7, 0x3c, 0xcc, 0x38, 0x0d, 0x46, 0x27, 0xb1, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c,
	0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x28, 0x16, 0x50, 0x94, 0x24,
	0xb1, 0x81, 0xf5, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x68, 0x67, 0x1f, 0xa6, 0x01,
	0x00, 0x00,
}
