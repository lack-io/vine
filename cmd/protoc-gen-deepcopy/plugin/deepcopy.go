// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

	"github.com/lack-io/vine/cmd/generator"
)

var TagString = "gen"

const (
	// message tag
	_deepcopy = "deepcopy"

	// field common tag
	_inline   = "inline"
)

type Tag struct {
	Key   string
	Value string
}

// deepcopy is an implementation of the Go protocol buffer compiler's
// plugin architecture. It generates bindings for deepcopy support.
type deepcopy struct {
	generator.PluginImports
	gen *generator.Generator
}

func New() *deepcopy {
	return &deepcopy{}
}

// Name returns the name of this plugin, "deepcopy".
func (g *deepcopy) Name() string {
	return "deepcopy"
}

// Init initializes the plugin.
func (g *deepcopy) Init(gen *generator.Generator) {
	g.gen = gen
	g.PluginImports = generator.NewPluginImports(g.gen)
}

// Given a type name defined in a .proto, return its object.
// Also record that we're using it, to guarantee the associated import.
func (g *deepcopy) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

// Given a type name defined in a .proto, return its name as we will print it.
func (g *deepcopy) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}

// P forwards to g.gen.P.
func (g *deepcopy) P(args ...interface{}) { g.gen.P(args...) }

// Generate generates code for the services in the given file.
func (g *deepcopy) Generate(file *generator.FileDescriptor) {
	if len(file.Comments()) == 0 {
		return
	}

	for _, msg := range file.Messages() {
		g.generateMessage(file, msg)
	}
}

func (g *deepcopy) generateMessage(file *generator.FileDescriptor, msg *generator.MessageDescriptor) {
	if msg.Proto.File() != file {
		return
	}
	if msg.Proto.Options != nil && msg.Proto.Options.GetMapEntry() {
		return
	}

	mname:= msg.Proto.GetName()
	g.P(fmt.Sprintf(`// DeepCopyInto is an auto-generated deepcopy function, coping the receiver, writing into out. in must be no-nil.`))
	g.P(fmt.Sprintf(`func (in *%s) DeepCopyInto(out *%s) {`, mname, mname))
	g.P(`*out = *in`)
	for _, field := range msg.Fields {
		if field.Proto.IsRepeated() {
			g.generateRepeatedField(file, msg, field)
			continue
		}
		if field.Proto.IsMessage() {
			g.generateMessageField(field)
			continue
		}
	}
	g.P(`}`)
	g.P()

	tags := g.extractTags(msg.Comments)
	if _, ok := tags[_deepcopy]; !ok {
		return
	}

	g.P(fmt.Sprintf(`// DeepCopy is an auto-generated deepcopy function, copying the receiver, creating a new %s.`, mname))
	g.P(fmt.Sprintf(`func (in *%s) DeepCopy() *%s {`, mname, mname))
	g.P(`if in == nil {`)
	g.P(`return nil`)
	g.P(`}`)
	g.P(fmt.Sprintf(`out := new(%s)`, mname))
	g.P(`in.DeepCopyInto(out)`)
	g.P(`return out`)
	g.P(`}`)
	g.P()
}

func (g *deepcopy) generateRepeatedField(file *generator.FileDescriptor, msg *generator.MessageDescriptor, field *generator.FieldDescriptor) {
	fname := generator.CamelCase(field.Proto.GetName())
	g.P(fmt.Sprintf(`if in.%s != nil {`, fname))
	g.P(fmt.Sprintf(`in, out := &in.%s, &out.%s`, fname, fname))
	if strings.HasSuffix(field.Proto.GetTypeName(), "Entry") {
		for _, nest := range msg.Proto.GetNestedType() {
			if strings.HasSuffix(field.Proto.GetTypeName(), nest.GetName()) {
				key, value := nest.GetMapFields()
				keyString, _ := g.buildFieldGoType(file, key)
				valueString, _ := g.buildFieldGoType(file, value)
				g.P(fmt.Sprintf(`*out = make(map[%s]%s, len(*in))`, keyString, valueString))
				g.P(`for key, val := range *in {`)
				if strings.HasPrefix(valueString, "*") {
					subMsg := g.gen.ExtractMessage(value.GetTypeName())
					g.P(fmt.Sprintf(`var outVal *%s`, subMsg.Proto.GetName()))
					g.P(`if val == nil {`)
					g.P(`(*out)[key] = nil`)
					g.P(`} else {`)
					g.P(`in, out := &val, &outVal`)
					g.P(fmt.Sprintf(`*out = new(%s)`, subMsg.Proto.GetName()))
					g.P(`(*in).DeepCopyInto(*out)`)
					g.P(`}`)
					g.P(`(*out)[key] = outVal`)
				} else {
					g.P(`(*out)[key] = val`)
				}
				g.P(`}`)
				break
			}
		}
	} else {
		alias, _ := g.buildFieldGoType(file, field.Proto)
		g.P(fmt.Sprintf(`*out = make([]%s, len(*in))`, alias))
		if strings.HasPrefix(alias, "*") {
			g.P(`for i := range *in {`)
			g.P(`if (*in)[i] != nil {`)
			g.P(`in, out := &(*out)[i], &(*out)[i]`)
			subMsg := g.gen.ExtractMessage(field.Proto.GetTypeName())
			g.P(fmt.Sprintf(`*out = new(%s)`, subMsg.Proto.GetName()))
			g.P(`(*in).DeepCopyInto(*out)`)
			g.P(`}`)
			g.P(`}`)
		} else {
			g.P(`copy(*out, *in)`)
		}
	}
	g.P("}")
}

func (g *deepcopy) generateMessageField(field *generator.FieldDescriptor) {
	fname := generator.CamelCase(field.Proto.GetName())
	subMsg := g.gen.ExtractMessage(field.Proto.GetTypeName())
	fpkg := subMsg.Proto.GetName()
	tags := g.extractTags(field.Comments)
	_, isInline := tags[_inline]
	if isInline {
		g.P(`{`)
		g.P(fmt.Sprintf(`in, out := &in.%s, &out.%s`, fname, fname))
		g.P(fmt.Sprintf(`out = new(%s)`, fpkg))
		g.P(`(*in).DeepCopyInto(out)`)
		g.P(`}`)
	} else {
		g.P(fmt.Sprintf(`if in.%s != nil {`, fname))
		g.P(fmt.Sprintf(`in, out := &in.%s, &out.%s`, fname, fname))
		g.P(fmt.Sprintf(`*out = new(%s)`, fpkg))
		g.P(`(*in).DeepCopyInto(*out)`)
		g.P(`}`)
	}
}


func (g *deepcopy) buildFieldGoType(file *generator.FileDescriptor, field *descriptor.FieldDescriptorProto) (string, error) {
	switch field.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return "float64", nil
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		return "float32", nil
	case descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_ENUM:
		return "int32", nil
	case descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_INT64:
		return "int64", nil
	case descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_UINT32:
		return "uint32", nil
	case descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_UINT64:
		return "uint64", nil
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		return "string", nil
	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		return "[]byte", nil
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		subMsg := g.gen.ExtractMessage(field.GetTypeName())
		if subMsg.Proto.File() == file {
			return "*" + subMsg.Proto.GetName(), nil
		}
		obj := g.gen.ObjectNamed(field.GetTypeName())
		v, ok := g.gen.ImportMap[obj.GoImportPath().String()]
		if !ok {
			v = string(g.gen.AddImport(obj.GoImportPath()))
		}
		return "*" + v + "." + subMsg.Proto.GetName(), nil
	default:
		return "", errors.New("invalid field type")
	}
}

func (g *deepcopy) extractTags(comments []*generator.Comment) map[string]*Tag {
	if comments == nil || len(comments) == 0 {
		return nil
	}
	tags := make(map[string]*Tag, 0)
	for _, c := range comments {
		if c.Tag != TagString || len(c.Text) == 0 {
			continue
		}
		parts := strings.Split(c.Text, ";")
		for _, p := range parts {
			tag := new(Tag)
			p = strings.TrimSpace(p)
			if i := strings.Index(p, "="); i > 0 {
				tag.Key = strings.TrimSpace(p[:i])
				v := strings.TrimSpace(p[i+1:])
				if v == "" {
					g.gen.Fail(fmt.Sprintf("tag '%s' missing value", tag.Key))
				}
				tag.Value = v
			} else {
				tag.Key = p
			}
			tags[tag.Key] = tag
		}
	}

	return tags
}