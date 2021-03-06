// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpcproto/rig_assignment.proto

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

// PoolConfig holds the options for the pool to mine to
type PoolConfig struct {
	// Endpoint of the pool, format pool.coin.com:3333
	Endpoint string `protobuf:"bytes,1,opt,name=Endpoint,proto3" json:"Endpoint,omitempty"`
	// Username for the pool, usually the wallet address to mine to
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	// Password for the pool, usually the miner name
	Password string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	// Variant of the given algorithm
	Variant              string   `protobuf:"bytes,4,opt,name=Variant,proto3" json:"Variant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PoolConfig) Reset()         { *m = PoolConfig{} }
func (m *PoolConfig) String() string { return proto.CompactTextString(m) }
func (*PoolConfig) ProtoMessage()    {}
func (*PoolConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b32283fa8ae293, []int{0}
}

func (m *PoolConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoolConfig.Unmarshal(m, b)
}
func (m *PoolConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoolConfig.Marshal(b, m, deterministic)
}
func (m *PoolConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolConfig.Merge(m, src)
}
func (m *PoolConfig) XXX_Size() int {
	return xxx_messageInfo_PoolConfig.Size(m)
}
func (m *PoolConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PoolConfig proto.InternalMessageInfo

func (m *PoolConfig) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *PoolConfig) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *PoolConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *PoolConfig) GetVariant() string {
	if m != nil {
		return m.Variant
	}
	return ""
}

// CPUConfig holds the options for CPU mining
type CPUConfig struct {
	// ThreadCount is the number of CPU threads to use
	ThreadCount          uint32   `protobuf:"varint,1,opt,name=ThreadCount,proto3" json:"ThreadCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPUConfig) Reset()         { *m = CPUConfig{} }
func (m *CPUConfig) String() string { return proto.CompactTextString(m) }
func (*CPUConfig) ProtoMessage()    {}
func (*CPUConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b32283fa8ae293, []int{1}
}

func (m *CPUConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPUConfig.Unmarshal(m, b)
}
func (m *CPUConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPUConfig.Marshal(b, m, deterministic)
}
func (m *CPUConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPUConfig.Merge(m, src)
}
func (m *CPUConfig) XXX_Size() int {
	return xxx_messageInfo_CPUConfig.Size(m)
}
func (m *CPUConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CPUConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CPUConfig proto.InternalMessageInfo

func (m *CPUConfig) GetThreadCount() uint32 {
	if m != nil {
		return m.ThreadCount
	}
	return 0
}

// NvidiaGPUConfig holds the options for Nvidia GPU mining
type NvidiaGPUConfig struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NvidiaGPUConfig) Reset()         { *m = NvidiaGPUConfig{} }
func (m *NvidiaGPUConfig) String() string { return proto.CompactTextString(m) }
func (*NvidiaGPUConfig) ProtoMessage()    {}
func (*NvidiaGPUConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b32283fa8ae293, []int{2}
}

func (m *NvidiaGPUConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NvidiaGPUConfig.Unmarshal(m, b)
}
func (m *NvidiaGPUConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NvidiaGPUConfig.Marshal(b, m, deterministic)
}
func (m *NvidiaGPUConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NvidiaGPUConfig.Merge(m, src)
}
func (m *NvidiaGPUConfig) XXX_Size() int {
	return xxx_messageInfo_NvidiaGPUConfig.Size(m)
}
func (m *NvidiaGPUConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NvidiaGPUConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NvidiaGPUConfig proto.InternalMessageInfo

// AMDGPUConfig holds the options for AMD GPU mining
type AMDGPUConfig struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AMDGPUConfig) Reset()         { *m = AMDGPUConfig{} }
func (m *AMDGPUConfig) String() string { return proto.CompactTextString(m) }
func (*AMDGPUConfig) ProtoMessage()    {}
func (*AMDGPUConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b32283fa8ae293, []int{3}
}

func (m *AMDGPUConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AMDGPUConfig.Unmarshal(m, b)
}
func (m *AMDGPUConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AMDGPUConfig.Marshal(b, m, deterministic)
}
func (m *AMDGPUConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AMDGPUConfig.Merge(m, src)
}
func (m *AMDGPUConfig) XXX_Size() int {
	return xxx_messageInfo_AMDGPUConfig.Size(m)
}
func (m *AMDGPUConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AMDGPUConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AMDGPUConfig proto.InternalMessageInfo

// GPUConfig holds the options for GPU mining
type GPUConfig struct {
	// NvidiaGPUConfig holds the options for Nvidia GPU mining
	NvidiaGPUConfigs []*NvidiaGPUConfig `protobuf:"bytes,1,rep,name=NvidiaGPUConfigs,proto3" json:"NvidiaGPUConfigs,omitempty"`
	// AMDGPUConfig holds the options for AMD GPU mining
	AMDGPUConfigs        []*AMDGPUConfig `protobuf:"bytes,2,rep,name=AMDGPUConfigs,proto3" json:"AMDGPUConfigs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GPUConfig) Reset()         { *m = GPUConfig{} }
func (m *GPUConfig) String() string { return proto.CompactTextString(m) }
func (*GPUConfig) ProtoMessage()    {}
func (*GPUConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b32283fa8ae293, []int{4}
}

func (m *GPUConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GPUConfig.Unmarshal(m, b)
}
func (m *GPUConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GPUConfig.Marshal(b, m, deterministic)
}
func (m *GPUConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GPUConfig.Merge(m, src)
}
func (m *GPUConfig) XXX_Size() int {
	return xxx_messageInfo_GPUConfig.Size(m)
}
func (m *GPUConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GPUConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GPUConfig proto.InternalMessageInfo

func (m *GPUConfig) GetNvidiaGPUConfigs() []*NvidiaGPUConfig {
	if m != nil {
		return m.NvidiaGPUConfigs
	}
	return nil
}

func (m *GPUConfig) GetAMDGPUConfigs() []*AMDGPUConfig {
	if m != nil {
		return m.AMDGPUConfigs
	}
	return nil
}

// MinerConfig holds the options for configuring a miner
type MinerConfig struct {
	// Key identifies this miner's configuration on MiningHQ
	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	// Alghorithm to use when mining
	Algorithm string `protobuf:"bytes,2,opt,name=Algorithm,proto3" json:"Algorithm,omitempty"`
	// Miner is the miner to use (like xmrig, xmr-stak, etc)
	Miner string `protobuf:"bytes,3,opt,name=Miner,proto3" json:"Miner,omitempty"`
	// PoolConfig holds the options for the pool to mine to
	PoolConfig *PoolConfig `protobuf:"bytes,4,opt,name=PoolConfig,proto3" json:"PoolConfig,omitempty"`
	// CPUConfig holds the options for CPU mining
	CPUConfig *CPUConfig `protobuf:"bytes,5,opt,name=CPUConfig,proto3" json:"CPUConfig,omitempty"`
	// GPUConfig holds the options for GPU mining
	GPUConfig            *GPUConfig `protobuf:"bytes,6,opt,name=GPUConfig,proto3" json:"GPUConfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MinerConfig) Reset()         { *m = MinerConfig{} }
func (m *MinerConfig) String() string { return proto.CompactTextString(m) }
func (*MinerConfig) ProtoMessage()    {}
func (*MinerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b32283fa8ae293, []int{5}
}

func (m *MinerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MinerConfig.Unmarshal(m, b)
}
func (m *MinerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MinerConfig.Marshal(b, m, deterministic)
}
func (m *MinerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinerConfig.Merge(m, src)
}
func (m *MinerConfig) XXX_Size() int {
	return xxx_messageInfo_MinerConfig.Size(m)
}
func (m *MinerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MinerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MinerConfig proto.InternalMessageInfo

func (m *MinerConfig) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MinerConfig) GetAlgorithm() string {
	if m != nil {
		return m.Algorithm
	}
	return ""
}

func (m *MinerConfig) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *MinerConfig) GetPoolConfig() *PoolConfig {
	if m != nil {
		return m.PoolConfig
	}
	return nil
}

func (m *MinerConfig) GetCPUConfig() *CPUConfig {
	if m != nil {
		return m.CPUConfig
	}
	return nil
}

func (m *MinerConfig) GetGPUConfig() *GPUConfig {
	if m != nil {
		return m.GPUConfig
	}
	return nil
}

// RigAssignmentRequest contains the information for a miner to start mining
type RigAssignmentRequest struct {
	// Configs hold the new configurations for the miner(s)
	MinerConfigs         []*MinerConfig `protobuf:"bytes,1,rep,name=MinerConfigs,proto3" json:"MinerConfigs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RigAssignmentRequest) Reset()         { *m = RigAssignmentRequest{} }
func (m *RigAssignmentRequest) String() string { return proto.CompactTextString(m) }
func (*RigAssignmentRequest) ProtoMessage()    {}
func (*RigAssignmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b32283fa8ae293, []int{6}
}

func (m *RigAssignmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RigAssignmentRequest.Unmarshal(m, b)
}
func (m *RigAssignmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RigAssignmentRequest.Marshal(b, m, deterministic)
}
func (m *RigAssignmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RigAssignmentRequest.Merge(m, src)
}
func (m *RigAssignmentRequest) XXX_Size() int {
	return xxx_messageInfo_RigAssignmentRequest.Size(m)
}
func (m *RigAssignmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RigAssignmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RigAssignmentRequest proto.InternalMessageInfo

func (m *RigAssignmentRequest) GetMinerConfigs() []*MinerConfig {
	if m != nil {
		return m.MinerConfigs
	}
	return nil
}

// RigAssignmentResponse is returned for a RigAssignmentRequest
type RigAssignmentResponse struct {
	// StatusCode is the HTTP status code
	StatusCode uint32 `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	// Status is the text representation of StatusCode
	Status string `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	// Reason for StatusCode when StatusCode is an error
	Reason string `protobuf:"bytes,3,opt,name=Reason,proto3" json:"Reason,omitempty"`
	// MinerVersions returns the versions of the miner assigned
	MinerVersions        []string `protobuf:"bytes,4,rep,name=MinerVersions,proto3" json:"MinerVersions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RigAssignmentResponse) Reset()         { *m = RigAssignmentResponse{} }
func (m *RigAssignmentResponse) String() string { return proto.CompactTextString(m) }
func (*RigAssignmentResponse) ProtoMessage()    {}
func (*RigAssignmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98b32283fa8ae293, []int{7}
}

func (m *RigAssignmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RigAssignmentResponse.Unmarshal(m, b)
}
func (m *RigAssignmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RigAssignmentResponse.Marshal(b, m, deterministic)
}
func (m *RigAssignmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RigAssignmentResponse.Merge(m, src)
}
func (m *RigAssignmentResponse) XXX_Size() int {
	return xxx_messageInfo_RigAssignmentResponse.Size(m)
}
func (m *RigAssignmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RigAssignmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RigAssignmentResponse proto.InternalMessageInfo

func (m *RigAssignmentResponse) GetStatusCode() uint32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *RigAssignmentResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *RigAssignmentResponse) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *RigAssignmentResponse) GetMinerVersions() []string {
	if m != nil {
		return m.MinerVersions
	}
	return nil
}

func init() {
	proto.RegisterType((*PoolConfig)(nil), "rpcproto.PoolConfig")
	proto.RegisterType((*CPUConfig)(nil), "rpcproto.CPUConfig")
	proto.RegisterType((*NvidiaGPUConfig)(nil), "rpcproto.NvidiaGPUConfig")
	proto.RegisterType((*AMDGPUConfig)(nil), "rpcproto.AMDGPUConfig")
	proto.RegisterType((*GPUConfig)(nil), "rpcproto.GPUConfig")
	proto.RegisterType((*MinerConfig)(nil), "rpcproto.MinerConfig")
	proto.RegisterType((*RigAssignmentRequest)(nil), "rpcproto.RigAssignmentRequest")
	proto.RegisterType((*RigAssignmentResponse)(nil), "rpcproto.RigAssignmentResponse")
}

func init() { proto.RegisterFile("rpcproto/rig_assignment.proto", fileDescriptor_98b32283fa8ae293) }

var fileDescriptor_98b32283fa8ae293 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0x36, 0xbb, 0x65, 0x33, 0xd9, 0xc2, 0x62, 0xba, 0x2b, 0x83, 0x00, 0x45, 0x11, 0x87,
	0x5e, 0x28, 0xda, 0x85, 0x0b, 0x12, 0x97, 0x2a, 0xac, 0x38, 0xa0, 0xa2, 0x62, 0x68, 0xaf, 0xc8,
	0x10, 0x93, 0x5a, 0x6a, 0xed, 0x60, 0xbb, 0x20, 0xf8, 0x07, 0x24, 0xfe, 0x94, 0x5f, 0x40, 0x71,
	0x9c, 0xd8, 0x0d, 0xb7, 0xbc, 0xf7, 0xe6, 0xcd, 0x4c, 0xde, 0x18, 0x1e, 0xa9, 0xfa, 0x4b, 0xad,
	0xa4, 0x91, 0xcf, 0x14, 0xaf, 0x3e, 0x51, 0xad, 0x79, 0x25, 0x76, 0x4c, 0x98, 0x99, 0x25, 0xd1,
	0x69, 0x27, 0xe7, 0xbf, 0x00, 0x96, 0x52, 0x6e, 0x0b, 0x29, 0xbe, 0xf2, 0x0a, 0x3d, 0x80, 0xd3,
	0x1b, 0x51, 0xd6, 0x92, 0x0b, 0x83, 0xa3, 0x2c, 0x9a, 0x26, 0xa4, 0xc7, 0x8d, 0xb6, 0xd2, 0x4c,
	0x09, 0xba, 0x63, 0xf8, 0xa8, 0xd5, 0x3a, 0xdc, 0x68, 0x4b, 0xaa, 0xf5, 0x0f, 0xa9, 0x4a, 0x1c,
	0xb7, 0x5a, 0x87, 0x11, 0x86, 0x5b, 0x6b, 0xaa, 0x38, 0x15, 0x06, 0x1f, 0x5b, 0xa9, 0x83, 0xf9,
	0x53, 0x48, 0x8a, 0xe5, 0xca, 0x8d, 0xce, 0x20, 0xfd, 0xb8, 0x51, 0x8c, 0x96, 0x85, 0xdc, 0xbb,
	0xe9, 0x63, 0x12, 0x52, 0xf9, 0x5d, 0xb8, 0xf3, 0xee, 0x3b, 0x2f, 0x39, 0x7d, 0xd3, 0x99, 0xf2,
	0xdb, 0x70, 0x36, 0x5f, 0xbc, 0xf6, 0xf8, 0x4f, 0x04, 0x49, 0x8f, 0xd0, 0x0d, 0x9c, 0x0f, 0x0c,
	0x1a, 0x47, 0x59, 0x3c, 0x4d, 0xaf, 0xef, 0xcf, 0xba, 0x00, 0x66, 0x83, 0x0a, 0xf2, 0x9f, 0x05,
	0xbd, 0x82, 0x71, 0x38, 0x44, 0xe3, 0x23, 0xdb, 0xe3, 0xd2, 0xf7, 0x08, 0x65, 0x72, 0x58, 0x9c,
	0xff, 0x8d, 0x20, 0x5d, 0x70, 0xc1, 0x94, 0x5b, 0xea, 0x1c, 0xe2, 0xb7, 0xec, 0xa7, 0x4b, 0xb7,
	0xf9, 0x44, 0x0f, 0x21, 0x99, 0x6f, 0x2b, 0xa9, 0xb8, 0xd9, 0xec, 0x5c, 0xb2, 0x9e, 0x40, 0x13,
	0x38, 0xb1, 0x76, 0x97, 0x6b, 0x0b, 0xd0, 0x8b, 0xf0, 0x6c, 0x36, 0xd7, 0xf4, 0x7a, 0xe2, 0x17,
	0xf2, 0x1a, 0x09, 0xcf, 0x7b, 0x15, 0x04, 0x8e, 0x4f, 0xac, 0xe9, 0x9e, 0x37, 0xf5, 0x12, 0x09,
	0xce, 0x72, 0x15, 0x04, 0x8a, 0x47, 0x43, 0x8b, 0xff, 0x6b, 0x5f, 0x95, 0xbf, 0x87, 0x09, 0xe1,
	0xd5, 0xbc, 0x7f, 0x73, 0x84, 0x7d, 0xdb, 0x33, 0x6d, 0xd0, 0x4b, 0x38, 0x0b, 0x82, 0xe8, 0x4e,
	0x71, 0xe1, 0xbb, 0x05, 0x2a, 0x39, 0x28, 0xcd, 0x7f, 0x47, 0x70, 0x31, 0xe8, 0xa9, 0x6b, 0x29,
	0x34, 0x43, 0x8f, 0x01, 0x3e, 0x18, 0x6a, 0xf6, 0xba, 0x90, 0x25, 0x73, 0xaf, 0x26, 0x60, 0xd0,
	0x25, 0x8c, 0x5a, 0xe4, 0x92, 0x75, 0xa8, 0xe1, 0x09, 0xa3, 0x5a, 0x0a, 0x97, 0xab, 0x43, 0xe8,
	0x09, 0x8c, 0xed, 0xe4, 0x35, 0x53, 0x9a, 0x4b, 0xa1, 0xf1, 0x71, 0x16, 0x4f, 0x13, 0x72, 0x48,
	0x7e, 0x1e, 0xd9, 0x85, 0x9f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x25, 0xdf, 0x36, 0x6f, 0x67,
	0x03, 0x00, 0x00,
}
