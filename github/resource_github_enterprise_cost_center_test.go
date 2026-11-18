package github

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestGithubEnterpriseCostCenterNameValidation(t *testing.T) {
	t.Parallel()

	nameSchema := resourceGithubEnterpriseCostCenter().Schema["name"]
	if diags := nameSchema.ValidateDiagFunc(strings.Repeat("a", 255), cty.GetAttrPath("name")); diags.HasError() {
		t.Fatalf("expected a 255-character name to pass validation: %s", diags[0].Summary)
	}
	if diags := nameSchema.ValidateDiagFunc(strings.Repeat("a", 256), cty.GetAttrPath("name")); !diags.HasError() {
		t.Fatal("expected a 256-character name to fail validation")
	}
}

func TestGithubEnterpriseCostCenterAssignmentValidation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		resource *schema.Resource
		field    string
	}{
		{name: "users", resource: resourceGithubEnterpriseCostCenterUsers(), field: "usernames"},
		{name: "organizations", resource: resourceGithubEnterpriseCostCenterOrganizations(), field: "organization_logins"},
		{name: "repositories", resource: resourceGithubEnterpriseCostCenterRepositories(), field: "repository_names"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			elem, ok := test.resource.Schema[test.field].Elem.(*schema.Schema)
			if !ok {
				t.Fatalf("expected %s elements to use a string schema", test.field)
			}
			if diags := elem.ValidateDiagFunc("", cty.GetAttrPath(test.field)); !diags.HasError() {
				t.Fatalf("expected an empty %s element to fail validation", test.field)
			}
			if diags := elem.ValidateDiagFunc("valid", cty.GetAttrPath(test.field)); diags.HasError() {
				t.Fatalf("expected a non-empty %s element to pass validation: %s", test.field, diags[0].Summary)
			}
		})
	}
}

func TestGithubEnterpriseCostCenterDataSourceReadTimeout(t *testing.T) {
	t.Parallel()

	timeout := dataSourceGithubEnterpriseCostCenter().Timeouts.Read
	if timeout == nil || *timeout != 5*time.Minute {
		t.Fatalf("expected a five-minute read timeout, got %v", timeout)
	}
}

func TestAccGithubEnterpriseCostCenter(t *testing.T) {
	t.Run("creates cost center without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEnterpriseCostCenterDestroy,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_cost_center" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%s%s"
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_cost_center.test", tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
						statecheck.ExpectKnownValue("github_enterprise_cost_center.test", tfjsonpath.New("name"), knownvalue.StringExact(testResourcePrefix+randomID)),
						statecheck.ExpectKnownValue("github_enterprise_cost_center.test", tfjsonpath.New("state"), knownvalue.StringExact("active")),
					},
				},
			},
		})
	})

	t.Run("updates cost center name without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEnterpriseCostCenterDestroy,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_cost_center" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%s%s"
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_cost_center.test", tfjsonpath.New("name"), knownvalue.StringExact(testResourcePrefix+randomID)),
					},
				},
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_cost_center" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%supdated-%s"
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_cost_center.test", tfjsonpath.New("name"), knownvalue.StringExact(testResourcePrefix+"updated-"+randomID)),
					},
				},
			},
		})
	})

	t.Run("imports cost center without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEnterpriseCostCenterDestroy,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_cost_center" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%s%s"
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
				},
				{
					ResourceName:        "github_enterprise_cost_center.test",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: testAccConf.enterpriseSlug + ":",
				},
			},
		})
	})
}

func testAccCheckGithubEnterpriseCostCenterDestroy(s *terraform.State) error {
	meta, err := getTestMeta(testAccConf)
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_enterprise_cost_center" {
			continue
		}

		enterpriseSlug := rs.Primary.Attributes["enterprise_slug"]
		costCenterID := rs.Primary.ID
		_, err := retryUntilOK(context.Background(), func() (bool, bool, error) {
			cc, err := getEnterpriseCostCenter(context.Background(), meta.v3client, enterpriseSlug, costCenterID, meta.maxPerPage)
			if errIs404(err) {
				return true, true, nil
			}
			if err != nil {
				return false, false, err
			}
			return true, cc.GetState() == "deleted", nil
		}, &retryOptions{
			delay:   5 * time.Second,
			timeout: 5 * time.Minute,
		})
		if err != nil {
			return fmt.Errorf("waiting for cost center %s to be archived: %w", costCenterID, err)
		}
	}

	return nil
}
