// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dke/v1/ekms.proto

package dke

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Gets the public key
type PublicKeyRequest struct {
	// Required. Key Name
	// e.g. `TestKey1`.
	KeyName              string   `protobuf:"bytes,1,opt,name=key_name,json=keyName,proto3" json:"key_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicKeyRequest) Reset()         { *m = PublicKeyRequest{} }
func (m *PublicKeyRequest) String() string { return proto.CompactTextString(m) }
func (*PublicKeyRequest) ProtoMessage()    {}
func (*PublicKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e13a1835b8d008a9, []int{0}
}

func (m *PublicKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKeyRequest.Unmarshal(m, b)
}
func (m *PublicKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKeyRequest.Marshal(b, m, deterministic)
}
func (m *PublicKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKeyRequest.Merge(m, src)
}
func (m *PublicKeyRequest) XXX_Size() int {
	return xxx_messageInfo_PublicKeyRequest.Size(m)
}
func (m *PublicKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKeyRequest proto.InternalMessageInfo

func (m *PublicKeyRequest) GetKeyName() string {
	if m != nil {
		return m.KeyName
	}
	return ""
}

// Response message for [AzureDKEKeyManagementService.GetPublicKey][].
type PublicKeyResponse struct {
	// The key
	Key                  *PublicKeyResponse_Key `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *PublicKeyResponse) Reset()         { *m = PublicKeyResponse{} }
func (m *PublicKeyResponse) String() string { return proto.CompactTextString(m) }
func (*PublicKeyResponse) ProtoMessage()    {}
func (*PublicKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e13a1835b8d008a9, []int{1}
}

func (m *PublicKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKeyResponse.Unmarshal(m, b)
}
func (m *PublicKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKeyResponse.Marshal(b, m, deterministic)
}
func (m *PublicKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKeyResponse.Merge(m, src)
}
func (m *PublicKeyResponse) XXX_Size() int {
	return xxx_messageInfo_PublicKeyResponse.Size(m)
}
func (m *PublicKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKeyResponse proto.InternalMessageInfo

func (m *PublicKeyResponse) GetKey() *PublicKeyResponse_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

// Key information
type PublicKeyResponse_Key struct {
	// Key type
	Kty string `protobuf:"bytes,1,opt,name=kty,proto3" json:"kty,omitempty"`
	// Key
	N string `protobuf:"bytes,2,opt,name=n,proto3" json:"n,omitempty"`
	// Key exponent (typically 65537)
	E int32 `protobuf:"varint,3,opt,name=e,proto3" json:"e,omitempty"`
	// Algorithm
	Alg string `protobuf:"bytes,4,opt,name=alg,proto3" json:"alg,omitempty"`
	// Key ID
	Kid                  string   `protobuf:"bytes,5,opt,name=kid,proto3" json:"kid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicKeyResponse_Key) Reset()         { *m = PublicKeyResponse_Key{} }
func (m *PublicKeyResponse_Key) String() string { return proto.CompactTextString(m) }
func (*PublicKeyResponse_Key) ProtoMessage()    {}
func (*PublicKeyResponse_Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_e13a1835b8d008a9, []int{1, 0}
}

func (m *PublicKeyResponse_Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKeyResponse_Key.Unmarshal(m, b)
}
func (m *PublicKeyResponse_Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKeyResponse_Key.Marshal(b, m, deterministic)
}
func (m *PublicKeyResponse_Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKeyResponse_Key.Merge(m, src)
}
func (m *PublicKeyResponse_Key) XXX_Size() int {
	return xxx_messageInfo_PublicKeyResponse_Key.Size(m)
}
func (m *PublicKeyResponse_Key) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKeyResponse_Key.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKeyResponse_Key proto.InternalMessageInfo

func (m *PublicKeyResponse_Key) GetKty() string {
	if m != nil {
		return m.Kty
	}
	return ""
}

func (m *PublicKeyResponse_Key) GetN() string {
	if m != nil {
		return m.N
	}
	return ""
}

func (m *PublicKeyResponse_Key) GetE() int32 {
	if m != nil {
		return m.E
	}
	return 0
}

func (m *PublicKeyResponse_Key) GetAlg() string {
	if m != nil {
		return m.Alg
	}
	return ""
}

func (m *PublicKeyResponse_Key) GetKid() string {
	if m != nil {
		return m.Kid
	}
	return ""
}

// Decrypt using the private key. The encrypted text must have been encrypted
// using the private key.
type DecryptRequest struct {
	// Required. Name of the key
	// e.g. `TestKey1`.
	KeyName string `protobuf:"bytes,1,opt,name=key_name,json=keyName,proto3" json:"key_name,omitempty"`
	// Required. The identifier of the key
	// (as provided by [AzureDKEKeyManagementService.GetKey][])
	// e.g.
	// `D798B899-3350-4F5C-A608-2EDA37CB0EBD`
	KeyId string `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// Required. Algorithm
	// e.g. `RSA-OAEP-256`
	Alg string `protobuf:"bytes,3,opt,name=alg,proto3" json:"alg,omitempty"`
	// Required. Encrypted value, base64 encoded
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecryptRequest) Reset()         { *m = DecryptRequest{} }
func (m *DecryptRequest) String() string { return proto.CompactTextString(m) }
func (*DecryptRequest) ProtoMessage()    {}
func (*DecryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e13a1835b8d008a9, []int{2}
}

func (m *DecryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptRequest.Unmarshal(m, b)
}
func (m *DecryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptRequest.Marshal(b, m, deterministic)
}
func (m *DecryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptRequest.Merge(m, src)
}
func (m *DecryptRequest) XXX_Size() int {
	return xxx_messageInfo_DecryptRequest.Size(m)
}
func (m *DecryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptRequest proto.InternalMessageInfo

func (m *DecryptRequest) GetKeyName() string {
	if m != nil {
		return m.KeyName
	}
	return ""
}

func (m *DecryptRequest) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *DecryptRequest) GetAlg() string {
	if m != nil {
		return m.Alg
	}
	return ""
}

func (m *DecryptRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Response message for [AzureDKEKeyManagementService.Decrypt][].
type DecryptResponse struct {
	// The decrypted data, base64 encoded
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecryptResponse) Reset()         { *m = DecryptResponse{} }
func (m *DecryptResponse) String() string { return proto.CompactTextString(m) }
func (*DecryptResponse) ProtoMessage()    {}
func (*DecryptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e13a1835b8d008a9, []int{3}
}

func (m *DecryptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptResponse.Unmarshal(m, b)
}
func (m *DecryptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptResponse.Marshal(b, m, deterministic)
}
func (m *DecryptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptResponse.Merge(m, src)
}
func (m *DecryptResponse) XXX_Size() int {
	return xxx_messageInfo_DecryptResponse.Size(m)
}
func (m *DecryptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptResponse proto.InternalMessageInfo

func (m *DecryptResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*PublicKeyRequest)(nil), "azure.dke.v1.PublicKeyRequest")
	proto.RegisterType((*PublicKeyResponse)(nil), "azure.dke.v1.PublicKeyResponse")
	proto.RegisterType((*PublicKeyResponse_Key)(nil), "azure.dke.v1.PublicKeyResponse.Key")
	proto.RegisterType((*DecryptRequest)(nil), "azure.dke.v1.DecryptRequest")
	proto.RegisterType((*DecryptResponse)(nil), "azure.dke.v1.DecryptResponse")
}

func init() { proto.RegisterFile("dke/v1/ekms.proto", fileDescriptor_e13a1835b8d008a9) }

var fileDescriptor_e13a1835b8d008a9 = []byte{
	// 669 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4f, 0x6f, 0xd3, 0x3c,
	0x18, 0x7f, 0xdd, 0xbc, 0xdd, 0xf6, 0x7a, 0xdd, 0x9f, 0xfa, 0xdd, 0x44, 0x89, 0x0a, 0x0b, 0xe5,
	0xc0, 0x34, 0xad, 0xc9, 0xda, 0xc1, 0x65, 0x9c, 0x3a, 0x36, 0xc1, 0x5a, 0x98, 0xaa, 0x01, 0x42,
	0x1a, 0x07, 0xe4, 0xd6, 0x8f, 0x52, 0x93, 0xc4, 0x0e, 0x89, 0xdb, 0x29, 0x9b, 0x76, 0x41, 0x9a,
	0x84, 0xc4, 0x6d, 0x9c, 0x38, 0x73, 0xe2, 0xcb, 0x70, 0xe1, 0x2b, 0xf0, 0x1d, 0xb8, 0x22, 0xa7,
	0x69, 0xe9, 0x34, 0x8d, 0x89, 0x53, 0x6c, 0x3f, 0xbf, 0x27, 0xbf, 0x3f, 0x96, 0x1f, 0x5c, 0x64,
	0x1e, 0x38, 0x83, 0x9a, 0x03, 0x5e, 0x10, 0xdb, 0x61, 0x24, 0x95, 0x24, 0x05, 0x7a, 0xdc, 0x8f,
	0xc0, 0x66, 0x1e, 0xd8, 0x83, 0x9a, 0x59, 0x76, 0xa5, 0x74, 0x7d, 0x70, 0x68, 0xc8, 0x1d, 0x2a,
	0x84, 0x54, 0x54, 0x71, 0x29, 0x32, 0xac, 0xb9, 0x3c, 0x51, 0xed, 0x29, 0x15, 0x66, 0xc7, 0xeb,
	0xe9, 0xa7, 0x5b, 0x75, 0x41, 0x54, 0xe3, 0x23, 0xea, 0xba, 0x10, 0x39, 0x32, 0x4c, 0x1b, 0x2f,
	0xff, 0xa4, 0x52, 0xc5, 0x8b, 0xed, 0x7e, 0xc7, 0xe7, 0xdd, 0x16, 0x24, 0x07, 0xf0, 0xae, 0x0f,
	0xb1, 0x22, 0x37, 0xf1, 0x8c, 0x07, 0xc9, 0x1b, 0x41, 0x03, 0x28, 0x21, 0x0b, 0xad, 0xfe, 0x77,
	0x30, 0xed, 0x41, 0xb2, 0x4f, 0x03, 0xa8, 0x7c, 0x41, 0xb8, 0x38, 0x81, 0x8f, 0x43, 0x29, 0x62,
	0x20, 0x0f, 0xb0, 0xe1, 0x41, 0x92, 0x62, 0x67, 0xeb, 0x77, 0xed, 0x49, 0x0f, 0xf6, 0x25, 0xb4,
	0xad, 0xd7, 0x1a, 0x6f, 0xbe, 0xc2, 0x46, 0x0b, 0x12, 0xb2, 0x88, 0x0d, 0x4f, 0x25, 0x19, 0x93,
	0x5e, 0x92, 0x02, 0x46, 0xa2, 0x94, 0x4b, 0xf7, 0x48, 0xe8, 0x1d, 0x94, 0x0c, 0x0b, 0xad, 0xe6,
	0x0f, 0x10, 0x68, 0x34, 0xf5, 0xdd, 0xd2, 0xbf, 0x43, 0x34, 0xf5, 0xdd, 0xb4, 0x9f, 0xb3, 0x52,
	0x3e, 0xeb, 0xe7, 0xac, 0xf2, 0x16, 0xcf, 0xef, 0x40, 0x37, 0x4a, 0x42, 0x75, 0xbd, 0x25, 0xb2,
	0x8c, 0xa7, 0x74, 0x89, 0xb3, 0x8c, 0x31, 0xef, 0x41, 0xb2, 0xc7, 0x46, 0x3c, 0xc6, 0x6f, 0x9e,
	0x25, 0x9c, 0x1f, 0x50, 0xbf, 0x0f, 0x19, 0xf7, 0x70, 0x53, 0xb9, 0x87, 0x17, 0xc6, 0x5c, 0x59,
	0x1c, 0x63, 0xa0, 0x66, 0x2a, 0x64, 0xc0, 0xfa, 0x47, 0x03, 0x97, 0x1b, 0x3a, 0x99, 0x9d, 0xd6,
	0x6e, 0x0b, 0x92, 0x67, 0x54, 0x50, 0x17, 0x02, 0x10, 0xea, 0x39, 0x44, 0x03, 0xde, 0x05, 0xf2,
	0x19, 0xe1, 0xc2, 0x63, 0x50, 0xe3, 0xc0, 0xc8, 0xed, 0x2b, 0x93, 0x4c, 0x4d, 0x99, 0x2b, 0xd7,
	0x24, 0x5d, 0x79, 0x72, 0xde, 0xa8, 0x37, 0x97, 0xb0, 0x71, 0x7f, 0x63, 0x83, 0xcc, 0xe1, 0xd9,
	0x3d, 0x31, 0xa0, 0x3e, 0x67, 0x96, 0x07, 0x49, 0xd3, 0xd4, 0xa7, 0x9b, 0xe4, 0x7f, 0x5c, 0x6c,
	0x43, 0x14, 0xf0, 0x38, 0xe6, 0x52, 0x58, 0x0c, 0x04, 0x07, 0xf6, 0xfe, 0xfb, 0x8f, 0x4f, 0xb9,
	0x39, 0x32, 0xeb, 0x9c, 0x8c, 0x22, 0x3b, 0x25, 0xdf, 0x10, 0x9e, 0xce, 0x6c, 0x92, 0xf2, 0x45,
	0xda, 0x8b, 0x49, 0x9b, 0xb7, 0xae, 0xa8, 0x66, 0x92, 0xce, 0xd0, 0x79, 0xe3, 0x75, 0xb3, 0x34,
	0xd4, 0x54, 0xc4, 0x0b, 0x23, 0x4d, 0xd1, 0xb0, 0xbd, 0xb9, 0xac, 0x2b, 0x35, 0x32, 0x8f, 0x0b,
	0x2f, 0x05, 0xed, 0xab, 0x9e, 0x8c, 0xf8, 0x31, 0xb0, 0x3f, 0xc9, 0xed, 0x14, 0x30, 0xc6, 0x53,
	0xdb, 0x40, 0x23, 0x88, 0xc8, 0x3f, 0xa9, 0xf8, 0x3b, 0x95, 0xf2, 0x84, 0xf8, 0xe1, 0x92, 0xb3,
	0x53, 0x87, 0x0d, 0xc5, 0x6c, 0xa1, 0xb5, 0xed, 0xaf, 0xb9, 0xf3, 0xc6, 0x4f, 0x44, 0xf6, 0x71,
	0x29, 0xbd, 0x13, 0x6b, 0x7c, 0x23, 0x91, 0xb5, 0xd3, 0xda, 0xb5, 0x1a, 0xed, 0xbd, 0x4a, 0x15,
	0xe3, 0x17, 0x3d, 0xea, 0x43, 0x6c, 0x3d, 0x6a, 0x3f, 0x25, 0x2b, 0xfa, 0x89, 0xc5, 0x5b, 0x8e,
	0xa3, 0x8d, 0x65, 0x16, 0x43, 0x29, 0x18, 0x04, 0x54, 0x30, 0x9b, 0xcb, 0xba, 0x51, 0xb3, 0x37,
	0xcc, 0x1b, 0x57, 0x54, 0xd7, 0x50, 0xee, 0xf0, 0x03, 0xc2, 0x67, 0x68, 0xac, 0xf4, 0x64, 0xc6,
	0x30, 0xe7, 0x1a, 0x99, 0xc3, 0xf4, 0x49, 0x5a, 0xb9, 0x55, 0xa3, 0x5e, 0x1b, 0xd1, 0xf8, 0xd2,
	0xe5, 0xc2, 0x3e, 0xe2, 0x82, 0xc9, 0xa3, 0xd8, 0x16, 0xa0, 0x9c, 0xae, 0x0c, 0x02, 0x29, 0x1c,
	0xa9, 0x63, 0xa9, 0x3b, 0xe3, 0x70, 0xb6, 0xfe, 0xbe, 0xe5, 0x70, 0xdd, 0xe5, 0xaa, 0xd7, 0xef,
	0xd8, 0x5d, 0x19, 0x38, 0xaa, 0x47, 0xfd, 0xb8, 0x1b, 0xfa, 0x55, 0x2e, 0x9d, 0x20, 0xae, 0x32,
	0x0f, 0xaa, 0x7a, 0xa2, 0x0c, 0x67, 0xd3, 0x43, 0xe6, 0x41, 0x67, 0x2a, 0x1d, 0x15, 0x9b, 0xbf,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x75, 0x01, 0xe6, 0xb0, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AzureDKEKeyManagementServiceClient is the client API for AzureDKEKeyManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AzureDKEKeyManagementServiceClient interface {
	// Gets the public key
	GetPublicKey(ctx context.Context, in *PublicKeyRequest, opts ...grpc.CallOption) (*PublicKeyResponse, error)
	// Decrypts using the private key. Credentials need to be provided via the
	// Authorization header with the Bearer JWT token.
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
}

type azureDKEKeyManagementServiceClient struct {
	cc *grpc.ClientConn
}

func NewAzureDKEKeyManagementServiceClient(cc *grpc.ClientConn) AzureDKEKeyManagementServiceClient {
	return &azureDKEKeyManagementServiceClient{cc}
}

func (c *azureDKEKeyManagementServiceClient) GetPublicKey(ctx context.Context, in *PublicKeyRequest, opts ...grpc.CallOption) (*PublicKeyResponse, error) {
	out := new(PublicKeyResponse)
	err := c.cc.Invoke(ctx, "/azure.dke.v1.AzureDKEKeyManagementService/GetPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *azureDKEKeyManagementServiceClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	out := new(DecryptResponse)
	err := c.cc.Invoke(ctx, "/azure.dke.v1.AzureDKEKeyManagementService/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AzureDKEKeyManagementServiceServer is the server API for AzureDKEKeyManagementService service.
type AzureDKEKeyManagementServiceServer interface {
	// Gets the public key
	GetPublicKey(context.Context, *PublicKeyRequest) (*PublicKeyResponse, error)
	// Decrypts using the private key. Credentials need to be provided via the
	// Authorization header with the Bearer JWT token.
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
}

func RegisterAzureDKEKeyManagementServiceServer(s *grpc.Server, srv AzureDKEKeyManagementServiceServer) {
	s.RegisterService(&_AzureDKEKeyManagementService_serviceDesc, srv)
}

func _AzureDKEKeyManagementService_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AzureDKEKeyManagementServiceServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azure.dke.v1.AzureDKEKeyManagementService/GetPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AzureDKEKeyManagementServiceServer).GetPublicKey(ctx, req.(*PublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AzureDKEKeyManagementService_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AzureDKEKeyManagementServiceServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/azure.dke.v1.AzureDKEKeyManagementService/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AzureDKEKeyManagementServiceServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AzureDKEKeyManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "azure.dke.v1.AzureDKEKeyManagementService",
	HandlerType: (*AzureDKEKeyManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPublicKey",
			Handler:    _AzureDKEKeyManagementService_GetPublicKey_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _AzureDKEKeyManagementService_Decrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dke/v1/ekms.proto",
}