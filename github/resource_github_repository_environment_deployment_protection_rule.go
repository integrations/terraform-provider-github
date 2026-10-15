package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryEnvironmentDeploymentProtectionRule() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a custom deployment protection rule for a GitHub repository environment. A custom deployment protection rule is a GitHub App that approves or rejects deployments to the environment.",

		Schema: map[string]*schema.Schema{
			"repository": {
				Description: "The name of the GitHub repository.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"repository_id": {
				Description: "The ID of the GitHub repository.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"environment": {
				Description: "The name of the environment.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"integration_id": {
				Description: "The ID of the GitHub App to enable as a custom deployment protection rule. The App must be installed on the repository and available as a custom deployment protection rule for the environment.",
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
			},
			"protection_rule_id": {
				Description: "The ID of the custom deployment protection rule.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},

		CustomizeDiff: diffRepository,

		CreateContext: resourceGithubRepositoryEnvironmentDeploymentProtectionRuleCreate,
		ReadContext:   resourceGithubRepositoryEnvironmentDeploymentProtectionRuleRead,
		UpdateContext: resourceGithubRepositoryEnvironmentDeploymentProtectionRuleUpdate,
		DeleteContext: resourceGithubRepositoryEnvironmentDeploymentProtectionRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryEnvironmentDeploymentProtectionRuleImport,
		},
	}
}

func resourceGithubRepositoryEnvironmentDeploymentProtectionRuleCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	envName, _ := d.Get("environment").(string)
	integrationID, _ := d.Get("integration_id").(int)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	rule, _, err := client.Repositories.CreateCustomDeploymentProtectionRule(ctx, owner, repoName, envName, &github.CustomDeploymentProtectionRuleRequest{
		IntegrationID: new(int64(integrationID)),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, escapeIDPart(envName), strconv.FormatInt(rule.GetID(), 10))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("protection_rule_id", int(rule.GetID())); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentProtectionRuleRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "id", d.Id())

	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	envName, _ := d.Get("environment").(string)
	ruleID, _ := d.Get("protection_rule_id").(int)

	rule, _, err := client.Repositories.GetCustomDeploymentProtectionRule(ctx, owner, repoName, envName, int64(ruleID))
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Custom deployment protection rule not found, removing from state.", map[string]any{"repository": repoName, "environment": envName, "protection_rule_id": ruleID})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := d.Set("integration_id", int(rule.GetApp().GetID())); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// Update only runs for a repository rename (see diffRepository); every other argument forces a new resource.
func resourceGithubRepositoryEnvironmentDeploymentProtectionRuleUpdate(_ context.Context, d *schema.ResourceData, _ any) diag.Diagnostics {
	repoName, _ := d.Get("repository").(string)
	envName, _ := d.Get("environment").(string)
	ruleID, _ := d.Get("protection_rule_id").(int)

	id, err := buildID(repoName, escapeIDPart(envName), strconv.Itoa(ruleID))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentProtectionRuleDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "id", d.Id())

	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	envName, _ := d.Get("environment").(string)
	ruleID, _ := d.Get("protection_rule_id").(int)

	_, err := client.Repositories.DisableCustomDeploymentProtectionRule(ctx, owner, repoName, envName, int64(ruleID))
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Custom deployment protection rule not found, treating it as deleted.", map[string]any{"repository": repoName, "environment": envName, "protection_rule_id": ruleID})
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentProtectionRuleImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, envNamePart, ruleIDStr, err := parseID3(d.Id())
	if err != nil {
		return nil, fmt.Errorf("invalid id (%s), expected format <repository>:<environment>:<protection_rule_id>", d.Id())
	}

	ruleID, err := strconv.Atoi(ruleIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid protection rule id: %s", ruleIDStr)
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository %s: %w", repoName, err)
	}

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return nil, err
	}
	if err := d.Set("environment", unescapeIDPart(envNamePart)); err != nil {
		return nil, err
	}
	if err := d.Set("protection_rule_id", ruleID); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
