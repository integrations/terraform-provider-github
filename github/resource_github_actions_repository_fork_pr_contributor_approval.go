package github

import (
	"context"

	"github.com/google/go-github/v86/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsRepositoryForkPRContributorApproval() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsRepositoryForkPRContributorApprovalCreateOrUpdate,
		ReadContext:   resourceGithubActionsRepositoryForkPRContributorApprovalRead,
		UpdateContext: resourceGithubActionsRepositoryForkPRContributorApprovalCreateOrUpdate,
		DeleteContext: resourceGithubActionsRepositoryForkPRContributorApprovalDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"approval_policy": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The policy controlling which fork PR contributors need maintainer approval before their workflows can run. Possible values are 'first_time_contributors_new_to_github', 'first_time_contributors', or 'all_external_contributors'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"first_time_contributors_new_to_github",
					"first_time_contributors",
					"all_external_contributors",
				}, false)),
			},
			"repository": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The GitHub repository.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 100)),
			},
		},
	}
}

func resourceGithubActionsRepositoryForkPRContributorApprovalCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	policy := github.ContributorApprovalPermissions{
		ApprovalPolicy: d.Get("approval_policy").(string),
	}

	if _, err := client.Actions.UpdateForkPRContributorApprovalPermissions(ctx, owner, repoName, policy); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(repoName)
	return resourceGithubActionsRepositoryForkPRContributorApprovalRead(ctx, d, meta)
}

func resourceGithubActionsRepositoryForkPRContributorApprovalRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Id()

	policy, _, err := client.Actions.GetForkPRContributorApprovalPermissions(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("repository", repoName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("approval_policy", policy.ApprovalPolicy); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsRepositoryForkPRContributorApprovalDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Id()

	// The API has no "off" state for this policy. Reset to the GitHub-documented
	// default (first_time_contributors) on Delete so the resource leaves no
	// residual non-default state behind.
	policy := github.ContributorApprovalPermissions{
		ApprovalPolicy: "first_time_contributors",
	}
	if _, err := client.Actions.UpdateForkPRContributorApprovalPermissions(ctx, owner, repoName, policy); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
