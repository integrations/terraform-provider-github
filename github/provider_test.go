package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestProvider(t *testing.T) {
	t.Run("runs internal validation without error", func(t *testing.T) {
		if err := Provider().InternalValidate(); err != nil {
			t.Fatalf("err: %s", err)
		}
	})

	t.Run("has an implementation", func(t *testing.T) {
		// FIXME: unsure if this is useful; refactored from:
		// func TestProvider_impl(t *testing.T) {
		// 	var _ terraform.ResourceProvider = Provider()
		// }

		_ = *Provider()
	})
}

func TestAccProviderConfigure(t *testing.T) {
	t.Run("can_be_configured_to_run_anonymously", func(t *testing.T) {
		config := `
		provider "github" {
		}
		data "github_ip_ranges" "test" {}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { t.Setenv("GITHUB_TOKEN", ""); t.Setenv("GH_PATH", "none-existent-path") },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:             config,
					PlanOnly:           true,
					ExpectNonEmptyPlan: false,
				},
			},
		})
	})

	t.Run("can_be_configured_with_app_auth_and_ignore_github_token", func(t *testing.T) {
		t.Skip("This test requires a valid app auth setup to run.")
		config := fmt.Sprintf(`
provider "github" {
	owner = "%s"
	app_auth {
		id = "1234567890"
		installation_id = "1234567890"
		pem_file = "1234567890"
	}
}

data "github_ip_ranges" "test" {}
`, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { t.Setenv("GITHUB_TOKEN", "1234567890") },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:             config,
					ExpectNonEmptyPlan: false,
				},
			},
		})
	})

	t.Run("can be configured to run insecurely", func(t *testing.T) {
		config := `
		provider "github" {
			insecure = true
		}
		data "github_ip_ranges" "test" {}
		`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:             config,
					PlanOnly:           true,
					ExpectNonEmptyPlan: false,
				},
			},
		})
	})

	t.Run("can be configured with an individual account", func(t *testing.T) {
		config := fmt.Sprintf(`
			provider "github" {
				token = "%s"
				owner = "%s"
			}
			`, testAccConf.token, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, individual) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:             config,
					PlanOnly:           true,
					ExpectNonEmptyPlan: false,
				},
			},
		})
	})

	t.Run("can be configured with an organization account", func(t *testing.T) {
		config := fmt.Sprintf(`
			provider "github" {
				token = "%s"
				owner = "%[2]s"
			}

			data "github_organization" "test" {
				name = "%[2]s"
			}
			`, testAccConf.token, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
			},
		})
	})

	t.Run("can be configured with an organization account legacy", func(t *testing.T) {
		config := fmt.Sprintf(`
			provider "github" {
				token = "%s"
				organization = "%s"
			}`, testAccConf.token, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:             config,
					PlanOnly:           true,
					ExpectNonEmptyPlan: false,
				},
			},
		})
	})

	t.Run("can be configured with a GHES deployment", func(t *testing.T) {
		config := fmt.Sprintf(`
			provider "github" {
				token = "%s"
				base_url = "%s"
			}`, testAccConf.token, testAccConf.baseURL)

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessMode(t, enterprise)
				if testAccConf.baseURL.Host != "api.github.com" {
					t.Skip("Skipping as test mode is not GHES")
				}
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:             config,
					ExpectNonEmptyPlan: false,
				},
			},
		})
	})

	t.Run("can be configured with max retries", func(t *testing.T) {
		testMaxRetries := -1
		config := fmt.Sprintf(`
			provider "github" {
				owner = "%s"
				max_retries = %d
			}

			data "github_ip_ranges" "test" {}
			`, testAccConf.owner, testMaxRetries)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:             config,
					ExpectNonEmptyPlan: false,
					ExpectError:        regexp.MustCompile("max_retries must be greater than or equal to 0"),
				},
			},
		})
	})

	t.Run("can be configured with max per page", func(t *testing.T) {
		testMaxPerPage := 101
		config := fmt.Sprintf(`
			provider "github" {
				owner = "%s"
				max_per_page = %d
			}

			data "github_ip_ranges" "test" {}
			`, testAccConf.owner, testMaxPerPage)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:             config,
					ExpectNonEmptyPlan: false,
					Check: func(_ *terraform.State) error {
						if maxPerPage != testMaxPerPage {
							return fmt.Errorf("max_per_page should be set to %d, got %d", testMaxPerPage, maxPerPage)
						}
						return nil
					},
				},
			},
		})
	})
	t.Run("should not allow both token and app_auth to be configured", func(t *testing.T) {
		t.Skip("This would be a semver breaking change, this will be reinstated for v7.")
		config := fmt.Sprintf(`
			provider "github" {
				owner = "%s"
				token = "%s"
				app_auth {
					id = "1234567890"
					installation_id = "1234567890"
					pem_file = "1234567890"
				}
			}

			data "github_ip_ranges" "test" {}
			`, testAccConf.owner, testAccConf.token)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`"app_auth": conflicts with token`),
				},
			},
		})
	})
}

func Test_ghCLIHostFromAPIHost(t *testing.T) {
	testCases := []struct {
		name         string
		host         string
		expectedHost string
	}{
		{
			name:         "dotcom API host is mapped to dotcom host",
			host:         "api.github.com",
			expectedHost: "github.com",
		},
		{
			name:         "ghec API host has api. prefix stripped",
			host:         "api.my-enterprise.ghe.com",
			expectedHost: "my-enterprise.ghe.com",
		},
		{
			name:         "ghec API host with numbers has api. prefix stripped",
			host:         "api.customer-123.ghe.com",
			expectedHost: "customer-123.ghe.com",
		},
		{
			name:         "ghes host is passed through unchanged",
			host:         "github.example.com",
			expectedHost: "github.example.com",
		},
		{
			name:         "ghes host with port is passed through unchanged",
			host:         "github.example.com:8443",
			expectedHost: "github.example.com:8443",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ghCLIHostFromAPIHost(tc.host)
			if got != tc.expectedHost {
				t.Errorf("ghCLIHostFromAPIHost(%q) = %q, want %q", tc.host, got, tc.expectedHost)
			}
		})
	}
}
