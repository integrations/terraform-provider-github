package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	githubv3 "github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseTeamOrganizations() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages organization assignments for a GitHub enterprise team.",
		CreateContext: resourceGithubEnterpriseTeamOrganizationsCreateOrUpdate,
		ReadContext:   resourceGithubEnterpriseTeamOrganizationsRead,
		UpdateContext: resourceGithubEnterpriseTeamOrganizationsCreateOrUpdate,
		DeleteContext: resourceGithubEnterpriseTeamOrganizationsDelete,
		Importer:      &schema.ResourceImporter{StateContext: resourceGithubEnterpriseTeamOrganizationsImport},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "The slug of the enterprise.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringIsNotWhiteSpace, validation.StringIsNotEmpty)),
			},
			"enterprise_team": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "The slug or ID of the enterprise team.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringIsNotWhiteSpace, validation.StringIsNotEmpty)),
			},
			"organization_slugs": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Set of organization slugs that the enterprise team should be assigned to.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
			},
		},
	}
}

func resourceGithubEnterpriseTeamOrganizationsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	enterpriseTeam := d.Get("enterprise_team").(string)

	desiredSet := map[string]struct{}{}
	if v, ok := d.GetOk("organization_slugs"); ok {
		for _, s := range v.(*schema.Set).List() {
			slug := strings.TrimSpace(s.(string))
			if slug != "" {
				desiredSet[slug] = struct{}{}
			}
		}
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())
	current, err := listEnterpriseTeamOrganizations(ctx, client, enterpriseSlug, enterpriseTeam)
	if err != nil {
		return diag.FromErr(err)
	}

	currentSet := map[string]struct{}{}
	for _, org := range current {
		if org.Login != "" {
			currentSet[org.Login] = struct{}{}
		}
	}

	toAdd := []string{}
	for slug := range desiredSet {
		if _, ok := currentSet[slug]; !ok {
			toAdd = append(toAdd, slug)
		}
	}

	toRemove := []string{}
	for slug := range currentSet {
		if _, ok := desiredSet[slug]; !ok {
			toRemove = append(toRemove, slug)
		}
	}

	// Perform adds before removes to avoid transient states where the team has no orgs
	if err := addEnterpriseTeamOrganizations(ctx, client, enterpriseSlug, enterpriseTeam, toAdd); err != nil {
		return diag.FromErr(err)
	}
	if _, err := removeEnterpriseTeamOrganizations(ctx, client, enterpriseSlug, enterpriseTeam, toRemove); err != nil {
		return diag.FromErr(err)
	}

	// NOTE: enterprise team slugs have the "ent:" prefix, so we must not use
	// colon-delimited IDs here.
	d.SetId(buildSlashTwoPartID(enterpriseSlug, enterpriseTeam))
	return resourceGithubEnterpriseTeamOrganizationsRead(context.WithValue(ctx, ctxId, d.Id()), d, meta)
}

func resourceGithubEnterpriseTeamOrganizationsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug, enterpriseTeam, err := parseSlashTwoPartID(d.Id(), "enterprise_slug", "enterprise_team")
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("enterprise_team", enterpriseTeam); err != nil {
		return diag.FromErr(err)
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())
	orgs, err := listEnterpriseTeamOrganizations(ctx, client, enterpriseSlug, enterpriseTeam)
	if err != nil {
		ghErr := &githubv3.ErrorResponse{}
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing enterprise team organizations %s from state because it no longer exists in GitHub", d.Id())
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	slugs := []string{}
	for _, org := range orgs {
		if org.Login != "" {
			slugs = append(slugs, org.Login)
		}
	}
	if err = d.Set("organization_slugs", slugs); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseTeamOrganizationsDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	enterpriseTeam := d.Get("enterprise_team").(string)

	ctx = context.WithValue(ctx, ctxId, d.Id())
	orgs, err := listEnterpriseTeamOrganizations(ctx, client, enterpriseSlug, enterpriseTeam)
	if err != nil {
		ghErr := &githubv3.ErrorResponse{}
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		return diag.FromErr(err)
	}

	toRemove := []string{}
	for _, org := range orgs {
		if org.Login != "" {
			toRemove = append(toRemove, org.Login)
		}
	}

	log.Printf("[INFO] Removing all organization assignments for enterprise team: %s/%s", enterpriseSlug, enterpriseTeam)
	_, err = removeEnterpriseTeamOrganizations(ctx, client, enterpriseSlug, enterpriseTeam, toRemove)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubEnterpriseTeamOrganizationsImport(_ context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	enterpriseSlug, enterpriseTeam, err := parseSlashTwoPartID(d.Id(), "enterprise_slug", "enterprise_team")
	if err != nil {
		return nil, fmt.Errorf("invalid import specified: supplied import must be written as <enterprise_slug>/<enterprise_team>")
	}
	d.SetId(buildSlashTwoPartID(enterpriseSlug, enterpriseTeam))
	_ = d.Set("enterprise_slug", enterpriseSlug)
	_ = d.Set("enterprise_team", enterpriseTeam)
	return []*schema.ResourceData{d}, nil
}
