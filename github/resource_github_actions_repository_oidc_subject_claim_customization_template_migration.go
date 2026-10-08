package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateResourceV0() *schema.Resource {
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

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateInstanceStateUpgradeV0(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	tflog.Debug(ctx, "Migrating repository OIDC subject claim customization template from v0 to v1.", rawState)

	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	migratedState, err := migrateRepositoryWithID(ctx, client, owner, rawState)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Migrated repository OIDC subject claim customization template to v1.", migratedState)
	return migratedState, nil
}
