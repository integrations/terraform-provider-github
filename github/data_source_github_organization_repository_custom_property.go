package github

import (
	"context"
	"errors"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubOrganizationRepositoryCustomProperty() *schema.Resource {
	return &schema.Resource{
		Description: "Looks up a single GitHub organization custom property definition by name.",
		ReadContext: dataSourceGithubOrganizationRepositoryCustomPropertyRead,

		Schema: map[string]*schema.Schema{
			"property_name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Name of the custom property to look up.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
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
	}
}

func dataSourceGithubOrganizationRepositoryCustomPropertyRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	client := meta.v3client
	owner := meta.name
	propertyName := d.Get("property_name").(string)

	tflog.Debug(ctx, "Reading organization custom property", map[string]any{"org": owner, "property": propertyName})

	cp, _, err := client.Organizations.GetCustomProperty(ctx, owner, propertyName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == 404 {
			return diag.Errorf("organization custom property %q not found in %q", propertyName, owner)
		}
		return diag.Errorf(organizationCustomPropertyReadErrorFormat, propertyName, err)
	}

	if cp.GetPropertyName() == "" {
		return diag.Errorf("organization %q returned a custom property with an empty name when reading %q", owner, propertyName)
	}

	return setOrganizationRepositoryCustomPropertyState(d, cp)
}
