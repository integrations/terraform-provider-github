package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects"
	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/project"
)

type projectStoreStub struct {
	created, updated, deleted bool
	updateErr, deleteErr      error
}

func (store *projectStoreStub) Create(context.Context, project.CreateInput) (project.Result, error) {
	store.created = true
	return project.Result{ID: "PVT_1"}, nil
}

func (store *projectStoreStub) Update(context.Context, project.UpdateInput) (project.Result, error) {
	store.updated = true
	return project.Result{ID: "PVT_1", Title: "Planning"}, store.updateErr
}

func (store *projectStoreStub) Delete(context.Context, string) error {
	store.deleted = true
	return store.deleteErr
}

func TestCreateReturnsUpdateMutationResult(t *testing.T) {
	store := &projectStoreStub{}
	result, err := NewCreate(store).Run(t.Context(), project.CreateInput{Title: "Planning"})
	if err != nil {
		t.Fatalf("creating project: %v", err)
	}
	if result.ID != "PVT_1" || result.Title != "Planning" || !store.created || !store.updated || store.deleted {
		t.Fatalf("unexpected orchestration: %#v", result)
	}
}

func TestCreateDeletesProjectWhenConfigurationFails(t *testing.T) {
	store := &projectStoreStub{updateErr: errors.New("update failed")}
	result, err := NewCreate(store).Run(t.Context(), project.CreateInput{Title: "Planning"})
	if err == nil || result.ID != "" || !store.deleted {
		t.Fatalf("expected compensated failure, result=%#v err=%v", result, err)
	}
}

func TestCreateTreatsMissingPartialProjectAsCompensated(t *testing.T) {
	store := &projectStoreStub{updateErr: errors.New("update failed"), deleteErr: projects.ErrNotFound}
	result, err := NewCreate(store).Run(t.Context(), project.CreateInput{Title: "Planning"})
	if err == nil || result.ID != "" || !store.deleted {
		t.Fatalf("expected compensated not-found, result=%#v err=%v", result, err)
	}
}

func TestCreatePreservesProjectWhenCompensationFails(t *testing.T) {
	store := &projectStoreStub{updateErr: errors.New("update failed"), deleteErr: errors.New("delete failed")}
	result, err := NewCreate(store).Run(t.Context(), project.CreateInput{Title: "Planning"})
	if err == nil || result.ID != "PVT_1" || !errors.Is(err, store.updateErr) || !errors.Is(err, store.deleteErr) {
		t.Fatalf("expected partial state and joined error, result=%#v err=%v", result, err)
	}
}
