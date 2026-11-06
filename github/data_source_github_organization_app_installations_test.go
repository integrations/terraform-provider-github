package github

import (
	"maps"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

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

func TestFlattenInstallationPermissions(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name  string
		perms *github.InstallationPermissions
		want  map[string]string
	}{
		{name: "nil", perms: nil, want: map[string]string{}},
		{name: "empty", perms: &github.InstallationPermissions{}, want: map[string]string{}},
		{
			name: "keeps permissions outside the old hand-picked subset",
			perms: &github.InstallationPermissions{
				Contents:                   new("read"),
				OrganizationAdministration: new("write"),
				Members:                    new("read"),
				OrganizationCustomRoles:    new("admin"),
			},
			want: map[string]string{
				"contents":                    "read",
				"organization_administration": "write",
				"members":                     "read",
				"organization_custom_roles":   "admin",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := flattenInstallationPermissions(tc.perms)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !maps.Equal(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}
