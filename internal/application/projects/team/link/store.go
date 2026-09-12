package link

import "context"

type Resolver interface {
	Resolve(context.Context, string, string) (Result, error)
}

type Reader interface {
	Get(context.Context, string, string) (Result, error)
}

type Writer interface {
	Attach(context.Context, string, string) (Result, error)
	Detach(context.Context, string, string) error
}

type Store interface {
	Resolver
	Reader
	Writer
}
