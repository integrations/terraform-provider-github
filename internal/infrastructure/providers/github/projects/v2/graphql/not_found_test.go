package graphql

import (
	"errors"
	"fmt"
	"testing"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects"
)

func TestErrorMapsGlobalNodeLookupFailure(t *testing.T) {
	err := Error("querying project", fmt.Errorf("Could not resolve to a node with the global id"))
	if !errors.Is(err, projects.ErrNotFound) {
		t.Fatalf("expected not found error, got %v", err)
	}
}
