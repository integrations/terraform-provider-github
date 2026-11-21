package github

import (
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestFlattenGitHubAppInstallations(t *testing.T) {
	t.Parallel()

	t.Run("skips null installations and keeps the other installations", func(t *testing.T) {
		t.Parallel()

		installations := []*github.Installation{
			{ID: new(int64(25381)), AppSlug: new("github-actions")},
			nil,
			{ID: new(int64(25383)), AppSlug: new("dependabot")},
		}

		gitHubAppInstallations := flattenGitHubAppInstallations(installations)

		if len(gitHubAppInstallations) != 2 {
			t.Fatalf("expected 2 installations, got %d", len(gitHubAppInstallations))
		}
		if gitHubAppInstallations[0]["app_slug"] != "github-actions" {
			t.Errorf("expected installations[0] app_slug github-actions, got %v", gitHubAppInstallations[0]["app_slug"])
		}
		if gitHubAppInstallations[1]["app_slug"] != "dependabot" {
			t.Errorf("expected installations[1] app_slug dependabot, got %v", gitHubAppInstallations[1]["app_slug"])
		}
	})
}

func TestAccGithubOrganizationAppInstallations(t *testing.T) {
	t.Parallel()

	t.Run("basic", func(t *testing.T) {
		t.Parallel()

		config := `data "github_organization_app_installations" "test" {}`

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessHasOrgs(t)
				skipUnlessHasAppInstallations(t)
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_organization_app_installations.test", "installations.0.id"),
						resource.TestCheckResourceAttrSet("data.github_organization_app_installations.test", "installations.0.app_slug"),
						resource.TestCheckResourceAttrSet("data.github_organization_app_installations.test", "installations.0.app_id"),
					),
				},
			},
		})
	})
}
