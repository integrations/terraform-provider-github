package github

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProviderFactories func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"github": testAccProvider,
	}
	testAccProviderFactories = func(providers *[]*schema.Provider) map[string]terraform.ResourceProviderFactory {
		return map[string]terraform.ResourceProviderFactory{
			"github": func() (terraform.ResourceProvider, error) {
				p := Provider()
				*providers = append(*providers, p.(*schema.Provider))
				return p, nil
			},
		}
	}
}

func TestProvider(t *testing.T) {

	t.Run("runs internal validation without error", func(t *testing.T) {

		if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
			t.Fatalf("err: %s", err)
		}

	})

	t.Run("has an implementation", func(t *testing.T) {
		// FIXME: unsure if this is useful; refactored from:
		// func TestProvider_impl(t *testing.T) {
		// 	var _ terraform.ResourceProvider = Provider()
		// }

		var _ terraform.ResourceProvider = Provider()
	})

}

func TestAccProviderConfigure(t *testing.T) {

	t.Run("can be configured to run insecurely", func(t *testing.T) {

		// Use ephemeral port range (49152–65535)
		port := fmt.Sprintf("%d", 49152+rand.Intn(16382))

		// Use self-signed certificate
		certFile := filepath.Join("test-fixtures", "cert.pem")
		keyFile := filepath.Join("test-fixtures", "key.pem")

		url, closeFunc := githubTLSApiMock(port, certFile, keyFile, t)
		defer func() {
			err := closeFunc()
			if err != nil {
				t.Fatal(err)
			}
		}()

		oldBaseUrl := os.Getenv("GITHUB_BASE_URL")
		defer os.Setenv("GITHUB_BASE_URL", oldBaseUrl)

		// Point provider to mock API with self-signed cert
		os.Setenv("GITHUB_BASE_URL", url)

		providerConfig := `provider "github" {}`

		username := "hashibot"
		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				testAccPreCheck(t)
			},
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config:      providerConfig + testAccCheckGithubUserDataSourceConfig(username),
					ExpectError: regexp.MustCompile("x509: certificate is valid for untrusted, not localhost"),
				},
			},
		})

	})

	t.Run("can be configured to run anonymously", func(t *testing.T) {

		anonymousConfiguration := `
			provider "github" {}
			data "github_ip_ranges" "test" {}
		`
		anonymousCheck := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "hooks.#"),
			resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "git.#"),
			resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "pages.#"),
			resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "importer.#"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() {},
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: anonymousConfiguration,
					Check:  anonymousCheck,
				},
			},
		})
	})

	t.Run("can be configured with an individual account", func(t *testing.T) {

		individualConfiguration := fmt.Sprintf(`
			provider "github" {
				owner = "%s"
				token = "%s"
			}
			data "github_user" "test" { username = "%s" }
		`, testOwner, testToken, testOwner)

		individualCheck := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_user.test", "name"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				requiredEnvironmentVariables := []string{
					"GITHUB_TOKEN",
					"GITHUB_OWNER",
				}
				for _, variable := range requiredEnvironmentVariables {
					if v := os.Getenv(variable); v == "" {
						t.Fatal(variable + " must be set for acceptance tests")
					}
				}
			},
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: individualConfiguration,
					Check:  individualCheck,
				},
			},
		})

	})

	t.Run("can be configured with an organization account", func(t *testing.T) {

		organizationConfiguration := fmt.Sprintf(`
			provider "github" {
				organization = "%s"
				token = "%s"
			}
			data "github_organization" "test" { name = "%s" }
		`, testOrganization, testToken, testOrganization)

		organizationCheck := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_organization.test", "plan"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				requiredEnvironmentVariables := []string{
					"GITHUB_TOKEN",
					"GITHUB_ORGANIZATION",
				}
				testAccPreCheckEnvironment(t, requiredEnvironmentVariables)
			},
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: organizationConfiguration,
					Check:  organizationCheck,
				},
			},
		})

	})
}
