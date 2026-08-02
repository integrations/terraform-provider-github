package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryAutolinkReferenceV1() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_prefix": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_url_template": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_alphanumeric": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubRepositoryAutolinkReferenceStateUpgradeV1(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	tflog.Debug(ctx, "repository autolink reference state before v1 migration", rawState)

	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	migratedState, err := migrateRepositoryWithID(ctx, client, owner, rawState)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "repository autolink reference state after v1 migration", migratedState)

	return migratedState, nil
}
