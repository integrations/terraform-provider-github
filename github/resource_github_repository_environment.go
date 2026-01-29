package github

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryEnvironment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubRepositoryEnvironmentCreate,
		ReadContext:   resourceGithubRepositoryEnvironmentRead,
		UpdateContext: resourceGithubRepositoryEnvironmentUpdate,
		DeleteContext: resourceGithubRepositoryEnvironmentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryEnvironmentImport,
		},
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The repository of the environment.",
			},
			"environment": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the environment.",
			},
			"can_admins_bypass": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Can Admins bypass deployment protections",
			},
			"prevent_self_review": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Prevent users from approving workflows runs that they triggered.",
			},
			"wait_timer": {
				Type:             schema.TypeInt,
				Optional:         true,
				ValidateDiagFunc: toDiagFunc(validation.IntBetween(0, 43200), "wait_timer"),
				Description:      "Amount of time to delay a job after the job is initially triggered.",
			},
			"reviewers": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    6,
				Description: "The environment reviewers configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"teams": {
							Type:        schema.TypeSet,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Description: "Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.",
						},
						"users": {
							Type:        schema.TypeSet,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Description: "Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.",
						},
					},
				},
			},
			"deployment_branch_policy": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The deployment branch policy configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protected_branches": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether only branches with branch protection rules can deploy to this environment.",
						},
						"custom_branch_policies": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether only branches that match the specified name patterns can deploy to this environment.",
						},
					},
				},
			},
		},
	}
}

func resourceGithubRepositoryEnvironmentCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	updateData := createUpdateEnvironmentData(d)

	_, _, err := client.Repositories.CreateUpdateEnvironment(ctx, owner, repoName, url.PathEscape(envName), &updateData)
	if err != nil {
		return diag.FromErr(err)
	}

	if id, err := buildID(repoName, escapeIDPart(envName)); err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId(id)
	}

	return nil
}

func resourceGithubRepositoryEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repoName, envNamePart, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	envName := unescapeIDPart(envNamePart)

	env, _, err := client.Repositories.GetEnvironment(ctx, owner, repoName, url.PathEscape(envName))
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing repository environment %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	_ = d.Set("repository", repoName)
	_ = d.Set("environment", envName)
	_ = d.Set("wait_timer", nil)
	_ = d.Set("can_admins_bypass", env.CanAdminsBypass)

	for _, pr := range env.ProtectionRules {
		switch *pr.Type {
		case "wait_timer":
			if err = d.Set("wait_timer", pr.WaitTimer); err != nil {
				return diag.FromErr(err)
			}

		case "required_reviewers":
			teams := make([]int64, 0)
			users := make([]int64, 0)

			for _, r := range pr.Reviewers {
				switch *r.Type {
				case "Team":
					if r.Reviewer.(*github.Team).ID != nil {
						teams = append(teams, *r.Reviewer.(*github.Team).ID)
					}
				case "User":
					if r.Reviewer.(*github.User).ID != nil {
						users = append(users, *r.Reviewer.(*github.User).ID)
					}
				}
			}
			if err = d.Set("reviewers", []any{
				map[string]any{
					"teams": teams,
					"users": users,
				},
			}); err != nil {
				return diag.FromErr(err)
			}

			if err = d.Set("prevent_self_review", pr.PreventSelfReview); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	if env.DeploymentBranchPolicy != nil {
		if err = d.Set("deployment_branch_policy", []any{
			map[string]any{
				"protected_branches":     env.DeploymentBranchPolicy.ProtectedBranches,
				"custom_branch_policies": env.DeploymentBranchPolicy.CustomBranchPolicies,
			},
		}); err != nil {
			return diag.FromErr(err)
		}
	} else {
		_ = d.Set("deployment_branch_policy", []any{})
	}

	return nil
}

func resourceGithubRepositoryEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	updateData := createUpdateEnvironmentData(d)

	_, _, err := client.Repositories.CreateUpdateEnvironment(ctx, owner, repoName, url.PathEscape(envName), &updateData)
	if err != nil {
		return diag.FromErr(err)
	}

	if id, err := buildID(repoName, escapeIDPart(envName)); err != nil {
		return diag.FromErr(err)
	} else {
		d.SetId(id)
	}

	return nil
}

func resourceGithubRepositoryEnvironmentDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repoName, envNamePart, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	envName := unescapeIDPart(envNamePart)

	_, err = client.Repositories.DeleteEnvironment(ctx, owner, repoName, url.PathEscape(envName))
	if err != nil {
		return diag.FromErr(deleteResourceOn404AndSwallow304OtherwiseReturnError(err, d, "environment (%s)", envName))
	}

	return nil
}

func resourceGithubRepositoryEnvironmentImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	repoName, envNamePart, err := parseID2(d.Id())
	if err != nil {
		return nil, err
	}

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err := d.Set("environment", unescapeIDPart(envNamePart)); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func createUpdateEnvironmentData(d *schema.ResourceData) github.CreateUpdateEnvironment {
	data := github.CreateUpdateEnvironment{}

	if v, ok := d.GetOk("wait_timer"); ok {
		data.WaitTimer = github.Ptr(v.(int))
	}

	data.CanAdminsBypass = github.Ptr(d.Get("can_admins_bypass").(bool))

	data.PreventSelfReview = github.Ptr(d.Get("prevent_self_review").(bool))

	if v, ok := d.GetOk("reviewers"); ok {
		envReviewers := make([]*github.EnvReviewers, 0)

		for _, team := range expandReviewers(v, "teams") {
			envReviewers = append(envReviewers, &github.EnvReviewers{
				Type: github.Ptr("Team"),
				ID:   github.Ptr(team),
			})
		}

		for _, user := range expandReviewers(v, "users") {
			envReviewers = append(envReviewers, &github.EnvReviewers{
				Type: github.Ptr("User"),
				ID:   github.Ptr(user),
			})
		}

		data.Reviewers = envReviewers
	}

	if v, ok := d.GetOk("deployment_branch_policy"); ok {
		policy := v.([]any)[0].(map[string]any)
		data.DeploymentBranchPolicy = &github.BranchPolicy{
			ProtectedBranches:    github.Ptr(policy["protected_branches"].(bool)),
			CustomBranchPolicies: github.Ptr(policy["custom_branch_policies"].(bool)),
		}
	}

	return data
}

func expandReviewers(v any, target string) []int64 {
	res := make([]int64, 0)
	m := v.([]any)[0]
	if m != nil {
		if v, ok := m.(map[string]any)[target]; ok {
			vL := v.(*schema.Set).List()
			for _, v := range vL {
				res = append(res, int64(v.(int)))
			}
		}
	}
	return res
}
