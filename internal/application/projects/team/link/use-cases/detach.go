package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/team/link"
)

type Detach struct{ store link.Store }

func NewDetach(store link.Store) *Detach { return &Detach{store: store} }

func (useCase *Detach) Run(ctx context.Context, projectID, teamID string) error {
	return useCase.store.Detach(ctx, projectID, teamID)
}
