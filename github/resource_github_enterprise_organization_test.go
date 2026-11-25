package github

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestIsSAMLEnforcementError(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			result := isSAMLEnforcementError(tc.err)
			if result != tc.expected {
				t.Errorf("isSAMLEnforcementError(%v) = %v, want %v", tc.err, result, tc.expected)
			}
		})
	}
}

func TestAccGithubEnterpriseOrganization(t *testing.T) {
	t.Parallel()

	t.Run("creates and updates an enterprise organization without error", func(t *testing.T) {
		t.Parallel()

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
				resource.TestCheckResourceAttrSet(
					"github_enterprise_organization.org", "database_id",
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
			PreCheck:          func() { skipUnlessEnterprise(t) },
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
		t.Parallel()

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
			PreCheck:          func() { skipUnlessEnterprise(t) },
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
		t.Parallel()

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
			PreCheck:          func() { skipUnlessEnterprise(t) },
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
						updatedDesc, 1,
					),
					Check: checks["after"],
				},
			},
		})
	})

	t.Run("creates org without display name, set and update display name", func(t *testing.T) {
		t.Parallel()

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
			PreCheck:          func() { skipUnlessEnterprise(t) },
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
						updatedDesc, 1,
					),
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
		t.Parallel()

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
			PreCheck:          func() { skipUnlessEnterprise(t) },
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
		t.Parallel()

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
			PreCheck:          func() { skipUnlessEnterprise(t) },
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
		t.Parallel()

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
			PreCheck:          func() { skipUnlessEnterprise(t) },
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

func TestResourceGithubEnterpriseOrganizationCreateSetsDatabaseID(t *testing.T) {
	t.Parallel()

	const createResponse = `{
  "data": {
    "createEnterpriseOrganization": {
      "organization": {
        "id": "O_kgDOCg7Zxw",
        "databaseId": 168828871
      }
    }
  }
}`

	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		body := mustRead(req.Body)
		if !strings.Contains(body, "createEnterpriseOrganization") {
			t.Errorf("unexpected GraphQL call: %s", body)
		}
		if !strings.Contains(body, "databaseId") {
			t.Errorf("mutation does not select databaseId: %s", body)
		}
		w.Header().Set("Content-Type", "application/json")
		mustWrite(w, createResponse)
	})

	meta := &Owner{
		v4client: newTestGraphQLClient(mux),
	}

	data := schema.TestResourceDataRaw(t, resourceGithubEnterpriseOrganization().Schema, map[string]any{
		"enterprise_id": "E_kgDNAbc",
		"name":          "some-awesome-org",
		"billing_email": "octocat@octo.cat",
		"admin_logins":  []any{"octocat"},
	})

	if err := resourceGithubEnterpriseOrganizationCreate(data, meta); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if got, want := data.Id(), "O_kgDOCg7Zxw"; got != want {
		t.Errorf("id = %q, want %q", got, want)
	}

	// database_id must be populated during create because the provider does not read after write,
	// and other resources reference it within the same apply.
	got, ok := data.Get("database_id").(int)
	if !ok {
		t.Fatalf("database_id is %T, want int", data.Get("database_id"))
	}
	if want := 168828871; got != want {
		t.Errorf("database_id = %d, want %d", got, want)
	}
}

func TestResourceGithubEnterpriseOrganizationReadBillingEmail(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		status       int
		body         string
		wantEmail    string
		wantError    string
		graphqlError string
		wantID       string
	}{
		{name: "REST billing email", status: http.StatusOK, body: `{"billing_email":"billing@example.com"}`, wantEmail: "billing@example.com", wantID: "O_example"},
		{name: "empty billing email", status: http.StatusOK, body: `{"billing_email":""}`, wantID: "O_example"},
		{name: "missing billing email", status: http.StatusOK, body: `{}`, wantError: "billing_email", wantID: "O_example"},
		{name: "null billing email", status: http.StatusOK, body: `{"billing_email":null}`, wantError: "billing_email", wantID: "O_example"},
		{name: "REST forbidden", status: http.StatusForbidden, body: `{"message":"Resource not accessible by integration"}`, wantError: "Resource not accessible by integration", wantID: "O_example"},
		{name: "REST not found after GraphQL success", status: http.StatusNotFound, body: `{"message":"Not Found"}`, wantError: "Not Found", wantID: "O_example"},
		{name: "GraphQL forbidden", graphqlError: "Resource not accessible by integration", wantError: "Resource not accessible by integration", wantID: "O_example"},
		{name: "GraphQL node not found", graphqlError: "Could not resolve to a node with the global id of 'O_example'"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			graphqlCalls := 0
			restCalls := 0
			mux := http.NewServeMux()
			mux.HandleFunc("/graphql", func(writer http.ResponseWriter, request *http.Request) {
				graphqlCalls++
				writer.Header().Set("Content-Type", "application/json")
				body := mustRead(request.Body)
				if strings.Contains(body, "organizationBillingEmail") {
					mustWrite(writer, `{"errors":[{"message":"Resource not accessible by integration","path":["node","organizationBillingEmail"]}]}`)
					return
				}
				if test.graphqlError != "" {
					mustWrite(writer, fmt.Sprintf(`{"errors":[{"message":%q}]}`, test.graphqlError))
					return
				}
				if graphqlCalls == 1 {
					if !strings.Contains(body, `"cursor":null`) {
						t.Errorf("expected initial null cursor: %s", body)
					}
					mustWrite(writer, `{"data":{"node":{"id":"O_example","databaseId":123,"login":"example-org","name":"Example Organization","description":"Example description","membersWithRole":{"edges":[{"node":{"login":"owner-one"},"role":"ADMIN"},{"node":{"login":"member"},"role":"MEMBER"}],"pageInfo":{"hasNextPage":true,"endCursor":"next-page"}}}}}`)
					return
				}
				if graphqlCalls != 2 || !strings.Contains(body, `"cursor":"next-page"`) {
					t.Errorf("unexpected pagination request: %s", body)
				}
				mustWrite(writer, `{"data":{"node":{"id":"O_example","databaseId":123,"login":"example-org","name":"Example Organization","description":"Example description","membersWithRole":{"edges":[{"node":{"login":"owner-two"},"role":"ADMIN"}],"pageInfo":{"hasNextPage":false,"endCursor":"last-page"}}}}}`)
			})
			mux.HandleFunc("/orgs/example-org", func(writer http.ResponseWriter, request *http.Request) {
				restCalls++
				if request.Method != http.MethodGet {
					t.Errorf("method = %s, want GET", request.Method)
				}
				writer.Header().Set("Content-Type", "application/json")
				writer.WriteHeader(test.status)
				mustWrite(writer, test.body)
			})

			meta := &Owner{
				v3client:   mustCreateTestGitHubClient(t, "https://api.github.com/", github.WithTransport(localRoundTripper{handler: mux})),
				v4client:   newTestGraphQLClient(mux),
				maxPerPage: 100,
			}
			data := schema.TestResourceDataRaw(t, resourceGithubEnterpriseOrganization().Schema, map[string]any{
				"enterprise_id": "E_example",
				"name":          "stale-org-name",
				"billing_email": "previous@example.com",
				"admin_logins":  []any{"previous-owner"},
			})
			data.SetId("O_example")

			err := resourceGithubEnterpriseOrganizationRead(data, meta)
			if test.wantError != "" {
				if err == nil || !strings.Contains(err.Error(), test.wantError) {
					t.Fatalf("error = %v, want containing %q", err, test.wantError)
				}
				if got := data.Get("billing_email"); got != "previous@example.com" {
					t.Errorf("billing_email = %q after error, want previous value", got)
				}
			} else if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if got := data.Id(); got != test.wantID {
				t.Errorf("id = %q, want %q", got, test.wantID)
			}
			if test.graphqlError != "" {
				if restCalls != 0 {
					t.Errorf("REST calls = %d after GraphQL error, want 0", restCalls)
				}
				return
			}
			if graphqlCalls != 2 || restCalls != 1 {
				t.Errorf("GraphQL calls = %d, REST calls = %d, want 2 and 1", graphqlCalls, restCalls)
			}
			if test.wantError != "" {
				return
			}
			for attribute, want := range map[string]any{
				"billing_email": test.wantEmail,
				"name":          "example-org",
				"database_id":   123,
				"display_name":  "Example Organization",
				"description":   "Example description",
			} {
				if got := data.Get(attribute); got != want {
					t.Errorf("%s = %v, want %v", attribute, got, want)
				}
			}
			owners := data.Get("admin_logins").(*schema.Set)
			if owners.Len() != 2 || !owners.Contains("owner-one") || !owners.Contains("owner-two") {
				t.Errorf("admin_logins = %v, want both owners from pagination", owners.List())
			}
		})
	}
}
