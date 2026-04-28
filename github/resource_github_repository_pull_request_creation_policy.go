package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryPullRequestCreationPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages the pull request creation policy for a repository.",
		CreateContext: resourceGithubRepositoryPullRequestCreationPolicyCreate,
		ReadContext:   resourceGithubRepositoryPullRequestCreationPolicyRead,
		UpdateContext: resourceGithubRepositoryPullRequestCreationPolicyUpdate,
		DeleteContext: resourceGithubRepositoryPullRequestCreationPolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryPullRequestCreationPolicyImport,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the GitHub repository.",
			},
			"policy": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Controls who can create pull requests for the repository. Can be `all` or `collaborators_only`.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "collaborators_only"}, false)),
			},
		},
	}
}

func resourceGithubRepositoryPullRequestCreationPolicyCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	repoName := d.Get("repository").(string)
	policy := d.Get("policy").(string)

	nodeID, err := getRepositoryID(repoName, meta)
	if err != nil {
		return diag.Errorf("error resolving repository node ID for %s: %s", repoName, err)
	}

	if err := updateRepositoryPullRequestCreationPolicy(ctx, nodeID, policy, meta); err != nil {
		return diag.Errorf("error setting pull request creation policy for %s: %s", repoName, err)
	}

	d.SetId(repoName)
	return resourceGithubRepositoryPullRequestCreationPolicyRead(ctx, d, meta)
}

func resourceGithubRepositoryPullRequestCreationPolicyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	owner := meta.(*Owner).name
	repoName := d.Id()

	policy, err := getRepositoryPullRequestCreationPolicy(ctx, owner, repoName, meta)
	if err != nil {
		return diag.Errorf("error reading pull request creation policy for %s/%s: %s", owner, repoName, err)
	}

	if err := d.Set("policy", policy); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("repository", repoName); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryPullRequestCreationPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	repoName := d.Id()
	policy := d.Get("policy").(string)

	nodeID, err := getRepositoryID(repoName, meta)
	if err != nil {
		return diag.Errorf("error resolving repository node ID for %s: %s", repoName, err)
	}

	if err := updateRepositoryPullRequestCreationPolicy(ctx, nodeID, policy, meta); err != nil {
		return diag.Errorf("error updating pull request creation policy for %s: %s", repoName, err)
	}

	return resourceGithubRepositoryPullRequestCreationPolicyRead(ctx, d, meta)
}

func resourceGithubRepositoryPullRequestCreationPolicyDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	repoName := d.Id()

	nodeID, err := getRepositoryID(repoName, meta)
	if err != nil {
		return diag.Errorf("error resolving repository node ID for %s: %s", repoName, err)
	}

	if err := updateRepositoryPullRequestCreationPolicy(ctx, nodeID, "all", meta); err != nil {
		return diag.Errorf("error resetting pull request creation policy for %s: %s", repoName, err)
	}

	return nil
}

func resourceGithubRepositoryPullRequestCreationPolicyImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	repoName := d.Id()

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}

	diags := resourceGithubRepositoryPullRequestCreationPolicyRead(ctx, d, meta)
	if diags.HasError() {
		return nil, fmt.Errorf("%s", diags[0].Summary)
	}

	return []*schema.ResourceData{d}, nil
}
