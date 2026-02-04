package github

import (
	"context"
	"strings"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseTeamOrganizations() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages organization assignments for a GitHub enterprise team.",
		CreateContext: resourceGithubEnterpriseTeamOrganizationsCreate,
		ReadContext:   resourceGithubEnterpriseTeamOrganizationsRead,
		UpdateContext: resourceGithubEnterpriseTeamOrganizationsUpdate,
		DeleteContext: resourceGithubEnterpriseTeamOrganizationsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "The slug of the enterprise.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringIsNotWhiteSpace, validation.StringIsNotEmpty)),
			},
			"team_slug": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Description:      "The slug of the enterprise team.",
				ExactlyOneOf:     []string{"team_slug", "team_id"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringIsNotWhiteSpace, validation.StringIsNotEmpty)),
			},
			"team_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				Description:  "The ID of the enterprise team.",
				ExactlyOneOf: []string{"team_slug", "team_id"},
			},
			"organization_slugs": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Set of organization slugs that the enterprise team should be assigned to.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				MinItems:    1,
			},
		},
	}
}

func resourceGithubEnterpriseTeamOrganizationsCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := strings.TrimSpace(d.Get("enterprise_slug").(string))

	// Get team by slug or ID
	var team *github.EnterpriseTeam
	var err error
	if v, ok := d.GetOk("team_slug"); ok {
		team, _, err = client.Enterprise.GetTeam(ctx, enterpriseSlug, v.(string))
	} else {
		teamID := int64(d.Get("team_id").(int))
		team, err = findEnterpriseTeamByID(ctx, client, enterpriseSlug, teamID)
	}
	if err != nil {
		return diag.FromErr(err)
	}
	if team == nil {
		return diag.Errorf("enterprise team not found")
	}

	orgSlugsSet := d.Get("organization_slugs").(*schema.Set)
	orgSlugs := make([]string, 0, orgSlugsSet.Len())
	for _, v := range orgSlugsSet.List() {
		orgSlugs = append(orgSlugs, v.(string))
	}

	// Add organizations to the team using the SDK
	_, _, err = client.Enterprise.AddMultipleAssignments(ctx, enterpriseSlug, team.Slug, orgSlugs)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildEnterpriseTeamOrganizationsID(enterpriseSlug, team.Slug))

	// Only set team_slug or team_id based on what user provided
	if _, ok := d.GetOk("team_slug"); ok {
		if err := d.Set("team_slug", team.Slug); err != nil {
			return diag.FromErr(err)
		}
	} else if _, ok := d.GetOk("team_id"); ok {
		if err := d.Set("team_id", int(team.ID)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubEnterpriseTeamOrganizationsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug, teamSlug, err := parseEnterpriseTeamOrganizationsID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	orgs, err := listAllEnterpriseTeamOrganizations(ctx, client, enterpriseSlug, teamSlug)
	if err != nil {
		return diag.FromErr(err)
	}

	slugs := make([]string, 0, len(orgs))
	for _, org := range orgs {
		if org.Login != nil && *org.Login != "" {
			slugs = append(slugs, *org.Login)
		}
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	// Only set team_slug if it was configured, or if neither team_slug nor team_id
	// is present (e.g., during import). This avoids drift when users configure team_id.
	if _, ok := d.GetOk("team_slug"); ok {
		if err := d.Set("team_slug", teamSlug); err != nil {
			return diag.FromErr(err)
		}
	} else if _, ok := d.GetOk("team_id"); !ok {
		// During import, neither is set, so we populate team_slug
		if err := d.Set("team_slug", teamSlug); err != nil {
			return diag.FromErr(err)
		}
	}
	if err := d.Set("organization_slugs", slugs); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseTeamOrganizationsUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug, teamSlug, err := parseEnterpriseTeamOrganizationsID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("organization_slugs") {
		oldVal, newVal := d.GetChange("organization_slugs")
		oldSet := oldVal.(*schema.Set)
		newSet := newVal.(*schema.Set)

		toAdd := newSet.Difference(oldSet)
		toRemove := oldSet.Difference(newSet)

		// Add new organizations
		if toAdd.Len() > 0 {
			addSlugs := make([]string, 0, toAdd.Len())
			for _, v := range toAdd.List() {
				addSlugs = append(addSlugs, v.(string))
			}
			_, _, err = client.Enterprise.AddMultipleAssignments(ctx, enterpriseSlug, teamSlug, addSlugs)
			if err != nil {
				return diag.FromErr(err)
			}
		}

		// Remove old organizations
		if toRemove.Len() > 0 {
			removeSlugs := make([]string, 0, toRemove.Len())
			for _, v := range toRemove.List() {
				removeSlugs = append(removeSlugs, v.(string))
			}
			_, _, err = client.Enterprise.RemoveMultipleAssignments(ctx, enterpriseSlug, teamSlug, removeSlugs)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return nil
}

func resourceGithubEnterpriseTeamOrganizationsDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug, teamSlug, err := parseEnterpriseTeamOrganizationsID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	// Get organizations from state
	orgSlugsSet := d.Get("organization_slugs").(*schema.Set)
	if orgSlugsSet.Len() > 0 {
		removeSlugs := make([]string, 0, orgSlugsSet.Len())
		for _, v := range orgSlugsSet.List() {
			removeSlugs = append(removeSlugs, v.(string))
		}
		_, resp, err := client.Enterprise.RemoveMultipleAssignments(ctx, enterpriseSlug, teamSlug, removeSlugs)
		if err != nil {
			// Already gone? That's fine, we wanted it deleted anyway.
			if resp != nil && resp.StatusCode == 404 {
				return nil
			}
			return diag.FromErr(err)
		}
	}

	return nil
}
