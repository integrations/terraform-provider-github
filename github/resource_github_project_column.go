package github

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubProjectColumn() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubProjectColumnCreate,
		Read:   resourceGithubProjectColumnRead,
		Update: resourceGithubProjectColumnUpdate,
		Delete: resourceGithubProjectColumnDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGithubProjectColumnCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	options := github.ProjectColumnOptions{
		Name: d.Get("name").(string),
	}

	projectIDStr := d.Get("project_id").(string)
	projectID, err := strconv.ParseInt(projectIDStr, 10, 64)
	if err != nil {
		return unconvertibleIdErr(projectIDStr, err)
	}

	orgName := meta.(*Organization).name
	log.Printf("[DEBUG] Creating project column (%s) in project %d (%s)", options.Name, projectID, orgName)
	column, _, err := client.Projects.CreateProjectColumn(context.TODO(),
		projectID,
		&options,
	)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(*column.ID, 10))

	return resourceGithubProjectColumnRead(d, meta)
}

func resourceGithubProjectColumnRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	columnID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	log.Printf("[DEBUG] Reading project column: %s", d.Id())
	column, _, err := client.Projects.GetProjectColumn(context.TODO(), columnID)
	if err != nil {
		if err, ok := err.(*github.ErrorResponse); ok {
			if err.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing project column %s from state because it no longer exists in GitHub", d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	projectURL := column.GetProjectURL()
	projectID := strings.TrimPrefix(projectURL, client.BaseURL.String()+`projects/`)

	d.Set("name", column.GetName())
	d.Set("project_id", projectID)
	return nil
}

func resourceGithubProjectColumnUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	options := github.ProjectColumnOptions{
		Name: d.Get("name").(string),
	}

	columnID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	log.Printf("[DEBUG] Updating project column: %s", d.Id())
	_, _, err = client.Projects.UpdateProjectColumn(context.TODO(), columnID, &options)
	if err != nil {
		return err
	}

	return resourceGithubProjectColumnRead(d, meta)
}

func resourceGithubProjectColumnDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	columnID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting project column: %s", d.Id())
	_, err = client.Projects.DeleteProjectColumn(context.TODO(), columnID)
	return err
}
