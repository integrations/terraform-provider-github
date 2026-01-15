package github

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceGithubSshKeys() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubSshKeysRead,

		Schema: map[string]*schema.Schema{
			"keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceGithubSshKeysRead(d *schema.ResourceData, meta any) error {
	owner := meta.(*Owner)

	api, _, err := owner.v3client.Meta.Get(owner.StopContext)
	if err != nil {
		return err
	}

	d.SetId("github-ssh-keys")
	if err = d.Set("keys", api.SSHKeys); err != nil {
		return err
	}

	return nil
}
