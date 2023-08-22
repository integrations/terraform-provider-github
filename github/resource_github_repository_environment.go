package github

import (
	"context"
	"log"
	"net/http"
	"net/url"

	"github.com/google/go-github/v54/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceGithubRepositoryEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryEnvironmentCreate,
		Read:   resourceGithubRepositoryEnvironmentRead,
		Update: resourceGithubRepositoryEnvironmentUpdate,
		Delete: resourceGithubRepositoryEnvironmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
			"wait_timer": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(0, 43200),
				Description:  "Amount of time to delay a job after the job is initially triggered.",
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

func resourceGithubRepositoryEnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	escapedEnvName := url.QueryEscape(envName)
	updateData := createUpdateEnvironmentData(d, meta)

	ctx := context.Background()

	_, _, err := client.Repositories.CreateUpdateEnvironment(ctx, owner, repoName, escapedEnvName, &updateData)

	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(repoName, envName))

	return resourceGithubRepositoryEnvironmentRead(d, meta)
}

func resourceGithubRepositoryEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName, envName, err := parseTwoPartID(d.Id(), "repository", "environment")
	escapedEnvName := url.QueryEscape(envName)
	if err != nil {
		return err
	}

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	env, _, err := client.Repositories.GetEnvironment(ctx, owner, repoName, escapedEnvName)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing repository environment %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
	}

	d.Set("repository", repoName)
	d.Set("environment", envName)

	for _, pr := range env.ProtectionRules {
		switch *pr.Type {
		case "wait_timer":
			d.Set("wait_timer", pr.WaitTimer)

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
			d.Set("reviewers", []interface{}{
				map[string]interface{}{
					"teams": teams,
					"users": users,
				},
			})
		}
	}

	if env.DeploymentBranchPolicy != nil {
		d.Set("deployment_branch_policy", []interface{}{
			map[string]interface{}{
				"protected_branches":     env.DeploymentBranchPolicy.ProtectedBranches,
				"custom_branch_policies": env.DeploymentBranchPolicy.CustomBranchPolicies,
			},
		})
	}

	return nil
}

func resourceGithubRepositoryEnvironmentUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	escapedEnvName := url.QueryEscape(envName)
	updateData := createUpdateEnvironmentData(d, meta)

	ctx := context.Background()

	resultKey, _, err := client.Repositories.CreateUpdateEnvironment(ctx, owner, repoName, escapedEnvName, &updateData)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(repoName, resultKey.GetName()))

	return resourceGithubRepositoryEnvironmentRead(d, meta)
}

func resourceGithubRepositoryEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName, envName, err := parseTwoPartID(d.Id(), "repository", "environment")
	escapedEnvName := url.QueryEscape(envName)
	if err != nil {
		return err
	}

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	_, err = client.Repositories.DeleteEnvironment(ctx, owner, repoName, escapedEnvName)
	return err
}

func createUpdateEnvironmentData(d *schema.ResourceData, meta interface{}) github.CreateUpdateEnvironment {
	data := github.CreateUpdateEnvironment{}

	if v, ok := d.GetOk("wait_timer"); ok {
		data.WaitTimer = github.Int(v.(int))
	}

	if v, ok := d.GetOk("reviewers"); ok {
		envReviewers := make([]*github.EnvReviewers, 0)

		for _, team := range expandReviewers(v, "teams") {
			envReviewers = append(envReviewers, &github.EnvReviewers{
				Type: github.String("Team"),
				ID:   github.Int64(team),
			})
		}

		for _, user := range expandReviewers(v, "users") {
			envReviewers = append(envReviewers, &github.EnvReviewers{
				Type: github.String("User"),
				ID:   github.Int64(user),
			})
		}

		data.Reviewers = envReviewers
	}

	if v, ok := d.GetOk("deployment_branch_policy"); ok {
		policy := v.([]interface{})[0].(map[string]interface{})
		data.DeploymentBranchPolicy = &github.BranchPolicy{
			ProtectedBranches:    github.Bool(policy["protected_branches"].(bool)),
			CustomBranchPolicies: github.Bool(policy["custom_branch_policies"].(bool)),
		}
	}

	return data
}

func expandReviewers(v interface{}, target string) []int64 {
	res := make([]int64, 0)
	m := v.([]interface{})[0]
	if m != nil {
		if v, ok := m.(map[string]interface{})[target]; ok {
			vL := v.(*schema.Set).List()
			for _, v := range vL {
				res = append(res, int64(v.(int)))
			}
		}
	}
	return res
}
