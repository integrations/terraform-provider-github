package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubIssueLabels() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubIssueLabelsRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the repository.",
			},
			"labels": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of this repository's labels.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the label.",
						},
						"color": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The hexadecimal color code for the label, without the leading #.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A short description of the label.",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the label.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubIssueLabelsRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repository := d.Get("repository").(string)
	ctx := context.Background()

	d.SetId(repository)

	labels, err := listLabels(client, ctx, owner, repository)
	if err != nil {
		return err
	}

	err = d.Set("labels", flattenLabels(labels))
	if err != nil {
		return err
	}

	return nil
}
