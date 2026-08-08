package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/field"
)

type Delete struct{ store field.Writer }

func NewDelete(store field.Writer) *Delete { return &Delete{store: store} }

func (useCase *Delete) Run(ctx context.Context, id string) error {
	return useCase.store.Delete(ctx, id)
}
