package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryEnvironmentDeploymentProtectionRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"repository": {
				Description: "The name of the GitHub repository.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"environment": {
				Description: "The name of the environment.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"integration_id": {
				Description: "The ID of the custom deployment protection rule integration — the GitHub App that gates deployments to the environment.",
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
			},
		},

		CreateContext: resourceGithubRepositoryEnvironmentDeploymentProtectionRuleCreate,
		ReadContext:   resourceGithubRepositoryEnvironmentDeploymentProtectionRuleRead,
		DeleteContext: resourceGithubRepositoryEnvironmentDeploymentProtectionRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryEnvironmentDeploymentProtectionRuleImport,
		},
	}
}

func resourceGithubRepositoryEnvironmentDeploymentProtectionRuleCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	integrationID := int64(d.Get("integration_id").(int))

	rule, _, err := client.Repositories.CreateCustomDeploymentProtectionRule(ctx, owner, repoName, url.PathEscape(envName), &github.CustomDeploymentProtectionRuleRequest{
		IntegrationID: new(integrationID),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, escapeIDPart(envName), strconv.FormatInt(rule.GetID(), 10))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	return resourceGithubRepositoryEnvironmentDeploymentProtectionRuleRead(ctx, d, m)
}

func resourceGithubRepositoryEnvironmentDeploymentProtectionRuleRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "id", d.Id())

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, envNamePart, ruleIDStr, err := parseID3(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid id (%s), expected format <repository>:<environment>:<rule_id>", d.Id()))
	}
	envName := unescapeIDPart(envNamePart)
	ruleID, err := strconv.ParseInt(ruleIDStr, 10, 64)
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid rule id: %s", ruleIDStr))
	}

	rule, _, err := client.Repositories.GetCustomDeploymentProtectionRule(ctx, owner, repoName, url.PathEscape(envName), ruleID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Custom deployment protection rule not found, removing from state.", map[string]any{"repository": repoName, "environment": envName, "rule_id": ruleID})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if err := d.Set("repository", repoName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("environment", envName); err != nil {
		return diag.FromErr(err)
	}
	if rule.App != nil {
		if err := d.Set("integration_id", int(rule.App.GetID())); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentProtectionRuleDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, envNamePart, ruleIDStr, err := parseID3(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid id (%s), expected format <repository>:<environment>:<rule_id>", d.Id()))
	}
	envName := unescapeIDPart(envNamePart)
	ruleID, err := strconv.ParseInt(ruleIDStr, 10, 64)
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid rule id: %s", ruleIDStr))
	}

	_, err = client.Repositories.DisableCustomDeploymentProtectionRule(ctx, owner, repoName, url.PathEscape(envName), ruleID)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentProtectionRuleImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	repoName, envNamePart, ruleIDStr, err := parseID3(d.Id())
	if err != nil {
		return nil, fmt.Errorf("invalid id (%s), expected format <repository>:<environment>:<rule_id>", d.Id())
	}

	if _, err := strconv.ParseInt(ruleIDStr, 10, 64); err != nil {
		return nil, fmt.Errorf("invalid rule id: %s", ruleIDStr)
	}

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err := d.Set("environment", unescapeIDPart(envNamePart)); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
