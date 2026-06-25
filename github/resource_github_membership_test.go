package github

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccGithubMembership(t *testing.T) {
	if len(testAccConf.testExternalUser1) == 0 {
		t.Skip("No external user provided")
	}

	t.Run("creates organization membership", func(t *testing.T) {
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
				},
			},
		})
	})

	t.Run("creates organization membership with downgrade", func(t *testing.T) {
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

	t.Run("creates organization membership with case insensitivity", func(t *testing.T) {
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
				},
			},
		})
	})

	t.Run("creates organization membership by user_id", func(t *testing.T) {
		ctx := t.Context()

		meta, err := getTestMeta()
		if err != nil {
			t.Fatalf("failed to get test meta: %s", err)
		}

		ghUser, _, err := meta.v3client.Users.Get(ctx, testAccConf.testExternalUser)
		if err != nil {
			t.Fatalf("failed to resolve external user id: %s", err)
		}

		var membership github.Membership
		rn := "github_membership.test_org_membership"

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubMembershipDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccGithubMembershipConfigByUserID(ghUser.GetID()),
					Check: resource.ComposeTestCheckFunc(
						testAccCheckGithubMembershipExists(ctx, rn, &membership),
						testAccCheckGithubMembershipRoleState(ctx, rn, &membership),
						resource.TestCheckResourceAttr(rn, "username", testAccConf.testExternalUser),
						resource.TestCheckResourceAttr(rn, "user_id", fmt.Sprintf("%d", ghUser.GetID())),
					),
				},
				{
					ResourceName:      rn,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("errors when neither username nor user_id is provided", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: `
						resource "github_membership" "test_org_membership" {
							role = "member"
						}
					`,
					ExpectError: regexp.MustCompile(`one of (\x60username\x60,\x60user_id\x60|\x60user_id\x60,\x60username\x60) must be specified`),
				},
			},
		})
	})

	t.Run("errors when both username and user_id are provided", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						resource "github_membership" "test_org_membership" {
							username = "%s"
							user_id  = 1
							role     = "member"
						}
					`, testAccConf.testExternalUser),
					ExpectError: regexp.MustCompile(`only one of (\x60user_id\x60,\x60username\x60|\x60username\x60,\x60user_id\x60) can be specified`),
				},
			},
		})
	})
}

// TestAccGithubMembershipRenameResilience verifies that when a membership is
// created via user_id, a subsequent rename of the GitHub account does not
// produce drift: the resource Reads the current login by numeric id and
// silently updates the username attribute. Requires GH_TEST_EXTERNAL_USER_TOKEN
// since renaming the external user account requires that user's own PAT.
func TestAccGithubMembershipRenameResilience(t *testing.T) {
	if len(testAccConf.testExternalUser) == 0 {
		t.Skip("No external user provided")
	}
	if len(testAccConf.testExternalUserToken) == 0 {
		t.Skip("No external user token provided (GH_TEST_EXTERNAL_USER_TOKEN) - skipping live-rename test")
	}

	ctx := t.Context()

	meta, err := getTestMeta()
	if err != nil {
		t.Fatalf("failed to get test meta: %s", err)
	}

	ghUser, _, err := meta.v3client.Users.Get(ctx, testAccConf.testExternalUser)
	if err != nil {
		t.Fatalf("failed to resolve external user id: %s", err)
	}

	originalLogin := testAccConf.testExternalUser
	renamedLogin := originalLogin + "-renamed"

	rn := "github_membership.test_org_membership"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessHasOrgs(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckGithubMembershipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubMembershipConfigByUserID(ghUser.GetID()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "username", originalLogin),
					resource.TestCheckResourceAttr(rn, "user_id", fmt.Sprintf("%d", ghUser.GetID())),
				),
			},
			{
				PreConfig: func() {
					if err := renameExternalUser(ctx, renamedLogin); err != nil {
						t.Fatalf("failed to rename external user before refresh step: %s", err)
					}
					t.Cleanup(func() {
						if err := renameExternalUser(context.Background(), originalLogin); err != nil {
							t.Logf("WARNING: failed to restore external user login back to %q: %s", originalLogin, err)
						}
					})
				},
				Config:             testAccGithubMembershipConfigByUserID(ghUser.GetID()),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				RefreshState: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "username", renamedLogin),
					resource.TestCheckResourceAttr(rn, "user_id", fmt.Sprintf("%d", ghUser.GetID())),
				),
			},
		},
	})
}

func testAccGithubMembershipConfigByUserID(userID int64) string {
	return fmt.Sprintf(`
  resource "github_membership" "test_org_membership" {
    user_id = %d
    role    = "member"
  }
`, userID)
}

// renameExternalUser renames the GH_TEST_EXTERNAL_USER account using that
// user's own PAT (GH_TEST_EXTERNAL_USER_TOKEN). PATCH /user only works for the
// authenticated user, so the org owner's token cannot rename them.
func renameExternalUser(ctx context.Context, newLogin string) error {
	cfg := Config{
		Token:   testAccConf.testExternalUserToken,
		BaseURL: testAccConf.baseURL,
	}
	owner, err := cfg.Meta()
	if err != nil {
		return fmt.Errorf("failed to build client for external user rename: %w", err)
	}
	client := owner.(*Owner).v3client

	_, _, err = client.Users.Edit(ctx, &github.User{Login: github.Ptr(newLogin)})
	return err
}

func testAccCheckGithubMembershipDestroy(s *terraform.State) error {
	ctx := context.Background()
	meta, err := getTestMeta()
	if err != nil {
		return err
	}
	conn := meta.v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_membership" {
			continue
		}

		orgName, username, err := parseID2(rs.Primary.ID)
		if err != nil {
			return err
		}

		downgradedOnDestroy := rs.Primary.Attributes["downgrade_on_destroy"] == "true"
		membership, resp, err := conn.Organizations.GetOrgMembership(ctx, username, orgName)
		responseIsSuccessful := err == nil && membership != nil && buildTwoPartID(orgName, username) == rs.Primary.ID

		if downgradedOnDestroy {
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

func testAccCheckGithubMembershipExists(ctx context.Context, n string, membership *github.Membership) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no membership ID is set")
		}

		meta, err := getTestMeta()
		if err != nil {
			return err
		}
		conn := meta.v3client

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

		meta, err := getTestMeta()
		if err != nil {
			return err
		}
		conn := meta.v3client

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

func testAccGithubMembershipTheSame(orig, other *github.Membership) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if orig.GetURL() != other.GetURL() {
			return errors.New("users are different")
		}

		return nil
	}
}
