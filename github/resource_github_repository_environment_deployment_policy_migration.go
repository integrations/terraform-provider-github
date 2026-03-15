package github

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryEnvironmentDeploymentPolicyV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"repository": {
				Description: "The name of the GitHub repository.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"environment": {
				Description: "The name of the environment.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"branch_pattern": {
				Description:      "The name pattern that branches must match in order to deploy to the environment.",
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         false,
				ExactlyOneOf:     []string{"branch_pattern", "tag_pattern"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
			"tag_pattern": {
				Description:      "The name pattern that tags must match in order to deploy to the environment.",
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         false,
				ExactlyOneOf:     []string{"branch_pattern", "tag_pattern"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
		},
	}
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyStateUpgradeV0(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	tflog.Debug(ctx, "Starting state upgrade for GitHub Repository Environment Deployment Policy.", map[string]any{"raw_state": rawState})

	_, _, policyIDStr, err := parseID3(rawState["id"].(string))
	if err != nil {
		return nil, err
	}

	repoName, ok := rawState["repository"].(string)
	if !ok {
		return nil, fmt.Errorf("repository not found or is not a string")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository %s: %w", repoName, err)
	}

	policyID, err := strconv.Atoi(policyIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid policy ID: %s", policyIDStr)
	}

	rawState["repository_id"] = int(repo.GetID())
	rawState["policy_id"] = policyID

	tflog.Debug(ctx, "Completed state upgrade for GitHub Repository Environment Deployment Policy.", map[string]any{"upgraded_state": rawState})

	return rawState, nil
}
