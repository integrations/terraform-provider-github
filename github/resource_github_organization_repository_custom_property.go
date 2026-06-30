package github

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/google/go-github/v88/github"
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

func setOrganizationRepositoryCustomPropertyState(d *schema.ResourceData, cp *github.CustomProperty) error {
	defaultValue, _ := cp.DefaultValueString()

	d.SetId(cp.GetPropertyName())

	for _, set := range []struct {
		key   string
		value any
	}{
		{"property_name", cp.GetPropertyName()},
		{"value_type", string(cp.ValueType)},
		{"required", cp.GetRequired()},
		{"description", cp.GetDescription()},
		{"default_value", defaultValue},
		{"allowed_values", cp.AllowedValues},
		{"values_editable_by", cp.GetValuesEditableBy()},
	} {
		if err := d.Set(set.key, set.value); err != nil {
			return err
		}
	}

	return nil
}

func resourceGithubOrganizationRepositoryCustomPropertyCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	propertyName := d.Get("property_name").(string)

	tflog.Debug(ctx, "Creating organization custom property", map[string]any{"org": orgName, "property": propertyName})

	cp, _, err := client.Organizations.CreateOrUpdateCustomProperty(ctx, orgName, propertyName, buildOrganizationRepositoryCustomProperty(d))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating organization custom property %q: %w", propertyName, err))
	}

	if err := setOrganizationRepositoryCustomPropertyState(d, cp); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationRepositoryCustomPropertyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	propertyName := d.Id()

	cp, _, err := client.Organizations.GetCustomProperty(ctx, orgName, propertyName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == 404 {
			tflog.Info(ctx, "Removing organization custom property from state because it no longer exists", map[string]any{"org": orgName, "property": propertyName})
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error reading organization custom property %q: %w", propertyName, err))
	}

	// Sentinel: the API ignores allowed_values for non-select types, so don't
	// import a phantom empty list into state. Only persist it when meaningful.
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

func resourceGithubOrganizationRepositoryCustomPropertyUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	propertyName := d.Get("property_name").(string)

	tflog.Debug(ctx, "Updating organization custom property", map[string]any{"org": orgName, "property": propertyName})

	cp, _, err := client.Organizations.CreateOrUpdateCustomProperty(ctx, orgName, propertyName, buildOrganizationRepositoryCustomProperty(d))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating organization custom property %q: %w", propertyName, err))
	}

	if err := setOrganizationRepositoryCustomPropertyState(d, cp); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationRepositoryCustomPropertyDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	propertyName := d.Get("property_name").(string)

	tflog.Debug(ctx, "Deleting organization custom property", map[string]any{"org": orgName, "property": propertyName})

	if _, err := client.Organizations.RemoveCustomProperty(ctx, orgName, propertyName); err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == 404 {
			return nil
		}
		return diag.FromErr(fmt.Errorf("error deleting organization custom property %q: %w", propertyName, err))
	}

	return nil
}
