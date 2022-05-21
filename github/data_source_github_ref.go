package github

import (
	"context"
	"log"
	"net/http"

	"github.com/google/go-github/v44/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubRef() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRefRead,

		Schema: map[string]*schema.Schema{
			"ref": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"etag": {
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

func dataSourceGithubRefRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	ref := d.Get("ref").(string)

	refData, resp, err := client.Git.GetRef(context.TODO(), orgName, repoName, ref)
	if err != nil {
		if err, ok := err.(*github.ErrorResponse); ok {
			if err.Response.StatusCode == http.StatusNotFound {
				log.Printf("[DEBUG] Missing GitHub ref %s/%s (%s)", orgName, repoName, ref)
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.SetId(buildTwoPartID(repoName, ref))
	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("sha", *refData.Object.SHA)

	return nil
}
