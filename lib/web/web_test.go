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

package web_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/lack-io/cli"

	"github.com/lack-io/vine"
	"github.com/lack-io/vine/lib/logger"
	"github.com/lack-io/vine/lib/web"
)

func TestWeb(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println("Test nr", i)
		testFunc()
	}
}

func testFunc() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*250)
	defer cancel()

	s := vine.NewService(
		vine.Name("test"),
		vine.Context(ctx),
		vine.HandleSignal(false),
		vine.Flags(
			&cli.StringFlag{
				Name: "test.timeout",
			},
			&cli.BoolFlag{
				Name: "test.v",
			},
			&cli.StringFlag{
				Name: "test.run",
			},
			&cli.StringFlag{
				Name: "test.testlogfile",
			},
		),
	)
	w := web.NewService(
		web.VineService(s),
		web.Context(ctx),
		web.HandleSignal(false),
	)
	//s.Init()
	//w.Init()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := s.Run()
		if err != nil {
			logger.Errorf("vine run error: %v", err)
		}
	}()
	go func() {
		defer wg.Done()
		err := w.Run()
		if err != nil {
			logger.Errorf("web run error: %v", err)
		}
	}()

	wg.Wait()
}
