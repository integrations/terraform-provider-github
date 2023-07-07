package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testAccProviders map[string]*schema.Provider
var testAccProviderFactories func(providers *[]*schema.Provider) map[string]
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"github": testAccProvider,
	}
	testAccProviderFactories = func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory {
		return map[string]terraform.ResourceProviderFactory{
			"github": func() (schema.Provider, error) {
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

		var _ schema.Provider = *Provider()
	})

}

func TestAccProviderConfigure(t *testing.T) {

	t.Run("can be configured to run anonymously", func(t *testing.T) {

		config := `
			provider "github" {}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, anonymous) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config:             config,
					ExpectNonEmptyPlan: false,
				},
			},
		})

	})

	t.Run("can be configured to run insecurely", func(t *testing.T) {

		config := fmt.Sprintf(`
				provider "github" {
					token = "%s"
					insecure = true
				}`,
			testToken,
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, anonymous) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config:             config,
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
			}`,
			testToken, testOwnerFunc(),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, individual) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config:             config,
					ExpectNonEmptyPlan: false,
				},
			},
		})

	})

	t.Run("can be configured with an organization account", func(t *testing.T) {

		config := fmt.Sprintf(`
			provider "github" {
				token = "%s"
				organization = "%s"
			}`,
			testToken, testOrganizationFunc(),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, organization) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config:             config,
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
			}`,
			testToken, testBaseURLGHES,
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, individual) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config:             config,
					ExpectNonEmptyPlan: false,
				},
			},
		})

	})
}
