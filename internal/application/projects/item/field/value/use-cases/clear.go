package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item/field/value"
)

type Clear struct{ store value.Store }

func NewClear(store value.Store) *Clear { return &Clear{store: store} }

func (useCase *Clear) Run(ctx context.Context, projectID, itemID, fieldID string) error {
	return useCase.store.Clear(ctx, projectID, itemID, fieldID)
}
