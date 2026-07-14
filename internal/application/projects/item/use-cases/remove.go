package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item"
)

type Remove struct{ store item.Store }

func NewRemove(store item.Store) *Remove { return &Remove{store: store} }

func (useCase *Remove) Run(ctx context.Context, projectID, itemID string) error {
	return useCase.store.Remove(ctx, projectID, itemID)
}
