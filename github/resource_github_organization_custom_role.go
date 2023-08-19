package github

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v54/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubOrganizationCustomRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubOrganizationCustomRoleCreate,
		Read:   resourceGithubOrganizationCustomRoleRead,
		Update: resourceGithubOrganizationCustomRoleUpdate,
		Delete: resourceGithubOrganizationCustomRoleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The organization custom repository role to create.",
			},
			"base_role": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The base role for the custom repository role.",
				ValidateFunc: validateValueFunc([]string{"read", "triage", "write", "maintain"}),
			},
			"permissions": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				MinItems:    1, // At least one permission should be passed.
				Description: "The permissions for the custom repository role.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the custom repository role.",
			},
		},
	}
}

func resourceGithubOrganizationCustomRoleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	permissions := d.Get("permissions").(*schema.Set).List()
	permissionsStr := make([]string, len(permissions))
	for i, v := range permissions {
		permissionsStr[i] = v.(string)
	}

	role, _, err := client.Organizations.CreateCustomRepoRole(ctx, orgName, &github.CreateOrUpdateCustomRoleOptions{
		Name:        github.String(d.Get("name").(string)),
		Description: github.String(d.Get("description").(string)),
		BaseRole:    github.String(d.Get("base_role").(string)),
		Permissions: permissionsStr,
	})

	if err != nil {
		return fmt.Errorf("error creating GitHub custom repository role %s (%s): %s", orgName, d.Get("name").(string), err)
	}

	d.SetId(fmt.Sprint(*role.ID))
	return resourceGithubOrganizationCustomRoleRead(d, meta)
}

func resourceGithubOrganizationCustomRoleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()
	orgName := meta.(*Owner).name

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	roleID := d.Id()

	// ListCustomRepoRoles returns a list of all custom repository roles for an organization.
	// There is an API endpoint for getting a single custom repository role, but is not
	// implemented in the go-github library.
	roleList, _, err := client.Organizations.ListCustomRepoRoles(ctx, orgName)
	if err != nil {
		return fmt.Errorf("error querying GitHub custom repository roles %s: %s", orgName, err)
	}

	var role *github.CustomRepoRoles
	for _, r := range roleList.CustomRepoRoles {
		if fmt.Sprint(*r.ID) == roleID {
			role = r
			break
		}
	}

	if role == nil {
		log.Printf("[WARN] GitHub custom repository role (%s/%s) not found, removing from state", orgName, roleID)
		d.SetId("")
		return nil
	}

	d.Set("name", role.Name)
	d.Set("description", role.Description)
	d.Set("base_role", role.BaseRole)
	d.Set("permissions", role.Permissions)

	return nil
}

func resourceGithubOrganizationCustomRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()
	orgName := meta.(*Owner).name

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	roleID := d.Id()
	permissions := d.Get("permissions").(*schema.Set).List()
	permissionsStr := make([]string, len(permissions))
	for i, v := range permissions {
		permissionsStr[i] = v.(string)
	}

	update := &github.CreateOrUpdateCustomRoleOptions{
		Name:        github.String(d.Get("name").(string)),
		Description: github.String(d.Get("description").(string)),
		BaseRole:    github.String(d.Get("base_role").(string)),
		Permissions: permissionsStr,
	}

	if _, _, err := client.Organizations.UpdateCustomRepoRole(ctx, orgName, roleID, update); err != nil {
		return fmt.Errorf("error updating GitHub custom repository role %s (%s): %s", orgName, roleID, err)
	}

	return resourceGithubOrganizationCustomRoleRead(d, meta)
}

func resourceGithubOrganizationCustomRoleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()
	orgName := meta.(*Owner).name

	err := checkOrganization(meta)
	if err != nil {
		return err
	}
	roleID := d.Id()

	_, err = client.Organizations.DeleteCustomRepoRole(ctx, orgName, roleID)
	if err != nil {
		return fmt.Errorf("Error deleting GitHub custom repository role %s (%s): %s", orgName, roleID, err)
	}

	return nil
}
