package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/project"
)

type Update struct {
	store project.Writer
}

func NewUpdate(store project.Writer) *Update {
	return &Update{store: store}
}

func (useCase *Update) Run(ctx context.Context, input project.UpdateInput) error {
	_, err := useCase.store.Update(ctx, input)
	return err
}
