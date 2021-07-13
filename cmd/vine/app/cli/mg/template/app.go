// MIT License
//
// Copyright (c) 2021 Lack
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package template

var (
	Doc = `// Code generated by vine. DO NOT EDIT.
package runtime

const Namespace = "{{.Toml.Package.Namespace}}"
{{if .Toml.Pkg}}
const {{title .Toml.Pkg.Name}}Name = "{{.Toml.Pkg.Alias}}"
const {{title .Toml.Pkg.Name}}Id = "{{uuid}}"
{{end}}
{{if eq .Toml.Package.Kind "cluster"}}
const ({{range .Toml.Mod}}
	{{title .Name}}Name = "{{.Alias}}"{{end}}
)

var ({{range .Toml.Mod}}
	{{title .Name}}Id = "{{uuid}}"{{end}}
)
{{end}}
var (
	GitTag    string
	GitCommit string
	BuildDate string
	GetVersion = func() string {
		v := GitTag
		if GitCommit != "" {
			v += "-" + GitCommit
		}
		if BuildDate != "" {
			v += "-" + BuildDate
		}
		if v == "" {
			return "latest"
		}
		return v
	}
)
`
	Inject = `package inject

import (
	"github.com/lack-io/pkg/inject"
	log "github.com/lack-io/vine/lib/logger"
)

type logger struct{}

func (l logger) Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

func init() {
	g = inject.Container{}
	g.Logger = logger{}
}

var g inject.Container

func Provide(vv ...interface{}) error {
	for _, v := range vv {
		if err := g.Provide(&inject.Object{
			Value: v,
		}); err != nil {
			return err
		}
	}

	return nil
}

func ProvidePanic(vv ...interface{}) {
	if err := Provide(vv...); err != nil {
		panic(err)
	}
}

func ProvideWithName(v interface{}, name string) error {
	return g.Provide(&inject.Object{Value: v, Name: name})
}

func ProvideWithNamePanic(v interface{}, name string) {
	if err := ProvideWithName(v, name); err != nil {
		panic(err)
	}
}

func PopulateTarget(target interface{}) error {
	return g.PopulateTarget(target)
}

func Populate() error {
	return g.Populate()
}

func Objects() []*inject.Object {
	return g.Objects()
}

func Resolve(dst interface{}) error {
	return g.Resolve(dst)
}

func ResolveByName(dst interface{}, name string) error {
	return g.ResolveByName(dst, name)
}`

	SinglePlugin = `package pkg
{{if .Plugins}}
import ({{range .Plugins}}
	_ "github.com/lack-io/plugins/{{.}}"{{end}}
){{end}}
`

	ClusterPlugin = `package {{.Name}}
{{if .Plugins}}
import ({{range .Plugins}}
	_ "github.com/lack-io/plugins/{{.}}"{{end}}
){{end}}
`

	SingleApp = `package pkg

import (
	"github.com/lack-io/vine"
	log "github.com/lack-io/vine/lib/logger"

	"{{.Dir}}/pkg/runtime"
	"{{.Dir}}/pkg/server"
)

func Run() {
	s := server.New(
		vine.Name(runtime.{{title .Name}}Name),
		vine.Id(runtime.{{title .Name}}Id),
		vine.Version(runtime.GetVersion()),
		vine.Metadata(map[string]string{
			"namespace": runtime.Namespace,
		}),
	)

	if err := s.Init(); err != nil {
		log.Fatal(err)
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}`

	ClusterApp = `package {{.Name}}

import (
	"github.com/lack-io/vine"
	log "github.com/lack-io/vine/lib/logger"

	"{{.Dir}}/pkg/runtime"
	"{{.Dir}}/pkg/{{.Name}}/server"
)

func Run() {
	s := server.New(
		vine.Name(runtime.{{title .Name}}Name),
		vine.Id(runtime.{{title .Name}}Id),
		vine.Version(runtime.GetVersion()),
		vine.Metadata(map[string]string{
			"namespace": runtime.Namespace,
		}),
	)

	if err := s.Init(); err != nil {
		log.Fatal(err)
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}`

	GatewayApp = `package {{.Name}}

import (
	"mime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/lack-io/cli"

	"github.com/lack-io/vine"
	ahandler "github.com/lack-io/vine/lib/api/handler"
	"github.com/lack-io/vine/lib/api/handler/openapi"
	arpc "github.com/lack-io/vine/lib/api/handler/rpc"
	"github.com/lack-io/vine/lib/api/resolver"
	"github.com/lack-io/vine/lib/api/resolver/grpc"
	"github.com/lack-io/vine/lib/api/router"
	regRouter "github.com/lack-io/vine/lib/api/router/registry"
	"github.com/lack-io/vine/lib/api/server"
	httpapi "github.com/lack-io/vine/lib/api/server/http"
	log "github.com/lack-io/vine/lib/logger"
	"github.com/lack-io/vine/util/helper"
	"github.com/lack-io/vine/util/namespace"
	"github.com/rakyll/statik/fs"

	"{{.Dir}}/pkg/runtime"

	_ "github.com/lack-io/vine/lib/api/handler/openapi/statik"
)

var (
	Address       = ":8080"
	Handler       = "rpc"
	Type          = "api"
	APIPath       = "/"
	enableOpenAPI = false

	flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "api-address",
			Usage:       "The specify for api address",
			EnvVars:     []string{"VINE_API_ADDRESS"},
			Required:    true,
			Value:       Address,
			Destination: &Address,
		},
		&cli.BoolFlag{
			Name:    "enable-openapi",
			Usage:   "Enable OpenAPI3",
			EnvVars: []string{"VINE_ENABLE_OPENAPI"},
			Value:   true,
		},
		&cli.BoolFlag{
			Name:    "enable-cors",
			Usage:   "Enable CORS, allowing the API to be called by frontend applications",
			EnvVars: []string{"VINE_API_ENABLE_CORS"},
			Value:   true,
		},
	}
)

func Run() {
	// Init API
	var opts []server.Option

	// initialise service
	svc := vine.NewService(
		vine.Name(runtime.{{title .Name}}Name),
		vine.Id(runtime.{{title .Name}}Id),
		vine.Version(runtime.GetVersion()),
		vine.Metadata(map[string]string{
			"api-address": Address,
			"namespace": runtime.Namespace,
		}),
		vine.Flags(flags...),
		vine.Action(func(ctx *cli.Context) error {
			if len(ctx.String("server-address")) > 0 {
				Address = ctx.String("server-address")
			}
			enableOpenAPI = ctx.Bool("enable-openapi")

			if ctx.Bool("enable-tls") {
				config, err := helper.TLSConfig(ctx)
				if err != nil {
					log.Errorf(err.Error())
					return err
				}

				opts = append(opts, server.EnableTLS(true))
				opts = append(opts, server.TLSConfig(config))
			}
			return nil
		}),
	)

	svc.Init()

	opts = append(opts, server.EnableCORS(true))

	// create the router
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	if enableOpenAPI {
		openAPI := openapi.New(svc)
		_ = mime.AddExtensionType(".svg", "image/svg+xml")
		sfs, err := fs.New()
		if err != nil {
			log.Fatalf("Starting OpenAPI: %v", err)
		}
		prefix := "/openapi-ui/"
		app.All(prefix, openAPI.OpenAPIHandler)
		app.Use(prefix, filesystem.New(filesystem.Config{Root: sfs}))
		app.Get("/openapi.json", openAPI.OpenAPIJOSNHandler)
		app.Get("/services", openAPI.OpenAPIServiceHandler)
		log.Infof("Starting OpenAPI at %v", prefix)
	}

	// create the namespace resolver
	nsResolver := namespace.NewResolver(Type, runtime.Namespace)
	// resolver options
	ropts := []resolver.Option{
		resolver.WithNamespace(nsResolver.ResolveWithType),
		resolver.WithHandler(Handler),
	}

	log.Infof("Registering API RPC Handler at %s", APIPath)
	rr := grpc.NewResolver(ropts...)
	rt := regRouter.NewRouter(
		router.WithHandler(arpc.Handler),
		router.WithResolver(rr),
		router.WithRegistry(svc.Options().Registry),
	)
	rp := arpc.NewHandler(
		ahandler.WithNamespace(runtime.Namespace),
		ahandler.WithRouter(rt),
		ahandler.WithClient(svc.Client()),
	)
	app.Group(APIPath, rp.Handle)

	api := httpapi.NewServer(Address)

	if err := api.Init(opts...); err != nil {
		log.Fatal(err)
    }
	api.Handle("/", app)

	// Start API
	if err := api.Start(); err != nil {
		log.Fatal(err)
	}

	// Run server
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}

	// Stop API
	if err := api.Stop(); err != nil {
		log.Fatal(err)
	}
}
`

	SingleWebSRV = `package pkg

import (
	"github.com/gofiber/fiber/v2"

	log "github.com/lack-io/vine/lib/logger"
	"github.com/lack-io/vine/lib/web"

	"{{.Dir}}/pkg/runtime"
)

func Run() {
	s := web.NewService(
		web.Name(runtime.{{title .Name}}Name),
		web.Id(runtime.{{title .Name}}Id),
		web.Metadata(map[string]string{
			"namespace": runtime.Namespace,
		}),
	)

	s.Handle(web.MethodGet, "/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	if err := s.Init(); err != nil {
		log.Fatal(err)
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}`

	ClusterWebSRV = `package {{.Name}}

import (
	"github.com/gofiber/fiber/v2"

	log "github.com/lack-io/vine/lib/logger"
	"github.com/lack-io/vine/lib/web"

	"{{.Dir}}/pkg/runtime"
)

func Run() {
	s := web.NewService(
		web.Name(runtime.{{title .Name}}Name),
		web.Id(runtime.{{title .Name}}Id),
		web.Metadata(map[string]string{
			"namespace": runtime.Namespace,
		}),
	)

	s.Handle(web.MethodGet, "/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	if err := s.Init(); err != nil {
		log.Fatal(err)
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}`

	DaoHandler = `package dao

import (
	"{{.Dir}}/pkg/runtime/inject"
	"github.com/lack-io/vine/util/runtime"
)

func init() {
	_ = inject.Provide(sets)
}

var sets = runtime.NewSchemaSet()
`
)
