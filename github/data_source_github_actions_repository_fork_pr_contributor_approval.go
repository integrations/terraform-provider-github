package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsRepositoryForkPRContributorApproval() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsRepositoryForkPRContributorApprovalRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GitHub repository.",
			},
			"approval_policy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The fork PR contributor approval policy currently configured on the repository. One of 'first_time_contributors_new_to_github', 'first_time_contributors', or 'all_external_contributors'.",
			},
		},
	}
}

func dataSourceGithubActionsRepositoryForkPRContributorApprovalRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repository := d.Get("repository").(string)

	policy, _, err := client.Actions.GetForkPRContributorApprovalPermissions(ctx, owner, repository)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(repository)
	if err := d.Set("approval_policy", policy.ApprovalPolicy); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
