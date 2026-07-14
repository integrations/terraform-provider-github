package link

import "context"

type Store interface {
	Resolve(context.Context, string, string) (Result, error)
	Attach(context.Context, string, string) error
	Get(context.Context, string, string) (Result, error)
	Detach(context.Context, string, string) error
}
