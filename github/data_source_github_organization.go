package github

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubOrganization() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrganizationRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"node_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"login": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plan": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubOrganizationRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	log.Printf("[INFO] Refreshing GitHub Organization: %s", name)

	client := meta.(*Owner).v3client
	ctx := meta.(*Owner).StopContext

	organization, _, err := client.Organizations.Get(ctx, name)
	if err != nil {
		return err
	}

	plan := organization.GetPlan()

	d.SetId(strconv.FormatInt(organization.GetID(), 10))
	d.Set("login", organization.GetLogin())
	d.Set("name", organization.GetName())
	d.Set("description", organization.GetDescription())
	d.Set("plan", plan.Name)

	return nil
}
