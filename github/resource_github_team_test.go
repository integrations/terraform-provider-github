package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubTeam(t *testing.T) {
	t.Parallel()

	configDefaults := `
resource "github_team" "test" {
  name = "%s"
}
`

	configVisible := `
resource "github_team" "test" {
  name    = "%s"
	privacy = "closed"
}
`

	configWithParent := `
resource "github_team" "test" {
  name           = "%s"
	privacy        = "closed"
  parent_team_id = "%v"
}
`

	configFull := `
resource "github_team" "test" {
  name                 = "%s"
	description          = "%s"
	privacy              = "%s"
	notification_setting = "%s"
}
`

	t.Run("full_lifecycle", func(t *testing.T) {
		t.Parallel()

		teamName := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))
		description := "Terraform acceptance tests."
		privacy := "closed"
		notificationSetting := "notifications_disabled"

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(configDefaults, teamName),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("description"), knownvalue.Null()),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("privacy"), knownvalue.StringExact("secret")),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("notification_setting"), knownvalue.StringExact("notifications_enabled")),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("parent_team_id"), knownvalue.StringExact("")),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("ldap_dn"), knownvalue.Null()),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("create_default_maintainer"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("slug"), knownvalue.StringExact(teamName)),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("members_count"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("parent_team_read_id"), knownvalue.StringExact("")),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("parent_team_read_slug"), knownvalue.StringExact("")),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("node_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("etag"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(configFull, teamName, description, privacy, notificationSetting),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("description"), knownvalue.StringExact(description)),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("privacy"), knownvalue.StringExact(privacy)),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("notification_setting"), knownvalue.StringExact(notificationSetting)),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("parent_team_id"), knownvalue.StringExact("")),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("ldap_dn"), knownvalue.Null()),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("create_default_maintainer"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("slug"), knownvalue.StringExact(teamName)),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("members_count"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("parent_team_read_id"), knownvalue.StringExact("")),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("parent_team_read_slug"), knownvalue.StringExact("")),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("node_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("etag"), knownvalue.NotNull()),
					},
				},
				{
					ResourceName:            "github_team.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"etag"},
				},
			},
		})
	})

	t.Run("change_name", func(t *testing.T) {
		t.Parallel()

		teamName := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))
		teamNameUpdated := fmt.Sprintf("%s-updated", teamName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(configDefaults, teamName),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("slug"), knownvalue.StringExact(teamName)),
					},
				},
				{
					Config: fmt.Sprintf(configDefaults, teamNameUpdated),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("slug"), knownvalue.StringExact(teamNameUpdated)),
					},
				},
			},
		})
	})

	t.Run("change_name_with_non_slug_characters", func(t *testing.T) {
		t.Parallel()

		teamName := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))
		teamNameUpdated := fmt.Sprintf("%s updated", teamName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(configDefaults, teamName),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("slug"), knownvalue.StringExact(teamName)),
					},
				},
				{
					Config: fmt.Sprintf(configDefaults, teamNameUpdated),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team.test", tfjsonpath.New("slug"), knownvalue.StringExact(strings.ReplaceAll(teamNameUpdated, " ", "-"))),
					},
				},
			},
		})
	})

	t.Run("create_with_parent_team", func(t *testing.T) {
		t.Parallel()

		parentTeam := mustCreateTestTeam(t, nil)
		teamName := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(configWithParent, teamName, parentTeam.GetID()),
				},
				{
					Config: fmt.Sprintf(configVisible, teamName),
				},
			},
		})
	})

	t.Run("update_with_parent_team", func(t *testing.T) {
		t.Parallel()

		parentTeam := mustCreateTestTeam(t, nil)
		teamName := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(configVisible, teamName),
				},
				{
					Config: fmt.Sprintf(configWithParent, teamName, parentTeam.GetID()),
				},
			},
		})
	})
}
