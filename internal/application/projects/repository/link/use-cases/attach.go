package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/repository/link"
)

type attachStore interface {
	link.Resolver
	link.Writer
}

type Attach struct{ store attachStore }

func NewAttach(store attachStore) *Attach { return &Attach{store: store} }

func (useCase *Attach) Run(ctx context.Context, input link.AttachInput) (link.Result, error) {
	repository, err := useCase.store.Resolve(ctx, input.Owner, input.Name)
	if err != nil {
		return link.Result{}, err
	}
	return useCase.store.Attach(ctx, input.ProjectID, repository.RepositoryID)
}
