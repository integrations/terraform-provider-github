package github

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubEnterpriseTeamMembership() *schema.Resource {
	return &schema.Resource{
		Description: "Manages membership in a GitHub enterprise team.",
		ReadContext: dataSourceGithubEnterpriseTeamMembershipRead,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The slug of the enterprise.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringIsNotWhiteSpace, validation.StringIsNotEmpty)),
			},
			"enterprise_team": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The slug or ID of the enterprise team.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringIsNotWhiteSpace, validation.StringIsNotEmpty)),
			},
			"username": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The GitHub username.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringIsNotWhiteSpace, validation.StringIsNotEmpty)),
			},
			"role": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The role of the user in the enterprise team, if returned by the API.",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The membership state, if returned by the API.",
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ETag of the membership response.",
			},
		},
	}
}

func dataSourceGithubEnterpriseTeamMembershipRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := strings.TrimSpace(d.Get("enterprise_slug").(string))
	enterpriseTeam := strings.TrimSpace(d.Get("enterprise_team").(string))
	username := strings.TrimSpace(d.Get("username").(string))
	m, resp, err := getEnterpriseTeamMembershipDetails(ctx, client, enterpriseSlug, enterpriseTeam, username)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildSlashThreePartID(enterpriseSlug, enterpriseTeam, username))
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enterprise_team", enterpriseTeam); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("username", username); err != nil {
		return diag.FromErr(err)
	}
	if m != nil {
		if err := d.Set("role", m.Role); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("state", m.State); err != nil {
			return diag.FromErr(err)
		}
	}
	if resp != nil {
		if err := d.Set("etag", resp.Header.Get("ETag")); err != nil {
			return diag.FromErr(err)
		}
	}
	return nil
}
