package github

import (
	"context"

	"github.com/google/go-github/v47/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubRepositoryDeployKeys() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRepositoryDeployKeysRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"title": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"verified": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubRepositoryDeployKeysRead(d *schema.ResourceData, meta interface{}) error {
	repository := d.Get("repository").(string)
	owner := meta.(*Owner).name

	client := meta.(*Owner).v3client
	ctx := context.Background()

	keys, _, err := client.Repositories.ListKeys(ctx, owner, repository, nil)
	if err != nil {
		return err
	}

	d.SetId(repository)
	d.Set("keys", flattenGitHubDeployKeys(keys))

	return nil
}

func flattenGitHubDeployKeys(keys []*github.Key) []interface{} {
	if keys == nil {
		return make([]interface{}, 0)
	}

	results := make([]interface{}, 0)

	for _, c := range keys {
		result := make(map[string]interface{})

		result["id"] = c.ID
		result["key"] = c.Key
		result["title"] = c.Title
		result["verified"] = c.Verified

		results = append(results, result)
	}

	return results
}
