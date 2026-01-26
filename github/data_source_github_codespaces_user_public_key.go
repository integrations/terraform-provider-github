package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubCodespacesUserPublicKey() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve the public key for a user's Codespaces secrets.",
		Read:        dataSourceGithubCodespacesUserPublicKeyRead,

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

func dataSourceGithubCodespacesUserPublicKeyRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client

	ctx := context.Background()

	publicKey, _, err := client.Codespaces.GetUserPublicKey(ctx)
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
