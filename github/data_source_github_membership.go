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

	client := meta.(*Organization).client
	orgName := meta.(*Organization).name

	ctx := context.Background()

	membership, resp, err := client.Organizations.GetOrgMembership(ctx,
		username, orgName)

	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(membership.Organization.Login, membership.User.Login))

	d.Set("username", membership.User.Login)
	d.Set("role", membership.Role)
	d.Set("etag", resp.Header.Get("ETag"))
	return nil
}
