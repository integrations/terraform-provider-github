package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/project"
)

type Update struct {
	store project.Store
}

func NewUpdate(store project.Store) *Update {
	return &Update{store: store}
}

func (useCase *Update) Run(ctx context.Context, input project.UpdateInput) error {
	return useCase.store.Update(ctx, input)
}
