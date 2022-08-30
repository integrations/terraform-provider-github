package github

import (
	"context"

	"github.com/google/go-github/v47/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubBranchDefault() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubBranchDefaultCreate,
		Read:   resourceGithubBranchDefaultRead,
		Delete: resourceGithubBranchDefaultDelete,
		Update: resourceGithubBranchDefaultUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"branch": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGithubBranchDefaultCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	defaultBranch := d.Get("branch").(string)

	repository := &github.Repository{
		DefaultBranch: &defaultBranch,
	}

	ctx := context.Background()

	if _, _, err := client.Repositories.Edit(ctx, owner, repoName, repository); err != nil {
		return err
	}

	d.SetId(repoName)

	return resourceGithubBranchDefaultRead(d, meta)
}

func resourceGithubBranchDefaultRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Id()

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	repository, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return err
	}

	if repository.DefaultBranch == nil {
		d.SetId("")
		return nil
	}

	d.Set("branch", *repository.DefaultBranch)
	d.Set("repository", *repository.Name)
	return nil
}

func resourceGithubBranchDefaultDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Id()

	repository := &github.Repository{
		DefaultBranch: nil,
	}

	ctx := context.Background()

	_, _, err := client.Repositories.Edit(ctx, owner, repoName, repository)
	return err
}

func resourceGithubBranchDefaultUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Id()
	defaultBranch := d.Get("branch").(string)

	repository := &github.Repository{
		DefaultBranch: &defaultBranch,
	}

	ctx := context.Background()

	if _, _, err := client.Repositories.Edit(ctx, owner, repoName, repository); err != nil {
		return err
	}

	return resourceGithubBranchDefaultRead(d, meta)
}
