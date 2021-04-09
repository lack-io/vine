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

// Package grpc provides a grpc transport
package grpc

import (
	"context"
	"crypto/tls"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/lack-io/vine/proto/services/transport"
	"github.com/lack-io/vine/service/transport"
	maddr "github.com/lack-io/vine/util/addr"
	mnet "github.com/lack-io/vine/util/net"
	mls "github.com/lack-io/vine/util/tls"
)

type grpcTransport struct {
	opts transport.Options
}

type grpcTransportListener struct {
	listener net.Listener
	secure   bool
	tls      *tls.Config
}

func getTLSConfig(addr string) (*tls.Config, error) {
	hosts := []string{addr}

	// check if its a valid host:port
	if host, _, err := net.SplitHostPort(addr); err == nil {
		if len(host) == 0 {
			hosts = maddr.IPs()
		} else {
			hosts = []string{host}
		}
	}

	// generate a certificate
	cert, err := mls.Certificate(hosts...)
	if err != nil {
		return nil, err
	}
	return &tls.Config{Certificates: []tls.Certificate{cert}}, nil
}

func (t *grpcTransportListener) Addr() string {
	return t.listener.Addr().String()
}

func (t *grpcTransportListener) Close() error {
	return t.listener.Close()
}

func (t *grpcTransportListener) Accept(fn func(transport.Socket)) error {
	var opts []grpc.ServerOption

	// setup tls if specified
	if t.secure || t.tls != nil {
		config := t.tls
		if config == nil {
			var err error
			addr := t.listener.Addr().String()
			config, err = getTLSConfig(addr)
			if err != nil {
				return err
			}
		}

		creds := credentials.NewTLS(config)
		opts = append(opts, grpc.Creds(creds))
	}

	// new service
	svc := grpc.NewServer(opts...)

	// register service
	pb.RegisterTransportServer(svc, &vineTransport{addr: t.listener.Addr().String(), fn: fn})

	// start serving
	return svc.Serve(t.listener)
}

func (t *grpcTransport) Dial(addr string, opts ...transport.DialOption) (transport.Client, error) {
	dopts := transport.DialOptions{
		Timeout: transport.DefaultDialTimeout,
	}

	for _, opt := range opts {
		opt(&dopts)
	}

	options := []grpc.DialOption{
		grpc.WithTimeout(dopts.Timeout),
	}

	if t.opts.Secure || t.opts.TLSConfig != nil {
		config := t.opts.TLSConfig
		if config == nil {
			config = &tls.Config{
				InsecureSkipVerify: true,
			}
		}
		creds := credentials.NewTLS(config)
		options = append(options, grpc.WithTransportCredentials(creds))
	} else {
		options = append(options, grpc.WithInsecure())
	}

	// dial the server
	conn, err := grpc.Dial(addr, options...)
	if err != nil {
		return nil, err
	}

	// create stream
	stream, err := pb.NewTransportClient(conn).Stream(context.Background())
	if err != nil {
		return nil, err
	}

	// return a client
	return &grpcTransportClient{
		conn:   conn,
		stream: stream,
		local:  "localhost",
		remote: addr,
	}, nil
}

func (t *grpcTransport) Listen(addr string, opts ...transport.ListenOption) (transport.Listener, error) {
	var options transport.ListenOptions
	for _, o := range opts {
		o(&options)
	}

	ln, err := mnet.Listen(addr, func(addr string) (net.Listener, error) {
		return net.Listen("tcp", addr)
	})
	if err != nil {
		return nil, err
	}

	return &grpcTransportListener{
		listener: ln,
		tls:      t.opts.TLSConfig,
		secure:   t.opts.Secure,
	}, nil
}

func (t *grpcTransport) Init(opts ...transport.Option) error {
	for _, o := range opts {
		o(&t.opts)
	}
	return nil
}

func (t *grpcTransport) Options() transport.Options {
	return t.opts
}

func (t *grpcTransport) String() string {
	return "grpc"
}

func NewTransport(opts ...transport.Option) transport.Transport {
	var options transport.Options
	for _, o := range opts {
		o(&options)
	}
	return &grpcTransport{opts: options}
}
