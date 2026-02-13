package github

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsSecretV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the repository.",
			},
			"secret_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Name of the secret.",
				ValidateDiagFunc: validateSecretNameFunc,
			},
			"encrypted_value": {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"plaintext_value"},
				Description:   "Encrypted value of the secret using the GitHub public key in Base64 format.",
			},
			"plaintext_value": {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"encrypted_value"},
				Description:   "Plaintext value of the secret to be encrypted.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'actions_secret' creation.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'actions_secret' update.",
			},
		},
	}
}

func resourceGithubActionsSecretV1() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the repository.",
			},
			"secret_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Name of the secret.",
				ValidateDiagFunc: validateSecretNameFunc,
			},
			"encrypted_value": {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"plaintext_value"},
				Description:   "Encrypted value of the secret using the GitHub public key in Base64 format.",
			},
			"plaintext_value": {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"encrypted_value"},
				Description:   "Plaintext value of the secret to be encrypted.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'actions_secret' creation.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'actions_secret' update.",
			},
			"destroy_on_drift": {
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
				ForceNew:    true,
				Description: "Boolean indicating whether to recreate the secret if it's modified outside of Terraform. When `true` (default), Terraform will delete and recreate the secret if it detects external changes. When `false`, Terraform will acknowledge external changes but not recreate the secret.",
			},
		},
	}
}

func resourceGithubActionsSecretStateUpgradeV0(ctx context.Context, rawState map[string]any, _ any) (map[string]any, error) {
	log.Printf("[DEBUG] GitHub Actions Secret State before migration: %#v", rawState)

	// Add the destroy_on_drift field with default value true if it doesn't exist
	if _, ok := rawState["destroy_on_drift"]; !ok {
		rawState["destroy_on_drift"] = true
	}

	log.Printf("[DEBUG] GitHub Actions Secret State after migration: %#v", rawState)

	return rawState, nil
}

func resourceGithubActionsSecretStateUpgradeV1(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	log.Printf("[DEBUG] GitHub Actions Secret Attributes before migration: %#v", rawState)

	repoName, ok := rawState["repository"].(string)
	if !ok {
		return nil, fmt.Errorf("repository not found or is not a string")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository %s: %w", repoName, err)
	}

	rawState["repository_id"] = int(repo.GetID())

	log.Printf("[DEBUG] GitHub Actions Secret Attributes after migration: %#v", rawState)

	return rawState, nil
}
