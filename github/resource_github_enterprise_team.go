package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseTeam() *schema.Resource {
	return &schema.Resource{
		Description:   "Creates and manages a GitHub enterprise team.",
		CreateContext: resourceGithubEnterpriseTeamCreate,
		ReadContext:   resourceGithubEnterpriseTeamRead,
		UpdateContext: resourceGithubEnterpriseTeamUpdate,
		DeleteContext: resourceGithubEnterpriseTeamDelete,
		Importer:      &schema.ResourceImporter{StateContext: resourceGithubEnterpriseTeamImport},

		CustomizeDiff: customdiff.All(
			customdiff.ComputedIf("slug", func(_ context.Context, d *schema.ResourceDiff, _ any) bool {
				return d.HasChange("name")
			}),
			// The SDK cannot encode an explicit null group_id to unlink an IdP group.
			customdiff.ForceNewIfChange("group_id", func(_ context.Context, old, next, _ any) bool {
				oldGroup, _ := old.(string)
				nextGroup, _ := next.(string)
				return oldGroup != "" && nextGroup == ""
			}),
		),

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "The slug of the enterprise (e.g. from the enterprise URL).",
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringLenBetween(1, 255), validation.StringIsNotWhiteSpace)),
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The name of the enterprise team.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringLenBetween(1, 255), validation.StringIsNotWhiteSpace)),
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
				Description: "The ID of the IdP group to assign team membership with. Removing an existing group ID replaces the team; changing it to another group ID updates the team in place.",
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

	req := buildEnterpriseTeamCreateRequest(name, description, orgSelection, groupID)

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
	enterpriseSlug := d.Get("enterprise_slug").(string)

	teamID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())

	owner, _ := meta.(*Owner)
	slug, _ := d.Get("slug").(string)
	te, err := findEnterpriseTeamByIdentity(owner, ctx, enterpriseSlug, slug, teamID)
	if err != nil {
		return diag.FromErr(err)
	}
	if te == nil {
		tflog.Info(ctx, "Removing missing enterprise team from state", map[string]any{"team_id": teamID})
		d.SetId("")
		return nil
	}

	if err = d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("name", te.Name); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("description", te.GetDescription()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("slug", te.Slug); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("team_id", int(te.ID)); err != nil {
		return diag.FromErr(err)
	}
	orgSelection := te.GetOrganizationSelectionType()
	if orgSelection == "" {
		orgSelection = "disabled"
	}
	if err = d.Set("organization_selection_type", orgSelection); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("group_id", te.GetGroupID()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseTeamUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	teamID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}
	owner, _ := meta.(*Owner)
	slug, _ := d.Get("slug").(string)
	team, err := findEnterpriseTeamByIdentity(owner, ctx, enterpriseSlug, slug, teamID)
	if err != nil {
		return diag.FromErr(err)
	}
	if team == nil {
		return diag.Errorf("enterprise team %d no longer exists", teamID)
	}
	teamSlug := team.Slug
	req := buildEnterpriseTeamUpdateRequest(d)

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
	teamID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Id(), err))
	}
	owner, _ := meta.(*Owner)
	slug, _ := d.Get("slug").(string)
	team, err := findEnterpriseTeamByIdentity(owner, ctx, enterpriseSlug, slug, teamID)
	if err != nil {
		return diag.FromErr(err)
	}
	if team == nil {
		return nil
	}
	teamSlug := team.Slug
	tflog.Info(ctx, "Deleting enterprise team", map[string]any{"team_id": teamID, "slug": teamSlug})

	_, err = client.Enterprise.DeleteTeam(ctx, enterpriseSlug, teamSlug)
	if err != nil {
		// Already gone? That's fine, we wanted it deleted anyway.
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}

// buildEnterpriseTeamCreateRequest omits unset optional values on creation.
func buildEnterpriseTeamCreateRequest(name, description, orgSelection, groupID string) github.EnterpriseTeamCreateOrUpdateRequest {
	req := github.EnterpriseTeamCreateOrUpdateRequest{
		Name:                      name,
		OrganizationSelectionType: new(orgSelection),
	}
	if description != "" {
		req.Description = new(description)
	}
	if groupID != "" {
		req.GroupID = new(groupID)
	}
	return req
}

// buildEnterpriseTeamUpdateRequest sends changed optional values, including an empty description.
func buildEnterpriseTeamUpdateRequest(d *schema.ResourceData) github.EnterpriseTeamCreateOrUpdateRequest {
	name, _ := d.Get("name").(string)
	req := github.EnterpriseTeamCreateOrUpdateRequest{Name: name}
	if d.HasChange("description") {
		description, _ := d.Get("description").(string)
		req.Description = new(description)
	}
	if d.HasChange("organization_selection_type") {
		selection, _ := d.Get("organization_selection_type").(string)
		req.OrganizationSelectionType = new(selection)
	}
	groupID, _ := d.Get("group_id").(string)
	if d.HasChange("group_id") && groupID != "" {
		req.GroupID = new(groupID)
	}
	return req
}

func resourceGithubEnterpriseTeamImport(_ context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	// Import format: <enterprise_slug>/<team_id>
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid import specified: supplied import must be written as <enterprise_slug>/<team_id>")
	}

	enterpriseSlug, teamID := parts[0], parts[1]
	if strings.TrimSpace(enterpriseSlug) == "" {
		return nil, fmt.Errorf("enterprise slug must not be empty")
	}
	id, err := strconv.ParseInt(teamID, 10, 64)
	if err != nil || id <= 0 {
		return nil, fmt.Errorf("team ID must be a positive integer: %q", teamID)
	}
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}
	d.SetId(teamID)
	return []*schema.ResourceData{d}, nil
}
