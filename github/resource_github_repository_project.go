package github

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubRepositoryProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryProjectCreate,
		Read:   resourceGithubRepositoryProjectRead,
		Update: resourceGithubRepositoryProjectUpdate,
		Delete: resourceGithubRepositoryProjectDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"body": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubRepositoryProjectCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	o := meta.(*Organization).name
	n := d.Get("name").(string)
	b := d.Get("body").(string)
	r := d.Get("repository").(string)

	options := github.ProjectOptions{
		Name: n,
		Body: b,
	}

	project, _, err := client.Repositories.CreateProject(context.TODO(), o, r, &options)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(*project.ID, 10))

	return resourceGithubRepositoryProjectRead(d, meta)
}

func resourceGithubRepositoryProjectRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	o := meta.(*Organization).name

	projectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	project, resp, err := client.Projects.GetProject(context.TODO(), projectID)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", project.GetName())
	d.Set("body", project.GetBody())
	d.Set("url", fmt.Sprintf("https://github.com/%s/%s/projects/%d", o, d.Get("repository"), project.GetNumber()))

	return nil
}

func resourceGithubRepositoryProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	n := d.Get("name").(string)
	b := d.Get("body").(string)

	options := github.ProjectOptions{
		Name: n,
		Body: b,
	}

	projectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	if _, _, err := client.Projects.UpdateProject(context.TODO(), projectID, &options); err != nil {
		return err
	}

	return resourceGithubRepositoryProjectRead(d, meta)
}

func resourceGithubRepositoryProjectDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	projectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	_, err = client.Projects.DeleteProject(context.TODO(), projectID)
	return err
}
