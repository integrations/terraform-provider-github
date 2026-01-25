package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubBranchProtectionV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the repository.",
			},
			"branch": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The branch name or pattern to protect.",
			},
		},
	}
}

func resourceGithubBranchProtectionUpgradeV0(_ context.Context, rawState map[string]any, meta any) (map[string]any, error) {
	repoName := rawState["repository"].(string)
	repoID, err := getRepositoryID(repoName, meta)
	if err != nil {
		return nil, err
	}

	branch := rawState["branch"].(string)
	protectionRuleID, err := getBranchProtectionID(repoID, branch, meta)
	if err != nil {
		return nil, err
	}

	rawState["id"] = protectionRuleID
	rawState[REPOSITORY_ID] = repoID
	rawState[PROTECTION_PATTERN] = branch

	return rawState, nil
}

func resourceGithubBranchProtectionV1() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"push_restrictions": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The list of actor Names/IDs that may push to the branch.",
			},
			"blocks_creations": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Setting this to true will block creating the branch.",
			},
		},
	}
}

func resourceGithubBranchProtectionUpgradeV1(_ context.Context, rawState map[string]any, meta any) (map[string]any, error) {
	blocksCreations := false

	if v, ok := rawState["blocks_creations"]; ok {
		blocksCreations = v.(bool)
	}

	if v, ok := rawState["push_restrictions"]; ok {
		rawState["restrict_pushes"] = []any{map[string]any{
			"blocks_creations": blocksCreations,
			"push_allowances":  v,
		}}
	}

	delete(rawState, "blocks_creations")
	delete(rawState, "push_restrictions")

	return rawState, nil
}
