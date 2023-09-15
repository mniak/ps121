// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: ping.proto

package grpc

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

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *string `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_ping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_ping_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalMessage    *string `protobuf:"bytes,1,req,name=original_message,json=originalMessage" json:"original_message,omitempty"`
	CapitalizedMessage *string `protobuf:"bytes,2,req,name=capitalized_message,json=capitalizedMessage" json:"capitalized_message,omitempty"`
	RandomNumber       *int32  `protobuf:"varint,3,req,name=random_number,json=randomNumber" json:"random_number,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_ping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_ping_proto_rawDescGZIP(), []int{1}
}

func (x *Pong) GetOriginalMessage() string {
	if x != nil && x.OriginalMessage != nil {
		return *x.OriginalMessage
	}
	return ""
}

func (x *Pong) GetCapitalizedMessage() string {
	if x != nil && x.CapitalizedMessage != nil {
		return *x.CapitalizedMessage
	}
	return ""
}

func (x *Pong) GetRandomNumber() int32 {
	if x != nil && x.RandomNumber != nil {
		return *x.RandomNumber
	}
	return 0
}

var File_ping_proto protoreflect.FileDescriptor

var file_ping_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x04,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x87,
	0x01, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x12, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x24, 0x0a, 0x06, 0x50, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x05,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x05, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x00, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63,
}

var (
	file_ping_proto_rawDescOnce sync.Once
	file_ping_proto_rawDescData = file_ping_proto_rawDesc
)

func file_ping_proto_rawDescGZIP() []byte {
	file_ping_proto_rawDescOnce.Do(func() {
		file_ping_proto_rawDescData = protoimpl.X.CompressGZIP(file_ping_proto_rawDescData)
	})
	return file_ping_proto_rawDescData
}

var file_ping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ping_proto_goTypes = []interface{}{
	(*Ping)(nil), // 0: Ping
	(*Pong)(nil), // 1: Pong
}
var file_ping_proto_depIdxs = []int32{
	0, // 0: Pinger.SendPing:input_type -> Ping
	1, // 1: Pinger.SendPing:output_type -> Pong
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ping_proto_init() }
func file_ping_proto_init() {
	if File_ping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_ping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
			RawDescriptor: file_ping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ping_proto_goTypes,
		DependencyIndexes: file_ping_proto_depIdxs,
		MessageInfos:      file_ping_proto_msgTypes,
	}.Build()
	File_ping_proto = out.File
	file_ping_proto_rawDesc = nil
	file_ping_proto_goTypes = nil
	file_ping_proto_depIdxs = nil
}
