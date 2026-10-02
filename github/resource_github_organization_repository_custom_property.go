package github

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

		CustomizeDiff: resourceGithubOrganizationRepositoryCustomPropertyDiff,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"property_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Name of the custom property.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
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
				Description: "Whether the custom property must be set on every repository. A `default_value` is required when this is `true`.",
			},
			"default_value": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Default value applied to repositories that do not explicitly set the property. Exactly one element for the `string`, `single_select`, `true_false` and `url` types; one or more for `multi_select`. Once set, a default cannot be removed via the API, only changed.",
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Short description of the custom property.",
			},
			// Deliberately not Computed: an omitted Optional+Computed list is
			// unknown at plan time, which would make the cross-field validation
			// in CustomizeDiff silently skip itself. Nothing needs to be read
			// back here either -- select types always set it in config, and Read
			// clears it for the other types.
			"allowed_values": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Allowed values for `single_select` and `multi_select` property types. Must be omitted for other types.",
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				},
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
	if err := validateOrganizationRepositoryCustomPropertyAllowedValues(d, valueType); err != nil {
		return err
	}

	return validateOrganizationRepositoryCustomPropertyDefaultValue(d, valueType)
}

func validateOrganizationRepositoryCustomPropertyAllowedValues(d *schema.ResourceDiff, valueType github.PropertyValueType) error {
	if !d.NewValueKnown("allowed_values") {
		return nil
	}

	allowedValues, _ := d.Get("allowed_values").([]any)
	selectType := valueType == github.PropertyValueTypeSingleSelect || valueType == github.PropertyValueTypeMultiSelect
	if selectType && len(allowedValues) == 0 {
		return fmt.Errorf("allowed_values is required when value_type is %q", valueType)
	}
	if !selectType && len(allowedValues) > 0 {
		return fmt.Errorf("allowed_values must not be set when value_type is %q", valueType)
	}

	return nil
}

func validateOrganizationRepositoryCustomPropertyDefaultValue(d *schema.ResourceDiff, valueType github.PropertyValueType) error {
	defaultValueKnown := d.NewValueKnown("default_value")
	defaultValue, _ := d.Get("default_value").([]any)
	required := d.NewValueKnown("required") && d.Get("required").(bool)
	if err := validateOrganizationRepositoryCustomPropertyDefaultValueUpdate(d, defaultValueKnown, defaultValue); err != nil {
		return err
	}
	if err := validateRequiredOrganizationRepositoryCustomPropertyDefaultValue(required, defaultValueKnown, defaultValue); err != nil {
		return err
	}
	if !defaultValueKnown {
		return nil
	}

	if valueType != github.PropertyValueTypeMultiSelect && len(defaultValue) > 1 {
		return fmt.Errorf("default_value must contain at most one element when value_type is %q, got %d", valueType, len(defaultValue))
	}

	if valueType == github.PropertyValueTypeTrueFalse {
		for _, v := range defaultValue {
			if s, _ := v.(string); s != "true" && s != "false" {
				return fmt.Errorf("default_value must be %q or %q when value_type is %q, got %q", "true", "false", valueType, s)
			}
		}
	}

	selectType := valueType == github.PropertyValueTypeSingleSelect || valueType == github.PropertyValueTypeMultiSelect
	if selectType && d.NewValueKnown("allowed_values") {
		allowedValues, _ := d.Get("allowed_values").([]any)
		if err := validateSelectPropertyDefaultValue(valueType, expandStringList(defaultValue), expandStringList(allowedValues)); err != nil {
			return err
		}
	}

	return nil
}

func validateOrganizationRepositoryCustomPropertyDefaultValueUpdate(d *schema.ResourceDiff, defaultValueKnown bool, defaultValue []any) error {
	if !defaultValueKnown || !d.HasChange("default_value") {
		return nil
	}

	oldValue, _ := d.GetChange("default_value")
	oldDefaultValue, _ := oldValue.([]any)
	return validateOrganizationRepositoryCustomPropertyDefaultValueNotRemoved(oldDefaultValue, defaultValue)
}

func validateOrganizationRepositoryCustomPropertyDefaultValueNotRemoved(oldValue, newValue []any) error {
	if len(oldValue) > 0 && len(newValue) == 0 {
		return fmt.Errorf("default_value cannot be removed once set; the GitHub API only allows changing it")
	}

	return nil
}

func validateRequiredOrganizationRepositoryCustomPropertyDefaultValue(required, defaultValueKnown bool, defaultValue []any) error {
	if required && (!defaultValueKnown || len(defaultValue) == 0) {
		return fmt.Errorf("default_value is required when required is true")
	}

	return nil
}

func validateSelectPropertyDefaultValue(valueType github.PropertyValueType, defaultValues, allowedValues []string) error {
	allowed := make(map[string]struct{}, len(allowedValues))
	for _, value := range allowedValues {
		allowed[value] = struct{}{}
	}

	for _, value := range defaultValues {
		if _, ok := allowed[value]; !ok {
			return fmt.Errorf("default_value %q must be one of allowed_values when value_type is %q", value, valueType)
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

	return setOrganizationRepositoryCustomPropertyState(d, cp)
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
		return diag.Errorf(organizationCustomPropertyReadErrorFormat, propertyName, err)
	}

	if cp.GetPropertyName() == "" {
		return diag.Errorf("organization %q returned a custom property with an empty name when reading %q", owner, propertyName)
	}

	return setOrganizationRepositoryCustomPropertyState(d, cp)
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

	return setOrganizationRepositoryCustomPropertyState(d, cp)
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
