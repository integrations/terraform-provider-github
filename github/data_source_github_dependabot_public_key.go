package github

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubDependabotPublicKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubDependabotPublicKeyRead,

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

func dataSourceGithubDependabotPublicKeyRead(d *schema.ResourceData, meta interface{}) error {
	repository := d.Get("repository").(string)
	owner := meta.(*Owner).name
	log.Printf("[INFO] Refreshing GitHub Dependabot Public Key from: %s/%s", owner, repository)

	client := meta.(*Owner).v3client
	ctx := context.Background()

	publicKey, _, err := client.Dependabot.GetRepoPublicKey(ctx, owner, repository)
	if err != nil {
		return err
	}

	d.SetId(publicKey.GetKeyID())
	d.Set("key_id", publicKey.GetKeyID())
	d.Set("key", publicKey.GetKey())

	return nil
}
