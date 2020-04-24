package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v29/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateTeamIDFunc,
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
				ValidateFunc: validateValueFunc([]string{"pull", "triage", "push", "maintain", "admin"}),
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubTeamRepositoryCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client
	orgId := meta.(*Organization).id

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

	opts := &github.TeamAddTeamRepoOptions{
		Permission: permission,
	}
	if meta.(*Organization).isEnterprise {
		_, err = AddEnterpriseTeamRepoByID(ctx, client, teamId, orgName, repoName, opts)
	} else {
		_, err = client.Teams.AddTeamRepoByID(ctx,
			orgId,
			teamId,
			orgName,
			repoName,
			opts,
		)
	}
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(teamIdString, repoName))

	return resourceGithubTeamRepositoryRead(d, meta)
}

func resourceGithubTeamRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client
	orgId := meta.(*Organization).id

	teamIdString, repoName, err := parseTwoPartID(d.Id(), "team_id", "repository")
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

	log.Printf("[DEBUG] Reading team repository association: %s (%s/%s)", teamIdString, orgName, repoName)
	var repo *github.Repository
	var resp *github.Response
	var repoErr error
	if meta.(*Organization).isEnterprise {
		repo, resp, repoErr = IsEnterpriseTeamRepoByID(ctx, client, teamId, orgName, repoName)
	} else {
		repo, resp, repoErr = client.Teams.IsTeamRepoByID(ctx, orgId, teamId, orgName, repoName)
	}
	if repoErr != nil {
		if ghErr, ok := repoErr.(*github.ErrorResponse); ok {
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
	d.Set("repository", repo.GetName())

	permName, permErr := getRepoPermission(repo.GetPermissions())
	if permErr != nil {
		return permErr
	}

	d.Set("permission", permName)

	return nil
}

func resourceGithubTeamRepositoryUpdate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client
	orgId := meta.(*Organization).id

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
	opts := &github.TeamAddTeamRepoOptions{
		Permission: permission,
	}
	if meta.(*Organization).isEnterprise {
		_, err = AddEnterpriseTeamRepoByID(ctx, client, teamId, orgName, repoName, opts)
	} else {
		_, err = client.Teams.AddTeamRepoByID(ctx,
			orgId,
			teamId,
			orgName,
			repoName,
			opts,
		)
	}
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(teamIdString, repoName))

	return resourceGithubTeamRepositoryRead(d, meta)
}

func resourceGithubTeamRepositoryDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).v3client
	orgId := meta.(*Organization).id

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
	if meta.(*Organization).isEnterprise {
		_, err = RemoveEnterpriseTeamRepoByID(ctx, client, teamId, orgName, repoName)
	} else {
		_, err = client.Teams.RemoveTeamRepoByID(ctx, orgId, teamId, orgName, repoName)
	}
	return err
}

// API functionality below is no longer available in go-github v29.0.3+.
// Naming conventions reflect Enterprise Github Account support.
// Code taken from go-github v29.0.2 as a temporary work-around to [GH-404] and [GH-434].
func IsEnterpriseTeamRepoByID(ctx context.Context, client *github.Client, teamId int64, owner, repo string) (*github.Repository, *github.Response, error) {
	u := fmt.Sprintf("teams/%v/repos/%v/%v", teamId, owner, repo)
	req, err := client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	mediaTypeOrgPermissionRepo := "application/vnd.github.v3.repository+json"
	headers := []string{mediaTypeOrgPermissionRepo}
	req.Header.Set("Accept", strings.Join(headers, ", "))

	repository := new(github.Repository)
	resp, err := client.Do(ctx, req, repository)
	if err != nil {
		return nil, resp, err
	}

	return repository, resp, nil
}

func AddEnterpriseTeamRepoByID(ctx context.Context, client *github.Client, teamId int64, owner, repo string, opts *github.TeamAddTeamRepoOptions) (*github.Response, error) {
	u := fmt.Sprintf("teams/%v/repos/%v/%v", teamId, owner, repo)
	req, err := client.NewRequest("PUT", u, opts)

	if err != nil {
		return nil, err
	}

	return client.Do(ctx, req, nil)

}

func RemoveEnterpriseTeamRepoByID(ctx context.Context, client *github.Client, teamId int64, owner, repo string) (*github.Response, error) {
	u := fmt.Sprintf("teams/%v/repos/%v/%v", teamId, owner, repo)
	req, err := client.NewRequest("DELETE", u, nil)

	if err != nil {
		return nil, err
	}

	return client.Do(ctx, req, nil)
}
