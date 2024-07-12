package github

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	abs "github.com/microsoft/kiota-abstractions-go"
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
				Type:     schema.TypeMap,
				Computed: true,
				Elem: schema.TypeString,
			},
		},
	}
}

func dataSourceGithubOrgaRepositoryCustomProperties(d *schema.ResourceData, meta interface{}) error {

	octokitClient := meta.(*Owner).octokitClient
	ctx := meta.(*Owner).StopContext

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	repoRequestConfig := &abs.RequestConfiguration[abs.DefaultQueryParameters]{
		QueryParameters: &abs.DefaultQueryParameters{},
	}
	repo, err := octokitClient.Repos().ByOwnerId(owner).ByRepoId(repoName).Get(ctx, repoRequestConfig)
	if err != nil {
		return err
	}

	properties := make(map[string]string)
	repoProps := repo.GetCustomProperties()
	for key, value := range repoProps.GetAdditionalData() {

		// The value of a custom property can be either a string, or a list of strings (https://docs.github.com/en/enterprise-cloud@latest/rest/repos/custom-properties?apiVersion=2022-11-28#get-all-custom-property-values-for-a-repository)
		valueAsString, ok := value.(*string)
		if ok {
			properties[key] = *valueAsString
		} else {
			var valueStringBuilder strings.Builder
			
			for _, valueInterface := range value.([]interface{}) {
				valueInterfaceAsString := valueInterface.(*string)
				valueStringBuilder.WriteString(*valueInterfaceAsString)
				valueStringBuilder.WriteString(",")

			}

			properties[key] = strings.TrimSuffix(valueStringBuilder.String(), ",") // Remove any trailing commas
		}
	}

	d.SetId(buildTwoPartID(owner, repoName)) // TODO: Maybe this should just be the repo name
	d.Set("repository", repoName)

	// test := make(map[string]string)
	// for k, value := range properties {
	// 	// valueSlice := value.([]string)
	// 	test[k] = value[0]
	// }

	// d.Set("properties", test)

	for key, value := range properties {
		// valueAsSlice := value.([]string)
		log.Printf("[DEBUG] HEREEEE %v: %v \n", key, value)
	}

	err = d.Set("properties", properties)
	if err != nil {
		return err
	}

	return nil
}
