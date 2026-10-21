package usecases

import (
	"context"
	"testing"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/repository/link"
)

type repositoryStoreStub struct{ attached bool }

func (*repositoryStoreStub) Resolve(context.Context, string, string) (link.Result, error) {
	return link.Result{RepositoryID: "R_1"}, nil
}

func (store *repositoryStoreStub) Attach(_ context.Context, projectID, repositoryID string) (link.Result, error) {
	store.attached = true
	return link.Result{ProjectID: projectID, RepositoryID: repositoryID}, nil
}
func (*repositoryStoreStub) Detach(context.Context, string, string) error { return nil }

func TestAttachReturnsMutationResult(t *testing.T) {
	store := &repositoryStoreStub{}
	result, err := NewAttach(store).Run(t.Context(), link.AttachInput{ProjectID: "PVT_1", Owner: "atls", Name: "planning"})
	if err != nil {
		t.Fatalf("attaching repository: %v", err)
	}
	if result.RepositoryID != "R_1" || !store.attached {
		t.Fatalf("unexpected orchestration: %#v", result)
	}
}
