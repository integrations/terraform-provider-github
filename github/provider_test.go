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

// TODO: this is failing.
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

	t.Run("can be configured with max retries", func(t *testing.T) {
		config := fmt.Sprintf(`
			provider "github" {
				token = "%s"
				owner = "%s"
				max_retries = 3
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

	t.Run("can be configured with max per page", func(t *testing.T) {
		config := fmt.Sprintf(`
			provider "github" {
				token = "%s"
				owner = "%s"
				max_per_page = 999
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
					Check: func(_ *terraform.State) error {
						if maxPerPage != 999 {
							return fmt.Errorf("max_per_page should be set to 999, got %d", maxPerPage)
						}
						return nil
					},
				},
			},
		})
	})
}

func Test_validateBaseURL(t *testing.T) {
	testCases := []struct {
		name  string
		url   string
		valid bool
	}{
		{
			name:  "dotcom",
			url:   "https://api.github.com/",
			valid: true,
		},
		{
			name:  "dotcom_no_slash",
			url:   "https://api.github.com",
			valid: true,
		},
		{
			name:  "dotcom_with_path",
			url:   "http://api.github.com/test/",
			valid: false,
		},
		{
			name:  "dotcom_http",
			url:   "http://api.github.com/",
			valid: false,
		},
		{
			name:  "dotcom_no_scheme",
			url:   "api.github.com/",
			valid: false,
		},
		{
			name:  "ghec",
			url:   "https://customer.ghe.com/",
			valid: true,
		},
		{
			name:  "ghec_no_slash",
			url:   "https://customer.ghe.com",
			valid: true,
		},
		{
			name:  "ghec_with_path",
			url:   "https://customer.ghe.com/test/",
			valid: false,
		},
		{
			name:  "ghec_http",
			url:   "http://customer.ghe.com/",
			valid: false,
		},
		{
			name:  "ghec_no_scheme",
			url:   "customer.ghe.com/",
			valid: false,
		},
		{
			name:  "ghes",
			url:   "https://example.com/",
			valid: true,
		},
		{
			name:  "ghes_no_slash",
			url:   "https://example.com",
			valid: true,
		},
		{
			name:  "ghes_with_path",
			url:   "https://example.com/test/",
			valid: true,
		},
		{
			name:  "ghes_http",
			url:   "http://example.com/",
			valid: true,
		},
		{
			name:  "ghes_no_scheme/",
			url:   "example.com",
			valid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := validateBaseURL(tc.url)
			if err != nil && tc.valid {
				t.Errorf("URL %q: expected valid URL, got error: %s", tc.url, err)
			} else if err == nil && !tc.valid {
				t.Errorf("URL %q: expected invalid URL, got no error", tc.url)
			}
		})
	}
}
