package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/team/link"
)

type Attach struct{ store link.Store }

func NewAttach(store link.Store) *Attach { return &Attach{store: store} }

func (useCase *Attach) Run(ctx context.Context, input link.AttachInput) (link.Result, error) {
	team, err := useCase.store.Resolve(ctx, input.Organization, input.Slug)
	if err != nil {
		return link.Result{}, err
	}
	if err := useCase.store.Attach(ctx, input.ProjectID, team.TeamID); err != nil {
		return link.Result{}, err
	}
	return useCase.store.Get(ctx, input.ProjectID, team.TeamID)
}
