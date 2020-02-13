package github

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/google/go-github/v28/github"
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
			"retrieve_by": {
				Type:     schema.TypeString,
				Required: true,
			},
			"release_tag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"release_id": {
				Type:     schema.TypeInt,
				Optional: true,
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

	client := meta.(*Organization).client
	ctx := context.Background()

	var err error
	var release *github.RepositoryRelease

	switch retrieveBy := strings.ToLower(d.Get("retrieve_by").(string)); retrieveBy {
	case "latest":
		log.Printf("[INFO] Refreshing GitHub latest release from repository %s", repository)
		release, _, err = client.Repositories.GetLatestRelease(ctx, owner, repository)
	case "id":
		releaseID := int64(d.Get("release_id").(int))
		if releaseID == 0 {
			return errors.New("'release_id' must be set when 'retrieve_by' = 'id'")
		}

		log.Printf("[INFO] Refreshing GitHub release by id %s from repository %s", releaseID, repository)
		release, _, err = client.Repositories.GetRelease(ctx, owner, repository, releaseID)
	case "tag":
		tag := d.Get("release_tag").(string)
		if tag == "" {
			return errors.New("'release_tag' must be set when 'retrieve_by' = 'tag'")
		}

		log.Printf("[INFO] Refreshing GitHub release by tag %s from repository %s", tag, repository)
		release, _, err = client.Repositories.GetReleaseByTag(ctx, owner, repository, tag)
	default:
		return errors.New("One of: 'latest', 'id', 'tag' must be set for 'retrieve_by'")
	}

	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(release.GetID(), 10))
	d.Set("url", release.GetURL())

	return nil
}
