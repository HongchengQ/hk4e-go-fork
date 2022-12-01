// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: PersonalLineAllDataRsp.proto

package proto

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

// CmdId: 476
// EnetChannelId: 0
// EnetIsReliable: true
type PersonalLineAllDataRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurFinishedDailyTaskCount     uint32                    `protobuf:"varint,5,opt,name=cur_finished_daily_task_count,json=curFinishedDailyTaskCount,proto3" json:"cur_finished_daily_task_count,omitempty"`
	CanBeUnlockedPersonalLineList []uint32                  `protobuf:"varint,13,rep,packed,name=can_be_unlocked_personal_line_list,json=canBeUnlockedPersonalLineList,proto3" json:"can_be_unlocked_personal_line_list,omitempty"`
	Retcode                       int32                     `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
	OngoingPersonalLineList       []uint32                  `protobuf:"varint,8,rep,packed,name=ongoing_personal_line_list,json=ongoingPersonalLineList,proto3" json:"ongoing_personal_line_list,omitempty"`
	LegendaryKeyCount             uint32                    `protobuf:"varint,11,opt,name=legendary_key_count,json=legendaryKeyCount,proto3" json:"legendary_key_count,omitempty"`
	LockedPersonalLineList        []*LockedPersonallineData `protobuf:"bytes,10,rep,name=locked_personal_line_list,json=lockedPersonalLineList,proto3" json:"locked_personal_line_list,omitempty"`
}

func (x *PersonalLineAllDataRsp) Reset() {
	*x = PersonalLineAllDataRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PersonalLineAllDataRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonalLineAllDataRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonalLineAllDataRsp) ProtoMessage() {}

func (x *PersonalLineAllDataRsp) ProtoReflect() protoreflect.Message {
	mi := &file_PersonalLineAllDataRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonalLineAllDataRsp.ProtoReflect.Descriptor instead.
func (*PersonalLineAllDataRsp) Descriptor() ([]byte, []int) {
	return file_PersonalLineAllDataRsp_proto_rawDescGZIP(), []int{0}
}

func (x *PersonalLineAllDataRsp) GetCurFinishedDailyTaskCount() uint32 {
	if x != nil {
		return x.CurFinishedDailyTaskCount
	}
	return 0
}

func (x *PersonalLineAllDataRsp) GetCanBeUnlockedPersonalLineList() []uint32 {
	if x != nil {
		return x.CanBeUnlockedPersonalLineList
	}
	return nil
}

func (x *PersonalLineAllDataRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *PersonalLineAllDataRsp) GetOngoingPersonalLineList() []uint32 {
	if x != nil {
		return x.OngoingPersonalLineList
	}
	return nil
}

func (x *PersonalLineAllDataRsp) GetLegendaryKeyCount() uint32 {
	if x != nil {
		return x.LegendaryKeyCount
	}
	return 0
}

func (x *PersonalLineAllDataRsp) GetLockedPersonalLineList() []*LockedPersonallineData {
	if x != nil {
		return x.LockedPersonalLineList
	}
	return nil
}

var File_PersonalLineAllDataRsp_proto protoreflect.FileDescriptor

var file_PersonalLineAllDataRsp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x65, 0x41, 0x6c,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x86, 0x03, 0x0a, 0x16, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x4c, 0x69, 0x6e, 0x65, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x73, 0x70, 0x12, 0x40,
	0x0a, 0x1d, 0x63, 0x75, 0x72, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x64,
	0x61, 0x69, 0x6c, 0x79, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x19, 0x63, 0x75, 0x72, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x49, 0x0a, 0x22, 0x63, 0x61, 0x6e, 0x5f, 0x62, 0x65, 0x5f, 0x75, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x1d, 0x63, 0x61,
	0x6e, 0x42, 0x65, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x1a, 0x6f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x17, 0x6f, 0x6e, 0x67, 0x6f, 0x69,
	0x6e, 0x67, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x11, 0x6c, 0x65, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x58, 0x0a, 0x19, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x16, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PersonalLineAllDataRsp_proto_rawDescOnce sync.Once
	file_PersonalLineAllDataRsp_proto_rawDescData = file_PersonalLineAllDataRsp_proto_rawDesc
)

func file_PersonalLineAllDataRsp_proto_rawDescGZIP() []byte {
	file_PersonalLineAllDataRsp_proto_rawDescOnce.Do(func() {
		file_PersonalLineAllDataRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_PersonalLineAllDataRsp_proto_rawDescData)
	})
	return file_PersonalLineAllDataRsp_proto_rawDescData
}

var file_PersonalLineAllDataRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PersonalLineAllDataRsp_proto_goTypes = []interface{}{
	(*PersonalLineAllDataRsp)(nil), // 0: proto.PersonalLineAllDataRsp
	(*LockedPersonallineData)(nil), // 1: proto.LockedPersonallineData
}
var file_PersonalLineAllDataRsp_proto_depIdxs = []int32{
	1, // 0: proto.PersonalLineAllDataRsp.locked_personal_line_list:type_name -> proto.LockedPersonallineData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PersonalLineAllDataRsp_proto_init() }
func file_PersonalLineAllDataRsp_proto_init() {
	if File_PersonalLineAllDataRsp_proto != nil {
		return
	}
	file_LockedPersonallineData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PersonalLineAllDataRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonalLineAllDataRsp); i {
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
			RawDescriptor: file_PersonalLineAllDataRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PersonalLineAllDataRsp_proto_goTypes,
		DependencyIndexes: file_PersonalLineAllDataRsp_proto_depIdxs,
		MessageInfos:      file_PersonalLineAllDataRsp_proto_msgTypes,
	}.Build()
	File_PersonalLineAllDataRsp_proto = out.File
	file_PersonalLineAllDataRsp_proto_rawDesc = nil
	file_PersonalLineAllDataRsp_proto_goTypes = nil
	file_PersonalLineAllDataRsp_proto_depIdxs = nil
}