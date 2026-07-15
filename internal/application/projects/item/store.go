package item

import "context"

type Reader interface {
	Get(context.Context, string) (Result, error)
}

type Writer interface {
	Add(context.Context, string, string) (Result, error)
	SetArchived(context.Context, string, string, bool) (Result, error)
	Remove(context.Context, string, string) error
}

type Store interface {
	Reader
	Writer
}
