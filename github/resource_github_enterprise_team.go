package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseTeam() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages a GitHub enterprise team.",
		CreateContext: resourceGithubEnterpriseTeamCreate,
		ReadContext:   resourceGithubEnterpriseTeamRead,
		UpdateContext: resourceGithubEnterpriseTeamUpdate,
		DeleteContext: resourceGithubEnterpriseTeamDelete,
		Importer:      &schema.ResourceImporter{StateContext: resourceGithubEnterpriseTeamImport},

		CustomizeDiff: customdiff.Sequence(
			customdiff.ComputedIf("slug", func(_ context.Context, d *schema.ResourceDiff, meta any) bool {
				return d.HasChange("name")
			}),
		),

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "The slug of the enterprise (e.g. from the enterprise URL).",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 255)),
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The name of the enterprise team.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 255)),
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A description of the enterprise team.",
			},
			"organization_selection_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "disabled",
				Description:      "Controls which organizations can see this team: `disabled`, `selected`, or `all`.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"disabled", "selected", "all"}, false)),
			},
			"group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the IdP group to assign team membership with.",
			},
			"slug": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The slug of the enterprise team. GitHub generates the slug from the team name and adds the ent: prefix.",
			},
			"team_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The numeric ID of the enterprise team.",
			},
		},
	}
}

func resourceGithubEnterpriseTeamCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	orgSelection := d.Get("organization_selection_type").(string)
	groupID := d.Get("group_id").(string)

	req := github.EnterpriseTeamCreateOrUpdateRequest{
		Name:                      name,
		OrganizationSelectionType: github.Ptr(orgSelection),
		GroupID:                   github.Ptr(groupID), // Empty string is valid for no group
	}
	if description != "" {
		req.Description = github.Ptr(description)
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())
	te, _, err := client.Enterprise.CreateTeam(ctx, enterpriseSlug, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(te.ID, 10))

	// Set computed fields directly from API response
	if err := d.Set("slug", te.Slug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("team_id", int(te.ID)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseTeamRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)

	teamID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())

	// Try to fetch by slug first (faster), but if the team was renamed we need
	// to fall back to listing all teams and matching by numeric ID.
	var te *github.EnterpriseTeam
	if slug, ok := d.GetOk("slug"); ok {
		if s := strings.TrimSpace(slug.(string)); s != "" {
			candidate, _, getErr := client.Enterprise.GetTeam(ctx, enterpriseSlug, s)
			if getErr == nil {
				te = candidate
			} else {
				ghErr := &github.ErrorResponse{}
				if errors.As(getErr, &ghErr) && ghErr.Response.StatusCode != http.StatusNotFound {
					return diag.FromErr(getErr)
				}
			}
		}
	}

	if te == nil {
		te, err = findEnterpriseTeamByID(ctx, client, enterpriseSlug, teamID)
		if err != nil {
			return diag.FromErr(err)
		}
		if te == nil {
			log.Printf("[INFO] Removing enterprise team %s/%s from state because it no longer exists in GitHub", enterpriseSlug, d.Id())
			d.SetId("")
			return nil
		}
	}

	if err = d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("name", te.Name); err != nil {
		return diag.FromErr(err)
	}
	if te.Description != nil {
		if err = d.Set("description", *te.Description); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err = d.Set("description", ""); err != nil {
			return diag.FromErr(err)
		}
	}
	if err = d.Set("slug", te.Slug); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("team_id", int(te.ID)); err != nil {
		return diag.FromErr(err)
	}
	orgSelection := ""
	if te.OrganizationSelectionType != nil {
		orgSelection = *te.OrganizationSelectionType
	}
	if orgSelection == "" {
		orgSelection = "disabled"
	}
	if err = d.Set("organization_selection_type", orgSelection); err != nil {
		return diag.FromErr(err)
	}
	if te.GroupID != "" {
		if err = d.Set("group_id", te.GroupID); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err = d.Set("group_id", ""); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubEnterpriseTeamUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	teamSlug := d.Get("slug").(string)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	orgSelection := d.Get("organization_selection_type").(string)
	groupID := d.Get("group_id").(string)

	req := github.EnterpriseTeamCreateOrUpdateRequest{
		Name:                      name,
		OrganizationSelectionType: github.Ptr(orgSelection),
		GroupID:                   github.Ptr(groupID), // Empty string clears the group
	}
	if description != "" {
		req.Description = github.Ptr(description)
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())
	te, _, err := client.Enterprise.UpdateTeam(ctx, enterpriseSlug, teamSlug, req)
	if err != nil {
		return diag.FromErr(err)
	}

	// Update slug in case it changed (e.g., team was renamed)
	if err := d.Set("slug", te.Slug); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseTeamDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)

	ctx = context.WithValue(ctx, ctxId, d.Id())
	teamSlug := strings.TrimSpace(d.Get("slug").(string))
	if teamSlug == "" {
		teamID, err := strconv.ParseInt(d.Id(), 10, 64)
		if err != nil {
			return diag.FromErr(unconvertibleIdErr(d.Id(), err))
		}
		te, err := findEnterpriseTeamByID(ctx, client, enterpriseSlug, teamID)
		if err != nil {
			return diag.FromErr(err)
		}
		if te == nil {
			return nil
		}
		teamSlug = te.Slug
	}

	log.Printf("[INFO] Deleting enterprise team: %s/%s (%s)", enterpriseSlug, teamSlug, d.Id())
	_, err := client.Enterprise.DeleteTeam(ctx, enterpriseSlug, teamSlug)
	if err != nil {
		// Already gone? That's fine, we wanted it deleted anyway.
		ghErr := &github.ErrorResponse{}
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseTeamImport(_ context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	// Import format: <enterprise_slug>/<team_id>
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid import specified: supplied import must be written as <enterprise_slug>/<team_id>")
	}

	enterpriseSlug, teamID := parts[0], parts[1]
	d.SetId(teamID)
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
