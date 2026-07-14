package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item"
)

type Restore struct{ store item.Store }

func NewRestore(store item.Store) *Restore { return &Restore{store: store} }

func (useCase *Restore) Run(ctx context.Context, projectID, itemID string) (item.Result, error) {
	if err := useCase.store.SetArchived(ctx, projectID, itemID, false); err != nil {
		return item.Result{}, err
	}
	return useCase.store.Get(ctx, itemID)
}
