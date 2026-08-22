package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsVariableV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository.",
			},
			"variable_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Name of the variable.",
				ValidateDiagFunc: validateSecretNameFunc,
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Value of the variable.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'actions_variable' creation.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'actions_variable' update.",
			},
		},
	}
}

func resourceGithubActionsVariableStateUpgradeV0(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	tflog.Debug(ctx, "GitHub Actions Variable migration from v0 to v1 started.", map[string]any{"raw_state": rawState})

	repoName, ok := rawState["repository"].(string)
	if !ok {
		return nil, fmt.Errorf("repository not found or is not a string")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository %s: %w", repoName, err)
	}

	repoID := int(repo.GetID())
	rawState["repository_id"] = repoID

	tflog.Debug(ctx, "GitHub Actions Variable migration from v0 to v1 completed.", map[string]any{"raw_state": rawState})

	return rawState, nil
}
