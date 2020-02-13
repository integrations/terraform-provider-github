package github

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceGithubRelease() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubReleaseRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Required: true,
			},
			"release_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubReleaseRead(d *schema.ResourceData, meta interface{}) error {
	repository := d.Get("repository").(string)
	owner := d.Get("owner").(string)
	releaseID := int64(d.Get("release_id").(int))
	log.Printf("[INFO] Refreshing GitHub release %s from repository %s", releaseID, repository)

	client := meta.(*Organization).client
	ctx := context.Background()

	release, _, err := client.Repositories.GetRelease(ctx, owner, repository, releaseID)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(release.GetID(), 10))
	d.Set("url", release.GetURL())

	return nil
}
