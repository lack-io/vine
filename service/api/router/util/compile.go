// MIT License
//
// Copyright (c) 2020 Lack
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

package util

// download from https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/master/protoc-gen-grpc-gateway/httprule/compile.go

const (
	opcodeVersion = 1
)

// Template is a compiled representation of path templates.
type Template struct {
	// Version is the version number of the format.
	Version int
	// OpCodes is a sequence of operations.
	OpCodes []int
	// Pool is a constant pool
	Pool []string
	// Verb is a VERB part in the template.
	Verb string
	// Fields is a list of field paths bound in this template.
	Fields []string
	// Original template (example: /v1/a_bit_of_everything)
	Template string
}

// Compiler compiles utilities representation of path templates into marshallable operations.
// They can be unmarshalled by runtime.NewPattern.
type Compiler interface {
	Compile() Template
}

type op struct {
	// code is the opcode of the operation
	code OpCode

	// str is a string operand of the code.
	// operand is ignored if str is not empty.
	str string

	// operand is a numeric operand of the code.
	operand int
}

func (w wildcard) compile() []op {
	return []op{
		{code: OpPush},
	}
}

func (w deepWildcard) compile() []op {
	return []op{
		{code: OpPushM},
	}
}

func (l literal) compile() []op {
	return []op{
		{
			code: OpLitPush,
			str:  string(l),
		},
	}
}

func (v variable) compile() []op {
	var ops []op
	for _, s := range v.segments {
		ops = append(ops, s.compile()...)
	}
	ops = append(ops, op{
		code:    OpConcatN,
		operand: len(v.segments),
	}, op{
		code: OpCapture,
		str:  v.path,
	})

	return ops
}

func (t template) Compile() Template {
	var rawOps []op
	for _, s := range t.segments {
		rawOps = append(rawOps, s.compile()...)
	}

	var (
		ops    []int
		pool   []string
		fields []string
	)
	consts := make(map[string]int)
	for _, op := range rawOps {
		ops = append(ops, int(op.code))
		if op.str == "" {
			ops = append(ops, op.operand)
		} else {
			if _, ok := consts[op.str]; !ok {
				consts[op.str] = len(pool)
				pool = append(pool, op.str)
			}
			ops = append(ops, consts[op.str])
		}
		if op.code == OpCapture {
			fields = append(fields, op.str)
		}
	}
	return Template{
		Version:  opcodeVersion,
		OpCodes:  ops,
		Pool:     pool,
		Verb:     t.verb,
		Fields:   fields,
		Template: t.template,
	}
}
