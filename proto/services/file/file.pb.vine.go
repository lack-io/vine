// Code generated by proto-gen-vine. DO NOT EDIT.
// source: github.com/lack-io/vine/proto/services/file/file.proto

package file

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	apipb "github.com/lack-io/vine/proto/apis/api"
	client "github.com/lack-io/vine/service/client"
	server "github.com/lack-io/vine/service/server"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// API Endpoints for File service
func NewFileEndpoints() []*apipb.Endpoint {
	return []*apipb.Endpoint{}
}

// Client API for File service
type FileService interface {
	Open(ctx context.Context, in *OpenRequest, opts ...client.CallOption) (*OpenResponse, error)
	Stat(ctx context.Context, in *StatRequest, opts ...client.CallOption) (*StatResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...client.CallOption) (*WriteResponse, error)
	Close(ctx context.Context, in *CloseRequest, opts ...client.CallOption) (*CloseResponse, error)
}

type fileService struct {
	c    client.Client
	name string
}

func NewFileService(name string, c client.Client) FileService {
	return &fileService{
		c:    c,
		name: name,
	}
}

func (c *fileService) Open(ctx context.Context, in *OpenRequest, opts ...client.CallOption) (*OpenResponse, error) {
	req := c.c.NewRequest(c.name, "File.Open", in)
	out := new(OpenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) Stat(ctx context.Context, in *StatRequest, opts ...client.CallOption) (*StatResponse, error) {
	req := c.c.NewRequest(c.name, "File.Stat", in)
	out := new(StatResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "File.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) Write(ctx context.Context, in *WriteRequest, opts ...client.CallOption) (*WriteResponse, error) {
	req := c.c.NewRequest(c.name, "File.Write", in)
	out := new(WriteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileService) Close(ctx context.Context, in *CloseRequest, opts ...client.CallOption) (*CloseResponse, error) {
	req := c.c.NewRequest(c.name, "File.Close", in)
	out := new(CloseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for File service
type FileHandler interface {
	Open(context.Context, *OpenRequest, *OpenResponse) error
	Stat(context.Context, *StatRequest, *StatResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Write(context.Context, *WriteRequest, *WriteResponse) error
	Close(context.Context, *CloseRequest, *CloseResponse) error
}

func RegisterFileHandler(s server.Server, hdlr FileHandler, opts ...server.HandlerOption) error {
	type fileImpl interface {
		Open(ctx context.Context, in *OpenRequest, out *OpenResponse) error
		Stat(ctx context.Context, in *StatRequest, out *StatResponse) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Write(ctx context.Context, in *WriteRequest, out *WriteResponse) error
		Close(ctx context.Context, in *CloseRequest, out *CloseResponse) error
	}
	type File struct {
		fileImpl
	}
	h := &fileHandler{hdlr}
	return s.Handle(s.NewHandler(&File{h}, opts...))
}

type fileHandler struct {
	FileHandler
}

func (h *fileHandler) Open(ctx context.Context, in *OpenRequest, out *OpenResponse) error {
	return h.FileHandler.Open(ctx, in, out)
}

func (h *fileHandler) Stat(ctx context.Context, in *StatRequest, out *StatResponse) error {
	return h.FileHandler.Stat(ctx, in, out)
}

func (h *fileHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.FileHandler.Read(ctx, in, out)
}

func (h *fileHandler) Write(ctx context.Context, in *WriteRequest, out *WriteResponse) error {
	return h.FileHandler.Write(ctx, in, out)
}

func (h *fileHandler) Close(ctx context.Context, in *CloseRequest, out *CloseResponse) error {
	return h.FileHandler.Close(ctx, in, out)
}
