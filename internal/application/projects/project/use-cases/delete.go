package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/project"
)

type Delete struct {
	store project.Writer
}

func NewDelete(store project.Writer) *Delete {
	return &Delete{store: store}
}

func (useCase *Delete) Run(ctx context.Context, id string) error {
	return useCase.store.Delete(ctx, id)
}
