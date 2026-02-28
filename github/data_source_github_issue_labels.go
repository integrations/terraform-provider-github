package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubIssueLabels() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubIssueLabelsRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"labels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"color": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubIssueLabelsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repository := d.Get("repository").(string)

	d.SetId(repository)

	labels, err := listLabels(client, ctx, owner, repository)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("labels", flattenLabels(labels))
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
