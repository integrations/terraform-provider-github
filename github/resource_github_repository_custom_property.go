package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v65/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
				ForceNew: true,
			},
			"property_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the custom property.",
				ForceNew: true,
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

func resourceGithubRepositoryCustomPropertyCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Owner).v3client
	ctx := context.Background()

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	propertyName := d.Get("property_name").(string)
	propertyValue := expandStringList(d.Get("property_value").(*schema.Set).List())

	customProperty := github.CustomPropertyValue{
		PropertyName: propertyName,
	}

	// The propertyValue can either be a list of strings or a string
	switch valueLength := len(propertyValue); valueLength {
	case 0:
		return fmt.Errorf("custom property value cannot be an empty list: %v", propertyValue)
	case 1:
		customProperty.Value = propertyValue[0]
	default:
		customProperty.Value = propertyValue
	}

	_, err := client.Repositories.CreateOrUpdateCustomProperties(ctx, owner, repoName, []*github.CustomPropertyValue{&customProperty})
	if err != nil {
		return err
	}

	d.SetId(buildThreePartID(owner, repoName, propertyName))

	return resourceGithubRepositoryCustomPropertyRead(d, meta)
}

func resourceGithubRepositoryCustomPropertyRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Owner).v3client
	ctx := context.Background()

	owner, repoName, propertyName, err := parseThreePartID(d.Id(), "owner", "repoName", "propertyName"); 
	if err != nil {
		return err
	}

	wantedCustomPropertyValue, err := readRepositoryCustomPropertyValue(ctx, client,  owner, repoName, propertyName)
	if err != nil {
		return err
	}

	d.SetId(buildThreePartID(owner, repoName, propertyName))
	d.Set("repository", repoName)
	d.Set("property_name", propertyName)
	d.Set("property_value", wantedCustomPropertyValue)

	return nil
}

func resourceGithubRepositoryCustomPropertyDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Owner).v3client
	ctx := context.Background()

	owner, repoName, propertyName, err := parseThreePartID(d.Id(), "owner", "repoName", "propertyName"); 
	if err != nil {
		return err
	}

	customProperty := github.CustomPropertyValue{
		PropertyName: propertyName,
		Value: nil, // TODO: setting this will remove the custom proprty, but it's currently blocked by https://github.com/google/go-github/pull/3309
	}

	_, err = client.Repositories.CreateOrUpdateCustomProperties(ctx, owner, repoName, []*github.CustomPropertyValue{&customProperty})
	if err != nil {
		return err
	}

	return nil
}

func readRepositoryCustomPropertyValue(ctx context.Context, client *github.Client, owner, repoName, propertyName string) ([]string, error){
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
		return nil, fmt.Errorf("could not find a custom property with name: %s", propertyName)
	}

	var wantedCustomPropertyValue []string
	switch value := wantedCustomProperty.Value.(type) {
	case string:
		wantedCustomPropertyValue = []string{value}
	case []string:
		wantedCustomPropertyValue = value
	default:
		return nil, fmt.Errorf("custom property value couldn't be parsed as a string or a list of strings: %s", value)
	}

	return wantedCustomPropertyValue, nil
}