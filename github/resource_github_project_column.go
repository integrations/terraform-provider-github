package github

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v47/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
			"column_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubProjectColumnCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client

	options := github.ProjectColumnOptions{
		Name: d.Get("name").(string),
	}

	projectIDStr := d.Get("project_id").(string)
	projectID, err := strconv.ParseInt(projectIDStr, 10, 64)
	if err != nil {
		return unconvertibleIdErr(projectIDStr, err)
	}
	ctx := context.Background()

	column, _, err := client.Projects.CreateProjectColumn(ctx,
		projectID,
		&options,
	)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(column.GetID(), 10))
	d.Set("column_id", column.GetID())

	return resourceGithubProjectColumnRead(d, meta)
}

func resourceGithubProjectColumnRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	columnID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	column, _, err := client.Projects.GetProjectColumn(ctx, columnID)
	if err != nil {
		if err, ok := err.(*github.ErrorResponse); ok {
			if err.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing project column %s from state because it no longer exists in GitHub", d.Id())
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
	d.Set("column_id", column.GetID())
	return nil
}

func resourceGithubProjectColumnUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	options := github.ProjectColumnOptions{
		Name: d.Get("name").(string),
	}

	columnID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	_, _, err = client.Projects.UpdateProjectColumn(ctx, columnID, &options)
	if err != nil {
		return err
	}

	return resourceGithubProjectColumnRead(d, meta)
}

func resourceGithubProjectColumnDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	columnID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	_, err = client.Projects.DeleteProjectColumn(ctx, columnID)
	return err
}
