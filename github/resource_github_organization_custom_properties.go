package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubOrganizationCustomProperties() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubCustomPropertiesCreate,
		ReadContext:   resourceGithubCustomPropertiesRead,
		UpdateContext: resourceGithubCustomPropertiesUpdate,
		DeleteContext: resourceGithubCustomPropertiesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubCustomPropertiesImport,
		},

		CustomizeDiff: customdiff.Sequence(
			customdiff.ComputedIf("slug", func(_ context.Context, d *schema.ResourceDiff, meta any) bool {
				return d.HasChange("name")
			}),
		),

		Schema: map[string]*schema.Schema{
			"property_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the custom property",
			},
			"value_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The type of the custom property",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{string(github.PropertyValueTypeString), string(github.PropertyValueTypeSingleSelect), string(github.PropertyValueTypeMultiSelect), string(github.PropertyValueTypeTrueFalse), string(github.PropertyValueTypeURL)}, false)),
			},
			"required": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the custom property is required",
			},
			"default_value": {
				Type:        schema.TypeString,
				Description: "The default value of the custom property. Not supported for multi_select properties.",
				Optional:    true,
				Computed:    true,
			},
			"description": {
				Description: "The description of the custom property",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"allowed_values": {
				Description: "The allowed values of the custom property",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"values_editable_by": {
				Description:      "Who can edit the values of the custom property. Can be one of 'org_actors' or 'org_and_repo_actors'. If not specified, the default is 'org_actors' (only organization owners can edit values)",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"org_actors", "org_and_repo_actors"}, false)),
			},
		},
	}
}

func resourceGithubCustomPropertiesCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	ownerName := meta.(*Owner).name

	propertyName := d.Get("property_name").(string)
	valueType := github.PropertyValueType(d.Get("value_type").(string))
	required := d.Get("required").(bool)
	defaultValue := d.Get("default_value").(string)
	description := d.Get("description").(string)
	allowedValues := d.Get("allowed_values").([]any)
	var allowedValuesString []string
	for _, v := range allowedValues {
		allowedValuesString = append(allowedValuesString, v.(string))
	}

	customProperty := &github.CustomProperty{
		PropertyName:  &propertyName,
		ValueType:     valueType,
		Required:      &required,
		DefaultValue:  &defaultValue,
		Description:   &description,
		AllowedValues: allowedValuesString,
	}

	if val, ok := d.GetOk("values_editable_by"); ok {
		str := val.(string)
		customProperty.ValuesEditableBy = &str
	}

	diags := multiSelectDefaultValueWarning(valueType, defaultValue)

	customProperty, _, err := client.Organizations.CreateOrUpdateCustomProperty(ctx, ownerName, d.Get("property_name").(string), customProperty)
	if err != nil {
		return append(diags, diag.FromErr(err)...)
	}

	d.SetId(*customProperty.PropertyName)
	return append(diags, resourceGithubCustomPropertiesRead(ctx, d, meta)...)
}

// multiSelectDefaultValueWarning warns when a default value is configured for a
// multi_select property. GitHub returns those defaults as a list of strings,
// which cannot be represented by the string default_value attribute, so the
// configured value is not reflected in state and shows up as a change on every
// plan.
func multiSelectDefaultValueWarning(valueType github.PropertyValueType, defaultValue string) diag.Diagnostics {
	if valueType != github.PropertyValueTypeMultiSelect || defaultValue == "" {
		return nil
	}

	return diag.Diagnostics{
		{
			Severity:      diag.Warning,
			Summary:       "default_value is not supported for multi_select properties",
			Detail:        "The default value of a multi_select property cannot be read back by this provider, so it is not stored in state and every plan will show a change for default_value. Remove default_value to avoid this.",
			AttributePath: cty.GetAttrPath("default_value"),
		},
	}
}

func resourceGithubCustomPropertiesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	ownerName := meta.(*Owner).name

	customProperty, _, err := client.Organizations.GetCustomProperty(ctx, ownerName, d.Get("property_name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	// multi_select is not supported: its default value is a []string, which
	// cannot round-trip through the TypeString default_value attribute.
	var defaultValue string
	switch customProperty.ValueType {
	case github.PropertyValueTypeTrueFalse:
		if b, ok := customProperty.DefaultValueBool(); ok {
			defaultValue = strconv.FormatBool(b)
		}
	default:
		if s, ok := customProperty.DefaultValueString(); ok {
			defaultValue = s
		}
	}

	d.SetId(*customProperty.PropertyName)
	_ = d.Set("allowed_values", customProperty.AllowedValues)
	_ = d.Set("default_value", defaultValue)
	_ = d.Set("description", customProperty.Description)
	_ = d.Set("property_name", customProperty.PropertyName)
	_ = d.Set("required", customProperty.Required)
	_ = d.Set("value_type", string(customProperty.ValueType))
	_ = d.Set("values_editable_by", customProperty.ValuesEditableBy)

	return nil
}

func resourceGithubCustomPropertiesUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// Create issues a PUT, which the API treats as an upsert, and reads the
	// property back afterwards.
	return resourceGithubCustomPropertiesCreate(ctx, d, meta)
}

func resourceGithubCustomPropertiesDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	ownerName := meta.(*Owner).name

	_, err := client.Organizations.RemoveCustomProperty(ctx, ownerName, d.Get("property_name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubCustomPropertiesImport(_ context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	if err := d.Set("property_name", d.Id()); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
