package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubTree() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve the file tree for a repository.",
		Read:        dataSourceGithubTreeRead,
		Schema: map[string]*schema.Schema{
			"recursive": {
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
				Description: "Setting this to true returns the objects or subtrees referenced by the tree specified in tree_sha.",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the repository.",
			},
			"entries": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Objects (of path, mode, type, size, and sha) specifying a tree structure.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The path of the tree entry relative to its parent tree.",
						},
						"mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The file mode (100644 for file, 100755 for executable, 040000 for subdirectory, 160000 for submodule, 120000 for symlink).",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the entry (blob, tree, or commit).",
						},
						"size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of the blob in bytes (only present for blobs).",
						},
						"sha": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SHA1 of the object this tree entry points to.",
						},
					},
				},
			},
			"tree_sha": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The SHA1 value for the tree.",
			},
		},
	}
}

func dataSourceGithubTreeRead(d *schema.ResourceData, meta any) error {
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

	entries := make([]any, 0, len(tree.Entries))

	for _, entry := range tree.Entries {
		entries = append(entries, map[string]any{
			"path": entry.Path,
			"mode": entry.Mode,
			"type": entry.Type,
			"size": entry.Size,
			"sha":  entry.SHA,
		})
	}

	d.SetId(tree.GetSHA())
	if err = d.Set("entries", entries); err != nil {
		return err
	}

	return nil
}
