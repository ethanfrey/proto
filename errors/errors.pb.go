// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errors/errors.proto

/*
Package options is a generated protocol buffer package.

It is generated from these files:
	errors/errors.proto

It has these top-level messages:
	Error
	Response
	ResponseInfo
	ResponseDeliverTx
	KVPair
*/
package options

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Error struct {
	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Log  string `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptorErrors, []int{0} }

func (m *Error) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

type Response struct {
	Error *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	// Types that are valid to be assigned to Result:
	//	*Response_Info
	//	*Response_DeliverTx
	Result isResponse_Result `protobuf_oneof:"result"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptorErrors, []int{1} }

type isResponse_Result interface {
	isResponse_Result()
}

type Response_Info struct {
	Info *ResponseInfo `protobuf:"bytes,2,opt,name=info,oneof"`
}
type Response_DeliverTx struct {
	DeliverTx *ResponseDeliverTx `protobuf:"bytes,3,opt,name=deliver_tx,json=deliverTx,oneof"`
}

func (*Response_Info) isResponse_Result()      {}
func (*Response_DeliverTx) isResponse_Result() {}

func (m *Response) GetResult() isResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetInfo() *ResponseInfo {
	if x, ok := m.GetResult().(*Response_Info); ok {
		return x.Info
	}
	return nil
}

func (m *Response) GetDeliverTx() *ResponseDeliverTx {
	if x, ok := m.GetResult().(*Response_DeliverTx); ok {
		return x.DeliverTx
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Response) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Response_OneofMarshaler, _Response_OneofUnmarshaler, _Response_OneofSizer, []interface{}{
		(*Response_Info)(nil),
		(*Response_DeliverTx)(nil),
	}
}

func _Response_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Response)
	// result
	switch x := m.Result.(type) {
	case *Response_Info:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Info); err != nil {
			return err
		}
	case *Response_DeliverTx:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DeliverTx); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Response.Result has unexpected type %T", x)
	}
	return nil
}

func _Response_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Response)
	switch tag {
	case 2: // result.info
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ResponseInfo)
		err := b.DecodeMessage(msg)
		m.Result = &Response_Info{msg}
		return true, err
	case 3: // result.deliver_tx
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ResponseDeliverTx)
		err := b.DecodeMessage(msg)
		m.Result = &Response_DeliverTx{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Response_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Response)
	// result
	switch x := m.Result.(type) {
	case *Response_Info:
		s := proto.Size(x.Info)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Response_DeliverTx:
		s := proto.Size(x.DeliverTx)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ResponseInfo struct {
	Data             string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Version          string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	LastBlockHeight  int64  `protobuf:"varint,3,opt,name=last_block_height,json=lastBlockHeight,proto3" json:"last_block_height,omitempty"`
	LastBlockAppHash []byte `protobuf:"bytes,4,opt,name=last_block_app_hash,json=lastBlockAppHash,proto3" json:"last_block_app_hash,omitempty"`
}

func (m *ResponseInfo) Reset()                    { *m = ResponseInfo{} }
func (m *ResponseInfo) String() string            { return proto.CompactTextString(m) }
func (*ResponseInfo) ProtoMessage()               {}
func (*ResponseInfo) Descriptor() ([]byte, []int) { return fileDescriptorErrors, []int{2} }

func (m *ResponseInfo) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ResponseInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ResponseInfo) GetLastBlockHeight() int64 {
	if m != nil {
		return m.LastBlockHeight
	}
	return 0
}

func (m *ResponseInfo) GetLastBlockAppHash() []byte {
	if m != nil {
		return m.LastBlockAppHash
	}
	return nil
}

type ResponseDeliverTx struct {
	Data []byte    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Tags []*KVPair `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty"`
}

func (m *ResponseDeliverTx) Reset()                    { *m = ResponseDeliverTx{} }
func (m *ResponseDeliverTx) String() string            { return proto.CompactTextString(m) }
func (*ResponseDeliverTx) ProtoMessage()               {}
func (*ResponseDeliverTx) Descriptor() ([]byte, []int) { return fileDescriptorErrors, []int{3} }

func (m *ResponseDeliverTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ResponseDeliverTx) GetTags() []*KVPair {
	if m != nil {
		return m.Tags
	}
	return nil
}

type KVPair struct {
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KVPair) Reset()                    { *m = KVPair{} }
func (m *KVPair) String() string            { return proto.CompactTextString(m) }
func (*KVPair) ProtoMessage()               {}
func (*KVPair) Descriptor() ([]byte, []int) { return fileDescriptorErrors, []int{4} }

func (m *KVPair) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KVPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "options.Error")
	proto.RegisterType((*Response)(nil), "options.Response")
	proto.RegisterType((*ResponseInfo)(nil), "options.ResponseInfo")
	proto.RegisterType((*ResponseDeliverTx)(nil), "options.ResponseDeliverTx")
	proto.RegisterType((*KVPair)(nil), "options.KVPair")
}

func init() { proto.RegisterFile("errors/errors.proto", fileDescriptorErrors) }

var fileDescriptorErrors = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x6d, 0x9a, 0xf4, 0xef, 0xb6, 0xdf, 0xd7, 0x76, 0xaa, 0x10, 0x5c, 0x85, 0xe8, 0x22, 0x28,
	0xad, 0x10, 0x97, 0xae, 0x2c, 0x0a, 0x11, 0x5d, 0xc8, 0x20, 0x6e, 0xc3, 0xb4, 0x99, 0x36, 0xc1,
	0x98, 0x19, 0x66, 0xa6, 0xa5, 0x6f, 0xe2, 0x0b, 0xf8, 0xa0, 0x32, 0x93, 0x1f, 0x2a, 0x5d, 0xe5,
	0x9e, 0x7b, 0xce, 0x3d, 0x9c, 0x13, 0x06, 0x66, 0x54, 0x08, 0x26, 0xe4, 0x6d, 0xf9, 0x59, 0x70,
	0xc1, 0x14, 0x43, 0x3d, 0xc6, 0x55, 0xc6, 0x0a, 0xe9, 0xcf, 0xa1, 0xf3, 0xa4, 0x09, 0x84, 0xc0,
	0x59, 0xb3, 0x84, 0xba, 0x96, 0x67, 0x05, 0xff, 0xb0, 0x99, 0xd1, 0x04, 0xec, 0x9c, 0x6d, 0xdd,
	0xb6, 0x67, 0x05, 0x03, 0xac, 0x47, 0xff, 0xc7, 0x82, 0x3e, 0xa6, 0x92, 0xb3, 0x42, 0x52, 0x74,
	0x05, 0x1d, 0x63, 0x6a, 0x6e, 0x86, 0xe1, 0xff, 0x45, 0x65, 0xba, 0x30, 0x8e, 0xb8, 0x24, 0xd1,
	0x0d, 0x38, 0x59, 0xb1, 0x61, 0xc6, 0x65, 0x18, 0x9e, 0x37, 0xa2, 0xda, 0xe6, 0xb9, 0xd8, 0xb0,
	0xa8, 0x85, 0x8d, 0x08, 0xdd, 0x03, 0x24, 0x34, 0xcf, 0xf6, 0x54, 0xc4, 0xea, 0xe0, 0xda, 0xe6,
	0xe4, 0xe2, 0xe4, 0xe4, 0xb1, 0x94, 0xbc, 0x1f, 0xa2, 0x16, 0x1e, 0x24, 0x35, 0x58, 0xf6, 0xa1,
	0x2b, 0xa8, 0xdc, 0xe5, 0xca, 0xff, 0xb6, 0x60, 0x74, 0xec, 0xaf, 0xdb, 0x25, 0x44, 0x11, 0x93,
	0x74, 0x80, 0xcd, 0x8c, 0x5c, 0xe8, 0xed, 0xa9, 0x90, 0x19, 0x2b, 0xaa, 0x86, 0x35, 0x44, 0xd7,
	0x30, 0xcd, 0x89, 0x54, 0xf1, 0x2a, 0x67, 0xeb, 0xcf, 0x38, 0xa5, 0xd9, 0x36, 0x55, 0x26, 0x8c,
	0x8d, 0xc7, 0x9a, 0x58, 0xea, 0x7d, 0x64, 0xd6, 0x68, 0x0e, 0xb3, 0x23, 0x2d, 0xe1, 0x3c, 0x4e,
	0x89, 0x4c, 0x5d, 0xc7, 0xb3, 0x82, 0x11, 0x9e, 0x34, 0xea, 0x07, 0xce, 0x23, 0x22, 0x53, 0xff,
	0x15, 0xa6, 0x27, 0x2d, 0xfe, 0xa4, 0x1b, 0x55, 0xe9, 0x2e, 0xc1, 0x51, 0x64, 0x2b, 0xdd, 0xb6,
	0x67, 0x07, 0xc3, 0x70, 0xdc, 0xfc, 0x83, 0x97, 0x8f, 0x37, 0x92, 0x09, 0x6c, 0x48, 0x3f, 0x84,
	0x6e, 0x89, 0xb5, 0x45, 0x41, 0xbe, 0x68, 0x5d, 0x50, 0xcf, 0xe8, 0x0c, 0x3a, 0x7b, 0x92, 0xef,
	0x68, 0x55, 0xaf, 0x04, 0xab, 0xae, 0x79, 0x01, 0x77, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x98,
	0x90, 0xdc, 0x2b, 0x18, 0x02, 0x00, 0x00,
}
