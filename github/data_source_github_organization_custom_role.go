package github

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v83/github"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationCustomRole() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "This data source is deprecated and will be removed in a future release. Use the github_organization_repository_role data source instead.",

		ReadContext: dataSourceGithubOrganizationCustomRoleRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"base_role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permissions": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubOrganizationCustomRoleRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	// ListCustomRepoRoles returns a list of all custom repository roles for an organization.
	// There is an API endpoint for getting a single custom repository role, but is not
	// implemented in the go-github library.
	roleList, _, err := client.Organizations.ListCustomRepoRoles(ctx, orgName)
	if err != nil {
		return diag.Errorf("error querying GitHub custom repository roles %s: %v", orgName, err)
	}

	var role *github.CustomRepoRoles
	for _, r := range roleList.CustomRepoRoles {
		if fmt.Sprint(*r.Name) == d.Get("name").(string) {
			role = r
			break
		}
	}

	if role == nil {
		log.Printf("[WARN] GitHub custom repository role (%s) not found.", d.Get("name").(string))
		d.SetId("")
		return nil
	}

	d.SetId(fmt.Sprint(*role.ID))
	err = d.Set("name", role.Name)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("description", role.Description)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("base_role", role.BaseRole)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("permissions", role.Permissions)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
