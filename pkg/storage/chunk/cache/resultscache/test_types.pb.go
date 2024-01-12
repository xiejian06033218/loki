// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/storage/chunk/cache/resultscache/test_types.proto

package resultscache

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MockRequest struct {
	Path           string         `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Start          time.Time      `protobuf:"bytes,2,opt,name=start,proto3,stdtime" json:"start"`
	End            time.Time      `protobuf:"bytes,3,opt,name=end,proto3,stdtime" json:"end"`
	Step           int64          `protobuf:"varint,4,opt,name=step,proto3" json:"step,omitempty"`
	Query          string         `protobuf:"bytes,6,opt,name=query,proto3" json:"query,omitempty"`
	CachingOptions CachingOptions `protobuf:"bytes,7,opt,name=cachingOptions,proto3" json:"cachingOptions"`
}

func (m *MockRequest) Reset()      { *m = MockRequest{} }
func (*MockRequest) ProtoMessage() {}
func (*MockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b2c489557407809, []int{0}
}
func (m *MockRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MockRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MockRequest.Merge(m, src)
}
func (m *MockRequest) XXX_Size() int {
	return m.Size()
}
func (m *MockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MockRequest proto.InternalMessageInfo

func (m *MockRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *MockRequest) GetStart() time.Time {
	if m != nil {
		return m.Start
	}
	return time.Time{}
}

func (m *MockRequest) GetEnd() time.Time {
	if m != nil {
		return m.End
	}
	return time.Time{}
}

func (m *MockRequest) GetStep() int64 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *MockRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *MockRequest) GetCachingOptions() CachingOptions {
	if m != nil {
		return m.CachingOptions
	}
	return CachingOptions{}
}

type MockResponse struct {
	Labels  []*MockLabelsPair `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	Samples []*MockSample     `protobuf:"bytes,2,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (m *MockResponse) Reset()      { *m = MockResponse{} }
func (*MockResponse) ProtoMessage() {}
func (*MockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b2c489557407809, []int{1}
}
func (m *MockResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MockResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MockResponse.Merge(m, src)
}
func (m *MockResponse) XXX_Size() int {
	return m.Size()
}
func (m *MockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MockResponse proto.InternalMessageInfo

func (m *MockResponse) GetLabels() []*MockLabelsPair {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *MockResponse) GetSamples() []*MockSample {
	if m != nil {
		return m.Samples
	}
	return nil
}

type MockLabelsPair struct {
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *MockLabelsPair) Reset()      { *m = MockLabelsPair{} }
func (*MockLabelsPair) ProtoMessage() {}
func (*MockLabelsPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b2c489557407809, []int{2}
}
func (m *MockLabelsPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MockLabelsPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MockLabelsPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MockLabelsPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MockLabelsPair.Merge(m, src)
}
func (m *MockLabelsPair) XXX_Size() int {
	return m.Size()
}
func (m *MockLabelsPair) XXX_DiscardUnknown() {
	xxx_messageInfo_MockLabelsPair.DiscardUnknown(m)
}

var xxx_messageInfo_MockLabelsPair proto.InternalMessageInfo

func (m *MockLabelsPair) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MockLabelsPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type MockSample struct {
	Value       float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	TimestampMs int64   `protobuf:"varint,2,opt,name=timestamp_ms,json=timestampMs,proto3" json:"timestamp_ms,omitempty"`
}

func (m *MockSample) Reset()      { *m = MockSample{} }
func (*MockSample) ProtoMessage() {}
func (*MockSample) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b2c489557407809, []int{3}
}
func (m *MockSample) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MockSample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MockSample.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MockSample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MockSample.Merge(m, src)
}
func (m *MockSample) XXX_Size() int {
	return m.Size()
}
func (m *MockSample) XXX_DiscardUnknown() {
	xxx_messageInfo_MockSample.DiscardUnknown(m)
}

var xxx_messageInfo_MockSample proto.InternalMessageInfo

func (m *MockSample) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *MockSample) GetTimestampMs() int64 {
	if m != nil {
		return m.TimestampMs
	}
	return 0
}

func init() {
	proto.RegisterType((*MockRequest)(nil), "resultscache.MockRequest")
	proto.RegisterType((*MockResponse)(nil), "resultscache.MockResponse")
	proto.RegisterType((*MockLabelsPair)(nil), "resultscache.MockLabelsPair")
	proto.RegisterType((*MockSample)(nil), "resultscache.MockSample")
}

func init() {
	proto.RegisterFile("pkg/storage/chunk/cache/resultscache/test_types.proto", fileDescriptor_5b2c489557407809)
}

var fileDescriptor_5b2c489557407809 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x3f, 0x6f, 0x13, 0x31,
	0x14, 0x3f, 0xe7, 0xd2, 0x94, 0x3a, 0x51, 0x07, 0xab, 0xc3, 0x29, 0x42, 0x4e, 0xc8, 0x94, 0xe9,
	0x2c, 0x95, 0x3f, 0x43, 0xc5, 0x14, 0xc4, 0x82, 0xa8, 0x40, 0x86, 0x89, 0xa5, 0x72, 0x0e, 0xd7,
	0x39, 0xe5, 0xee, 0xec, 0xde, 0xf3, 0x21, 0xba, 0xb1, 0xb3, 0xf4, 0x63, 0xf0, 0x51, 0x3a, 0x66,
	0xec, 0x04, 0xe4, 0xb2, 0x30, 0xf6, 0x23, 0x20, 0xfb, 0x92, 0x36, 0xa5, 0x0b, 0xdd, 0xde, 0xf3,
	0xfb, 0xfd, 0xb1, 0x7e, 0xef, 0xe1, 0xe7, 0x66, 0xae, 0x18, 0x58, 0x5d, 0x0a, 0x25, 0x59, 0x32,
	0xab, 0x8a, 0x39, 0x4b, 0x44, 0x32, 0x93, 0xac, 0x94, 0x50, 0x65, 0x16, 0x9a, 0xc6, 0x4a, 0xb0,
	0x27, 0xf6, 0xdc, 0x48, 0x88, 0x4d, 0xa9, 0xad, 0x26, 0xbd, 0xed, 0x71, 0xff, 0x40, 0x69, 0xa5,
	0xfd, 0x80, 0xb9, 0xaa, 0xc1, 0xf4, 0x07, 0x4a, 0x6b, 0x95, 0x49, 0xe6, 0xbb, 0x69, 0x75, 0xca,
	0x6c, 0x9a, 0x4b, 0xb0, 0x22, 0x37, 0x6b, 0x40, 0x77, 0x4b, 0x71, 0xf4, 0xbd, 0x85, 0xbb, 0xc7,
	0x3a, 0x99, 0x73, 0x79, 0x56, 0x49, 0xb0, 0x84, 0xe0, 0xb6, 0x11, 0x76, 0x16, 0xa1, 0x21, 0x1a,
	0xef, 0x71, 0x5f, 0x93, 0x23, 0xbc, 0x03, 0x56, 0x94, 0x36, 0x6a, 0x0d, 0xd1, 0xb8, 0x7b, 0xd8,
	0x8f, 0x1b, 0x87, 0x78, 0xe3, 0x10, 0x7f, 0xdc, 0x38, 0x4c, 0x1e, 0x5d, 0xfe, 0x1c, 0x04, 0x17,
	0xbf, 0x06, 0x88, 0x37, 0x14, 0xf2, 0x02, 0x87, 0xb2, 0xf8, 0x1c, 0x85, 0x0f, 0x60, 0x3a, 0x82,
	0xfb, 0x07, 0x58, 0x69, 0xa2, 0xf6, 0x10, 0x8d, 0x43, 0xee, 0x6b, 0x72, 0x80, 0x77, 0xce, 0x2a,
	0x59, 0x9e, 0x47, 0x1d, 0xff, 0xb9, 0xa6, 0x21, 0x6f, 0xf0, 0xbe, 0x8b, 0x23, 0x2d, 0xd4, 0x3b,
	0x63, 0x53, 0x5d, 0x40, 0xb4, 0xeb, 0xcd, 0x1e, 0xc7, 0xdb, 0x61, 0xc5, 0xaf, 0xee, 0x60, 0x26,
	0x6d, 0x67, 0xc7, 0xff, 0x61, 0x8e, 0xbe, 0xe2, 0x5e, 0x13, 0x06, 0x18, 0x5d, 0x80, 0x24, 0xcf,
	0x70, 0x27, 0x13, 0x53, 0x99, 0x41, 0x84, 0x86, 0xe1, 0x7d, 0x4d, 0x87, 0x7d, 0xeb, 0xe7, 0xef,
	0x45, 0x5a, 0xf2, 0x35, 0x96, 0x1c, 0xe2, 0x5d, 0x10, 0xb9, 0xc9, 0x24, 0x44, 0x2d, 0x4f, 0x8b,
	0xee, 0xd3, 0x3e, 0x78, 0x00, 0xdf, 0x00, 0x47, 0x47, 0x78, 0xff, 0xae, 0x9a, 0x4b, 0xa0, 0x10,
	0xb9, 0xdc, 0x6c, 0xc2, 0xd5, 0x2e, 0x81, 0x2f, 0x22, 0xab, 0xa4, 0xdf, 0xc4, 0x1e, 0x6f, 0x9a,
	0xd1, 0x6b, 0x8c, 0x6f, 0x25, 0x6f, 0x31, 0x8e, 0x88, 0xd6, 0x18, 0xf2, 0x04, 0xf7, 0x6e, 0xee,
	0xe0, 0x24, 0x07, 0x2f, 0x10, 0xf2, 0xee, 0xcd, 0xdb, 0x31, 0x4c, 0xca, 0xc5, 0x92, 0x06, 0x57,
	0x4b, 0x1a, 0x5c, 0x2f, 0x29, 0xfa, 0x56, 0x53, 0xf4, 0xa3, 0xa6, 0xe8, 0xb2, 0xa6, 0x68, 0x51,
	0x53, 0xf4, 0xbb, 0xa6, 0xe8, 0x4f, 0x4d, 0x83, 0xeb, 0x9a, 0xa2, 0x8b, 0x15, 0x0d, 0x16, 0x2b,
	0x1a, 0x5c, 0xad, 0x68, 0xf0, 0xe9, 0xa5, 0x4a, 0xed, 0xac, 0x9a, 0xc6, 0x89, 0xce, 0x99, 0x2a,
	0xc5, 0xa9, 0x28, 0x04, 0xcb, 0xf4, 0x3c, 0x65, 0xff, 0x73, 0xe1, 0xd3, 0x8e, 0xbf, 0x84, 0xa7,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x20, 0x73, 0x6a, 0xfb, 0x10, 0x03, 0x00, 0x00,
}

func (this *MockRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockRequest)
	if !ok {
		that2, ok := that.(MockRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Path != that1.Path {
		return false
	}
	if !this.Start.Equal(that1.Start) {
		return false
	}
	if !this.End.Equal(that1.End) {
		return false
	}
	if this.Step != that1.Step {
		return false
	}
	if this.Query != that1.Query {
		return false
	}
	if !this.CachingOptions.Equal(&that1.CachingOptions) {
		return false
	}
	return true
}
func (this *MockResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResponse)
	if !ok {
		that2, ok := that.(MockResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if !this.Labels[i].Equal(that1.Labels[i]) {
			return false
		}
	}
	if len(this.Samples) != len(that1.Samples) {
		return false
	}
	for i := range this.Samples {
		if !this.Samples[i].Equal(that1.Samples[i]) {
			return false
		}
	}
	return true
}
func (this *MockLabelsPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockLabelsPair)
	if !ok {
		that2, ok := that.(MockLabelsPair)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *MockSample) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockSample)
	if !ok {
		that2, ok := that.(MockSample)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.TimestampMs != that1.TimestampMs {
		return false
	}
	return true
}
func (this *MockRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&resultscache.MockRequest{")
	s = append(s, "Path: "+fmt.Sprintf("%#v", this.Path)+",\n")
	s = append(s, "Start: "+fmt.Sprintf("%#v", this.Start)+",\n")
	s = append(s, "End: "+fmt.Sprintf("%#v", this.End)+",\n")
	s = append(s, "Step: "+fmt.Sprintf("%#v", this.Step)+",\n")
	s = append(s, "Query: "+fmt.Sprintf("%#v", this.Query)+",\n")
	s = append(s, "CachingOptions: "+strings.Replace(this.CachingOptions.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *MockResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&resultscache.MockResponse{")
	if this.Labels != nil {
		s = append(s, "Labels: "+fmt.Sprintf("%#v", this.Labels)+",\n")
	}
	if this.Samples != nil {
		s = append(s, "Samples: "+fmt.Sprintf("%#v", this.Samples)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *MockLabelsPair) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&resultscache.MockLabelsPair{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *MockSample) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&resultscache.MockSample{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "TimestampMs: "+fmt.Sprintf("%#v", this.TimestampMs)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTestTypes(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *MockRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MockRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MockRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CachingOptions.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTestTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.Query) > 0 {
		i -= len(m.Query)
		copy(dAtA[i:], m.Query)
		i = encodeVarintTestTypes(dAtA, i, uint64(len(m.Query)))
		i--
		dAtA[i] = 0x32
	}
	if m.Step != 0 {
		i = encodeVarintTestTypes(dAtA, i, uint64(m.Step))
		i--
		dAtA[i] = 0x20
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.End, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.End):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTestTypes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Start, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Start):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintTestTypes(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintTestTypes(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MockResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MockResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MockResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Samples) > 0 {
		for iNdEx := len(m.Samples) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Samples[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTestTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Labels) > 0 {
		for iNdEx := len(m.Labels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Labels[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTestTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MockLabelsPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MockLabelsPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MockLabelsPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTestTypes(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTestTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MockSample) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MockSample) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MockSample) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimestampMs != 0 {
		i = encodeVarintTestTypes(dAtA, i, uint64(m.TimestampMs))
		i--
		dAtA[i] = 0x10
	}
	if m.Value != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func encodeVarintTestTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTestTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MockRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovTestTypes(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Start)
	n += 1 + l + sovTestTypes(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.End)
	n += 1 + l + sovTestTypes(uint64(l))
	if m.Step != 0 {
		n += 1 + sovTestTypes(uint64(m.Step))
	}
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovTestTypes(uint64(l))
	}
	l = m.CachingOptions.Size()
	n += 1 + l + sovTestTypes(uint64(l))
	return n
}

func (m *MockResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.Size()
			n += 1 + l + sovTestTypes(uint64(l))
		}
	}
	if len(m.Samples) > 0 {
		for _, e := range m.Samples {
			l = e.Size()
			n += 1 + l + sovTestTypes(uint64(l))
		}
	}
	return n
}

func (m *MockLabelsPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTestTypes(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTestTypes(uint64(l))
	}
	return n
}

func (m *MockSample) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 9
	}
	if m.TimestampMs != 0 {
		n += 1 + sovTestTypes(uint64(m.TimestampMs))
	}
	return n
}

func sovTestTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTestTypes(x uint64) (n int) {
	return sovTestTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *MockRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MockRequest{`,
		`Path:` + fmt.Sprintf("%v", this.Path) + `,`,
		`Start:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Start), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`End:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.End), "Timestamp", "types.Timestamp", 1), `&`, ``, 1) + `,`,
		`Step:` + fmt.Sprintf("%v", this.Step) + `,`,
		`Query:` + fmt.Sprintf("%v", this.Query) + `,`,
		`CachingOptions:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.CachingOptions), "CachingOptions", "CachingOptions", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MockResponse) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForLabels := "[]*MockLabelsPair{"
	for _, f := range this.Labels {
		repeatedStringForLabels += strings.Replace(f.String(), "MockLabelsPair", "MockLabelsPair", 1) + ","
	}
	repeatedStringForLabels += "}"
	repeatedStringForSamples := "[]*MockSample{"
	for _, f := range this.Samples {
		repeatedStringForSamples += strings.Replace(f.String(), "MockSample", "MockSample", 1) + ","
	}
	repeatedStringForSamples += "}"
	s := strings.Join([]string{`&MockResponse{`,
		`Labels:` + repeatedStringForLabels + `,`,
		`Samples:` + repeatedStringForSamples + `,`,
		`}`,
	}, "")
	return s
}
func (this *MockLabelsPair) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MockLabelsPair{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MockSample) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MockSample{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`TimestampMs:` + fmt.Sprintf("%v", this.TimestampMs) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTestTypes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *MockRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MockRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MockRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTestTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTestTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTestTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTestTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Start, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTestTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTestTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.End, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			m.Step = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Step |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTestTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTestTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CachingOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTestTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTestTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CachingOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTestTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTestTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTestTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MockResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MockResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MockResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTestTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTestTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = append(m.Labels, &MockLabelsPair{})
			if err := m.Labels[len(m.Labels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Samples", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTestTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTestTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Samples = append(m.Samples, &MockSample{})
			if err := m.Samples[len(m.Samples)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTestTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTestTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTestTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MockLabelsPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MockLabelsPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MockLabelsPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTestTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTestTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTestTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTestTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTestTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTestTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTestTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MockSample) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTestTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MockSample: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MockSample: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimestampMs", wireType)
			}
			m.TimestampMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTestTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimestampMs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTestTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTestTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTestTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTestTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTestTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTestTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTestTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTestTypes
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTestTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTestTypes
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTestTypes(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTestTypes
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTestTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTestTypes   = fmt.Errorf("proto: integer overflow")
)