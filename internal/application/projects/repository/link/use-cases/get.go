package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/repository/link"
)

type Get struct{ store link.Store }

func NewGet(store link.Store) *Get { return &Get{store: store} }

func (useCase *Get) Run(ctx context.Context, projectID, repositoryID string) (link.Result, error) {
	return useCase.store.Get(ctx, projectID, repositoryID)
}
