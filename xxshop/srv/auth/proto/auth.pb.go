
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"


import (
  client "github.com/micro/go-micro/client"
  server "github.com/micro/go-micro/server"
  context "golang.org/x/net/context"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf


const _ = proto.ProtoPackageIsVersion2


//对应 proto 文件的 Request 参数
//----------------------------------------------------
type Request struct {
  AuthToken string `protobuf:"bytes,1,opt,name=authToken" json:"authToken,omitempty"`
}


//下面这个其实是一些针对参数的常规的， protobuf要用到的方法
func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }
//------------------------------------------------------



//对应 proto 文件的 Result 参数
//------------------------------------------------------
type Result struct {
  Customer *Customer `protobuf:"bytes,1,opt,name=customer" json:"customer,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }
//----------------------------------------------------



//对应 proto 文件的  Customer 参数
type Customer struct {
  Id        int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
  AuthToken string `protobuf:"bytes,2,opt,name=authToken" json:"authToken,omitempty"`
}


func (m *Result) GetCustomer() *Customer {
  if m != nil {
    return m.Customer
  }
  return nil
}


func (m *Customer) Reset()                    { *m = Customer{} }
func (m *Customer) String() string            { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()               {}
func (*Customer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

//----参数的设置完毕---------------------------------------


//---- 开始初始化程序 -----------------------
func init() {
  proto.RegisterType((*Request)(nil), "auth.Request")
  proto.RegisterType((*Result)(nil), "auth.Result")
  proto.RegisterType((*Customer)(nil), "auth.Customer")
}

// Reference imports to suppress errors if they are not otherwise used.
//屏蔽一些变量错误， 假如程序没有用到的话
var _ context.Context
var _ client.Option
var _ server.Option



// Client API for Auth service
//---------------------------------------------------------------

//写service 的声明了
type AuthClient interface {
  VerifyToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Result, error)
}


type authClient struct {
  c                   client.Client
  serviceName         string
}

func NewAuthClient(serviceName string, c client.Client) AuthClient {
  if c == nil {
    c = client.NewClient()
  }
  if len(serviceName) == 0 {
    serviceName = "auth"
  }
  return &authClient{
    c:           c,
    serviceName: serviceName,
  }
}

func (c *authClient) VerifyToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Result, error) {
  req := c.c.NewRequest(c.serviceName, "Auth.VerifyToken", in)
  out := new(Result)
  err := c.c.Call(ctx, req, out, opts...)
  if err != nil {
    return nil, err
  }
  return out, nil
}


//---------------------------------------------------------------



// Server API for Auth service
// ---------------------------------------------------------------
type AuthHandler interface {
  VerifyToken(context.Context, *Request, *Result) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) {
  s.Handle(s.NewHandler(&Auth{hdlr}, opts...))
}

type Auth struct {
  AuthHandler
}

func (h *Auth) VerifyToken(ctx context.Context, in *Request, out *Result) error {
  return h.AuthHandler.VerifyToken(ctx, in, out)
}

func init() { proto.RegisterFile("srv/auth/proto/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
  // 179 bytes of a gzipped FileDescriptorProto
  0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x2e, 0x2a, 0xd3,
  0x4f, 0x2c, 0x2d, 0xc9, 0xd0, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x07, 0x33, 0xf5, 0xc0, 0x4c, 0x21,
  0x16, 0x10, 0x5b, 0x49, 0x9d, 0x8b, 0x3d, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x86,
  0x8b, 0x13, 0x24, 0x14, 0x92, 0x9f, 0x9d, 0x9a, 0x27, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84,
  0x10, 0x50, 0x32, 0xe1, 0x62, 0x0b, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x11, 0xd2, 0xe2, 0xe2, 0x48,
  0x2e, 0x2d, 0x2e, 0xc9, 0xcf, 0x4d, 0x2d, 0x02, 0x2b, 0xe3, 0x36, 0xe2, 0xd3, 0x03, 0x9b, 0xeb,
  0x0c, 0x15, 0x0d, 0x82, 0xcb, 0x2b, 0x59, 0x70, 0x71, 0xc0, 0x44, 0x85, 0xf8, 0xb8, 0x98, 0x32,
  0x53, 0xc0, 0x3a, 0x58, 0x83, 0x98, 0x32, 0x53, 0x50, 0xed, 0x63, 0x42, 0xb3, 0xcf, 0xc8, 0x84,
  0x8b, 0xc5, 0xb1, 0xb4, 0x24, 0x43, 0x48, 0x87, 0x8b, 0x3b, 0x2c, 0xb5, 0x28, 0x33, 0xad, 0x12,
  0x2c, 0x2c, 0xc4, 0x0b, 0xb1, 0x0a, 0xea, 0x66, 0x29, 0x1e, 0x18, 0x17, 0xe4, 0x32, 0x25, 0x86,
  0x24, 0x36, 0xb0, 0xdf, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x92, 0x53, 0x1b, 0x1d, 0xf8,
  0x00, 0x00, 0x00,
}

//--------------------------------------------------




