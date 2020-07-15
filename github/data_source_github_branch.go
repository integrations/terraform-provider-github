package github

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubBranch() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubBranchRead,

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

func dataSourceGithubBranchRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	branchName := d.Get("branch").(string)
	branchRefName := "refs/heads/" + branchName

	log.Printf("[DEBUG] Reading GitHub branch reference %s/%s (%s)",
		orgName, repoName, branchRefName)
	ref, resp, err := client.Git.GetRef(
		context.TODO(), orgName, repoName, branchRefName)
	if err != nil {
		return fmt.Errorf("Error reading GitHub branch reference %s/%s (%s): %s",
			orgName, repoName, branchRefName, err)
	}

	d.SetId(buildTwoPartID(repoName, branchName))
	d.Set("etag", resp.Header.Get("ETag"))
	d.Set("ref", *ref.Ref)
	d.Set("sha", *ref.Object.SHA)

	return nil
}
