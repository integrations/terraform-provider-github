package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/team/link"
)

type Resolve struct{ resolver link.Resolver }

func NewResolve(resolver link.Resolver) *Resolve { return &Resolve{resolver: resolver} }

func (useCase *Resolve) Run(ctx context.Context, organization, slug string) (link.Result, error) {
	return useCase.resolver.Resolve(ctx, organization, slug)
}
