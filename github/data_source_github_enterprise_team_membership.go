package github

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseTeamMembership() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubEnterpriseTeamMembershipRead,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
			},
			"enterprise_team": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug or ID of the enterprise team.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GitHub username.",
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

func dataSourceGithubEnterpriseTeamMembershipRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	enterpriseSlug := strings.TrimSpace(d.Get("enterprise_slug").(string))
	enterpriseTeam := strings.TrimSpace(d.Get("enterprise_team").(string))
	username := strings.TrimSpace(d.Get("username").(string))
	if enterpriseSlug == "" {
		return fmt.Errorf("enterprise_slug must not be empty")
	}
	if enterpriseTeam == "" {
		return fmt.Errorf("enterprise_team must not be empty")
	}
	if username == "" {
		return fmt.Errorf("username must not be empty")
	}

	ctx := context.Background()
	m, resp, err := getEnterpriseTeamMembershipDetails(ctx, client, enterpriseSlug, enterpriseTeam, username)
	if err != nil {
		return err
	}

	d.SetId(buildSlashThreePartID(enterpriseSlug, enterpriseTeam, username))
	_ = d.Set("enterprise_slug", enterpriseSlug)
	_ = d.Set("enterprise_team", enterpriseTeam)
	_ = d.Set("username", username)
	if m != nil {
		_ = d.Set("role", m.Role)
		_ = d.Set("state", m.State)
	}
	if resp != nil {
		_ = d.Set("etag", resp.Header.Get("ETag"))
	}
	return nil
}
