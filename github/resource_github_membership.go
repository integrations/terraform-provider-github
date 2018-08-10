package github

import (
	"context"
	"log"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubMembership() *schema.Resource {

	return &schema.Resource{
		Create: resourceGithubMembershipCreate,
		Read:   resourceGithubMembershipRead,
		Update: resourceGithubMembershipUpdate,
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

func resourceGithubMembershipCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	membership, _, err := client.Organizations.EditOrgMembership(context.TODO(),
		d.Get("username").(string),
		meta.(*Organization).name,
		&github.Membership{
			Role: github.String(d.Get("role").(string)),
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
	_, username, err := parseTwoPartID(d.Id())
	if err != nil {
		return err
	}

	membership, _, err := client.Organizations.GetOrgMembership(context.TODO(),
		username, meta.(*Organization).name)
	if err != nil {
		log.Printf("[WARN] GitHub Membership (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("username", membership.User.Login)
	d.Set("role", membership.Role)
	return nil
}

func resourceGithubMembershipUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	membership, _, err := client.Organizations.EditOrgMembership(context.TODO(),
		d.Get("username").(string),
		meta.(*Organization).name,
		&github.Membership{
			Role: github.String(d.Get("role").(string)),
		},
	)
	if err != nil {
		return err
	}
	d.SetId(buildTwoPartID(membership.Organization.Login, membership.User.Login))

	return nil
}

func resourceGithubMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	_, err := client.Organizations.RemoveOrgMembership(context.TODO(),
		d.Get("username").(string), meta.(*Organization).name)

	return err
}
