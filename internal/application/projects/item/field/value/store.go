package value

import "context"

type Reader interface {
	Get(context.Context, string, string) (Result, error)
}

type Writer interface {
	Set(context.Context, SetInput) (Result, error)
	Clear(context.Context, string, string, string) error
}

type Store interface {
	Reader
	Writer
}
