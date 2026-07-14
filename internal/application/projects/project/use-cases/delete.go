package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/project"
)

type Delete struct {
	store project.Store
}

func NewDelete(store project.Store) *Delete {
	return &Delete{store: store}
}

func (useCase *Delete) Run(ctx context.Context, id string) error {
	return useCase.store.Delete(ctx, id)
}
