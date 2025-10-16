package github

import (
	"context"
	"net/url"

	"github.com/google/go-github/v66/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryEnvironmentCustomDeploymentProtectionRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRepositoryEnvironmentCustomDeploymentProtectionRulesRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the repository. The name is not case sensitive.",
			},
			"environment": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the environment.",
			},
			"custom_deployment_protection_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"integration_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"node_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubRepositoryEnvironmentCustomDeploymentProtectionRulesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	escapedEnvName := url.PathEscape(envName)

	availableCustomDeploymentRules, _, err := client.Repositories.ListCustomDeploymentRuleIntegrations(ctx, owner, repoName, escapedEnvName)
	if err != nil {
		return err
	}
	enabledCustomDeploymentRules, _, err := client.Repositories.GetAllDeploymentProtectionRules(ctx, owner, repoName, escapedEnvName)
	if err != nil {
		return err
	}

	results := make([]map[string]interface{}, 0)
	results = append(results, flattenCustomDeploymentRules(availableCustomDeploymentRules, enabledCustomDeploymentRules)...)
	d.SetId(escapedEnvName)
	err = d.Set("custom_deployment_protection_rules", results)
	if err != nil {
		return err
	}
	return nil
}

func flattenCustomDeploymentRules(availableCustomDeploymentRules *github.ListCustomDeploymentRuleIntegrationsResponse, enabledCustomDeploymentRules *github.ListDeploymentProtectionRuleResponse) []map[string]interface{} {
	results := make([]map[string]interface{}, 0)
	if availableCustomDeploymentRules == nil && enabledCustomDeploymentRules == nil {
		return results
	}
	for _, customDeploymentRule := range availableCustomDeploymentRules.AvailableIntegrations {
		customDeploymentRuleMap := make(map[string]interface{})
		customDeploymentRuleMap["id"] = customDeploymentRule.GetID()
		customDeploymentRuleMap["slug"] = customDeploymentRule.GetSlug()
		customDeploymentRuleMap["integration_url"] = customDeploymentRule.GetIntegrationURL()
		customDeploymentRuleMap["node_id"] = customDeploymentRule.GetNodeID()
		results = append(results, customDeploymentRuleMap)
	}

	for _, enabledCustomenabledCustomDeploymentRule := range enabledCustomDeploymentRules.ProtectionRules {
		enabledCustomDeploymentRuleMap := make(map[string]interface{})
		enabledCustomDeploymentRuleMap["id"] = enabledCustomenabledCustomDeploymentRule.App.GetID()
		enabledCustomDeploymentRuleMap["slug"] = enabledCustomenabledCustomDeploymentRule.App.GetSlug()
		enabledCustomDeploymentRuleMap["integration_url"] = enabledCustomenabledCustomDeploymentRule.App.GetIntegrationURL()
		enabledCustomDeploymentRuleMap["node_id"] = enabledCustomenabledCustomDeploymentRule.App.GetNodeID()
		results = append(results, enabledCustomDeploymentRuleMap)

	}
	return results
}
