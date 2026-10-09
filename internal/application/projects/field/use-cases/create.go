package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/field"
)

type Create struct{ store field.Writer }

func NewCreate(store field.Writer) *Create { return &Create{store: store} }

func (useCase *Create) Run(ctx context.Context, input field.CreateInput) (field.Result, error) {
	return useCase.store.Create(ctx, input)
}
