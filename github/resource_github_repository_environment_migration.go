package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryEnvironmentV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
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
	}
}

func resourceGithubRepositoryEnvironmentStateUpgradeV0(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	tflog.Debug(ctx, "Starting state upgrade for GitHub Repository Environment.", map[string]any{"raw_state": rawState})

	repoName, ok := rawState["repository"].(string)
	if !ok {
		return nil, fmt.Errorf("repository not found or is not a string")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository %s: %w", repoName, err)
	}

	rawState["repository_id"] = int(repo.GetID())

	tflog.Debug(ctx, "Completed state upgrade for GitHub Repository Environment.", map[string]any{"upgraded_state": rawState})

	return rawState, nil
}
