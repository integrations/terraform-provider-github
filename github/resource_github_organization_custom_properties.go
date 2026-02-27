package github

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubOrganizationCustomProperties() *schema.Resource {
	return &schema.Resource{
		Description: "Creates and manages a custom property for a GitHub Organization.",
		Create:      resourceGithubCustomPropertiesCreate,
		Read:        resourceGithubCustomPropertiesRead,
		Update:      resourceGithubCustomPropertiesUpdate,
		Delete:      resourceGithubCustomPropertiesDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGithubCustomPropertiesImport,
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

func resourceGithubCustomPropertiesCreate(d *schema.ResourceData, meta any) error {
	if err := checkOrganization(meta); err != nil {
		return err
	}

	ctx := context.Background()
	client := meta.(*Owner).v3client
	ownerName := meta.(*Owner).name

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
			return fmt.Errorf("allowed_values can only be set for single_select or multi_select value types")
		}
	}

	// Validate that allowed_values is provided for select types
	if (valueType == github.PropertyValueTypeSingleSelect || valueType == github.PropertyValueTypeMultiSelect) && len(customProperty.AllowedValues) == 0 {
		return fmt.Errorf("allowed_values is required for %s value type", valueType)
	}

	if val, ok := d.GetOk("values_editable_by"); ok {
		str := val.(string)
		customProperty.ValuesEditableBy = &str
	}

	customProperty, _, err := client.Organizations.CreateOrUpdateCustomProperty(ctx, ownerName, propertyName, customProperty)
	if err != nil {
		return fmt.Errorf("error creating organization custom property %s: %w", propertyName, err)
	}

	d.SetId(*customProperty.PropertyName)
	return resourceGithubCustomPropertiesRead(d, meta)
}

func resourceGithubCustomPropertiesRead(d *schema.ResourceData, meta any) error {
	if err := checkOrganization(meta); err != nil {
		return err
	}

	ctx := context.Background()
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
		return fmt.Errorf("error reading organization custom property %s: %w", propertyName, err)
	}

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

	return nil
}

func resourceGithubCustomPropertiesUpdate(d *schema.ResourceData, meta any) error {
	// Create uses the same upsert API, and already calls Read at the end
	return resourceGithubCustomPropertiesCreate(d, meta)
}

func resourceGithubCustomPropertiesDelete(d *schema.ResourceData, meta any) error {
	if err := checkOrganization(meta); err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	ownerName := meta.(*Owner).name
	propertyName := d.Get("property_name").(string)

	_, err := client.Organizations.RemoveCustomProperty(context.Background(), ownerName, propertyName)
	if err != nil {
		return fmt.Errorf("error deleting organization custom property %s: %w", propertyName, err)
	}

	return nil
}

func resourceGithubCustomPropertiesImport(d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	if err := d.Set("property_name", d.Id()); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
