package github

import (
	"context"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubMembershipCreateOrUpdate,
		Read:   resourceGithubMembershipRead,
		Update: resourceGithubMembershipCreateOrUpdate,
		Delete: resourceGithubMembershipDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"role": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateValueFunc([]string{"member", "admin"}),
				Default:      "member",
			},
		},
	}
}

func resourceGithubMembershipCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	orgName := meta.(*Organization).name
	username := d.Get("username").(string)
	roleName := d.Get("role").(string)

	log.Printf("[DEBUG] Creating membership: %s/%s", orgName, username)
	membership, _, err := client.Organizations.EditOrgMembership(context.TODO(),
		username,
		orgName,
		&github.Membership{
			Role: github.String(roleName),
		},
	)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(membership.Organization.Login, membership.User.Login))

	return resourceGithubMembershipRead(d, meta)
}

func resourceGithubMembershipRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	orgName := meta.(*Organization).name
	_, username, err := parseTwoPartID(d.Id())
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Reading membership: %s", d.Id())
	membership, _, err := client.Organizations.GetOrgMembership(context.TODO(),
		username, orgName)
	if err != nil {
		if err, ok := err.(*github.ErrorResponse); ok {
			if err.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing membership %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("username", membership.User.Login)
	d.Set("role", membership.Role)

	return nil
}

func resourceGithubMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	orgName := meta.(*Organization).name

	log.Printf("[DEBUG] Deleting membership: %s", d.Id())
	_, err := client.Organizations.RemoveOrgMembership(context.TODO(),
		d.Get("username").(string), orgName)

	return err
}
