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

package update

import (
	"testing"

	"github.com/chenmingyong0423/go-mongox/builder/query"
	"github.com/chenmingyong0423/go-mongox/pkg"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_arrayUpdateBuilder_AddToSet(t *testing.T) {

	testCases := []struct {
		name      string
		keyValues []any

		want bson.D
	}{
		{
			name: "zero params",
			want: bson.D{bson.E{Key: "$addToSet", Value: bson.D{}}},
		},
		{
			name:      "odd params",
			keyValues: []any{"colors", "mauve", "letters"},
			want:      bson.D{bson.E{Key: "$addToSet", Value: bson.D{}}},
		},
		{
			name:      "keys contain non-string",
			keyValues: []any{"colors", "mauve", 2, "letters"},
			want:      bson.D{bson.E{Key: "$addToSet", Value: bson.D{bson.E{Key: "colors", Value: "mauve"}}}},
		},
		{
			name:      "normal params",
			keyValues: []any{"colors", "mauve", "letters", []string{"a", "b", "c"}},
			want:      bson.D{bson.E{Key: "$addToSet", Value: bson.D{bson.E{Key: "colors", Value: "mauve"}, bson.E{Key: "letters", Value: []string{"a", "b", "c"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().AddToSet(tc.keyValues...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_AddToSetForMap(t *testing.T) {

	testCases := []struct {
		name      string
		keyValues map[string]any
		want      bson.D
	}{
		{
			name: "nil map",
			want: bson.D{},
		},
		{
			name:      "empty map",
			keyValues: map[string]any{},
			want:      bson.D{bson.E{Key: "$addToSet", Value: bson.D{}}},
		},
		{
			name:      "normal map",
			keyValues: map[string]any{"colors": "mauve", "letters": []string{"a", "b", "c"}},
			want:      bson.D{bson.E{Key: "$addToSet", Value: bson.D{bson.E{Key: "colors", Value: "mauve"}, bson.E{Key: "letters", Value: []string{"a", "b", "c"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().AddToSetForMap(tc.keyValues).Build())
		})
	}
}

func Test_arrayUpdateBuilder_Pop(t *testing.T) {
	testCases := []struct {
		name      string
		keyValues []any
		want      bson.D
	}{
		{
			name: "zero params",
			want: bson.D{bson.E{Key: "$pop", Value: bson.D{}}},
		},
		{
			name:      "odd params",
			keyValues: []any{"scores", -1, "letters"},
			want:      bson.D{bson.E{Key: "$pop", Value: bson.D{}}},
		},
		{
			name:      "keys contain non-string",
			keyValues: []any{"scores", 1, 1, "letters"},
			want:      bson.D{bson.E{Key: "$pop", Value: bson.D{bson.E{Key: "scores", Value: 1}}}},
		},
		{
			name:      "value contain non-int",
			keyValues: []any{"scores", 1, "letters", "1"},
			want:      bson.D{bson.E{Key: "$pop", Value: bson.D{bson.E{Key: "scores", Value: 1}}}},
		},
		{
			name:      "normal params",
			keyValues: []any{"scores", 1, "letters", -1},
			want:      bson.D{bson.E{Key: "$pop", Value: bson.D{bson.E{Key: "scores", Value: 1}, bson.E{Key: "letters", Value: -1}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Pop(tc.keyValues...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_PopForMap(t *testing.T) {
	testCases := []struct {
		name      string
		keyValues map[string]int
		want      bson.D
	}{
		{
			name: "nil map",
			want: bson.D{},
		},
		{
			name:      "empty map",
			keyValues: map[string]int{},
			want:      bson.D{bson.E{Key: "$pop", Value: bson.D{}}},
		},
		{
			name:      "normal map",
			keyValues: map[string]int{"scores": 1, "letters": -1},
			want:      bson.D{bson.E{Key: "$pop", Value: bson.D{bson.E{Key: "scores", Value: 1}, bson.E{Key: "letters", Value: -1}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().PopForMap(tc.keyValues).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_Pull(t *testing.T) {
	// { $pull: { fruits: { $in: [ "apples", "oranges" ] }, votes: { $gte: 6 }, vegetables: "carrots" } }
	assert.Equal(t, bson.D{bson.E{Key: "$pull", Value: bson.D{bson.E{Key: "fruits", Value: bson.M{"$in": []any{"apples", "oranges"}}}, bson.E{Key: "votes", Value: bson.M{"$gte": 6}}, bson.E{Key: "vegetables", Value: "carrots"}}}}, BsonBuilder().Pull(
		// 可以使用 query 构造
		query.BsonBuilder().InString("fruits", []string{"apples", "oranges"}...).Gte("votes", 6).Add("vegetables", "carrots").Build(),
	).Build())
}

func Test_arrayUpdateBuilder_Push(t *testing.T) {
	testCases := []struct {
		name      string
		keyValues []any
		want      bson.D
	}{
		{
			name: "zero params",
			want: bson.D{bson.E{Key: "$push", Value: bson.D{}}},
		},
		{
			name:      "odd params",
			keyValues: []any{"scores", 1, "letters"},
			want:      bson.D{bson.E{Key: "$push", Value: bson.D{}}},
		},
		{
			name:      "keys contain non-string",
			keyValues: []any{"scores", BsonBuilder().Add("$each", []int{90, 82, 85}).Build(), 1, "letters"},
			want:      bson.D{bson.E{Key: "$push", Value: bson.D{bson.E{Key: "scores", Value: bson.D{bson.E{Key: "$each", Value: []int{90, 82, 85}}}}}}},
		},
		{
			name:      "normal params",
			keyValues: []any{"scores", BsonBuilder().Add("$each", []int{90, 82, 85}).Build(), "scores", 1},
			want:      bson.D{bson.E{Key: "$push", Value: bson.D{bson.E{Key: "scores", Value: bson.D{bson.E{Key: "$each", Value: []int{90, 82, 85}}}}, bson.E{Key: "scores", Value: 1}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Push(tc.keyValues...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_PullAll(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []any
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]any)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []any{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []any{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []any{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []any{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []any{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []any{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().PullAll(tc.key, tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_PullAllInt(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]int)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []int{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []int{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []int{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().PullAllInt(tc.key, tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_PullAllInt8(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int8
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]int8)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []int8{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int8{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []int8{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int8{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []int8{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int8{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().PullAllInt8(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllInt16(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int16
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]int16)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []int16{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int16{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []int16{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int16{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []int16{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int16{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().PullAllInt16(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllInt32(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int32
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]int32)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []int32{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int32{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []int32{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int32{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []int32{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int32{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().PullAllInt32(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllInt64(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int64
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]int64)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []int64{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int64{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []int64{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int64{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []int64{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int64{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().PullAllInt64(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllString(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []string
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]string)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []string{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []string{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []string{"1"},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []string{"1"}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []string{"1", "2", "3"},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []string{"1", "2", "3"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().PullAllString(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllFloat32(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []float32
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]float32)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []float32{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []float32{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []float32{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []float32{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []float32{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []float32{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().PullAllFloat32(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllFloat64(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []float64
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]float64)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []float64{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []float64{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []float64{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []float64{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []float64{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []float64{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().PullAllFloat64(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllUint(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]uint)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []uint{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []uint{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []uint{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().PullAllUint(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllUint8(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint8
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]uint8)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []uint8{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint8{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []uint8{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint8{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []uint8{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint8{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().PullAllUint8(tc.key, tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_PullAllUint16(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint16
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]uint16)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []uint16{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint16{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []uint16{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint16{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []uint16{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint16{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().PullAllUint16(tc.key, tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_PullAllUint32(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint32
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]uint32)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []uint32{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint32{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []uint32{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint32{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []uint32{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint32{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().PullAllUint32(tc.key, tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_PullAllUint64(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint64
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]uint64)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []uint64{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint64{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []uint64{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint64{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []uint64{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint64{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().PullAllUint64(tc.key, tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_Each(t *testing.T) {
	testCases := []struct {
		name   string
		values []any
		want   bson.D
	}{
		{
			name: "nil values",
			want: bson.D{bson.E{Key: "$each", Value: []any(nil)}},
		},
		{
			name:   "empty values",
			values: []any{},
			want:   bson.D{bson.E{Key: "$each", Value: []any{}}},
		},
		{
			name:   "single values",
			values: []any{"99"},
			want:   bson.D{bson.E{Key: "$each", Value: []any{"99"}}},
		},
		{
			name:   "multiple values",
			values: []any{"99", "98", "97"},
			want:   bson.D{bson.E{Key: "$each", Value: []any{"99", "98", "97"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Each(tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_EachInt(t *testing.T) {
	testCases := []struct {
		name   string
		values []int
		want   bson.D
	}{
		{
			name: "nil values",
			want: bson.D{bson.E{Key: "$each", Value: []int(nil)}},
		},
		{
			name:   "empty values",
			values: []int{},
			want:   bson.D{bson.E{Key: "$each", Value: []int{}}},
		},
		{
			name:   "single values",
			values: []int{99},
			want:   bson.D{bson.E{Key: "$each", Value: []int{99}}},
		},
		{
			name:   "multiple values",
			values: []int{99, 98, 97},
			want:   bson.D{bson.E{Key: "$each", Value: []int{99, 98, 97}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().EachInt(tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_EachInt8(t *testing.T) {
	testCases := []struct {
		name   string
		values []int8
		want   bson.D
	}{
		{
			name: "nil values",
			want: bson.D{bson.E{Key: "$each", Value: []int8(nil)}},
		},
		{
			name:   "empty values",
			values: []int8{},
			want:   bson.D{bson.E{Key: "$each", Value: []int8{}}},
		},
		{
			name:   "single values",
			values: []int8{99},
			want:   bson.D{bson.E{Key: "$each", Value: []int8{99}}},
		},
		{
			name:   "multiple values",
			values: []int8{99, 98, 97},
			want:   bson.D{bson.E{Key: "$each", Value: []int8{99, 98, 97}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().EachInt8(tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_EachInt16(t *testing.T) {
	testCases := []struct {
		name   string
		values []int16
		want   bson.D
	}{
		{
			name: "nil values",
			want: bson.D{bson.E{Key: "$each", Value: []int16(nil)}},
		},
		{
			name:   "empty values",
			values: []int16{},
			want:   bson.D{bson.E{Key: "$each", Value: []int16{}}},
		},
		{
			name:   "single values",
			values: []int16{99},
			want:   bson.D{bson.E{Key: "$each", Value: []int16{99}}},
		},
		{
			name:   "multiple values",
			values: []int16{99, 98, 97},
			want:   bson.D{bson.E{Key: "$each", Value: []int16{99, 98, 97}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().EachInt16(tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_EachInt32(t *testing.T) {
	testCases := []struct {
		name   string
		values []int32
		want   bson.D
	}{
		{
			name: "nil values",
			want: bson.D{bson.E{Key: "$each", Value: []int32(nil)}},
		},
		{
			name:   "empty values",
			values: []int32{},
			want:   bson.D{bson.E{Key: "$each", Value: []int32{}}},
		},
		{
			name:   "single values",
			values: []int32{99},
			want:   bson.D{bson.E{Key: "$each", Value: []int32{99}}},
		},
		{
			name:   "multiple values",
			values: []int32{99, 98, 97},
			want:   bson.D{bson.E{Key: "$each", Value: []int32{99, 98, 97}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().EachInt32(tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_EachInt64(t *testing.T) {
	testCases := []struct {
		name   string
		values []int64
		want   bson.D
	}{
		{
			name: "nil values",
			want: bson.D{bson.E{Key: "$each", Value: []int64(nil)}},
		},
		{
			name:   "empty values",
			values: []int64{},
			want:   bson.D{bson.E{Key: "$each", Value: []int64{}}},
		},
		{
			name:   "single values",
			values: []int64{99},
			want:   bson.D{bson.E{Key: "$each", Value: []int64{99}}},
		},
		{
			name:   "multiple values",
			values: []int64{99, 98, 97},
			want:   bson.D{bson.E{Key: "$each", Value: []int64{99, 98, 97}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().EachInt64(tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_EachString(t *testing.T) {
	testCases := []struct {
		name   string
		values []string
		want   bson.D
	}{
		{
			name: "nil values",
			want: bson.D{bson.E{Key: "$each", Value: []string(nil)}},
		},
		{
			name:   "empty values",
			values: []string{},
			want:   bson.D{bson.E{Key: "$each", Value: []string{}}},
		},
		{
			name:   "single values",
			values: []string{"99"},
			want:   bson.D{bson.E{Key: "$each", Value: []string{"99"}}},
		},
		{
			name:   "multiple values",
			values: []string{"99", "98", "97"},
			want:   bson.D{bson.E{Key: "$each", Value: []string{"99", "98", "97"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().EachString(tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_EachFloat32(t *testing.T) {
	testCases := []struct {
		name   string
		values []float32
		want   bson.D
	}{
		{
			name: "nil values",
			want: bson.D{bson.E{Key: "$each", Value: []float32(nil)}},
		},
		{
			name:   "empty values",
			values: []float32{},
			want:   bson.D{bson.E{Key: "$each", Value: []float32{}}},
		},
		{
			name:   "single values",
			values: []float32{99},
			want:   bson.D{bson.E{Key: "$each", Value: []float32{99}}},
		},
		{
			name:   "multiple values",
			values: []float32{99, 98, 97},
			want:   bson.D{bson.E{Key: "$each", Value: []float32{99, 98, 97}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().EachFloat32(tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_EachFloat64(t *testing.T) {
	testCases := []struct {
		name   string
		values []float64
		want   bson.D
	}{
		{
			name: "nil values",
			want: bson.D{bson.E{Key: "$each", Value: []float64(nil)}},
		},
		{
			name:   "empty values",
			values: []float64{},
			want:   bson.D{bson.E{Key: "$each", Value: []float64{}}},
		},
		{
			name:   "single values",
			values: []float64{99},
			want:   bson.D{bson.E{Key: "$each", Value: []float64{99}}},
		},
		{
			name:   "multiple values",
			values: []float64{99, 98, 97},
			want:   bson.D{bson.E{Key: "$each", Value: []float64{99, 98, 97}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().EachFloat64(tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_EachUint(t *testing.T) {
	testCases := []struct {
		name   string
		values []uint
		want   bson.D
	}{
		{
			name: "nil values",
			want: bson.D{bson.E{Key: "$each", Value: []uint(nil)}},
		},
		{
			name:   "empty values",
			values: []uint{},
			want:   bson.D{bson.E{Key: "$each", Value: []uint{}}},
		},
		{
			name:   "single values",
			values: []uint{99},
			want:   bson.D{bson.E{Key: "$each", Value: []uint{99}}},
		},
		{
			name:   "multiple values",
			values: []uint{99, 98, 97},
			want:   bson.D{bson.E{Key: "$each", Value: []uint{99, 98, 97}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().EachUint(tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_EachUint8(t *testing.T) {
	testCases := []struct {
		name   string
		values []uint8
		want   bson.D
	}{
		{
			name: "nil values",
			want: bson.D{bson.E{Key: "$each", Value: []uint8(nil)}},
		},
		{
			name:   "empty values",
			values: []uint8{},
			want:   bson.D{bson.E{Key: "$each", Value: []uint8{}}},
		},
		{
			name:   "single values",
			values: []uint8{99},
			want:   bson.D{bson.E{Key: "$each", Value: []uint8{99}}},
		},
		{
			name:   "multiple values",
			values: []uint8{99, 98, 97},
			want:   bson.D{bson.E{Key: "$each", Value: []uint8{99, 98, 97}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().EachUint8(tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_EachUint16(t *testing.T) {
	testCases := []struct {
		name   string
		values []uint16
		want   bson.D
	}{
		{
			name: "nil values",
			want: bson.D{bson.E{Key: "$each", Value: []uint16(nil)}},
		},
		{
			name:   "empty values",
			values: []uint16{},
			want:   bson.D{bson.E{Key: "$each", Value: []uint16{}}},
		},
		{
			name:   "single values",
			values: []uint16{99},
			want:   bson.D{bson.E{Key: "$each", Value: []uint16{99}}},
		},
		{
			name:   "multiple values",
			values: []uint16{99, 98, 97},
			want:   bson.D{bson.E{Key: "$each", Value: []uint16{99, 98, 97}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().EachUint16(tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_EachUint32(t *testing.T) {
	testCases := []struct {
		name   string
		values []uint32
		want   bson.D
	}{
		{
			name: "nil values",
			want: bson.D{bson.E{Key: "$each", Value: []uint32(nil)}},
		},
		{
			name:   "empty values",
			values: []uint32{},
			want:   bson.D{bson.E{Key: "$each", Value: []uint32{}}},
		},
		{
			name:   "single values",
			values: []uint32{99},
			want:   bson.D{bson.E{Key: "$each", Value: []uint32{99}}},
		},
		{
			name:   "multiple values",
			values: []uint32{99, 98, 97},
			want:   bson.D{bson.E{Key: "$each", Value: []uint32{99, 98, 97}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().EachUint32(tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_EachUint64(t *testing.T) {
	testCases := []struct {
		name   string
		values []uint64
		want   bson.D
	}{
		{
			name: "nil values",
			want: bson.D{bson.E{Key: "$each", Value: []uint64(nil)}},
		},
		{
			name:   "empty values",
			values: []uint64{},
			want:   bson.D{bson.E{Key: "$each", Value: []uint64{}}},
		},
		{
			name:   "single values",
			values: []uint64{99},
			want:   bson.D{bson.E{Key: "$each", Value: []uint64{99}}},
		},
		{
			name:   "multiple values",
			values: []uint64{99, 98, 97},
			want:   bson.D{bson.E{Key: "$each", Value: []uint64{99, 98, 97}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().EachUint64(tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_Position(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "$position", Value: 1}}, BsonBuilder().Position(1).Build())
}

func Test_arrayUpdateBuilder_Slice(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "$slice", Value: 1}}, BsonBuilder().Slice(1).Build())
}

func Test_arrayUpdateBuilder_Sort(t *testing.T) {
	testCases := []struct {
		name      string
		keyValues []any

		want bson.D
	}{
		{
			name: "zero keyValues",
			want: bson.D{bson.E{Key: "$sort", Value: bson.D{}}},
		},
		{
			name:      "key contains non-string",
			keyValues: []any{"score", -1, 1, "name"},
			want:      bson.D{bson.E{Key: "$sort", Value: bson.D{bson.E{Key: "score", Value: -1}}}},
		},
		{
			name:      "values contains non-int",
			keyValues: []any{"score", -1, "name", "1"},
			want:      bson.D{bson.E{Key: "$sort", Value: bson.D{bson.E{Key: "score", Value: -1}}}},
		},
		{
			name:      "normal keyValues",
			keyValues: []any{"score", -1, "name", 1},
			want:      bson.D{bson.E{Key: "$sort", Value: bson.D{bson.E{Key: "score", Value: -1}, bson.E{Key: "name", Value: 1}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Sort(tc.keyValues...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_SortForMap(t *testing.T) {
	testCases := []struct {
		name  string
		value map[string]int

		want bson.D
	}{
		{
			name:  "nil value",
			value: nil,
			want:  bson.D{},
		},
		{
			name:  "empty value",
			value: map[string]int{},
			want:  bson.D{bson.E{Key: "$sort", Value: bson.D{}}},
		},
		{
			name:  "normal value",
			value: map[string]int{"score": -1, "name": 1},
			want:  bson.D{bson.E{Key: "$sort", Value: bson.D{bson.E{Key: "score", Value: -1}, bson.E{Key: "name", Value: 1}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().SortForMap(tc.value).Build())
		})
	}
}
