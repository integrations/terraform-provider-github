package github

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryCustomProperties() *schema.Resource {
	return &schema.Resource{
		Description: "Manages custom properties for a GitHub repository. This resource allows you to set multiple custom property values on a single repository in a single resource block, with in-place updates when values change.",
		Create:      resourceGithubRepositoryCustomPropertiesCreateOrUpdate,
		Read:        resourceGithubRepositoryCustomPropertiesRead,
		Update:      resourceGithubRepositoryCustomPropertiesCreateOrUpdate,
		Delete:      resourceGithubRepositoryCustomPropertiesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryCustomPropertiesImport,
		},

		Schema: map[string]*schema.Schema{
			"repository_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the repository.",
			},
			"property": {
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Description: "Set of custom property values for this repository.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the custom property (must be defined at the organization level).",
						},
						"value": {
							Type:        schema.TypeSet,
							Required:    true,
							MinItems:    1,
							Description: "Value(s) of the custom property. For multi_select properties, multiple values can be specified.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
				Set: resourceGithubRepositoryCustomPropertiesHash,
			},
		},
	}
}

// resourceGithubRepositoryCustomPropertiesHash creates a hash for a property block
// using only the property name, so that value changes are detected as in-place
// updates rather than remove+add within the set.
func resourceGithubRepositoryCustomPropertiesHash(v any) int {
	raw := v.(map[string]any)
	name := raw["name"].(string)
	return schema.HashString(name)
}

func resourceGithubRepositoryCustomPropertiesCreateOrUpdate(d *schema.ResourceData, meta any) error {
	if err := checkOrganization(meta); err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	ctx := context.Background()
	owner := meta.(*Owner).name
	repoName := d.Get("repository_name").(string)
	properties := d.Get("property").(*schema.Set).List()

	// Get all organization custom property definitions to determine types
	orgProperties, _, err := client.Organizations.GetAllCustomProperties(ctx, owner)
	if err != nil {
		return fmt.Errorf("error reading organization custom property definitions: %w", err)
	}

	// Create a map of property names to their types
	propertyTypes := make(map[string]github.PropertyValueType)
	for _, prop := range orgProperties {
		if prop.PropertyName != nil {
			propertyTypes[*prop.PropertyName] = prop.ValueType
		}
	}

	// Build custom property values for this repository
	customProperties := make([]*github.CustomPropertyValue, 0, len(properties))

	for _, propBlock := range properties {
		propMap := propBlock.(map[string]any)
		propertyName := propMap["name"].(string)
		propertyValues := expandStringList(propMap["value"].(*schema.Set).List())

		propertyType, ok := propertyTypes[propertyName]
		if !ok {
			return fmt.Errorf("custom property %q is not defined at the organization level", propertyName)
		}

		customProperty := &github.CustomPropertyValue{
			PropertyName: propertyName,
		}

		switch propertyType {
		case github.PropertyValueTypeMultiSelect:
			customProperty.Value = propertyValues
		case github.PropertyValueTypeString, github.PropertyValueTypeSingleSelect,
			github.PropertyValueTypeTrueFalse, github.PropertyValueTypeURL:
			if len(propertyValues) > 0 {
				customProperty.Value = propertyValues[0]
			}
		default:
			return fmt.Errorf("unsupported property type %q for property %q", propertyType, propertyName)
		}

		customProperties = append(customProperties, customProperty)
	}

	_, err = client.Repositories.CreateOrUpdateCustomProperties(ctx, owner, repoName, customProperties)
	if err != nil {
		return fmt.Errorf("error setting custom properties for repository %s/%s: %w", owner, repoName, err)
	}

	d.SetId(buildTwoPartID(owner, repoName))

	return resourceGithubRepositoryCustomPropertiesRead(d, meta)
}

func resourceGithubRepositoryCustomPropertiesRead(d *schema.ResourceData, meta any) error {
	if err := checkOrganization(meta); err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	ctx := context.Background()

	owner, repoName, err := parseTwoPartID(d.Id(), "owner", "repository")
	if err != nil {
		return err
	}

	// Get current properties from state to know which ones we're managing.
	// On import this will be empty, which is handled below.
	propertiesFromState := d.Get("property").(*schema.Set).List()
	managedPropertyNames := make(map[string]bool)
	for _, propBlock := range propertiesFromState {
		propMap := propBlock.(map[string]any)
		managedPropertyNames[propMap["name"].(string)] = true
	}

	isImport := len(managedPropertyNames) == 0

	// Read actual properties from GitHub
	allCustomProperties, _, err := client.Repositories.GetAllCustomPropertyValues(ctx, owner, repoName)
	if err != nil {
		return fmt.Errorf("error reading custom properties for repository %s/%s: %w", owner, repoName, err)
	}

	// Build the property set — either all properties (import) or only managed ones
	managedProperties := make([]any, 0)
	for _, prop := range allCustomProperties {
		if !isImport && !managedPropertyNames[prop.PropertyName] {
			continue
		}

		// Skip properties with nil/null values (unset)
		if prop.Value == nil {
			continue
		}

		propertyValue, err := parseRepositoryCustomPropertyValueToStringSlice(prop)
		if err != nil {
			return fmt.Errorf("error parsing property %q for repository %s/%s: %w", prop.PropertyName, owner, repoName, err)
		}

		if len(propertyValue) == 0 {
			continue
		}

		managedProperties = append(managedProperties, map[string]any{
			"name":  prop.PropertyName,
			"value": propertyValue,
		})
	}

	// If no properties exist, remove resource from state
	if len(managedProperties) == 0 {
		log.Printf("[WARN] No custom properties found for %s/%s, removing from state", owner, repoName)
		d.SetId("")
		return nil
	}

	_ = d.Set("repository_name", repoName)
	_ = d.Set("property", managedProperties)

	return nil
}

func resourceGithubRepositoryCustomPropertiesDelete(d *schema.ResourceData, meta any) error {
	if err := checkOrganization(meta); err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	ctx := context.Background()

	owner, repoName, err := parseTwoPartID(d.Id(), "owner", "repository")
	if err != nil {
		return err
	}

	properties := d.Get("property").(*schema.Set).List()
	if len(properties) == 0 {
		return nil
	}

	// Set all managed properties to nil (removes them)
	customProperties := make([]*github.CustomPropertyValue, 0, len(properties))
	for _, propBlock := range properties {
		propMap := propBlock.(map[string]any)
		customProperties = append(customProperties, &github.CustomPropertyValue{
			PropertyName: propMap["name"].(string),
			Value:        nil,
		})
	}

	_, err = client.Repositories.CreateOrUpdateCustomProperties(ctx, owner, repoName, customProperties)
	if err != nil {
		return fmt.Errorf("error deleting custom properties for repository %s/%s: %w", owner, repoName, err)
	}

	return nil
}

func resourceGithubRepositoryCustomPropertiesImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	// Import ID format: owner/repo (using standard two-part ID)
	// On import, Read will detect empty state and import ALL properties
	parts := strings.SplitN(d.Id(), "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return nil, fmt.Errorf("invalid import ID %q, expected format: owner/repository", d.Id())
	}

	d.SetId(buildTwoPartID(parts[0], parts[1]))
	return []*schema.ResourceData{d}, nil
}
