package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/field"
)

type Update struct{ store field.Store }

func NewUpdate(store field.Store) *Update { return &Update{store: store} }

func (useCase *Update) Run(ctx context.Context, input field.UpdateInput) (field.Result, error) {
	if err := useCase.store.Update(ctx, input); err != nil {
		return field.Result{}, err
	}
	return useCase.store.Get(ctx, input.ID)
}
