package github

import (
	"context"

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
			State: resourceGithubMembershipImport,
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
	n := d.Get("username").(string)
	r := d.Get("role").(string)

	membership, _, err := client.Organizations.EditOrgMembership(context.TODO(), n, meta.(*Organization).name,
		&github.Membership{Role: &r})
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(membership.Organization.Login, membership.User.Login))

	return resourceGithubMembershipRead(d, meta)
}

func resourceGithubMembershipRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	_, n := parseTwoPartID(d.Id())

	membership, _, err := client.Organizations.GetOrgMembership(context.TODO(), n, meta.(*Organization).name)
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("username", membership.User.Login)
	d.Set("role", membership.Role)
	return nil
}

func resourceGithubMembershipUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	n := d.Get("username").(string)
	r := d.Get("role").(string)

	membership, _, err := client.Organizations.EditOrgMembership(context.TODO(), n, meta.(*Organization).name, &github.Membership{
		Role: &r,
	})
	if err != nil {
		return err
	}
	d.SetId(buildTwoPartID(membership.Organization.Login, membership.User.Login))

	return nil
}

func resourceGithubMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client
	n := d.Get("username").(string)

	_, err := client.Organizations.RemoveOrgMembership(context.TODO(), n, meta.(*Organization).name)

	return err
}

func resourceGithubMembershipImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// All we do here is validate that the import string is in a correct enough
	// format to be parsed.  parseTwoPartID will panic if it's missing elements,
	// and is used otherwise in places where that should never happen, so we want
	// to keep it that way.
	if err := validateTwoPartID(d.Id()); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
