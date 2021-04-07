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

package snapshot

import (
	"encoding/gob"
	"fmt"
	"io"
	"net/url"
	"os"
	"time"

	"github.com/lack-io/vine/service/store"
)

// Restore emits records from a vine store snapshot
type Restore interface {
	// Init validates the RestoreOptions and returns an error if they are invalid.
	// Init must be called before a Restore is used
	Init(opts ...RestoreOption) error
	// Start opens a channel over which records from the snapshot are retrieved.
	// The channel will be closed when the entire snapshot has been read.
	Start() (<-chan *store.Record, error)
}

// RestoreOptions configure a Restore
type RestoreOptions struct {
	Source string
}

// RestoreOption is an individual option
type RestoreOption func(r *RestoreOptions)

// Source is the source URL of a snapshot, e.g. file:///path/to/file
func Source(source string) RestoreOption {
	return func(r *RestoreOptions) {
		r.Source = source
	}
}

// FileRestore reads records from a file
type FileRestore struct {
	Options RestoreOptions

	path string
}

func NewFileRestore(opts ...RestoreOption) Restore {
	r := &FileRestore{}
	for _, o := range opts {
		o(&r.Options)
	}
	return r
}

func (f *FileRestore) Init(opts ...RestoreOption) error {
	for _, o := range opts {
		o(&f.Options)
	}
	u, err := url.Parse(f.Options.Source)
	if err != nil {
		return fmt.Errorf("source is invalid: %w", err)
	}
	if u.Scheme != "file" {
		return fmt.Errorf("unsupported scheme %s (wanted file)", u.Scheme)
	}
	f.path = u.Path
	return nil
}

// Start starts reading records from a file. The returned channel is closed when complete
func (f *FileRestore) Start() (<-chan *store.Record, error) {
	fi, err := os.Open(f.path)
	if err != nil {
		return nil, fmt.Errorf("Couldn't open file %s: %w", f.path, err)
	}
	recordChan := make(chan *store.Record)
	go func(records chan<- *store.Record, reader io.ReadCloser) {
		defer close(recordChan)
		defer reader.Close()
		dec := gob.NewDecoder(fi)
		var r record
		for {
			err := dec.Decode(&r)
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			rec := &store.Record{
				Key: r.Key,
			}
			rec.Value = make([]byte, len(r.Value))
			copy(rec.Value, r.Value)
			if !r.ExpiresAt.IsZero() {
				rec.Expiry = time.Until(r.ExpiresAt)
			}
			recordChan <- rec
		}
	}(recordChan, fi)
	return recordChan, nil
}
