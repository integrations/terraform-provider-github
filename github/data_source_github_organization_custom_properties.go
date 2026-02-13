package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationCustomProperties() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve information about an organization's custom property.",
		Read:        dataSourceGithubOrganizationCustomPropertiesRead,

		Schema: map[string]*schema.Schema{
			"property_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the custom property.",
			},
			"value_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the custom property. Can be one of 'string', 'single_select', 'multi_select', or 'true_false'.",
			},
			"required": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the custom property is required.",
			},
			"default_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The default value of the custom property.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the custom property.",
			},
			"allowed_values": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of allowed values for the custom property. Only populated when value_type is 'single_select' or 'multi_select'.",
			},
			"values_editable_by": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Who can edit the values of the custom property. Can be one of 'org_actors' or 'org_and_repo_actors'.",
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

	// TODO: Add support for other types of default values
	defaultValue, _ := propertyAttributes.DefaultValueString()

	d.SetId("org-custom-properties")
	_ = d.Set("allowed_values", propertyAttributes.AllowedValues)
	_ = d.Set("default_value", defaultValue)
	_ = d.Set("description", propertyAttributes.Description)
	_ = d.Set("property_name", propertyAttributes.PropertyName)
	_ = d.Set("required", propertyAttributes.Required)
	_ = d.Set("value_type", string(propertyAttributes.ValueType))
	_ = d.Set("values_editable_by", propertyAttributes.ValuesEditableBy)

	return nil
}
