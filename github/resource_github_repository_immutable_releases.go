package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryImmutableReleases() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryImmutableReleasesCreateOrUpdate,
		Read:   resourceGithubRepositoryImmutableReleasesRead,
		Update: resourceGithubRepositoryImmutableReleasesCreateOrUpdate,
		Delete: resourceGithubRepositoryImmutableReleasesDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGithubRepositoryImmutableReleasesImport,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the repository.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether immutable releases are enabled for the repository.",
			},
		},
	}
}

func resourceGithubRepositoryImmutableReleasesCreateOrUpdate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	enabled := d.Get("enabled").(bool)

	ctx := context.Background()
	var err error
	if enabled {
		_, err = client.Repositories.EnableImmutableReleases(ctx, owner, repoName)
	} else {
		_, err = client.Repositories.DisableImmutableReleases(ctx, owner, repoName)
	}

	if err != nil {
		return err
	}

	d.SetId(repoName)
	return resourceGithubRepositoryImmutableReleasesRead(d, meta)
}

func resourceGithubRepositoryImmutableReleasesRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	ctx := context.Background()

	status, _, err := client.Repositories.AreImmutableReleasesEnabled(ctx, owner, repoName)
	if err != nil {
		return err
	}

	enabled := status.GetEnabled()
	if err := d.Set("enabled", enabled); err != nil {
		return err
	}

	if d.Id() == "" {
		d.SetId(repoName)
	}

	return nil
}

func resourceGithubRepositoryImmutableReleasesDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	ctx := context.Background()

	_, err := client.Repositories.DisableImmutableReleases(ctx, owner, repoName)
	if err != nil {
		return err
	}

	return resourceGithubRepositoryImmutableReleasesRead(d, meta)
}

func resourceGithubRepositoryImmutableReleasesImport(d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	repoName := d.Id()

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}

	if err := resourceGithubRepositoryImmutableReleasesRead(d, meta); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
