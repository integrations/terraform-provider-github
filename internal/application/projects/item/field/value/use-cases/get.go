package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item/field/value"
)

type Get struct{ store value.Reader }

func NewGet(store value.Reader) *Get { return &Get{store: store} }

func (useCase *Get) Run(ctx context.Context, itemID, fieldID string) (value.Result, error) {
	return useCase.store.Get(ctx, itemID, fieldID)
}
