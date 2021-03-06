// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/tls/v3/tls.proto

package envoy_extensions_transport_sockets_tls_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type UpstreamTlsContext struct {
	CommonTlsContext     *CommonTlsContext     `protobuf:"bytes,1,opt,name=common_tls_context,json=commonTlsContext,proto3" json:"common_tls_context,omitempty"`
	Sni                  string                `protobuf:"bytes,2,opt,name=sni,proto3" json:"sni,omitempty"`
	AllowRenegotiation   bool                  `protobuf:"varint,3,opt,name=allow_renegotiation,json=allowRenegotiation,proto3" json:"allow_renegotiation,omitempty"`
	MaxSessionKeys       *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=max_session_keys,json=maxSessionKeys,proto3" json:"max_session_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpstreamTlsContext) Reset()         { *m = UpstreamTlsContext{} }
func (m *UpstreamTlsContext) String() string { return proto.CompactTextString(m) }
func (*UpstreamTlsContext) ProtoMessage()    {}
func (*UpstreamTlsContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_57bc8d868b0a9df1, []int{0}
}

func (m *UpstreamTlsContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamTlsContext.Unmarshal(m, b)
}
func (m *UpstreamTlsContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamTlsContext.Marshal(b, m, deterministic)
}
func (m *UpstreamTlsContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamTlsContext.Merge(m, src)
}
func (m *UpstreamTlsContext) XXX_Size() int {
	return xxx_messageInfo_UpstreamTlsContext.Size(m)
}
func (m *UpstreamTlsContext) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamTlsContext.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamTlsContext proto.InternalMessageInfo

func (m *UpstreamTlsContext) GetCommonTlsContext() *CommonTlsContext {
	if m != nil {
		return m.CommonTlsContext
	}
	return nil
}

func (m *UpstreamTlsContext) GetSni() string {
	if m != nil {
		return m.Sni
	}
	return ""
}

func (m *UpstreamTlsContext) GetAllowRenegotiation() bool {
	if m != nil {
		return m.AllowRenegotiation
	}
	return false
}

func (m *UpstreamTlsContext) GetMaxSessionKeys() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxSessionKeys
	}
	return nil
}

type DownstreamTlsContext struct {
	CommonTlsContext         *CommonTlsContext   `protobuf:"bytes,1,opt,name=common_tls_context,json=commonTlsContext,proto3" json:"common_tls_context,omitempty"`
	RequireClientCertificate *wrappers.BoolValue `protobuf:"bytes,2,opt,name=require_client_certificate,json=requireClientCertificate,proto3" json:"require_client_certificate,omitempty"`
	RequireSni               *wrappers.BoolValue `protobuf:"bytes,3,opt,name=require_sni,json=requireSni,proto3" json:"require_sni,omitempty"`
	// Types that are valid to be assigned to SessionTicketKeysType:
	//	*DownstreamTlsContext_SessionTicketKeys
	//	*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig
	//	*DownstreamTlsContext_DisableStatelessSessionResumption
	SessionTicketKeysType isDownstreamTlsContext_SessionTicketKeysType `protobuf_oneof:"session_ticket_keys_type"`
	SessionTimeout        *duration.Duration                           `protobuf:"bytes,6,opt,name=session_timeout,json=sessionTimeout,proto3" json:"session_timeout,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                                     `json:"-"`
	XXX_unrecognized      []byte                                       `json:"-"`
	XXX_sizecache         int32                                        `json:"-"`
}

func (m *DownstreamTlsContext) Reset()         { *m = DownstreamTlsContext{} }
func (m *DownstreamTlsContext) String() string { return proto.CompactTextString(m) }
func (*DownstreamTlsContext) ProtoMessage()    {}
func (*DownstreamTlsContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_57bc8d868b0a9df1, []int{1}
}

func (m *DownstreamTlsContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownstreamTlsContext.Unmarshal(m, b)
}
func (m *DownstreamTlsContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownstreamTlsContext.Marshal(b, m, deterministic)
}
func (m *DownstreamTlsContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownstreamTlsContext.Merge(m, src)
}
func (m *DownstreamTlsContext) XXX_Size() int {
	return xxx_messageInfo_DownstreamTlsContext.Size(m)
}
func (m *DownstreamTlsContext) XXX_DiscardUnknown() {
	xxx_messageInfo_DownstreamTlsContext.DiscardUnknown(m)
}

var xxx_messageInfo_DownstreamTlsContext proto.InternalMessageInfo

func (m *DownstreamTlsContext) GetCommonTlsContext() *CommonTlsContext {
	if m != nil {
		return m.CommonTlsContext
	}
	return nil
}

func (m *DownstreamTlsContext) GetRequireClientCertificate() *wrappers.BoolValue {
	if m != nil {
		return m.RequireClientCertificate
	}
	return nil
}

func (m *DownstreamTlsContext) GetRequireSni() *wrappers.BoolValue {
	if m != nil {
		return m.RequireSni
	}
	return nil
}

type isDownstreamTlsContext_SessionTicketKeysType interface {
	isDownstreamTlsContext_SessionTicketKeysType()
}

type DownstreamTlsContext_SessionTicketKeys struct {
	SessionTicketKeys *TlsSessionTicketKeys `protobuf:"bytes,4,opt,name=session_ticket_keys,json=sessionTicketKeys,proto3,oneof"`
}

type DownstreamTlsContext_SessionTicketKeysSdsSecretConfig struct {
	SessionTicketKeysSdsSecretConfig *SdsSecretConfig `protobuf:"bytes,5,opt,name=session_ticket_keys_sds_secret_config,json=sessionTicketKeysSdsSecretConfig,proto3,oneof"`
}

type DownstreamTlsContext_DisableStatelessSessionResumption struct {
	DisableStatelessSessionResumption bool `protobuf:"varint,7,opt,name=disable_stateless_session_resumption,json=disableStatelessSessionResumption,proto3,oneof"`
}

func (*DownstreamTlsContext_SessionTicketKeys) isDownstreamTlsContext_SessionTicketKeysType() {}

func (*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig) isDownstreamTlsContext_SessionTicketKeysType() {
}

func (*DownstreamTlsContext_DisableStatelessSessionResumption) isDownstreamTlsContext_SessionTicketKeysType() {
}

func (m *DownstreamTlsContext) GetSessionTicketKeysType() isDownstreamTlsContext_SessionTicketKeysType {
	if m != nil {
		return m.SessionTicketKeysType
	}
	return nil
}

func (m *DownstreamTlsContext) GetSessionTicketKeys() *TlsSessionTicketKeys {
	if x, ok := m.GetSessionTicketKeysType().(*DownstreamTlsContext_SessionTicketKeys); ok {
		return x.SessionTicketKeys
	}
	return nil
}

func (m *DownstreamTlsContext) GetSessionTicketKeysSdsSecretConfig() *SdsSecretConfig {
	if x, ok := m.GetSessionTicketKeysType().(*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig); ok {
		return x.SessionTicketKeysSdsSecretConfig
	}
	return nil
}

func (m *DownstreamTlsContext) GetDisableStatelessSessionResumption() bool {
	if x, ok := m.GetSessionTicketKeysType().(*DownstreamTlsContext_DisableStatelessSessionResumption); ok {
		return x.DisableStatelessSessionResumption
	}
	return false
}

func (m *DownstreamTlsContext) GetSessionTimeout() *duration.Duration {
	if m != nil {
		return m.SessionTimeout
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DownstreamTlsContext) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DownstreamTlsContext_SessionTicketKeys)(nil),
		(*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig)(nil),
		(*DownstreamTlsContext_DisableStatelessSessionResumption)(nil),
	}
}

type CommonTlsContext struct {
	TlsParams                         *TlsParameters                        `protobuf:"bytes,1,opt,name=tls_params,json=tlsParams,proto3" json:"tls_params,omitempty"`
	TlsCertificates                   []*TlsCertificate                     `protobuf:"bytes,2,rep,name=tls_certificates,json=tlsCertificates,proto3" json:"tls_certificates,omitempty"`
	TlsCertificateSdsSecretConfigs    []*SdsSecretConfig                    `protobuf:"bytes,6,rep,name=tls_certificate_sds_secret_configs,json=tlsCertificateSdsSecretConfigs,proto3" json:"tls_certificate_sds_secret_configs,omitempty"`
	TlsCertificateCertificateProvider *CommonTlsContext_CertificateProvider `protobuf:"bytes,9,opt,name=tls_certificate_certificate_provider,json=tlsCertificateCertificateProvider,proto3" json:"tls_certificate_certificate_provider,omitempty"`
	// Types that are valid to be assigned to ValidationContextType:
	//	*CommonTlsContext_ValidationContext
	//	*CommonTlsContext_ValidationContextSdsSecretConfig
	//	*CommonTlsContext_CombinedValidationContext
	//	*CommonTlsContext_ValidationContextCertificateProvider
	ValidationContextType isCommonTlsContext_ValidationContextType `protobuf_oneof:"validation_context_type"`
	AlpnProtocols         []string                                 `protobuf:"bytes,4,rep,name=alpn_protocols,json=alpnProtocols,proto3" json:"alpn_protocols,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                                 `json:"-"`
	XXX_unrecognized      []byte                                   `json:"-"`
	XXX_sizecache         int32                                    `json:"-"`
}

func (m *CommonTlsContext) Reset()         { *m = CommonTlsContext{} }
func (m *CommonTlsContext) String() string { return proto.CompactTextString(m) }
func (*CommonTlsContext) ProtoMessage()    {}
func (*CommonTlsContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_57bc8d868b0a9df1, []int{2}
}

func (m *CommonTlsContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonTlsContext.Unmarshal(m, b)
}
func (m *CommonTlsContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonTlsContext.Marshal(b, m, deterministic)
}
func (m *CommonTlsContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonTlsContext.Merge(m, src)
}
func (m *CommonTlsContext) XXX_Size() int {
	return xxx_messageInfo_CommonTlsContext.Size(m)
}
func (m *CommonTlsContext) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonTlsContext.DiscardUnknown(m)
}

var xxx_messageInfo_CommonTlsContext proto.InternalMessageInfo

func (m *CommonTlsContext) GetTlsParams() *TlsParameters {
	if m != nil {
		return m.TlsParams
	}
	return nil
}

func (m *CommonTlsContext) GetTlsCertificates() []*TlsCertificate {
	if m != nil {
		return m.TlsCertificates
	}
	return nil
}

func (m *CommonTlsContext) GetTlsCertificateSdsSecretConfigs() []*SdsSecretConfig {
	if m != nil {
		return m.TlsCertificateSdsSecretConfigs
	}
	return nil
}

func (m *CommonTlsContext) GetTlsCertificateCertificateProvider() *CommonTlsContext_CertificateProvider {
	if m != nil {
		return m.TlsCertificateCertificateProvider
	}
	return nil
}

type isCommonTlsContext_ValidationContextType interface {
	isCommonTlsContext_ValidationContextType()
}

type CommonTlsContext_ValidationContext struct {
	ValidationContext *CertificateValidationContext `protobuf:"bytes,3,opt,name=validation_context,json=validationContext,proto3,oneof"`
}

type CommonTlsContext_ValidationContextSdsSecretConfig struct {
	ValidationContextSdsSecretConfig *SdsSecretConfig `protobuf:"bytes,7,opt,name=validation_context_sds_secret_config,json=validationContextSdsSecretConfig,proto3,oneof"`
}

type CommonTlsContext_CombinedValidationContext struct {
	CombinedValidationContext *CommonTlsContext_CombinedCertificateValidationContext `protobuf:"bytes,8,opt,name=combined_validation_context,json=combinedValidationContext,proto3,oneof"`
}

type CommonTlsContext_ValidationContextCertificateProvider struct {
	ValidationContextCertificateProvider *CommonTlsContext_CertificateProvider `protobuf:"bytes,10,opt,name=validation_context_certificate_provider,json=validationContextCertificateProvider,proto3,oneof"`
}

func (*CommonTlsContext_ValidationContext) isCommonTlsContext_ValidationContextType() {}

func (*CommonTlsContext_ValidationContextSdsSecretConfig) isCommonTlsContext_ValidationContextType() {}

func (*CommonTlsContext_CombinedValidationContext) isCommonTlsContext_ValidationContextType() {}

func (*CommonTlsContext_ValidationContextCertificateProvider) isCommonTlsContext_ValidationContextType() {
}

func (m *CommonTlsContext) GetValidationContextType() isCommonTlsContext_ValidationContextType {
	if m != nil {
		return m.ValidationContextType
	}
	return nil
}

func (m *CommonTlsContext) GetValidationContext() *CertificateValidationContext {
	if x, ok := m.GetValidationContextType().(*CommonTlsContext_ValidationContext); ok {
		return x.ValidationContext
	}
	return nil
}

func (m *CommonTlsContext) GetValidationContextSdsSecretConfig() *SdsSecretConfig {
	if x, ok := m.GetValidationContextType().(*CommonTlsContext_ValidationContextSdsSecretConfig); ok {
		return x.ValidationContextSdsSecretConfig
	}
	return nil
}

func (m *CommonTlsContext) GetCombinedValidationContext() *CommonTlsContext_CombinedCertificateValidationContext {
	if x, ok := m.GetValidationContextType().(*CommonTlsContext_CombinedValidationContext); ok {
		return x.CombinedValidationContext
	}
	return nil
}

func (m *CommonTlsContext) GetValidationContextCertificateProvider() *CommonTlsContext_CertificateProvider {
	if x, ok := m.GetValidationContextType().(*CommonTlsContext_ValidationContextCertificateProvider); ok {
		return x.ValidationContextCertificateProvider
	}
	return nil
}

func (m *CommonTlsContext) GetAlpnProtocols() []string {
	if m != nil {
		return m.AlpnProtocols
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CommonTlsContext) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CommonTlsContext_ValidationContext)(nil),
		(*CommonTlsContext_ValidationContextSdsSecretConfig)(nil),
		(*CommonTlsContext_CombinedValidationContext)(nil),
		(*CommonTlsContext_ValidationContextCertificateProvider)(nil),
	}
}

type CommonTlsContext_CertificateProvider struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Config:
	//	*CommonTlsContext_CertificateProvider_TypedConfig
	Config               isCommonTlsContext_CertificateProvider_Config `protobuf_oneof:"config"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *CommonTlsContext_CertificateProvider) Reset()         { *m = CommonTlsContext_CertificateProvider{} }
func (m *CommonTlsContext_CertificateProvider) String() string { return proto.CompactTextString(m) }
func (*CommonTlsContext_CertificateProvider) ProtoMessage()    {}
func (*CommonTlsContext_CertificateProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_57bc8d868b0a9df1, []int{2, 0}
}

func (m *CommonTlsContext_CertificateProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonTlsContext_CertificateProvider.Unmarshal(m, b)
}
func (m *CommonTlsContext_CertificateProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonTlsContext_CertificateProvider.Marshal(b, m, deterministic)
}
func (m *CommonTlsContext_CertificateProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonTlsContext_CertificateProvider.Merge(m, src)
}
func (m *CommonTlsContext_CertificateProvider) XXX_Size() int {
	return xxx_messageInfo_CommonTlsContext_CertificateProvider.Size(m)
}
func (m *CommonTlsContext_CertificateProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonTlsContext_CertificateProvider.DiscardUnknown(m)
}

var xxx_messageInfo_CommonTlsContext_CertificateProvider proto.InternalMessageInfo

func (m *CommonTlsContext_CertificateProvider) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isCommonTlsContext_CertificateProvider_Config interface {
	isCommonTlsContext_CertificateProvider_Config()
}

type CommonTlsContext_CertificateProvider_TypedConfig struct {
	TypedConfig *v3.TypedExtensionConfig `protobuf:"bytes,2,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*CommonTlsContext_CertificateProvider_TypedConfig) isCommonTlsContext_CertificateProvider_Config() {
}

func (m *CommonTlsContext_CertificateProvider) GetConfig() isCommonTlsContext_CertificateProvider_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *CommonTlsContext_CertificateProvider) GetTypedConfig() *v3.TypedExtensionConfig {
	if x, ok := m.GetConfig().(*CommonTlsContext_CertificateProvider_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CommonTlsContext_CertificateProvider) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CommonTlsContext_CertificateProvider_TypedConfig)(nil),
	}
}

type CommonTlsContext_CombinedCertificateValidationContext struct {
	DefaultValidationContext             *CertificateValidationContext         `protobuf:"bytes,1,opt,name=default_validation_context,json=defaultValidationContext,proto3" json:"default_validation_context,omitempty"`
	ValidationContextSdsSecretConfig     *SdsSecretConfig                      `protobuf:"bytes,2,opt,name=validation_context_sds_secret_config,json=validationContextSdsSecretConfig,proto3" json:"validation_context_sds_secret_config,omitempty"`
	ValidationContextCertificateProvider *CommonTlsContext_CertificateProvider `protobuf:"bytes,3,opt,name=validation_context_certificate_provider,json=validationContextCertificateProvider,proto3" json:"validation_context_certificate_provider,omitempty"`
	XXX_NoUnkeyedLiteral                 struct{}                              `json:"-"`
	XXX_unrecognized                     []byte                                `json:"-"`
	XXX_sizecache                        int32                                 `json:"-"`
}

func (m *CommonTlsContext_CombinedCertificateValidationContext) Reset() {
	*m = CommonTlsContext_CombinedCertificateValidationContext{}
}
func (m *CommonTlsContext_CombinedCertificateValidationContext) String() string {
	return proto.CompactTextString(m)
}
func (*CommonTlsContext_CombinedCertificateValidationContext) ProtoMessage() {}
func (*CommonTlsContext_CombinedCertificateValidationContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_57bc8d868b0a9df1, []int{2, 1}
}

func (m *CommonTlsContext_CombinedCertificateValidationContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonTlsContext_CombinedCertificateValidationContext.Unmarshal(m, b)
}
func (m *CommonTlsContext_CombinedCertificateValidationContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonTlsContext_CombinedCertificateValidationContext.Marshal(b, m, deterministic)
}
func (m *CommonTlsContext_CombinedCertificateValidationContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonTlsContext_CombinedCertificateValidationContext.Merge(m, src)
}
func (m *CommonTlsContext_CombinedCertificateValidationContext) XXX_Size() int {
	return xxx_messageInfo_CommonTlsContext_CombinedCertificateValidationContext.Size(m)
}
func (m *CommonTlsContext_CombinedCertificateValidationContext) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonTlsContext_CombinedCertificateValidationContext.DiscardUnknown(m)
}

var xxx_messageInfo_CommonTlsContext_CombinedCertificateValidationContext proto.InternalMessageInfo

func (m *CommonTlsContext_CombinedCertificateValidationContext) GetDefaultValidationContext() *CertificateValidationContext {
	if m != nil {
		return m.DefaultValidationContext
	}
	return nil
}

func (m *CommonTlsContext_CombinedCertificateValidationContext) GetValidationContextSdsSecretConfig() *SdsSecretConfig {
	if m != nil {
		return m.ValidationContextSdsSecretConfig
	}
	return nil
}

func (m *CommonTlsContext_CombinedCertificateValidationContext) GetValidationContextCertificateProvider() *CommonTlsContext_CertificateProvider {
	if m != nil {
		return m.ValidationContextCertificateProvider
	}
	return nil
}

func init() {
	proto.RegisterType((*UpstreamTlsContext)(nil), "envoy.extensions.transport_sockets.tls.v3.UpstreamTlsContext")
	proto.RegisterType((*DownstreamTlsContext)(nil), "envoy.extensions.transport_sockets.tls.v3.DownstreamTlsContext")
	proto.RegisterType((*CommonTlsContext)(nil), "envoy.extensions.transport_sockets.tls.v3.CommonTlsContext")
	proto.RegisterType((*CommonTlsContext_CertificateProvider)(nil), "envoy.extensions.transport_sockets.tls.v3.CommonTlsContext.CertificateProvider")
	proto.RegisterType((*CommonTlsContext_CombinedCertificateValidationContext)(nil), "envoy.extensions.transport_sockets.tls.v3.CommonTlsContext.CombinedCertificateValidationContext")
}

func init() {
	proto.RegisterFile("envoy/extensions/transport_sockets/tls/v3/tls.proto", fileDescriptor_57bc8d868b0a9df1)
}

var fileDescriptor_57bc8d868b0a9df1 = []byte{
	// 1132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xcf, 0x6f, 0x1b, 0xc5,
	0x17, 0xcf, 0xda, 0x89, 0xe3, 0x4c, 0xbe, 0x49, 0xfc, 0x9d, 0x20, 0x75, 0xb3, 0x2d, 0x91, 0x63,
	0xb9, 0xc5, 0x2d, 0x62, 0x57, 0x72, 0x24, 0x7e, 0xa4, 0x07, 0xa4, 0x4d, 0xa1, 0x29, 0x08, 0x1a,
	0x36, 0x69, 0xe1, 0xb6, 0x4c, 0x76, 0x27, 0x66, 0xd5, 0xdd, 0x99, 0xed, 0xcc, 0xd8, 0xb1, 0x6f,
	0x39, 0x22, 0x04, 0x04, 0xf5, 0x84, 0x38, 0x21, 0x4e, 0x88, 0x3f, 0x80, 0x03, 0x37, 0x0e, 0x48,
	0x88, 0x1b, 0xff, 0x0a, 0x27, 0xc8, 0x01, 0xd0, 0xcc, 0xee, 0xda, 0x89, 0x77, 0xa5, 0xd8, 0x4d,
	0x2b, 0x4e, 0xf1, 0xce, 0x7b, 0xef, 0x33, 0x9f, 0xf7, 0xde, 0xe7, 0xbd, 0x09, 0xd8, 0xc4, 0xa4,
	0x47, 0x07, 0x16, 0xee, 0x0b, 0x4c, 0x78, 0x40, 0x09, 0xb7, 0x04, 0x43, 0x84, 0xc7, 0x94, 0x09,
	0x97, 0x53, 0xef, 0x11, 0x16, 0xdc, 0x12, 0x21, 0xb7, 0x7a, 0x9b, 0xf2, 0x8f, 0x19, 0x33, 0x2a,
	0x28, 0xbc, 0xa9, 0x82, 0xcc, 0x51, 0x90, 0x99, 0x0b, 0x32, 0xa5, 0x77, 0x6f, 0xd3, 0x68, 0x26,
	0xf8, 0x1e, 0x25, 0x87, 0x41, 0xc7, 0xf2, 0x28, 0xc3, 0x12, 0x6a, 0x18, 0x99, 0x00, 0x1a, 0xaf,
	0x4e, 0xce, 0xc2, 0xa3, 0x51, 0xf4, 0x34, 0x71, 0x1c, 0x7b, 0x0c, 0x8b, 0x34, 0x6e, 0xad, 0x43,
	0x69, 0x27, 0xc4, 0x96, 0xfa, 0x3a, 0xe8, 0x1e, 0x5a, 0x88, 0x0c, 0x52, 0xd3, 0xfa, 0xb8, 0xc9,
	0xef, 0x32, 0x24, 0x46, 0x54, 0x73, 0xf6, 0x23, 0x86, 0xe2, 0x18, 0x33, 0x9e, 0xd9, 0xbb, 0x7e,
	0x8c, 0x2c, 0x44, 0x08, 0x15, 0x2a, 0x8c, 0x5b, 0x51, 0xd0, 0x61, 0x48, 0xe0, 0xd4, 0xfe, 0x62,
	0xce, 0xce, 0x05, 0x12, 0xdd, 0x2c, 0x7c, 0x23, 0x67, 0xee, 0x61, 0x26, 0x53, 0x0b, 0x48, 0x27,
	0x75, 0xb9, 0xd2, 0x43, 0x61, 0xe0, 0x23, 0x81, 0xad, 0xec, 0x47, 0x62, 0x68, 0xfc, 0x5c, 0x02,
	0xf0, 0x41, 0xcc, 0x05, 0xc3, 0x28, 0xda, 0x0f, 0xf9, 0x36, 0x25, 0x02, 0xf7, 0x05, 0x0c, 0x00,
	0x4c, 0x8a, 0xe6, 0x8a, 0x90, 0xbb, 0x5e, 0x72, 0xaa, 0x6b, 0x75, 0xad, 0xb5, 0xd8, 0xbe, 0x6d,
	0x4e, 0xdc, 0x4a, 0x73, 0x5b, 0x81, 0x8c, 0x80, 0x9d, 0x9a, 0x37, 0x76, 0x02, 0x0d, 0x50, 0xe6,
	0x24, 0xd0, 0x4b, 0x75, 0xad, 0xb5, 0x60, 0x57, 0x4f, 0xed, 0x39, 0x56, 0x6e, 0xfd, 0xa3, 0x39,
	0xf2, 0x10, 0x5a, 0x60, 0x15, 0x85, 0x21, 0x3d, 0x72, 0x19, 0x26, 0xb8, 0x43, 0x45, 0xa0, 0xf2,
	0xd3, 0xcb, 0x75, 0xad, 0x55, 0x75, 0xa0, 0x32, 0x39, 0x67, 0x2d, 0xf0, 0x6d, 0x50, 0x8b, 0x50,
	0xdf, 0xe5, 0x98, 0x4b, 0x62, 0xee, 0x23, 0x3c, 0xe0, 0xfa, 0xac, 0x62, 0x7d, 0xcd, 0x4c, 0x9a,
	0x60, 0x66, 0x4d, 0x30, 0x1f, 0xdc, 0x23, 0x62, 0xb3, 0xfd, 0x10, 0x85, 0x5d, 0xec, 0x2c, 0x47,
	0xa8, 0xbf, 0x97, 0x04, 0xbd, 0x8b, 0x07, 0x7c, 0xeb, 0xe5, 0x6f, 0x7e, 0xf9, 0x74, 0xfd, 0x06,
	0x48, 0x94, 0x68, 0xa2, 0x38, 0x30, 0x7b, 0x6d, 0x13, 0x75, 0xc5, 0x27, 0x66, 0xbe, 0x58, 0x8d,
	0x93, 0x0a, 0x78, 0xe1, 0x0e, 0x3d, 0x22, 0xff, 0x65, 0x15, 0x3f, 0x02, 0x06, 0xc3, 0x8f, 0xbb,
	0x01, 0xc3, 0xae, 0x17, 0x06, 0x98, 0x08, 0xd7, 0xc3, 0x4c, 0x04, 0x87, 0x81, 0x87, 0x04, 0x56,
	0xc5, 0x5d, 0x6c, 0x1b, 0xb9, 0x12, 0xd8, 0x94, 0x86, 0x49, 0x01, 0xf4, 0x34, 0x7a, 0x5b, 0x05,
	0x6f, 0x8f, 0x62, 0xe1, 0x6d, 0xb0, 0x98, 0x21, 0xcb, 0x3e, 0x95, 0x2f, 0x84, 0x02, 0xa9, 0xfb,
	0x1e, 0x09, 0xe0, 0x63, 0xb0, 0x9a, 0xf5, 0x42, 0x04, 0x32, 0xa7, 0xb3, 0x2d, 0x79, 0x73, 0x8a,
	0x12, 0xec, 0x87, 0x3c, 0xed, 0xcf, 0xbe, 0xc2, 0x91, 0x5d, 0xda, 0x99, 0x71, 0xfe, 0xcf, 0xc7,
	0x0f, 0xe1, 0x17, 0x1a, 0xb8, 0x5e, 0x70, 0xa7, 0xcb, 0x7d, 0xee, 0x26, 0x03, 0xed, 0x26, 0x9b,
	0x45, 0x9f, 0x53, 0x2c, 0xb6, 0xa6, 0x60, 0xb1, 0xe7, 0xf3, 0x3d, 0x05, 0xb1, 0xad, 0x10, 0x76,
	0x66, 0x9c, 0x7a, 0x8e, 0xc0, 0x98, 0x0f, 0xfc, 0x00, 0x34, 0xfd, 0x80, 0xa3, 0x83, 0x10, 0xbb,
	0x72, 0x6a, 0x71, 0x88, 0x39, 0x1f, 0x0a, 0x94, 0x61, 0xde, 0x8d, 0x62, 0x25, 0xea, 0x79, 0x29,
	0xea, 0x9d, 0x19, 0x67, 0x23, 0xf5, 0xde, 0xcb, 0x9c, 0xd3, 0xbc, 0x9d, 0xa1, 0x2b, 0xdc, 0x07,
	0x2b, 0xa3, 0x0c, 0x23, 0x4c, 0xbb, 0x42, 0xaf, 0xa8, 0x5c, 0xd6, 0x72, 0x6d, 0xb9, 0x93, 0x6e,
	0x22, 0xbb, 0x76, 0x6a, 0x2f, 0xfd, 0xa0, 0x01, 0xa3, 0x52, 0x3d, 0x3e, 0x3e, 0x3e, 0xae, 0xb5,
	0x67, 0x9c, 0xe5, 0x21, 0x75, 0x05, 0xb1, 0xf5, 0x8a, 0xd4, 0x7c, 0x0b, 0xdc, 0xc8, 0x6b, 0xbe,
	0x48, 0xdc, 0xb6, 0x01, 0xf4, 0xa2, 0x32, 0x8b, 0x41, 0x8c, 0x1b, 0x7f, 0xad, 0x80, 0xda, 0xb8,
	0x68, 0xe1, 0x87, 0x00, 0xc8, 0x31, 0x88, 0x11, 0x43, 0x11, 0x4f, 0xa7, 0xe0, 0xf5, 0xe9, 0x24,
	0xb0, 0x2b, 0x63, 0xb1, 0xc0, 0x8c, 0x3b, 0x0b, 0x22, 0xfd, 0xe4, 0xd0, 0x07, 0x35, 0x35, 0x5f,
	0x23, 0xd1, 0x72, 0xbd, 0x54, 0x2f, 0xb7, 0x16, 0xdb, 0x6f, 0x4c, 0x07, 0x7f, 0x46, 0xf6, 0xce,
	0x8a, 0x38, 0xf7, 0xcd, 0xe1, 0x57, 0x1a, 0x68, 0x8c, 0x5d, 0x93, 0xd7, 0x14, 0xd7, 0x2b, 0xea,
	0xe2, 0x4b, 0x88, 0x4a, 0xed, 0xc0, 0x27, 0x5a, 0xa9, 0xa6, 0x39, 0xeb, 0xe7, 0x39, 0x8c, 0x39,
	0x72, 0xf8, 0xad, 0x06, 0x9a, 0xe3, 0x94, 0xce, 0xfe, 0x8e, 0x19, 0xed, 0x05, 0x3e, 0x66, 0xfa,
	0x82, 0x2a, 0xf6, 0xfd, 0x4b, 0xac, 0x1c, 0xf3, 0x0c, 0x8f, 0xdd, 0x14, 0xd6, 0xd9, 0x38, 0xcf,
	0xaf, 0xc0, 0x05, 0xf6, 0x01, 0x4c, 0x5f, 0x1c, 0x29, 0x94, 0x6c, 0x05, 0x26, 0x4b, 0xe4, 0xee,
	0x34, 0x7c, 0x46, 0xd8, 0x0f, 0x87, 0x78, 0x29, 0x37, 0xb9, 0x07, 0x7a, 0xe3, 0x87, 0xf0, 0x73,
	0x0d, 0x34, 0xf3, 0x57, 0x17, 0xac, 0x81, 0xf9, 0x67, 0xb1, 0x06, 0x72, 0xf7, 0x8f, 0xaf, 0x81,
	0xef, 0x34, 0x70, 0xd5, 0xa3, 0xd1, 0x41, 0x40, 0xb0, 0xef, 0x16, 0x94, 0xa4, 0xaa, 0x58, 0x7c,
	0x7c, 0xa9, 0x16, 0xa5, 0xf0, 0x17, 0xd4, 0x6a, 0x2d, 0xa3, 0x91, 0x33, 0xc2, 0xef, 0x35, 0xf0,
	0x52, 0x41, 0xcd, 0x0a, 0x35, 0x05, 0x9e, 0x8b, 0xa6, 0x76, 0x66, 0x9c, 0x66, 0xae, 0x96, 0x45,
	0xc2, 0xba, 0x0e, 0x96, 0x51, 0x18, 0x13, 0x57, 0x6d, 0x3a, 0x8f, 0x86, 0xf2, 0x51, 0x29, 0xb7,
	0x16, 0x9c, 0x25, 0x79, 0xba, 0x9b, 0x1d, 0x1a, 0x4f, 0x34, 0xb0, 0x5a, 0x14, 0x7e, 0x15, 0xcc,
	0x12, 0x14, 0x61, 0xb5, 0x86, 0x16, 0xec, 0xf9, 0x53, 0x7b, 0x96, 0x95, 0xea, 0x9a, 0xa3, 0x0e,
	0xe1, 0x7d, 0xf0, 0x3f, 0xb9, 0xc6, 0xfc, 0x4c, 0x21, 0xc9, 0xf3, 0x79, 0x2b, 0x4d, 0x35, 0x39,
	0x34, 0xe5, 0xff, 0xa5, 0x6a, 0x6f, 0x48, 0xcf, 0xb7, 0xb2, 0xf4, 0x87, 0x8a, 0x58, 0x54, 0x08,
	0xe9, 0x48, 0x2f, 0x81, 0x4a, 0x12, 0x05, 0xcb, 0x7f, 0xda, 0x9a, 0xf1, 0xe5, 0x1c, 0x68, 0x4e,
	0xd2, 0x2c, 0x78, 0xa2, 0x01, 0xc3, 0xc7, 0x87, 0xa8, 0x1b, 0x8a, 0x22, 0xcd, 0x68, 0xcf, 0x74,
	0x8c, 0xd4, 0xe2, 0xf9, 0x4c, 0x2d, 0x1e, 0x3d, 0xbd, 0x34, 0xcf, 0xe8, 0xc7, 0x49, 0xa7, 0xaa,
	0x74, 0xd9, 0xa9, 0xb2, 0x6f, 0x65, 0x74, 0xfe, 0xf8, 0xfa, 0xef, 0x93, 0xb9, 0x6b, 0xd0, 0xf0,
	0x07, 0x04, 0x45, 0x81, 0x57, 0x50, 0x89, 0x09, 0xe6, 0xef, 0xb7, 0x29, 0xa4, 0x5d, 0x7e, 0x2e,
	0xd2, 0xb6, 0x1b, 0x13, 0x24, 0x32, 0x91, 0xf8, 0xb7, 0xde, 0x97, 0x4f, 0xf5, 0x3d, 0x70, 0x37,
	0xff, 0x54, 0x3f, 0xd5, 0x52, 0xd8, 0xba, 0x29, 0xf1, 0x9a, 0xa0, 0x71, 0x31, 0x9e, 0xbd, 0x06,
	0xae, 0x14, 0x94, 0x51, 0x8a, 0xfd, 0x9d, 0xd9, 0xea, 0x5c, 0xad, 0x62, 0xbf, 0xf7, 0xd3, 0xf1,
	0xaf, 0xbf, 0x57, 0x4a, 0xb5, 0x12, 0x78, 0x2d, 0xa0, 0x49, 0x09, 0x63, 0x46, 0xfb, 0x83, 0xc9,
	0xab, 0x69, 0x57, 0xe5, 0x53, 0x2f, 0x47, 0x78, 0x57, 0x3b, 0xa8, 0xa8, 0x01, 0xdf, 0xfc, 0x37,
	0x00, 0x00, 0xff, 0xff, 0xe2, 0x5d, 0x65, 0x92, 0x75, 0x0e, 0x00, 0x00,
}
