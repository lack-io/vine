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

package noop

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lack-io/vine/service/metrics"
)

func TestNoopReporter(t *testing.T) {

	// Make a Reporter:
	reporter := New(metrics.Path("/noop"))
	assert.NotNil(t, reporter)
	assert.Equal(t, "/noop", reporter.options.Path)

	// Check that our implementation is valid:
	assert.Implements(t, new(metrics.Reporter), reporter)
}
