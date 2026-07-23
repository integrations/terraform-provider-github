package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TestBuildOrganizationSettingsCreateIncludesFalseBool verifies that a boolean
// attribute explicitly configured as false is included in the create payload.
// Regression test for the create path dropping false booleans (only fixed by a
// second apply through the update/HasChange path).
func TestBuildOrganizationSettingsCreateIncludesFalseBool(t *testing.T) {
	t.Parallel()

	r := resourceGithubOrganizationSettings()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
		"billing_email":               "org@example.com",
		"has_organization_projects":   false,
		"web_commit_signoff_required": false,
	})
	// TestResourceDataRaw leaves the ID empty, so this exercises the create path.

	settings := buildOrganizationSettings(d, false)

	if settings.HasOrganizationProjects == nil {
		t.Fatal("has_organization_projects set to false was dropped on create; expected it to be included")
	}
	if got := settings.GetHasOrganizationProjects(); got != false {
		t.Errorf("has_organization_projects = %v, want false", got)
	}

	if settings.WebCommitSignoffRequired == nil {
		t.Fatal("web_commit_signoff_required set to false was dropped on create; expected it to be included")
	}
	if got := settings.GetWebCommitSignoffRequired(); got != false {
		t.Errorf("web_commit_signoff_required = %v, want false", got)
	}
}

// TestBuildOrganizationSettingsCreateIncludesTrueBool verifies a bool set to
// true is still included on create.
func TestBuildOrganizationSettingsCreateIncludesTrueBool(t *testing.T) {
	t.Parallel()

	r := resourceGithubOrganizationSettings()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
		"billing_email":               "org@example.com",
		"web_commit_signoff_required": true,
	})

	settings := buildOrganizationSettings(d, false)

	if settings.WebCommitSignoffRequired == nil {
		t.Fatal("web_commit_signoff_required set to true was dropped on create")
	}
	if got := settings.GetWebCommitSignoffRequired(); got != true {
		t.Errorf("web_commit_signoff_required = %v, want true", got)
	}
}

// TestBuildOrganizationSettingsCreateStringUnsetOmitted verifies the fix does
// not change string handling: an unconfigured optional string stays omitted on
// create (only booleans are always included).
func TestBuildOrganizationSettingsCreateStringUnsetOmitted(t *testing.T) {
	t.Parallel()

	r := resourceGithubOrganizationSettings()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
		"billing_email": "org@example.com",
	})

	settings := buildOrganizationSettings(d, false)

	if settings.Company != nil {
		t.Errorf("company was not configured but got included as %q", settings.GetCompany())
	}
	if settings.BillingEmail == nil || settings.GetBillingEmail() != "org@example.com" {
		t.Errorf("billing_email = %v, want org@example.com", settings.BillingEmail)
	}
}

// TestBuildOrganizationSettingsUpdateGatesOnHasChange verifies the update path
// is unaffected by the create-path fix: with an ID set and no diff, unchanged
// boolean fields are still omitted (only changed fields are included on update).
func TestBuildOrganizationSettingsUpdateGatesOnHasChange(t *testing.T) {
	t.Parallel()

	r := resourceGithubOrganizationSettings()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
		"billing_email":             "org@example.com",
		"has_organization_projects": false,
	})
	d.SetId("example-org") // non-empty ID exercises the update path

	settings := buildOrganizationSettings(d, false)

	if settings.HasOrganizationProjects != nil {
		t.Error("has_organization_projects with no change should be omitted on update")
	}
}

// TestBuildOrganizationSettingsCreateEnterpriseBool verifies the enterprise-only
// boolean is included on create only when isEnterprise is true.
func TestBuildOrganizationSettingsCreateEnterpriseBool(t *testing.T) {
	t.Parallel()

	r := resourceGithubOrganizationSettings()

	dEnterprise := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
		"billing_email": "org@example.com",
		"members_can_create_internal_repositories": false,
	})
	if got := buildOrganizationSettings(dEnterprise, true); got.MembersCanCreateInternalRepos == nil {
		t.Fatal("members_can_create_internal_repositories set to false was dropped on enterprise create")
	}

	dNonEnterprise := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
		"billing_email": "org@example.com",
		"members_can_create_internal_repositories": false,
	})
	if got := buildOrganizationSettings(dNonEnterprise, false); got.MembersCanCreateInternalRepos != nil {
		t.Error("members_can_create_internal_repositories should not be set when isEnterprise is false")
	}
}
