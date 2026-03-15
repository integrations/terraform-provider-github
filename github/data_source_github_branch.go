package github

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubBranch() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubBranchRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"branch": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sha": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubBranchRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId(buildTwoPartID(repoName, branchName))
	err = d.Set("etag", resp.Header.Get("ETag"))
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("ref", *ref.Ref)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("sha", *ref.Object.SHA)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
