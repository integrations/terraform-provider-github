package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubBranchDefaultV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"branch": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rename": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"etag": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceGithubBranchDefaultStateUpgradeV0(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	tflog.Debug(ctx, "state upgrade v0: State before v0 migration", rawState)

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
		if v, ok := rawState["id"]; ok {
			if s, ok := v.(string); ok && s != "" {
				repoName = s
			}
		}
	}
	if repoName == "" {
		return nil, fmt.Errorf("state upgrade v0: repository is not a string or not set")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("state upgrade v0: failed to retrieve repository '%s': %w", repoName, err)
	}

	rawState["repository"] = repoName
	rawState["repository_id"] = int(repo.GetID())

	tflog.Debug(ctx, "state upgrade v0: State after v0 migration", rawState)

	return rawState, nil
}
