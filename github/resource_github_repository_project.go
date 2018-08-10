package github

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				parts := strings.Split(d.Id(), "/")
				if len(parts) != 2 {
					return nil, fmt.Errorf("Invalid ID specified. Supplied ID must be written as <repository>/<project_id>")
				}
				d.Set("repository", parts[0])
				d.SetId(parts[1])
				return []*schema.ResourceData{d}, nil
			},
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

	options := github.ProjectOptions{
		Name: d.Get("name").(string),
		Body: d.Get("body").(string),
	}

	project, _, err := client.Repositories.CreateProject(context.TODO(),
		meta.(*Organization).name,
		d.Get("repository").(string),
		&options,
	)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(*project.ID, 10))

	return resourceGithubRepositoryProjectRead(d, meta)
}

func resourceGithubRepositoryProjectRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	orgName := meta.(*Organization).name

	projectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
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
	d.Set("url", fmt.Sprintf("https://github.com/%s/%s/projects/%d",
		orgName, d.Get("repository"), project.GetNumber()))

	return nil
}

func resourceGithubRepositoryProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	options := github.ProjectOptions{
		Name: d.Get("name").(string),
		Body: d.Get("body").(string),
	}

	projectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	_, _, err = client.Projects.UpdateProject(context.TODO(), projectID, &options)
	if err != nil {
		return err
	}

	return resourceGithubRepositoryProjectRead(d, meta)
}

func resourceGithubRepositoryProjectDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	projectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	_, err = client.Projects.DeleteProject(context.TODO(), projectID)
	return err
}
