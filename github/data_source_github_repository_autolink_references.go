package github

import (
	"context"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryAutolinkReferences() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRepositoryAutolinkReferencesRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository to retrieve the autolink references from.",
			},
			"autolink_references": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of this repository's autolink references.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key_prefix": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Key prefix.",
						},
						"target_url_template": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Target URL template.",
						},
						"is_alphanumeric": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether this autolink reference matches alphanumeric characters.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubRepositoryAutolinkReferencesRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	results := make([]map[string]any, 0)

	for {
		listOptions := &github.ListOptions{}
		autoLinks, resp, err := client.Repositories.ListAutolinks(context.Background(), orgName, repoName, listOptions)
		if err != nil {
			return err
		}

		results = append(results, flattenAutolinkReferences(autoLinks)...)

		if resp.NextPage == 0 {
			break
		}

		listOptions.Page = resp.NextPage
	}

	d.SetId(repoName)
	err := d.Set("autolink_references", results)
	if err != nil {
		return err
	}

	return nil
}

func flattenAutolinkReferences(autoLinks []*github.Autolink) []map[string]any {
	results := make([]map[string]any, 0)
	if autoLinks == nil {
		return results
	}

	for _, autolink := range autoLinks {
		linkMap := make(map[string]any)
		linkMap["key_prefix"] = autolink.GetKeyPrefix()
		linkMap["target_url_template"] = autolink.GetURLTemplate()
		linkMap["is_alphanumeric"] = autolink.GetIsAlphanumeric()
		results = append(results, linkMap)
	}

	return results
}
