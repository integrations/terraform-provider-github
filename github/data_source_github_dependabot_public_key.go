package github

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubDependabotPublicKey() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve the public key for a repository's Dependabot secrets.",
		Read:        dataSourceGithubDependabotPublicKeyRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the repository.",
			},
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

func dataSourceGithubDependabotPublicKeyRead(d *schema.ResourceData, meta any) error {
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
