package github

import (
	"context"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseCustomProperty() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve information about a custom property definition for a GitHub enterprise.",

		ReadContext: dataSourceGithubEnterpriseCustomPropertyRead,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
			},
			"property_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the custom property.",
			},
			"value_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the value for the property.",
			},
			"required": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the custom property is required.",
			},
			"default_values": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The default value(s) of the custom property.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A short description of the custom property.",
			},
			"allowed_values": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "An ordered list of allowed values for the property.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"values_editable_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Who can edit the values of the property. Can be one of 'org_actors' or 'org_and_repo_actors'.",
			},
		},
	}
}

func dataSourceGithubEnterpriseCustomPropertyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug := d.Get("enterprise_slug").(string)
	propertyName := d.Get("property_name").(string)

	property, _, err := client.Enterprise.GetCustomProperty(ctx, enterpriseSlug, propertyName)
	if err != nil {
		return diag.Errorf("error reading enterprise custom property %s/%s: %v", enterpriseSlug, propertyName, err)
	}

	var defaultValues []string
	if property.ValueType == github.PropertyValueTypeMultiSelect {
		if vals, ok := property.DefaultValueStrings(); ok {
			defaultValues = vals
		}
	} else {
		if val, ok := property.DefaultValueString(); ok {
			defaultValues = []string{val}
		}
	}

	d.SetId(buildTwoPartID(enterpriseSlug, propertyName))

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("property_name", property.GetPropertyName()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("value_type", string(property.ValueType)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("required", property.GetRequired()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default_values", defaultValues); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", property.GetDescription()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_values", property.AllowedValues); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("values_editable_by", property.GetValuesEditableBy()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
