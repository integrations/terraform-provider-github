package usecases

import (
	"context"
	"testing"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/team/link"
)

type teamStoreStub struct{ attached, loaded bool }

func (*teamStoreStub) Resolve(context.Context, string, string) (link.Result, error) {
	return link.Result{TeamID: "T_1"}, nil
}

func (store *teamStoreStub) Attach(context.Context, string, string) error {
	store.attached = true
	return nil
}

func (store *teamStoreStub) Get(context.Context, string, string) (link.Result, error) {
	store.loaded = true
	return link.Result{ProjectID: "PVT_1", TeamID: "T_1"}, nil
}
func (*teamStoreStub) Detach(context.Context, string, string) error { return nil }

func TestAttachResolvesAndLoadsLink(t *testing.T) {
	store := &teamStoreStub{}
	result, err := NewAttach(store).Run(t.Context(), link.AttachInput{ProjectID: "PVT_1", Organization: "atls", Slug: "platform"})
	if err != nil {
		t.Fatalf("attaching team: %v", err)
	}
	if result.TeamID != "T_1" || !store.attached || !store.loaded {
		t.Fatalf("unexpected orchestration: %#v", result)
	}
}
