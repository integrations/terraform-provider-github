package github

import (
	"context"
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
	"github.com/shurcooL/githubv4"
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

func TestResourceGithubEnterpriseOrganizationCreateUsesDisplayNameAsProfileName(t *testing.T) {
	t.Parallel()

	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		body := mustRead(req.Body)
		if !strings.Contains(body, `"profileName":"Some Awesome Org"`) {
			t.Errorf("mutation does not use display_name as profileName: %s", body)
		}
		w.Header().Set("Content-Type", "application/json")
		mustWrite(w, `{"data":{"createEnterpriseOrganization":{"organization":{"id":"O_kgDOCg7Zxw","databaseId":1}}}}`)
	})

	meta := &Owner{v4client: newTestGraphQLClient(mux)}

	data := schema.TestResourceDataRaw(t, resourceGithubEnterpriseOrganization().Schema, map[string]any{
		"enterprise_id": "E_kgDNAbc",
		"name":          "some-awesome-org",
		"display_name":  "Some Awesome Org",
		"billing_email": "octocat@octo.cat",
		"admin_logins":  []any{"octocat"},
	})

	if err := resourceGithubEnterpriseOrganizationCreate(data, meta); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResourceGithubEnterpriseOrganizationReadFailsWithoutOrganizationAccess(t *testing.T) {
	t.Parallel()

	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		body := mustRead(req.Body)
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(body, "membersWithRole") {
			mustWrite(w, `{"data":{"node":{"membersWithRole":null}},"errors":[{"type":"FORBIDDEN","message":"Resource not accessible by integration"}]}`)
			return
		}
		mustWrite(w, `{"data":{"node":{"id":"O_kgDOCg7Zxw","databaseId":168828871,"name":"Some Awesome Org","login":"some-awesome-org","description":"Created with terraform"}}}`)
	})

	meta := &Owner{v4client: newTestGraphQLClient(mux), maxPerPage: 100}

	data := schema.TestResourceDataRaw(t, resourceGithubEnterpriseOrganization().Schema, map[string]any{
		"enterprise_id": "E_kgDNAbc",
		"name":          "some-awesome-org",
		"billing_email": "octocat@octo.cat",
		"admin_logins":  []any{"octocat"},
	})
	data.SetId("O_kgDOCg7Zxw")

	err := resourceGithubEnterpriseOrganizationRead(data, meta)
	if err == nil || !strings.Contains(err.Error(), "Resource not accessible by integration") {
		t.Fatalf("expected not accessible error, got %v", err)
	}
}

func TestResourceGithubEnterpriseOrganizationReadWithOrganizationAccess(t *testing.T) {
	t.Parallel()

	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		body := mustRead(req.Body)
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(body, "membersWithRole") {
			mustWrite(w, `{"data":{"node":{"membersWithRole":{"edges":[{"node":{"login":"octocat"},"role":"ADMIN"},{"node":{"login":"hubot"},"role":"MEMBER"}],"pageInfo":{"hasNextPage":false,"endCursor":""}}}}}`)
			return
		}
		mustWrite(w, `{"data":{"node":{"id":"O_kgDOCg7Zxw","databaseId":168828871,"name":"some-awesome-org","login":"some-awesome-org","description":""}}}`)
	})
	mux.HandleFunc("/orgs/some-awesome-org", func(w http.ResponseWriter, req *http.Request) {
		mustWrite(w, `{"login":"some-awesome-org","billing_email":"billing@octo.cat"}`)
	})

	v3, _ := github.NewClient(github.WithHTTPClient(&http.Client{Transport: localRoundTripper{handler: mux}}))
	meta := &Owner{v3client: v3, v4client: newTestGraphQLClient(mux), maxPerPage: 100}

	data := schema.TestResourceDataRaw(t, resourceGithubEnterpriseOrganization().Schema, map[string]any{
		"enterprise_id": "E_kgDNAbc",
		"name":          "some-awesome-org",
		"billing_email": "old@octo.cat",
		"admin_logins":  []any{"old"},
	})
	data.SetId("O_kgDOCg7Zxw")

	if err := resourceGithubEnterpriseOrganizationRead(data, meta); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if got, want := data.Get("billing_email").(string), "billing@octo.cat"; got != want {
		t.Errorf("billing_email = %q, want %q", got, want)
	}
	if got := data.Get("admin_logins").(*schema.Set).List(); len(got) != 1 || got[0] != "octocat" {
		t.Errorf("admin_logins = %v, want [octocat]", got)
	}
}

// fakeAppSource routes every client to the same test server; the test server tells the calls apart by
// the Authorization header the fake sets per installation.
type fakeAppSource struct {
	handler http.Handler
}

func (s *fakeAppSource) client(auth string) *http.Client {
	return &http.Client{Transport: authRoundTripper{auth: auth, next: localRoundTripper{handler: s.handler}}}
}

func (s *fakeAppSource) RESTClient() (*github.Client, error) {
	return github.NewClient(github.WithHTTPClient(s.client("app")))
}

func (s *fakeAppSource) OwnerRESTClient(_ context.Context, owner string) (*github.Client, error) {
	return github.NewClient(github.WithHTTPClient(s.client("org:" + owner)))
}

func (s *fakeAppSource) GraphQLClient() (*githubv4.Client, error) {
	return githubv4.NewClient(s.client("app")), nil
}

func (s *fakeAppSource) OwnerGraphQLClient(_ context.Context, owner string) (*githubv4.Client, error) {
	return githubv4.NewClient(s.client("org:" + owner)), nil
}

type authRoundTripper struct {
	auth string
	next http.RoundTripper
}

func (rt authRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-Test-Auth", rt.auth)
	return rt.next.RoundTrip(req)
}

// With an enterprise-installed GitHub App the organization is created through the enterprise
// installation, the app is then installed on the new organization, and organization-level calls use
// that installation.
func TestResourceGithubEnterpriseOrganizationCreateWithEnterpriseAppInstallation(t *testing.T) {
	t.Parallel()

	var calls []string
	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		body := mustRead(req.Body)
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.Contains(body, "createEnterpriseOrganization"):
			calls = append(calls, "create:"+req.Header.Get("X-Test-Auth"))
			mustWrite(w, `{"data":{"createEnterpriseOrganization":{"organization":{"id":"O_kgDOCg7Zxw","databaseId":1}}}}`)
		case strings.Contains(body, "Enterprise"):
			calls = append(calls, "slug:"+req.Header.Get("X-Test-Auth"))
			mustWrite(w, `{"data":{"node":{"slug":"octo-enterprise"}}}`)
		default:
			t.Errorf("unexpected GraphQL call: %s", body)
		}
	})
	mux.HandleFunc("/app", func(w http.ResponseWriter, req *http.Request) {
		calls = append(calls, "app:"+req.Header.Get("X-Test-Auth"))
		mustWrite(w, `{"id":1,"slug":"octo-app","client_id":"Iv1.abc"}`)
	})
	mux.HandleFunc("/enterprises/octo-enterprise/apps/organizations/some-awesome-org/installations", func(w http.ResponseWriter, req *http.Request) {
		body := mustRead(req.Body)
		if req.Method != http.MethodPost || !strings.Contains(body, `"client_id":"Iv1.abc"`) {
			t.Errorf("unexpected install call: %s %s", req.Method, body)
		}
		calls = append(calls, "install:"+req.Header.Get("X-Test-Auth"))
		mustWrite(w, `{"id":42}`)
	})
	mux.HandleFunc("/orgs/some-awesome-org", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPatch {
			t.Errorf("unexpected org call: %s", req.Method)
		}
		calls = append(calls, "edit:"+req.Header.Get("X-Test-Auth"))
		mustWrite(w, `{"login":"some-awesome-org"}`)
	})

	enterprise := &fakeAppSource{handler: mux}
	enterpriseV3, _ := enterprise.OwnerRESTClient(context.Background(), "enterprise")
	enterpriseV4, _ := enterprise.OwnerGraphQLClient(context.Background(), "enterprise")
	meta := &Owner{v3client: enterpriseV3, v4client: enterpriseV4, appSource: enterprise}

	data := schema.TestResourceDataRaw(t, resourceGithubEnterpriseOrganization().Schema, map[string]any{
		"enterprise_id": "E_kgDNAbc",
		"name":          "some-awesome-org",
		"description":   "Created with terraform",
		"billing_email": "octocat@octo.cat",
		"admin_logins":  []any{"octocat"},
	})

	if err := resourceGithubEnterpriseOrganizationCreate(data, meta); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	want := []string{"create:org:enterprise", "app:app", "slug:org:enterprise", "install:org:enterprise", "edit:org:some-awesome-org"}
	if strings.Join(calls, ",") != strings.Join(want, ",") {
		t.Errorf("calls = %v, want %v", calls, want)
	}
}
