package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v77/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubOrganizationProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubOrganizationProjectCreate,
		Read:   resourceGithubOrganizationProjectRead,
		Update: resourceGithubOrganizationProjectUpdate,
		Delete: resourceGithubOrganizationProjectDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The title of the project.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the project.",
			},
			"public": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether the project should be public or private.",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the project.",
			},
			"number": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of the project.",
			},
			"node_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The node ID of the project.",
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubOrganizationProjectCreate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	// Projects V2 API doesn't support creating projects via REST API
	// Projects must be created through the GitHub web interface
	// This resource can only import and manage existing projects
	return fmt.Errorf("Projects V2 cannot be created via the REST API. Please create the project through the GitHub web interface and import it using terraform import")
}

func resourceGithubOrganizationProjectRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	projectNumber, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	project, resp, err := client.Projects.GetOrganizationProject(ctx, orgName, projectNumber)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing organization project %s/%s from state because it no longer exists in GitHub",
					orgName, d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return err
	}
	if err = d.Set("title", project.GetTitle()); err != nil {
		return err
	}
	if err = d.Set("description", project.GetDescription()); err != nil {
		return err
	}
	if err = d.Set("public", project.GetPublic()); err != nil {
		return err
	}
	if err = d.Set("number", project.GetNumber()); err != nil {
		return err
	}
	if err = d.Set("node_id", project.GetNodeID()); err != nil {
		return err
	}
	if err = d.Set("url", fmt.Sprintf("https://github.com/orgs/%s/projects/%d",
		orgName, project.GetNumber())); err != nil {
		return err
	}

	return nil
}

func resourceGithubOrganizationProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	// Projects V2 API doesn't support updating projects via REST API
	// Projects must be updated through the GitHub web interface
	return fmt.Errorf("Projects V2 cannot be updated via the REST API. Please update the project through the GitHub web interface")
}

func resourceGithubOrganizationProjectDelete(d *schema.ResourceData, meta interface{}) error {
	// Projects V2 API doesn't support deleting projects via REST API
	// Projects must be deleted through the GitHub web interface
	return fmt.Errorf("Projects V2 cannot be deleted via the REST API. Please delete the project through the GitHub web interface")
}
