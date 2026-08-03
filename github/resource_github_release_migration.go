package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubReleaseV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the repository.",
			},
			"tag_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the tag.",
			},
			"target_commitish": {
				Type:        schema.TypeString,
				Default:     "main",
				Optional:    true,
				ForceNew:    true,
				Description: " The branch name or commit SHA the tag is created from.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    false,
				Description: "The name of the release.",
			},
			"body": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    false,
				Description: "Text describing the contents of the tag.",
			},
			"draft": {
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
				ForceNew:    true,
				Description: "Set to 'false' to create a published release.",
			},
			"prerelease": {
				Type:        schema.TypeBool,
				Default:     true,
				Optional:    true,
				Description: "Set to 'false' to identify the release as a full release.",
			},
			"generate_release_notes": {
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
				Description: "Set to 'true' to automatically generate the name and body for this release. If 'name' is specified, the specified name will be used; otherwise, a name will be automatically generated. If 'body' is specified, the body will be pre-pended to the automatically generated notes.",
			},
			"discussion_category_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository.",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"release_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the release.",
			},
			"node_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The node ID of the release.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the release was created.",
			},
			"published_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the release was published.",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for the release.",
			},
			"html_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HTML URL for the release.",
			},
			"assets_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for the release assets.",
			},
			"upload_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for the uploaded assets of release.",
			},
			"zipball_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for the zipball of the release.",
			},
			"tarball_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for the tarball of the release.",
			},
		},
	}
}

func resourceGithubReleaseStateUpgradeV0(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	tflog.Debug(ctx, "Migrating GitHub Release from v0 to v1.", map[string]any{"raw_state": rawState})

	state, err := migrateRepositoryWithID(ctx, client, owner, rawState)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "GitHub Release migrated to v1.", map[string]any{"raw_state": state})

	return state, nil
}
