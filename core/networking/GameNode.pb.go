// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: GameNode.proto

// Version 4

package networking

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

// Client requests available lobbies
type LobbyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VersionMain int32 `protobuf:"varint,1,opt,name=version_main,json=versionMain,proto3" json:"version_main,omitempty"`
	VersionMin  int32 `protobuf:"varint,2,opt,name=version_min,json=versionMin,proto3" json:"version_min,omitempty"` // Possibly an return address?
}

func (x *LobbyRequest) Reset() {
	*x = LobbyRequest{}
	mi := &file_GameNode_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LobbyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LobbyRequest) ProtoMessage() {}

func (x *LobbyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_GameNode_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LobbyRequest.ProtoReflect.Descriptor instead.
func (*LobbyRequest) Descriptor() ([]byte, []int) {
	return file_GameNode_proto_rawDescGZIP(), []int{0}
}

func (x *LobbyRequest) GetVersionMain() int32 {
	if x != nil {
		return x.VersionMain
	}
	return 0
}

func (x *LobbyRequest) GetVersionMin() int32 {
	if x != nil {
		return x.VersionMin
	}
	return 0
}

// Active servers respond
type LobbyProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner     *Player `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	IsFull    bool    `protobuf:"varint,4,opt,name=is_full,json=isFull,proto3" json:"is_full,omitempty"`
	IsRunning bool    `protobuf:"varint,5,opt,name=is_running,json=isRunning,proto3" json:"is_running,omitempty"`
}

func (x *LobbyProposal) Reset() {
	*x = LobbyProposal{}
	mi := &file_GameNode_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LobbyProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LobbyProposal) ProtoMessage() {}

func (x *LobbyProposal) ProtoReflect() protoreflect.Message {
	mi := &file_GameNode_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LobbyProposal.ProtoReflect.Descriptor instead.
func (*LobbyProposal) Descriptor() ([]byte, []int) {
	return file_GameNode_proto_rawDescGZIP(), []int{1}
}

func (x *LobbyProposal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LobbyProposal) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LobbyProposal) GetOwner() *Player {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *LobbyProposal) GetIsFull() bool {
	if x != nil {
		return x.IsFull
	}
	return false
}

func (x *LobbyProposal) GetIsRunning() bool {
	if x != nil {
		return x.IsRunning
	}
	return false
}

type JoinLobbyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *JoinLobbyRequest) Reset() {
	*x = JoinLobbyRequest{}
	mi := &file_GameNode_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinLobbyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinLobbyRequest) ProtoMessage() {}

func (x *JoinLobbyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_GameNode_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinLobbyRequest.ProtoReflect.Descriptor instead.
func (*JoinLobbyRequest) Descriptor() ([]byte, []int) {
	return file_GameNode_proto_rawDescGZIP(), []int{2}
}

func (x *JoinLobbyRequest) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type JoinLobbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *JoinLobbyResponse) Reset() {
	*x = JoinLobbyResponse{}
	mi := &file_GameNode_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinLobbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinLobbyResponse) ProtoMessage() {}

func (x *JoinLobbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_GameNode_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinLobbyResponse.ProtoReflect.Descriptor instead.
func (*JoinLobbyResponse) Descriptor() ([]byte, []int) {
	return file_GameNode_proto_rawDescGZIP(), []int{3}
}

func (x *JoinLobbyResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JoinLobbyResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type RegisterOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RegisterOwnerRequest) Reset() {
	*x = RegisterOwnerRequest{}
	mi := &file_GameNode_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterOwnerRequest) ProtoMessage() {}

func (x *RegisterOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_GameNode_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterOwnerRequest.ProtoReflect.Descriptor instead.
func (*RegisterOwnerRequest) Descriptor() ([]byte, []int) {
	return file_GameNode_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterOwnerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RegisterOwnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accepted bool `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
}

func (x *RegisterOwnerResponse) Reset() {
	*x = RegisterOwnerResponse{}
	mi := &file_GameNode_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterOwnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterOwnerResponse) ProtoMessage() {}

func (x *RegisterOwnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_GameNode_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterOwnerResponse.ProtoReflect.Descriptor instead.
func (*RegisterOwnerResponse) Descriptor() ([]byte, []int) {
	return file_GameNode_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterOwnerResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Shape    int32  `protobuf:"varint,3,opt,name=Shape,proto3" json:"Shape,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	mi := &file_GameNode_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_GameNode_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_GameNode_proto_rawDescGZIP(), []int{6}
}

func (x *Player) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Player) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Player) GetShape() int32 {
	if x != nil {
		return x.Shape
	}
	return 0
}

var File_GameNode_proto protoreflect.FileDescriptor

var file_GameNode_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x47, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x52, 0x0a, 0x0c,
	0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x69, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6e,
	0x22, 0x95, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x69, 0x73, 0x46, 0x75, 0x6c, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x3e, 0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e,
	0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x35, 0x0a, 0x11, 0x4a, 0x6f, 0x69, 0x6e,
	0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x26, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x22, 0x4a, 0x0a, 0x06,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x68, 0x61, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x53, 0x68, 0x61, 0x70, 0x65, 0x32, 0x8a, 0x02, 0x0a, 0x12, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4d, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x4c, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4c,
	0x6f, 0x62, 0x62, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x30, 0x01, 0x12, 0x4f,
	0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4a, 0x6f,
	0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x4a, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x54, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x20, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GameNode_proto_rawDescOnce sync.Once
	file_GameNode_proto_rawDescData = file_GameNode_proto_rawDesc
)

func file_GameNode_proto_rawDescGZIP() []byte {
	file_GameNode_proto_rawDescOnce.Do(func() {
		file_GameNode_proto_rawDescData = protoimpl.X.CompressGZIP(file_GameNode_proto_rawDescData)
	})
	return file_GameNode_proto_rawDescData
}

var file_GameNode_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_GameNode_proto_goTypes = []any{
	(*LobbyRequest)(nil),          // 0: networking.LobbyRequest
	(*LobbyProposal)(nil),         // 1: networking.LobbyProposal
	(*JoinLobbyRequest)(nil),      // 2: networking.JoinLobbyRequest
	(*JoinLobbyResponse)(nil),     // 3: networking.JoinLobbyResponse
	(*RegisterOwnerRequest)(nil),  // 4: networking.RegisterOwnerRequest
	(*RegisterOwnerResponse)(nil), // 5: networking.RegisterOwnerResponse
	(*Player)(nil),                // 6: networking.Player
}
var file_GameNode_proto_depIdxs = []int32{
	6, // 0: networking.LobbyProposal.owner:type_name -> networking.Player
	6, // 1: networking.JoinLobbyRequest.player:type_name -> networking.Player
	0, // 2: networking.LobbySearchService.RequestActiveLobbies:input_type -> networking.LobbyRequest
	2, // 3: networking.LobbySearchService.RequestLobbyJoin:input_type -> networking.JoinLobbyRequest
	4, // 4: networking.LobbySearchService.RegisterOwner:input_type -> networking.RegisterOwnerRequest
	1, // 5: networking.LobbySearchService.RequestActiveLobbies:output_type -> networking.LobbyProposal
	3, // 6: networking.LobbySearchService.RequestLobbyJoin:output_type -> networking.JoinLobbyResponse
	5, // 7: networking.LobbySearchService.RegisterOwner:output_type -> networking.RegisterOwnerResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GameNode_proto_init() }
func file_GameNode_proto_init() {
	if File_GameNode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GameNode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_GameNode_proto_goTypes,
		DependencyIndexes: file_GameNode_proto_depIdxs,
		MessageInfos:      file_GameNode_proto_msgTypes,
	}.Build()
	File_GameNode_proto = out.File
	file_GameNode_proto_rawDesc = nil
	file_GameNode_proto_goTypes = nil
	file_GameNode_proto_depIdxs = nil
}
