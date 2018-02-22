// Code generated by protoc-gen-go. DO NOT EDIT.
// source: IRMAIssueSpecification.proto

/*
Package irmaproto is a generated protocol buffer package.

It is generated from these files:
	IRMAIssueSpecification.proto
	IRMAIssuerMetadata.proto
	IRMAIssuerPublicKey.proto
	IRMASchemeMetadata.proto

It has these top-level messages:
	IRMAIssueSpecification
	IRMAIssuerMetadata
	IRMAIssuerPublicKey
	IRMASchemeMetadata
*/
package irmaproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// This mimicks IRMA issuer description.xml
type IRMAIssueSpecification struct {
	Version           int32                                          `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	Name              []*IRMAIssueSpecification_LocalizedName        `protobuf:"bytes,2,rep,name=Name" json:"Name,omitempty"`
	ShortName         []*IRMAIssueSpecification_LocalizedName        `protobuf:"bytes,3,rep,name=ShortName" json:"ShortName,omitempty"`
	SchemeManager     string                                         `protobuf:"bytes,4,opt,name=SchemeManager" json:"SchemeManager,omitempty"`
	IssuerId          string                                         `protobuf:"bytes,5,opt,name=IssuerId" json:"IssuerId,omitempty"`
	CredentialId      string                                         `protobuf:"bytes,6,opt,name=CredentialId" json:"CredentialId,omitempty"`
	Description       []*IRMAIssueSpecification_LocalizedDescription `protobuf:"bytes,7,rep,name=Description" json:"Description,omitempty"`
	ShouldBeSingleton bool                                           `protobuf:"varint,8,opt,name=ShouldBeSingleton" json:"ShouldBeSingleton,omitempty"`
	Attributes        []*IRMAIssueSpecification_Attribute            `protobuf:"bytes,9,rep,name=Attributes" json:"Attributes,omitempty"`
}

func (m *IRMAIssueSpecification) Reset()                    { *m = IRMAIssueSpecification{} }
func (m *IRMAIssueSpecification) String() string            { return proto.CompactTextString(m) }
func (*IRMAIssueSpecification) ProtoMessage()               {}
func (*IRMAIssueSpecification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IRMAIssueSpecification) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *IRMAIssueSpecification) GetName() []*IRMAIssueSpecification_LocalizedName {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *IRMAIssueSpecification) GetShortName() []*IRMAIssueSpecification_LocalizedName {
	if m != nil {
		return m.ShortName
	}
	return nil
}

func (m *IRMAIssueSpecification) GetSchemeManager() string {
	if m != nil {
		return m.SchemeManager
	}
	return ""
}

func (m *IRMAIssueSpecification) GetIssuerId() string {
	if m != nil {
		return m.IssuerId
	}
	return ""
}

func (m *IRMAIssueSpecification) GetCredentialId() string {
	if m != nil {
		return m.CredentialId
	}
	return ""
}

func (m *IRMAIssueSpecification) GetDescription() []*IRMAIssueSpecification_LocalizedDescription {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *IRMAIssueSpecification) GetShouldBeSingleton() bool {
	if m != nil {
		return m.ShouldBeSingleton
	}
	return false
}

func (m *IRMAIssueSpecification) GetAttributes() []*IRMAIssueSpecification_Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type IRMAIssueSpecification_LocalizedName struct {
	Lang string `protobuf:"bytes,1,opt,name=Lang" json:"Lang,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
}

func (m *IRMAIssueSpecification_LocalizedName) Reset()         { *m = IRMAIssueSpecification_LocalizedName{} }
func (m *IRMAIssueSpecification_LocalizedName) String() string { return proto.CompactTextString(m) }
func (*IRMAIssueSpecification_LocalizedName) ProtoMessage()    {}
func (*IRMAIssueSpecification_LocalizedName) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

func (m *IRMAIssueSpecification_LocalizedName) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *IRMAIssueSpecification_LocalizedName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type IRMAIssueSpecification_LocalizedDescription struct {
	Lang string `protobuf:"bytes,1,opt,name=Lang" json:"Lang,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
}

func (m *IRMAIssueSpecification_LocalizedDescription) Reset() {
	*m = IRMAIssueSpecification_LocalizedDescription{}
}
func (m *IRMAIssueSpecification_LocalizedDescription) String() string {
	return proto.CompactTextString(m)
}
func (*IRMAIssueSpecification_LocalizedDescription) ProtoMessage() {}
func (*IRMAIssueSpecification_LocalizedDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1}
}

func (m *IRMAIssueSpecification_LocalizedDescription) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *IRMAIssueSpecification_LocalizedDescription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type IRMAIssueSpecification_Attribute struct {
	Id          string                                         `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Name        []*IRMAIssueSpecification_LocalizedName        `protobuf:"bytes,2,rep,name=Name" json:"Name,omitempty"`
	Description []*IRMAIssueSpecification_LocalizedDescription `protobuf:"bytes,3,rep,name=Description" json:"Description,omitempty"`
}

func (m *IRMAIssueSpecification_Attribute) Reset()         { *m = IRMAIssueSpecification_Attribute{} }
func (m *IRMAIssueSpecification_Attribute) String() string { return proto.CompactTextString(m) }
func (*IRMAIssueSpecification_Attribute) ProtoMessage()    {}
func (*IRMAIssueSpecification_Attribute) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 2}
}

func (m *IRMAIssueSpecification_Attribute) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *IRMAIssueSpecification_Attribute) GetName() []*IRMAIssueSpecification_LocalizedName {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *IRMAIssueSpecification_Attribute) GetDescription() []*IRMAIssueSpecification_LocalizedDescription {
	if m != nil {
		return m.Description
	}
	return nil
}

func init() {
	proto.RegisterType((*IRMAIssueSpecification)(nil), "irmaproto.IRMAIssueSpecification")
	proto.RegisterType((*IRMAIssueSpecification_LocalizedName)(nil), "irmaproto.IRMAIssueSpecification.LocalizedName")
	proto.RegisterType((*IRMAIssueSpecification_LocalizedDescription)(nil), "irmaproto.IRMAIssueSpecification.LocalizedDescription")
	proto.RegisterType((*IRMAIssueSpecification_Attribute)(nil), "irmaproto.IRMAIssueSpecification.Attribute")
}

func init() { proto.RegisterFile("IRMAIssueSpecification.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x14, 0x85, 0x49, 0xbf, 0x73, 0xfb, 0xf6, 0x05, 0x2f, 0x22, 0x43, 0x71, 0x11, 0x8a, 0x8b, 0x80,
	0x52, 0x41, 0x41, 0x77, 0x42, 0xad, 0x9b, 0x60, 0xeb, 0x62, 0x02, 0xe2, 0x76, 0x9a, 0x5c, 0xdb,
	0x81, 0x74, 0xa6, 0x4c, 0xa6, 0x1b, 0xff, 0x98, 0xe0, 0xaf, 0x93, 0x0c, 0x35, 0xb6, 0x5a, 0xf1,
	0x83, 0xee, 0xee, 0x9c, 0xb9, 0xcf, 0xc9, 0x49, 0x4e, 0xe0, 0x30, 0xe2, 0xe3, 0x41, 0x94, 0xe7,
	0x4b, 0x8a, 0x17, 0x94, 0xc8, 0x47, 0x99, 0x08, 0x2b, 0xb5, 0xea, 0x2f, 0x8c, 0xb6, 0x1a, 0x7d,
	0x69, 0xe6, 0xc2, 0x8d, 0xbd, 0xe7, 0x06, 0x1c, 0x6c, 0xdf, 0x45, 0x06, 0xcd, 0x7b, 0x32, 0xb9,
	0xd4, 0x8a, 0x79, 0x81, 0x17, 0xd6, 0xf9, 0xdb, 0x11, 0x87, 0x50, 0xbb, 0x13, 0x73, 0x62, 0x95,
	0xa0, 0x1a, 0xb6, 0xcf, 0x4e, 0xfb, 0xa5, 0x5d, 0xff, 0x8b, 0xc7, 0x8e, 0x74, 0x22, 0x32, 0xf9,
	0x44, 0x69, 0x81, 0x71, 0x07, 0xe3, 0x18, 0xfc, 0x78, 0xa6, 0x8d, 0x75, 0x4e, 0xd5, 0xbf, 0x39,
	0xbd, 0x3b, 0xe0, 0x11, 0x74, 0xe2, 0x64, 0x46, 0x73, 0x1a, 0x0b, 0x25, 0xa6, 0x64, 0x58, 0x2d,
	0xf0, 0x42, 0x9f, 0x6f, 0x8a, 0xd8, 0x85, 0x96, 0x33, 0x35, 0x51, 0xca, 0xea, 0x6e, 0xa1, 0x3c,
	0x63, 0x0f, 0xfe, 0x0d, 0x0d, 0xa5, 0xa4, 0xac, 0x14, 0x59, 0x94, 0xb2, 0x86, 0xbb, 0xdf, 0xd0,
	0xf0, 0x01, 0xda, 0x37, 0x94, 0x27, 0x46, 0x2e, 0x8a, 0x34, 0xac, 0xe9, 0x62, 0x5f, 0xfc, 0x22,
	0xf6, 0x1a, 0xcd, 0xd7, 0xad, 0xf0, 0x04, 0xf6, 0xe2, 0x99, 0x5e, 0x66, 0xe9, 0x35, 0xc5, 0x52,
	0x4d, 0x33, 0xb2, 0x5a, 0xb1, 0x56, 0xe0, 0x85, 0x2d, 0xfe, 0xf9, 0x02, 0x6f, 0x01, 0x06, 0xd6,
	0x1a, 0x39, 0x59, 0x5a, 0xca, 0x99, 0xef, 0x62, 0x1c, 0x7f, 0x1f, 0xa3, 0x64, 0xf8, 0x1a, 0xde,
	0xbd, 0x84, 0xce, 0xc6, 0x67, 0x45, 0x84, 0xda, 0x48, 0xa8, 0xa9, 0xab, 0xdd, 0xe7, 0x6e, 0x2e,
	0xb4, 0x55, 0xe7, 0x4e, 0x2b, 0xe6, 0xee, 0x15, 0xec, 0x6f, 0x7b, 0xb1, 0x1f, 0xf3, 0x2f, 0x1e,
	0xf8, 0x65, 0x0e, 0xfc, 0x0f, 0x95, 0x28, 0x5d, 0x31, 0x95, 0x28, 0xdd, 0xcd, 0x5f, 0xf6, 0xa1,
	0xb0, 0xea, 0xce, 0x0a, 0x9b, 0x34, 0x1c, 0x7f, 0xfe, 0x1a, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xa7,
	0x55, 0x6c, 0x6b, 0x03, 0x00, 0x00,
}