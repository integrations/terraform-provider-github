package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsSecretResourceV0() *schema.Resource {
	return &schema.Resource{
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
