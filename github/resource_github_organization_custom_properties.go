package github

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/v85/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubOrganizationCustomProperties() *schema.Resource {
	return &schema.Resource{
		Description:   "Creates and manages a custom property for a GitHub Organization.",
		CreateContext: resourceGithubCustomPropertiesCreate,
		ReadContext:   resourceGithubCustomPropertiesRead,
		UpdateContext: resourceGithubCustomPropertiesUpdate,
		DeleteContext: resourceGithubCustomPropertiesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubCustomPropertiesImport,
		},

		Schema: map[string]*schema.Schema{
			"property_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the custom property",
			},
			"value_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "The type of the custom property. Can be one of: 'string', 'single_select', 'multi_select', 'true_false', or 'url'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{string(github.PropertyValueTypeString), string(github.PropertyValueTypeSingleSelect), string(github.PropertyValueTypeMultiSelect), string(github.PropertyValueTypeTrueFalse), string(github.PropertyValueTypeURL)}, false)),
			},
			"required": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the custom property is required",
			},
			"default_value": {
				Type:        schema.TypeString,
				Description: "The default value of the custom property",
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

// buildCustomProperty constructs a github.CustomProperty from the resource data.
func buildCustomProperty(d *schema.ResourceData) (*github.CustomProperty, error) {
	propertyName := d.Get("property_name").(string)
	valueType := github.PropertyValueType(d.Get("value_type").(string))
	required := d.Get("required").(bool)
	description := d.Get("description").(string)

	customProperty := &github.CustomProperty{
		PropertyName: &propertyName,
		ValueType:    valueType,
		Required:     &required,
		Description:  &description,
	}

	// Set default value if provided
	if v, ok := d.GetOk("default_value"); ok {
		defaultValue := v.(string)
		customProperty.DefaultValue = &defaultValue
	}

	// Set allowed values if provided (only valid for select types)
	if v, ok := d.GetOk("allowed_values"); ok {
		allowedValues := expandStringList(v.([]any))
		if valueType == github.PropertyValueTypeSingleSelect || valueType == github.PropertyValueTypeMultiSelect {
			customProperty.AllowedValues = allowedValues
		} else {
			return nil, fmt.Errorf("allowed_values can only be set for single_select or multi_select value types")
		}
	}

	// Validate that allowed_values is provided for select types
	if (valueType == github.PropertyValueTypeSingleSelect || valueType == github.PropertyValueTypeMultiSelect) && len(customProperty.AllowedValues) == 0 {
		return nil, fmt.Errorf("allowed_values is required for %s value type", valueType)
	}

	if val, ok := d.GetOk("values_editable_by"); ok {
		str := val.(string)
		customProperty.ValuesEditableBy = &str
	}

	return customProperty, nil
}

// setCustomPropertyState sets all resource data fields from a CustomProperty API response.
func setCustomPropertyState(d *schema.ResourceData, customProperty *github.CustomProperty) {
	// TODO: Add support for other types of default values
	defaultValue, _ := customProperty.DefaultValueString()

	d.SetId(*customProperty.PropertyName)
	_ = d.Set("allowed_values", customProperty.AllowedValues)
	_ = d.Set("default_value", defaultValue)
	_ = d.Set("description", customProperty.Description)
	_ = d.Set("property_name", customProperty.PropertyName)
	_ = d.Set("required", customProperty.Required)
	_ = d.Set("value_type", string(customProperty.ValueType))
	_ = d.Set("values_editable_by", customProperty.ValuesEditableBy)
}

func resourceGithubCustomPropertiesCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	ownerName := meta.(*Owner).name

	customProperty, err := buildCustomProperty(d)
	if err != nil {
		return diag.FromErr(err)
	}

	propertyName := d.Get("property_name").(string)
	customProperty, _, err = client.Organizations.CreateOrUpdateCustomProperty(ctx, ownerName, propertyName, customProperty)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating organization custom property %s: %w", propertyName, err))
	}

	setCustomPropertyState(d, customProperty)

	return nil
}

func resourceGithubCustomPropertiesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	ownerName := meta.(*Owner).name

	propertyName := d.Id()
	if pn, ok := d.GetOk("property_name"); ok {
		propertyName = pn.(string)
	}

	customProperty, resp, err := client.Organizations.GetCustomProperty(ctx, ownerName, propertyName)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			log.Printf("[WARN] Removing organization custom property %s from state because it no longer exists in GitHub", propertyName)
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error reading organization custom property %s: %w", propertyName, err))
	}

	setCustomPropertyState(d, customProperty)

	return nil
}

func resourceGithubCustomPropertiesUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	ownerName := meta.(*Owner).name

	customProperty, err := buildCustomProperty(d)
	if err != nil {
		return diag.FromErr(err)
	}

	propertyName := d.Get("property_name").(string)
	customProperty, _, err = client.Organizations.CreateOrUpdateCustomProperty(ctx, ownerName, propertyName, customProperty)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating organization custom property %s: %w", propertyName, err))
	}

	setCustomPropertyState(d, customProperty)

	return nil
}

func resourceGithubCustomPropertiesDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	ownerName := meta.(*Owner).name
	propertyName := d.Get("property_name").(string)

	_, err := client.Organizations.RemoveCustomProperty(ctx, ownerName, propertyName)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting organization custom property %s: %w", propertyName, err))
	}

	return nil
}

func resourceGithubCustomPropertiesImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	if err := d.Set("property_name", d.Id()); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
