package github

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

func TestAccGithubEnterpriseOrganization(t *testing.T) {
	t.Run("creates and updates an enterprise organization without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		orgName := fmt.Sprintf("tf-acc-test-%s", randomID)

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
			`, testEnterprise, orgName, desc)

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

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
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
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			if isEnterprise != "true" {
				t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
			}
			if testEnterprise == "" {
				t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
			}
			testCase(t, enterprise)
		})
	})

	t.Run("deletes an enterprise organization without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		orgName := fmt.Sprintf("tf-acc-test-%s", randomID)

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
			`, testEnterprise, orgName)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:  config,
						Destroy: true,
					},
				},
			})
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			if isEnterprise != "true" {
				t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
			}
			if testEnterprise == "" {
				t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
			}
			testCase(t, enterprise)
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
			`, testEnterprise, orgName, displayName, desc)

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

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
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
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			if isEnterprise != "true" {
				t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
			}
			if testEnterprise == "" {
				t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
			}
			testCase(t, enterprise)
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
			`, testEnterprise, orgName, desc)

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
			  `, testEnterprise, orgName, displayName, desc)

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

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
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
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			if isEnterprise != "true" {
				t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
			}
			if testEnterprise == "" {
				t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
			}
			testCase(t, enterprise)
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
			  `, testEnterprise, orgName)

		check := resource.ComposeTestCheckFunc()

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
					{
						ResourceName:      "github_enterprise_organization.org",
						ImportState:       true,
						ImportStateVerify: true,
						ImportStateId:     fmt.Sprintf(`%s/%s`, testEnterprise, orgName),
					},
				},
			})
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			if isEnterprise != "true" {
				t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
			}
			if testEnterprise == "" {
				t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
			}
			testCase(t, enterprise)
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
			  `, testEnterprise, orgName)

		check := resource.ComposeTestCheckFunc()

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
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
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			if isEnterprise != "true" {
				t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
			}
			if testEnterprise == "" {
				t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
			}
			testCase(t, enterprise)
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
			  `, testEnterprise, orgName)

		check := resource.ComposeTestCheckFunc()

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
					{
						ResourceName:  "github_enterprise_organization.org",
						ImportState:   true,
						ImportStateId: fmt.Sprintf(`%s/%s`, testEnterprise, randomID),
						ExpectError:   regexp.MustCompile("Could not resolve to an Organization with the login of .*"),
					},
				},
			})
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			if isEnterprise != "true" {
				t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
			}
			if testEnterprise == "" {
				t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
			}
			testCase(t, enterprise)
		})
	})
}

// TestEnterpriseOrganizationReadGraphQLErrorHandling tests the Read function's
// behavior when GraphQL returns "Could not resolve to a node" error.
// This can happen when:
// 1. The org was actually deleted
// 2. The org exists but the PAT hasn't been authorized for it yet (EMU/SSO)
func TestEnterpriseOrganizationReadGraphQLErrorHandling(t *testing.T) {
	graphqlNotFoundResponse := `{
		"data": { "node": null },
		"errors": [{
			"type": "NOT_FOUND",
			"path": ["node"],
			"message": "Could not resolve to a node with the global id of 'O_test123'"
		}]
	}`

	t.Run("returns error and preserves state when REST returns 404 (could be deleted or unauthorized)", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(graphqlNotFoundResponse))
		})
		mux.HandleFunc("/api/v3/orgs/test-org", func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "Not Found"}`))
		})

		server := httptest.NewServer(mux)
		defer server.Close()

		v3client, _ := github.NewClient(nil).WithEnterpriseURLs(server.URL+"/api/v3/", server.URL+"/")
		meta := &Owner{
			v4client: githubv4.NewClient(&http.Client{Transport: localRoundTripper{handler: mux}}),
			v3client: v3client,
		}

		resourceData := schema.TestResourceDataRaw(t, resourceGithubEnterpriseOrganization().Schema, map[string]interface{}{
			"name":          "test-org",
			"enterprise_id": "E_test",
			"billing_email": "test@example.com",
			"admin_logins":  []interface{}{"admin"},
		})
		resourceData.SetId("O_test123")

		err := resourceGithubEnterpriseOrganizationRead(resourceData, meta)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !strings.Contains(err.Error(), "cannot read organization") {
			t.Fatalf("expected 'cannot read organization' error, got: %v", err)
		}
		if !strings.Contains(err.Error(), "terraform state rm") {
			t.Fatalf("expected guidance to use 'terraform state rm', got: %v", err)
		}
		if resourceData.Id() == "" {
			t.Fatal("expected resource ID to NOT be cleared (to prevent accidental destruction)")
		}
	})

	t.Run("returns error when org exists but GraphQL can't access it (REST 200)", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(graphqlNotFoundResponse))
		})
		mux.HandleFunc("/api/v3/orgs/test-org", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"login": "test-org", "id": 123}`))
		})

		server := httptest.NewServer(mux)
		defer server.Close()

		v3client, _ := github.NewClient(nil).WithEnterpriseURLs(server.URL+"/api/v3/", server.URL+"/")
		meta := &Owner{
			v4client: githubv4.NewClient(&http.Client{Transport: localRoundTripper{handler: mux}}),
			v3client: v3client,
		}

		resourceData := schema.TestResourceDataRaw(t, resourceGithubEnterpriseOrganization().Schema, map[string]interface{}{
			"name":          "test-org",
			"enterprise_id": "E_test",
			"billing_email": "test@example.com",
			"admin_logins":  []interface{}{"admin"},
		})
		resourceData.SetId("O_test123")

		err := resourceGithubEnterpriseOrganizationRead(resourceData, meta)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !strings.Contains(err.Error(), "exists but cannot be read via GraphQL") {
			t.Fatalf("expected PAT authorization error message, got: %v", err)
		}
		if !strings.Contains(err.Error(), "authorize the PAT") {
			t.Fatalf("expected guidance to authorize PAT, got: %v", err)
		}
		if resourceData.Id() == "" {
			t.Fatal("expected resource ID to NOT be cleared")
		}
	})

	t.Run("returns error and preserves state when REST fails with 403", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(graphqlNotFoundResponse))
		})
		mux.HandleFunc("/api/v3/orgs/test-org", func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"message": "Forbidden"}`))
		})

		server := httptest.NewServer(mux)
		defer server.Close()

		v3client, _ := github.NewClient(nil).WithEnterpriseURLs(server.URL+"/api/v3/", server.URL+"/")
		meta := &Owner{
			v4client: githubv4.NewClient(&http.Client{Transport: localRoundTripper{handler: mux}}),
			v3client: v3client,
		}

		resourceData := schema.TestResourceDataRaw(t, resourceGithubEnterpriseOrganization().Schema, map[string]interface{}{
			"name":          "test-org",
			"enterprise_id": "E_test",
			"billing_email": "test@example.com",
			"admin_logins":  []interface{}{"admin"},
		})
		resourceData.SetId("O_test123")

		err := resourceGithubEnterpriseOrganizationRead(resourceData, meta)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !strings.Contains(err.Error(), "cannot read organization") {
			t.Fatalf("expected 'cannot read organization' error, got: %v", err)
		}
		if !strings.Contains(err.Error(), "GraphQL error") {
			t.Fatalf("expected GraphQL error in message, got: %v", err)
		}
		if !strings.Contains(err.Error(), "REST error") {
			t.Fatalf("expected REST error in message, got: %v", err)
		}
		if resourceData.Id() == "" {
			t.Fatal("expected resource ID to NOT be cleared")
		}
	})
}
