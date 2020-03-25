package github

import (
	"context"
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
				Optional: true,
				ForceNew: true,
				Default:  "master",
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

	client := meta.(*Organization).client
	orgName := meta.(*Organization).name
	repoName := d.Get("repository").(string)
	branchName := d.Get("branch").(string)

	log.Printf("[DEBUG] Reading repository branch: %s/%s (%s)",
		orgName, repoName, branchName)
	ref, _, err := client.Git.GetRef(
		context.TODO(), orgName, repoName, "refs/heads/"+branchName)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(&repoName, &branchName))
	d.Set("ref", *ref.Ref)
	d.Set("sha", *ref.Object.SHA)

	return nil
}
