package github

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubBranch() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve information about a branch in a repository.",
		Read:        dataSourceGithubBranchRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GitHub repository name.",
			},
			"branch": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The repository branch to retrieve.",
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An etag representing the branch object.",
			},
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A string representing a branch reference, in the form of refs/heads/<branch>.",
			},
			"sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SHA1 of the branch HEAD commit.",
			},
		},
	}
}

func dataSourceGithubBranchRead(d *schema.ResourceData, meta any) error {
	ctx := context.Background()
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	branchName := d.Get("branch").(string)
	branchRefName := "refs/heads/" + branchName

	ref, resp, err := client.Git.GetRef(ctx, orgName, repoName, branchRefName)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[DEBUG] Missing GitHub branch %s/%s (%s)", orgName, repoName, branchRefName)
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.SetId(buildTwoPartID(repoName, branchName))
	err = d.Set("etag", resp.Header.Get("ETag"))
	if err != nil {
		return err
	}
	err = d.Set("ref", ref.Ref)
	if err != nil {
		return err
	}
	err = d.Set("sha", ref.Object.SHA)
	if err != nil {
		return err
	}

	return nil
}
