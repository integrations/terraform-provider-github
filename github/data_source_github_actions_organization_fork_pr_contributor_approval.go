package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsOrganizationForkPRContributorApproval() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsOrganizationForkPRContributorApprovalRead,

		Schema: map[string]*schema.Schema{
			"approval_policy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The organization-wide fork PR contributor approval policy currently configured. One of 'first_time_contributors_new_to_github', 'first_time_contributors', or 'all_external_contributors'.",
			},
		},
	}
}

func dataSourceGithubActionsOrganizationForkPRContributorApprovalRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	policy, _, err := client.Actions.GetOrganizationForkPRContributorApprovalPermissions(ctx, orgName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(orgName)
	if err := d.Set("approval_policy", policy.ApprovalPolicy); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
