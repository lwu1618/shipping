// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv/user/proto/user/user.proto

package go_micro_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (User_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (User_PingPongService, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "User.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (User_StreamService, error) {
	req := c.c.NewRequest(c.name, "User.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &userServiceStream{stream}, nil
}

type User_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type userServiceStream struct {
	stream client.Stream
}

func (x *userServiceStream) Close() error {
	return x.stream.Close()
}

func (x *userServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userService) PingPong(ctx context.Context, opts ...client.CallOption) (User_PingPongService, error) {
	req := c.c.NewRequest(c.name, "User.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &userServicePingPong{stream}, nil
}

type User_PingPongService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type userServicePingPong struct {
	stream client.Stream
}

func (x *userServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *userServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *userServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for User service

type UserHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, User_StreamStream) error
	PingPong(context.Context, User_PingPongStream) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.UserHandler.Call(ctx, in, out)
}

func (h *userHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.UserHandler.Stream(ctx, m, &userStreamStream{stream})
}

type User_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type userStreamStream struct {
	stream server.Stream
}

func (x *userStreamStream) Close() error {
	return x.stream.Close()
}

func (x *userStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *userHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.UserHandler.PingPong(ctx, &userPingPongStream{stream})
}

type User_PingPongStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type userPingPongStream struct {
	stream server.Stream
}

func (x *userPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *userPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *userPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *userPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *userPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
