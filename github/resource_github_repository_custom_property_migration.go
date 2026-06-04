package github

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryCustomPropertyV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,

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

func resourceGithubRepositoryCustomPropertyStateUpgradeV0(ctx context.Context, rawState map[string]any, m any) (map[string]any, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	log.Printf("[DEBUG] GitHub Repository Custom Property Attributes before migration: %#v", rawState)

	repoName, ok := rawState["repository"].(string)
	if !ok {
		return nil, fmt.Errorf("repository not found or is not a string")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository %s: %w", repoName, err)
	}

	repoID := int(repo.GetID())
	rawState["repository_id"] = repoID

	log.Printf("[DEBUG] GitHub Repository Custom Property Attributes after migration: %#v", rawState)

	return rawState, nil
}
