package usecases

import (
	"context"
	"testing"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item/field/value"
)

type valueStoreStub struct{ set, loaded bool }

func (store *valueStoreStub) Set(context.Context, value.SetInput) error { store.set = true; return nil }
func (store *valueStoreStub) Get(context.Context, string, string) (value.Result, error) {
	store.loaded = true
	return value.Result{Kind: value.KindNumber}, nil
}
func (*valueStoreStub) Clear(context.Context, string, string, string) error { return nil }

func TestSetReturnsLoadedValue(t *testing.T) {
	store := &valueStoreStub{}
	result, err := NewSet(store).Run(t.Context(), value.SetInput{ItemID: "PVTI_1", FieldID: "PVTF_1"})
	if err != nil {
		t.Fatalf("setting value: %v", err)
	}
	if result.Kind != value.KindNumber || !store.set || !store.loaded {
		t.Fatalf("unexpected orchestration: %#v", result)
	}
}
