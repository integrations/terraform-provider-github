package github

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	abs "github.com/microsoft/kiota-abstractions-go"
	"github.com/octokit/go-sdk/pkg/github/models"
)

func dataSourceGithubRepositoryCustomProperties() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrgaRepositoryCustomProperties,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository which the custom properties should be on.",
			},
			"properties": {
				Type:        schema.TypeMap,
				Computed:    true,
				Elem:        schema.TypeString, // This should arguably be a TypeSet/TypeList with strings as their sub-element. Played around with it a bit, but never got it to work...
				Description: "Map of property keys and their corresponding values formatted as comma separated strings. I.e., multi_select properties will have values similar to `option1,option2`",
			},
		},
	}
}

func parseRepositoryCustomProperties(repo models.FullRepositoryable) (map[string]string, error) {
	repoFullName := repo.GetFullName()
	repoProps := repo.GetCustomProperties().GetAdditionalData()

	properties := make(map[string]string)
	for key, value := range repoProps {

		typeAssertionErr := errors.New(fmt.Sprintf("error reading custom property `%v` in %s. Value couldn't be parsed as a string, or a list of strings", key, *repoFullName))

		// The value of a custom property can be either a string, or a list of strings (https://docs.github.com/en/enterprise-cloud@latest/rest/repos/custom-properties?apiVersion=2022-11-28#get-all-custom-property-values-for-a-repository)
		if valueAsString, ok := value.(*string); ok {
			properties[key] = *valueAsString
		} else if valueAsInterfaceSlice, ok := value.([]interface{}); ok {
			// Format the multi_select props as comma separated values
			var valueStringBuilder strings.Builder
			for _, valueAsInterface := range valueAsInterfaceSlice {
				if valueAsString, ok := valueAsInterface.(*string); ok {
					valueStringBuilder.WriteString(*valueAsString)
					valueStringBuilder.WriteString(",")
				} else {
					return nil, typeAssertionErr
				}

			}
			properties[key] = strings.TrimSuffix(valueStringBuilder.String(), ",") // Remove any trailing commas
		} else {
			return nil, typeAssertionErr
		}
	}

	return properties, nil
}

func dataSourceGithubOrgaRepositoryCustomProperties(d *schema.ResourceData, meta interface{}) error {

	octokitClient := meta.(*Owner).octokitClient
	ctx := context.Background()

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

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

	d.SetId(buildTwoPartID(owner, repoName)) // TODO: Maybe this should just be the repo name
	d.Set("repository", repoName)
	d.Set("properties", properties)

	return nil
}
