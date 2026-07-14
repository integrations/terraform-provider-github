package usecases

import (
	"context"
	"testing"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects/project"
)

type projectStoreStub struct{ created, loaded bool }

func (store *projectStoreStub) Create(context.Context, project.CreateInput) (string, error) {
	store.created = true
	return "PVT_1", nil
}

func (store *projectStoreStub) Get(context.Context, string) (project.Result, error) {
	store.loaded = true
	return project.Result{ID: "PVT_1"}, nil
}
func (*projectStoreStub) Update(context.Context, project.UpdateInput) error { return nil }
func (*projectStoreStub) Delete(context.Context, string) error              { return nil }

func TestCreateReturnsLoadedProject(t *testing.T) {
	store := &projectStoreStub{}
	result, err := NewCreate(store).Run(t.Context(), project.CreateInput{Title: "Planning"})
	if err != nil {
		t.Fatalf("creating project: %v", err)
	}
	if result.ID != "PVT_1" || !store.created || !store.loaded {
		t.Fatalf("unexpected orchestration: %#v", result)
	}
}
