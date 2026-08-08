package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item"
)

type Get struct{ store item.Reader }

func NewGet(store item.Reader) *Get { return &Get{store: store} }

func (useCase *Get) Run(ctx context.Context, id string) (item.Result, error) {
	return useCase.store.Get(ctx, id)
}
