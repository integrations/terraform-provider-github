package usecases

import (
	"context"
	"testing"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/field"
)

type fieldStoreStub struct{ created bool }

func (store *fieldStoreStub) Create(context.Context, field.CreateInput) (field.Result, error) {
	store.created = true
	return field.Result{ID: "PVTF_1", Name: "Estimate"}, nil
}

func (*fieldStoreStub) Update(context.Context, field.UpdateInput) (field.Result, error) {
	return field.Result{}, nil
}
func (*fieldStoreStub) Delete(context.Context, string) error { return nil }

func TestCreateReturnsMutationResult(t *testing.T) {
	store := &fieldStoreStub{}
	result, err := NewCreate(store).Run(t.Context(), field.CreateInput{Name: "Estimate"})
	if err != nil {
		t.Fatalf("creating field: %v", err)
	}
	if result.ID != "PVTF_1" || result.Name != "Estimate" || !store.created {
		t.Fatalf("unexpected orchestration: %#v", result)
	}
}
