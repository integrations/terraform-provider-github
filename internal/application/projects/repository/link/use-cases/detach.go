package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/repository/link"
)

type Detach struct{ store link.Store }

func NewDetach(store link.Store) *Detach { return &Detach{store: store} }

func (useCase *Detach) Run(ctx context.Context, projectID, repositoryID string) error {
	return useCase.store.Detach(ctx, projectID, repositoryID)
}
