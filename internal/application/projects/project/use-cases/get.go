package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/project"
)

type Get struct {
	store project.Store
}

func NewGet(store project.Store) *Get {
	return &Get{store: store}
}

func (useCase *Get) Run(ctx context.Context, id string) (project.Result, error) {
	return useCase.store.Get(ctx, id)
}
