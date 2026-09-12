package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccGithubMembership(t *testing.T) {
	t.Parallel()

	if len(testAccConf.testExternalUser1) == 0 {
		t.Skip("No external user provided")
	}

	t.Run("creates organization membership", func(t *testing.T) {
		// IMPORTANT: Do not run this sub test in parallel is it uses shared state.

		ctx := t.Context()

		var membership github.Membership
		rn := "github_membership.test_org_membership"

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubMembershipDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccGithubMembershipConfig(testAccConf.testExternalUser1),
					Check: resource.ComposeTestCheckFunc(
						testAccCheckGithubMembershipExists(ctx, rn, &membership),
						testAccCheckGithubMembershipRoleState(ctx, rn, &membership),
					),
				},
				{
					ResourceName:      rn,
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateVerifyIgnore: []string{
						"downgrade_on_destroy",
						"downgrade_to",
					},
				},
			},
		})
	})

	t.Run("creates organization membership with downgrade", func(t *testing.T) {
		// IMPORTANT: Do not run this sub test in parallel is it uses shared state.

		ctx := t.Context()

		var membership github.Membership
		rn := "github_membership.test_org_membership"

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubMembershipDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccGithubMembershipConfigDowngradable(testAccConf.testExternalUser1),
					Check: resource.ComposeTestCheckFunc(
						testAccCheckGithubMembershipExists(ctx, rn, &membership),
						testAccCheckGithubMembershipRoleState(ctx, rn, &membership),
					),
				},
				{
					ResourceName: rn,
					ImportState:  true,
				},
			},
		})
	})

	t.Run("downgrades organization membership to outside collaborator", func(t *testing.T) {
		// IMPORTANT: Do not run this sub test in parallel is it uses shared state.

		if len(testAccConf.testExternalUser1Token) == 0 {
			t.Skip("No external user token provided")
		}

		ctx := t.Context()
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%smembership-downgrade-%s", testResourcePrefix, randomID)
		username := testAccConf.testExternalUser1
		rn := "github_membership.test_org_membership"

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessHasOrgs(t)
				t.Cleanup(func() {
					_, _ = testAccConf.meta.v3client.Organizations.RemoveOutsideCollaborator(context.Background(), testAccConf.owner, username)
				})
			},
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubMembershipDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccGithubMembershipConfigOutsideCollaboratorDowngradePending(username, repoName),
					Check: resource.ComposeTestCheckFunc(
						testAccCheckGithubMembershipState(t, ctx, rn, "pending"),
					),
				},
				{
					PreConfig: func() {
						testAccGithubMembershipAcceptOrgInvitation(t, testAccConf.testExternalUser1Token)
						testAccGithubMembershipAddRepositoryCollaborator(t, repoName, username)
					},
					Config: testAccGithubMembershipConfigOutsideCollaboratorDowngrade(username, repoName),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(rn, "downgrade_to", membershipDowngradeToOutsideCollaborator),
						testAccCheckGithubMembershipState(t, ctx, rn, membershipStateActive),
					),
				},
				{
					Config: testAccGithubMembershipConfigOutsideCollaboratorDowngradeCleanup(username, repoName),
					Check:  testAccCheckGithubMembershipOutsideCollaborator(t, ctx, username),
				},
			},
		})
	})

	t.Run("creates organization membership with case insensitivity", func(t *testing.T) {
		// IMPORTANT: Do not run this sub test in parallel is it uses shared state.

		ctx := t.Context()

		var membership github.Membership
		var otherMembership github.Membership

		rn := "github_membership.test_org_membership"
		otherCase := flipUsernameCase(testAccConf.testExternalUser1)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubMembershipDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccGithubMembershipConfig(testAccConf.testExternalUser1),
					Check: resource.ComposeTestCheckFunc(
						testAccCheckGithubMembershipExists(ctx, rn, &membership),
					),
				},
				{
					Config: testAccGithubMembershipConfig(otherCase),
					Check: resource.ComposeTestCheckFunc(
						testAccCheckGithubMembershipExists(ctx, rn, &otherMembership),
						testAccGithubMembershipTheSame(&membership, &otherMembership),
					),
				},
				{
					ResourceName:      rn,
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateVerifyIgnore: []string{
						"downgrade_on_destroy",
						"downgrade_to",
					},
				},
			},
		})
	})
}

func testAccCheckGithubMembershipDestroy(s *terraform.State) error {
	ctx := context.Background()
	conn := testAccConf.meta.v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_membership" {
			continue
		}

		orgName, username, err := parseID2(rs.Primary.ID)
		if err != nil {
			return err
		}

		downgradedOnDestroy := rs.Primary.Attributes["downgrade_on_destroy"] == "true"
		downgradeTo := rs.Primary.Attributes["downgrade_to"]
		membership, resp, err := conn.Organizations.GetOrgMembership(ctx, username, orgName)
		responseIsSuccessful := err == nil && membership != nil && buildTwoPartID(orgName, username) == rs.Primary.ID

		if downgradedOnDestroy {
			if downgradeTo == membershipDowngradeToOutsideCollaborator {
				if responseIsSuccessful {
					return fmt.Errorf("organization membership %q still exists", rs.Primary.ID)
				}
				if resp == nil || resp.StatusCode != http.StatusNotFound {
					return err
				}

				isOutsideCollaborator, err := testAccGithubMembershipIsOutsideCollaborator(ctx, conn, orgName, username)
				if err != nil {
					return err
				}
				if !isOutsideCollaborator {
					return fmt.Errorf("organization membership %q was not converted to an outside collaborator", rs.Primary.ID)
				}

				// Now actually remove them from the org to clean up
				_, removeErr := conn.Organizations.RemoveOutsideCollaborator(ctx, orgName, username)
				if removeErr != nil {
					return fmt.Errorf("outside collaborator %q could not be removed during membership downgrade test case cleanup: %w", rs.Primary.ID, removeErr)
				}

				return nil
			}

			if !responseIsSuccessful {
				return fmt.Errorf("could not load organization membership for %q", rs.Primary.ID)
			}

			if *membership.Role != "member" {
				return fmt.Errorf("organization membership %q is not a member of the org or is not the 'member' role", rs.Primary.ID)
			}

			// Now actually remove them from the org to clean up
			_, removeErr := conn.Organizations.RemoveOrgMembership(ctx, username, orgName)
			if removeErr != nil {
				return fmt.Errorf("organization membership %q could not be removed during membership downgrade test case cleanup: %w", rs.Primary.ID, removeErr)
			}
		} else if responseIsSuccessful {
			return fmt.Errorf("organization membership %q still exists", rs.Primary.ID)
		} else if resp.StatusCode != 404 {
			return err
		}

		return nil
	}
	return nil
}

func testAccCheckGithubMembershipOutsideCollaborator(t *testing.T, ctx context.Context, username string) resource.TestCheckFunc {
	t.Helper()

	return func(s *terraform.State) error {
		conn := testAccConf.meta.v3client
		orgName := testAccConf.owner

		return retry.RetryContext(ctx, 30*time.Second, func() *retry.RetryError {
			membership, resp, err := conn.Organizations.GetOrgMembership(ctx, username, orgName)
			if err == nil && membership != nil {
				return retry.RetryableError(fmt.Errorf("organization membership %q still exists", buildTwoPartID(orgName, username)))
			}
			if resp == nil || resp.StatusCode != http.StatusNotFound {
				return retry.NonRetryableError(err)
			}

			isOutsideCollaborator, err := testAccGithubMembershipIsOutsideCollaborator(ctx, conn, orgName, username)
			if err != nil {
				return retry.NonRetryableError(err)
			}
			if !isOutsideCollaborator {
				return retry.RetryableError(fmt.Errorf("%q is not an outside collaborator for organization %q", username, orgName))
			}

			return nil
		})
	}
}

func testAccGithubMembershipAcceptOrgInvitation(t *testing.T, inviteeToken string) {
	t.Helper()

	client, err := testAccGithubMembershipInviteeClient(inviteeToken)
	if err != nil {
		t.Fatalf("failed to create invitee GitHub client: %s", err)
	}

	err = retry.RetryContext(t.Context(), 30*time.Second, func() *retry.RetryError {
		_, resp, err := client.Organizations.EditOrgMembership(t.Context(), "", testAccConf.owner, &github.Membership{
			State: new(membershipStateActive),
		})
		if err != nil {
			statusCode := 0
			if resp != nil {
				statusCode = resp.StatusCode
			}
			if statusCode == http.StatusForbidden || statusCode == http.StatusNotFound || statusCode == http.StatusUnauthorized {
				return retry.NonRetryableError(fmt.Errorf("failed to accept organization invitation: %w", err))
			}
			return retry.RetryableError(fmt.Errorf("failed to accept organization invitation: %w", err))
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func testAccGithubMembershipAddRepositoryCollaborator(t *testing.T, repoName, username string) {
	t.Helper()

	_, _, err := testAccConf.meta.v3client.Repositories.AddCollaborator(t.Context(),
		testAccConf.owner,
		repoName,
		username,
		&github.RepositoryAddCollaboratorOptions{Permission: "push"},
	)
	if err != nil {
		t.Fatalf("failed to add repository collaborator %q on %q: %s", username, repoName, err)
	}
}

func testAccGithubMembershipInviteeClient(inviteeToken string) (*github.Client, error) {
	config := &Config{
		BaseURL: testAccConf.baseURL,
		IsGHES:  testAccConf.isGHES,
		Token:   inviteeToken,
	}
	return config.NewRESTClient(config.AuthenticatedHTTPClient())
}

func testAccGithubMembershipIsOutsideCollaborator(ctx context.Context, conn *github.Client, orgName, username string) (bool, error) {
	opts := &github.ListOutsideCollaboratorsOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	for {
		users, resp, err := conn.Organizations.ListOutsideCollaborators(ctx, orgName, opts)
		if err != nil {
			return false, err
		}
		for _, user := range users {
			if strings.EqualFold(user.GetLogin(), username) {
				return true, nil
			}
		}
		if resp == nil || resp.NextPage == 0 {
			return false, nil
		}
		opts.Page = resp.NextPage
	}
}

func testAccCheckGithubMembershipState(t *testing.T, ctx context.Context, n, expectedState string) resource.TestCheckFunc {
	t.Helper()

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no membership ID is set")
		}

		conn := testAccConf.meta.v3client

		orgName, username, err := parseID2(rs.Primary.ID)
		if err != nil {
			return err
		}

		membership, _, err := conn.Organizations.GetOrgMembership(ctx, username, orgName)
		if err != nil {
			return err
		}

		if membership.GetState() != expectedState {
			return fmt.Errorf("expected membership state %q, got %q", expectedState, membership.GetState())
		}

		return nil
	}
}

func testAccCheckGithubMembershipExists(ctx context.Context, n string, membership *github.Membership) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no membership ID is set")
		}

		conn := testAccConf.meta.v3client

		orgName, username, err := parseID2(rs.Primary.ID)
		if err != nil {
			return err
		}

		githubMembership, _, err := conn.Organizations.GetOrgMembership(ctx, username, orgName)
		if err != nil {
			return err
		}
		*membership = *githubMembership
		return nil
	}
}

func testAccCheckGithubMembershipRoleState(ctx context.Context, n string, membership *github.Membership) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no membership ID is set")
		}

		conn := testAccConf.meta.v3client

		orgName, username, err := parseID2(rs.Primary.ID)
		if err != nil {
			return err
		}

		githubMembership, _, err := conn.Organizations.GetOrgMembership(ctx, username, orgName)
		if err != nil {
			return err
		}

		resourceRole := membership.GetRole()
		actualRole := githubMembership.GetRole()

		if resourceRole != actualRole {
			return fmt.Errorf("membership role %v in resource does match actual state of %v",
				resourceRole, actualRole)
		}
		return nil
	}
}

func testAccGithubMembershipConfig(username string) string {
	return fmt.Sprintf(`
  resource "github_membership" "test_org_membership" {
    username = "%s"
    role = "member"
  }
`, username)
}

func testAccGithubMembershipConfigDowngradable(username string) string {
	return fmt.Sprintf(`
  resource "github_membership" "test_org_membership" {
    username = "%s"
    role = "admin"
    downgrade_on_destroy = %t
  }
`, username, true)
}

func testAccGithubMembershipConfigOutsideCollaboratorDowngradePending(username, repoName string) string {
	return fmt.Sprintf(`
  resource "github_membership" "test_org_membership" {
    username = "%[1]s"
    role = "member"
  }

  resource "github_repository" "test" {
    name      = "%[2]s"
    auto_init = true
  }
`, username, repoName)
}

func testAccGithubMembershipConfigOutsideCollaboratorDowngrade(username, repoName string) string {
	return fmt.Sprintf(`
  resource "github_membership" "test_org_membership" {
    username = "%[1]s"
    role = "member"
    downgrade_on_destroy = true
    downgrade_to = "%[3]s"
  }

  resource "github_repository" "test" {
    name      = "%[2]s"
    auto_init = true
  }
`, username, repoName, membershipDowngradeToOutsideCollaborator)
}

func testAccGithubMembershipConfigOutsideCollaboratorDowngradeCleanup(username, repoName string) string {
	return fmt.Sprintf(`
  resource "github_repository" "test" {
    name      = "%[2]s"
    auto_init = true
  }
`, username, repoName)
}

func testAccGithubMembershipTheSame(orig, other *github.Membership) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if orig.GetURL() != other.GetURL() {
			return errors.New("users are different")
		}

		return nil
	}
}

func Test_resourceGithubMembershipDelete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		downgradeOnDestroy bool
		downgradeTo        string
		setDowngradeTo     bool
		expectError        string
		responses          []*mockResponse
	}{
		{
			name: "removes membership",
			responses: []*mockResponse{
				{
					ExpectedMethod: http.MethodDelete,
					ExpectedUri:    "/orgs/test-org/memberships/octocat",
					StatusCode:     http.StatusNoContent,
				},
			},
		},
		{
			name:               "member default",
			downgradeOnDestroy: true,
			responses: []*mockResponse{
				{
					ExpectedMethod: http.MethodGet,
					ExpectedUri:    "/orgs/test-org/memberships/octocat",
					StatusCode:     http.StatusOK,
					ResponseBody:   `{"role":"admin"}`,
				},
				{
					ExpectedMethod: http.MethodPut,
					ExpectedUri:    "/orgs/test-org/memberships/octocat",
					StatusCode:     http.StatusOK,
					ResponseBody:   `{"role":"member"}`,
				},
			},
		},
		{
			name:               "outside collaborator",
			downgradeOnDestroy: true,
			downgradeTo:        membershipDowngradeToOutsideCollaborator,
			setDowngradeTo:     true,
			responses: []*mockResponse{
				{
					ExpectedMethod: http.MethodGet,
					ExpectedUri:    "/orgs/test-org/memberships/octocat",
					StatusCode:     http.StatusOK,
					ResponseBody:   `{"role":"member","state":"active"}`,
				},
				{
					ExpectedMethod: http.MethodPut,
					ExpectedUri:    "/orgs/test-org/outside_collaborators/octocat",
					StatusCode:     http.StatusNoContent,
				},
			},
		},
		{
			name:               "outside collaborator pending invitation",
			downgradeOnDestroy: true,
			downgradeTo:        membershipDowngradeToOutsideCollaborator,
			setDowngradeTo:     true,
			expectError:        `cannot downgrade "octocat" to outside collaborator because organization membership is "pending", not active; the user must accept the organization invitation first`,
			responses: []*mockResponse{
				{
					ExpectedMethod: http.MethodGet,
					ExpectedUri:    "/orgs/test-org/memberships/octocat",
					StatusCode:     http.StatusOK,
					ResponseBody:   `{"role":"member","state":"pending"}`,
				},
			},
		},
		{
			name:               "membership not found",
			downgradeOnDestroy: true,
			downgradeTo:        membershipDowngradeToOutsideCollaborator,
			setDowngradeTo:     true,
			responses: []*mockResponse{
				{
					ExpectedMethod: http.MethodGet,
					ExpectedUri:    "/orgs/test-org/memberships/octocat",
					StatusCode:     http.StatusNotFound,
					ResponseBody:   `{"message":"Not Found"}`,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ts := githubApiMock(tt.responses)
			defer ts.Close()

			raw := map[string]any{
				"username":             "octocat",
				"downgrade_on_destroy": tt.downgradeOnDestroy,
			}
			if tt.setDowngradeTo {
				raw["downgrade_to"] = tt.downgradeTo
			}
			d := schema.TestResourceDataRaw(t, resourceGithubMembership().Schema, raw)
			d.SetId(buildTwoPartID("test-org", "octocat"))

			meta := &Owner{
				name:           "test-org",
				v3client:       mustCreateTestGitHubClient(t, ts.URL),
				IsOrganization: true,
			}

			diags := resourceGithubMembershipDelete(t.Context(), d, meta)
			if tt.expectError != "" {
				if !diags.HasError() {
					t.Fatalf("expected diagnostics containing %q, got none", tt.expectError)
				}
				if !strings.Contains(diags[0].Summary, tt.expectError) {
					t.Fatalf("expected diagnostics containing %q, got: %v", tt.expectError, diags)
				}
				return
			}
			if diags.HasError() {
				t.Fatalf("expected no diagnostics, got: %v", diags)
			}
		})
	}
}

func Test_resourceGithubMembershipDowngradeToValidation(t *testing.T) {
	t.Parallel()

	validate := resourceGithubMembership().Schema["downgrade_to"].ValidateDiagFunc
	for _, value := range []string{membershipDowngradeToMember, membershipDowngradeToOutsideCollaborator} {
		if diags := validate(value, cty.Path{cty.GetAttrStep{Name: "downgrade_to"}}); diags.HasError() {
			t.Fatalf("expected %q to be valid, got: %v", value, diags)
		}
	}

	for _, value := range []string{"outside-collaborator", "outside collaborator"} {
		if diags := validate(value, cty.Path{cty.GetAttrStep{Name: "downgrade_to"}}); !diags.HasError() {
			t.Fatalf("expected %q to be invalid", value)
		}
	}
}
