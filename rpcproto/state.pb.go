// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpcproto/state.proto

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

// StateRequest updates a miner's state
type StateRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateRequest) Reset()         { *m = StateRequest{} }
func (m *StateRequest) String() string { return proto.CompactTextString(m) }
func (*StateRequest) ProtoMessage()    {}
func (*StateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd92d90f4dceef24, []int{0}
}

func (m *StateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateRequest.Unmarshal(m, b)
}
func (m *StateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateRequest.Marshal(b, m, deterministic)
}
func (m *StateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateRequest.Merge(m, src)
}
func (m *StateRequest) XXX_Size() int {
	return xxx_messageInfo_StateRequest.Size(m)
}
func (m *StateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StateRequest proto.InternalMessageInfo

// StateResponse is returned for a StateRequest
type StateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateResponse) Reset()         { *m = StateResponse{} }
func (m *StateResponse) String() string { return proto.CompactTextString(m) }
func (*StateResponse) ProtoMessage()    {}
func (*StateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd92d90f4dceef24, []int{1}
}

func (m *StateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateResponse.Unmarshal(m, b)
}
func (m *StateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateResponse.Marshal(b, m, deterministic)
}
func (m *StateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateResponse.Merge(m, src)
}
func (m *StateResponse) XXX_Size() int {
	return xxx_messageInfo_StateResponse.Size(m)
}
func (m *StateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StateRequest)(nil), "rpcproto.StateRequest")
	proto.RegisterType((*StateResponse)(nil), "rpcproto.StateResponse")
}

func init() { proto.RegisterFile("rpcproto/state.proto", fileDescriptor_bd92d90f4dceef24) }

var fileDescriptor_bd92d90f4dceef24 = []byte{
	// 77 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x2a, 0x48, 0x2e,
	0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2e, 0x49, 0x2c, 0x49, 0xd5, 0x03, 0xb3, 0x85, 0x38, 0x60,
	0xa2, 0x4a, 0x7c, 0x5c, 0x3c, 0xc1, 0x20, 0x89, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x25,
	0x7e, 0x2e, 0x5e, 0x28, 0xbf, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x35, 0x89, 0x0d, 0xac, 0xce, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x76, 0x12, 0x2d, 0x49, 0x00, 0x00, 0x00,
}