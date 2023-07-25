package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubCodespacesUserPublicKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubCodespacesUserPublicKeyRead,

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

func dataSourceGithubCodespacesUserPublicKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	ctx := context.Background()

	publicKey, _, err := client.Codespaces.GetUserPublicKey(ctx)
	if err != nil {
		return err
	}

	d.SetId(publicKey.GetKeyID())
	d.Set("key_id", publicKey.GetKeyID())
	d.Set("key", publicKey.GetKey())

	return nil
}
