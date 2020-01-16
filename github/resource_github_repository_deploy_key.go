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
	client := meta.(*Owner).client

	repo := d.Get("repository").(string)

	k := d.Get("key").(string)
	t := d.Get("title").(string)
	r := d.Get("read_only").(bool)
	key := &github.Key{
		Key:      &k,
		Title:    &t,
		ReadOnly: &r,
	}

	owner := meta.(*Owner).name
	resultKey, _, err := client.Repositories.CreateKey(context.TODO(), owner, repo, key)

	if err != nil {
		return err
	}

	i := strconv.FormatInt(*resultKey.ID, 10)
	id := buildTwoPartID(&repo, &i)

	d.SetId(id)

	return resourceGithubRepositoryDeployKeyRead(d, meta)
}

func resourceGithubRepositoryDeployKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).client

	owner := meta.(*Owner).name
	repo, id := parseTwoPartID(d.Id())

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	key, _, err := client.Repositories.GetKey(context.TODO(), owner, repo, i)
	if err != nil {
		return err
	}

	d.Set("key", key.Key)
	d.Set("read_only", key.ReadOnly)
	d.Set("repository", repo)
	d.Set("title", key.Title)

	return nil
}

func resourceGithubRepositoryDeployKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).client

	owner := meta.(*Owner).name
	repo, id := parseTwoPartID(d.Id())

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	_, err = client.Repositories.DeleteKey(context.TODO(), owner, repo, i)
	if err != nil {
		return err
	}

	return err
}
