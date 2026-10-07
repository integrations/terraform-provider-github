package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

const (
	testCustomPropsRepoNameFmt    = "%srepo-custom-props-%s"
	testCustomPropsEnvPropNameFmt = "tf-acc-env-%s"
	testCustomPropsResourceAddr   = "github_repository_custom_properties.test"
)

func TestAccGithubRepositoryCustomProperties(t *testing.T) {
	t.Run("creates and reads multiple custom properties", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf(testCustomPropsRepoNameFmt, testResourcePrefix, randomID)
		envPropName := fmt.Sprintf(testCustomPropsEnvPropNameFmt, randomID)
		teamPropName := fmt.Sprintf("tf-acc-team-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "environment" {
				allowed_values = ["production", "staging", "development"]
				description    = "Deployment environment"
				property_name  = "%s"
				value_type     = "single_select"
			}

			resource "github_organization_custom_properties" "team" {
				description   = "Team responsible"
				property_name = "%s"
				value_type    = "string"
			}

			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_custom_properties" "test" {
				repository = github_repository.test.name

				property {
					name  = github_organization_custom_properties.environment.property_name
					value = ["production"]
				}

				property {
					name  = github_organization_custom_properties.team.property_name
					value = ["platform-team"]
				}
			}
		`, envPropName, teamPropName, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(testCustomPropsResourceAddr, tfjsonpath.New("repository"), knownvalue.StringExact(repoName)),
						statecheck.ExpectKnownValue(testCustomPropsResourceAddr, tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(testCustomPropsResourceAddr, tfjsonpath.New("property"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("updates property value in place", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf(testCustomPropsRepoNameFmt, testResourcePrefix, randomID)
		propName := fmt.Sprintf(testCustomPropsEnvPropNameFmt, randomID)

		configTmpl := `
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["production", "staging", "development"]
				description    = "Deployment environment"
				property_name  = "%s"
				value_type     = "single_select"
			}

			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_custom_properties" "test" {
				repository = github_repository.test.name

				property {
					name  = github_organization_custom_properties.test.property_name
					value = ["%s"]
				}
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(configTmpl, propName, repoName, "production"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(testCustomPropsResourceAddr, tfjsonpath.New("property"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(configTmpl, propName, repoName, "staging"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(testCustomPropsResourceAddr, plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(testCustomPropsResourceAddr, tfjsonpath.New("property"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("imports all properties for a repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf(testCustomPropsRepoNameFmt, testResourcePrefix, randomID)
		propName := fmt.Sprintf(testCustomPropsEnvPropNameFmt, randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["production", "staging"]
				description    = "Deployment environment"
				property_name  = "%s"
				value_type     = "single_select"
			}

			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_custom_properties" "test" {
				repository = github_repository.test.name

				property {
					name  = github_organization_custom_properties.test.property_name
					value = ["production"]
				}
			}
		`, propName, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:      testCustomPropsResourceAddr,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("creates multi_select property", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf(testCustomPropsRepoNameFmt, testResourcePrefix, randomID)
		propName := fmt.Sprintf("tf-acc-tags-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["go", "python", "rust", "typescript"]
				description    = "Language tags"
				property_name  = "%s"
				value_type     = "multi_select"
			}

			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_custom_properties" "test" {
				repository = github_repository.test.name

				property {
					name  = github_organization_custom_properties.test.property_name
					value = ["go", "rust"]
				}
			}
		`, propName, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(testCustomPropsResourceAddr, tfjsonpath.New("property"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"name": knownvalue.StringExact(propName),
								"value": knownvalue.SetExact([]knownvalue.Check{
									knownvalue.StringExact("go"),
									knownvalue.StringExact("rust"),
								}),
							}),
						})),
					},
				},
			},
		})
	})
}
