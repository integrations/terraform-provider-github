package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryFileV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"file": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content": {
				Type:     schema.TypeString,
				Required: true,
			},
			"branch": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ref": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"commit_sha": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"commit_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"commit_author": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"commit_email"},
			},
			"commit_email": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"commit_author"},
			},
			"sha": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"overwrite_on_create": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"autocreate_branch": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"autocreate_branch_source_branch": {
				Type:         schema.TypeString,
				Default:      "main",
				Optional:     true,
				RequiredWith: []string{"autocreate_branch"},
			},
			"autocreate_branch_source_sha": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				RequiredWith: []string{"autocreate_branch"},
			},
		},
	}
}

func resourceGithubRepositoryFileStateUpgradeV0(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	tflog.Debug(ctx, "GitHub Repository File State before v0 migration", map[string]any{
		"rawState": rawState,
	})

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, ok := rawState["repository"].(string)
	if !ok {
		return nil, fmt.Errorf("state upgrade v0: repository not found or is not a string")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("state upgrade v0: failed to retrieve repository '%s': %w", repoName, err)
	}

	rawState["repository_id"] = int(repo.GetID())

	branch, ok := rawState["branch"].(string)
	// If branch is missing or empty, fetch the default branch from the repository
	if !ok || branch == "" {
		branch = repo.GetDefaultBranch()
		rawState["branch"] = branch
	}

	filePath, ok := rawState["file"].(string)
	if !ok {
		return nil, fmt.Errorf("state upgrade v0: file path is not a string")
	}

	newResourceID, err := buildID(repoName, escapeIDPart(filePath), branch)
	if err != nil {
		return nil, fmt.Errorf("state upgrade v0: failed to build ID: %w", err)
	}
	rawState["id"] = newResourceID

	tflog.Debug(ctx, "GitHub Repository File State after v0 migration", map[string]any{
		"rawState": rawState,
	})
	return rawState, nil
}
