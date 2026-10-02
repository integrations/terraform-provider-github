package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationRepositoryCustomProperties() *schema.Resource {
	return &schema.Resource{
		Description: "Lists all custom property definitions for a GitHub organization. See the [GitHub REST API documentation](https://docs.github.com/en/rest/orgs/custom-properties?apiVersion=2022-11-28#get-all-custom-properties-for-an-organization) for required permissions.",
		ReadContext: dataSourceGithubOrganizationRepositoryCustomPropertiesRead,

		Schema: map[string]*schema.Schema{
			"properties": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Custom property definitions in the organization.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"property_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the custom property.",
						},
						"value_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the custom property.",
						},
						"required": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the custom property must be set on every repository.",
						},
						"default_value": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Default value applied to repositories that do not explicitly set the property. Holds multiple elements only when `value_type` is `multi_select`.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Short description of the custom property.",
						},
						"allowed_values": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Allowed values when `value_type` is `single_select` or `multi_select`.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"values_editable_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Who can edit values of this property on repositories.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationRepositoryCustomPropertiesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	properties, _, err := meta.v3client.Organizations.GetAllCustomProperties(ctx, meta.name)
	if err != nil {
		return diag.Errorf("error listing organization custom properties: %v", err)
	}

	propertyValues := make([]any, 0, len(properties))
	for _, property := range properties {
		if property == nil || property.GetPropertyName() == "" {
			return diag.Errorf("organization %q returned a custom property with an empty name", meta.name)
		}

		values, err := organizationRepositoryCustomPropertyToMap(property)
		if err != nil {
			return diag.Errorf(organizationCustomPropertyReadErrorFormat, property.GetPropertyName(), err)
		}

		propertyValues = append(propertyValues, values)
	}

	d.SetId(meta.name)
	if err := d.Set("properties", propertyValues); err != nil {
		return diag.Errorf("error setting organization custom properties: %v", err)
	}

	return nil
}
