package item

import "context"

type Store interface {
	Add(context.Context, string, string) (string, error)
	Get(context.Context, string) (Result, error)
	SetArchived(context.Context, string, string, bool) error
	Remove(context.Context, string, string) error
}
