package project

import "context"

type Store interface {
	Create(context.Context, CreateInput) (string, error)
	Get(context.Context, string) (Result, error)
	Update(context.Context, UpdateInput) error
	Delete(context.Context, string) error
}
