package github

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubOrganizationRoleUser() *schema.Resource {
	return &schema.Resource{
		Description: "Manage an association between an organization role and a user.",

		CreateContext: resourceGithubOrganizationRoleUserCreate,
		ReadContext:   resourceGithubOrganizationRoleUserRead,
		DeleteContext: resourceGithubOrganizationRoleUserDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"role_id": {
				Description: "The unique identifier of the organization role.",
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
			},
			"login": {
				Description: "The login for the GitHub user account.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceGithubOrganizationRoleUserCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	roleId := int64(d.Get("role_id").(int))
	login := d.Get("login").(string)

	_, err = client.Organizations.AssignOrgRoleToUser(ctx, orgName, login, roleId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildTwoPartID(strconv.FormatInt(roleId, 10), login))

	return nil
}

func resourceGithubOrganizationRoleUserRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	roleIdString, login, err := parseTwoPartID(d.Id(), "role_id", "login")
	if err != nil {
		return diag.FromErr(err)
	}
	roleId, err := strconv.ParseInt(roleIdString, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	opts := &github.ListOptions{
		PerPage: maxPerPage,
	}

	var user *github.User
	for {
		users, resp, err := client.Organizations.ListUsersAssignedToOrgRole(ctx, orgName, roleId, opts)
		if err != nil {
			return diag.FromErr(err)
		}

		for _, u := range users {
			if u.GetLogin() == login {
				user = u
				break
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	if user == nil {
		log.Printf("[INFO] Removing organization role User (%d:%s) from state because it no longer exists in GitHub", roleId, login)
		d.SetId("")
		return nil
	}

	if err = d.Set("role_id", roleId); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("login", login); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationRoleUserDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	roleId := int64(d.Get("role_id").(int))
	login := d.Get("login").(string)

	_, err = client.Organizations.RemoveOrgRoleFromUser(ctx, orgName, login, roleId)
	if err != nil {
		return diag.FromErr(fmt.Errorf("Error deleting organization role user %d %s: %w", roleId, login, err))
	}

	return nil
}
