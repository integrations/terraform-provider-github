package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubDependabotOrganizationPublicKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubDependabotOrganizationPublicKeyRead,

		Schema: map[string]*schema.Schema{
			"key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubDependabotOrganizationPublicKeyRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	ctx := context.Background()

	publicKey, _, err := client.Dependabot.GetOrgPublicKey(ctx, owner)
	if err != nil {
		return err
	}

	d.SetId(publicKey.GetKeyID())
	d.Set("key_id", publicKey.GetKeyID())
	d.Set("key", publicKey.GetKey())

	return nil
}
