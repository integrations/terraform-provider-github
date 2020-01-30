package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-github/v28/github"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceGithubBranchProtection() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubBranchProtectionCreate,
		Read:   resourceGithubBranchProtectionRead,
		Update: resourceGithubBranchProtectionUpdate,
		Delete: resourceGithubBranchProtectionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"branch": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"required_status_checks": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"include_admins": {
							Type:       schema.TypeBool,
							Optional:   true,
							Default:    false,
							Deprecated: "Use enforce_admins instead",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return true
							},
						},
						"strict": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"contexts": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"required_pull_request_reviews": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"include_admins": {
							Type:       schema.TypeBool,
							Optional:   true,
							Default:    false,
							Deprecated: "Use enforce_admins instead",
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								return true
							},
						},
						"dismiss_stale_reviews": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"dismissal_users": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"dismissal_teams": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"require_code_owner_reviews": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"required_approving_review_count": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      1,
							ValidateFunc: validation.IntBetween(1, 6),
						},
					},
				},
			},
			"restrictions": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"users": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"teams": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"enforce_admins": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"require_signed_commits": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubBranchProtectionCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client

	orgName := meta.(*Organization).name
	repoName := d.Get("repository").(string)
	branch := d.Get("branch").(string)

	protectionRequest, err := buildProtectionRequest(d)
	if err != nil {
		return err
	}
	ctx := context.Background()

	log.Printf("[DEBUG] Creating branch protection: %s/%s (%s)",
		orgName, repoName, branch)
	protection, _, err := client.Repositories.UpdateBranchProtection(ctx,
		orgName,
		repoName,
		branch,
		protectionRequest,
	)
	if err != nil {
		return err
	}

	if err := checkBranchRestrictionsUsers(protection.GetRestrictions(), protectionRequest.GetRestrictions()); err != nil {
		return err
	}

	d.SetId(buildTwoPartID(&repoName, &branch))

	if err = requireSignedCommitsUpdate(d, meta); err != nil {
		return err
	}

	return resourceGithubBranchProtectionRead(d, meta)
}

func resourceGithubBranchProtectionRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client

	repoName, branch, err := parseTwoPartID(d.Id(), "repository", "branch")
	if err != nil {
		return err
	}
	orgName := meta.(*Organization).name

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	log.Printf("[DEBUG] Reading branch protection: %s/%s (%s)",
		orgName, repoName, branch)
	githubProtection, resp, err := client.Repositories.GetBranchProtection(ctx,
		orgName, repoName, branch)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				if err := requireSignedCommitsRead(d, meta); err != nil {
					return fmt.Errorf("Error setting signed commit restriction: %v", err)
				}
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing branch protection %s/%s (%s) from state because it no longer exists in GitHub",
					orgName, repoName, branch)
				d.SetId("")
				return nil
			}
		}

		return err
	}

	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("repository", repoName)
	d.Set("branch", branch)
	d.Set("enforce_admins", githubProtection.EnforceAdmins.Enabled)

	if err := flattenAndSetRequiredStatusChecks(d, githubProtection); err != nil {
		return fmt.Errorf("Error setting required_status_checks: %v", err)
	}

	if err := flattenAndSetRequiredPullRequestReviews(d, githubProtection); err != nil {
		return fmt.Errorf("Error setting required_pull_request_reviews: %v", err)
	}

	if err := flattenAndSetRestrictions(d, githubProtection); err != nil {
		return fmt.Errorf("Error setting restrictions: %v", err)
	}

	if err := requireSignedCommitsRead(d, meta); err != nil {
		return fmt.Errorf("Error setting signed commit restriction: %v", err)
	}

	return nil
}

func resourceGithubBranchProtectionUpdate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client
	repoName, branch, err := parseTwoPartID(d.Id(), "repository", "branch")
	if err != nil {
		return err
	}

	protectionRequest, err := buildProtectionRequest(d)
	if err != nil {
		return err
	}

	orgName := meta.(*Organization).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Updating branch protection: %s/%s (%s)",
		orgName, repoName, branch)
	protection, _, err := client.Repositories.UpdateBranchProtection(ctx,
		orgName,
		repoName,
		branch,
		protectionRequest,
	)
	if err != nil {
		return err
	}

	if err := checkBranchRestrictionsUsers(protection.GetRestrictions(), protectionRequest.GetRestrictions()); err != nil {
		return err
	}

	if protectionRequest.RequiredPullRequestReviews == nil {
		_, err = client.Repositories.RemovePullRequestReviewEnforcement(ctx,
			orgName,
			repoName,
			branch,
		)
		if err != nil {
			return err
		}
	}

	d.SetId(buildTwoPartID(&repoName, &branch))

	if err = requireSignedCommitsUpdate(d, meta); err != nil {
		return err
	}

	return resourceGithubBranchProtectionRead(d, meta)
}

func resourceGithubBranchProtectionDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client
	repoName, branch, err := parseTwoPartID(d.Id(), "repository", "branch")
	if err != nil {
		return err
	}

	orgName := meta.(*Organization).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Deleting branch protection: %s/%s (%s)", orgName, repoName, branch)
	_, err = client.Repositories.RemoveBranchProtection(ctx,
		orgName, repoName, branch)
	return err
}

func buildProtectionRequest(d *schema.ResourceData) (*github.ProtectionRequest, error) {
	req := &github.ProtectionRequest{
		EnforceAdmins: d.Get("enforce_admins").(bool),
	}

	rsc, err := expandRequiredStatusChecks(d)
	if err != nil {
		return nil, err
	}
	req.RequiredStatusChecks = rsc

	rprr, err := expandRequiredPullRequestReviews(d)
	if err != nil {
		return nil, err
	}
	req.RequiredPullRequestReviews = rprr

	res, err := expandRestrictions(d)
	if err != nil {
		return nil, err
	}
	req.Restrictions = res

	return req, nil
}

func flattenAndSetRequiredStatusChecks(d *schema.ResourceData, protection *github.Protection) error {
	rsc := protection.RequiredStatusChecks
	if rsc != nil {
		contexts := make([]interface{}, 0, len(rsc.Contexts))
		for _, c := range rsc.Contexts {
			contexts = append(contexts, c)
		}

		return d.Set("required_status_checks", []interface{}{
			map[string]interface{}{
				"strict":   rsc.Strict,
				"contexts": schema.NewSet(schema.HashString, contexts),
			},
		})
	}

	return d.Set("required_status_checks", []interface{}{})
}

func requireSignedCommitsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).v3client

	repoName, branch, err := parseTwoPartID(d.Id(), "repository", "branch")
	if err != nil {
		return err
	}
	orgName := meta.(*Organization).name

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	log.Printf("[DEBUG] Reading branch protection signed commit status: %s/%s (%s)", orgName, repoName, branch)
	signedCommitStatus, _, err := client.Repositories.GetSignaturesProtectedBranch(ctx,
		orgName, repoName, branch)
	if err != nil {
		log.Printf("[WARN] Not able to read signature protection: %s/%s (%s)", orgName, repoName, branch)
		return nil
	}

	return d.Set("require_signed_commits", signedCommitStatus.Enabled)
}

func requireSignedCommitsUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	requiredSignedCommit := d.Get("require_signed_commits").(bool)
	client := meta.(*Organization).v3client

	repoName, branch, err := parseTwoPartID(d.Id(), "repository", "branch")
	if err != nil {
		return err
	}
	orgName := meta.(*Organization).name

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	if requiredSignedCommit {
		log.Printf("[DEBUG] Enabling branch protection signed commit: %s/%s (%s) - $s", orgName, repoName, branch)
		_, _, err = client.Repositories.RequireSignaturesOnProtectedBranch(ctx, orgName, repoName, branch)
	} else {
		log.Printf("[DEBUG] Removing branch protection signed commit: %s/%s (%s) - $s", orgName, repoName, branch)
		_, err = client.Repositories.OptionalSignaturesOnProtectedBranch(ctx, orgName, repoName, branch)
	}
	return err
}

func flattenAndSetRequiredPullRequestReviews(d *schema.ResourceData, protection *github.Protection) error {
	rprr := protection.RequiredPullRequestReviews
	if rprr != nil {
		users := make([]interface{}, 0, len(rprr.DismissalRestrictions.Users))
		for _, u := range rprr.DismissalRestrictions.Users {
			if u.Login != nil {
				users = append(users, *u.Login)
			}
		}

		teams := make([]interface{}, 0, len(rprr.DismissalRestrictions.Teams))
		for _, t := range rprr.DismissalRestrictions.Teams {
			if t.Slug != nil {
				teams = append(teams, *t.Slug)
			}
		}

		return d.Set("required_pull_request_reviews", []interface{}{
			map[string]interface{}{
				"dismiss_stale_reviews":           rprr.DismissStaleReviews,
				"dismissal_users":                 schema.NewSet(schema.HashString, users),
				"dismissal_teams":                 schema.NewSet(schema.HashString, teams),
				"require_code_owner_reviews":      rprr.RequireCodeOwnerReviews,
				"required_approving_review_count": rprr.RequiredApprovingReviewCount,
			},
		})
	}

	return d.Set("required_pull_request_reviews", []interface{}{})
}

func flattenAndSetRestrictions(d *schema.ResourceData, protection *github.Protection) error {
	restrictions := protection.Restrictions
	if restrictions != nil {
		users := make([]interface{}, 0, len(restrictions.Users))
		for _, u := range restrictions.Users {
			if u.Login != nil {
				users = append(users, *u.Login)
			}
		}

		teams := make([]interface{}, 0, len(restrictions.Teams))
		for _, t := range restrictions.Teams {
			if t.Slug != nil {
				teams = append(teams, *t.Slug)
			}
		}

		return d.Set("restrictions", []interface{}{
			map[string]interface{}{
				"users": schema.NewSet(schema.HashString, users),
				"teams": schema.NewSet(schema.HashString, teams),
			},
		})
	}

	return d.Set("restrictions", []interface{}{})
}

func expandRequiredStatusChecks(d *schema.ResourceData) (*github.RequiredStatusChecks, error) {
	if v, ok := d.GetOk("required_status_checks"); ok {
		vL := v.([]interface{})
		if len(vL) > 1 {
			return nil, errors.New("cannot specify required_status_checks more than one time")
		}
		rsc := new(github.RequiredStatusChecks)

		for _, v := range vL {
			// List can only have one item, safe to early return here
			if v == nil {
				return nil, nil
			}
			m := v.(map[string]interface{})
			rsc.Strict = m["strict"].(bool)

			contexts := expandNestedSet(m, "contexts")
			rsc.Contexts = contexts
		}
		return rsc, nil
	}

	return nil, nil
}

func expandRequiredPullRequestReviews(d *schema.ResourceData) (*github.PullRequestReviewsEnforcementRequest, error) {
	if v, ok := d.GetOk("required_pull_request_reviews"); ok {
		vL := v.([]interface{})
		if len(vL) > 1 {
			return nil, errors.New("cannot specify required_pull_request_reviews more than one time")
		}

		rprr := new(github.PullRequestReviewsEnforcementRequest)
		drr := new(github.DismissalRestrictionsRequest)

		for _, v := range vL {
			// List can only have one item, safe to early return here
			if v == nil {
				return nil, nil
			}
			m := v.(map[string]interface{})

			users := expandNestedSet(m, "dismissal_users")
			drr.Users = &users
			teams := expandNestedSet(m, "dismissal_teams")
			drr.Teams = &teams

			rprr.DismissalRestrictionsRequest = drr
			rprr.DismissStaleReviews = m["dismiss_stale_reviews"].(bool)
			rprr.RequireCodeOwnerReviews = m["require_code_owner_reviews"].(bool)
			rprr.RequiredApprovingReviewCount = m["required_approving_review_count"].(int)
		}

		return rprr, nil
	}

	return nil, nil
}

func expandRestrictions(d *schema.ResourceData) (*github.BranchRestrictionsRequest, error) {
	if v, ok := d.GetOk("restrictions"); ok {
		vL := v.([]interface{})
		if len(vL) > 1 {
			return nil, errors.New("cannot specify restrictions more than one time")
		}
		restrictions := new(github.BranchRestrictionsRequest)

		for _, v := range vL {
			// Restrictions only have set attributes nested, need to return nil values for these.
			// The API won't initialize these as nil
			if v == nil {
				restrictions.Users = []string{}
				restrictions.Teams = []string{}
				return restrictions, nil
			}
			m := v.(map[string]interface{})

			users := expandNestedSet(m, "users")
			restrictions.Users = users
			teams := expandNestedSet(m, "teams")
			restrictions.Teams = teams
		}
		return restrictions, nil
	}

	return nil, nil
}

func expandNestedSet(m map[string]interface{}, target string) []string {
	res := []string{}
	if v, ok := m[target]; ok {
		vL := v.(*schema.Set).List()
		for _, v := range vL {
			res = append(res, v.(string))
		}
	}
	return res
}

func checkBranchRestrictionsUsers(actual *github.BranchRestrictions, expected *github.BranchRestrictionsRequest) error {
	if expected == nil {
		return nil
	}

	expectedUsers := expected.Users

	if actual == nil {
		return fmt.Errorf("unable to add users in restrictions: %s", strings.Join(expectedUsers, ", "))
	}

	actualLoopUp := make(map[string]struct{}, len(actual.Users))
	for _, a := range actual.Users {
		actualLoopUp[a.GetLogin()] = struct{}{}
	}

	notFounds := make([]string, 0, len(actual.Users))

	for _, e := range expectedUsers {
		if _, ok := actualLoopUp[e]; !ok {
			notFounds = append(notFounds, e)
		}
	}

	if len(notFounds) == 0 {
		return nil
	}

	return fmt.Errorf("unable to add users in restrictions: %s", strings.Join(notFounds, ", "))
}
