package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubBranchProtectionV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"branch": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGithubBranchProtectionUpgradeV0(_ context.Context, rawState map[string]any, meta any) (map[string]any, error) {
	repoName, ok := rawState["repository"].(string)
	if !ok {
		return nil, fmt.Errorf("repository not found or is not a string")
	}
	repoID, err := getRepositoryID(repoName, meta)
	if err != nil {
		return nil, err
	}

	branch, ok := rawState["branch"].(string)
	if !ok {
		return nil, fmt.Errorf("branch not found or is not a string")
	}
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
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"blocks_creations": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceGithubBranchProtectionUpgradeV1(_ context.Context, rawState map[string]any, meta any) (map[string]any, error) {
	blocksCreations := false

	if v, ok := rawState["blocks_creations"]; ok {
		if b, ok := v.(bool); ok {
			blocksCreations = b
		}
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
