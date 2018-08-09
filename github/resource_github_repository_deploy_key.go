package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubRepositoryDeployKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryDeployKeyCreate,
		Read:   resourceGithubRepositoryDeployKeyRead,
		// Deploy keys are defined immutable in the API. Updating results in force new.
		Delete: resourceGithubRepositoryDeployKeyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"read_only": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  true,
			},
			"repository": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGithubRepositoryDeployKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	repoName := d.Get("repository").(string)
	key := d.Get("key").(string)
	title := d.Get("title").(string)
	readOnly := d.Get("read_only").(bool)

	owner := meta.(*Organization).name
	resultKey, _, err := client.Repositories.CreateKey(context.TODO(), owner, repoName, &github.Key{
		Key:      &key,
		Title:    &title,
		ReadOnly: &readOnly,
	})

	if err != nil {
		return err
	}

	id := strconv.FormatInt(*resultKey.ID, 10)

	d.SetId(buildTwoPartID(&repoName, &id))

	return resourceGithubRepositoryDeployKeyRead(d, meta)
}

func resourceGithubRepositoryDeployKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	owner := meta.(*Organization).name
	repoName, idString, err := parseTwoPartID(d.Id())
	if err != nil {
		return err
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(idString, err)
	}

	key, _, err := client.Repositories.GetKey(context.TODO(), owner, repoName, id)
	if err != nil {
		return err
	}

	d.Set("key", key.Key)
	d.Set("read_only", key.ReadOnly)
	d.Set("repository", repoName)
	d.Set("title", key.Title)

	return nil
}

func resourceGithubRepositoryDeployKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	owner := meta.(*Organization).name
	repoName, idString, err := parseTwoPartID(d.Id())
	if err != nil {
		return err
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(idString, err)
	}

	_, err = client.Repositories.DeleteKey(context.TODO(), owner, repoName, id)
	if err != nil {
		return err
	}

	return err
}
