package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item/field/value"
)

type Set struct{ store value.Store }

func NewSet(store value.Store) *Set { return &Set{store: store} }

func (useCase *Set) Run(ctx context.Context, input value.SetInput) (value.Result, error) {
	if err := useCase.store.Set(ctx, input); err != nil {
		return value.Result{}, err
	}
	return useCase.store.Get(ctx, input.ItemID, input.FieldID)
}
