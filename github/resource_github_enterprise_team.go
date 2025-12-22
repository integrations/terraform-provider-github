package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	githubv3 "github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseTeam() *schema.Resource {
	return &schema.Resource{
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
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise (e.g. from the enterprise URL).",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 255)),
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the enterprise team.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 255)),
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A description of the enterprise team.",
			},
			"organization_selection_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "disabled",
				Description: "Controls which organizations can see this team: `disabled`, `selected`, or `all`.",
				ValidateDiagFunc: toDiagFunc(
					validation.StringInSlice([]string{"disabled", "selected", "all"}, false),
					"organization_selection_type",
				),
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

	req := enterpriseTeamCreateRequest{
		Name:                      name,
		Description:               githubv3.String(description),
		OrganizationSelectionType: githubv3.String(orgSelection),
	}
	if groupID != "" {
		req.GroupID = githubv3.String(groupID)
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())
	te, _, err := createEnterpriseTeam(ctx, client, enterpriseSlug, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(te.ID, 10))
	return resourceGithubEnterpriseTeamRead(context.WithValue(ctx, ctxId, d.Id()), d, meta)
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
	var te *enterpriseTeam
	if slug, ok := d.GetOk("slug"); ok {
		if s := strings.TrimSpace(slug.(string)); s != "" {
			candidate, _, getErr := getEnterpriseTeamBySlug(ctx, client, enterpriseSlug, s)
			if getErr == nil {
				te = candidate
			} else {
				ghErr := &githubv3.ErrorResponse{}
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
	orgSelection := te.OrganizationSelectionType
	if orgSelection == "" {
		orgSelection = "disabled"
	}
	if err = d.Set("organization_selection_type", orgSelection); err != nil {
		return diag.FromErr(err)
	}
	if te.GroupID != nil {
		if err = d.Set("group_id", *te.GroupID); err != nil {
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

	// We need a team slug for the API. If state is missing, re-discover it by ID.
	teamSlug := strings.TrimSpace(d.Get("slug").(string))
	if teamSlug == "" {
		teamID, err := strconv.ParseInt(d.Id(), 10, 64)
		if err != nil {
			return diag.FromErr(unconvertibleIdErr(d.Id(), err))
		}
		ctx = context.WithValue(ctx, ctxId, d.Id())
		te, err := findEnterpriseTeamByID(ctx, client, enterpriseSlug, teamID)
		if err != nil {
			return diag.FromErr(err)
		}
		if te == nil {
			return diag.FromErr(fmt.Errorf("enterprise team %s no longer exists", d.Id()))
		}
		teamSlug = te.Slug
	}

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	orgSelection := d.Get("organization_selection_type").(string)
	groupID := d.Get("group_id").(string)

	req := enterpriseTeamUpdateRequest{
		Name:                      githubv3.String(name),
		Description:               githubv3.String(description),
		OrganizationSelectionType: githubv3.String(orgSelection),
	}
	if groupID != "" {
		req.GroupID = githubv3.String(groupID)
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())
	_, _, err := updateEnterpriseTeam(ctx, client, enterpriseSlug, teamSlug, req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubEnterpriseTeamRead(ctx, d, meta)
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
	resp, err := deleteEnterpriseTeam(ctx, client, enterpriseSlug, teamSlug)
	if err != nil {
		// Already gone? That's fine, we wanted it deleted anyway.
		ghErr := &githubv3.ErrorResponse{}
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		_ = resp
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
	_ = d.Set("enterprise_slug", enterpriseSlug)
	return []*schema.ResourceData{d}, nil
}
