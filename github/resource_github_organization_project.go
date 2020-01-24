package github

import (
	"fmt"
	"log"
	"strconv"

	"github.com/google/go-github/v28/github"
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
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubOrganizationProjectCreate(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client
	name := d.Get("name").(string)
	body := d.Get("body").(string)

	ctx := prepareResourceContext(d)

	log.Printf("[DEBUG] Creating organization project: %s (%s)", name, orgName)
	project, _, err := client.Organizations.CreateProject(ctx,
		orgName,
		&github.ProjectOptions{
			Name: &name,
			Body: &body,
		},
	)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(*project.ID, 10))

	return resourceGithubOrganizationProjectRead(d, meta)
}

func resourceGithubOrganizationProjectRead(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client

	projectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	ctx := prepareResourceContext(d)

	log.Printf("[DEBUG] Reading organization project: %s (%s)", d.Id(), orgName)
	project, resp, err := client.Projects.GetProject(ctx, projectID)
	switch apires, apierr := apiResult(resp, err); apires {
	case APINotModified:
		return nil
	case APINotFound:
		log.Printf("[WARN] Removing organization project %s/%s from state because it no longer exists in GitHub", orgName, d.Id())
		d.SetId("")
		return nil
	case APIError:
		return apierr
	default:
		d.Set("etag", resp.Header.Get("ETag"))
		d.Set("name", project.GetName())
		d.Set("body", project.GetBody())
		d.Set("url", fmt.Sprintf("https://github.com/orgs/%s/projects/%d", orgName, project.GetNumber()))

		return nil
	}
}

func resourceGithubOrganizationProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client

	name := d.Get("name").(string)
	body := d.Get("body").(string)

	options := github.ProjectOptions{
		Name: &name,
		Body: &body,
	}

	projectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	ctx := prepareResourceContext(d)

	log.Printf("[DEBUG] Updating organization project: %s (%s)", d.Id(), orgName)
	if _, _, err := client.Projects.UpdateProject(ctx, projectID, &options); err != nil {
		return err
	}

	return resourceGithubOrganizationProjectRead(d, meta)
}

func resourceGithubOrganizationProjectDelete(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client
	projectID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	ctx := prepareResourceContext(d)

	log.Printf("[DEBUG] Deleting organization project: %s (%s)", d.Id(), orgName)
	_, err = client.Projects.DeleteProject(ctx, projectID)
	return err
}
