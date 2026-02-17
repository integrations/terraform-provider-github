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
			// lintignore:XS001 // No changes to old schema versions
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// lintignore:XS001 // No changes to old schema versions
			"file": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// lintignore:XS001 // No changes to old schema versions
			"content": {
				Type:     schema.TypeString,
				Required: true,
			},
			// lintignore:XS001 // No changes to old schema versions
			"branch": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			// lintignore:XS001,S020 // No changes to old schema versions
			"ref": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			// lintignore:XS001 // No changes to old schema versions
			"commit_sha": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// lintignore:XS001 // No changes to old schema versions
			"commit_message": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// lintignore:XS001 // No changes to old schema versions
			"commit_author": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"commit_email"},
			},
			// lintignore:XS001 // No changes to old schema versions
			"commit_email": {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{"commit_author"},
			},
			// lintignore:XS001 // No changes to old schema versions
			"sha": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// lintignore:XS001 // No changes to old schema versions
			"overwrite_on_create": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// lintignore:XS001 // No changes to old schema versions
			"autocreate_branch": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// lintignore:XS001 // No changes to old schema versions
			"autocreate_branch_source_branch": {
				Type:         schema.TypeString,
				Default:      "main",
				Optional:     true,
				RequiredWith: []string{"autocreate_branch"},
			},
			// lintignore:XS001 // No changes to old schema versions
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
	tflog.Debug(ctx, "GitHub Repository File State before migration", map[string]any{
		"rawState": rawState,
	})

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, ok := rawState["repository"].(string)
	if !ok {
		return nil, fmt.Errorf("repository not found or is not a string")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository '%s': %w", repoName, err)
	}

	rawState["repository_id"] = int(repo.GetID())

	// If branch is missing or empty, fetch the default branch from the repository
	if branch, ok := rawState["branch"].(string); !ok || branch == "" {
		rawState["branch"] = repo.GetDefaultBranch()
	}

	newResourceID, err := buildID(rawState["repository"].(string), rawState["file"].(string), rawState["branch"].(string))
	if err != nil {
		return nil, fmt.Errorf("failed to build ID: %w", err)
	}
	rawState["id"] = newResourceID

	tflog.Debug(ctx, "GitHub Repository File State after migration", map[string]any{
		"rawState": rawState,
	})
	return rawState, nil
}
