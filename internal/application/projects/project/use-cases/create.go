package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects"
	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/project"
)

type Create struct {
	store project.Writer
}

func NewCreate(store project.Writer) *Create {
	return &Create{store: store}
}

func (useCase *Create) Run(ctx context.Context, input project.CreateInput) (project.Result, error) {
	created, err := useCase.store.Create(ctx, input)
	if err != nil {
		return project.Result{}, err
	}
	updated, err := useCase.store.Update(ctx, project.UpdateInput{
		ID:               created.ID,
		Title:            input.Title,
		ShortDescription: input.ShortDescription,
		Readme:           input.Readme,
		Public:           input.Public,
		Closed:           input.Closed,
	})
	if err == nil {
		return updated, nil
	}

	cleanupErr := useCase.store.Delete(ctx, created.ID)
	if cleanupErr == nil || errors.Is(cleanupErr, projects.ErrNotFound) {
		return project.Result{}, fmt.Errorf("configuring created project: %w", err)
	}

	return created, errors.Join(
		fmt.Errorf("configuring created project: %w", err),
		fmt.Errorf("deleting partially created project %q: %w", created.ID, cleanupErr),
	)
}
