// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: dropsonde-protocol/events/uuid.proto

package events

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

// Type representing a 128-bit UUID.
//
// The bytes of the UUID should be packed in little-endian **byte** (not bit)
// order. For example, the UUID `f47ac10b-58cc-4372-a567-0e02b2c3d479` should
// be encoded as `UUID{ low: 0x7243cc580bc17af4, high: 0x79d4c3b2020e67a5 }`
type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Low  *uint64 `protobuf:"varint,1,req,name=low" json:"low,omitempty"`
	High *uint64 `protobuf:"varint,2,req,name=high" json:"high,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dropsonde_protocol_events_uuid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_dropsonde_protocol_events_uuid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_dropsonde_protocol_events_uuid_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetLow() uint64 {
	if x != nil && x.Low != nil {
		return *x.Low
	}
	return 0
}

func (x *UUID) GetHigh() uint64 {
	if x != nil && x.High != nil {
		return *x.High
	}
	return 0
}

var File_dropsonde_protocol_events_uuid_proto protoreflect.FileDescriptor

var file_dropsonde_protocol_events_uuid_proto_rawDesc = []byte{
	0x0a, 0x24, 0x64, 0x72, 0x6f, 0x70, 0x73, 0x6f, 0x6e, 0x64, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x75, 0x75, 0x69, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x2c,
	0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x04, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x42, 0x59, 0x0a, 0x21,
	0x6f, 0x72, 0x67, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79,
	0x2e, 0x64, 0x72, 0x6f, 0x70, 0x73, 0x6f, 0x6e, 0x64, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x42, 0x0b, 0x55, 0x75, 0x69, 0x64, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5a, 0x27,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x2f, 0x73, 0x6f, 0x6e, 0x64, 0x65, 0x2d, 0x67, 0x6f,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
}

var (
	file_dropsonde_protocol_events_uuid_proto_rawDescOnce sync.Once
	file_dropsonde_protocol_events_uuid_proto_rawDescData = file_dropsonde_protocol_events_uuid_proto_rawDesc
)

func file_dropsonde_protocol_events_uuid_proto_rawDescGZIP() []byte {
	file_dropsonde_protocol_events_uuid_proto_rawDescOnce.Do(func() {
		file_dropsonde_protocol_events_uuid_proto_rawDescData = protoimpl.X.CompressGZIP(file_dropsonde_protocol_events_uuid_proto_rawDescData)
	})
	return file_dropsonde_protocol_events_uuid_proto_rawDescData
}

var file_dropsonde_protocol_events_uuid_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_dropsonde_protocol_events_uuid_proto_goTypes = []interface{}{
	(*UUID)(nil), // 0: events.UUID
}
var file_dropsonde_protocol_events_uuid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dropsonde_protocol_events_uuid_proto_init() }
func file_dropsonde_protocol_events_uuid_proto_init() {
	if File_dropsonde_protocol_events_uuid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dropsonde_protocol_events_uuid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
			RawDescriptor: file_dropsonde_protocol_events_uuid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dropsonde_protocol_events_uuid_proto_goTypes,
		DependencyIndexes: file_dropsonde_protocol_events_uuid_proto_depIdxs,
		MessageInfos:      file_dropsonde_protocol_events_uuid_proto_msgTypes,
	}.Build()
	File_dropsonde_protocol_events_uuid_proto = out.File
	file_dropsonde_protocol_events_uuid_proto_rawDesc = nil
	file_dropsonde_protocol_events_uuid_proto_goTypes = nil
	file_dropsonde_protocol_events_uuid_proto_depIdxs = nil
}
