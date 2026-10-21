package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects"
	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item"
)

type Add struct{ store item.Writer }

func NewAdd(store item.Writer) *Add { return &Add{store: store} }

func (useCase *Add) Run(ctx context.Context, input item.AddInput) (item.Result, error) {
	created, err := useCase.store.Add(ctx, input.ProjectID, input.ContentID)
	if err != nil {
		return item.Result{}, err
	}
	if !input.Archived {
		return created, nil
	}

	archived, err := useCase.store.SetArchived(ctx, input.ProjectID, created.ID, true)
	if err == nil {
		return archived, nil
	}

	cleanupErr := useCase.store.Remove(ctx, input.ProjectID, created.ID)
	if cleanupErr == nil || errors.Is(cleanupErr, projects.ErrNotFound) {
		return item.Result{}, fmt.Errorf("archiving created project item: %w", err)
	}

	return created, errors.Join(
		fmt.Errorf("archiving created project item: %w", err),
		fmt.Errorf("removing partially created project item %q: %w", created.ID, cleanupErr),
	)
}
