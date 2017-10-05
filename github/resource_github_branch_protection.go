package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/schema"
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
		},
	}
}

func resourceGithubBranchProtectionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	r := d.Get("repository").(string)
	b := d.Get("branch").(string)

	protectionRequest, err := buildProtectionRequest(d)
	if err != nil {
		return err
	}

	_, _, err = client.Repositories.UpdateBranchProtection(context.TODO(), meta.(*Organization).name, r, b, protectionRequest)
	if err != nil {
		return err
	}
	d.SetId(buildTwoPartID(&r, &b))

	return resourceGithubBranchProtectionRead(d, meta)
}

func resourceGithubBranchProtectionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	r, b := parseTwoPartID(d.Id())

	githubProtection, _, err := client.Repositories.GetBranchProtection(context.TODO(), meta.(*Organization).name, r, b)
	if err != nil {
		if err, ok := err.(*github.ErrorResponse); ok && err.Response.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}

		return err
	}

	d.Set("repository", r)
	d.Set("branch", b)
	d.Set("enforce_admins", githubProtection.EnforceAdmins.Enabled)

	if err := flattenRequiredStatusChecks(d, githubProtection); err != nil {
		return fmt.Errorf("Error setting required_status_checks: %v", err)
	}

	if err := flattenRequiredPullRequestReviews(d, githubProtection); err != nil {
		return fmt.Errorf("Error setting required_pull_request_reviews: %v", err)
	}

	if err := flattenRestrictions(d, githubProtection); err != nil {
		return fmt.Errorf("Error setting restrictions: %v", err)
	}

	return nil
}

func resourceGithubBranchProtectionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	r, b := parseTwoPartID(d.Id())

	protectionRequest, err := buildProtectionRequest(d)
	if err != nil {
		return err
	}

	_, _, err = client.Repositories.UpdateBranchProtection(context.TODO(), meta.(*Organization).name, r, b, protectionRequest)
	if err != nil {
		return err
	}

	if protectionRequest.RequiredPullRequestReviews == nil {
		_, err = client.Repositories.RemovePullRequestReviewEnforcement(context.TODO(), meta.(*Organization).name, r, b)
		if err != nil {
			return err
		}
	}

	d.SetId(buildTwoPartID(&r, &b))

	return resourceGithubBranchProtectionRead(d, meta)
}

func resourceGithubBranchProtectionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	r, b := parseTwoPartID(d.Id())

	_, err := client.Repositories.RemoveBranchProtection(context.TODO(), meta.(*Organization).name, r, b)
	return err
}

func buildProtectionRequest(d *schema.ResourceData) (*github.ProtectionRequest, error) {
	protectionRequest := new(github.ProtectionRequest)

	rsc, err := expandRequiredStatusChecks(d)
	if err != nil {
		return nil, err
	}
	protectionRequest.RequiredStatusChecks = rsc

	rprr, err := expandRequiredPullRequestReviews(d)
	if err != nil {
		return nil, err
	}
	protectionRequest.RequiredPullRequestReviews = rprr

	res, err := expandRestrictions(d)
	if err != nil {
		return nil, err
	}
	protectionRequest.Restrictions = res

	protectionRequest.EnforceAdmins = d.Get("enforce_admins").(bool)

	return protectionRequest, nil
}

func flattenRequiredStatusChecks(d *schema.ResourceData, protection *github.Protection) error {
	rsc := protection.RequiredStatusChecks
	if rsc != nil {
		contexts := make([]interface{}, 0, len(rsc.Contexts))
		for _, c := range rsc.Contexts {
			contexts = append(contexts, c)
		}

		if err := d.Set("required_status_checks", []interface{}{
			map[string]interface{}{
				"strict":   rsc.Strict,
				"contexts": schema.NewSet(schema.HashString, contexts),
			},
		}); err != nil {
			return err
		}
	} else {
		d.Set("required_status_checks", []interface{}{})
	}

	return nil
}

func flattenRequiredPullRequestReviews(d *schema.ResourceData, protection *github.Protection) error {
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

		if err := d.Set("required_pull_request_reviews", []interface{}{
			map[string]interface{}{
				"dismiss_stale_reviews":      rprr.DismissStaleReviews,
				"dismissal_users":            schema.NewSet(schema.HashString, users),
				"dismissal_teams":            schema.NewSet(schema.HashString, teams),
				"require_code_owner_reviews": rprr.RequireCodeOwnerReviews,
			},
		}); err != nil {
			return err
		}
	} else {
		d.Set("required_pull_request_reviews", []interface{}{})
	}

	return nil
}

func flattenRestrictions(d *schema.ResourceData, protection *github.Protection) error {
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

		if err := d.Set("restrictions", []interface{}{
			map[string]interface{}{
				"users": schema.NewSet(schema.HashString, users),
				"teams": schema.NewSet(schema.HashString, teams),
			},
		}); err != nil {
			return fmt.Errorf("Error setting restrictions: %v", err)
		}
	} else {
		d.Set("restrictions", []interface{}{})
	}

	return nil
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
			drr.Users = users
			teams := expandNestedSet(m, "dismissal_teams")
			drr.Teams = teams

			rprr.DismissalRestrictionsRequest = drr
			rprr.DismissStaleReviews = m["dismiss_stale_reviews"].(bool)
			rprr.RequireCodeOwnerReviews = m["require_code_owner_reviews"].(bool)
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
