// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpcproto/packet.proto

// Copyright (c) 2018 Donovan Solms / MiningHQ. Licensed under the MIT license.
// See LICENSE in the root of this repository for details.

package rpcproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Method is the supported methods for RPC calls
type Method int32

const (
	Method_UnknownMethod Method = 0
	// RigWarning information
	Method_RigWarning Method = 1
	// RigError information
	Method_RigError Method = 2
	// RigAssignment sets a miner's mining assignment
	Method_RigAssignment Method = 3
	// State updates a miner's state
	Method_State Method = 4
	// Logs requests logs from a miner
	Method_Logs Method = 5
	// Stats requests and send miner stats
	Method_Stats Method = 6
	// RigInfo requests rig information from MiningHQ
	Method_RigInfo Method = 7
)

var Method_name = map[int32]string{
	0: "UnknownMethod",
	1: "RigWarning",
	2: "RigError",
	3: "RigAssignment",
	4: "State",
	5: "Logs",
	6: "Stats",
	7: "RigInfo",
}

var Method_value = map[string]int32{
	"UnknownMethod": 0,
	"RigWarning":    1,
	"RigError":      2,
	"RigAssignment": 3,
	"State":         4,
	"Logs":          5,
	"Stats":         6,
	"RigInfo":       7,
}

func (x Method) String() string {
	return proto.EnumName(Method_name, int32(x))
}

func (Method) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e775f4d4d61f99b, []int{0}
}

// Packet contains the RPC call information sent over websockets
type Packet struct {
	// Method is the method to execute
	Method Method `protobuf:"varint,1,opt,name=Method,proto3,enum=rpcproto.Method" json:"Method,omitempty"`
	// Params contain the information needed for Method
	//
	// Types that are valid to be assigned to Params:
	//	*Packet_RigWarning
	//	*Packet_RigError
	//	*Packet_LogsRequest
	//	*Packet_LogsResponse
	//	*Packet_StatsRequest
	//	*Packet_StatsResponse
	//	*Packet_StateRequest
	//	*Packet_StateResponse
	//	*Packet_RigAssignmentRequest
	//	*Packet_RigAssignmentResponse
	//	*Packet_RigInfoRequest
	//	*Packet_RigInfoResponse
	Params               isPacket_Params `protobuf_oneof:"Params"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e775f4d4d61f99b, []int{0}
}

func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetMethod() Method {
	if m != nil {
		return m.Method
	}
	return Method_UnknownMethod
}

type isPacket_Params interface {
	isPacket_Params()
}

type Packet_RigWarning struct {
	RigWarning *RigWarningDetail `protobuf:"bytes,2,opt,name=RigWarning,proto3,oneof"`
}

type Packet_RigError struct {
	RigError *RigErrorDetail `protobuf:"bytes,3,opt,name=RigError,proto3,oneof"`
}

type Packet_LogsRequest struct {
	LogsRequest *LogsRequest `protobuf:"bytes,4,opt,name=LogsRequest,proto3,oneof"`
}

type Packet_LogsResponse struct {
	LogsResponse *LogsResponse `protobuf:"bytes,5,opt,name=LogsResponse,proto3,oneof"`
}

type Packet_StatsRequest struct {
	StatsRequest *StatsRequest `protobuf:"bytes,6,opt,name=StatsRequest,proto3,oneof"`
}

type Packet_StatsResponse struct {
	StatsResponse *StatsResponse `protobuf:"bytes,7,opt,name=StatsResponse,proto3,oneof"`
}

type Packet_StateRequest struct {
	StateRequest *StateRequest `protobuf:"bytes,8,opt,name=StateRequest,proto3,oneof"`
}

type Packet_StateResponse struct {
	StateResponse *StateResponse `protobuf:"bytes,9,opt,name=StateResponse,proto3,oneof"`
}

type Packet_RigAssignmentRequest struct {
	RigAssignmentRequest *RigAssignmentRequest `protobuf:"bytes,10,opt,name=RigAssignmentRequest,proto3,oneof"`
}

type Packet_RigAssignmentResponse struct {
	RigAssignmentResponse *RigAssignmentResponse `protobuf:"bytes,11,opt,name=RigAssignmentResponse,proto3,oneof"`
}

type Packet_RigInfoRequest struct {
	RigInfoRequest *RigInfoRequest `protobuf:"bytes,12,opt,name=RigInfoRequest,proto3,oneof"`
}

type Packet_RigInfoResponse struct {
	RigInfoResponse *RigInfoResponse `protobuf:"bytes,13,opt,name=RigInfoResponse,proto3,oneof"`
}

func (*Packet_RigWarning) isPacket_Params() {}

func (*Packet_RigError) isPacket_Params() {}

func (*Packet_LogsRequest) isPacket_Params() {}

func (*Packet_LogsResponse) isPacket_Params() {}

func (*Packet_StatsRequest) isPacket_Params() {}

func (*Packet_StatsResponse) isPacket_Params() {}

func (*Packet_StateRequest) isPacket_Params() {}

func (*Packet_StateResponse) isPacket_Params() {}

func (*Packet_RigAssignmentRequest) isPacket_Params() {}

func (*Packet_RigAssignmentResponse) isPacket_Params() {}

func (*Packet_RigInfoRequest) isPacket_Params() {}

func (*Packet_RigInfoResponse) isPacket_Params() {}

func (m *Packet) GetParams() isPacket_Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Packet) GetRigWarning() *RigWarningDetail {
	if x, ok := m.GetParams().(*Packet_RigWarning); ok {
		return x.RigWarning
	}
	return nil
}

func (m *Packet) GetRigError() *RigErrorDetail {
	if x, ok := m.GetParams().(*Packet_RigError); ok {
		return x.RigError
	}
	return nil
}

func (m *Packet) GetLogsRequest() *LogsRequest {
	if x, ok := m.GetParams().(*Packet_LogsRequest); ok {
		return x.LogsRequest
	}
	return nil
}

func (m *Packet) GetLogsResponse() *LogsResponse {
	if x, ok := m.GetParams().(*Packet_LogsResponse); ok {
		return x.LogsResponse
	}
	return nil
}

func (m *Packet) GetStatsRequest() *StatsRequest {
	if x, ok := m.GetParams().(*Packet_StatsRequest); ok {
		return x.StatsRequest
	}
	return nil
}

func (m *Packet) GetStatsResponse() *StatsResponse {
	if x, ok := m.GetParams().(*Packet_StatsResponse); ok {
		return x.StatsResponse
	}
	return nil
}

func (m *Packet) GetStateRequest() *StateRequest {
	if x, ok := m.GetParams().(*Packet_StateRequest); ok {
		return x.StateRequest
	}
	return nil
}

func (m *Packet) GetStateResponse() *StateResponse {
	if x, ok := m.GetParams().(*Packet_StateResponse); ok {
		return x.StateResponse
	}
	return nil
}

func (m *Packet) GetRigAssignmentRequest() *RigAssignmentRequest {
	if x, ok := m.GetParams().(*Packet_RigAssignmentRequest); ok {
		return x.RigAssignmentRequest
	}
	return nil
}

func (m *Packet) GetRigAssignmentResponse() *RigAssignmentResponse {
	if x, ok := m.GetParams().(*Packet_RigAssignmentResponse); ok {
		return x.RigAssignmentResponse
	}
	return nil
}

func (m *Packet) GetRigInfoRequest() *RigInfoRequest {
	if x, ok := m.GetParams().(*Packet_RigInfoRequest); ok {
		return x.RigInfoRequest
	}
	return nil
}

func (m *Packet) GetRigInfoResponse() *RigInfoResponse {
	if x, ok := m.GetParams().(*Packet_RigInfoResponse); ok {
		return x.RigInfoResponse
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Packet) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Packet_OneofMarshaler, _Packet_OneofUnmarshaler, _Packet_OneofSizer, []interface{}{
		(*Packet_RigWarning)(nil),
		(*Packet_RigError)(nil),
		(*Packet_LogsRequest)(nil),
		(*Packet_LogsResponse)(nil),
		(*Packet_StatsRequest)(nil),
		(*Packet_StatsResponse)(nil),
		(*Packet_StateRequest)(nil),
		(*Packet_StateResponse)(nil),
		(*Packet_RigAssignmentRequest)(nil),
		(*Packet_RigAssignmentResponse)(nil),
		(*Packet_RigInfoRequest)(nil),
		(*Packet_RigInfoResponse)(nil),
	}
}

func _Packet_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Packet)
	// Params
	switch x := m.Params.(type) {
	case *Packet_RigWarning:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RigWarning); err != nil {
			return err
		}
	case *Packet_RigError:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RigError); err != nil {
			return err
		}
	case *Packet_LogsRequest:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LogsRequest); err != nil {
			return err
		}
	case *Packet_LogsResponse:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LogsResponse); err != nil {
			return err
		}
	case *Packet_StatsRequest:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StatsRequest); err != nil {
			return err
		}
	case *Packet_StatsResponse:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StatsResponse); err != nil {
			return err
		}
	case *Packet_StateRequest:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StateRequest); err != nil {
			return err
		}
	case *Packet_StateResponse:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StateResponse); err != nil {
			return err
		}
	case *Packet_RigAssignmentRequest:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RigAssignmentRequest); err != nil {
			return err
		}
	case *Packet_RigAssignmentResponse:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RigAssignmentResponse); err != nil {
			return err
		}
	case *Packet_RigInfoRequest:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RigInfoRequest); err != nil {
			return err
		}
	case *Packet_RigInfoResponse:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RigInfoResponse); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Packet.Params has unexpected type %T", x)
	}
	return nil
}

func _Packet_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Packet)
	switch tag {
	case 2: // Params.RigWarning
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RigWarningDetail)
		err := b.DecodeMessage(msg)
		m.Params = &Packet_RigWarning{msg}
		return true, err
	case 3: // Params.RigError
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RigErrorDetail)
		err := b.DecodeMessage(msg)
		m.Params = &Packet_RigError{msg}
		return true, err
	case 4: // Params.LogsRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogsRequest)
		err := b.DecodeMessage(msg)
		m.Params = &Packet_LogsRequest{msg}
		return true, err
	case 5: // Params.LogsResponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogsResponse)
		err := b.DecodeMessage(msg)
		m.Params = &Packet_LogsResponse{msg}
		return true, err
	case 6: // Params.StatsRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StatsRequest)
		err := b.DecodeMessage(msg)
		m.Params = &Packet_StatsRequest{msg}
		return true, err
	case 7: // Params.StatsResponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StatsResponse)
		err := b.DecodeMessage(msg)
		m.Params = &Packet_StatsResponse{msg}
		return true, err
	case 8: // Params.StateRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StateRequest)
		err := b.DecodeMessage(msg)
		m.Params = &Packet_StateRequest{msg}
		return true, err
	case 9: // Params.StateResponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StateResponse)
		err := b.DecodeMessage(msg)
		m.Params = &Packet_StateResponse{msg}
		return true, err
	case 10: // Params.RigAssignmentRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RigAssignmentRequest)
		err := b.DecodeMessage(msg)
		m.Params = &Packet_RigAssignmentRequest{msg}
		return true, err
	case 11: // Params.RigAssignmentResponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RigAssignmentResponse)
		err := b.DecodeMessage(msg)
		m.Params = &Packet_RigAssignmentResponse{msg}
		return true, err
	case 12: // Params.RigInfoRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RigInfoRequest)
		err := b.DecodeMessage(msg)
		m.Params = &Packet_RigInfoRequest{msg}
		return true, err
	case 13: // Params.RigInfoResponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RigInfoResponse)
		err := b.DecodeMessage(msg)
		m.Params = &Packet_RigInfoResponse{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Packet_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Packet)
	// Params
	switch x := m.Params.(type) {
	case *Packet_RigWarning:
		s := proto.Size(x.RigWarning)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Packet_RigError:
		s := proto.Size(x.RigError)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Packet_LogsRequest:
		s := proto.Size(x.LogsRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Packet_LogsResponse:
		s := proto.Size(x.LogsResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Packet_StatsRequest:
		s := proto.Size(x.StatsRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Packet_StatsResponse:
		s := proto.Size(x.StatsResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Packet_StateRequest:
		s := proto.Size(x.StateRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Packet_StateResponse:
		s := proto.Size(x.StateResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Packet_RigAssignmentRequest:
		s := proto.Size(x.RigAssignmentRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Packet_RigAssignmentResponse:
		s := proto.Size(x.RigAssignmentResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Packet_RigInfoRequest:
		s := proto.Size(x.RigInfoRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Packet_RigInfoResponse:
		s := proto.Size(x.RigInfoResponse)
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
	proto.RegisterEnum("rpcproto.Method", Method_name, Method_value)
	proto.RegisterType((*Packet)(nil), "rpcproto.Packet")
}

func init() { proto.RegisterFile("rpcproto/packet.proto", fileDescriptor_6e775f4d4d61f99b) }

var fileDescriptor_6e775f4d4d61f99b = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x5f, 0x6b, 0x14, 0x31,
	0x14, 0xc5, 0x67, 0xdb, 0xdd, 0xd9, 0xe9, 0xdd, 0x3f, 0xc6, 0x6b, 0xd7, 0x8e, 0x0b, 0x6a, 0xf1,
	0x69, 0xf1, 0x61, 0x85, 0x0a, 0x82, 0x50, 0x10, 0x8b, 0x85, 0x11, 0x14, 0x24, 0x2a, 0x7d, 0x94,
	0xb1, 0xa6, 0x31, 0xb4, 0x4d, 0xd6, 0x64, 0x44, 0xfc, 0xbc, 0x7e, 0x11, 0x99, 0x4c, 0x92, 0x9d,
	0xc9, 0x4e, 0x1f, 0x73, 0xce, 0xfd, 0xdd, 0x73, 0x76, 0x27, 0x81, 0x85, 0xde, 0x5c, 0x6e, 0xb4,
	0xaa, 0xd4, 0x8b, 0x4d, 0x79, 0x79, 0xcd, 0xaa, 0xb5, 0x3d, 0x60, 0xe6, 0xe5, 0xe5, 0x83, 0x30,
	0x70, 0xa3, 0xb8, 0x69, 0xec, 0xe5, 0x61, 0x10, 0x99, 0xd6, 0x4a, 0xef, 0xa8, 0xa6, 0x2a, 0x2b,
	0xd6, 0xab, 0xfa, 0x0d, 0x47, 0x41, 0xd5, 0x82, 0x7f, 0x13, 0xf2, 0x4a, 0x39, 0xe3, 0x71, 0xc7,
	0x28, 0x8d, 0x11, 0x5c, 0xde, 0x32, 0xe9, 0x8a, 0x3d, 0xfb, 0x97, 0x42, 0xfa, 0xc9, 0x36, 0xc5,
	0x15, 0xa4, 0x1f, 0x59, 0xf5, 0x53, 0xfd, 0xc8, 0x07, 0xc7, 0x83, 0xd5, 0xfc, 0x84, 0xac, 0x3d,
	0xba, 0x6e, 0x74, 0xea, 0x7c, 0x3c, 0x05, 0xa0, 0x82, 0x5f, 0x94, 0x5a, 0x0a, 0xc9, 0xf3, 0xbd,
	0xe3, 0xc1, 0x6a, 0x72, 0xb2, 0xdc, 0x4e, 0x6f, 0xbd, 0x77, 0xac, 0x2a, 0xc5, 0x4d, 0x91, 0xd0,
	0xd6, 0x3c, 0xbe, 0x82, 0x8c, 0x0a, 0x7e, 0x5e, 0xff, 0xd0, 0x7c, 0xdf, 0xb2, 0x79, 0x87, 0xb5,
	0x4e, 0x20, 0xc3, 0x2c, 0xbe, 0x86, 0xc9, 0x07, 0xc5, 0x0d, 0x65, 0xbf, 0x7e, 0x33, 0x53, 0xe5,
	0x43, 0x8b, 0x2e, 0xb6, 0x68, 0xcb, 0x2c, 0x12, 0xda, 0x9e, 0xc5, 0x53, 0x98, 0x36, 0x47, 0xb3,
	0x51, 0xd2, 0xb0, 0x7c, 0x64, 0xd9, 0x87, 0x31, 0xdb, 0xb8, 0x45, 0x42, 0x3b, 0xd3, 0x35, 0xfd,
	0xb9, 0xfe, 0xab, 0x7d, 0x72, 0x1a, 0xd3, 0x6d, 0xb7, 0xa6, 0xdb, 0x67, 0x7c, 0x03, 0x33, 0x77,
	0x76, 0xe1, 0x63, 0x8b, 0x1f, 0xed, 0xe0, 0x21, 0xbd, 0x3b, 0xef, 0xe3, 0x99, 0x8f, 0xcf, 0xfa,
	0xe2, 0x59, 0x14, 0xcf, 0xa2, 0x78, 0x16, 0xe2, 0x0f, 0xfa, 0xe2, 0x59, 0x1c, 0x1f, 0x04, 0xfc,
	0x02, 0x87, 0x54, 0xf0, 0xb7, 0xe1, 0xe2, 0xf8, 0x1a, 0x60, 0xf7, 0x3c, 0xe9, 0x7c, 0xba, 0x9d,
	0xa9, 0x22, 0xa1, 0xbd, 0x34, 0x5e, 0xc0, 0x22, 0xd2, 0x5d, 0xbd, 0x89, 0x5d, 0xfb, 0xf4, 0xce,
	0xb5, 0xa1, 0x66, 0x3f, 0x8f, 0x67, 0x30, 0xa7, 0x82, 0xbf, 0x97, 0x57, 0xca, 0x17, 0x9d, 0xf6,
	0xdc, 0xb1, 0x96, 0x5f, 0x24, 0x34, 0x22, 0xf0, 0x1c, 0xee, 0x05, 0xc5, 0xd5, 0x9a, 0xd9, 0x25,
	0x8f, 0x7a, 0x96, 0x84, 0x42, 0x31, 0x73, 0x96, 0xd5, 0x4f, 0x4b, 0x97, 0xb7, 0xe6, 0xf9, 0x5f,
	0xff, 0xb4, 0xf0, 0x3e, 0xcc, 0xbe, 0xca, 0x6b, 0xa9, 0xfe, 0xc8, 0x46, 0x20, 0x09, 0xce, 0xdb,
	0xaf, 0x89, 0x0c, 0x70, 0xba, 0x7d, 0x1f, 0x64, 0xaf, 0x06, 0x3a, 0x3f, 0x94, 0xec, 0xe3, 0x01,
	0x8c, 0xec, 0x27, 0x22, 0x43, 0xcc, 0x60, 0x58, 0x5f, 0x55, 0x32, 0xf2, 0xa2, 0x21, 0x29, 0x4e,
	0x60, 0xec, 0xaa, 0x90, 0xf1, 0xf7, 0xd4, 0xd6, 0x7d, 0xf9, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xd8,
	0xe7, 0x7c, 0x03, 0x99, 0x04, 0x00, 0x00,
}
