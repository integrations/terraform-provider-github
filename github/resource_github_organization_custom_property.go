package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	ORG_ACTORS          = "org_actors"
	ORG_AND_REPO_ACTORS = "org_and_repo_actors"
)

// This file implements a custom schema and (un)marshalling logic for GitHub organization custom properties.
// The upstream GitHub SDK (github.CustomProperty) supports only a single string value for the `default_value` field,
// which is not compatible with the MULTI_SELECT property type that requires a list of default values.
// To work around this limitation, we define a customPropertyExtended struct that embeds github.CustomProperty
// and overrides the `default_value` field with a flexible interface that can handle string, []string, or null values.
// Custom JSON marshalling and unmarshalling methods are implemented to ensure compatibility with the GitHub API
// while preserving type correctness and Terraform expectations.
type customPropertyExtended struct {
	github.CustomProperty

	// Overrides the original DefaultValue to support string, null, and []string.
	DefaultValueOverride interface{} `json:"default_value,omitempty"`
}

func (c *customPropertyExtended) MarshalJSON() ([]byte, error) {
	// Marshal base struct to map
	base, err := json.Marshal(c.CustomProperty)
	if err != nil {
		return nil, err
	}

	var baseMap map[string]interface{}
	if err := json.Unmarshal(base, &baseMap); err != nil {
		return nil, err
	}

	// Override default_value
	if c.DefaultValueOverride != nil {
		baseMap["default_value"] = c.DefaultValueOverride
	}

	return json.Marshal(baseMap)
}

func (c *customPropertyExtended) UnmarshalJSON(data []byte) error {
	// Unmarshal the JSON into a map to isolate default_value
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	// Extract and remove default_value before unmarshalling the embedded struct
	var rawDefault json.RawMessage
	if v, ok := m["default_value"]; ok {
		rawDefault = v
		delete(m, "default_value")
	}

	// Re-marshal the map without default_value
	sanitized, err := json.Marshal(m)
	if err != nil {
		return err
	}

	// Unmarshal the sanitized JSON into the embedded CustomProperty struct
	if err := json.Unmarshal(sanitized, &c.CustomProperty); err != nil {
		return err
	}

	// Manually unmarshal default_value based on its type
	if len(rawDefault) > 0 {
		// Try to unmarshal as a string
		var s string
		if err := json.Unmarshal(rawDefault, &s); err == nil {
			c.DefaultValueOverride = s
			return nil
		}

		// Try to unmarshal as a []string
		var list []string
		if err := json.Unmarshal(rawDefault, &list); err == nil {
			c.DefaultValueOverride = list
			return nil
		}

		// Handle null value
		if string(rawDefault) == "null" {
			c.DefaultValueOverride = nil
			return nil
		}

		return fmt.Errorf("invalid format for default_value: %s", string(rawDefault))
	}

	return nil
}

func (c *customPropertyExtended) GetDefaultValueOverride() ([]string, error) {
	switch value := c.DefaultValueOverride.(type) {
	case string:
		return []string{value}, nil
	case []string:
		return value, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("custom property value couldn't be parsed as a string or a list of strings: %s", value)
	}
}

func resourceGithubOrganizationCustomProperty() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubOrganizationCustomPropertyCreateOrUpdate,
		Update: resourceGithubOrganizationCustomPropertyCreateOrUpdate,
		Read:   resourceGithubOrganizationCustomPropertyRead,
		Delete: resourceGithubOrganizationCustomPropertyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				if err := resourceGithubOrganizationCustomPropertyRead(d, meta); err != nil {
					return nil, err
				}
				return []*schema.ResourceData{d}, nil
			},
		},
		CustomizeDiff: func(ctx context.Context, diff *schema.ResourceDiff, meta interface{}) error {
			// Validate the relationship between required and default_value.
			// If the property is marked as required, a non-empty default value must be provided.
			// If the property is not required, then a default value must not be set.
			required := diff.Get("required").(bool)
			defaultValue := expandStringList(diff.Get("default_value").(*schema.Set).List())
			if required {
				if len(defaultValue) == 0 {
					return errors.New("default_value can not be empty")
				}
			} else {
				if len(defaultValue) != 0 {
					return errors.New("default_value is only allowed if required is true")
				}
			}

			// Validate that for MULTI_SELECT and SINGLE_SELECT types,
			// all default values must be included in the list of allowed values.
			propertyType := diff.Get("type").(string)
			allowedValues := expandStringList(diff.Get("allowed_values").(*schema.Set).List())
			if propertyType == MULTI_SELECT || propertyType == SINGLE_SELECT {
				if !isSubset(defaultValue, allowedValues) {
					return errors.New("default_value must be a subset of allowed_values")
				}
			}

			// Validate that for STRING or TRUE_FALSE properties, no allowed values are permitted.
			// STRING type should not define allowed_values, as it's meant to be free-form text.
			if propertyType == STRING || propertyType == TRUE_FALSE {
				if len(allowedValues) != 0 {
					return errors.New("allowed_values must be empty when type is STRING or TRUE_FALSE")
				}
			}

			// Validate that for SINGLE_SELECT and STRING properties, at most one default value is permitted.
			// An empty list or a single option is allowed, but more than one value is not supported in this context.
			if propertyType == SINGLE_SELECT || propertyType == STRING {
				if len(defaultValue) > 1 {
					return errors.New("default_value must contain zero or one item when type is SINGLE_SELECT or STRING")
				}
			}

			// Validate that for TRUE_FALSE properties, at most one default value is permitted,
			// and if provided, it must be either "true" or "false".
			if propertyType == TRUE_FALSE && len(defaultValue) == 1 {
				if defaultValue[0] != "true" && defaultValue[0] != "false" {
					return errors.New("default_value must be either \"true\" or \"false\" when type is TRUE_FALSE")
				}
			}

			return nil
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the custom property.",
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
			},
			"type": {
				Description:      "Type of the custom property",
				ForceNew:         true,
				Required:         true,
				Type:             schema.TypeString,
				ValidateDiagFunc: toDiagFunc(validation.StringInSlice([]string{SINGLE_SELECT, MULTI_SELECT, STRING, TRUE_FALSE}, false), "type"),
			},
			"required": {
				Default:     false,
				Description: "Whether the property is required.",
				Optional:    true,
				Type:        schema.TypeBool,
			},
			"default_value": {
				Description: "Default value of the property if required.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				MinItems: 1,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"description": {
				Description: "Short description of the property.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"allowed_values": {
				Description: "An ordered list of the allowed values of the property. The property can have up to 200 allowed values.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				MaxItems: 200,
				Optional: true,
				Type:     schema.TypeSet,
			},
			"values_editable_by": {
				Default:          ORG_ACTORS,
				Description:      "Who can edit the values of the property.",
				Optional:         true,
				Type:             schema.TypeString,
				ValidateDiagFunc: toDiagFunc(validation.StringInSlice([]string{ORG_ACTORS, ORG_AND_REPO_ACTORS}, false), "values_editable_by"),
			},
		},
	}
}

func resourceGithubOrganizationCustomPropertyCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	propertyName := d.Get("name").(string)
	propertyType := d.Get("type").(string)
	propertyRequired := d.Get("required").(bool)
	propertyDefaultValue := expandStringList(d.Get("default_value").(*schema.Set).List())
	propertyDescription := d.Get("description").(string)
	propertyAllowedValues := expandStringList(d.Get("allowed_values").(*schema.Set).List())
	propertyValuesEditableBy := d.Get("values_editable_by").(string)

	customProperty := customPropertyExtended{
		CustomProperty: github.CustomProperty{
			PropertyName:     &propertyName,
			ValueType:        propertyType,
			Required:         &propertyRequired,
			Description:      &propertyDescription,
			AllowedValues:    propertyAllowedValues,
			ValuesEditableBy: &propertyValuesEditableBy,
		},
	}

	if len(propertyDefaultValue) > 0 {
		// The propertyDefaultValue can either be a list of strings or a string
		switch propertyType {
		case SINGLE_SELECT, TRUE_FALSE, STRING:
			customProperty.DefaultValueOverride = &propertyDefaultValue[0]
		case MULTI_SELECT:
			customProperty.DefaultValueOverride = propertyDefaultValue
		default:
			return fmt.Errorf("custom property type is not valid: %v", propertyType)
		}
	}

	u := fmt.Sprintf("orgs/%v/properties/schema/%v", orgName, propertyName)
	req, err := client.NewRequest("PUT", u, customProperty)
	if err != nil {
		return err
	}

	_, err = client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(orgName, propertyName))

	return resourceGithubOrganizationCustomPropertyRead(d, meta)
}

func resourceGithubOrganizationCustomPropertyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	orgName, propertyName, err := parseTwoPartID(d.Id(), "orgName", "propertyName")
	if err != nil {
		return err
	}

	u := fmt.Sprintf("orgs/%v/properties/schema/%v", orgName, propertyName)
	req, err := client.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}

	var customProperty *customPropertyExtended
	_, err = client.Do(ctx, req, &customProperty)
	if err != nil {
		return err
	}

	defaultValue, err := customProperty.GetDefaultValueOverride()
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(orgName, customProperty.GetPropertyName()))
	d.Set("name", customProperty.GetPropertyName())
	d.Set("type", customProperty.ValueType)
	d.Set("required", customProperty.Required)
	d.Set("default_value", defaultValue)
	d.Set("description", customProperty.GetDescription())
	d.Set("allowed_values", customProperty.AllowedValues)
	d.Set("values_editable_by", customProperty.GetValuesEditableBy())

	return nil
}

func resourceGithubOrganizationCustomPropertyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	orgName, propertyName, err := parseTwoPartID(d.Id(), "orgName", "propertyName")
	if err != nil {
		return err
	}

	_, err = client.Organizations.RemoveCustomProperty(ctx, orgName, propertyName)
	if err != nil {
		return err
	}

	return nil
}
