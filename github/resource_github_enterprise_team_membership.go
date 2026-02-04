package github

import (
	"context"
	"strings"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseTeamMembership() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages membership of a user in a GitHub enterprise team.",
		CreateContext: resourceGithubEnterpriseTeamMembershipCreate,
		ReadContext:   resourceGithubEnterpriseTeamMembershipRead,
		DeleteContext: resourceGithubEnterpriseTeamMembershipDelete,
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
			"username": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "The username of the user to add to the team.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringIsNotWhiteSpace, validation.StringIsNotEmpty)),
			},
			"user_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the user.",
			},
		},
	}
}

func resourceGithubEnterpriseTeamMembershipCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := strings.TrimSpace(d.Get("enterprise_slug").(string))
	username := strings.TrimSpace(d.Get("username").(string))

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

	// Add the user to the team using the SDK
	user, _, err := client.Enterprise.AddTeamMember(ctx, enterpriseSlug, team.Slug, username)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildEnterpriseTeamMembershipID(enterpriseSlug, team.Slug, username))

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

	if user != nil && user.ID != nil {
		if err := d.Set("user_id", int(*user.ID)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubEnterpriseTeamMembershipRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug, teamSlug, username, err := parseEnterpriseTeamMembershipID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if v, ok := d.GetOk("team_id"); ok {
		teamID := int64(v.(int))
		if teamID > 0 {
			team, findErr := findEnterpriseTeamByID(ctx, client, enterpriseSlug, teamID)
			if findErr != nil {
				return diag.FromErr(findErr)
			}
			if team == nil {
				d.SetId("")
				return nil
			}
			teamSlug = team.Slug
		}
	}

	// Get the membership using the SDK
	user, resp, err := client.Enterprise.GetTeamMembership(ctx, enterpriseSlug, teamSlug, username)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
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
	if err := d.Set("username", username); err != nil {
		return diag.FromErr(err)
	}
	if user != nil && user.ID != nil {
		if err := d.Set("user_id", int(*user.ID)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubEnterpriseTeamMembershipDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug, teamSlug, username, err := parseEnterpriseTeamMembershipID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if v, ok := d.GetOk("team_id"); ok {
		teamID := int64(v.(int))
		if teamID > 0 {
			team, findErr := findEnterpriseTeamByID(ctx, client, enterpriseSlug, teamID)
			if findErr != nil {
				return diag.FromErr(findErr)
			}
			if team == nil {
				return nil
			}
			teamSlug = team.Slug
		}
	}

	// Remove the user from the team using the SDK
	resp, err := client.Enterprise.RemoveTeamMember(ctx, enterpriseSlug, teamSlug, username)
	if err != nil {
		// Already gone? That's fine, we wanted it deleted anyway.
		if resp != nil && resp.StatusCode == 404 {
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}
