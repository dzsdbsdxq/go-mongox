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

	"github.com/chenmingyong0423/go-mongox/builder/query"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_logicalBuilder_And(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil expressions",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$and", Value: []any{nil}}},
		},
		{
			name:        "empty expressions",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$and", Value: []any{}}},
		},
		{
			name:        "normal expressions",
			expressions: []any{BsonBuilder().Gt("$qty", 100).Build(), BsonBuilder().Lt("$qty", 250).Build()},
			expected:    bson.D{bson.E{Key: "$and", Value: []any{bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 100}}}, bson.D{bson.E{Key: "$lt", Value: []any{"$qty", 250}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().And(tc.expressions...).Build())
		})
	}
}

func Test_logicalBuilder_Not(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil expressions",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$not", Value: []any{nil}}},
		},
		{
			name:        "empty expressions",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$not", Value: []any{}}},
		},
		{
			name:        "normal expressions",
			expressions: []any{BsonBuilder().Gt("$qty", 250).Build()},
			expected:    bson.D{bson.E{Key: "$not", Value: []any{bson.D{bson.E{Key: "$gt", Value: []any{"$qty", 250}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Not(tc.expressions...).Build())
		})
	}
}

func Test_logicalBuilder_Or(t *testing.T) {
	testCases := []struct {
		name        string
		expressions []any
		expected    bson.D
	}{
		{
			name:        "nil expressions",
			expressions: []any{nil},
			expected:    bson.D{bson.E{Key: "$or", Value: []any{nil}}},
		},
		{
			name:        "empty expressions",
			expressions: []any{},
			expected:    bson.D{bson.E{Key: "$or", Value: []any{}}},
		},
		{
			name:        "normal expressions",
			expressions: []any{query.BsonBuilder().Eq("x", 0).Build(), query.BsonBuilder().Expr(BsonBuilder().Eq(BsonBuilder().Divide(1, "$x").Build(), 3).Build()).Build()},
			expected:    bson.D{bson.E{Key: "$or", Value: []any{bson.D{bson.E{Key: "x", Value: bson.M{"$eq": 0}}}, bson.D{bson.E{Key: "$expr", Value: bson.D{bson.E{Key: "$eq", Value: []any{bson.D{bson.E{Key: "$divide", Value: []any{1, "$x"}}}, 3}}}}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, BsonBuilder().Or(tc.expressions...).Build())
		})
	}
}