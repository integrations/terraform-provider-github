package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubToken() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubTokenRead,

		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubTokenRead(d *schema.ResourceData, meta interface{}) error {
	owner := meta.(*Owner)
	d.Set("value", owner.Token)
	return nil
}
