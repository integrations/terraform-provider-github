package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubTeamRepository() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubTeamRepositoryCreate,
		ReadContext:   resourceGithubTeamRepositoryRead,
		UpdateContext: resourceGithubTeamRepositoryUpdate,
		DeleteContext: resourceGithubTeamRepositoryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
				meta := m.(*Owner)
				teamIdString, username, err := parseTwoPartID(d.Id(), "team_id", "username")
				if err != nil {
					return nil, err
				}

				teamId, err := getTeamID(ctx, meta, teamIdString)
				if err != nil {
					return nil, err
				}

				d.SetId(buildTwoPartID(strconv.FormatInt(teamId, 10), username))
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID or slug of team",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The repository to add to the team.",
			},
			"permission": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "pull",
				Description: "The permissions of team members regarding the repository. Must be one of 'pull', 'triage', 'push', 'maintain', 'admin' or the name of an existing custom repository role within the organisation.",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubTeamRepositoryCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	orgId := meta.id
	orgName := meta.name

	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	// The given team id could be an id or a slug
	givenTeamId := d.Get("team_id").(string)
	teamId, err := getTeamID(ctx, meta, givenTeamId)
	if err != nil {
		return diag.FromErr(err)
	}

	repoName := d.Get("repository").(string)
	permission := d.Get("permission").(string)

	_, err = client.Teams.AddTeamRepoByID(ctx,
		orgId,
		teamId,
		orgName,
		repoName,
		&github.TeamAddTeamRepoOptions{
			Permission: permission,
		},
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildTwoPartID(strconv.FormatInt(teamId, 10), repoName))

	return resourceGithubTeamRepositoryRead(ctx, d, meta)
}

func resourceGithubTeamRepositoryRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	orgId := meta.id
	orgName := meta.name

	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	teamIdString, repoName, err := parseTwoPartID(d.Id(), "team_id", "repository")
	if err != nil {
		return diag.FromErr(err)
	}
	teamId, err := getTeamID(ctx, meta, teamIdString)
	if err != nil {
		return diag.FromErr(err)
	}

	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	repo, resp, repoErr := client.Teams.IsTeamRepoByID(ctx, orgId, teamId, orgName, repoName)
	if repoErr != nil {
		var ghErr *github.ErrorResponse
		if errors.As(repoErr, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing team repository association %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(repoErr)
	}

	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if d.Get("team_id") == "" {
		// If team_id is empty, that means we are importing the resource.
		// Set the team_id to be the id of the team.
		if err = d.Set("team_id", teamIdString); err != nil {
			return diag.FromErr(err)
		}
	}
	if err = d.Set("repository", repo.GetName()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("permission", getPermission(repo.GetRoleName())); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubTeamRepositoryUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgId := meta.(*Owner).id

	teamIdString, repoName, err := parseTwoPartID(d.Id(), "team_id", "repository")
	if err != nil {
		return diag.FromErr(err)
	}
	teamId, err := strconv.ParseInt(teamIdString, 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(teamIdString, err))
	}
	orgName := meta.(*Owner).name
	permission := d.Get("permission").(string)

	// the go-github library's AddTeamRepo method uses the add/update endpoint from GitHub API
	_, err = client.Teams.AddTeamRepoByID(ctx,
		orgId,
		teamId,
		orgName,
		repoName,
		&github.TeamAddTeamRepoOptions{
			Permission: permission,
		},
	)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(buildTwoPartID(teamIdString, repoName))

	return resourceGithubTeamRepositoryRead(ctx, d, meta)
}

func resourceGithubTeamRepositoryDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgId := meta.(*Owner).id

	teamIdString, repoName, err := parseTwoPartID(d.Id(), "team_id", "repository")
	if err != nil {
		return diag.FromErr(err)
	}
	teamId, err := strconv.ParseInt(teamIdString, 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(teamIdString, err))
	}
	orgName := meta.(*Owner).name

	resp, err := client.Teams.RemoveTeamRepoByID(ctx, orgId, teamId, orgName, repoName)

	if resp.StatusCode == 404 {
		log.Printf("[DEBUG] Failed to find team %s to delete for repo: %s.", teamIdString, repoName)
		repo, _, err := client.Repositories.Get(ctx, orgName, repoName)
		if err != nil {
			return diag.FromErr(err)
		}
		newRepoName := repo.GetName()
		if newRepoName != repoName {
			log.Printf("[INFO] Repo name has changed %s -> %s. "+
				"Try deleting team repository again.",
				repoName, newRepoName)
			_, err := client.Teams.RemoveTeamRepoByID(ctx, orgId, teamId, orgName, newRepoName)
			return diag.FromErr(handleArchivedRepoDelete(err, "team repository access", fmt.Sprintf("team %s", teamIdString), orgName, newRepoName))
		}
	}

	return diag.FromErr(handleArchivedRepoDelete(err, "team repository access", fmt.Sprintf("team %s", teamIdString), orgName, repoName))
}
