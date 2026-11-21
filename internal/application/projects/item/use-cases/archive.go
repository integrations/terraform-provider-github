package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item"
)

type Archive struct{ store item.Writer }

func NewArchive(store item.Writer) *Archive { return &Archive{store: store} }

func (useCase *Archive) Run(ctx context.Context, projectID, itemID string) (item.Result, error) {
	return useCase.store.SetArchived(ctx, projectID, itemID, true)
}
