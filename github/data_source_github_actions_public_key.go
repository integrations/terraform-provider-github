package github

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubActionsPublicKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubActionsPublicKeyRead,

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

func dataSourceGithubActionsPublicKeyRead(d *schema.ResourceData, meta interface{}) error {
	repository := d.Get("repository").(string)
	owner := meta.(*Owner).name
	log.Printf("[INFO] Refreshing GitHub Actions Public Key from: %s/%s", owner, repository)

	client := meta.(*Owner).v3client
	ctx := context.Background()

	publicKey, _, err := client.Actions.GetPublicKey(ctx, owner, repository)
	if err != nil {
		return err
	}

	d.SetId(publicKey.GetKeyID())
	d.Set("key_id", publicKey.GetKeyID())
	d.Set("key", publicKey.GetKey())

	return nil
}
