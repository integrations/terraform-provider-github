package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item"
)

type Archive struct{ store item.Store }

func NewArchive(store item.Store) *Archive { return &Archive{store: store} }

func (useCase *Archive) Run(ctx context.Context, projectID, itemID string) (item.Result, error) {
	if err := useCase.store.SetArchived(ctx, projectID, itemID, true); err != nil {
		return item.Result{}, err
	}
	return useCase.store.Get(ctx, itemID)
}
