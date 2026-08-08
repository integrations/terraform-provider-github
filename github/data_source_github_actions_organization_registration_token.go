package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsOrganizationRegistrationToken() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsOrganizationRegistrationTokenRead,

		Schema: map[string]*schema.Schema{
			"token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expires_at": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubActionsOrganizationRegistrationTokenRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	tflog.Debug(ctx, "Creating a GitHub Actions organization registration token", map[string]any{
		"owner": owner,
	})
	token, _, err := client.Actions.CreateOrganizationRegistrationToken(ctx, owner)
	if err != nil {
		return diag.Errorf("error creating a GitHub Actions organization registration token for %s: %v", owner, err)
	}

	d.SetId(owner)
	err = d.Set("token", token.Token)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("expires_at", token.ExpiresAt.Unix())
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
