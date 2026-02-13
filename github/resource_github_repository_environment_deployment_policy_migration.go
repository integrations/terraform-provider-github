package github

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryEnvironmentDeploymentPolicyV0() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,

		Schema: map[string]*schema.Schema{
			"repository": {
				Description: "The name of the GitHub repository.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"environment": {
				Description: "The name of the environment.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"branch_pattern": {
				Description:      "The name pattern that branches must match in order to deploy to the environment.",
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         false,
				ExactlyOneOf:     []string{"branch_pattern", "tag_pattern"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
			"tag_pattern": {
				Description:      "The name pattern that tags must match in order to deploy to the environment.",
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         false,
				ExactlyOneOf:     []string{"branch_pattern", "tag_pattern"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
		},
	}
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyStateUpgradeV0(_ context.Context, rawState map[string]any, _ any) (map[string]any, error) {
	log.Printf("[DEBUG] GitHub Repository Environment Deployment Policy Attributes before migration: %#v", rawState)

	_, _, policyIDStr, err := parseID3(rawState["id"].(string))
	if err != nil {
		return nil, err
	}

	policyID, err := strconv.Atoi(policyIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid policy ID: %s", policyIDStr)
	}

	rawState["policy_id"] = policyID

	log.Printf("[DEBUG] GitHub Repository Environment Deployment Policy Attributes after migration: %#v", rawState)

	return rawState, nil
}
