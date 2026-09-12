package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/field"
)

type Get struct{ store field.Reader }

func NewGet(store field.Reader) *Get { return &Get{store: store} }

func (useCase *Get) Run(ctx context.Context, id string) (field.Result, error) {
	return useCase.store.Get(ctx, id)
}
