package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubCodespacesOrganizationPublicKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubCodespacesOrganizationPublicKeyRead,

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

func dataSourceGithubCodespacesOrganizationPublicKeyRead(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	ctx := context.Background()

	publicKey, _, err := client.Codespaces.GetOrgPublicKey(ctx, owner)
	if err != nil {
		return err
	}

	d.SetId(publicKey.GetKeyID())
	err = d.Set("key_id", publicKey.GetKeyID())
	if err != nil {
		return err
	}
	err = d.Set("key", publicKey.GetKey())
	if err != nil {
		return err
	}

	return nil
}
