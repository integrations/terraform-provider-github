package github

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryEnvironmentDeploymentPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubRepositoryEnvironmentDeploymentPolicyCreate,
		ReadContext:   resourceGithubRepositoryEnvironmentDeploymentPolicyRead,
		UpdateContext: resourceGithubRepositoryEnvironmentDeploymentPolicyUpdate,
		DeleteContext: resourceGithubRepositoryEnvironmentDeploymentPolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryEnvironmentDeploymentPolicyImport,
		},
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
			"branch_pattern": {
				Description:  "The name pattern that branches must match in order to deploy to the environment.",
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     false,
				ExactlyOneOf: []string{"branch_pattern", "tag_pattern"},
				ValidateDiagFunc: func(i any, _ cty.Path) diag.Diagnostics {
					str, ok := i.(string)
					if ok && len(str) > 0 {
						return nil
					}
					return diag.Errorf("`branch_pattern` must be a valid non-empty string")
				},
			},
			"tag_pattern": {
				Description:  "The name pattern that tags must match in order to deploy to the environment.",
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     false,
				ExactlyOneOf: []string{"branch_pattern", "tag_pattern"},
				ValidateDiagFunc: func(i any, _ cty.Path) diag.Diagnostics {
					str, ok := i.(string)
					if ok && len(str) > 0 {
						return nil
					}
					return diag.Errorf("`tag_pattern` must be a valid non-empty string")
				},
			},
		},
		CustomizeDiff: customDeploymentPolicyDiffFunction,
	}
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	escapedEnvName := url.PathEscape(envName)

	var createData github.DeploymentBranchPolicyRequest
	var patternType string
	if v, ok := d.GetOk("branch_pattern"); ok {
		patternType = "branch"
		createData = github.DeploymentBranchPolicyRequest{
			Name: github.Ptr(v.(string)),
			Type: github.Ptr("branch"),
		}
	} else if v, ok := d.GetOk("tag_pattern"); ok {
		patternType = "tag"
		createData = github.DeploymentBranchPolicyRequest{
			Name: github.Ptr(v.(string)),
			Type: github.Ptr("tag"),
		}
	}

	tflog.Debug(ctx, "Creating repository environment deployment policy", map[string]any{
		"owner":        owner,
		"repository":   repoName,
		"environment":  envName,
		"pattern_type": patternType,
	})

	resultKey, _, err := client.Repositories.CreateDeploymentBranchPolicy(ctx, owner, repoName, escapedEnvName, &createData)
	if err != nil {
		tflog.Error(ctx, "Failed to create repository environment deployment policy", map[string]any{
			"owner":       owner,
			"repository":  repoName,
			"environment": envName,
			"error":       err.Error(),
		})
		return diag.FromErr(err)
	}

	policyID := resultKey.GetID()
	d.SetId(buildThreePartID(repoName, escapedEnvName, strconv.FormatInt(policyID, 10)))

	tflog.Info(ctx, "Created repository environment deployment policy", map[string]any{
		"owner":       owner,
		"repository":  repoName,
		"environment": envName,
		"policy_id":   policyID,
	})

	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	tflog.Debug(ctx, "Reading repository environment deployment policy", map[string]any{
		"id": d.Id(),
	})
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName, envName, branchPolicyIdString, err := parseThreePartID(d.Id(), "repository", "environment", "branchPolicyId")
	if err != nil {
		return diag.FromErr(err)
	}

	branchPolicyId, err := strconv.ParseInt(branchPolicyIdString, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	branchPolicy, _, err := client.Repositories.GetDeploymentBranchPolicy(ctx, owner, repoName, envName, branchPolicyId)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				tflog.Debug(ctx, "API responded with StatusNotModified, not refreshing state")
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing branch deployment policy from state because it no longer exists in GitHub", map[string]any{
					"owner":       owner,
					"repository":  repoName,
					"environment": envName,
				})
				d.SetId("")
				return nil
			}
		}
		tflog.Error(ctx, "Failed to read repository environment deployment policy", map[string]any{
			"owner":       owner,
			"repository":  repoName,
			"environment": envName,
			"policy_id":   branchPolicyId,
			"error":       err.Error(),
		})
		return diag.FromErr(err)
	}

	patternType := branchPolicy.GetType()
	if patternType == "branch" {
		_ = d.Set("branch_pattern", branchPolicy.GetName())
	} else {
		_ = d.Set("tag_pattern", branchPolicy.GetName())
	}

	tflog.Debug(ctx, "Successfully read repository environment deployment policy", map[string]any{
		"owner":        owner,
		"repository":   repoName,
		"environment":  envName,
		"policy_id":    branchPolicyId,
		"pattern_type": patternType,
	})

	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	branchPattern := d.Get("branch_pattern").(string)
	tagPattern := d.Get("tag_pattern").(string)
	escapedEnvName := url.PathEscape(envName)
	_, _, branchPolicyIdString, err := parseThreePartID(d.Id(), "repository", "environment", "branchPolicyId")
	if err != nil {
		return diag.FromErr(err)
	}

	branchPolicyId, err := strconv.ParseInt(branchPolicyIdString, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, "Updating repository environment deployment policy", map[string]any{
		"owner":       owner,
		"repository":  repoName,
		"environment": envName,
		"policy_id":   branchPolicyId,
	})

	pattern := branchPattern
	if branchPattern == "" {
		pattern = tagPattern
	}

	updateData := github.DeploymentBranchPolicyRequest{
		Name: github.Ptr(pattern),
	}

	resultKey, _, err := client.Repositories.UpdateDeploymentBranchPolicy(ctx, owner, repoName, escapedEnvName, branchPolicyId, &updateData)
	if err != nil {
		tflog.Error(ctx, "Failed to update repository environment deployment policy", map[string]any{
			"owner":       owner,
			"repository":  repoName,
			"environment": envName,
			"policy_id":   branchPolicyId,
			"error":       err.Error(),
		})
		return diag.FromErr(err)
	}

	policyID := resultKey.GetID()
	d.SetId(buildThreePartID(repoName, escapedEnvName, strconv.FormatInt(policyID, 10)))

	tflog.Info(ctx, "Updated repository environment deployment policy", map[string]any{
		"owner":       owner,
		"repository":  repoName,
		"environment": envName,
		"policy_id":   policyID,
	})

	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName, envName, branchPolicyIdString, err := parseThreePartID(d.Id(), "repository", "environment", "branchPolicyId")
	if err != nil {
		return diag.FromErr(err)
	}

	branchPolicyId, err := strconv.ParseInt(branchPolicyIdString, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, "Deleting repository environment deployment policy", map[string]any{
		"owner":       owner,
		"repository":  repoName,
		"environment": envName,
		"policy_id":   branchPolicyId,
	})

	_, err = client.Repositories.DeleteDeploymentBranchPolicy(ctx, owner, repoName, envName, branchPolicyId)
	if err != nil {
		tflog.Error(ctx, "Failed to delete repository environment deployment policy", map[string]any{
			"owner":       owner,
			"repository":  repoName,
			"environment": envName,
			"policy_id":   branchPolicyId,
			"error":       err.Error(),
		})
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Deleted repository environment deployment policy", map[string]any{
		"owner":       owner,
		"repository":  repoName,
		"environment": envName,
		"policy_id":   branchPolicyId,
	})

	return nil
}

func customDeploymentPolicyDiffFunction(_ context.Context, diff *schema.ResourceDiff, v any) error {
	if diff.HasChange("branch_pattern") && diff.HasChange("tag_pattern") {
		if err := diff.ForceNew("branch_pattern"); err != nil {
			return err
		}
		if err := diff.ForceNew("tag_pattern"); err != nil {
			return err
		}
	}

	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	tflog.Debug(ctx, "Importing repository environment deployment policy", map[string]any{
		"id": d.Id(),
	})

	repoName, envName, policyId, err := parseThreePartID(d.Id(), "repository", "environment", "branchPolicyId")
	if err != nil {
		tflog.Error(ctx, "Failed to parse import ID", map[string]any{
			"id":    d.Id(),
			"error": err.Error(),
		})
		return nil, err
	}

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}

	envNameUnescaped, err := url.PathUnescape(envName)
	if err != nil {
		tflog.Error(ctx, "Failed to unescape environment name", map[string]any{
			"id":          d.Id(),
			"environment": envName,
			"error":       err.Error(),
		})
		return nil, err
	}

	if err := d.Set("environment", envNameUnescaped); err != nil {
		return nil, err
	}

	tflog.Info(ctx, "Imported repository environment deployment policy", map[string]any{
		"repository":  repoName,
		"environment": envNameUnescaped,
		"policy_id":   policyId,
	})

	return []*schema.ResourceData{d}, nil
}
