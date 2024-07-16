package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	abs "github.com/microsoft/kiota-abstractions-go"
	"github.com/octokit/go-sdk/pkg/github/models"
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
				Type: schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func parseRepositoryCustomProperties(repo models.FullRepositoryable) (map[string][]string, error) {
	repoFullName := repo.GetFullName()
	repoProps := repo.GetCustomProperties().GetAdditionalData()

	properties := make(map[string][]string)
	for key, value := range repoProps {

		typeAssertionErr := fmt.Errorf("error reading custom property `%v` in %s. Value couldn't be parsed as a string, or a list of strings", key, *repoFullName)

		// The value of a custom property can be either a string, or a list of strings (https://docs.github.com/en/enterprise-cloud@latest/rest/repos/custom-properties?apiVersion=2022-11-28#get-all-custom-property-values-for-a-repository)
		switch valueStringOrSlice := value.(type) {
		case *string:
			interfaceSlice := make([]string, 1)
			interfaceSlice[0] = *valueStringOrSlice
			properties[key] = interfaceSlice

		case []interface{}:
			interfaceSlice := make([]string, len(valueStringOrSlice))
			for idx, valInterface := range valueStringOrSlice {
				switch valString := valInterface.(type) {
				case *string:
					interfaceSlice[idx] = *valString

				default:
					return nil, typeAssertionErr
				}
			}
			properties[key] = interfaceSlice

		default:
			return nil, typeAssertionErr
		}
	}

	return properties, nil
}

func dataSourceGithubOrgaRepositoryCustomProperty(d *schema.ResourceData, meta interface{}) error {

	octokitClient := meta.(*Owner).octokitClient
	ctx := context.Background()

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	propertyName := d.Get("property_name").(string)

	repoRequestConfig := &abs.RequestConfiguration[abs.DefaultQueryParameters]{
		QueryParameters: &abs.DefaultQueryParameters{},
	}
	repo, err := octokitClient.Repos().ByOwnerId(owner).ByRepoId(repoName).Get(ctx, repoRequestConfig)
	if err != nil {
		return err
	}

	properties, err := parseRepositoryCustomProperties(repo)
	if err != nil {
		return err
	}

	d.SetId(buildThreePartID(owner, repoName, propertyName))
	d.Set("repository", repoName)
	d.Set("property_name", propertyName)
	d.Set("property_value", properties[propertyName])

	return nil
}
