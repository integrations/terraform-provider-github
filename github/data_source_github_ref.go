package github

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRef() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve information about a repository ref (branch or tag).",
		Read:        dataSourceGithubRefRead,

		Schema: map[string]*schema.Schema{
			"ref": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The repository ref to look up. Must be formatted 'heads/<ref>' for branches, and 'tags/<ref>' for tags.",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The GitHub repository name.",
			},
			"owner": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Owner of the repository.",
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An etag representing the ref.",
			},
			"sha": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The reference's HEAD commit's SHA1.",
			},
		},
	}
}

func dataSourceGithubRefRead(d *schema.ResourceData, meta any) error {
	ctx := context.Background()
	client := meta.(*Owner).v3client
	owner, ok := d.Get("owner").(string)
	if !ok {
		owner = meta.(*Owner).name
	}
	repoName := d.Get("repository").(string)
	ref := d.Get("ref").(string)

	refData, resp, err := client.Git.GetRef(ctx, owner, repoName, ref)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[DEBUG] Missing GitHub ref %s/%s (%s)", owner, repoName, ref)
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.SetId(buildTwoPartID(repoName, ref))
	err = d.Set("etag", resp.Header.Get("ETag"))
	if err != nil {
		return err
	}
	err = d.Set("sha", *refData.Object.SHA)
	if err != nil {
		return err
	}

	return nil
}
