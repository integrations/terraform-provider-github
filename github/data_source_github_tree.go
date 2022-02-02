package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubTree() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubTreeRead,
		Schema: map[string]*schema.Schema{
			"recursive": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entries": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"sha": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tree_sha": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGithubTreeRead(d *schema.ResourceData, meta interface{}) error {
	owner := meta.(*Owner).name
	repository := d.Get("repository").(string)
	sha := d.Get("tree_sha").(string)
	recursive := d.Get("recursive").(bool)

	client := meta.(*Owner).v3client
	ctx := context.Background()

	tree, _, err := client.Git.GetTree(ctx, owner, repository, sha, recursive)

	if err != nil {
		return err
	}

	entries := make([]interface{}, 0, len(tree.Entries))

	for _, entry := range tree.Entries {
		entries = append(entries, map[string]interface{}{
			"path": entry.Path,
			"mode": entry.Mode,
			"type": entry.Type,
			"size": entry.Size,
			"sha":  entry.SHA,
		})
	}

	d.SetId(tree.GetSHA())
	d.Set("entries", entries)

	return nil
}
