package github

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubEnterpriseTeam() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubEnterpriseTeamRead,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The slug of the enterprise.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringIsNotWhiteSpace, validation.StringIsNotEmpty)),
			},
			"slug": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"team_id"},
				Description:  "The slug of the enterprise team.",
			},
			"team_id": {
				Type:             schema.TypeInt,
				Optional:         true,
				Computed:         true,
				ConflictsWith:    []string{"slug"},
				Description:      "The numeric ID of the enterprise team.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
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

func dataSourceGithubEnterpriseTeamRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := strings.TrimSpace(d.Get("enterprise_slug").(string))

	var te *enterpriseTeam
	if v, ok := d.GetOk("team_id"); ok {
		teamID := int64(v.(int))
		if teamID != 0 {
			found, err := findEnterpriseTeamByID(ctx, client, enterpriseSlug, teamID)
			if err != nil {
				return diag.FromErr(err)
			}
			if found == nil {
				return diag.FromErr(fmt.Errorf("could not find enterprise team %d in enterprise %s", teamID, enterpriseSlug))
			}
			te = found
		}
	}

	if te == nil {
		teamSlug := strings.TrimSpace(d.Get("slug").(string))
		if teamSlug == "" {
			return diag.FromErr(fmt.Errorf("one of slug or team_id must be set"))
		}
		found, _, err := getEnterpriseTeamBySlug(ctx, client, enterpriseSlug, teamSlug)
		if err != nil {
			return diag.FromErr(err)
		}
		te = found
	}

	d.SetId(buildSlashTwoPartID(enterpriseSlug, strconv.FormatInt(te.ID, 10)))
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("slug", te.Slug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("team_id", int(te.ID)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", te.Name); err != nil {
		return diag.FromErr(err)
	}
	if te.Description != nil {
		if err := d.Set("description", *te.Description); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err := d.Set("description", ""); err != nil {
			return diag.FromErr(err)
		}
	}
	orgSel := te.OrganizationSelectionType
	if orgSel == "" {
		orgSel = "disabled"
	}
	if err := d.Set("organization_selection_type", orgSel); err != nil {
		return diag.FromErr(err)
	}
	if te.GroupID != nil {
		if err := d.Set("group_id", *te.GroupID); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err := d.Set("group_id", ""); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}
