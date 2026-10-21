package usecases

import (
	"context"
	"testing"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item/field/value"
)

type valueStoreStub struct{ set bool }

func (store *valueStoreStub) Set(_ context.Context, input value.SetInput) (value.Result, error) {
	store.set = true
	return input.Value, nil
}
func (*valueStoreStub) Clear(context.Context, string, string, string) error { return nil }

func TestSetReturnsMutationInput(t *testing.T) {
	store := &valueStoreStub{}
	result, err := NewSet(store).Run(t.Context(), value.SetInput{ItemID: "PVTI_1", FieldID: "PVTF_1", Value: value.Result{Kind: value.KindNumber, Number: 3}})
	if err != nil {
		t.Fatalf("setting value: %v", err)
	}
	if result.Kind != value.KindNumber || result.Number != 3 || !store.set {
		t.Fatalf("unexpected orchestration: %#v", result)
	}
}
