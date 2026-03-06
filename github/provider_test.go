package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestProvider(t *testing.T) {
	t.Run("validate", func(t *testing.T) {
		if err := Provider().InternalValidate(); err != nil {
			t.Fatalf("err: %s", err)
		}
	})
}

func TestAccProviderConfigure(t *testing.T) {
	t.Run("anonymous", func(t *testing.T) {
		config := `
provider "github" {}

data "github_ip_ranges" "test" {}
`

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				t.Setenv("GITHUB_TOKEN", "")
				t.Setenv("GITHUB_APP_PEM_FILE", "")
				t.Setenv("GH_PATH", "none-existent-path")
			},
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

	t.Run("insecure", func(t *testing.T) {
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

	t.Run("max_retries", func(t *testing.T) {
		testMaxRetries := -1
		config := fmt.Sprintf(`
provider "github" {
	max_retries = %d
}

data "github_ip_ranges" "test" {}
`, testMaxRetries)

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

	t.Run("max_per_page", func(t *testing.T) {
		testMaxPerPage := 101
		config := fmt.Sprintf(`
provider "github" {
	max_per_page = %d
}

data "github_ip_ranges" "test" {}
`, testMaxPerPage)

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

	t.Run("app_auth", func(t *testing.T) {
		config := fmt.Sprintf(`
provider "github" {
	owner = "%s"
	app_auth {
		id = "%s"
		installation_id = "%s"
		pem_file = "%s"
	}
}

data "github_ip_ranges" "test" {}
`, testAccConf.owner, testAccConf.appID, testAccConf.appInstallationID, testAccConf.appPEM)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipNoApp(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:             config,
					ExpectNonEmptyPlan: false,
				},
			},
		})
	})

	t.Run("app_auth_ignore_token_env", func(t *testing.T) {
		config := fmt.Sprintf(`
provider "github" {
	owner = "%s"
	app_auth {
		id = "%s"
		installation_id = "%s"
		pem_file = "%s"
	}
}

data "github_ip_ranges" "test" {}
`, testAccConf.owner, testAccConf.appID, testAccConf.appInstallationID, testAccConf.appPEM)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipNoApp(t); t.Setenv("GITHUB_TOKEN", "1234567890") },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:             config,
					ExpectNonEmptyPlan: false,
				},
			},
		})
	})

	t.Run("token_auth", func(t *testing.T) {
		config := fmt.Sprintf(`
provider "github" {
	owner = "%s"
	token = "%s"
}
`, testAccConf.token, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipNoToken(t) },
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

	t.Run("organization_account_with_token_legacy", func(t *testing.T) {
		config := fmt.Sprintf(`
provider "github" {
	token = "%s"
	organization = "%s"
}`, testAccConf.token, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipNoToken(t)
				skipUnlessHasOrgs(t)
			},
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

	t.Run("ghes_with_token", func(t *testing.T) {
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
			PreCheck: func() {
				skipNoToken(t)
				skipNoApp(t)
			},
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
