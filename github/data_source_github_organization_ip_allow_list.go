package github

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubOrganizationIpAllowList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrganizationIpAllowListRead,

		Schema: map[string]*schema.Schema{
			"ip_allow_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_list_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_active": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationIpAllowListRead(d *schema.ResourceData, meta interface{}) error {
	orgId := meta.(*Owner).id

	ipAllowListEntries, err := getOrganizationIpAllowListEntries(meta)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(orgId, 10))
	d.Set("ip_allow_list", ipAllowListEntries)

	return nil
}
