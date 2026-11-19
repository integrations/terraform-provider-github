package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryIssueCreationPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages the issue creation policy for a repository.",
		CreateContext: resourceGithubRepositoryIssueCreationPolicyCreate,
		ReadContext:   resourceGithubRepositoryIssueCreationPolicyRead,
		UpdateContext: resourceGithubRepositoryIssueCreationPolicyUpdate,
		DeleteContext: resourceGithubRepositoryIssueCreationPolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CustomizeDiff: diffRepository,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the GitHub repository.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The numeric ID of the GitHub repository.",
			},
			"policy": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Controls who can create issues for the repository. Can be `all` or `collaborators_only`.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "collaborators_only"}, false)),
			},
		},
	}
}

func resourceGithubRepositoryIssueCreationPolicyCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ownerMeta, _ := meta.(*Owner)
	owner := ownerMeta.name
	repoName, _ := d.Get("repository").(string)
	policy, _ := d.Get("policy").(string)

	nodeID, databaseID, err := getRepositoryNodeAndDatabaseID(ctx, owner, repoName, meta)
	if err != nil {
		return diag.Errorf("error resolving repository ID for %s: %s", repoName, err)
	}

	if err := updateRepositoryIssueCreationPolicy(ctx, nodeID, policy, meta); err != nil {
		return diag.Errorf("error setting issue creation policy for %s: %s", repoName, err)
	}

	d.SetId(repoName)
	if err := d.Set("repository_id", databaseID); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryIssueCreationPolicyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ownerMeta, _ := meta.(*Owner)
	owner := ownerMeta.name
	repoName := d.Id()

	policy, databaseID, err := getRepositoryIssueCreationPolicy(ctx, owner, repoName, meta)
	if err != nil {
		if isRepositoryNotFoundError(err) {
			tflog.Info(ctx, "Repository no longer exists, removing issue creation policy from state", map[string]any{
				"owner":      owner,
				"repository": repoName,
			})
			d.SetId("")
			return nil
		}
		return diag.Errorf("error reading issue creation policy for %s/%s: %s", owner, repoName, err)
	}

	if err := d.Set("policy", policy); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("repository", repoName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("repository_id", databaseID); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryIssueCreationPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ownerMeta, _ := meta.(*Owner)
	owner := ownerMeta.name
	repoName, _ := d.Get("repository").(string)
	policy, _ := d.Get("policy").(string)

	nodeID, databaseID, err := getRepositoryNodeAndDatabaseID(ctx, owner, repoName, meta)
	if err != nil {
		return diag.Errorf("error resolving repository ID for %s: %s", repoName, err)
	}

	if err := updateRepositoryIssueCreationPolicy(ctx, nodeID, policy, meta); err != nil {
		return diag.Errorf("error updating issue creation policy for %s: %s", repoName, err)
	}

	d.SetId(repoName)
	if err := d.Set("repository_id", databaseID); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryIssueCreationPolicyDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	repoName := d.Id()

	nodeID, err := getRepositoryID(repoName, meta)
	if err != nil {
		if isRepositoryNotFoundError(err) {
			tflog.Info(ctx, "Repository no longer exists, nothing to reset for the issue creation policy", map[string]any{
				"repository": repoName,
			})
			return nil
		}
		return diag.Errorf("error resolving repository node ID for %s: %s", repoName, err)
	}

	if err := updateRepositoryIssueCreationPolicy(ctx, nodeID, "all", meta); err != nil {
		return diag.Errorf("error resetting issue creation policy for %s: %s", repoName, err)
	}

	return nil
}
