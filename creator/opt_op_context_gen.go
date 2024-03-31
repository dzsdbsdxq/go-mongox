// Generated by optioner -type OpContext; DO NOT EDIT
// If you have any questions, please create issues and submit contributions at:
// https://github.com/chenmingyong0423/go-optioner

package creator

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type OpContextOption[T any] func(*OpContext[T])

func NewOpContext[T any](col *mongo.Collection, opts ...OpContextOption[T]) *OpContext[T] {
	opContext := &OpContext[T]{
		Col: col,
	}

	for _, opt := range opts {
		opt(opContext)
	}

	return opContext
}

func WithDoc[T any](doc *T) OpContextOption[T] {
	return func(opContext *OpContext[T]) {
		opContext.Doc = doc
	}
}

func WithDocs[T any](docs []*T) OpContextOption[T] {
	return func(opContext *OpContext[T]) {
		opContext.Docs = docs
	}
}