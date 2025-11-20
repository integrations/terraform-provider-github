package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationCustomProperties() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrganizationCustomPropertiesRead,

		Schema: map[string]*schema.Schema{
			"property_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"required": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default_value": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowed_values": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceGithubOrganizationCustomPropertiesRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()
	orgName := meta.(*Owner).name

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	propertyAttributes, _, err := client.Organizations.GetCustomProperty(ctx, orgName, d.Get("property_name").(string))
	if err != nil {
		return fmt.Errorf("error querying GitHub custom properties %s: %w", orgName, err)
	}

	d.SetId("org-custom-properties")
	_ = d.Set("allowed_values", propertyAttributes.AllowedValues)
	_ = d.Set("default_value", propertyAttributes.DefaultValue)
	_ = d.Set("description", propertyAttributes.Description)
	_ = d.Set("property_name", propertyAttributes.PropertyName)
	_ = d.Set("required", propertyAttributes.Required)
	_ = d.Set("value_type", propertyAttributes.ValueType)

	return nil
}
