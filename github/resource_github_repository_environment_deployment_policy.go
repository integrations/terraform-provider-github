package github

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/go-cty/cty"
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
			StateContext: schema.ImportStatePassthroughContext,
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
	if v, ok := d.GetOk("branch_pattern"); ok {
		createData = github.DeploymentBranchPolicyRequest{
			Name: github.String(v.(string)),
			Type: github.String("branch"),
		}
	} else if v, ok := d.GetOk("tag_pattern"); ok {
		createData = github.DeploymentBranchPolicyRequest{
			Name: github.String(v.(string)),
			Type: github.String("tag"),
		}
	} else {
		return diag.Errorf("only one of 'branch_pattern' or 'tag_pattern' must be specified")
	}

	resultKey, _, err := client.Repositories.CreateDeploymentBranchPolicy(ctx, owner, repoName, escapedEnvName, &createData)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildThreePartID(repoName, escapedEnvName, strconv.FormatInt(resultKey.GetID(), 10)))
	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
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
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[INFO] Removing branch deployment policy for %s/%s/%s from state because it no longer exists in GitHub",
					owner, repoName, envName)
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if branchPolicy.GetType() == "branch" {
		_ = d.Set("branch_pattern", branchPolicy.GetName())
	} else {
		_ = d.Set("tag_pattern", branchPolicy.GetName())
	}
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

	pattern := branchPattern
	if branchPattern == "" {
		pattern = tagPattern
	}

	updateData := github.DeploymentBranchPolicyRequest{
		Name: github.String(pattern),
	}

	resultKey, _, err := client.Repositories.UpdateDeploymentBranchPolicy(ctx, owner, repoName, escapedEnvName, branchPolicyId, &updateData)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(buildThreePartID(repoName, escapedEnvName, strconv.FormatInt(resultKey.GetID(), 10)))
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

	_, err = client.Repositories.DeleteDeploymentBranchPolicy(ctx, owner, repoName, envName, branchPolicyId)
	if err != nil {
		return diag.FromErr(err)
	}

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
