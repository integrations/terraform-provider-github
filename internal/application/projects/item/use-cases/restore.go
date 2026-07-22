package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item"
)

type Restore struct{ store item.Writer }

func NewRestore(store item.Writer) *Restore { return &Restore{store: store} }

func (useCase *Restore) Run(ctx context.Context, projectID, itemID string) (item.Result, error) {
	return useCase.store.SetArchived(ctx, projectID, itemID, false)
}
