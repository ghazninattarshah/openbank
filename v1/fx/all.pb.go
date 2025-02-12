// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/fx/all.proto

package fx

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// FX holds all the details about a foreign exchange object.
type FX struct {
	// BankID is an identifier for the bank on this fx transaction.
	BankID string `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"BankID,omitempty"`
	// FromCurrencyCode is the currency to transfer from.
	FromCurrencyCode string `protobuf:"bytes,2,opt,name=FromCurrencyCode,json=from_currency_code,proto3" json:"FromCurrencyCode,omitempty"`
	// ToCurrencyCode is the currency that we are transferring to.
	ToCurrencyCode string `protobuf:"bytes,3,opt,name=ToCurrencyCode,json=to_currency_code,proto3" json:"ToCurrencyCode,omitempty"`
	// Rate is the exchange rate of the foreign exchange.
	Rate string `protobuf:"bytes,4,opt,name=Rate,json=rate,proto3" json:"Rate,omitempty"`
	// InverseRate is the inverse of the exchange rate of the foreign exchange.
	InverseRate string `protobuf:"bytes,5,opt,name=InverseRate,json=inverse_rate,proto3" json:"InverseRate,omitempty"`
	// EffectiveDate is the effective date of the foreign exchange quote.
	EffectiveDate        string   `protobuf:"bytes,6,opt,name=EffectiveDate,json=effective_date,proto3" json:"EffectiveDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FX) Reset()         { *m = FX{} }
func (m *FX) String() string { return proto.CompactTextString(m) }
func (*FX) ProtoMessage()    {}
func (*FX) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3bdb6bf3dbfaace, []int{0}
}

func (m *FX) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FX.Unmarshal(m, b)
}
func (m *FX) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FX.Marshal(b, m, deterministic)
}
func (m *FX) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FX.Merge(m, src)
}
func (m *FX) XXX_Size() int {
	return xxx_messageInfo_FX.Size(m)
}
func (m *FX) XXX_DiscardUnknown() {
	xxx_messageInfo_FX.DiscardUnknown(m)
}

var xxx_messageInfo_FX proto.InternalMessageInfo

func (m *FX) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *FX) GetFromCurrencyCode() string {
	if m != nil {
		return m.FromCurrencyCode
	}
	return ""
}

func (m *FX) GetToCurrencyCode() string {
	if m != nil {
		return m.ToCurrencyCode
	}
	return ""
}

func (m *FX) GetRate() string {
	if m != nil {
		return m.Rate
	}
	return ""
}

func (m *FX) GetInverseRate() string {
	if m != nil {
		return m.InverseRate
	}
	return ""
}

func (m *FX) GetEffectiveDate() string {
	if m != nil {
		return m.EffectiveDate
	}
	return ""
}

// CreateFXRequest is a request envelope to create a foreign exchange quote.
type CreateFXRequest struct {
	// FX is the foreign exchange information to create.
	FX                   *FX      `protobuf:"bytes,1,opt,name=FX,json=fx,proto3" json:"FX,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFXRequest) Reset()         { *m = CreateFXRequest{} }
func (m *CreateFXRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFXRequest) ProtoMessage()    {}
func (*CreateFXRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3bdb6bf3dbfaace, []int{1}
}

func (m *CreateFXRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFXRequest.Unmarshal(m, b)
}
func (m *CreateFXRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFXRequest.Marshal(b, m, deterministic)
}
func (m *CreateFXRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFXRequest.Merge(m, src)
}
func (m *CreateFXRequest) XXX_Size() int {
	return xxx_messageInfo_CreateFXRequest.Size(m)
}
func (m *CreateFXRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFXRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFXRequest proto.InternalMessageInfo

func (m *CreateFXRequest) GetFX() *FX {
	if m != nil {
		return m.FX
	}
	return nil
}

// UpdateFXRequest is a request envelope to update a foreign exchange quote.
type UpdateFXRequest struct {
	// FX is the foreign exchange information to update.
	FX                   *FX      `protobuf:"bytes,1,opt,name=FX,json=fx,proto3" json:"FX,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateFXRequest) Reset()         { *m = UpdateFXRequest{} }
func (m *UpdateFXRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateFXRequest) ProtoMessage()    {}
func (*UpdateFXRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3bdb6bf3dbfaace, []int{2}
}

func (m *UpdateFXRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFXRequest.Unmarshal(m, b)
}
func (m *UpdateFXRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFXRequest.Marshal(b, m, deterministic)
}
func (m *UpdateFXRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFXRequest.Merge(m, src)
}
func (m *UpdateFXRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateFXRequest.Size(m)
}
func (m *UpdateFXRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFXRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFXRequest proto.InternalMessageInfo

func (m *UpdateFXRequest) GetFX() *FX {
	if m != nil {
		return m.FX
	}
	return nil
}

// GetCurrentFXRateRequest is a request envelope to get a foreign exchange rate.
type GetCurrentFXRateRequest struct {
	// FromCurrencyCode is the currency to transfer from.
	FromCurrencyCode string `protobuf:"bytes,1,opt,name=FromCurrencyCode,json=from_currency_code,proto3" json:"FromCurrencyCode,omitempty"`
	// ToCurrencyCode is the currency that we are transferring to.
	ToCurrencyCode       string   `protobuf:"bytes,2,opt,name=ToCurrencyCode,json=to_currency_code,proto3" json:"ToCurrencyCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentFXRateRequest) Reset()         { *m = GetCurrentFXRateRequest{} }
func (m *GetCurrentFXRateRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurrentFXRateRequest) ProtoMessage()    {}
func (*GetCurrentFXRateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3bdb6bf3dbfaace, []int{3}
}

func (m *GetCurrentFXRateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentFXRateRequest.Unmarshal(m, b)
}
func (m *GetCurrentFXRateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentFXRateRequest.Marshal(b, m, deterministic)
}
func (m *GetCurrentFXRateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentFXRateRequest.Merge(m, src)
}
func (m *GetCurrentFXRateRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurrentFXRateRequest.Size(m)
}
func (m *GetCurrentFXRateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentFXRateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentFXRateRequest proto.InternalMessageInfo

func (m *GetCurrentFXRateRequest) GetFromCurrencyCode() string {
	if m != nil {
		return m.FromCurrencyCode
	}
	return ""
}

func (m *GetCurrentFXRateRequest) GetToCurrencyCode() string {
	if m != nil {
		return m.ToCurrencyCode
	}
	return ""
}

func init() {
	proto.RegisterType((*FX)(nil), "fx.FX")
	proto.RegisterType((*CreateFXRequest)(nil), "fx.CreateFXRequest")
	proto.RegisterType((*UpdateFXRequest)(nil), "fx.UpdateFXRequest")
	proto.RegisterType((*GetCurrentFXRateRequest)(nil), "fx.GetCurrentFXRateRequest")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/fx/all.proto", fileDescriptor_e3bdb6bf3dbfaace)
}

var fileDescriptor_e3bdb6bf3dbfaace = []byte{
	// 1141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xde, 0x5d, 0x27, 0xae, 0x3b, 0x85, 0x62, 0x26, 0xa8, 0x44, 0x6e, 0x55, 0x46, 0xae, 0x04,
	0x51, 0x94, 0xac, 0x7f, 0x34, 0x95, 0xaa, 0x48, 0x55, 0x71, 0xd2, 0x38, 0x24, 0x42, 0x22, 0x32,
	0xa5, 0xb2, 0x7a, 0xb1, 0xc6, 0xbb, 0x6f, 0xed, 0x21, 0xf6, 0xcc, 0x76, 0x66, 0xd6, 0x76, 0x28,
	0x95, 0x50, 0x25, 0xa4, 0x5c, 0x90, 0x4a, 0xf8, 0x13, 0xb8, 0x70, 0xe2, 0x0c, 0xff, 0x05, 0x52,
	0x39, 0x70, 0x42, 0x42, 0xca, 0x0d, 0x4e, 0x9c, 0x38, 0xa2, 0xd9, 0x5d, 0x3b, 0x89, 0x9d, 0xf0,
	0xa3, 0x27, 0xef, 0xfa, 0x7d, 0xdf, 0x37, 0xef, 0x7d, 0xef, 0xbd, 0x59, 0xb4, 0xd2, 0x61, 0xba,
	0x1b, 0xb5, 0x5d, 0x4f, 0xf4, 0x4b, 0x22, 0x04, 0xde, 0xa6, 0x7c, 0xff, 0xe4, 0x61, 0x50, 0x29,
	0x05, 0xa3, 0x12, 0xed, 0xf5, 0xdc, 0x50, 0x0a, 0x2d, 0xb0, 0x13, 0x8c, 0x0a, 0x37, 0x3a, 0x42,
	0x74, 0x7a, 0x50, 0xa2, 0x21, 0x2b, 0x51, 0xce, 0x85, 0xa6, 0x9a, 0x09, 0xae, 0x12, 0x44, 0x61,
	0x25, 0xfe, 0xf1, 0x56, 0x3b, 0xc0, 0x57, 0xd5, 0x90, 0x76, 0x3a, 0x20, 0x4b, 0x22, 0x8c, 0x11,
	0xe7, 0xa0, 0xaf, 0xa7, 0x5a, 0xf1, 0x5b, 0x3b, 0x0a, 0x4a, 0xd0, 0x0f, 0xf5, 0x41, 0x12, 0x2c,
	0x7e, 0xeb, 0x20, 0xa7, 0xde, 0xc4, 0xb7, 0x50, 0x76, 0x83, 0xf2, 0xfd, 0x9d, 0x07, 0x8b, 0x36,
	0xb1, 0x97, 0x2e, 0x6f, 0xa0, 0x9c, 0xb5, 0x68, 0x2d, 0x59, 0x65, 0x6b, 0xcf, 0x6a, 0x5c, 0x32,
	0x59, 0xb6, 0x98, 0x8f, 0xef, 0xa2, 0x7c, 0x5d, 0x8a, 0xfe, 0x66, 0x24, 0x25, 0x70, 0xef, 0x60,
	0x53, 0xf8, 0xb0, 0xe8, 0xcc, 0xc0, 0x71, 0x20, 0x45, 0xbf, 0xe5, 0xa5, 0xa0, 0x96, 0x27, 0x7c,
	0xc0, 0x6b, 0xe8, 0xea, 0x43, 0x71, 0x86, 0x97, 0x99, 0xe1, 0xe5, 0xb5, 0x98, 0x62, 0xdd, 0x44,
	0x73, 0x0d, 0xaa, 0x61, 0x71, 0x6e, 0x06, 0x3b, 0x27, 0xa9, 0x06, 0xbc, 0x8a, 0xae, 0xec, 0xf0,
	0x01, 0x48, 0x05, 0x31, 0x6c, 0x7e, 0x06, 0xf6, 0x1a, 0x4b, 0xc2, 0xad, 0x18, 0x5e, 0x41, 0xaf,
	0x6f, 0x05, 0x01, 0x78, 0x9a, 0x0d, 0xe0, 0x81, 0x21, 0x64, 0x67, 0x08, 0x57, 0x61, 0x0c, 0x68,
	0xf9, 0x54, 0xc3, 0x7a, 0x36, 0x67, 0xe5, 0xad, 0x45, 0xab, 0x78, 0x0f, 0xbd, 0xb1, 0x29, 0x81,
	0x6a, 0xa8, 0x37, 0x1b, 0xf0, 0x24, 0x02, 0xa5, 0x71, 0xd1, 0xf8, 0x16, 0xbb, 0x75, 0xa5, 0x9a,
	0x75, 0x83, 0x91, 0x5b, 0x6f, 0x9e, 0x91, 0x72, 0x82, 0xd1, 0x69, 0xfa, 0x27, 0xa1, 0xff, 0xca,
	0xf4, 0xaf, 0x6d, 0xf4, 0xf6, 0x36, 0xe8, 0xc4, 0x3f, 0x5d, 0x6f, 0x9a, 0x6a, 0xc7, 0x3a, 0xe7,
	0xf5, 0xc4, 0x7e, 0xc5, 0x9e, 0x38, 0xff, 0xde, 0x93, 0x71, 0x4e, 0xd5, 0xdf, 0xb3, 0xe8, 0x72,
	0xbd, 0xf9, 0x31, 0xc8, 0x01, 0xf3, 0x00, 0xff, 0xe1, 0xa0, 0xfc, 0x74, 0x86, 0xf8, 0xba, 0x29,
	0xeb, 0x82, 0xbc, 0x0b, 0x69, 0xcd, 0xc5, 0xef, 0x9d, 0xa3, 0xda, 0x9f, 0x76, 0x3c, 0x83, 0xab,
	0x0d, 0xd0, 0x92, 0xc1, 0x00, 0x08, 0x8c, 0xbc, 0x2e, 0xe5, 0x1d, 0x20, 0xa6, 0x6f, 0xa4, 0x0d,
	0x7a, 0x08, 0xc0, 0x89, 0x1e, 0x0a, 0x92, 0x26, 0xc3, 0x40, 0x15, 0xee, 0x4c, 0xe0, 0xba, 0x3b,
	0x4d, 0x31, 0xe5, 0x12, 0x3a, 0x46, 0x1f, 0x10, 0x2d, 0x08, 0xe5, 0x42, 0x77, 0x41, 0xee, 0xde,
	0x47, 0x99, 0x6a, 0xb9, 0x8c, 0xef, 0xa2, 0x9b, 0x69, 0x2a, 0x04, 0x46, 0xe0, 0x45, 0x1a, 0x7c,
	0xa2, 0x22, 0xcf, 0x03, 0xa5, 0x82, 0xa8, 0xd7, 0x3b, 0x70, 0xf1, 0x35, 0xf4, 0x56, 0x01, 0xdf,
	0x2a, 0xf9, 0x10, 0x30, 0xce, 0x92, 0xdd, 0x0a, 0x46, 0xf5, 0xe6, 0x6e, 0x05, 0x65, 0xd6, 0xca,
	0x6b, 0x78, 0x19, 0x2d, 0x35, 0x40, 0x47, 0x92, 0x83, 0x4f, 0x86, 0x5d, 0x93, 0x5e, 0x17, 0x88,
	0x04, 0x25, 0x22, 0xe9, 0x01, 0x61, 0x8a, 0x70, 0xa1, 0x49, 0x20, 0x22, 0xee, 0xbb, 0x6d, 0x8c,
	0xf2, 0x28, 0xfb, 0x51, 0x2d, 0xd2, 0xdd, 0x2a, 0xce, 0xa2, 0x39, 0x09, 0xd4, 0x7f, 0xfe, 0xf2,
	0xf8, 0x1b, 0x67, 0x05, 0x2f, 0xa7, 0xdb, 0xff, 0x74, 0xba, 0x89, 0xcf, 0x4a, 0x4f, 0xcf, 0x76,
	0xe7, 0xd9, 0xa1, 0x63, 0xbd, 0x70, 0xe2, 0xc6, 0xe0, 0xe7, 0x0e, 0xca, 0x8d, 0xe7, 0x11, 0x2f,
	0x18, 0x27, 0xa7, 0xa6, 0x73, 0x62, 0xef, 0xcf, 0xf6, 0x51, 0xed, 0xc7, 0xc4, 0xde, 0x77, 0x12,
	0x0c, 0xa1, 0x24, 0x10, 0x12, 0x58, 0x87, 0x9f, 0x78, 0xf6, 0x24, 0x12, 0x1a, 0x0a, 0x6b, 0x09,
	0x40, 0x11, 0x4a, 0x38, 0x0c, 0x2f, 0x40, 0x11, 0xca, 0x7d, 0x22, 0xe3, 0xc2, 0x15, 0x61, 0xda,
	0xdd, 0xdd, 0x36, 0x7e, 0x56, 0xf0, 0xfb, 0xe8, 0xdd, 0x7a, 0x4a, 0xd8, 0x1a, 0x13, 0xbc, 0x58,
	0xef, 0x3f, 0xfa, 0xda, 0x5e, 0x40, 0x6f, 0x4e, 0x4c, 0xba, 0x84, 0xe6, 0x87, 0x92, 0x69, 0x88,
	0x5d, 0xba, 0xb2, 0x6e, 0x2f, 0x17, 0xb3, 0x89, 0x51, 0xa7, 0x4c, 0x38, 0xb6, 0x51, 0x6e, 0xbc,
	0x55, 0x89, 0x09, 0x53, 0x3b, 0x56, 0xb8, 0xe6, 0x26, 0x37, 0x9f, 0x3b, 0xbe, 0xf9, 0xdc, 0x2d,
	0x73, 0xf3, 0x15, 0xbf, 0xb3, 0x8f, 0x6a, 0x5f, 0xa5, 0xa6, 0x24, 0x9c, 0x8b, 0x4d, 0x21, 0x09,
	0x40, 0x5d, 0x88, 0xd8, 0x2d, 0x19, 0x03, 0xd6, 0xf0, 0xd2, 0x39, 0x06, 0x9c, 0x2e, 0x9c, 0x44,
	0xb1, 0x90, 0xef, 0xfe, 0x73, 0xa1, 0x85, 0xa9, 0x42, 0x0b, 0x99, 0x43, 0xc7, 0xda, 0xf8, 0x32,
	0x7b, 0x54, 0x3b, 0x9e, 0x47, 0x99, 0xaa, 0x5b, 0xc6, 0x8f, 0x50, 0xb6, 0xde, 0x24, 0xb5, 0xbd,
	0x1d, 0xbc, 0xb5, 0x27, 0xc5, 0x80, 0xf9, 0xa0, 0x52, 0xa7, 0xd3, 0xde, 0x50, 0x9f, 0x88, 0x10,
	0x64, 0xf2, 0x0d, 0x20, 0x22, 0x99, 0xcc, 0x99, 0xe4, 0xc7, 0xa3, 0xea, 0x56, 0xe7, 0x2b, 0x6e,
	0xd9, 0x2d, 0x2f, 0xdb, 0x4e, 0x35, 0x4f, 0xc3, 0xb0, 0xc7, 0xbc, 0x98, 0x59, 0xfa, 0x54, 0x09,
	0xbe, 0x3e, 0xf3, 0x4f, 0xa3, 0x65, 0xc6, 0xbf, 0x8c, 0x9b, 0xe8, 0xd1, 0x79, 0xe3, 0x9f, 0x6c,
	0x54, 0x5b, 0xf8, 0x07, 0x66, 0x05, 0xfa, 0xb4, 0x17, 0x08, 0xd9, 0xa7, 0xda, 0x4c, 0x81, 0x90,
	0xc4, 0x17, 0x90, 0xec, 0x45, 0x9f, 0x6a, 0xaf, 0x9b, 0xee, 0x6d, 0x08, 0x9e, 0x09, 0xa7, 0x5c,
	0xb7, 0xf1, 0xa1, 0x39, 0xa0, 0x82, 0xb7, 0xd0, 0xe6, 0xc5, 0x07, 0x4c, 0x84, 0x3c, 0xc1, 0x35,
	0x65, 0x5c, 0xc5, 0xd1, 0x48, 0x81, 0x7c, 0x2f, 0x36, 0xc3, 0x07, 0xae, 0x19, 0xed, 0x29, 0xb7,
	0xb1, 0x67, 0xd4, 0x6e, 0xe3, 0x1d, 0xb4, 0x3d, 0xab, 0x66, 0xf0, 0x27, 0x52, 0x5d, 0x3a, 0x00,
	0x12, 0x82, 0xec, 0x33, 0xa5, 0x98, 0x71, 0x4d, 0x10, 0x1a, 0xb7, 0xef, 0xcc, 0x66, 0xbb, 0x8d,
	0xff, 0xbf, 0xff, 0x8d, 0x3a, 0xca, 0xdc, 0x29, 0x97, 0xf1, 0x7d, 0x74, 0xef, 0x2c, 0x85, 0x72,
	0x12, 0xf1, 0x89, 0x03, 0x20, 0xa5, 0x90, 0x44, 0x78, 0x5e, 0x24, 0x8d, 0x5d, 0x89, 0xa2, 0x02,
	0x39, 0x00, 0x49, 0x14, 0xf3, 0xc1, 0x7d, 0xfc, 0x9b, 0x8d, 0x7e, 0xb5, 0x27, 0xb3, 0xf3, 0xd2,
	0xce, 0x65, 0xf0, 0xe7, 0xb5, 0x34, 0x47, 0x41, 0x82, 0xd1, 0xe4, 0x7c, 0x65, 0x12, 0x90, 0xa0,
	0xb4, 0x64, 0xb1, 0xb4, 0xa9, 0x25, 0xd2, 0x5d, 0xe3, 0x8a, 0x17, 0xaf, 0xa4, 0x29, 0x5d, 0xb9,
	0xe4, 0x61, 0x17, 0x4e, 0x07, 0x4c, 0xd9, 0xa1, 0x14, 0xb1, 0x60, 0x20, 0x7a, 0x3d, 0x31, 0x4c,
	0x8a, 0x37, 0xe7, 0x09, 0xc9, 0x3e, 0x4b, 0x10, 0xe6, 0x36, 0x22, 0x41, 0x4f, 0x0c, 0xdd, 0xa5,
	0xb9, 0x6a, 0xce, 0x8c, 0xab, 0x91, 0x58, 0xbf, 0x6c, 0x9e, 0xb4, 0xd8, 0x07, 0xbe, 0xb1, 0x8e,
	0x0a, 0xc9, 0x2d, 0x87, 0xf1, 0xb6, 0xa4, 0x5c, 0xab, 0x64, 0x2a, 0x13, 0x3b, 0xd1, 0x8d, 0x74,
	0xe6, 0xf1, 0x42, 0x1a, 0x8c, 0xdf, 0xd2, 0xe8, 0x07, 0xf6, 0x9e, 0xf5, 0xd8, 0x09, 0x46, 0x5f,
	0xd8, 0xd6, 0xa1, 0x6d, 0xbd, 0xb0, 0xad, 0x1f, 0x6c, 0xeb, 0x17, 0xdb, 0xfa, 0xcb, 0xb6, 0x7e,
	0x72, 0xac, 0x76, 0x36, 0xde, 0xe4, 0xdb, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x3b, 0x34,
	0x06, 0x51, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FXServiceClient is the client API for FXService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FXServiceClient interface {
	// GetCurrentFXRate retrieves the current foreign exchange rate of the two
	//specified currency code.
	GetCurrentFXRate(ctx context.Context, in *GetCurrentFXRateRequest, opts ...grpc.CallOption) (*FX, error)
	// CreateFX creates a new foreign exchange quote.
	CreateFX(ctx context.Context, in *CreateFXRequest, opts ...grpc.CallOption) (*FX, error)
	// UpdateFX updates a foreign exchange quote.
	UpdateFX(ctx context.Context, in *UpdateFXRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type fXServiceClient struct {
	cc *grpc.ClientConn
}

func NewFXServiceClient(cc *grpc.ClientConn) FXServiceClient {
	return &fXServiceClient{cc}
}

func (c *fXServiceClient) GetCurrentFXRate(ctx context.Context, in *GetCurrentFXRateRequest, opts ...grpc.CallOption) (*FX, error) {
	out := new(FX)
	err := c.cc.Invoke(ctx, "/fx.FXService/GetCurrentFXRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fXServiceClient) CreateFX(ctx context.Context, in *CreateFXRequest, opts ...grpc.CallOption) (*FX, error) {
	out := new(FX)
	err := c.cc.Invoke(ctx, "/fx.FXService/CreateFX", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fXServiceClient) UpdateFX(ctx context.Context, in *UpdateFXRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/fx.FXService/UpdateFX", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FXServiceServer is the server API for FXService service.
type FXServiceServer interface {
	// GetCurrentFXRate retrieves the current foreign exchange rate of the two
	//specified currency code.
	GetCurrentFXRate(context.Context, *GetCurrentFXRateRequest) (*FX, error)
	// CreateFX creates a new foreign exchange quote.
	CreateFX(context.Context, *CreateFXRequest) (*FX, error)
	// UpdateFX updates a foreign exchange quote.
	UpdateFX(context.Context, *UpdateFXRequest) (*empty.Empty, error)
}

// UnimplementedFXServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFXServiceServer struct {
}

func (*UnimplementedFXServiceServer) GetCurrentFXRate(ctx context.Context, req *GetCurrentFXRateRequest) (*FX, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentFXRate not implemented")
}
func (*UnimplementedFXServiceServer) CreateFX(ctx context.Context, req *CreateFXRequest) (*FX, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFX not implemented")
}
func (*UnimplementedFXServiceServer) UpdateFX(ctx context.Context, req *UpdateFXRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFX not implemented")
}

func RegisterFXServiceServer(s *grpc.Server, srv FXServiceServer) {
	s.RegisterService(&_FXService_serviceDesc, srv)
}

func _FXService_GetCurrentFXRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentFXRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FXServiceServer).GetCurrentFXRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fx.FXService/GetCurrentFXRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FXServiceServer).GetCurrentFXRate(ctx, req.(*GetCurrentFXRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FXService_CreateFX_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFXRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FXServiceServer).CreateFX(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fx.FXService/CreateFX",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FXServiceServer).CreateFX(ctx, req.(*CreateFXRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FXService_UpdateFX_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFXRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FXServiceServer).UpdateFX(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fx.FXService/UpdateFX",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FXServiceServer).UpdateFX(ctx, req.(*UpdateFXRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FXService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fx.FXService",
	HandlerType: (*FXServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentFXRate",
			Handler:    _FXService_GetCurrentFXRate_Handler,
		},
		{
			MethodName: "CreateFX",
			Handler:    _FXService_CreateFX_Handler,
		},
		{
			MethodName: "UpdateFX",
			Handler:    _FXService_UpdateFX_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/fx/all.proto",
}
