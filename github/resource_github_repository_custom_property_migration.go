package github

import (
	"context"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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

	tflog.Debug(ctx, "Migrating GitHub Repository Custom Property from v0 to v1.", map[string]any{"raw_state": rawState})

	state, err := migrateRepositoryWithID(ctx, client, owner, rawState)
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "GitHub Repository Custom Property migrated to v1.", map[string]any{"raw_state": state})

	return state, nil
}
