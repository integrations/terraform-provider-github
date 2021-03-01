package github

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

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

func resourceGithubBranchProtectionUpgradeV0(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
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
