package github

import (
	"context"
	"log"
	"strconv"

	"github.com/google/go-github/v36/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubOrganization() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubOrganizationCreate,
		Read:   resourceGithubOrganizationRead,
		Update: resourceGithubOrganizationUpdate,
		Delete: resourceGithubOrganizationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"login": {
				Type:     schema.TypeString,
				Required: true,
			},
			"admin": {
				Type:     schema.TypeString,
				Required: true,
			},
			"profile_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceGithubOrganizationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	ownerName := meta.(*Owner).name
	login := github.String(d.Get("login").(string))
	admin := github.String(d.Get("admin").(string))

	newOrganization := github.Organization{
		Login: login,
	}

	ctx := context.Background()

	log.Printf("[DEBUG] Creating organization: %s (%s)", *login, ownerName)
	githubOrganization, _, err := client.Admin.CreateOrg(ctx, &newOrganization, *admin)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(githubOrganization.GetID(), 10))
	return resourceGithubOrganizationRead(d, meta)
}

func resourceGithubOrganizationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	ctx := context.Background()

	log.Printf("[DEBUG] Reading organization: %s", orgName)
	org, _, err := client.Organizations.Get(ctx, orgName)
	if err != nil {
		return err
	}

	d.Set("login", org.GetLogin())
	d.Set("profile_name", org.GetName())

	return nil
}

func resourceGithubOrganizationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	ctx := context.Background()

	newName := github.String(d.Get("login").(string))
	newOrganization := github.Organization{
		Login: &orgName,
		Name:  github.String(d.Get("profile_name").(string)),
	}

	log.Printf("[DEBUG] Updating organization: %s", orgName)
	_, _, err := client.Admin.RenameOrg(ctx, &newOrganization, *newName)
	if err != nil {
		return err
	}

	org, _, err := client.Organizations.Edit(ctx, orgName, &newOrganization)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(org.GetID(), 10))
	return resourceGithubOrganizationRead(d, meta)
}

func resourceGithubOrganizationDelete(d *schema.ResourceData, meta interface{}) error {
	// no github api for that?
	return nil
}
