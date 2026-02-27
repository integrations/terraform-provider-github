package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryEnvironmentDeploymentPolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubRepositoryEnvironmentDeploymentPolicyV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubRepositoryEnvironmentDeploymentPolicyStateUpgradeV0,
				Version: 0,
			},
		},

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
			"policy_id": {
				Description: "The ID of the deployment policy.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},

		CustomizeDiff: customdiff.All(
			diffRepository,
			resourceGithubRepositoryEnvironmentDeploymentPolicyDiff,
		),

		CreateContext: resourceGithubRepositoryEnvironmentDeploymentPolicyCreate,
		ReadContext:   resourceGithubRepositoryEnvironmentDeploymentPolicyRead,
		UpdateContext: resourceGithubRepositoryEnvironmentDeploymentPolicyUpdate,
		DeleteContext: resourceGithubRepositoryEnvironmentDeploymentPolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryEnvironmentDeploymentPolicyImport,
		},
	}
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyDiff(_ context.Context, d *schema.ResourceDiff, _ any) error {
	if d.Id() == "" {
		return nil
	}

	if d.HasChange("branch_pattern") && d.HasChange("tag_pattern") {
		if err := d.ForceNew("branch_pattern"); err != nil {
			return err
		}
		if err := d.ForceNew("tag_pattern"); err != nil {
			return err
		}
	}

	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	branchPattern := d.Get("branch_pattern").(string)
	tagPattern := d.Get("tag_pattern").(string)

	policyType := "branch"
	pattern := branchPattern
	if branchPattern == "" {
		policyType = "tag"
		pattern = tagPattern
	}

	createData := github.DeploymentBranchPolicyRequest{
		Name: github.Ptr(pattern),
		Type: github.Ptr(policyType),
	}

	policy, _, err := client.Repositories.CreateDeploymentBranchPolicy(ctx, owner, repoName, url.PathEscape(envName), &createData)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, escapeIDPart(envName), strconv.FormatInt(policy.GetID(), 10))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("policy_id", policy.GetID()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "id", d.Id())

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	policyID := d.Get("policy_id").(int)

	policy, _, err := client.Repositories.GetDeploymentBranchPolicy(ctx, owner, repoName, url.PathEscape(envName), int64(policyID))
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Deployment branch policy not found, removing from state.", map[string]any{"repository": repoName, "environment": envName, "policy_id": policyID})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if policy.GetType() == "branch" {
		if err := d.Set("branch_pattern", policy.GetName()); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err := d.Set("tag_pattern", policy.GetName()); err != nil {
			return diag.FromErr(err)
		}
	}
	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	branchPattern := d.Get("branch_pattern").(string)
	tagPattern := d.Get("tag_pattern").(string)
	policyID := d.Get("policy_id").(int)

	pattern := branchPattern
	if branchPattern == "" {
		pattern = tagPattern
	}

	updateData := github.DeploymentBranchPolicyRequest{
		Name: github.Ptr(pattern),
	}

	_, _, err := client.Repositories.UpdateDeploymentBranchPolicy(ctx, owner, repoName, url.PathEscape(envName), int64(policyID), &updateData)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(repoName, escapeIDPart(envName), strconv.Itoa(policyID))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	policyID := d.Get("policy_id").(int)

	_, err := client.Repositories.DeleteDeploymentBranchPolicy(ctx, owner, repoName, url.PathEscape(envName), int64(policyID))
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, envNamePart, policyIDStr, err := parseID3(d.Id())
	if err != nil {
		return nil, fmt.Errorf("invalid id (%s), expected format <repository>:<environment>:<policy_id>", d.Id())
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve repository %s: %w", repoName, err)
	}

	policyID, err := strconv.Atoi(policyIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid policy ID: %s", policyIDStr)
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
	if err := d.Set("policy_id", policyID); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
