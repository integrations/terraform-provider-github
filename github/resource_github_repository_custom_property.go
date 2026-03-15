package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryCustomProperty() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryCustomPropertyCreate,
		Read:   resourceGithubRepositoryCustomPropertyRead,
		Delete: resourceGithubRepositoryCustomPropertyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository which the custom properties should be on.",
				ForceNew:    true,
			},
			"property_type": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Type of the custom property",
				ForceNew:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{string(github.PropertyValueTypeString), string(github.PropertyValueTypeSingleSelect), string(github.PropertyValueTypeMultiSelect), string(github.PropertyValueTypeTrueFalse), string(github.PropertyValueTypeURL)}, false)),
			},
			"property_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the custom property.",
				ForceNew:    true,
			},
			"property_value": {
				Type:        schema.TypeSet,
				MinItems:    1,
				Required:    true,
				Description: "Value of the custom property.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ForceNew: true,
			},
		},
	}
}

func resourceGithubRepositoryCustomPropertyCreate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	propertyName := d.Get("property_name").(string)
	propertyType := github.PropertyValueType(d.Get("property_type").(string))
	propertyValue := expandStringList(d.Get("property_value").(*schema.Set).List())

	customProperty := github.CustomPropertyValue{
		PropertyName: propertyName,
	}

	// The propertyValue can either be a list of strings or a string
	switch propertyType {
	case github.PropertyValueTypeString, github.PropertyValueTypeSingleSelect, github.PropertyValueTypeURL, github.PropertyValueTypeTrueFalse:
		customProperty.Value = propertyValue[0]
	case github.PropertyValueTypeMultiSelect:
		customProperty.Value = propertyValue
	default:
		return fmt.Errorf("custom property type is not valid: %v", propertyType)
	}

	_, err := client.Repositories.CreateOrUpdateCustomProperties(ctx, owner, repoName, []*github.CustomPropertyValue{&customProperty})
	if err != nil {
		return err
	}

	d.SetId(buildThreePartID(owner, repoName, propertyName))

	return resourceGithubRepositoryCustomPropertyRead(d, meta)
}

func resourceGithubRepositoryCustomPropertyRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	owner, repoName, propertyName, err := parseThreePartID(d.Id(), "owner", "repoName", "propertyName")
	if err != nil {
		return err
	}

	wantedCustomPropertyValue, err := readRepositoryCustomPropertyValue(ctx, client, owner, repoName, propertyName)
	if err != nil {
		return err
	}

	if wantedCustomPropertyValue == nil {
		d.SetId("")
		return nil
	}

	d.SetId(buildThreePartID(owner, repoName, propertyName))
	_ = d.Set("repository", repoName)
	_ = d.Set("property_name", propertyName)
	_ = d.Set("property_value", wantedCustomPropertyValue)

	return nil
}

func resourceGithubRepositoryCustomPropertyDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	owner, repoName, propertyName, err := parseThreePartID(d.Id(), "owner", "repoName", "propertyName")
	if err != nil {
		return err
	}

	customProperty := github.CustomPropertyValue{
		PropertyName: propertyName,
		Value:        nil,
	}

	_, err = client.Repositories.CreateOrUpdateCustomProperties(ctx, owner, repoName, []*github.CustomPropertyValue{&customProperty})
	if err != nil {
		return err
	}

	return nil
}

func readRepositoryCustomPropertyValue(ctx context.Context, client *github.Client, owner, repoName, propertyName string) ([]string, error) {
	allCustomProperties, _, err := client.Repositories.GetAllCustomPropertyValues(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}

	var wantedCustomProperty *github.CustomPropertyValue
	for _, customProperty := range allCustomProperties {
		if customProperty.PropertyName == propertyName {
			wantedCustomProperty = customProperty
		}
	}

	if wantedCustomProperty == nil {
		return nil, nil
	}

	wantedPropertyValue, err := parseRepositoryCustomPropertyValueToStringSlice(wantedCustomProperty)
	if err != nil {
		return nil, err
	}

	return wantedPropertyValue, nil
}
