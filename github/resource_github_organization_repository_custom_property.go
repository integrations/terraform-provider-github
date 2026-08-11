package github

import (
	"context"
	"errors"
	"fmt"

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
			StateContext: schema.ImportStatePassthroughContext,
		},

		CustomizeDiff: customdiff.All(resourceGithubOrganizationRepositoryCustomPropertyDiff),

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
				Description: "Whether the custom property must be set on every repository. When true, `default_value` must be provided.",
			},
			"default_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Default value applied to repositories that do not explicitly set the property.",
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
	if !d.NewValueKnown("value_type") || !d.NewValueKnown("allowed_values") {
		return nil
	}

	valueType := github.PropertyValueType(d.Get("value_type").(string))
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
		s := v.(string)
		cp.DefaultValue = &s
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
		return diag.FromErr(fmt.Errorf("error creating organization custom property %q: %w", propertyName, err))
	}

	defaultValue, _ := cp.DefaultValueString()
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
	propertyName := d.Id()

	cp, _, err := client.Organizations.GetCustomProperty(ctx, owner, propertyName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == 404 {
			tflog.Info(ctx, "Removing organization custom property from state because it no longer exists", map[string]any{"org": owner, "property": propertyName})
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error reading organization custom property %q: %w", propertyName, err))
	}

	switch cp.ValueType {
	case github.PropertyValueTypeSingleSelect, github.PropertyValueTypeMultiSelect:
	default:
		cp.AllowedValues = nil
	}

	defaultValue, _ := cp.DefaultValueString()
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
		return diag.FromErr(fmt.Errorf("error updating organization custom property %q: %w", propertyName, err))
	}

	defaultValue, _ := cp.DefaultValueString()
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
		return diag.FromErr(fmt.Errorf("error deleting organization custom property %q: %w", propertyName, err))
	}

	return nil
}
