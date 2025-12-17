package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	githubv3 "github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubEnterpriseTeamMembership() *schema.Resource {
	return &schema.Resource{
		Create:   resourceGithubEnterpriseTeamMembershipCreate,
		Read:     resourceGithubEnterpriseTeamMembershipRead,
		Delete:   resourceGithubEnterpriseTeamMembershipDelete,
		Importer: &schema.ResourceImporter{State: resourceGithubEnterpriseTeamMembershipImport},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise.",
			},
			"enterprise_team": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug or ID of the enterprise team.",
			},
			"username": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: caseInsensitive(),
				Description:      "The login handle of the user.",
			},
		},
	}
}

func resourceGithubEnterpriseTeamMembershipCreate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	enterpriseTeam := d.Get("enterprise_team").(string)
	username := d.Get("username").(string)

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	// The API is idempotent, so we don't need to check if they're already a member
	_, err := addEnterpriseTeamMember(ctx, client, enterpriseSlug, enterpriseTeam, username)
	if err != nil {
		return err
	}

	// NOTE: enterprise team slugs have the "ent:" prefix, so we must not use
	// colon-delimited IDs here.
	d.SetId(buildSlashThreePartID(enterpriseSlug, enterpriseTeam, username))
	return resourceGithubEnterpriseTeamMembershipRead(d, meta)
}

func resourceGithubEnterpriseTeamMembershipRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	enterpriseSlug, enterpriseTeam, username, err := parseSlashThreePartID(d.Id(), "enterprise_slug", "enterprise_team", "username")
	if err != nil {
		return err
	}

	if err = d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return err
	}
	if err = d.Set("enterprise_team", enterpriseTeam); err != nil {
		return err
	}
	if err = d.Set("username", username); err != nil {
		return err
	}

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	_, err = getEnterpriseTeamMembership(ctx, client, enterpriseSlug, enterpriseTeam, username)
	if err != nil {
		ghErr := &githubv3.ErrorResponse{}
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing enterprise team membership %s from state because it no longer exists in GitHub", d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	return nil
}

func resourceGithubEnterpriseTeamMembershipDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	enterpriseTeam := d.Get("enterprise_team").(string)
	username := d.Get("username").(string)

	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	resp, err := removeEnterpriseTeamMember(ctx, client, enterpriseSlug, enterpriseTeam, username)
	if err != nil {
		ghErr := &githubv3.ErrorResponse{}
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		_ = resp
		return err
	}

	return nil
}

func resourceGithubEnterpriseTeamMembershipImport(d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	enterpriseSlug, enterpriseTeam, username, err := parseSlashThreePartID(d.Id(), "enterprise_slug", "enterprise_team", "username")
	if err != nil {
		return nil, fmt.Errorf("invalid import specified: supplied import must be written as <enterprise_slug>/<enterprise_team>/<username>")
	}
	d.SetId(buildSlashThreePartID(enterpriseSlug, enterpriseTeam, username))
	_ = d.Set("enterprise_slug", enterpriseSlug)
	_ = d.Set("enterprise_team", enterpriseTeam)
	_ = d.Set("username", username)
	return []*schema.ResourceData{d}, nil
}
