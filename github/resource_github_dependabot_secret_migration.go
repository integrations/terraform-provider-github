package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubDependabotSecretV0() *schema.Resource {
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
				Description: "Date of 'dependabot_secret' creation.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of 'dependabot_secret' update.",
			},
		},
	}
}

func resourceGithubDependabotSecretStateUpgradeV0(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	tflog.Debug(ctx, "GitHub Dependabot Secret migration from v0 to v1 started.", map[string]any{"raw_state": rawState})

	repoName, ok := rawState["repository"].(string)
	if !ok {
		return nil, fmt.Errorf("repository not found or is not a string")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository %s: %w", repoName, err)
	}

	rawState["repository_id"] = int(repo.GetID())

	tflog.Debug(ctx, "GitHub Dependabot Secret migration from v0 to v1 completed.", map[string]any{"raw_state": rawState})

	return rawState, nil
}
