package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item/field/value"
)

type Set struct{ store value.Writer }

func NewSet(store value.Writer) *Set { return &Set{store: store} }

func (useCase *Set) Run(ctx context.Context, input value.SetInput) (value.Result, error) {
	return useCase.store.Set(ctx, input)
}
