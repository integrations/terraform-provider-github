package github

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsEnvironmentVariableV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository.",
			},
			"environment": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the environment.",
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

func resourceGithubActionsEnvironmentVariableStateUpgradeV0(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	log.Printf("[DEBUG] GitHub Actions Environment Variable Attributes before migration: %#v", rawState)

	repoName, ok := rawState["repository"].(string)
	if !ok {
		return nil, fmt.Errorf("repository not found or is not a string")
	}

	envName, ok := rawState["environment"].(string)
	if !ok {
		return nil, fmt.Errorf("environment not found or is not a string")
	}

	varName, ok := rawState["variable_name"].(string)
	if !ok {
		return nil, fmt.Errorf("variable_name not found or is not a string")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository %s: %w", repoName, err)
	}

	repoID := int(repo.GetID())

	id, err := buildID(repoName, escapeIDPart(envName), varName)
	if err != nil {
		return nil, fmt.Errorf("failed to build id for repository %s, environment %s, variable %s: %w", repoName, envName, varName, err)
	}
	rawState["id"] = id
	rawState["repository_id"] = repoID

	log.Printf("[DEBUG] GitHub Actions Environment Variable Attributes after migration: %#v", rawState)

	return rawState, nil
}
