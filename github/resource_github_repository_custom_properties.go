package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryCustomProperties() *schema.Resource {
	return &schema.Resource{
		Description: "Manages custom properties for a GitHub repository. This resource allows you to set multiple custom property values on a single repository in a single resource block, with in-place updates when values change.",

		CreateContext: resourceGithubRepositoryCustomPropertiesCreate,
		ReadContext:   resourceGithubRepositoryCustomPropertiesRead,
		UpdateContext: resourceGithubRepositoryCustomPropertiesUpdate,
		DeleteContext: resourceGithubRepositoryCustomPropertiesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryCustomPropertiesImport,
		},

		CustomizeDiff: customdiff.All(
			diffRepository,
		),

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the GitHub repository.",
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

func resourceGithubRepositoryCustomPropertiesApply(ctx context.Context, d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
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

	return nil
}

func resourceGithubRepositoryCustomPropertiesCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	owner := meta.(*Owner).name
	client := meta.(*Owner).v3client
	repoName := d.Get("repository").(string)

	if err := resourceGithubRepositoryCustomPropertiesApply(ctx, d, meta); err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryCustomPropertiesUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := resourceGithubRepositoryCustomPropertiesApply(ctx, d, meta); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryCustomPropertiesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	ctx = tflog.SetField(ctx, "id", d.Id())

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	_, repoName, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
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
		return diag.FromErr(fmt.Errorf("error reading custom properties for repository %s/%s: %w", owner, repoName, err))
	}

	managedProperties, err := filterManagedCustomProperties(allCustomProperties, managedPropertyNames, isImport)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error processing custom properties for repository %s/%s: %w", owner, repoName, err))
	}

	// If no properties exist, remove resource from state
	if len(managedProperties) == 0 {
		tflog.Warn(ctx, "No custom properties found, removing from state", map[string]any{"owner": owner, "repository": repoName})
		d.SetId("")
		return nil
	}

	if err := d.Set("repository", repoName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("property", managedProperties); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// filterManagedCustomProperties builds the property set from GitHub API results,
// filtering to only managed properties (or all properties during import).
func filterManagedCustomProperties(allProps []*github.CustomPropertyValue, managed map[string]bool, isImport bool) ([]any, error) {
	result := make([]any, 0)
	for _, prop := range allProps {
		if !isImport && !managed[prop.PropertyName] {
			continue
		}

		if prop.Value == nil {
			continue
		}

		propertyValue, err := parseRepositoryCustomPropertyValueToStringSlice(prop)
		if err != nil {
			return nil, fmt.Errorf("error parsing property %q: %w", prop.PropertyName, err)
		}

		if len(propertyValue) == 0 {
			continue
		}

		result = append(result, map[string]any{
			"name":  prop.PropertyName,
			"value": propertyValue,
		})
	}
	return result, nil
}

func resourceGithubRepositoryCustomPropertiesDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	_, repoName, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
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
		return diag.FromErr(fmt.Errorf("error deleting custom properties for repository %s/%s: %w", owner, repoName, err))
	}

	return nil
}

func resourceGithubRepositoryCustomPropertiesImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	// Import ID format: <repository> — owner is inferred from the provider config.
	// On import, Read will detect empty state and import ALL properties.
	repoName := d.Id()

	owner := meta.(*Owner).name
	client := meta.(*Owner).v3client

	id, err := buildID(owner, repoName)
	if err != nil {
		return nil, err
	}
	d.SetId(id)

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository %s: %w", repoName, err)
	}

	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
