package usecases

import (
	"context"
	"testing"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item"
)

type itemStoreStub struct{ added, archived, loaded bool }

func (store *itemStoreStub) Add(context.Context, string, string) (string, error) {
	store.added = true
	return "PVTI_1", nil
}

func (store *itemStoreStub) Get(context.Context, string) (item.Result, error) {
	store.loaded = true
	return item.Result{ID: "PVTI_1", Archived: store.archived}, nil
}

func (store *itemStoreStub) SetArchived(_ context.Context, _, _ string, archived bool) error {
	store.archived = archived
	return nil
}
func (*itemStoreStub) Remove(context.Context, string, string) error { return nil }

func TestAddAppliesArchivedStateBeforeLoading(t *testing.T) {
	store := &itemStoreStub{}
	result, err := NewAdd(store).Run(t.Context(), item.AddInput{ProjectID: "PVT_1", ContentID: "I_1", Archived: true})
	if err != nil {
		t.Fatalf("adding item: %v", err)
	}
	if !result.Archived || !store.added || !store.loaded {
		t.Fatalf("unexpected orchestration: %#v", result)
	}
}
