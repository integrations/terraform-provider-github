package github

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/v45/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubTeamRepository() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubTeamRepositoryRead,

		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID or slug of team",
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"permission": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "pull",
			},
		},
	}
}

func dataSourceGithubTeamRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgId := meta.(*Owner).id

	// The given team id could be an id or a slug
	givenTeamId := d.Get("team_id").(string)
	teamId, err := getTeamID(givenTeamId, meta)
	if err != nil {
		return err
	}

	orgName := meta.(*Owner).name
	permission := d.Get("permission").(string)
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	var repoName string

	if name, ok := d.GetOk("name"); ok {
		repoName = name.(string)
	}

	if repoName == "" {
		return fmt.Errorf("%q has to be provided", "name")
	}

	repo, resp, repoErr := client.Teams.IsTeamRepoByID(ctx, orgId, teamId, orgName, repoName)
	if repoErr != nil {
		if ghErr, ok := repoErr.(*github.ErrorResponse); ok {
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
		return err
	}

	permName, permErr := getRepoPermission(repo.GetPermissions())
	if permErr != nil {
		return permErr
	}
	if permName == permission {
		d.Set("permission", permName)

		d.Set("etag", resp.Header.Get("ETag"))
		if d.Get("team_id") == "" {
			// If team_id is empty, that means we are importing the resource.
			// Set the team_id to be the id of the team.
			d.Set("team_id", teamId)
		}
		d.Set("repository", repo.GetName())

	}

	return nil
}
