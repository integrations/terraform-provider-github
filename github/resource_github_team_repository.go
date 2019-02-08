package github

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v21/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubTeamRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubTeamRepositoryCreate,
		Read:   resourceGithubTeamRepositoryRead,
		Update: resourceGithubTeamRepositoryUpdate,
		Delete: resourceGithubTeamRepositoryDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"permission": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "pull",
				ValidateFunc: validateValueFunc([]string{"pull", "push", "admin"}),
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubTeamRepositoryCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	teamIdString := d.Get("team_id").(string)
	teamId, err := strconv.ParseInt(teamIdString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(teamIdString, err)
	}
	orgName := meta.(*Organization).name
	repoName := d.Get("repository").(string)
	permission := d.Get("permission").(string)
	ctx := context.Background()

	log.Printf("[DEBUG] Creating team repository association: %s:%s (%s/%s)",
		teamIdString, permission, orgName, repoName)
	_, err = client.Teams.AddTeamRepo(ctx,
		teamId,
		orgName,
		repoName,
		&github.TeamAddTeamRepoOptions{
			Permission: permission,
		},
	)

	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(&teamIdString, &repoName))

	return resourceGithubTeamRepositoryRead(d, meta)
}

func resourceGithubTeamRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	teamIdString, repoName, err := parseTwoPartID(d.Id())
	if err != nil {
		return err
	}

	teamId, err := strconv.ParseInt(teamIdString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(teamIdString, err)
	}
	orgName := meta.(*Organization).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	log.Printf("[DEBUG] Reading team repository association: %s (%s/%s)",
		teamIdString, orgName, repoName)
	repo, resp, repoErr := client.Teams.IsTeamRepo(ctx,
		teamId, orgName, repoName)
	if repoErr != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing team repository association %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("team_id", teamIdString)
	d.Set("repository", repo.Name)

	permName, permErr := getRepoPermission(repo.Permissions)
	if permErr != nil {
		return permErr
	}

	d.Set("permission", permName)

	return nil
}

func resourceGithubTeamRepositoryUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	teamIdString := d.Get("team_id").(string)
	teamId, err := strconv.ParseInt(teamIdString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(teamIdString, err)
	}
	orgName := meta.(*Organization).name
	repoName := d.Get("repository").(string)
	permission := d.Get("permission").(string)
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Updating team repository association: %s:%s (%s/%s)",
		teamIdString, permission, orgName, repoName)
	// the go-github library's AddTeamRepo method uses the add/update endpoint from Github API
	_, err = client.Teams.AddTeamRepo(ctx,
		teamId,
		orgName,
		repoName,
		&github.TeamAddTeamRepoOptions{
			Permission: permission,
		},
	)

	if err != nil {
		return err
	}
	d.SetId(buildTwoPartID(&teamIdString, &repoName))

	return resourceGithubTeamRepositoryRead(d, meta)
}

func resourceGithubTeamRepositoryDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	teamIdString := d.Get("team_id").(string)

	teamId, err := strconv.ParseInt(teamIdString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(teamIdString, err)
	}
	orgName := meta.(*Organization).name
	repoName := d.Get("repository").(string)
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	log.Printf("[DEBUG] Deleting team repository association: %s (%s/%s)",
		teamIdString, orgName, repoName)
	_, err = client.Teams.RemoveTeamRepo(ctx,
		teamId, orgName, repoName)
	return err
}
