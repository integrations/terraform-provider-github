package project

import (
	"testing"

	application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/project"
)

func TestResultFromNodeMapsOrganizationOwner(t *testing.T) {
	value := node{ID: "PVT_1"}
	value.Owner.Typename = "Organization"
	value.Owner.Organization.DatabaseID = 101
	value.Owner.Organization.Login = "atls"
	result, err := resultFromNode(value)
	if err != nil {
		t.Fatalf("mapping project: %v", err)
	}
	if result.ID != "PVT_1" || result.OwnerKind != application.OwnerOrganization || result.Owner != "atls" || result.OwnerID != 101 {
		t.Fatalf("unexpected project mapping: %#v", result)
	}
}

func TestResultFromNodeRejectsUnsupportedOwner(t *testing.T) {
	value := node{ID: "PVT_1"}
	value.Owner.Typename = "Issue"
	if _, err := resultFromNode(value); err == nil {
		t.Fatal("unsupported project owner was accepted")
	}
}
