package github

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsEnvironmentSecretV0() *schema.Resource {
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
			"secret_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Name of the secret.",
				ValidateDiagFunc: validateSecretNameFunc,
			},
			"encrypted_value": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Description:      "Encrypted value of the secret using the GitHub public key in Base64 format.",
				ConflictsWith:    []string{"plaintext_value"},
				ValidateDiagFunc: toDiagFunc(validation.StringIsBase64, "encrypted_value"),
			},
			"plaintext_value": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				Description:   "Plaintext value of the secret to be encrypted.",
				ConflictsWith: []string{"encrypted_value"},
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'actions_environment_secret' creation.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'actions_environment_secret' update.",
			},
		},
	}
}

func resourceGithubActionsEnvironmentSecretStateUpgradeV0(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	log.Printf("[DEBUG] GitHub Actions Environment Secret Attributes before migration: %#v", rawState)

	repoName, ok := rawState["repository"].(string)
	if !ok {
		return nil, fmt.Errorf("repository not found or is not a string")
	}

	envName, ok := rawState["environment"].(string)
	if !ok {
		return nil, fmt.Errorf("environment not found or is not a string")
	}

	secretName, ok := rawState["secret_name"].(string)
	if !ok {
		return nil, fmt.Errorf("secret_name not found or is not a string")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository %s: %w", repoName, err)
	}

	repoID := int(repo.GetID())

	id, err := buildID(repoName, escapeIDPart(envName), secretName)
	if err != nil {
		return nil, fmt.Errorf("failed to build id for repository %s, environment %s, secret %s: %w", repoName, envName, secretName, err)
	}
	rawState["id"] = id
	rawState["repository_id"] = repoID

	log.Printf("[DEBUG] GitHub Actions Environment Secret Attributes after migration: %#v", rawState)

	return rawState, nil
}
