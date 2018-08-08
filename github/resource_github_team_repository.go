package github

import (
	"context"
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
	client := meta.(*Organization).client

	teamIdString := d.Get("team_id").(string)
	teamId, err := strconv.ParseInt(teamIdString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(teamIdString, err)
	}
	repoName := d.Get("repository").(string)

	_, err = client.Organizations.AddTeamRepo(context.TODO(),
		teamId,
		meta.(*Organization).name,
		repoName,
		&github.OrganizationAddTeamRepoOptions{
			Permission: d.Get("permission").(string),
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

	repo, _, repoErr := client.Organizations.IsTeamRepo(context.TODO(),
		teamId,
		meta.(*Organization).name,
		repoName,
	)

	if repoErr != nil {
		d.SetId("")
		return nil
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
	client := meta.(*Organization).client

	teamIdString := d.Get("team_id").(string)
	teamId, err := strconv.ParseInt(teamIdString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(teamIdString, err)
	}

	repoName := d.Get("repository").(string)

	// the go-github library's AddTeamRepo method uses the add/update endpoint from Github API
	_, err = client.Organizations.AddTeamRepo(context.TODO(),
		teamId,
		meta.(*Organization).name,
		repoName,
		&github.OrganizationAddTeamRepoOptions{
			Permission: d.Get("permission").(string),
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

	_, err = client.Organizations.RemoveTeamRepo(context.TODO(),
		teamId,
		meta.(*Organization).name,
		d.Get("repository").(string),
	)

	return err
}
