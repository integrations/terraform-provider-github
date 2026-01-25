package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsOrganizationPublicKey() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve the public key for an organization's GitHub Actions secrets.",
		Read:        dataSourceGithubActionsOrganizationPublicKeyRead,

		Schema: map[string]*schema.Schema{
			"key_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The identifier for the key.",
			},
			"key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Base64 encoded public key.",
			},
		},
	}
}

func dataSourceGithubActionsOrganizationPublicKeyRead(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	ctx := context.Background()

	publicKey, _, err := client.Actions.GetOrgPublicKey(ctx, owner)
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
