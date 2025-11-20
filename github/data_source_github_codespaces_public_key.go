package github

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubCodespacesPublicKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubCodespacesPublicKeyRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
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

func dataSourceGithubCodespacesPublicKeyRead(d *schema.ResourceData, meta any) error {
	repository := d.Get("repository").(string)
	owner := meta.(*Owner).name
	log.Printf("[INFO] Refreshing GitHub Codespaces Public Key from: %s/%s", owner, repository)

	client := meta.(*Owner).v3client
	ctx := context.Background()

	publicKey, _, err := client.Codespaces.GetRepoPublicKey(ctx, owner, repository)
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
