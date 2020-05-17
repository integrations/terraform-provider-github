package github

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubMembership() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubMembershipRead,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role": {
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

func dataSourceGithubMembershipRead(d *schema.ResourceData, meta interface{}) error {
	username := d.Get("username").(string)
	log.Printf("[INFO] Refreshing GitHub membership: %s", username)

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	ctx := context.Background()

	membership, resp, err := client.Organizations.GetOrgMembership(ctx,
		username, orgName)

	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(membership.GetOrganization().GetLogin(), membership.GetUser().GetLogin()))

	d.Set("username", membership.GetUser().GetLogin())
	d.Set("role", membership.GetRole())
	d.Set("etag", resp.Header.Get("ETag"))
	return nil
}
