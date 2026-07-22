package usecases

import (
	"context"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/field"
)

type Update struct{ store field.Writer }

func NewUpdate(store field.Writer) *Update { return &Update{store: store} }

func (useCase *Update) Run(ctx context.Context, input field.UpdateInput) (field.Result, error) {
	return useCase.store.Update(ctx, input)
}
