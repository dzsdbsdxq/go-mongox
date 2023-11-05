// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aggregation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestBuilder_AddKeyValues(t *testing.T) {
	testCases := []struct {
		name      string
		keyValues []any
		expected  bson.D
	}{
		{
			name:      "nil keyValues",
			keyValues: nil,
			expected:  bson.D{},
		},
		{
			name:      "empty keyValues",
			keyValues: []any{},
			expected:  bson.D{},
		},
		{
			name:      "keyValues with odd length",
			keyValues: []any{"a"},
			expected:  bson.D{},
		},
		{
			name:      "key contains non-string",
			keyValues: []any{1, "a"},
			expected:  bson.D{},
		},
		{
			name:      "normal",
			keyValues: []any{"name", "cmy"},
			expected:  bson.D{{Key: "name", Value: "cmy"}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().AddKeyValues(tc.keyValues...).Build())
		})
	}
}