package github

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseTeamMembership() *schema.Resource {
	return &schema.Resource{
		Description:   "Creates and manages membership of a user in a GitHub enterprise team.",
		CreateContext: resourceGithubEnterpriseTeamMembershipCreate,
		ReadContext:   resourceGithubEnterpriseTeamMembershipRead,
		DeleteContext: resourceGithubEnterpriseTeamMembershipDelete,
		Importer: &schema.ResourceImporter{
			StateContext: func(_ context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
				enterprise, selector, _, err := parseEnterpriseTeamMembershipID(d.Id())
				if err != nil {
					return nil, err
				}
				if err := importEnterpriseTeamSelector(d, enterprise, selector); err != nil {
					return nil, err
				}
				return []*schema.ResourceData{d}, nil
			},
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
				Description:      "The slug of the enterprise team. Specify exactly one of team_slug or team_id.",
				ExactlyOneOf:     []string{"team_slug", "team_id"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringIsNotWhiteSpace, validation.StringIsNotEmpty)),
			},
			"resolved_team_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The stable numeric identity of the team, retained across renames even when team_slug is configured.",
			},
			"team_id": {
				Type:             schema.TypeInt,
				Optional:         true,
				ForceNew:         true,
				Description:      "The positive numeric ID of the enterprise team. Specify exactly one of team_slug or team_id.",
				ExactlyOneOf:     []string{"team_slug", "team_id"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
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

	teamID, slug, err := resolveEnterpriseTeamForCreate(meta.(*Owner), ctx, enterpriseSlug, d)
	if err != nil {
		return diag.FromErr(err)
	}
	if teamID <= 0 {
		return diag.Errorf("enterprise team not found")
	}
	// The membership endpoints accept the numeric team ID directly.
	teamSelector := strconv.FormatInt(teamID, 10)

	user, _, err := client.Enterprise.AddTeamMember(ctx, enterpriseSlug, teamSelector, username)
	if err != nil {
		return diag.FromErr(err)
	}

	// The resource ID keeps the slug (or team_id) selector for compatibility.
	idSlug := slug
	if idSlug == "" {
		idSlug = teamSelector
	}
	d.SetId(buildEnterpriseTeamMembershipID(enterpriseSlug, idSlug, username))

	if err := d.Set("resolved_team_id", int(teamID)); err != nil {
		return diag.FromErr(err)
	}

	// Only set team_slug or team_id based on what user provided
	if _, ok := d.GetOk("team_slug"); ok {
		if err := d.Set("team_slug", slug); err != nil {
			return diag.FromErr(err)
		}
	} else if _, ok := d.GetOk("team_id"); ok {
		if err := d.Set("team_id", int(teamID)); err != nil {
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
	owner, _ := meta.(*Owner)
	teamID, err := enterpriseTeamIDForOperations(owner, ctx, enterpriseSlug, teamSlug, d)
	if err != nil {
		return diag.FromErr(err)
	}
	if teamID == 0 {
		d.SetId("")
		return nil
	}
	teamSelector := strconv.FormatInt(teamID, 10)

	user, resp, err := client.Enterprise.GetTeamMembership(ctx, enterpriseSlug, teamSelector, username)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := d.Set("resolved_team_id", int(teamID)); err != nil {
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
	owner, _ := meta.(*Owner)
	teamID, err := enterpriseTeamIDForOperations(owner, ctx, enterpriseSlug, teamSlug, d)
	if err != nil {
		return diag.FromErr(err)
	}
	if teamID == 0 {
		return nil
	}
	teamSelector := strconv.FormatInt(teamID, 10)

	resp, err := client.Enterprise.RemoveTeamMember(ctx, enterpriseSlug, teamSelector, username)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}
