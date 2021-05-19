//
// (C) Copyright 2019-2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

// This file defines the messages used by DRPC_MODULE_SRV.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: srv/srv.proto

package srv

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NotifyReadyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri              string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`                           // CaRT URI
	Nctxs            uint32 `protobuf:"varint,2,opt,name=nctxs,proto3" json:"nctxs,omitempty"`                      // Number of CaRT contexts
	DrpcListenerSock string `protobuf:"bytes,3,opt,name=drpcListenerSock,proto3" json:"drpcListenerSock,omitempty"` // Path to I/O Engine's dRPC listener socket
	InstanceIdx      uint32 `protobuf:"varint,4,opt,name=instanceIdx,proto3" json:"instanceIdx,omitempty"`          // I/O Engine instance index
	Ntgts            uint32 `protobuf:"varint,5,opt,name=ntgts,proto3" json:"ntgts,omitempty"`                      // number of VOS targets allocated in I/O Engine
}

func (x *NotifyReadyReq) Reset() {
	*x = NotifyReadyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_srv_srv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyReadyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyReadyReq) ProtoMessage() {}

func (x *NotifyReadyReq) ProtoReflect() protoreflect.Message {
	mi := &file_srv_srv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyReadyReq.ProtoReflect.Descriptor instead.
func (*NotifyReadyReq) Descriptor() ([]byte, []int) {
	return file_srv_srv_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyReadyReq) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *NotifyReadyReq) GetNctxs() uint32 {
	if x != nil {
		return x.Nctxs
	}
	return 0
}

func (x *NotifyReadyReq) GetDrpcListenerSock() string {
	if x != nil {
		return x.DrpcListenerSock
	}
	return ""
}

func (x *NotifyReadyReq) GetInstanceIdx() uint32 {
	if x != nil {
		return x.InstanceIdx
	}
	return 0
}

func (x *NotifyReadyReq) GetNtgts() uint32 {
	if x != nil {
		return x.Ntgts
	}
	return 0
}

type BioErrorReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnmapErr         bool   `protobuf:"varint,1,opt,name=unmapErr,proto3" json:"unmapErr,omitempty"`                // unmap I/O error
	ReadErr          bool   `protobuf:"varint,2,opt,name=readErr,proto3" json:"readErr,omitempty"`                  // read I/O error
	WriteErr         bool   `protobuf:"varint,3,opt,name=writeErr,proto3" json:"writeErr,omitempty"`                // write I/O error
	TgtId            int32  `protobuf:"varint,4,opt,name=tgtId,proto3" json:"tgtId,omitempty"`                      // VOS target ID
	InstanceIdx      uint32 `protobuf:"varint,5,opt,name=instanceIdx,proto3" json:"instanceIdx,omitempty"`          // I/O Engine instance index
	DrpcListenerSock string `protobuf:"bytes,6,opt,name=drpcListenerSock,proto3" json:"drpcListenerSock,omitempty"` // Path to I/O Engine's dRPC listener socket
	Uri              string `protobuf:"bytes,7,opt,name=uri,proto3" json:"uri,omitempty"`                           // CaRT URI
}

func (x *BioErrorReq) Reset() {
	*x = BioErrorReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_srv_srv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BioErrorReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BioErrorReq) ProtoMessage() {}

func (x *BioErrorReq) ProtoReflect() protoreflect.Message {
	mi := &file_srv_srv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BioErrorReq.ProtoReflect.Descriptor instead.
func (*BioErrorReq) Descriptor() ([]byte, []int) {
	return file_srv_srv_proto_rawDescGZIP(), []int{1}
}

func (x *BioErrorReq) GetUnmapErr() bool {
	if x != nil {
		return x.UnmapErr
	}
	return false
}

func (x *BioErrorReq) GetReadErr() bool {
	if x != nil {
		return x.ReadErr
	}
	return false
}

func (x *BioErrorReq) GetWriteErr() bool {
	if x != nil {
		return x.WriteErr
	}
	return false
}

func (x *BioErrorReq) GetTgtId() int32 {
	if x != nil {
		return x.TgtId
	}
	return 0
}

func (x *BioErrorReq) GetInstanceIdx() uint32 {
	if x != nil {
		return x.InstanceIdx
	}
	return 0
}

func (x *BioErrorReq) GetDrpcListenerSock() string {
	if x != nil {
		return x.DrpcListenerSock
	}
	return ""
}

func (x *BioErrorReq) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type GetPoolSvcReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // Pool UUID
}

func (x *GetPoolSvcReq) Reset() {
	*x = GetPoolSvcReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_srv_srv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPoolSvcReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPoolSvcReq) ProtoMessage() {}

func (x *GetPoolSvcReq) ProtoReflect() protoreflect.Message {
	mi := &file_srv_srv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPoolSvcReq.ProtoReflect.Descriptor instead.
func (*GetPoolSvcReq) Descriptor() ([]byte, []int) {
	return file_srv_srv_proto_rawDescGZIP(), []int{2}
}

func (x *GetPoolSvcReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetPoolSvcResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`          // DAOS error code
	Svcreps []uint32 `protobuf:"varint,2,rep,packed,name=svcreps,proto3" json:"svcreps,omitempty"` // Pool service replica ranks
}

func (x *GetPoolSvcResp) Reset() {
	*x = GetPoolSvcResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_srv_srv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPoolSvcResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPoolSvcResp) ProtoMessage() {}

func (x *GetPoolSvcResp) ProtoReflect() protoreflect.Message {
	mi := &file_srv_srv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPoolSvcResp.ProtoReflect.Descriptor instead.
func (*GetPoolSvcResp) Descriptor() ([]byte, []int) {
	return file_srv_srv_proto_rawDescGZIP(), []int{3}
}

func (x *GetPoolSvcResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetPoolSvcResp) GetSvcreps() []uint32 {
	if x != nil {
		return x.Svcreps
	}
	return nil
}

var File_srv_srv_proto protoreflect.FileDescriptor

var file_srv_srv_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x72, 0x76, 0x2f, 0x73, 0x72, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x73, 0x72, 0x76, 0x22, 0x9c, 0x01, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x63, 0x74,
	0x78, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e, 0x63, 0x74, 0x78, 0x73, 0x12,
	0x2a, 0x0a, 0x10, 0x64, 0x72, 0x70, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53,
	0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x72, 0x70, 0x63, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x6f, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x78, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x74, 0x67, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e, 0x74,
	0x67, 0x74, 0x73, 0x22, 0xd5, 0x01, 0x0a, 0x0b, 0x42, 0x69, 0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x6d, 0x61, 0x70, 0x45, 0x72, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x75, 0x6e, 0x6d, 0x61, 0x70, 0x45, 0x72, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x64, 0x45, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x72, 0x65, 0x61, 0x64, 0x45, 0x72, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x45, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x45, 0x72, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x67, 0x74, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x67, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x78, 0x12, 0x2a, 0x0a,
	0x10, 0x64, 0x72, 0x70, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x6f, 0x63,
	0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x72, 0x70, 0x63, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x6f, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x22, 0x23, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x76, 0x63, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x22, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x76, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x76,
	0x63, 0x72, 0x65, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x76, 0x63,
	0x72, 0x65, 0x70, 0x73, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x6f, 0x73, 0x2d, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x64, 0x61,
	0x6f, 0x73, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x76, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_srv_srv_proto_rawDescOnce sync.Once
	file_srv_srv_proto_rawDescData = file_srv_srv_proto_rawDesc
)

func file_srv_srv_proto_rawDescGZIP() []byte {
	file_srv_srv_proto_rawDescOnce.Do(func() {
		file_srv_srv_proto_rawDescData = protoimpl.X.CompressGZIP(file_srv_srv_proto_rawDescData)
	})
	return file_srv_srv_proto_rawDescData
}

var file_srv_srv_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_srv_srv_proto_goTypes = []interface{}{
	(*NotifyReadyReq)(nil), // 0: srv.NotifyReadyReq
	(*BioErrorReq)(nil),    // 1: srv.BioErrorReq
	(*GetPoolSvcReq)(nil),  // 2: srv.GetPoolSvcReq
	(*GetPoolSvcResp)(nil), // 3: srv.GetPoolSvcResp
}
var file_srv_srv_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_srv_srv_proto_init() }
func file_srv_srv_proto_init() {
	if File_srv_srv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_srv_srv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyReadyReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_srv_srv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BioErrorReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_srv_srv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPoolSvcReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_srv_srv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPoolSvcResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_srv_srv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_srv_srv_proto_goTypes,
		DependencyIndexes: file_srv_srv_proto_depIdxs,
		MessageInfos:      file_srv_srv_proto_msgTypes,
	}.Build()
	File_srv_srv_proto = out.File
	file_srv_srv_proto_rawDesc = nil
	file_srv_srv_proto_goTypes = nil
	file_srv_srv_proto_depIdxs = nil
}
