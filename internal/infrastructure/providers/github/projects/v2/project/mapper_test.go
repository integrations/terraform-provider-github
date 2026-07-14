package project

import (
	"testing"

	application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/project"
)

func TestResultFromNodeMapsOrganizationOwner(t *testing.T) {
	value := node{ID: "PVT_1"}
	value.Owner.Typename = "Organization"
	value.Owner.Organization.Login = "atls"
	result := resultFromNode(value)
	if result.ID != "PVT_1" || result.OwnerKind != application.OwnerOrganization || result.Owner != "atls" {
		t.Fatalf("unexpected project mapping: %#v", result)
	}
}
