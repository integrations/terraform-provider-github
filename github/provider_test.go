package github

import (
	"fmt"
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
				owner = "%s"
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
		config := fmt.Sprintf(`
			provider "github" {
				owner = "%s"
				max_retries = 3
			}`, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:             config,
					ExpectNonEmptyPlan: false,
				},
			},
		})
	})

	t.Run("can be configured with max per page", func(t *testing.T) {
		config := `
			provider "github" {
				owner = "%s"
				max_per_page = 100
			}`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:             config,
					ExpectNonEmptyPlan: false,
					Check: func(_ *terraform.State) error {
						if maxPerPage != 100 {
							return fmt.Errorf("max_per_page should be set to 100, got %d", maxPerPage)
						}
						return nil
					},
				},
			},
		})
	})
}
