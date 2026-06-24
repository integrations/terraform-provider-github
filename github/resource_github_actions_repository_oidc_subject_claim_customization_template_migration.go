package github

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The name of the repository.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 100)),
			},
			"use_default": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether to use the default template or not. If 'true', 'include_claim_keys' must not be set.",
			},
			"include_claim_keys": {
				Type:        schema.TypeList,
				Optional:    true,
				MinItems:    1,
				Description: "A list of OpenID Connect claims.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateStateUpgradeV0(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	log.Printf("[DEBUG] OIDC Subject Claim Customization Template state before migration: %#v", rawState)

	repoName, ok := rawState["repository"].(string)
	if !ok {
		return nil, fmt.Errorf("repository not found or is not a string")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository %s: %w", repoName, err)
	}

	rawState["repository_id"] = int(repo.GetID())

	log.Printf("[DEBUG] OIDC Subject Claim Customization Template state after migration: %#v", rawState)

	return rawState, nil
}
