// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/keyword_plan_idea_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [KeywordIdeaService.GenerateKeywordIdeas][].
type GenerateKeywordIdeasRequest struct {
	// The ID of the customer with the recommendation.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The resource name of the language to target.
	// Required
	Language *wrappers.StringValue `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	// The resource names of the location to target.
	// Max 10
	GeoTargetConstants []*wrappers.StringValue `protobuf:"bytes,8,rep,name=geo_target_constants,json=geoTargetConstants,proto3" json:"geo_target_constants,omitempty"`
	// Targeting network.
	KeywordPlanNetwork enums.KeywordPlanNetworkEnum_KeywordPlanNetwork `protobuf:"varint,9,opt,name=keyword_plan_network,json=keywordPlanNetwork,proto3,enum=google.ads.googleads.v1.enums.KeywordPlanNetworkEnum_KeywordPlanNetwork" json:"keyword_plan_network,omitempty"`
	// The type of seed to generate keyword ideas.
	//
	// Types that are valid to be assigned to Seed:
	//	*GenerateKeywordIdeasRequest_KeywordAndUrlSeed
	//	*GenerateKeywordIdeasRequest_KeywordSeed
	//	*GenerateKeywordIdeasRequest_UrlSeed
	Seed                 isGenerateKeywordIdeasRequest_Seed `protobuf_oneof:"seed"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *GenerateKeywordIdeasRequest) Reset()         { *m = GenerateKeywordIdeasRequest{} }
func (m *GenerateKeywordIdeasRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeasRequest) ProtoMessage()    {}
func (*GenerateKeywordIdeasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_idea_service_79f4adf1893c8d4e, []int{0}
}
func (m *GenerateKeywordIdeasRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Marshal(b, m, deterministic)
}
func (dst *GenerateKeywordIdeasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeasRequest.Merge(dst, src)
}
func (m *GenerateKeywordIdeasRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Size(m)
}
func (m *GenerateKeywordIdeasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeasRequest proto.InternalMessageInfo

func (m *GenerateKeywordIdeasRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *GenerateKeywordIdeasRequest) GetLanguage() *wrappers.StringValue {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetGeoTargetConstants() []*wrappers.StringValue {
	if m != nil {
		return m.GeoTargetConstants
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordPlanNetwork() enums.KeywordPlanNetworkEnum_KeywordPlanNetwork {
	if m != nil {
		return m.KeywordPlanNetwork
	}
	return enums.KeywordPlanNetworkEnum_UNSPECIFIED
}

type isGenerateKeywordIdeasRequest_Seed interface {
	isGenerateKeywordIdeasRequest_Seed()
}

type GenerateKeywordIdeasRequest_KeywordAndUrlSeed struct {
	KeywordAndUrlSeed *KeywordAndUrlSeed `protobuf:"bytes,2,opt,name=keyword_and_url_seed,json=keywordAndUrlSeed,proto3,oneof"`
}

type GenerateKeywordIdeasRequest_KeywordSeed struct {
	KeywordSeed *KeywordSeed `protobuf:"bytes,3,opt,name=keyword_seed,json=keywordSeed,proto3,oneof"`
}

type GenerateKeywordIdeasRequest_UrlSeed struct {
	UrlSeed *UrlSeed `protobuf:"bytes,5,opt,name=url_seed,json=urlSeed,proto3,oneof"`
}

func (*GenerateKeywordIdeasRequest_KeywordAndUrlSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (*GenerateKeywordIdeasRequest_KeywordSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (*GenerateKeywordIdeasRequest_UrlSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (m *GenerateKeywordIdeasRequest) GetSeed() isGenerateKeywordIdeasRequest_Seed {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordAndUrlSeed() *KeywordAndUrlSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_KeywordAndUrlSeed); ok {
		return x.KeywordAndUrlSeed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordSeed() *KeywordSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_KeywordSeed); ok {
		return x.KeywordSeed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetUrlSeed() *UrlSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_UrlSeed); ok {
		return x.UrlSeed
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GenerateKeywordIdeasRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GenerateKeywordIdeasRequest_OneofMarshaler, _GenerateKeywordIdeasRequest_OneofUnmarshaler, _GenerateKeywordIdeasRequest_OneofSizer, []interface{}{
		(*GenerateKeywordIdeasRequest_KeywordAndUrlSeed)(nil),
		(*GenerateKeywordIdeasRequest_KeywordSeed)(nil),
		(*GenerateKeywordIdeasRequest_UrlSeed)(nil),
	}
}

func _GenerateKeywordIdeasRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GenerateKeywordIdeasRequest)
	// seed
	switch x := m.Seed.(type) {
	case *GenerateKeywordIdeasRequest_KeywordAndUrlSeed:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KeywordAndUrlSeed); err != nil {
			return err
		}
	case *GenerateKeywordIdeasRequest_KeywordSeed:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KeywordSeed); err != nil {
			return err
		}
	case *GenerateKeywordIdeasRequest_UrlSeed:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UrlSeed); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GenerateKeywordIdeasRequest.Seed has unexpected type %T", x)
	}
	return nil
}

func _GenerateKeywordIdeasRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GenerateKeywordIdeasRequest)
	switch tag {
	case 2: // seed.keyword_and_url_seed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KeywordAndUrlSeed)
		err := b.DecodeMessage(msg)
		m.Seed = &GenerateKeywordIdeasRequest_KeywordAndUrlSeed{msg}
		return true, err
	case 3: // seed.keyword_seed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KeywordSeed)
		err := b.DecodeMessage(msg)
		m.Seed = &GenerateKeywordIdeasRequest_KeywordSeed{msg}
		return true, err
	case 5: // seed.url_seed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UrlSeed)
		err := b.DecodeMessage(msg)
		m.Seed = &GenerateKeywordIdeasRequest_UrlSeed{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GenerateKeywordIdeasRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GenerateKeywordIdeasRequest)
	// seed
	switch x := m.Seed.(type) {
	case *GenerateKeywordIdeasRequest_KeywordAndUrlSeed:
		s := proto.Size(x.KeywordAndUrlSeed)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GenerateKeywordIdeasRequest_KeywordSeed:
		s := proto.Size(x.KeywordSeed)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GenerateKeywordIdeasRequest_UrlSeed:
		s := proto.Size(x.UrlSeed)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Keyword And Url Seed
type KeywordAndUrlSeed struct {
	// The URL to crawl in order to generate keyword ideas.
	Url *wrappers.StringValue `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Requires at least one keyword.
	Keywords             []*wrappers.StringValue `protobuf:"bytes,2,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordAndUrlSeed) Reset()         { *m = KeywordAndUrlSeed{} }
func (m *KeywordAndUrlSeed) String() string { return proto.CompactTextString(m) }
func (*KeywordAndUrlSeed) ProtoMessage()    {}
func (*KeywordAndUrlSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_idea_service_79f4adf1893c8d4e, []int{1}
}
func (m *KeywordAndUrlSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordAndUrlSeed.Unmarshal(m, b)
}
func (m *KeywordAndUrlSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordAndUrlSeed.Marshal(b, m, deterministic)
}
func (dst *KeywordAndUrlSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordAndUrlSeed.Merge(dst, src)
}
func (m *KeywordAndUrlSeed) XXX_Size() int {
	return xxx_messageInfo_KeywordAndUrlSeed.Size(m)
}
func (m *KeywordAndUrlSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordAndUrlSeed.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordAndUrlSeed proto.InternalMessageInfo

func (m *KeywordAndUrlSeed) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *KeywordAndUrlSeed) GetKeywords() []*wrappers.StringValue {
	if m != nil {
		return m.Keywords
	}
	return nil
}

// Keyword Seed
type KeywordSeed struct {
	// Requires at least one keyword.
	Keywords             []*wrappers.StringValue `protobuf:"bytes,1,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordSeed) Reset()         { *m = KeywordSeed{} }
func (m *KeywordSeed) String() string { return proto.CompactTextString(m) }
func (*KeywordSeed) ProtoMessage()    {}
func (*KeywordSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_idea_service_79f4adf1893c8d4e, []int{2}
}
func (m *KeywordSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordSeed.Unmarshal(m, b)
}
func (m *KeywordSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordSeed.Marshal(b, m, deterministic)
}
func (dst *KeywordSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordSeed.Merge(dst, src)
}
func (m *KeywordSeed) XXX_Size() int {
	return xxx_messageInfo_KeywordSeed.Size(m)
}
func (m *KeywordSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordSeed.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordSeed proto.InternalMessageInfo

func (m *KeywordSeed) GetKeywords() []*wrappers.StringValue {
	if m != nil {
		return m.Keywords
	}
	return nil
}

// Url Seed
type UrlSeed struct {
	// The URL to crawl in order to generate keyword ideas.
	Url                  *wrappers.StringValue `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UrlSeed) Reset()         { *m = UrlSeed{} }
func (m *UrlSeed) String() string { return proto.CompactTextString(m) }
func (*UrlSeed) ProtoMessage()    {}
func (*UrlSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_idea_service_79f4adf1893c8d4e, []int{3}
}
func (m *UrlSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UrlSeed.Unmarshal(m, b)
}
func (m *UrlSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UrlSeed.Marshal(b, m, deterministic)
}
func (dst *UrlSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UrlSeed.Merge(dst, src)
}
func (m *UrlSeed) XXX_Size() int {
	return xxx_messageInfo_UrlSeed.Size(m)
}
func (m *UrlSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_UrlSeed.DiscardUnknown(m)
}

var xxx_messageInfo_UrlSeed proto.InternalMessageInfo

func (m *UrlSeed) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

// Response message for [KeywordIdeaService.GenerateKeywordIdeas][].
type GenerateKeywordIdeaResponse struct {
	// Results of generating keyword ideas.
	Results              []*GenerateKeywordIdeaResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GenerateKeywordIdeaResponse) Reset()         { *m = GenerateKeywordIdeaResponse{} }
func (m *GenerateKeywordIdeaResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeaResponse) ProtoMessage()    {}
func (*GenerateKeywordIdeaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_idea_service_79f4adf1893c8d4e, []int{4}
}
func (m *GenerateKeywordIdeaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Marshal(b, m, deterministic)
}
func (dst *GenerateKeywordIdeaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeaResponse.Merge(dst, src)
}
func (m *GenerateKeywordIdeaResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Size(m)
}
func (m *GenerateKeywordIdeaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeaResponse proto.InternalMessageInfo

func (m *GenerateKeywordIdeaResponse) GetResults() []*GenerateKeywordIdeaResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result of generating keyword ideas.
type GenerateKeywordIdeaResult struct {
	// Text of the keyword idea.
	// As in Keyword Plan historical metrics, this text may not be an actual
	// keyword, but the canonical form of multiple keywords.
	// See KeywordPlanKeywordHistoricalMetrics message in KeywordPlanService.
	Text *wrappers.StringValue `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// The historical metrics for the keyword
	KeywordIdeaMetrics   *common.KeywordPlanHistoricalMetrics `protobuf:"bytes,3,opt,name=keyword_idea_metrics,json=keywordIdeaMetrics,proto3" json:"keyword_idea_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *GenerateKeywordIdeaResult) Reset()         { *m = GenerateKeywordIdeaResult{} }
func (m *GenerateKeywordIdeaResult) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeaResult) ProtoMessage()    {}
func (*GenerateKeywordIdeaResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_keyword_plan_idea_service_79f4adf1893c8d4e, []int{5}
}
func (m *GenerateKeywordIdeaResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeaResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Marshal(b, m, deterministic)
}
func (dst *GenerateKeywordIdeaResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeaResult.Merge(dst, src)
}
func (m *GenerateKeywordIdeaResult) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Size(m)
}
func (m *GenerateKeywordIdeaResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeaResult.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeaResult proto.InternalMessageInfo

func (m *GenerateKeywordIdeaResult) GetText() *wrappers.StringValue {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *GenerateKeywordIdeaResult) GetKeywordIdeaMetrics() *common.KeywordPlanHistoricalMetrics {
	if m != nil {
		return m.KeywordIdeaMetrics
	}
	return nil
}

func init() {
	proto.RegisterType((*GenerateKeywordIdeasRequest)(nil), "google.ads.googleads.v1.services.GenerateKeywordIdeasRequest")
	proto.RegisterType((*KeywordAndUrlSeed)(nil), "google.ads.googleads.v1.services.KeywordAndUrlSeed")
	proto.RegisterType((*KeywordSeed)(nil), "google.ads.googleads.v1.services.KeywordSeed")
	proto.RegisterType((*UrlSeed)(nil), "google.ads.googleads.v1.services.UrlSeed")
	proto.RegisterType((*GenerateKeywordIdeaResponse)(nil), "google.ads.googleads.v1.services.GenerateKeywordIdeaResponse")
	proto.RegisterType((*GenerateKeywordIdeaResult)(nil), "google.ads.googleads.v1.services.GenerateKeywordIdeaResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeywordPlanIdeaServiceClient is the client API for KeywordPlanIdeaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordPlanIdeaServiceClient interface {
	// Returns a list of keyword ideas.
	GenerateKeywordIdeas(ctx context.Context, in *GenerateKeywordIdeasRequest, opts ...grpc.CallOption) (*GenerateKeywordIdeaResponse, error)
}

type keywordPlanIdeaServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeywordPlanIdeaServiceClient(cc *grpc.ClientConn) KeywordPlanIdeaServiceClient {
	return &keywordPlanIdeaServiceClient{cc}
}

func (c *keywordPlanIdeaServiceClient) GenerateKeywordIdeas(ctx context.Context, in *GenerateKeywordIdeasRequest, opts ...grpc.CallOption) (*GenerateKeywordIdeaResponse, error) {
	out := new(GenerateKeywordIdeaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.KeywordPlanIdeaService/GenerateKeywordIdeas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanIdeaServiceServer is the server API for KeywordPlanIdeaService service.
type KeywordPlanIdeaServiceServer interface {
	// Returns a list of keyword ideas.
	GenerateKeywordIdeas(context.Context, *GenerateKeywordIdeasRequest) (*GenerateKeywordIdeaResponse, error)
}

func RegisterKeywordPlanIdeaServiceServer(s *grpc.Server, srv KeywordPlanIdeaServiceServer) {
	s.RegisterService(&_KeywordPlanIdeaService_serviceDesc, srv)
}

func _KeywordPlanIdeaService_GenerateKeywordIdeas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKeywordIdeasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordIdeas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.KeywordPlanIdeaService/GenerateKeywordIdeas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordIdeas(ctx, req.(*GenerateKeywordIdeasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordPlanIdeaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.KeywordPlanIdeaService",
	HandlerType: (*KeywordPlanIdeaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateKeywordIdeas",
			Handler:    _KeywordPlanIdeaService_GenerateKeywordIdeas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/keyword_plan_idea_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/keyword_plan_idea_service.proto", fileDescriptor_keyword_plan_idea_service_79f4adf1893c8d4e)
}

var fileDescriptor_keyword_plan_idea_service_79f4adf1893c8d4e = []byte{
	// 727 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xfe, 0x39, 0xe9, 0xaf, 0x69, 0x37, 0x08, 0xa9, 0xab, 0x0a, 0x85, 0xb6, 0x82, 0x28, 0xea,
	0xa1, 0x54, 0xc2, 0x26, 0xe9, 0x81, 0xe2, 0x12, 0x89, 0x14, 0x41, 0x52, 0x21, 0xaa, 0xca, 0xa5,
	0x39, 0xa0, 0x48, 0xd6, 0x36, 0x9e, 0x5a, 0x56, 0x9c, 0xdd, 0xb0, 0xbb, 0x6e, 0xf9, 0xa3, 0x5e,
	0xfa, 0x0a, 0xbc, 0x01, 0x47, 0xde, 0x81, 0x17, 0xe0, 0x84, 0xd4, 0x57, 0xe0, 0xc4, 0x81, 0x67,
	0x40, 0xb6, 0x77, 0x93, 0xb4, 0x75, 0x48, 0xe9, 0x6d, 0x3c, 0xfb, 0xcd, 0x37, 0xdf, 0xce, 0xcc,
	0x8e, 0xd1, 0x33, 0x9f, 0x31, 0x3f, 0x04, 0x8b, 0x78, 0xc2, 0x4a, 0xcd, 0xd8, 0x3a, 0xae, 0x5a,
	0x02, 0xf8, 0x71, 0xd0, 0x05, 0x61, 0xf5, 0xe0, 0xc3, 0x09, 0xe3, 0x9e, 0x3b, 0x08, 0x09, 0x75,
	0x03, 0x0f, 0x88, 0xab, 0x8e, 0xcc, 0x01, 0x67, 0x92, 0xe1, 0x72, 0x1a, 0x66, 0x12, 0x4f, 0x98,
	0x43, 0x06, 0xf3, 0xb8, 0x6a, 0x6a, 0x86, 0xa5, 0xcd, 0x49, 0x39, 0xba, 0xac, 0xdf, 0x67, 0xf4,
	0x62, 0x86, 0xd4, 0x97, 0x72, 0x4f, 0x8e, 0x04, 0x1a, 0xf5, 0x2f, 0x49, 0xa3, 0x20, 0x4f, 0x18,
	0xef, 0xa9, 0xc8, 0x15, 0x1d, 0x39, 0x08, 0x2c, 0x42, 0x29, 0x93, 0x44, 0x06, 0x8c, 0x0a, 0x75,
	0x7a, 0x4f, 0x9d, 0x26, 0x5f, 0x87, 0xd1, 0x91, 0x75, 0xc2, 0xc9, 0x60, 0x00, 0x5c, 0x9d, 0x57,
	0x7e, 0xcc, 0xa0, 0xe5, 0x26, 0x50, 0xe0, 0x44, 0xc2, 0xab, 0x34, 0xc9, 0x8e, 0x07, 0x44, 0x38,
	0xf0, 0x2e, 0x02, 0x21, 0xf1, 0x7d, 0x54, 0xec, 0x46, 0x42, 0xb2, 0x3e, 0x70, 0x37, 0xf0, 0x4a,
	0x46, 0xd9, 0x58, 0x9b, 0x77, 0x90, 0x76, 0xed, 0x78, 0x78, 0x13, 0xcd, 0x85, 0x84, 0xfa, 0x11,
	0xf1, 0xa1, 0x54, 0x28, 0x1b, 0x6b, 0xc5, 0xda, 0x8a, 0x2a, 0x8e, 0xa9, 0x73, 0x9a, 0xfb, 0x92,
	0x07, 0xd4, 0x6f, 0x93, 0x30, 0x02, 0x67, 0x88, 0xc6, 0xbb, 0x68, 0xd1, 0x07, 0xe6, 0x4a, 0xc2,
	0x7d, 0x90, 0x6e, 0x97, 0x51, 0x21, 0x09, 0x95, 0xa2, 0x34, 0x57, 0xce, 0x4f, 0x65, 0xc1, 0x3e,
	0xb0, 0x37, 0x49, 0xe0, 0x73, 0x1d, 0x87, 0x3f, 0xa2, 0xc5, 0xac, 0x32, 0x95, 0xe6, 0xcb, 0xc6,
	0xda, 0xed, 0x5a, 0xcb, 0x9c, 0xd4, 0xbd, 0xa4, 0xc2, 0xa6, 0xba, 0xfc, 0x5e, 0x48, 0xe8, 0x6e,
	0x1a, 0xf8, 0x82, 0x46, 0xfd, 0x0c, 0xb7, 0x83, 0x7b, 0x57, 0x7c, 0xf8, 0x68, 0x94, 0x9b, 0x50,
	0xcf, 0x8d, 0x78, 0xe8, 0x0a, 0x00, 0xaf, 0x94, 0x4b, 0x2a, 0xb2, 0x61, 0x4e, 0x9b, 0x1c, 0x9d,
	0xa7, 0x41, 0xbd, 0x03, 0x1e, 0xee, 0x03, 0x78, 0xad, 0xff, 0x9c, 0x85, 0xde, 0x65, 0x27, 0x76,
	0xd0, 0x2d, 0x9d, 0x27, 0xe1, 0xcf, 0x27, 0xfc, 0x0f, 0xaf, 0xcd, 0xaf, 0x98, 0x8b, 0xbd, 0xd1,
	0x27, 0x7e, 0x89, 0xe6, 0x86, 0x7a, 0xff, 0x4f, 0xf8, 0x1e, 0x4c, 0xe7, 0x1b, 0xa9, 0x2c, 0x44,
	0xa9, 0xb9, 0x3d, 0x8b, 0x66, 0x62, 0x8e, 0xca, 0x29, 0x5a, 0xb8, 0x72, 0x1b, 0x6c, 0xa2, 0x7c,
	0xc4, 0xc3, 0x64, 0x7e, 0xa6, 0xf5, 0x36, 0x06, 0xc6, 0x63, 0xa5, 0x34, 0x8a, 0x52, 0xee, 0x1a,
	0x03, 0x31, 0x44, 0x57, 0x9a, 0xa8, 0x38, 0x76, 0xd9, 0x0b, 0x44, 0xc6, 0x3f, 0x11, 0x3d, 0x41,
	0x85, 0x1b, 0xaa, 0xaf, 0xc8, 0xcc, 0x47, 0xe5, 0x80, 0x18, 0x30, 0x2a, 0x00, 0x1f, 0xa0, 0x02,
	0x07, 0x11, 0x85, 0x52, 0x4b, 0xda, 0x9a, 0x5e, 0xf0, 0x6c, 0xbe, 0x28, 0x94, 0x8e, 0xe6, 0xaa,
	0x7c, 0x33, 0xd0, 0xdd, 0x89, 0x30, 0xfc, 0x08, 0xcd, 0x48, 0x78, 0x2f, 0xd5, 0x48, 0xfe, 0xfd,
	0x12, 0x09, 0x12, 0xd3, 0xd1, 0x50, 0x27, 0xdb, 0xb0, 0x0f, 0x92, 0x07, 0x5d, 0xa1, 0x86, 0xee,
	0xe9, 0x44, 0xcd, 0x6a, 0xb1, 0x8d, 0x3d, 0x9d, 0x56, 0x20, 0x24, 0xe3, 0x41, 0x97, 0x84, 0xaf,
	0x53, 0x8e, 0xe1, 0x23, 0x8a, 0x05, 0x2a, 0x5f, 0xed, 0xb7, 0x81, 0xee, 0x8c, 0x05, 0xc5, 0x47,
	0xfb, 0xe9, 0xf5, 0xf1, 0xb9, 0x81, 0x16, 0xb3, 0xd6, 0x14, 0xae, 0xdf, 0xa8, 0x72, 0x7a, 0xbd,
	0x2d, 0xd5, 0x6f, 0x5a, 0xf8, 0xa4, 0x91, 0x95, 0xfa, 0xd9, 0xf9, 0xcf, 0xcf, 0xb9, 0xc7, 0x95,
	0x5a, 0xb2, 0xe3, 0xd5, 0x52, 0x14, 0xd6, 0xa7, 0xb1, 0x95, 0x59, 0x5f, 0x3f, 0xb5, 0xfd, 0x0c,
	0x05, 0xb6, 0xb1, 0xbe, 0x7d, 0x96, 0x43, 0xab, 0x5d, 0xd6, 0x9f, 0xaa, 0x61, 0x7b, 0x39, 0xbb,
	0x2c, 0x7b, 0x71, 0xef, 0xf6, 0x8c, 0xb7, 0x2d, 0x45, 0xe0, 0xb3, 0x78, 0xb9, 0x9a, 0x8c, 0xfb,
	0x96, 0x0f, 0x34, 0xe9, 0xac, 0xfe, 0x99, 0x0c, 0x02, 0x31, 0xf9, 0xcf, 0xb7, 0xa5, 0x8d, 0x2f,
	0xb9, 0x7c, 0xb3, 0xd1, 0xf8, 0x9a, 0x2b, 0x37, 0x53, 0xc2, 0x86, 0x27, 0xcc, 0xd4, 0x8c, 0xad,
	0x76, 0xd5, 0x54, 0x89, 0xc5, 0x77, 0x0d, 0xe9, 0x34, 0x3c, 0xd1, 0x19, 0x42, 0x3a, 0xed, 0x6a,
	0x47, 0x43, 0x7e, 0xe5, 0x56, 0x53, 0xbf, 0x6d, 0x37, 0x3c, 0x61, 0xdb, 0x43, 0x90, 0x6d, 0xb7,
	0xab, 0xb6, 0xad, 0x61, 0x87, 0xb3, 0x89, 0xce, 0x8d, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfd,
	0x13, 0x9c, 0xd6, 0xa0, 0x07, 0x00, 0x00,
}