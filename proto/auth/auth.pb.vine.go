// Code generated by proto-gen-vine. DO NOT EDIT.
// source: github.com/lack-io/vine/proto/auth/auth.proto

package auth

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

import (
	context "context"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option
var _ registry.OpenAPI

// API Endpoints for Auth service
func NewAuthEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Swagger OpenAPI 3.0 for Auth service
func NewAuthOpenAPI() *registry.OpenAPI {
	return &registry.OpenAPI{}
}

// Client API for Auth service
type AuthService interface {
	Generate(ctx context.Context, in *GenerateRequest, opts ...client.CallOption) (*GenerateResponse, error)
	Inspect(ctx context.Context, in *InspectRequest, opts ...client.CallOption) (*InspectResponse, error)
	Token(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*TokenResponse, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Generate(ctx context.Context, in *GenerateRequest, opts ...client.CallOption) (*GenerateResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Generate", in)
	out := new(GenerateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Inspect(ctx context.Context, in *InspectRequest, opts ...client.CallOption) (*InspectResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Inspect", in)
	out := new(InspectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Token(ctx context.Context, in *TokenRequest, opts ...client.CallOption) (*TokenResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Token", in)
	out := new(TokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service
type AuthHandler interface {
	Generate(context.Context, *GenerateRequest, *GenerateResponse) error
	Inspect(context.Context, *InspectRequest, *InspectResponse) error
	Token(context.Context, *TokenRequest, *TokenResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type authImpl interface {
		Generate(ctx context.Context, in *GenerateRequest, out *GenerateResponse) error
		Inspect(ctx context.Context, in *InspectRequest, out *InspectResponse) error
		Token(ctx context.Context, in *TokenRequest, out *TokenResponse) error
	}
	type Auth struct {
		authImpl
	}
	h := &authHandler{hdlr}
	opts = append(opts, server.OpenAPIHandler(NewAuthOpenAPI()))
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) Generate(ctx context.Context, in *GenerateRequest, out *GenerateResponse) error {
	return h.AuthHandler.Generate(ctx, in, out)
}

func (h *authHandler) Inspect(ctx context.Context, in *InspectRequest, out *InspectResponse) error {
	return h.AuthHandler.Inspect(ctx, in, out)
}

func (h *authHandler) Token(ctx context.Context, in *TokenRequest, out *TokenResponse) error {
	return h.AuthHandler.Token(ctx, in, out)
}

// API Endpoints for Accounts service
func NewAccountsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Swagger OpenAPI 3.0 for Accounts service
func NewAccountsOpenAPI() *registry.OpenAPI {
	return &registry.OpenAPI{}
}

// Client API for Accounts service
type AccountsService interface {
	List(ctx context.Context, in *ListAccountsRequest, opts ...client.CallOption) (*ListAccountsResponse, error)
}

type accountsService struct {
	c    client.Client
	name string
}

func NewAccountsService(name string, c client.Client) AccountsService {
	return &accountsService{
		c:    c,
		name: name,
	}
}

func (c *accountsService) List(ctx context.Context, in *ListAccountsRequest, opts ...client.CallOption) (*ListAccountsResponse, error) {
	req := c.c.NewRequest(c.name, "Accounts.List", in)
	out := new(ListAccountsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Accounts service
type AccountsHandler interface {
	List(context.Context, *ListAccountsRequest, *ListAccountsResponse) error
}

func RegisterAccountsHandler(s server.Server, hdlr AccountsHandler, opts ...server.HandlerOption) error {
	type accountsImpl interface {
		List(ctx context.Context, in *ListAccountsRequest, out *ListAccountsResponse) error
	}
	type Accounts struct {
		accountsImpl
	}
	h := &accountsHandler{hdlr}
	opts = append(opts, server.OpenAPIHandler(NewAccountsOpenAPI()))
	return s.Handle(s.NewHandler(&Accounts{h}, opts...))
}

type accountsHandler struct {
	AccountsHandler
}

func (h *accountsHandler) List(ctx context.Context, in *ListAccountsRequest, out *ListAccountsResponse) error {
	return h.AccountsHandler.List(ctx, in, out)
}

// API Endpoints for Rules service
func NewRulesEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Swagger OpenAPI 3.0 for Rules service
func NewRulesOpenAPI() *registry.OpenAPI {
	return &registry.OpenAPI{}
}

// Client API for Rules service
type RulesService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
}

type rulesService struct {
	c    client.Client
	name string
}

func NewRulesService(name string, c client.Client) RulesService {
	return &rulesService{
		c:    c,
		name: name,
	}
}

func (c *rulesService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Rules.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Rules.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Rules.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rules service
type RulesHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
}

func RegisterRulesHandler(s server.Server, hdlr RulesHandler, opts ...server.HandlerOption) error {
	type rulesImpl interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
	}
	type Rules struct {
		rulesImpl
	}
	h := &rulesHandler{hdlr}
	opts = append(opts, server.OpenAPIHandler(NewRulesOpenAPI()))
	return s.Handle(s.NewHandler(&Rules{h}, opts...))
}

type rulesHandler struct {
	RulesHandler
}

func (h *rulesHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.RulesHandler.Create(ctx, in, out)
}

func (h *rulesHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.RulesHandler.Delete(ctx, in, out)
}

func (h *rulesHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.RulesHandler.List(ctx, in, out)
}
