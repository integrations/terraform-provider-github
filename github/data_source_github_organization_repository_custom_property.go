package github

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationRepositoryCustomProperty() *schema.Resource {
	return &schema.Resource{
		Description: "Looks up a single GitHub organization custom property definition by name.",
		ReadContext: dataSourceGithubOrganizationRepositoryCustomPropertyRead,

		Schema: map[string]*schema.Schema{
			"property_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the custom property to look up.",
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Default value applied to repositories that do not explicitly set the property.",
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
	}
}

func dataSourceGithubOrganizationRepositoryCustomPropertyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	propertyName := d.Get("property_name").(string)

	tflog.Debug(ctx, "Reading organization custom property", map[string]any{"org": orgName, "property": propertyName})

	cp, _, err := client.Organizations.GetCustomProperty(ctx, orgName, propertyName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == 404 {
			return diag.FromErr(fmt.Errorf("organization custom property %q not found in %q", propertyName, orgName))
		}
		return diag.FromErr(fmt.Errorf("error reading organization custom property %q: %w", propertyName, err))
	}

	if !slices.Contains([]github.PropertyValueType{
		github.PropertyValueTypeSingleSelect,
		github.PropertyValueTypeMultiSelect,
	}, cp.ValueType) {
		cp.AllowedValues = nil
	}

	if err := setOrganizationRepositoryCustomPropertyState(d, cp); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
