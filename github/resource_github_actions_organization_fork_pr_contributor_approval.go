package github

import (
	"context"

	"github.com/google/go-github/v86/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsOrganizationForkPRContributorApproval() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsOrganizationForkPRContributorApprovalCreateOrUpdate,
		ReadContext:   resourceGithubActionsOrganizationForkPRContributorApprovalRead,
		UpdateContext: resourceGithubActionsOrganizationForkPRContributorApprovalCreateOrUpdate,
		DeleteContext: resourceGithubActionsOrganizationForkPRContributorApprovalDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"approval_policy": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The organization-wide policy controlling which fork PR contributors need maintainer approval before their workflows can run. Possible values are 'first_time_contributors_new_to_github', 'first_time_contributors', or 'all_external_contributors'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"first_time_contributors_new_to_github",
					"first_time_contributors",
					"all_external_contributors",
				}, false)),
			},
		},
	}
}

func resourceGithubActionsOrganizationForkPRContributorApprovalCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	policy := github.ContributorApprovalPermissions{
		ApprovalPolicy: d.Get("approval_policy").(string),
	}

	if _, err := client.Actions.UpdateOrganizationForkPRContributorApprovalPermissions(ctx, orgName, policy); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(orgName)
	return resourceGithubActionsOrganizationForkPRContributorApprovalRead(ctx, d, meta)
}

func resourceGithubActionsOrganizationForkPRContributorApprovalRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	policy, _, err := client.Actions.GetOrganizationForkPRContributorApprovalPermissions(ctx, orgName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("approval_policy", policy.ApprovalPolicy); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsOrganizationForkPRContributorApprovalDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	// The API has no "off" state for this policy. Reset to the GitHub-documented
	// default (first_time_contributors) on Delete so the resource leaves no
	// residual non-default state behind.
	policy := github.ContributorApprovalPermissions{
		ApprovalPolicy: "first_time_contributors",
	}
	if _, err := client.Actions.UpdateOrganizationForkPRContributorApprovalPermissions(ctx, orgName, policy); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
