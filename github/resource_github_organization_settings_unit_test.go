package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestBuildOrganizationSettings_OmittedFieldsNotSent(t *testing.T) {
	resource := resourceGithubOrganizationSettings()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"billing_email": "test@example.com",
	})

	settings := buildOrganizationSettings(d, false)

	if settings.MembersCanForkPrivateRepos != nil {
		t.Error("MembersCanForkPrivateRepos should be nil when not configured")
	}
	if settings.BillingEmail == nil {
		t.Error("BillingEmail should be set when configured")
	}
}

func TestBuildOrganizationSettings_ExplicitTrueFieldsSent(t *testing.T) {
	resource := resourceGithubOrganizationSettings()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"billing_email":                        "test@example.com",
		"members_can_fork_private_repositories": true,
	})

	settings := buildOrganizationSettings(d, false)

	if settings.MembersCanForkPrivateRepos == nil {
		t.Fatal("MembersCanForkPrivateRepos should not be nil when explicitly set to true")
	}
	if *settings.MembersCanForkPrivateRepos != true {
		t.Errorf("MembersCanForkPrivateRepos = %v, want true", *settings.MembersCanForkPrivateRepos)
	}
}

func TestBuildOrganizationSettings_ExplicitFalseFieldsSent(t *testing.T) {
	resource := resourceGithubOrganizationSettings()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"billing_email":                        "test@example.com",
		"members_can_fork_private_repositories": false,
	})

	settings := buildOrganizationSettings(d, false)

	if settings.MembersCanForkPrivateRepos == nil {
		t.Fatal("MembersCanForkPrivateRepos should not be nil when explicitly set to false")
	}
	if *settings.MembersCanForkPrivateRepos != false {
		t.Errorf("MembersCanForkPrivateRepos = %v, want false", *settings.MembersCanForkPrivateRepos)
	}
}

func TestBuildOrganizationSettings_UpdateOmittedFieldsNotSent(t *testing.T) {
	resource := resourceGithubOrganizationSettings()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"billing_email": "test@example.com",
	})
	d.SetId("test-org")

	settings := buildOrganizationSettings(d, false)

	if settings.MembersCanForkPrivateRepos != nil {
		t.Error("MembersCanForkPrivateRepos should be nil on update when field has not changed")
	}
}

func TestBuildOrganizationSettings_NonEnterpriseExcludesInternalRepos(t *testing.T) {
	resource := resourceGithubOrganizationSettings()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"billing_email":                           "test@example.com",
		"members_can_create_internal_repositories": true,
	})

	settings := buildOrganizationSettings(d, false)

	if settings.MembersCanCreateInternalRepos != nil {
		t.Error("MembersCanCreateInternalRepos should be nil when not enterprise")
	}
}

func TestOrganizationSettingsSchemaProperties(t *testing.T) {
	resource := resourceGithubOrganizationSettings()

	field := resource.Schema["members_can_fork_private_repositories"]
	if field == nil {
		t.Fatal("members_can_fork_private_repositories not found in schema")
	}

	if !field.Optional {
		t.Error("members_can_fork_private_repositories should be Optional")
	}
	if !field.Computed {
		t.Error("members_can_fork_private_repositories should be Computed")
	}
	if field.Default != nil {
		t.Errorf("members_can_fork_private_repositories should have no Default, got %v", field.Default)
	}
}
