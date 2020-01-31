package github

import "github.com/hashicorp/terraform/helper/schema"

func resourceGithubRepositoryV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGithubRepositoryUpgradeV0(rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	repositoryName := rawState["name"].(string)
	repositoryID, err := getRepositoryID(repositoryName, meta)
	if err != nil {
		return nil, err
	}

	rawState["id"] = repositoryID

	return rawState, nil
}
