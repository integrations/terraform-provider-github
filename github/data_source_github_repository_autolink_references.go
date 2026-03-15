package github

import (
	"context"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryAutolinkReferences() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubRepositoryAutolinkReferencesRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"autolink_references": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key_prefix": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_url_template": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_alphanumeric": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubRepositoryAutolinkReferencesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)

	autoLinks, _, err := client.Repositories.ListAutolinks(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(repoName)
	if err := d.Set("autolink_references", flattenAutolinkReferences(autoLinks)); err != nil {
		return diag.FromErr(err)
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
