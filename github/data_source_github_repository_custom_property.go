package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v65/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryCustomProperty() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrgaRepositoryCustomProperty,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository which the custom properties should be on.",
			},
			"property_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository which the custom properties should be on.",
			},
			"property_value": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceGithubOrgaRepositoryCustomProperty(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Owner).v3client
	ctx := context.Background()

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	propertyName := d.Get("property_name").(string)

	allCustomProperties, _, err := client.Repositories.GetAllCustomPropertyValues(ctx, owner, repoName)
	if err != nil {
		return err
	}

	var wantedCustomProperty *github.CustomPropertyValue
	for _, customProperty := range allCustomProperties {
		if customProperty.PropertyName == propertyName {
			wantedCustomProperty = customProperty
		}
	}

	if wantedCustomProperty == nil {
		return fmt.Errorf("could not find a custom property with name: %s", propertyName)
	}

	var wantedCustomPropertyValue []string // := make([]string, 0)
	switch value := wantedCustomProperty.Value.(type) {
	case string:
		wantedCustomPropertyValue = []string{value}
	case []string:
		wantedCustomPropertyValue = value
	default:
		return fmt.Errorf("custom property value couldn't be parsed as a string or a list of strings: %s", value)
	}

	d.SetId(buildThreePartID(owner, repoName, propertyName))
	d.Set("repository", repoName)
	d.Set("property_name", wantedCustomProperty.PropertyName)
	d.Set("property_value", wantedCustomPropertyValue)

	return nil
}
