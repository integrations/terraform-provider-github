package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects"
	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/item"
)

type itemStoreStub struct {
	added, archived, removed bool
	archiveErr, removeErr    error
}

func (store *itemStoreStub) Add(_ context.Context, projectID, contentID string) (item.Result, error) {
	store.added = true
	return item.Result{ID: "PVTI_1", ProjectID: projectID, ContentID: contentID}, nil
}

func (store *itemStoreStub) SetArchived(_ context.Context, projectID, itemID string, archived bool) (item.Result, error) {
	store.archived = archived
	return item.Result{ID: itemID, ProjectID: projectID, ContentID: "I_1", Archived: archived}, store.archiveErr
}

func (store *itemStoreStub) Remove(context.Context, string, string) error {
	store.removed = true
	return store.removeErr
}

func TestAddReturnsArchiveMutationResult(t *testing.T) {
	store := &itemStoreStub{}
	result, err := NewAdd(store).Run(t.Context(), item.AddInput{ProjectID: "PVT_1", ContentID: "I_1", Archived: true})
	if err != nil {
		t.Fatalf("adding item: %v", err)
	}
	if !result.Archived || !store.added || store.removed {
		t.Fatalf("unexpected orchestration: %#v", result)
	}
}

func TestAddRemovesItemWhenArchiveFails(t *testing.T) {
	store := &itemStoreStub{archiveErr: errors.New("archive failed")}
	result, err := NewAdd(store).Run(t.Context(), item.AddInput{ProjectID: "PVT_1", ContentID: "I_1", Archived: true})
	if err == nil || result.ID != "" || !store.removed {
		t.Fatalf("expected compensated failure, result=%#v err=%v", result, err)
	}
}

func TestAddTreatsMissingPartialItemAsCompensated(t *testing.T) {
	store := &itemStoreStub{archiveErr: errors.New("archive failed"), removeErr: projects.ErrNotFound}
	result, err := NewAdd(store).Run(t.Context(), item.AddInput{ProjectID: "PVT_1", ContentID: "I_1", Archived: true})
	if err == nil || result.ID != "" || !store.removed {
		t.Fatalf("expected compensated not-found, result=%#v err=%v", result, err)
	}
}

func TestAddPreservesItemWhenCompensationFails(t *testing.T) {
	store := &itemStoreStub{archiveErr: errors.New("archive failed"), removeErr: errors.New("remove failed")}
	result, err := NewAdd(store).Run(t.Context(), item.AddInput{ProjectID: "PVT_1", ContentID: "I_1", Archived: true})
	if err == nil || result.ID != "PVTI_1" || !errors.Is(err, store.archiveErr) || !errors.Is(err, store.removeErr) {
		t.Fatalf("expected partial state and joined error, result=%#v err=%v", result, err)
	}
}
