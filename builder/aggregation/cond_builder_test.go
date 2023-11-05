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

func Test_condBuilder_Cond(t *testing.T) {
	testCases := []struct {
		name      string
		boolExpr  any
		trueExpr  any
		falseExpr any
		expected  bson.D
	}{
		{
			name:      "normal",
			boolExpr:  BsonBuilder().Gte("$qty", 250).Build(),
			trueExpr:  30,
			falseExpr: 20,
			expected:  bson.D{{Key: "$cond", Value: []any{bson.D{{Key: "$gte", Value: []any{"$qty", 250}}}, 30, 20}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Cond(tc.boolExpr, tc.trueExpr, tc.falseExpr).Build())
		})
	}
}

func Test_condBuilder_IfNull(t *testing.T) {
	testCases := []struct {
		name        string
		expr        any
		replacement any
		expected    bson.D
	}{
		{
			name:        "nil expr",
			expr:        nil,
			replacement: "Unspecified",
			expected:    bson.D{{Key: "$ifNull", Value: []any{nil, "Unspecified"}}},
		},
		{
			name:        "nil replacement",
			expr:        "$description",
			replacement: nil,
			expected:    bson.D{{Key: "$ifNull", Value: []any{"$description", nil}}},
		},
		{
			name:        "normal",
			expr:        "$description",
			replacement: "Unspecified",
			expected:    bson.D{{Key: "$ifNull", Value: []any{"$description", "Unspecified"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().IfNull(tc.expr, tc.replacement).Build())
		})
	}
}

func Test_condBuilder_Switch(t *testing.T) {
	testCases := []struct {
		name        string
		cases       []any
		defaultCase any
		expected    bson.D
	}{
		{
			name:        "nil cases",
			cases:       nil,
			defaultCase: "Did not match",
			expected:    bson.D{},
		},
		{
			name:        "empty cases",
			cases:       []any{},
			defaultCase: "Did not match",
			expected:    bson.D{},
		},
		{
			name: "normal",
			cases: []any{
				BsonBuilder().Eq(0, 5).Build(), "equals",
				BsonBuilder().Gt(0, 5).Build(), "greater than",
			},
			defaultCase: "Did not match",
			expected: bson.D{
				{Key: "$switch", Value: bson.D{
					{Key: "branches", Value: bson.A{
						bson.D{{Key: "case", Value: bson.D{{Key: "$eq", Value: []any{0, 5}}}}, {Key: "then", Value: "equals"}},
						bson.D{{Key: "case", Value: bson.D{{Key: "$gt", Value: []any{0, 5}}}}, {Key: "then", Value: "greater than"}},
					}},
					{Key: "default", Value: "Did not match"},
				}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Switch(tc.cases, tc.defaultCase).Build())
		})
	}
}