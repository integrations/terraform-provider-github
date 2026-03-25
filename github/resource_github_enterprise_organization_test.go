package github

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestIsSAMLEnforcementError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name: "GitHub ErrorResponse with SAML enforcement",
			err: &github.ErrorResponse{
				Response: &http.Response{StatusCode: 403},
				Message:  "Resource protected by organization SAML enforcement. You must grant your Personal Access token access to this organization.",
			},
			expected: true,
		},
		{
			name: "GitHub ErrorResponse 403 without SAML message",
			err: &github.ErrorResponse{
				Response: &http.Response{StatusCode: 403},
				Message:  "Forbidden",
			},
			expected: false,
		},
		{
			name: "GitHub ErrorResponse 404",
			err: &github.ErrorResponse{
				Response: &http.Response{StatusCode: 404},
				Message:  "Not Found",
			},
			expected: false,
		},
		{
			name:     "plain error with SAML enforcement message",
			err:      errors.New("Resource protected by organization SAML enforcement"),
			expected: true,
		},
		{
			name:     "plain error without SAML message",
			err:      errors.New("some other error"),
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isSAMLEnforcementError(tc.err)
			if result != tc.expected {
				t.Errorf("isSAMLEnforcementError(%v) = %v, want %v", tc.err, result, tc.expected)
			}
		})
	}
}

func TestAccGithubEnterpriseOrganization(t *testing.T) {
	t.Run("creates and updates an enterprise organization without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		orgName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		desc := "Initial org description"
		updatedDesc := "Updated org description"

		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
			slug = "%s"
			}

			data "github_user" "current" {
			username = ""
			}

			resource "github_enterprise_organization" "org" {
			enterprise_id = data.github_enterprise.enterprise.id
			name          = "%s"
			description   = "%s"
			billing_email = data.github_user.current.email
			admin_logins  = [
				data.github_user.current.login
			]
			}
			`, testAccConf.enterpriseSlug, orgName, desc)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(
					"github_enterprise_organization.org", "enterprise_id",
				),
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "name",
					orgName,
				),
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "description",
					desc,
				),
				resource.TestCheckResourceAttrSet(
					"github_enterprise_organization.org", "billing_email",
				),
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "admin_logins.#",
					"1",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "description",
					updatedDesc,
				),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  checks["before"],
				},
				{
					Config: strings.Replace(config,
						desc,
						updatedDesc, 1),
					Check: checks["after"],
				},
			},
		})
	})

	t.Run("deletes an enterprise organization without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		orgName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_user" "current" {
				username = ""
			}

			resource "github_enterprise_organization" "org" {
				enterprise_id = data.github_enterprise.enterprise.id
				name          = "%s"
				billing_email = data.github_user.current.email
				admin_logins  = [
					data.github_user.current.login
				]
			}
			`, testAccConf.enterpriseSlug, orgName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:  config,
					Destroy: true,
				},
			},
		})
	})

	t.Run("creates and updates org with display name", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		orgName := fmt.Sprintf("tf-acc-test-displayname%s", randomID)

		displayName := fmt.Sprintf("Tf Acc Test displayname %s", randomID)
		updatedDisplayName := fmt.Sprintf("Updated Tf Acc Test Display Name %s", randomID)

		desc := "Initial org description"
		updatedDesc := "Updated org description"

		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
			slug = "%s"
			}

			data "github_user" "current" {
			username = ""
			}

			resource "github_enterprise_organization" "org" {
			enterprise_id = data.github_enterprise.enterprise.id
			name          = "%s"
			display_name  = "%s"
			description   = "%s"
			billing_email = data.github_user.current.email
			admin_logins  = [
				data.github_user.current.login
			]
			}
			`, testAccConf.enterpriseSlug, orgName, displayName, desc)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(
					"github_enterprise_organization.org", "enterprise_id",
				),
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "name",
					orgName,
				),
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "display_name",
					displayName,
				),
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "description",
					desc,
				),
				resource.TestCheckResourceAttrSet(
					"github_enterprise_organization.org", "billing_email",
				),
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "admin_logins.#",
					"1",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "description",
					updatedDesc,
				),
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "display_name",
					updatedDisplayName,
				),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  checks["before"],
				},
				{
					Config: strings.Replace(
						strings.Replace(config,
							displayName,
							updatedDisplayName, 1),
						desc,
						updatedDesc, 1),
					Check: checks["after"],
				},
			},
		})
	})

	t.Run("creates org without display name, set and update display name", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		orgName := fmt.Sprintf("tf-acc-test-adddisplayname%s", randomID)

		displayName := fmt.Sprintf("Tf Acc Test Add displayname %s", randomID)
		updatedDisplayName := fmt.Sprintf("Updated Tf Acc Test Add Display Name %s", randomID)

		desc := "Initial org description"
		updatedDesc := "Updated org description"

		configWithoutDisplayName := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
			slug = "%s"
			}

			data "github_user" "current" {
			username = ""
			}

			resource "github_enterprise_organization" "org" {
			enterprise_id = data.github_enterprise.enterprise.id
			name          = "%s"
			description   = "%s"
			billing_email = data.github_user.current.email
			admin_logins  = [
				data.github_user.current.login
			]
			}
			`, testAccConf.enterpriseSlug, orgName, desc)

		configWithDisplayName := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_user" "current" {
				username = ""
			}

			resource "github_enterprise_organization" "org" {
				enterprise_id = data.github_enterprise.enterprise.id
				name          = "%s"
				display_name  = "%s"
				description   = "%s"
				billing_email = data.github_user.current.email
				admin_logins  = [
				data.github_user.current.login
				]
			}
				`, testAccConf.enterpriseSlug, orgName, displayName, desc)

		checks := map[string]resource.TestCheckFunc{
			"create": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(
					"github_enterprise_organization.org", "enterprise_id",
				),
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "name",
					orgName,
				),
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "description",
					desc,
				),
				resource.TestCheckResourceAttrSet(
					"github_enterprise_organization.org", "billing_email",
				),
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "admin_logins.#",
					"1",
				),
			),
			"set": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "description",
					desc,
				),
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "display_name",
					displayName,
				),
			),
			"updateDisplayName": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "description",
					desc,
				),
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "display_name",
					updatedDisplayName,
				),
			),
			"updateDesc": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "description",
					updatedDesc,
				),
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "display_name",
					updatedDisplayName,
				),
			),
			"unset": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_enterprise_organization.org", "description",
					desc,
				),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configWithoutDisplayName,
					Check:  checks["create"],
				},
				{
					Config: configWithDisplayName,
					Check:  checks["set"],
				},
				{
					Config: strings.Replace(configWithDisplayName,
						displayName,
						updatedDisplayName, 1),
					Check: checks["updateDisplayName"],
				},
				{
					Config: strings.Replace(
						strings.Replace(configWithDisplayName,
							displayName,
							updatedDisplayName, 1),
						desc,
						updatedDesc, 1),
					Check: checks["updateDesc"],
				},
				{
					Config: configWithoutDisplayName,
					Check:  checks["unset"],
				},
			},
		})
	})

	t.Run("imports enterprise organization without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		orgName := fmt.Sprintf("tf-acc-test-import%s", randomID)

		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_user" "current" {
				username = ""
			}

			resource "github_enterprise_organization" "org" {
				enterprise_id = data.github_enterprise.enterprise.id
				name          = "%s"
				billing_email = data.github_user.current.email
				admin_logins  = [
				data.github_user.current.login
				]
			}
				`, testAccConf.enterpriseSlug, orgName)

		check := resource.ComposeTestCheckFunc()

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:      "github_enterprise_organization.org",
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateId:     fmt.Sprintf(`%s/%s`, testAccConf.enterpriseSlug, orgName),
				},
			},
		})
	})

	t.Run("imports enterprise organization invalid enterprise name", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		orgName := fmt.Sprintf("tf-acc-test-adddisplayname%s", randomID)

		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_user" "current" {
				username = ""
			}

			resource "github_enterprise_organization" "org" {
				enterprise_id = data.github_enterprise.enterprise.id
				name          = "%s"
				description   = "org description"
				billing_email = data.github_user.current.email
				admin_logins  = [
				data.github_user.current.login
				]
			}
				`, testAccConf.enterpriseSlug, orgName)

		check := resource.ComposeTestCheckFunc()

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:  "github_enterprise_organization.org",
					ImportState:   true,
					ImportStateId: fmt.Sprintf(`%s/%s`, randomID, orgName),
					ExpectError:   regexp.MustCompile("Could not resolve to a Business with the URL slug of .*"),
				},
			},
		})
	})

	t.Run("imports enterprise organization invalid organization name", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		orgName := fmt.Sprintf("tf-acc-test-adddisplayname%s", randomID)

		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_user" "current" {
				username = ""
			}

			resource "github_enterprise_organization" "org" {
				enterprise_id = data.github_enterprise.enterprise.id
				name          = "%s"
				description   = "org description"
				billing_email = data.github_user.current.email
				admin_logins  = [
				data.github_user.current.login
				]
			}
				`, testAccConf.enterpriseSlug, orgName)

		check := resource.ComposeTestCheckFunc()

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:  "github_enterprise_organization.org",
					ImportState:   true,
					ImportStateId: fmt.Sprintf(`%s/%s`, testAccConf.enterpriseSlug, randomID),
					ExpectError:   regexp.MustCompile("Could not resolve to an Organization with the login of .*"),
				},
			},
		})
	})
}
