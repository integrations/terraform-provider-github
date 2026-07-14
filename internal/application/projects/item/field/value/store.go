package value

import "context"

type Store interface {
	Set(context.Context, SetInput) error
	Get(context.Context, string, string) (Result, error)
	Clear(context.Context, string, string, string) error
}
