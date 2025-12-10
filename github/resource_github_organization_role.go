package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubOrganizationRole() *schema.Resource {
	return &schema.Resource{
		Description: "Manage a custom organization role.",

		CreateContext: resourceGithubOrganizationRoleCreate,
		ReadContext:   resourceGithubOrganizationRoleRead,
		UpdateContext: resourceGithubOrganizationRoleUpdate,
		DeleteContext: resourceGithubOrganizationRoleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"role_id": {
				Description: "The ID of the organization role.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"name": {
				Description: "The name of the organization role.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "The description of the organization role.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"base_role": {
				Description:      "The system role from which this role inherits permissions.",
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "none",
				ValidateDiagFunc: validateValueFunc([]string{"none", "read", "triage", "write", "maintain", "admin"}),
			},
			"permissions": {
				Description: "The permissions for the organization role.",
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceGithubOrganizationRoleCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	permissions := d.Get("permissions").(*schema.Set).List()
	permissionsStr := make([]string, len(permissions))
	for i, v := range permissions {
		permissionsStr[i] = v.(string)
	}

	role, _, err := client.Organizations.CreateCustomOrgRole(ctx, orgName, &github.CreateOrUpdateOrgRoleOptions{
		Name:        github.String(d.Get("name").(string)),
		Description: github.String(d.Get("description").(string)),
		BaseRole:    github.String(d.Get("base_role").(string)),
		Permissions: permissionsStr,
	})
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating organization role (%s/%s): %w", orgName, d.Get("name").(string), err))
	}

	d.SetId(fmt.Sprint(role.GetID()))
	return nil
}

func resourceGithubOrganizationRoleRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	roleId, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	role, _, err := client.Organizations.GetOrgRole(ctx, orgName, roleId)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] organization role (%s/%d) not found, removing from state", orgName, roleId)
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if err = d.Set("role_id", role.GetID()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("name", role.Name); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("description", role.Description); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("base_role", role.BaseRole); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("permissions", role.Permissions); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationRoleUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	roleId, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	permissions := d.Get("permissions").(*schema.Set).List()
	permissionsStr := make([]string, len(permissions))
	for i, v := range permissions {
		permissionsStr[i] = v.(string)
	}

	update := &github.CreateOrUpdateOrgRoleOptions{
		Name:        github.String(d.Get("name").(string)),
		Description: github.String(d.Get("description").(string)),
		BaseRole:    github.String(d.Get("base_role").(string)),
		Permissions: permissionsStr,
	}

	_, _, err = client.Organizations.UpdateCustomOrgRole(ctx, orgName, roleId, update)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating organization role (%s/%s): %w", orgName, d.Get("name").(string), err))
	}

	return nil
}

func resourceGithubOrganizationRoleDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	roleId, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.Organizations.DeleteCustomOrgRole(ctx, orgName, roleId)
	if err != nil {
		return diag.FromErr(fmt.Errorf("Error deleting organization role %d: %w", roleId, err))
	}

	return nil
}
