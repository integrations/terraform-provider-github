package github

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsOrganizationSecretV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Name of the secret.",
				ValidateDiagFunc: validateSecretNameFunc,
			},
			"encrypted_value": {
				Type:             schema.TypeString,
				ForceNew:         true,
				Optional:         true,
				Sensitive:        true,
				ConflictsWith:    []string{"plaintext_value"},
				Description:      "Encrypted value of the secret using the GitHub public key in Base64 format.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsBase64),
			},
			"plaintext_value": {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"encrypted_value"},
				Description:   "Plaintext value of the secret to be encrypted.",
			},
			"visibility": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validateValueFunc([]string{"all", "private", "selected"}),
				Description:      "Configures the access that repositories have to the organization secret. Must be one of 'all', 'private', or 'selected'. 'selected_repository_ids' is required if set to 'selected'.",
			},
			"selected_repository_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Set:         schema.HashInt,
				Optional:    true,
				ForceNew:    true,
				Description: "An array of repository ids that can access the organization secret.",
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

func resourceGithubActionsOrganizationSecretStateUpgradeV0(ctx context.Context, rawState map[string]any, _ any) (map[string]any, error) {
	log.Printf("[DEBUG] GitHub Actions Organization Secret Attributes before migration: %#v", rawState)

	// Add the destroy_on_drift field with default value true if it doesn't exist
	if _, ok := rawState["destroy_on_drift"]; !ok {
		rawState["destroy_on_drift"] = true
	}

	log.Printf("[DEBUG] GitHub Actions Organization Secret Attributes after migration: %#v", rawState)

	return rawState, nil
}
