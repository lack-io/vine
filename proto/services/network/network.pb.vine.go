// Code generated by proto-gen-vine. DO NOT EDIT.
// source: github.com/lack-io/vine/proto/services/network/network.proto

package network

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/lack-io/vine/proto/services/router"
	math "math"
)

import (
	context "context"
	apipb "github.com/lack-io/vine/proto/apis/api"
	registry "github.com/lack-io/vine/proto/apis/registry"
	api "github.com/lack-io/vine/service/api"
	client "github.com/lack-io/vine/service/client"
	server "github.com/lack-io/vine/service/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ apipb.Endpoint
var _ api.Option
var _ context.Context
var _ client.Option
var _ server.Option
var _ registry.OpenAPI

// API Endpoints for Network service
func NewNetworkEndpoints() []*apipb.Endpoint {
	return []*apipb.Endpoint{}
}

// Client API for Network service
// Network service is usesd to gain visibility into networks
type NetworkService interface {
	// Connect to the network
	Connect(ctx context.Context, in *ConnectRequest, opts ...client.CallOption) (*ConnectResponse, error)
	// Returns the entire network graph
	Graph(ctx context.Context, in *GraphRequest, opts ...client.CallOption) (*GraphResponse, error)
	// Returns a list of known nodes in the network
	Nodes(ctx context.Context, in *NodesRequest, opts ...client.CallOption) (*NodesResponse, error)
	// Returns a list of known routes in the network
	Routes(ctx context.Context, in *RoutesRequest, opts ...client.CallOption) (*RoutesResponse, error)
	// Returns a list of known services based on routes
	Services(ctx context.Context, in *ServicesRequest, opts ...client.CallOption) (*ServicesResponse, error)
	// Status returns network status
	Status(ctx context.Context, in *StatusRequest, opts ...client.CallOption) (*StatusResponse, error)
}

type networkService struct {
	c    client.Client
	name string
}

func NewNetworkService(name string, c client.Client) NetworkService {
	return &networkService{
		c:    c,
		name: name,
	}
}

func (c *networkService) Connect(ctx context.Context, in *ConnectRequest, opts ...client.CallOption) (*ConnectResponse, error) {
	req := c.c.NewRequest(c.name, "Network.Connect", in)
	out := new(ConnectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkService) Graph(ctx context.Context, in *GraphRequest, opts ...client.CallOption) (*GraphResponse, error) {
	req := c.c.NewRequest(c.name, "Network.Graph", in)
	out := new(GraphResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkService) Nodes(ctx context.Context, in *NodesRequest, opts ...client.CallOption) (*NodesResponse, error) {
	req := c.c.NewRequest(c.name, "Network.Nodes", in)
	out := new(NodesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkService) Routes(ctx context.Context, in *RoutesRequest, opts ...client.CallOption) (*RoutesResponse, error) {
	req := c.c.NewRequest(c.name, "Network.Routes", in)
	out := new(RoutesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkService) Services(ctx context.Context, in *ServicesRequest, opts ...client.CallOption) (*ServicesResponse, error) {
	req := c.c.NewRequest(c.name, "Network.Services", in)
	out := new(ServicesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkService) Status(ctx context.Context, in *StatusRequest, opts ...client.CallOption) (*StatusResponse, error) {
	req := c.c.NewRequest(c.name, "Network.Status", in)
	out := new(StatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Network service
// Network service is usesd to gain visibility into networks
type NetworkHandler interface {
	// Connect to the network
	Connect(context.Context, *ConnectRequest, *ConnectResponse) error
	// Returns the entire network graph
	Graph(context.Context, *GraphRequest, *GraphResponse) error
	// Returns a list of known nodes in the network
	Nodes(context.Context, *NodesRequest, *NodesResponse) error
	// Returns a list of known routes in the network
	Routes(context.Context, *RoutesRequest, *RoutesResponse) error
	// Returns a list of known services based on routes
	Services(context.Context, *ServicesRequest, *ServicesResponse) error
	// Status returns network status
	Status(context.Context, *StatusRequest, *StatusResponse) error
}

func RegisterNetworkHandler(s server.Server, hdlr NetworkHandler, opts ...server.HandlerOption) error {
	type networkImpl interface {
		Connect(ctx context.Context, in *ConnectRequest, out *ConnectResponse) error
		Graph(ctx context.Context, in *GraphRequest, out *GraphResponse) error
		Nodes(ctx context.Context, in *NodesRequest, out *NodesResponse) error
		Routes(ctx context.Context, in *RoutesRequest, out *RoutesResponse) error
		Services(ctx context.Context, in *ServicesRequest, out *ServicesResponse) error
		Status(ctx context.Context, in *StatusRequest, out *StatusResponse) error
	}
	type Network struct {
		networkImpl
	}
	h := &networkHandler{hdlr}
	return s.Handle(s.NewHandler(&Network{h}, opts...))
}

type networkHandler struct {
	NetworkHandler
}

func (h *networkHandler) Connect(ctx context.Context, in *ConnectRequest, out *ConnectResponse) error {
	return h.NetworkHandler.Connect(ctx, in, out)
}

func (h *networkHandler) Graph(ctx context.Context, in *GraphRequest, out *GraphResponse) error {
	return h.NetworkHandler.Graph(ctx, in, out)
}

func (h *networkHandler) Nodes(ctx context.Context, in *NodesRequest, out *NodesResponse) error {
	return h.NetworkHandler.Nodes(ctx, in, out)
}

func (h *networkHandler) Routes(ctx context.Context, in *RoutesRequest, out *RoutesResponse) error {
	return h.NetworkHandler.Routes(ctx, in, out)
}

func (h *networkHandler) Services(ctx context.Context, in *ServicesRequest, out *ServicesResponse) error {
	return h.NetworkHandler.Services(ctx, in, out)
}

func (h *networkHandler) Status(ctx context.Context, in *StatusRequest, out *StatusResponse) error {
	return h.NetworkHandler.Status(ctx, in, out)
}
