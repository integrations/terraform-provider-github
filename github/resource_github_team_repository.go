package github

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/github"
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
		},
	}
}

func resourceGithubTeamRepositoryCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).client

	teamIdString := d.Get("team_id").(string)
	teamId, err := strconv.ParseInt(teamIdString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(teamIdString, err)
	}
	orgName := meta.(*Organization).name
	repoName := d.Get("repository").(string)
	permission := d.Get("permission").(string)

	log.Printf("[DEBUG] Creating team repository association: %s:%s (%s/%s)",
		teamIdString, permission, orgName, repoName)
	_, err = client.Organizations.AddTeamRepo(context.TODO(),
		teamId,
		orgName,
		repoName,
		&github.OrganizationAddTeamRepoOptions{
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
	client := meta.(*Owner).client

	teamIdString, repoName, err := parseTwoPartID(d.Id())
	if err != nil {
		return err
	}

	teamId, err := strconv.ParseInt(teamIdString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(teamIdString, err)
	}
	orgName := meta.(*Organization).name

	log.Printf("[DEBUG] Reading team repository association: %s (%s/%s)",
		teamIdString, orgName, repoName)
	repo, _, repoErr := client.Organizations.IsTeamRepo(context.TODO(),
		teamId, orgName, repoName)
	if repoErr != nil {
		if err, ok := err.(*github.ErrorResponse); ok {
			if err.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing team repository association %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

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
	client := meta.(*Owner).client

	teamIdString := d.Get("team_id").(string)
	teamId, err := strconv.ParseInt(teamIdString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(teamIdString, err)
	}
	orgName := meta.(*Organization).name
	repoName := d.Get("repository").(string)
	permission := d.Get("permission").(string)

	log.Printf("[DEBUG] Updating team repository association: %s:%s (%s/%s)",
		teamIdString, permission, orgName, repoName)
	// the go-github library's AddTeamRepo method uses the add/update endpoint from Github API
	_, err = client.Organizations.AddTeamRepo(context.TODO(),
		teamId,
		orgName,
		repoName,
		&github.OrganizationAddTeamRepoOptions{
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
	client := meta.(*Owner).client

	teamIdString := d.Get("team_id").(string)

	teamId, err := strconv.ParseInt(teamIdString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(teamIdString, err)
	}
	orgName := meta.(*Organization).name
	repoName := d.Get("repository").(string)

	log.Printf("[DEBUG] Deleting team repository association: %s (%s/%s)",
		teamIdString, orgName, repoName)
	_, err = client.Organizations.RemoveTeamRepo(context.TODO(),
		teamId, orgName, repoName)
	return err
}
