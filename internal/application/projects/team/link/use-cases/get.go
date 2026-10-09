package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/team/link"
)

type Get struct{ store link.Reader }

func NewGet(store link.Reader) *Get { return &Get{store: store} }

func (useCase *Get) Run(ctx context.Context, projectID, teamID string) (link.Result, error) {
	return useCase.store.Get(ctx, projectID, teamID)
}
