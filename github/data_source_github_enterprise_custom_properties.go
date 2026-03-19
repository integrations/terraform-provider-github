package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseCustomProperties() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubEnterpriseCustomPropertiesRead,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"values_editable_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubEnterpriseCustomPropertiesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	entSlug := d.Get("enterprise_slug").(string)

	propertyAttributes, _, err := client.Enterprise.GetCustomProperty(ctx, entSlug, d.Get("property_name").(string))
	if err != nil {
		return diag.Errorf("error querying GitHub custom properties %s: %v", entSlug, err)
	}

	// TODO: Add support for other types of default values
	defaultValue, _ := propertyAttributes.DefaultValueString()

	d.SetId("enterprise-custom-properties")
	_ = d.Set("enterprise_slug", entSlug)
	_ = d.Set("allowed_values", propertyAttributes.AllowedValues)
	_ = d.Set("default_value", defaultValue)
	_ = d.Set("description", propertyAttributes.Description)
	_ = d.Set("property_name", propertyAttributes.PropertyName)
	_ = d.Set("required", propertyAttributes.Required)
	_ = d.Set("value_type", string(propertyAttributes.ValueType))
	_ = d.Set("values_editable_by", propertyAttributes.ValuesEditableBy)

	return nil
}
