package graphql

import (
	"fmt"
	"strings"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects"
)

func Error(operation string, err error) error {
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), "Could not resolve to a node with the global id") {
		return fmt.Errorf("%s: %w: %w", operation, projects.ErrNotFound, err)
	}
	return fmt.Errorf("%s: %w", operation, err)
}
