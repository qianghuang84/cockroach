// Code generated by protoc-gen-gogo.
// source: echo.proto
// DO NOT EDIT!

package message

import proto "code.google.com/p/gogoprotobuf/proto"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type EchoRequest struct {
	Msg              string `protobuf:"bytes,1,opt,name=msg" json:"msg"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}

func (m *EchoRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type EchoResponse struct {
	Msg              string `protobuf:"bytes,1,opt,name=msg" json:"msg"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}

func (m *EchoResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
}
