package usecases

import (
	"context"
	"testing"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/field"
)

type fieldStoreStub struct{ created, loaded bool }

func (store *fieldStoreStub) Create(context.Context, field.CreateInput) (string, error) {
	store.created = true
	return "PVTF_1", nil
}

func (store *fieldStoreStub) Get(context.Context, string) (field.Result, error) {
	store.loaded = true
	return field.Result{ID: "PVTF_1"}, nil
}
func (*fieldStoreStub) Update(context.Context, field.UpdateInput) error { return nil }
func (*fieldStoreStub) Delete(context.Context, string) error            { return nil }

func TestCreateReturnsLoadedField(t *testing.T) {
	store := &fieldStoreStub{}
	result, err := NewCreate(store).Run(t.Context(), field.CreateInput{Name: "Estimate"})
	if err != nil {
		t.Fatalf("creating field: %v", err)
	}
	if result.ID != "PVTF_1" || !store.created || !store.loaded {
		t.Fatalf("unexpected orchestration: %#v", result)
	}
}
