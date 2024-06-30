// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: poruke/poruke.proto

package poruke

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

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poruke_poruke_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_poruke_poruke_proto_msgTypes[0]
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
	return file_poruke_poruke_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poruke_poruke_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_poruke_poruke_proto_msgTypes[1]
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
	return file_poruke_poruke_proto_rawDescGZIP(), []int{1}
}

func (x *Pong) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Welcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Welcome) Reset() {
	*x = Welcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poruke_poruke_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Welcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Welcome) ProtoMessage() {}

func (x *Welcome) ProtoReflect() protoreflect.Message {
	mi := &file_poruke_poruke_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Welcome.ProtoReflect.Descriptor instead.
func (*Welcome) Descriptor() ([]byte, []int) {
	return file_poruke_poruke_proto_rawDescGZIP(), []int{2}
}

func (x *Welcome) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Adrese struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Primalac   *AktorId `protobuf:"bytes,1,opt,name=primalac,proto3" json:"primalac,omitempty"`
	Posiljalac *AktorId `protobuf:"bytes,2,opt,name=posiljalac,proto3" json:"posiljalac,omitempty"`
}

func (x *Adrese) Reset() {
	*x = Adrese{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poruke_poruke_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Adrese) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Adrese) ProtoMessage() {}

func (x *Adrese) ProtoReflect() protoreflect.Message {
	mi := &file_poruke_poruke_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Adrese.ProtoReflect.Descriptor instead.
func (*Adrese) Descriptor() ([]byte, []int) {
	return file_poruke_poruke_proto_rawDescGZIP(), []int{3}
}

func (x *Adrese) GetPrimalac() *AktorId {
	if x != nil {
		return x.Primalac
	}
	return nil
}

func (x *Adrese) GetPosiljalac() *AktorId {
	if x != nil {
		return x.Posiljalac
	}
	return nil
}

type AktorId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adresa string `protobuf:"bytes,1,opt,name=adresa,proto3" json:"adresa,omitempty"`
	Port   int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Id     int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AktorId) Reset() {
	*x = AktorId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poruke_poruke_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AktorId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AktorId) ProtoMessage() {}

func (x *AktorId) ProtoReflect() protoreflect.Message {
	mi := &file_poruke_poruke_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AktorId.ProtoReflect.Descriptor instead.
func (*AktorId) Descriptor() ([]byte, []int) {
	return file_poruke_poruke_proto_rawDescGZIP(), []int{4}
}

func (x *AktorId) GetAdresa() string {
	if x != nil {
		return x.Adresa
	}
	return ""
}

func (x *AktorId) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *AktorId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Poruka struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posiljalac string `protobuf:"bytes,1,opt,name=posiljalac,proto3" json:"posiljalac,omitempty"`
	// Types that are assignable to Msg:
	//
	//	*Poruka_Ping
	//	*Poruka_Pong
	//	*Poruka_Stop
	Msg isPoruka_Msg `protobuf_oneof:"msg"`
}

func (x *Poruka) Reset() {
	*x = Poruka{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poruke_poruke_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Poruka) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Poruka) ProtoMessage() {}

func (x *Poruka) ProtoReflect() protoreflect.Message {
	mi := &file_poruke_poruke_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Poruka.ProtoReflect.Descriptor instead.
func (*Poruka) Descriptor() ([]byte, []int) {
	return file_poruke_poruke_proto_rawDescGZIP(), []int{5}
}

func (x *Poruka) GetPosiljalac() string {
	if x != nil {
		return x.Posiljalac
	}
	return ""
}

func (m *Poruka) GetMsg() isPoruka_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *Poruka) GetPing() *Ping {
	if x, ok := x.GetMsg().(*Poruka_Ping); ok {
		return x.Ping
	}
	return nil
}

func (x *Poruka) GetPong() *Pong {
	if x, ok := x.GetMsg().(*Poruka_Pong); ok {
		return x.Pong
	}
	return nil
}

func (x *Poruka) GetStop() *Stop {
	if x, ok := x.GetMsg().(*Poruka_Stop); ok {
		return x.Stop
	}
	return nil
}

type isPoruka_Msg interface {
	isPoruka_Msg()
}

type Poruka_Ping struct {
	Ping *Ping `protobuf:"bytes,2,opt,name=ping,proto3,oneof"`
}

type Poruka_Pong struct {
	Pong *Pong `protobuf:"bytes,3,opt,name=pong,proto3,oneof"`
}

type Poruka_Stop struct {
	Stop *Stop `protobuf:"bytes,4,opt,name=stop,proto3,oneof"`
}

func (*Poruka_Ping) isPoruka_Msg() {}

func (*Poruka_Pong) isPoruka_Msg() {}

func (*Poruka_Stop) isPoruka_Msg() {}

type Stop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Stop) Reset() {
	*x = Stop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poruke_poruke_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stop) ProtoMessage() {}

func (x *Stop) ProtoReflect() protoreflect.Message {
	mi := &file_poruke_poruke_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stop.ProtoReflect.Descriptor instead.
func (*Stop) Descriptor() ([]byte, []int) {
	return file_poruke_poruke_proto_rawDescGZIP(), []int{6}
}

var File_poruke_poruke_proto protoreflect.FileDescriptor

var file_poruke_poruke_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6f, 0x72, 0x75, 0x6b, 0x65, 0x2f, 0x70, 0x6f, 0x72, 0x75, 0x6b, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x6f, 0x72, 0x75, 0x6b, 0x65, 0x22, 0x16, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x19, 0x0a,
	0x07, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x66, 0x0a, 0x06, 0x41, 0x64, 0x72, 0x65,
	0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x6c, 0x61, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x72, 0x75, 0x6b, 0x65, 0x2e, 0x41, 0x6b,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x6c, 0x61, 0x63, 0x12,
	0x2f, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x6c, 0x6a, 0x61, 0x6c, 0x61, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6f, 0x72, 0x75, 0x6b, 0x65, 0x2e, 0x41, 0x6b, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x6c, 0x6a, 0x61, 0x6c, 0x61, 0x63,
	0x22, 0x45, 0x0a, 0x07, 0x41, 0x6b, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x64, 0x72, 0x65, 0x73, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x64, 0x72,
	0x65, 0x73, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9b, 0x01, 0x0a, 0x06, 0x50, 0x6f, 0x72, 0x75,
	0x6b, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x6c, 0x6a, 0x61, 0x6c, 0x61, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x6c, 0x6a, 0x61, 0x6c,
	0x61, 0x63, 0x12, 0x22, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x6f, 0x72, 0x75, 0x6b, 0x65, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x48, 0x00,
	0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6f, 0x72, 0x75, 0x6b, 0x65, 0x2e, 0x50, 0x6f,
	0x6e, 0x67, 0x48, 0x00, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x04, 0x73, 0x74,
	0x6f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x6f, 0x72, 0x75, 0x6b,
	0x65, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x48, 0x00, 0x52, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x42, 0x05,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x06, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x42, 0x0d, 0x5a,
	0x0b, 0x41, 0x54, 0x32, 0x34, 0x2f, 0x70, 0x6f, 0x72, 0x75, 0x6b, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_poruke_poruke_proto_rawDescOnce sync.Once
	file_poruke_poruke_proto_rawDescData = file_poruke_poruke_proto_rawDesc
)

func file_poruke_poruke_proto_rawDescGZIP() []byte {
	file_poruke_poruke_proto_rawDescOnce.Do(func() {
		file_poruke_poruke_proto_rawDescData = protoimpl.X.CompressGZIP(file_poruke_poruke_proto_rawDescData)
	})
	return file_poruke_poruke_proto_rawDescData
}

var file_poruke_poruke_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_poruke_poruke_proto_goTypes = []interface{}{
	(*Ping)(nil),    // 0: poruke.Ping
	(*Pong)(nil),    // 1: poruke.Pong
	(*Welcome)(nil), // 2: poruke.Welcome
	(*Adrese)(nil),  // 3: poruke.Adrese
	(*AktorId)(nil), // 4: poruke.AktorId
	(*Poruka)(nil),  // 5: poruke.Poruka
	(*Stop)(nil),    // 6: poruke.Stop
}
var file_poruke_poruke_proto_depIdxs = []int32{
	4, // 0: poruke.Adrese.primalac:type_name -> poruke.AktorId
	4, // 1: poruke.Adrese.posiljalac:type_name -> poruke.AktorId
	0, // 2: poruke.Poruka.ping:type_name -> poruke.Ping
	1, // 3: poruke.Poruka.pong:type_name -> poruke.Pong
	6, // 4: poruke.Poruka.stop:type_name -> poruke.Stop
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_poruke_poruke_proto_init() }
func file_poruke_poruke_proto_init() {
	if File_poruke_poruke_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_poruke_poruke_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_poruke_poruke_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_poruke_poruke_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Welcome); i {
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
		file_poruke_poruke_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Adrese); i {
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
		file_poruke_poruke_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AktorId); i {
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
		file_poruke_poruke_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Poruka); i {
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
		file_poruke_poruke_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stop); i {
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
	file_poruke_poruke_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Poruka_Ping)(nil),
		(*Poruka_Pong)(nil),
		(*Poruka_Stop)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_poruke_poruke_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_poruke_poruke_proto_goTypes,
		DependencyIndexes: file_poruke_poruke_proto_depIdxs,
		MessageInfos:      file_poruke_poruke_proto_msgTypes,
	}.Build()
	File_poruke_poruke_proto = out.File
	file_poruke_poruke_proto_rawDesc = nil
	file_poruke_poruke_proto_goTypes = nil
	file_poruke_poruke_proto_depIdxs = nil
}