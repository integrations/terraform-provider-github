package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubReleaseResource(t *testing.T) {
	t.Parallel()

	t.Run("with_defaults", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_release" "test" {
  repository = "%s"
  tag_name   = "v1.0.0"
}
`, repo.GetName())

		configLive := fmt.Sprintf(`
resource "github_release" "test" {
  repository = "%s"
  tag_name   = "v1.0.0"
  draft      = false
	prerelease = false
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("release_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("node_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("created_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("published_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("html_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("assets_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("upload_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("tarball_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("zipball_url"), knownvalue.NotNull()),
					},
				},
				{
					ResourceName:            "github_release.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateIdPrefix:     fmt.Sprintf("%s:", repo.GetName()),
					ImportStateVerifyIgnore: []string{"name", "body"},
				},
				{
					Config: configLive,
				},
			},
		})
	})

	t.Run("with_options", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		branchName := mustCreateTestBranch(t, repo)

		config := fmt.Sprintf(`
resource "github_release" "test" {
  repository       = "%s"
  tag_name         = "v1.0.0"
	target_commitish = "%s"
  name             = "My Release"
	body             = "Release notes"
  draft            = false
	prerelease       = false
}
`, repo.GetName(), branchName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("release_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("node_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("created_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("published_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("html_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("assets_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("upload_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("tarball_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_release.test", tfjsonpath.New("zipball_url"), knownvalue.NotNull()),
					},
				},
				{
					ResourceName:            "github_release.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateIdPrefix:     fmt.Sprintf("%s:", repo.GetName()),
					ImportStateVerifyIgnore: []string{"name", "body"},
				},
			},
		})
	})
}
