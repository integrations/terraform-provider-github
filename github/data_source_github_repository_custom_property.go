package github

import (
	"context"

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
				Description: "Name of the custom property.",
			},
			"property_value": {
				Type:     schema.TypeList,
				Computed: true,
				Description: "Value of the custom property.",
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
