package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/field"
)

type Create struct{ store field.Store }

func NewCreate(store field.Store) *Create { return &Create{store: store} }

func (useCase *Create) Run(ctx context.Context, input field.CreateInput) (field.Result, error) {
	id, err := useCase.store.Create(ctx, input)
	if err != nil {
		return field.Result{}, err
	}
	return useCase.store.Get(ctx, id)
}
