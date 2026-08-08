package github

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsEnterpriseRunnerGroupV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
			},
			"allows_public_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether public repositories can be added to the runner group.",
			},
			"default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether this is the default runner group.",
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An etag representing the runner group object",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the runner group.",
			},
			"runners_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The GitHub API URL for the runner group's runners.",
			},
			"visibility": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The visibility of the runner group.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "selected"}, false)),
			},
			"restricted_to_workflows": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "If 'true', the runner group will be restricted to running only the workflows specified in the 'selected_workflows' array. Defaults to 'false'.",
			},
			"selected_workflows": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to 'true'.",
			},
			"selected_organization_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Set:         schema.HashInt,
				Optional:    true,
				Description: "List of organization IDs that can access the runner group.",
			},
			"selected_organizations_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GitHub API URL for the runner group's organizations.",
			},
		},
	}
}

func resourceGithubActionsEnterpriseRunnerGroupStateUpgradeV0(_ context.Context, rawState map[string]any, _ any) (map[string]any, error) {
	log.Printf("[DEBUG] GitHub Enterprise Actions Runner Group Attributes before migration: %#v", rawState)

	// No transformation needed. The SDK re-encodes selected_workflows from
	// TypeList (index-based keys) to TypeSet (hash-based keys) automatically
	// when it persists the state with the new schema version.

	log.Printf("[DEBUG] GitHub Enterprise Actions Runner Group Attributes after migration: %#v", rawState)

	return rawState, nil
}
