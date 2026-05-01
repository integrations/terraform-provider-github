package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestBuildOrganizationSettings_OmittedInternalReposNotSent(t *testing.T) {
	resource := resourceGithubOrganizationSettings()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"billing_email": "test@example.com",
	})

	settings := buildOrganizationSettings(d, true)

	if settings.MembersCanCreateInternalRepos != nil {
		t.Error("MembersCanCreateInternalRepos should be nil when not configured")
	}
}

func TestBuildOrganizationSettings_ExplicitTrueInternalReposSent(t *testing.T) {
	resource := resourceGithubOrganizationSettings()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"billing_email":                           "test@example.com",
		"members_can_create_internal_repositories": true,
	})

	settings := buildOrganizationSettings(d, true)

	if settings.MembersCanCreateInternalRepos == nil {
		t.Fatal("MembersCanCreateInternalRepos should not be nil when explicitly set to true")
	}
	if *settings.MembersCanCreateInternalRepos != true {
		t.Errorf("MembersCanCreateInternalRepos = %v, want true", *settings.MembersCanCreateInternalRepos)
	}
}

func TestBuildOrganizationSettings_ExplicitFalseInternalReposSent(t *testing.T) {
	resource := resourceGithubOrganizationSettings()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"billing_email":                           "test@example.com",
		"members_can_create_internal_repositories": false,
	})

	settings := buildOrganizationSettings(d, true)

	if settings.MembersCanCreateInternalRepos == nil {
		t.Fatal("MembersCanCreateInternalRepos should not be nil when explicitly set to false")
	}
	if *settings.MembersCanCreateInternalRepos != false {
		t.Errorf("MembersCanCreateInternalRepos = %v, want false", *settings.MembersCanCreateInternalRepos)
	}
}

func TestBuildOrganizationSettings_UpdateOmittedInternalReposNotSent(t *testing.T) {
	resource := resourceGithubOrganizationSettings()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"billing_email": "test@example.com",
	})
	d.SetId("test-org")

	settings := buildOrganizationSettings(d, true)

	if settings.MembersCanCreateInternalRepos != nil {
		t.Error("MembersCanCreateInternalRepos should be nil on update when field has not changed")
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

func TestOrganizationSettingsSchemaInternalRepos(t *testing.T) {
	resource := resourceGithubOrganizationSettings()

	field := resource.Schema["members_can_create_internal_repositories"]
	if field == nil {
		t.Fatal("members_can_create_internal_repositories not found in schema")
	}

	if !field.Optional {
		t.Error("members_can_create_internal_repositories should be Optional")
	}
	if !field.Computed {
		t.Error("members_can_create_internal_repositories should be Computed")
	}
	if field.Default != nil {
		t.Errorf("members_can_create_internal_repositories should have no Default, got %v", field.Default)
	}
}
