package github

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

func resourceGithubBranchProtectionUpgradeV0(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	repoName := rawState["repository"].(string)
	repoID, err := getRepositoryID(repoName, meta)
	if err != nil {
		return nil, err
	}

	branch := rawState["branch"].(string)
	protectionRuleID, err := getBranchProtectionID(repoName, branch, meta)
	if err != nil {
		return nil, err
	}

	rawState["id"] = protectionRuleID
	rawState["repository_id"] = repoID
	rawState["pattern"] = branch

	return rawState, nil
}

func resourceGithubBranchProtectionV1() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"pattern": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGithubBranchProtectionUpgradeV1(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {

	repoName, err := getRepositoryName(rawState["repository_id"].(string), meta)
	if err != nil {
		return nil, err
	}

	rawState["repository"] = repoName
	rawState["branch"] = rawState["pattern"]
	rawState["id"] = fmt.Sprintf("%s:%s", rawState["repository"], rawState["branch"])

	return rawState, nil
}
