package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/repository/link"
)

type Resolve struct{ resolver link.Resolver }

func NewResolve(resolver link.Resolver) *Resolve { return &Resolve{resolver: resolver} }

func (useCase *Resolve) Run(ctx context.Context, owner, name string) (link.Result, error) {
	return useCase.resolver.Resolve(ctx, owner, name)
}
