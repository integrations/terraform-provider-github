package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var (
	testAccProviders         map[string]*schema.Provider
	testAccProviderFactories func(providers *[]*schema.Provider) map[string]func() (*schema.Provider, error)
	testAccProvider          *schema.Provider
)

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"github": testAccProvider,
	}
	testAccProviderFactories = func(providers *[]*schema.Provider) map[string]func() (*schema.Provider, error) {
		return map[string]func() (*schema.Provider, error){
			//nolint:unparam
			"github": func() (*schema.Provider, error) {
				p := Provider()
				*providers = append(*providers, p)
				return p, nil
			},
		}
	}
}

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
	t.Run("can be configured to run anonymously", func(t *testing.T) {
		config := `
		provider "github" {
			token = ""
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

	t.Run("can be configured to run insecurely", func(t *testing.T) {
		config := `
		provider "github" {
			token = ""
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
			}`, testAccConf.token, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, individual) },
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
			PreCheck:          func() { skipUnauthenticated(t) },
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
			PreCheck:          func() { skipUnauthenticated(t) },
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
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("only one of `app_auth,token` can be specified"),
				},
			},
		})
	})
	t.Run("should not allow app_auth and GITHUB_TOKEN to be configured", func(t *testing.T) {
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
			PreCheck:          func() { skipUnauthenticated(t); t.Setenv("GITHUB_TOKEN", "1234567890") },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("only one of `app_auth,token` can be specified"),
				},
			},
		})
	})
}
