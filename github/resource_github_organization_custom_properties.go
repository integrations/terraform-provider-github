package github

import (
	"context"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubOrganizationCustomProperties() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubCustomPropertiesCreate,
		Read:   resourceGithubCustomPropertiesRead,
		Update: resourceGithubCustomPropertiesUpdate,
		Delete: resourceGithubCustomPropertiesDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGithubCustomPropertiesImport,
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
	ctx := context.Background()
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

	customProperty, _, err := client.Organizations.CreateOrUpdateCustomProperty(ctx, ownerName, d.Get("property_name").(string), customProperty)
	if err != nil {
		return err
	}

	d.SetId(*customProperty.PropertyName)
	return resourceGithubCustomPropertiesRead(d, meta)
}

func resourceGithubCustomPropertiesRead(d *schema.ResourceData, meta any) error {
	ctx := context.Background()
	client := meta.(*Owner).v3client
	ownerName := meta.(*Owner).name

	customProperty, _, err := client.Organizations.GetCustomProperty(ctx, ownerName, d.Get("property_name").(string))
	if err != nil {
		return err
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
	if err := resourceGithubCustomPropertiesCreate(d, meta); err != nil {
		return err
	}
	return resourceGithubCustomPropertiesRead(d, meta)
}

func resourceGithubCustomPropertiesDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ownerName := meta.(*Owner).name

	_, err := client.Organizations.RemoveCustomProperty(context.Background(), ownerName, d.Get("property_name").(string))
	if err != nil {
		return err
	}

	return nil
}

func resourceGithubCustomPropertiesImport(d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	if err := d.Set("property_name", d.Id()); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
