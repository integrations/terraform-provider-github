package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item"
)

type Add struct{ store item.Store }

func NewAdd(store item.Store) *Add { return &Add{store: store} }

func (useCase *Add) Run(ctx context.Context, input item.AddInput) (item.Result, error) {
	id, err := useCase.store.Add(ctx, input.ProjectID, input.ContentID)
	if err != nil {
		return item.Result{}, err
	}
	if input.Archived {
		if err := useCase.store.SetArchived(ctx, input.ProjectID, id, true); err != nil {
			return item.Result{}, err
		}
	}
	return useCase.store.Get(ctx, id)
}
