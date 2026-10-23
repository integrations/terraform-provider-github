package github

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryCustomProperty(t *testing.T) {
	t.Parallel()

	t.Run("single_select", func(t *testing.T) {
		t.Parallel()

		prop := mustCreateTestOrganizationRepositoryCustomProperty(t, "single_select", []string{"option1", "option2"})
		repo := mustCreateTestRepository(t)
		allowed := prop.GetAllowedValues()

		config := fmt.Sprintf(`
resource "github_repository_custom_property" "test" {
  repository     = "%s"
  property_name  = "%s"
  property_type  = "%s"
  property_value = %%s
}
`, repo.GetName(), prop.GetPropertyName(), prop.GetValueType())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, `[]`),
					ExpectError: regexp.MustCompile(`Not enough list items`),
					PlanOnly:    true,
				},
				{
					Config: fmt.Sprintf(config, fmt.Sprintf(`["%s"]`, allowed[0])),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_custom_property.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				{
					Config:      fmt.Sprintf(config, `["invalid_option"]`),
					ExpectError: regexp.MustCompile(`is not allowed for property`),
				},
				{
					Config: fmt.Sprintf(config, fmt.Sprintf(`["%s"]`, allowed[1])),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_custom_property.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_repository_custom_property.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("multi_select", func(t *testing.T) {
		t.Parallel()

		prop := mustCreateTestOrganizationRepositoryCustomProperty(t, "multi_select", []string{"option1", "option2", "option3"})
		repo := mustCreateTestRepository(t)
		allowed := prop.GetAllowedValues()

		config := fmt.Sprintf(`
resource "github_repository_custom_property" "test" {
  repository     = "%s"
  property_name  = "%s"
  property_type  = "%s"
  property_value = %%s
}
`, repo.GetName(), prop.GetPropertyName(), prop.GetValueType())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, `[]`),
					ExpectError: regexp.MustCompile(`Not enough list items`),
					PlanOnly:    true,
				},
				{
					Config: fmt.Sprintf(config, fmt.Sprintf(`["%s", "%s"]`, allowed[0], allowed[1])),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_custom_property.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				{
					Config:      fmt.Sprintf(config, `["invalid_option"]`),
					ExpectError: regexp.MustCompile(`is not allowed for property`),
				},
				{
					Config: fmt.Sprintf(config, fmt.Sprintf(`["%s"]`, allowed[2])),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_custom_property.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_repository_custom_property.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("true_false", func(t *testing.T) {
		t.Parallel()

		prop := mustCreateTestOrganizationRepositoryCustomProperty(t, "true_false", nil)
		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_custom_property" "test" {
  repository    = "%s"
  property_name = "%s"
  property_type = "%s"
  property_value = %%s
}
`, repo.GetName(), prop.GetPropertyName(), prop.GetValueType())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, `[]`),
					ExpectError: regexp.MustCompile(`Not enough list items`),
					PlanOnly:    true,
				},
				{
					Config: fmt.Sprintf(config, `["true"]`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_custom_property.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				{
					Config:      fmt.Sprintf(config, `["invalid_option"]`),
					ExpectError: regexp.MustCompile(`is not allowed for property`),
				},
				{
					Config: fmt.Sprintf(config, `["false"]`),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_custom_property.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_repository_custom_property.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("url", func(t *testing.T) {
		t.Parallel()

		prop := mustCreateTestOrganizationRepositoryCustomProperty(t, "url", nil)
		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_custom_property" "test" {
  repository     = "%s"
  property_name  = "%s"
  property_type  = "%s"
  property_value = %%s
}
`, repo.GetName(), prop.GetPropertyName(), prop.GetValueType())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, `[]`),
					ExpectError: regexp.MustCompile(`Not enough list items`),
					PlanOnly:    true,
				},
				{
					Config: fmt.Sprintf(config, `["https://example.com"]`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_custom_property.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				{
					Config:      fmt.Sprintf(config, `["xxxx"]`),
					ExpectError: regexp.MustCompile(`URL must be absolute`),
				},
				{
					Config: fmt.Sprintf(config, `["https://example.com/test"]`),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_custom_property.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_repository_custom_property.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("enterprise_scoped_schema", func(t *testing.T) {
		t.Parallel()

		skipUnlessEnterprise(t)

		prop := mustCreateTestEnterpriseRepositoryCustomProperty(t, "true_false", nil)
		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_custom_property" "test" {
  repository    = "%s"
  property_name = "%s"
  property_type = "%s"
  property_value = %%s
}
`, repo.GetName(), prop.GetPropertyName(), prop.GetValueType())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, `["true"]`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_custom_property.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				{
					ResourceName:      "github_repository_custom_property.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()

		prop := mustCreateTestOrganizationRepositoryCustomProperty(t, "string", nil)
		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_custom_property" "test" {
  repository     = "%s"
  property_name  = "%s"
  property_type  = "%s"
  property_value = %%s
}
`, repo.GetName(), prop.GetPropertyName(), prop.GetValueType())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, `[]`),
					ExpectError: regexp.MustCompile(`Not enough list items`),
					PlanOnly:    true,
				},
				{
					Config:      fmt.Sprintf(config, `[""]`),
					ExpectError: regexp.MustCompile(`to not be an empty string`),
					PlanOnly:    true,
				},
				{
					Config: fmt.Sprintf(config, `["text"]`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_custom_property.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, `["new text"]`),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_custom_property.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_repository_custom_property.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}

func Test_resourceGithubRepositoryCustomPropertyImport(t *testing.T) {
	t.Run("succeeds_for_property_sourced_from_enterprise", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			mustGetTestMockResponse(t, "/repos/my-org/my-repo", 200, &github.Repository{
				ID: new(int64(123456)),
			}),
			mustGetTestMockResponse(t, "/orgs/my-org/properties/schema", 200, []*github.CustomProperty{
				{
					PropertyName: new("hcp_prod"),
					ValueType:    github.PropertyValueTypeTrueFalse,
					SourceType:   new("enterprise"),
				},
			}),
		})
		defer ts.Close()

		meta := &Owner{
			name:     "my-org",
			v3client: mustCreateTestGitHubClient(t, ts.URL),
		}

		d := resourceGithubRepositoryCustomProperty().Data(&terraform.InstanceState{
			ID: "my-org:my-repo:hcp_prod",
		})

		got, err := resourceGithubRepositoryCustomPropertyImport(t.Context(), d, meta)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if len(got) != 1 {
			t.Fatalf("expected 1 resource data, got %d", len(got))
		}

		if v := got[0].Get("repository"); v != "my-repo" {
			t.Errorf("repository = %v, want my-repo", v)
		}
		if v := got[0].Get("repository_id"); v != 123456 {
			t.Errorf("repository_id = %v, want 123456", v)
		}
		if v := got[0].Get("property_name"); v != "hcp_prod" {
			t.Errorf("property_name = %v, want hcp_prod", v)
		}
		if v := got[0].Get("property_type"); v != string(github.PropertyValueTypeTrueFalse) {
			t.Errorf("property_type = %v, want %v", v, github.PropertyValueTypeTrueFalse)
		}
	})

	t.Run("fails_when_property_not_found", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			mustGetTestMockResponse(t, "/repos/my-org/my-repo", 200, &github.Repository{
				ID: new(int64(123456)),
			}),
			mustGetTestMockResponse(t, "/orgs/my-org/properties/schema", 200, []*github.CustomProperty{}),
		})
		defer ts.Close()

		meta := &Owner{
			name:     "my-org",
			v3client: mustCreateTestGitHubClient(t, ts.URL),
		}

		d := resourceGithubRepositoryCustomProperty().Data(&terraform.InstanceState{
			ID: "my-org:my-repo:hcp_prod",
		})

		_, err := resourceGithubRepositoryCustomPropertyImport(t.Context(), d, meta)
		if err == nil {
			t.Fatal("expected error, got none")
		}
		if !strings.Contains(err.Error(), "not found") {
			t.Fatalf("unexpected error: %s", err)
		}
	})

	t.Run("fails_when_owner_does_not_match_authenticated_owner", func(t *testing.T) {
		meta := &Owner{name: "my-org"}

		d := resourceGithubRepositoryCustomProperty().Data(&terraform.InstanceState{
			ID: "other-org:my-repo:hcp_prod",
		})

		_, err := resourceGithubRepositoryCustomPropertyImport(t.Context(), d, meta)
		if err == nil {
			t.Fatal("expected error, got none")
		}
		if !strings.Contains(err.Error(), "owner in id must match authenticated owner") {
			t.Fatalf("unexpected error: %s", err)
		}
	})
}
