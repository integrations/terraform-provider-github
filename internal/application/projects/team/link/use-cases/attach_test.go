package usecases

import (
	"context"
	"testing"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/team/link"
)

type teamStoreStub struct{ attached bool }

func (*teamStoreStub) Resolve(context.Context, string, string) (link.Result, error) {
	return link.Result{TeamID: "T_1"}, nil
}

func (store *teamStoreStub) Attach(_ context.Context, projectID, teamID string) (link.Result, error) {
	store.attached = true
	return link.Result{ProjectID: projectID, TeamID: teamID}, nil
}
func (*teamStoreStub) Detach(context.Context, string, string) error { return nil }

func TestAttachReturnsMutationResult(t *testing.T) {
	store := &teamStoreStub{}
	result, err := NewAttach(store).Run(t.Context(), link.AttachInput{ProjectID: "PVT_1", Organization: "atls", Slug: "platform"})
	if err != nil {
		t.Fatalf("attaching team: %v", err)
	}
	if result.TeamID != "T_1" || !store.attached {
		t.Fatalf("unexpected orchestration: %#v", result)
	}
}
