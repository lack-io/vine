// Copyright 2020 The vine Authors
//
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

// Package handler is the handler for the `vine debug stats` service
package stats

import (
	"context"
	"sync"
	"time"

	"github.com/lack-io/vine/proto/debug"
	"github.com/lack-io/vine/proto/debug/stats"
	"github.com/lack-io/vine/proto/errors"
	"github.com/lack-io/vine/service/client"
	"github.com/lack-io/vine/service/config/cmd"
	"github.com/lack-io/vine/service/registry"
	"github.com/lack-io/vine/util/ring"
)

// New initialises and returns a new Stats service handler
func New(done <-chan bool, windowSize int, services func() []*registry.Service) (*Stats, error) {
	s := &Stats{
		client:    *cmd.DefaultOptions().Client,
		snapshots: ring.New(windowSize),
		services:  services,
	}

	s.Start(done)
	return s, nil
}

// Stats is the Debug.Stats handler
type Stats struct {
	client client.Client

	sync.RWMutex
	// historical snapshots from the start
	snapshots *ring.Buffer
	// returns list of services
	services func() []*registry.Service
}

// Read returns gets a snapshot of all current stats
func (s *Stats) Read(ctx context.Context, req *stats.ReadRequest, rsp *stats.ReadResponse) error {
	allSnapshots := []*stats.Snapshot{}

	func() {
		s.RLock()
		defer s.RUnlock()

		// get last snapshot
		numEntries := 1

		if req.Past {
			numEntries = -1
		}

		entries := s.snapshots.Get(numEntries)

		for _, entry := range entries {
			allSnapshots = append(allSnapshots, entry.Value.([]*stats.Snapshot)...)
		}
	}()
	if req.Service == nil {
		rsp.Stats = allSnapshots
		return nil
	}
	filter := func(a, b string) bool {
		if len(b) == 0 {
			return true
		}
		return a == b
	}
	filteredSnapshots := []*stats.Snapshot{}
	for _, s := range allSnapshots {
		if !filter(s.Service.Name, req.Service.Name) {
			continue
		}
		if !filter(s.Service.Version, req.Service.Version) {
			continue
		}
		filteredSnapshots = append(filteredSnapshots, s)
	}
	rsp.Stats = filteredSnapshots
	return nil
}

func (s *Stats) Write(ctx context.Context, req *stats.WriteRequest, rsp *stats.WriteResponse) error {
	return errors.BadRequest("go.vine.debug.stats", "not implemented")
}

// Stream starts streaming stats
func (s *Stats) Stream(ctx context.Context, req *stats.StreamRequest, rsp stats.Stats_StreamStream) error {
	return errors.BadRequest("go.vine.debug.stats", "not implemented")
}

// Start Starts scraping other services until the provided channel is closed
func (s *Stats) Start(done <-chan bool) {
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				s.scrape()
				time.Sleep(time.Second)
			}
		}
	}()
}

func (s *Stats) scrape() {
	// get services
	cached := s.services()

	s.RLock()
	// Create a local copy of cached services
	services := make([]*registry.Service, len(cached))
	copy(services, cached)
	s.RUnlock()

	// Start building the next list of snapshots
	var mtx sync.Mutex
	next := make([]*stats.Snapshot, 0)

	// Call each node of each service in goroutines
	var wg sync.WaitGroup

	protocol := s.client.String()

	for _, svc := range services {
		// Ignore nodeless and non mucp services
		if len(svc.Nodes) == 0 {
			continue
		}
		// Call every node
		for _, node := range svc.Nodes {
			if node.Metadata["protocol"] != protocol {
				continue
			}

			wg.Add(1)

			go func(st *Stats, service *registry.Service, node *registry.Node) {
				defer wg.Done()

				// create new context to cancel within a few seconds
				ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
				defer cancel()

				req := s.client.NewRequest(service.Name, "Debug.Stats", &debug.StatsRequest{})
				rsp := new(debug.StatsResponse)
				if err := s.client.Call(ctx, req, rsp, client.WithAddress(node.Address)); err != nil {
					// Don't report an error to the user if stats can't be collected, just continue
					return
				}

				// Append the new snapshot
				snap := &stats.Snapshot{
					Service: &stats.Service{
						Name:    service.Name,
						Version: service.Version,
						Node: &stats.Node{
							Id:      node.Id,
							Address: node.Address,
						},
					},
					Started:  int64(rsp.Started),
					Uptime:   rsp.Uptime,
					Memory:   rsp.Memory,
					Threads:  rsp.Threads,
					Gc:       rsp.Gc,
					Requests: rsp.Requests,
					Errors:   rsp.Errors,
				}
				timestamp := time.Now().Unix()
				snap.Timestamp = uint64(timestamp)
				mtx.Lock()
				next = append(next, snap)
				mtx.Unlock()
			}(s, svc, node)
		}
	}
	wg.Wait()

	// Swap in the snapshots
	s.Lock()
	s.snapshots.Put(next)
	s.Unlock()
}