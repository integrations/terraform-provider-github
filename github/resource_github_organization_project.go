package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubOrganizationProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubOrganizationProjectCreate,
		Read:   resourceGithubOrganizationProjectRead,
		Update: resourceGithubOrganizationProjectUpdate,
		Delete: resourceGithubOrganizationProjectDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
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

func resourceGithubOrganizationProjectCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	orgName := meta.(*Organization).name
	name := d.Get("name").(string)

	log.Printf("[DEBUG] Creating organization project: %s (%s)", name, orgName)
	project, _, err := client.Organizations.CreateProject(context.TODO(),
		orgName,
		&github.ProjectOptions{
			Name: name,
			Body: d.Get("body").(string),
		},
	)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(*project.ID, 10))

	return resourceGithubOrganizationProjectRead(d, meta)
}

func resourceGithubOrganizationProjectRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	orgName := meta.(*Organization).name

	projectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Reading organization project: %s (%s)", d.Id(), orgName)
	project, _, err := client.Projects.GetProject(context.TODO(), projectID)
	if err != nil {
		if err, ok := err.(*github.ErrorResponse); ok {
			if err.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing organization project %s/%s from state because it no longer exists in GitHub",
					orgName, d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("name", project.GetName())
	d.Set("body", project.GetBody())
	d.Set("url", fmt.Sprintf("https://github.com/orgs/%s/projects/%d",
		orgName, project.GetNumber()))

	return nil
}

func resourceGithubOrganizationProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	orgName := meta.(*Organization).name
	name := d.Get("name").(string)
	options := github.ProjectOptions{
		Name: name,
		Body: d.Get("body").(string),
	}

	projectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating organization project: %s (%s)", d.Id(), orgName)
	if _, _, err := client.Projects.UpdateProject(context.TODO(), projectID, &options); err != nil {
		return err
	}

	return resourceGithubOrganizationProjectRead(d, meta)
}

func resourceGithubOrganizationProjectDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	orgName := meta.(*Organization).name
	projectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting organization project: %s (%s)", d.Id(), orgName)
	_, err = client.Projects.DeleteProject(context.TODO(), projectID)
	return err
}
