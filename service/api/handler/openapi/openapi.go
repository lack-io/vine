// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openapi

import (
	"html/template"
	"net/http"
	"strings"

	json "github.com/json-iterator/go"

	"github.com/lack-io/vine"
	openapipb "github.com/lack-io/vine/proto/apis/openapi"
	maddr "github.com/lack-io/vine/util/addr"
)

type openAPI struct {
	svc    vine.Service
	prefix string
}

func (o *openAPI) OpenAPIHandler(w http.ResponseWriter, r *http.Request) {
	if ct := r.Header.Get("Content-Type"); ct == "application/json" {
		w.Header().Set("Content-Type", ct)
		//w.Write(b)
		return
	}
	var tmpl string
	style := r.URL.Query().Get("style")
	switch style {
	case "redoc":
		tmpl = redocTmpl
	default:
		tmpl = swaggerTmpl
	}

	render(w, r, tmpl, nil)
}

func (o *openAPI) OpenAPIJOSNHandler(w http.ResponseWriter, r *http.Request) {
	services, err := o.svc.Options().Registry.ListServices()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	tags := make(map[string]*openapipb.OpenAPITag, 0)
	paths := make(map[string]*openapipb.OpenAPIPath, 0)
	schemas := make(map[string]*openapipb.Model, 0)
	security := &openapipb.SecuritySchemes{}
	servers := make([]*openapipb.OpenAPIServer, 0)
	for _, item := range services {
		list, err := o.svc.Options().Registry.GetService(item.Name)
		if err != nil {
			continue
		}
		if item.Name == "go.vine.api" {
			for _, node := range item.Nodes {
				if v, ok := node.Metadata["api-address"]; ok {
					if strings.HasPrefix(v, ":") {
						for _, ip := range maddr.IPv4s() {
							if ip == "localhost" || ip == "127.0.0.1" {
								continue
							}
							v = ip + v
						}
					}
					if !strings.HasPrefix(v, "http://") || !strings.HasPrefix(v, "https://") {
						v = "http://" + v
					}
					servers = append(servers, &openapipb.OpenAPIServer{
						Url:         v,
						Description: item.Name,
					})
				}
			}
		}
		for _, i := range list {
			if len(i.Apis) == 0 {
				continue
			}
			for _, api := range i.Apis {
				if api == nil || api.Components.SecuritySchemes == nil {
					continue
				}
				for _, tag := range api.Tags {
					tags[tag.Name] = tag
				}
				for name, path := range api.Paths {
					paths[name] = path
				}
				for name, schema := range api.Components.Schemas {
					schemas[name] = schema
				}
				if api.Components.SecuritySchemes.Basic != nil {
					security.Basic = api.Components.SecuritySchemes.Basic
				}
				if api.Components.SecuritySchemes.Bearer != nil {
					security.Bearer = api.Components.SecuritySchemes.Bearer
				}
				if api.Components.SecuritySchemes.ApiKeys != nil {
					security.ApiKeys = api.Components.SecuritySchemes.ApiKeys
				}
			}
		}
	}
	openapi := &openapipb.OpenAPI{
		Openapi: "3.0.1",
		Info: &openapipb.OpenAPIInfo{
			Title:       "Vine Document",
			Description: "OpenAPI3.0",
		},
		Tags:    []*openapipb.OpenAPITag{},
		Paths:   paths,
		Servers: servers,
		Components: &openapipb.OpenAPIComponents{
			SecuritySchemes: security,
			Schemas:         schemas,
		},
	}
	for _, tag := range tags {
		openapi.Tags = append(openapi.Tags, tag)
	}
	v, _ := json.Marshal(openapi)
	w.Write(v)
	w.WriteHeader(200)
}

func (o *openAPI) ServeHTTP(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}

func New(svc vine.Service) *openAPI {
	return &openAPI{svc: svc}
}

func render(w http.ResponseWriter, r *http.Request, tmpl string, data interface{}) {
	t, err := template.New("template").Funcs(template.FuncMap{
		//		"format": format,
	}).Parse(layoutTemplate)
	if err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
		return
	}
	t, err = t.Parse(tmpl)
	if err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
		return
	}
	if err := t.ExecuteTemplate(w, "layout", data); err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
	}
}