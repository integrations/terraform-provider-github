package github

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var organizationCustomPropertyValueTypes = []string{
	string(github.PropertyValueTypeString),
	string(github.PropertyValueTypeSingleSelect),
	string(github.PropertyValueTypeMultiSelect),
	string(github.PropertyValueTypeTrueFalse),
	string(github.PropertyValueTypeURL),
}

var organizationCustomPropertyValuesEditableBy = []string{"org_actors", "org_and_repo_actors"}

func resourceGithubOrganizationRepositoryCustomProperty() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a GitHub organization custom property definition. Custom properties defined here can later be assigned values on individual repositories.",

		CreateContext: resourceGithubOrganizationRepositoryCustomPropertyCreate,
		ReadContext:   resourceGithubOrganizationRepositoryCustomPropertyRead,
		UpdateContext: resourceGithubOrganizationRepositoryCustomPropertyUpdate,
		DeleteContext: resourceGithubOrganizationRepositoryCustomPropertyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubOrganizationRepositoryCustomPropertyImport,
		},

		CustomizeDiff: customdiff.All(resourceGithubOrganizationRepositoryCustomPropertyDiff),

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"property_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the custom property.",
			},
			"value_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      fmt.Sprintf("Type of the custom property. One of: %v.", organizationCustomPropertyValueTypes),
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(organizationCustomPropertyValueTypes, false)),
			},
			"required": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the custom property must be set on every repository. GitHub may reject `required = true` unless a `default_value` is also provided.",
			},
			"default_value": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Default value applied to repositories that do not explicitly set the property. Exactly one element for the `string`, `single_select`, `true_false` and `url` types; one or more for `multi_select`.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Short description of the custom property.",
			},
			"allowed_values": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Allowed values for `single_select` and `multi_select` property types. Must be omitted for other types.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"values_editable_by": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Description:      fmt.Sprintf("Who can edit values of this property on repositories. One of: %v. Defaults to `org_actors` server-side.", organizationCustomPropertyValuesEditableBy),
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(organizationCustomPropertyValuesEditableBy, false)),
			},
		},
	}
}

func resourceGithubOrganizationRepositoryCustomPropertyDiff(ctx context.Context, d *schema.ResourceDiff, _ any) error {
	if !d.NewValueKnown("value_type") {
		return nil
	}

	valueType := github.PropertyValueType(d.Get("value_type").(string))
	selectType := valueType == github.PropertyValueTypeSingleSelect || valueType == github.PropertyValueTypeMultiSelect

	if d.NewValueKnown("allowed_values") {
		allowedValues, _ := d.Get("allowed_values").([]any)

		if selectType && len(allowedValues) == 0 {
			return fmt.Errorf("allowed_values is required when value_type is %q", valueType)
		}
		if !selectType && len(allowedValues) > 0 {
			return fmt.Errorf("allowed_values must not be set when value_type is %q", valueType)
		}
	}

	// Only multi_select accepts a list-valued default; every other type is scalar.
	if d.NewValueKnown("default_value") && valueType != github.PropertyValueTypeMultiSelect {
		if defaultValue, _ := d.Get("default_value").([]any); len(defaultValue) > 1 {
			return fmt.Errorf("default_value must contain at most one element when value_type is %q, got %d", valueType, len(defaultValue))
		}
	}

	return nil
}

func buildOrganizationRepositoryCustomProperty(d *schema.ResourceData) *github.CustomProperty {
	propertyName := d.Get("property_name").(string)
	valueType := github.PropertyValueType(d.Get("value_type").(string))
	required := d.Get("required").(bool)
	description := d.Get("description").(string)

	cp := &github.CustomProperty{
		PropertyName: &propertyName,
		ValueType:    valueType,
		Required:     &required,
		Description:  &description,
	}

	if v, ok := d.GetOk("default_value"); ok {
		if defaultValue := expandStringList(v.([]any)); len(defaultValue) > 0 {
			// Only multi_select sends an array; the other types send a bare string.
			switch valueType {
			case github.PropertyValueTypeMultiSelect:
				cp.DefaultValue = defaultValue
			default:
				cp.DefaultValue = defaultValue[0]
			}
		}
	}

	if v, ok := d.GetOk("allowed_values"); ok {
		cp.AllowedValues = expandStringList(v.([]any))
	}

	if v, ok := d.GetOk("values_editable_by"); ok {
		s := v.(string)
		cp.ValuesEditableBy = &s
	}

	return cp
}

// flattenOrganizationRepositoryCustomPropertyDefaultValue normalises the
// polymorphic default_value returned by the API into a list of strings. The
// wire type depends on value_type: multi_select is an array, true_false is a
// stringified bool and the rest are plain strings.
func flattenOrganizationRepositoryCustomPropertyDefaultValue(cp *github.CustomProperty) ([]string, error) {
	if cp.DefaultValue == nil {
		return nil, nil
	}

	switch cp.ValueType {
	case github.PropertyValueTypeMultiSelect:
		if v, ok := cp.DefaultValueStrings(); ok {
			return v, nil
		}
	case github.PropertyValueTypeTrueFalse:
		if v, ok := cp.DefaultValueBool(); ok {
			return []string{strconv.FormatBool(v)}, nil
		}
	default:
		if v, ok := cp.DefaultValueString(); ok {
			return []string{v}, nil
		}
	}

	return nil, fmt.Errorf("default_value %#v could not be parsed for value_type %q", cp.DefaultValue, cp.ValueType)
}

func resourceGithubOrganizationRepositoryCustomPropertyCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	client := meta.v3client
	owner := meta.name
	propertyName := d.Get("property_name").(string)

	tflog.Debug(ctx, "Creating organization custom property", map[string]any{"org": owner, "property": propertyName})

	cp, _, err := client.Organizations.CreateOrUpdateCustomProperty(ctx, owner, propertyName, buildOrganizationRepositoryCustomProperty(d))
	if err != nil {
		return diag.Errorf("error creating organization custom property %q: %v", propertyName, err)
	}

	if cp.GetPropertyName() == "" {
		return diag.Errorf("organization %q returned a custom property with an empty name when creating %q", owner, propertyName)
	}

	defaultValue, err := flattenOrganizationRepositoryCustomPropertyDefaultValue(cp)
	if err != nil {
		return diag.Errorf("error reading organization custom property %q: %v", propertyName, err)
	}

	d.SetId(cp.GetPropertyName())
	if err := d.Set("property_name", cp.GetPropertyName()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("value_type", string(cp.ValueType)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("required", cp.GetRequired()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default_value", defaultValue); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", cp.GetDescription()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_values", cp.AllowedValues); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("values_editable_by", cp.GetValuesEditableBy()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationRepositoryCustomPropertyRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	client := meta.v3client
	owner := meta.name
	propertyName := d.Get("property_name").(string)

	cp, _, err := client.Organizations.GetCustomProperty(ctx, owner, propertyName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == 404 {
			tflog.Info(ctx, "Removing organization custom property from state because it no longer exists", map[string]any{"org": owner, "property": propertyName})
			d.SetId("")
			return nil
		}
		return diag.Errorf("error reading organization custom property %q: %v", propertyName, err)
	}

	if cp.GetPropertyName() == "" {
		return diag.Errorf("organization %q returned a custom property with an empty name when reading %q", owner, propertyName)
	}

	switch cp.ValueType {
	case github.PropertyValueTypeSingleSelect, github.PropertyValueTypeMultiSelect:
	default:
		cp.AllowedValues = nil
	}

	defaultValue, err := flattenOrganizationRepositoryCustomPropertyDefaultValue(cp)
	if err != nil {
		return diag.Errorf("error reading organization custom property %q: %v", propertyName, err)
	}

	d.SetId(cp.GetPropertyName())
	if err := d.Set("property_name", cp.GetPropertyName()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("value_type", string(cp.ValueType)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("required", cp.GetRequired()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default_value", defaultValue); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", cp.GetDescription()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_values", cp.AllowedValues); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("values_editable_by", cp.GetValuesEditableBy()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationRepositoryCustomPropertyUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	client := meta.v3client
	owner := meta.name
	propertyName := d.Get("property_name").(string)

	tflog.Debug(ctx, "Updating organization custom property", map[string]any{"org": owner, "property": propertyName})

	cp, _, err := client.Organizations.CreateOrUpdateCustomProperty(ctx, owner, propertyName, buildOrganizationRepositoryCustomProperty(d))
	if err != nil {
		return diag.Errorf("error updating organization custom property %q: %v", propertyName, err)
	}

	if cp.GetPropertyName() == "" {
		return diag.Errorf("organization %q returned a custom property with an empty name when updating %q", owner, propertyName)
	}

	defaultValue, err := flattenOrganizationRepositoryCustomPropertyDefaultValue(cp)
	if err != nil {
		return diag.Errorf("error reading organization custom property %q: %v", propertyName, err)
	}

	d.SetId(cp.GetPropertyName())
	if err := d.Set("property_name", cp.GetPropertyName()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("value_type", string(cp.ValueType)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("required", cp.GetRequired()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default_value", defaultValue); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", cp.GetDescription()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_values", cp.AllowedValues); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("values_editable_by", cp.GetValuesEditableBy()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationRepositoryCustomPropertyDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	client := meta.v3client
	owner := meta.name
	propertyName := d.Get("property_name").(string)

	tflog.Debug(ctx, "Deleting organization custom property", map[string]any{"org": owner, "property": propertyName})

	if _, err := client.Organizations.RemoveCustomProperty(ctx, owner, propertyName); err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == 404 {
			return nil
		}
		return diag.Errorf("error deleting organization custom property %q: %v", propertyName, err)
	}

	return nil
}

func resourceGithubOrganizationRepositoryCustomPropertyImport(ctx context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
	propertyName := d.Id()
	if propertyName == "" {
		return nil, errors.New("custom property name must not be empty")
	}

	// Read looks the property up by attribute, so seed it from the import ID.
	if err := d.Set("property_name", propertyName); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
