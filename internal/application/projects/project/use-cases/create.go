package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/project"
)

type Create struct {
	store project.Store
}

func NewCreate(store project.Store) *Create {
	return &Create{store: store}
}

func (useCase *Create) Run(ctx context.Context, input project.CreateInput) (project.Result, error) {
	id, err := useCase.store.Create(ctx, input)
	if err != nil {
		return project.Result{}, err
	}
	return useCase.store.Get(ctx, id)
}
