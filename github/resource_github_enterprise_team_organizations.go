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

func resourceGithubEnterpriseTeamOrganizations() *schema.Resource {
	return &schema.Resource{
		Description:   "Creates and manages organization assignments for a GitHub enterprise team.",
		CreateContext: resourceGithubEnterpriseTeamOrganizationsCreate,
		ReadContext:   resourceGithubEnterpriseTeamOrganizationsRead,
		UpdateContext: resourceGithubEnterpriseTeamOrganizationsUpdate,
		DeleteContext: resourceGithubEnterpriseTeamOrganizationsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: func(_ context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
				enterprise, selector, err := parseEnterpriseTeamOrganizationsID(d.Id())
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
				Description:      "The slug of the enterprise team. Specify exactly one of team_slug or team_id. Not ForceNew: updates verify it still resolves to the managed team's numeric identity.",
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
			"organization_slugs": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Non-empty set of non-blank organization slugs that the enterprise team should be assigned to. Slugs are case-insensitive and stored in lowercase.",
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringIsNotWhiteSpace, validation.StringIsNotEmpty)),
					StateFunc:        func(v any) string { slug, _ := v.(string); return strings.ToLower(slug) },
				},
				Set:      func(v any) int { slug, _ := v.(string); return schema.HashString(strings.ToLower(slug)) },
				MinItems: 1,
			},
		},
	}
}

func resourceGithubEnterpriseTeamOrganizationsCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := strings.TrimSpace(d.Get("enterprise_slug").(string))

	teamID, slug, err := resolveEnterpriseTeamForCreate(meta.(*Owner), ctx, enterpriseSlug, d)
	if err != nil {
		return diag.FromErr(err)
	}
	if teamID <= 0 {
		return diag.Errorf("enterprise team not found")
	}
	// The assignment endpoints accept the numeric team ID directly.
	teamSelector := strconv.FormatInt(teamID, 10)

	// Verify no organizations are already assigned (authoritative resource).
	// A 404 here means the team has no assignments yet — treat as empty and proceed.
	existing, err := listAllEnterpriseTeamOrganizations(meta.(*Owner), ctx, enterpriseSlug, teamSelector)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotFound {
			existing = nil
		} else {
			return diag.FromErr(err)
		}
	}
	if len(existing) > 0 {
		return diag.Errorf("team %d already has organizations assigned; import first or remove manually", teamID)
	}

	orgSlugsSet := d.Get("organization_slugs").(*schema.Set)
	orgSlugs := make([]string, 0, orgSlugsSet.Len())
	for _, item := range orgSlugsSet.List() {
		slug, _ := item.(string)
		orgSlugs = append(orgSlugs, strings.ToLower(slug))
	}

	_, _, err = client.Enterprise.AddMultipleAssignments(ctx, enterpriseSlug, teamSelector, orgSlugs)
	if err != nil {
		return diag.FromErr(err)
	}

	// The resource ID keeps the slug (or team_id) selector for compatibility.
	idSlug := slug
	if idSlug == "" {
		idSlug = teamSelector
	}
	d.SetId(buildEnterpriseTeamOrganizationsID(enterpriseSlug, idSlug))

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

	return nil
}

func resourceGithubEnterpriseTeamOrganizationsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	enterpriseSlug, teamSlug, err := parseEnterpriseTeamOrganizationsID(d.Id())
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

	orgs, err := listAllEnterpriseTeamOrganizations(meta.(*Owner), ctx, enterpriseSlug, teamSelector)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	slugs := organizationSlugs(orgs)
	for i, slug := range slugs {
		slugs[i] = strings.ToLower(slug)
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

	owner, _ := meta.(*Owner)
	teamID, err := enterpriseTeamIDForOperations(owner, ctx, enterpriseSlug, teamSlug, d)
	if err != nil {
		return diag.FromErr(err)
	}
	if teamID == 0 {
		return diag.Errorf("enterprise team no longer exists")
	}
	teamSelector := strconv.FormatInt(teamID, 10)

	// Allow renamed slugs only if they still identify the managed team.
	// Validate before mutations; a 404 does not prove a rename.
	idSlug := teamSlug
	if v, ok := d.GetOk("team_slug"); ok {
		requestedSlug := strings.TrimSpace(v.(string))
		if requestedSlug == "" {
			return diag.Errorf("team_slug must not be empty")
		}
		requested, _, err := client.Enterprise.GetTeam(ctx, enterpriseSlug, requestedSlug)
		if err != nil {
			return diag.FromErr(err)
		}
		if requested == nil || requested.ID != teamID {
			if requested != nil {
				return diag.Errorf(
					"team_slug %q resolves to a different enterprise team (ID %d, slug %q) than the managed team (resolved_team_id %d); to manage a different team, import it instead of changing team_slug",
					requestedSlug, requested.ID, requested.Slug, teamID)
			}
			return diag.Errorf("team_slug %q does not resolve to the managed enterprise team (resolved_team_id %d)", requestedSlug, teamID)
		}
		idSlug = requestedSlug
	}

	if d.HasChange("organization_slugs") {
		oldVal, newVal := d.GetChange("organization_slugs")
		oldSet := oldVal.(*schema.Set)
		newSet := newVal.(*schema.Set)

		toAdd := newSet.Difference(oldSet)
		toRemove := oldSet.Difference(newSet)

		if toAdd.Len() > 0 {
			addSlugs := make([]string, 0, toAdd.Len())
			for _, item := range toAdd.List() {
				slug, _ := item.(string)
				addSlugs = append(addSlugs, strings.ToLower(slug))
			}
			_, _, err = client.Enterprise.AddMultipleAssignments(ctx, enterpriseSlug, teamSelector, addSlugs)
			if err != nil {
				return diag.FromErr(err)
			}
		}

		if toRemove.Len() > 0 {
			removeSlugs := make([]string, 0, toRemove.Len())
			for _, item := range toRemove.List() {
				slug, _ := item.(string)
				removeSlugs = append(removeSlugs, strings.ToLower(slug))
			}
			_, _, err = client.Enterprise.RemoveMultipleAssignments(ctx, enterpriseSlug, teamSelector, removeSlugs)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	d.SetId(buildEnterpriseTeamOrganizationsID(enterpriseSlug, idSlug))
	return nil
}

func resourceGithubEnterpriseTeamOrganizationsDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug, teamSlug, err := parseEnterpriseTeamOrganizationsID(d.Id())
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

	orgs, err := listAllEnterpriseTeamOrganizations(meta.(*Owner), ctx, enterpriseSlug, teamSelector)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		return diag.FromErr(err)
	}

	removeSlugs := organizationSlugs(orgs)

	if len(removeSlugs) > 0 {
		_, resp, err := client.Enterprise.RemoveMultipleAssignments(ctx, enterpriseSlug, teamSelector, removeSlugs)
		if err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotFound {
				return nil
			}
			if resp != nil && resp.StatusCode == http.StatusNotFound {
				return nil
			}
			return diag.FromErr(err)
		}
	}

	return nil
}
