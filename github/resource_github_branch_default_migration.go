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
	tflog.Debug(ctx, "Migrating GitHub Branch Default from v0 to v1.", rawState)

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
		return nil, fmt.Errorf("repository is not a string or not set")
	}

	rawState["repository"] = repoName

	migratedState, err := migrateRepositoryWithID(ctx, client, owner, rawState)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Migrated GitHub Branch Default from v1.", migratedState)
	return migratedState, nil
}
