package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryEnvironment() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubRepositoryEnvironmentV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubRepositoryEnvironmentStateUpgradeV0,
				Version: 0,
			},
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The repository of the environment.",
			},
			"repository_id": {
				Description: "The ID of the GitHub repository.",
				Type:        schema.TypeInt,
				Computed:    true,
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
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(0, 43200)),
				Description:      "Amount of time to delay a job after the job is initially triggered.",
			},
			"reviewers": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The environment reviewers configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"teams": {
							Type:        schema.TypeSet,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Optional:    true,
							MaxItems:    6,
							Description: "Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.",
						},
						"users": {
							Type:        schema.TypeSet,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Optional:    true,
							MaxItems:    6,
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

		CustomizeDiff: customdiff.All(
			diffRepository,
			resourceGithubRepositoryEnvironmentDiff,
		),

		CreateContext: resourceGithubRepositoryEnvironmentCreate,
		ReadContext:   resourceGithubRepositoryEnvironmentRead,
		UpdateContext: resourceGithubRepositoryEnvironmentUpdate,
		DeleteContext: resourceGithubRepositoryEnvironmentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryEnvironmentImport,
		},
	}
}

func resourceGithubRepositoryEnvironmentDiff(_ context.Context, d *schema.ResourceDiff, _ any) error {
	if d.Id() == "" {
		return nil
	}

	if v, ok := d.GetOk("reviewers"); ok {
		count := 0
		o := v.([]any)[0]
		if t, ok := o.(map[string]any)["teams"]; ok {
			count += t.(*schema.Set).Len()
		}

		if t, ok := o.(map[string]any)["users"]; ok {
			count += t.(*schema.Set).Len()
		}

		if count > 6 {
			return fmt.Errorf("a maximum of 6 reviewers (users and teams combined) can be set for an environment")
		}
	}

	return nil
}

func resourceGithubRepositoryEnvironmentCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	updateData := createUpdateEnvironmentData(d)

	_, _, err := client.Repositories.CreateUpdateEnvironment(ctx, owner, repoName, url.PathEscape(envName), &updateData)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, escapeIDPart(envName))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryEnvironmentRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "id", d.Id())

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)

	env, _, err := client.Repositories.GetEnvironment(ctx, owner, repoName, url.PathEscape(envName))
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Repository environment not found, removing from state.", map[string]any{"repository": repoName, "environment": envName})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if err := d.Set("wait_timer", nil); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("can_admins_bypass", env.CanAdminsBypass); err != nil {
		return diag.FromErr(err)
	}

	for _, pr := range env.ProtectionRules {
		switch pr.GetType() {
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
		if err := d.Set("deployment_branch_policy", []any{}); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubRepositoryEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	updateData := createUpdateEnvironmentData(d)

	_, _, err := client.Repositories.CreateUpdateEnvironment(ctx, owner, repoName, url.PathEscape(envName), &updateData)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, escapeIDPart(envName))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	return nil
}

func resourceGithubRepositoryEnvironmentDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)

	_, err := client.Repositories.DeleteEnvironment(ctx, owner, repoName, url.PathEscape(envName))
	if err != nil {
		return diag.FromErr(deleteResourceOn404AndSwallow304OtherwiseReturnError(err, d, "environment (%s)", envName))
	}

	return nil
}

func resourceGithubRepositoryEnvironmentImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, envNamePart, err := parseID2(d.Id())
	if err != nil {
		return nil, fmt.Errorf("invalid id (%s), expected format <repository>:<environment>", d.Id())
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository %s: %w", repoName, err)
	}

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
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
