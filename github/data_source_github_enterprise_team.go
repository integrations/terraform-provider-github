package github

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseTeam() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubEnterpriseTeamRead,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
			},
			"slug": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"team_id"},
				Description:   "The slug of the enterprise team.",
			},
			"team_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"slug"},
				Description:   "The numeric ID of the enterprise team.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the enterprise team.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A description of the enterprise team.",
			},
			"organization_selection_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies which organizations in the enterprise should have access to this team.",
			},
			"group_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the IdP group to assign team membership with.",
			},
		},
	}
}

func dataSourceGithubEnterpriseTeamRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	enterpriseSlug := strings.TrimSpace(d.Get("enterprise_slug").(string))
	if enterpriseSlug == "" {
		return fmt.Errorf("enterprise_slug must not be empty")
	}

	ctx := context.Background()

	var te *enterpriseTeam
	if v, ok := d.GetOk("team_id"); ok {
		teamID := int64(v.(int))
		if teamID != 0 {
			found, err := findEnterpriseTeamByID(ctx, client, enterpriseSlug, teamID)
			if err != nil {
				return err
			}
			if found == nil {
				return fmt.Errorf("could not find enterprise team %d in enterprise %s", teamID, enterpriseSlug)
			}
			te = found
		}
	}

	if te == nil {
		teamSlug := strings.TrimSpace(d.Get("slug").(string))
		if teamSlug == "" {
			return fmt.Errorf("one of slug or team_id must be set")
		}
		found, _, err := getEnterpriseTeamBySlug(ctx, client, enterpriseSlug, teamSlug)
		if err != nil {
			return err
		}
		te = found
	}

	d.SetId(buildSlashTwoPartID(enterpriseSlug, strconv.FormatInt(te.ID, 10)))
	_ = d.Set("enterprise_slug", enterpriseSlug)
	_ = d.Set("slug", te.Slug)
	_ = d.Set("team_id", int(te.ID))
	_ = d.Set("name", te.Name)
	if te.Description != nil {
		_ = d.Set("description", *te.Description)
	} else {
		_ = d.Set("description", "")
	}
	orgSel := te.OrganizationSelectionType
	if orgSel == "" {
		orgSel = "disabled"
	}
	_ = d.Set("organization_selection_type", orgSel)
	if te.GroupID != nil {
		_ = d.Set("group_id", *te.GroupID)
	} else {
		_ = d.Set("group_id", "")
	}

	return nil
}
