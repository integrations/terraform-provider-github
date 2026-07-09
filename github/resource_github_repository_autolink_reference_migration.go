package github

import (
	"context"
	"fmt"

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
	tflog.Debug(ctx, "GitHub Repository Autolink Reference state before v1 migration", rawState)

	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := ""
	if v, ok := rawState["repository"]; ok {
		if s, ok := v.(string); ok && s != "" {
			repoName = s
		}
	}
	if repoName == "" {
		return nil, fmt.Errorf("state upgrade v1: repository is not a string or not set")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("state upgrade v1: failed to retrieve repository '%s/%s': %w", owner, repoName, err)
	}

	rawState["repository_id"] = int(repo.GetID())

	tflog.Debug(ctx, "GitHub Repository Autolink Reference state after v1 migration", rawState)

	return rawState, nil
}
