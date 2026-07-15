package project

import "context"

type Reader interface {
	Get(context.Context, string) (Result, error)
}

type Writer interface {
	Create(context.Context, CreateInput) (Result, error)
	Update(context.Context, UpdateInput) (Result, error)
	Delete(context.Context, string) error
}

type Store interface {
	Reader
	Writer
}
