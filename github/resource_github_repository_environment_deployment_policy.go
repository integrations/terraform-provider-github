package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/go-github/v74/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryEnvironmentDeploymentPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryEnvironmentDeploymentPolicyCreate,
		Read:   resourceGithubRepositoryEnvironmentDeploymentPolicyRead,
		Update: resourceGithubRepositoryEnvironmentDeploymentPolicyUpdate,
		Delete: resourceGithubRepositoryEnvironmentDeploymentPolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the repository. The name is not case sensitive.",
			},
			"environment": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the environment.",
			},
			"branch_pattern": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      false,
				ConflictsWith: []string{"tag_pattern"},
				Description:   "The name pattern that branches must match in order to deploy to the environment.",
			},
			"tag_pattern": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      false,
				ConflictsWith: []string{"branch_pattern"},
				Description:   "The name pattern that tags must match in order to deploy to the environment.",
			},
		},
		CustomizeDiff: customDeploymentPolicyDiffFunction,
	}

}

func resourceGithubRepositoryEnvironmentDeploymentPolicyCreate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	escapedEnvName := url.PathEscape(envName)

	var createData github.DeploymentBranchPolicyRequest
	if v, ok := d.GetOk("branch_pattern"); ok {
		createData = github.DeploymentBranchPolicyRequest{
			Name: github.Ptr(v.(string)),
			Type: github.Ptr("branch"),
		}
	} else if v, ok := d.GetOk("tag_pattern"); ok {
		createData = github.DeploymentBranchPolicyRequest{
			Name: github.Ptr(v.(string)),
			Type: github.Ptr("tag"),
		}
	} else {
		return fmt.Errorf("exactly one of %q and %q must be specified", "branch_pattern", "tag_pattern")
	}

	resultKey, _, err := client.Repositories.CreateDeploymentBranchPolicy(ctx, owner, repoName, escapedEnvName, &createData)
	if err != nil {
		return err
	}

	d.SetId(buildThreePartID(repoName, escapedEnvName, strconv.FormatInt(resultKey.GetID(), 10)))
	return resourceGithubRepositoryEnvironmentDeploymentPolicyRead(d, meta)
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	owner := meta.(*Owner).name
	repoName, envName, branchPolicyIdString, err := parseThreePartID(d.Id(), "repository", "environment", "branchPolicyId")
	if err != nil {
		return err
	}

	branchPolicyId, err := strconv.ParseInt(branchPolicyIdString, 10, 64)
	if err != nil {
		return err
	}

	branchPolicy, _, err := client.Repositories.GetDeploymentBranchPolicy(ctx, owner, repoName, envName, branchPolicyId)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
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
		return err
	}

	if branchPolicy.GetType() == "branch" {
		_ = d.Set("branch_pattern", branchPolicy.GetName())
	} else {
		_ = d.Set("tag_pattern", branchPolicy.GetName())
	}
	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyUpdate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	branchPattern := d.Get("branch_pattern").(string)
	tagPattern := d.Get("tag_pattern").(string)
	escapedEnvName := url.PathEscape(envName)
	_, _, branchPolicyIdString, err := parseThreePartID(d.Id(), "repository", "environment", "branchPolicyId")
	if err != nil {
		return err
	}

	branchPolicyId, err := strconv.ParseInt(branchPolicyIdString, 10, 64)
	if err != nil {
		return err
	}

	pattern := branchPattern
	if branchPattern == "" {
		pattern = tagPattern
	}

	updateData := github.DeploymentBranchPolicyRequest{
		Name: github.Ptr(pattern),
	}

	resultKey, _, err := client.Repositories.UpdateDeploymentBranchPolicy(ctx, owner, repoName, escapedEnvName, branchPolicyId, &updateData)
	if err != nil {
		return err
	}
	d.SetId(buildThreePartID(repoName, escapedEnvName, strconv.FormatInt(resultKey.GetID(), 10)))
	return resourceGithubRepositoryEnvironmentDeploymentPolicyRead(d, meta)
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	owner := meta.(*Owner).name
	repoName, envName, branchPolicyIdString, err := parseThreePartID(d.Id(), "repository", "environment", "branchPolicyId")
	if err != nil {
		return err
	}

	branchPolicyId, err := strconv.ParseInt(branchPolicyIdString, 10, 64)
	if err != nil {
		return err
	}

	_, err = client.Repositories.DeleteDeploymentBranchPolicy(ctx, owner, repoName, envName, branchPolicyId)
	if err != nil {
		return err
	}

	return nil
}

func customDeploymentPolicyDiffFunction(_ context.Context, diff *schema.ResourceDiff, v any) error {
	oldBranchPattern, newBranchPattern := diff.GetChange("branch_pattern")

	if oldBranchPattern != "" && newBranchPattern == "" {
		if err := diff.ForceNew("branch_pattern"); err != nil {
			return err
		}
	}
	if oldBranchPattern == "" && newBranchPattern != "" {
		if err := diff.ForceNew("branch_pattern"); err != nil {
			return err
		}
	}

	oldTagPattern, newTagPattern := diff.GetChange("tag_pattern")
	if oldTagPattern != "" && newTagPattern == "" {
		if err := diff.ForceNew("tag_pattern"); err != nil {
			return err
		}
	}
	if oldTagPattern == "" && newTagPattern != "" {
		if err := diff.ForceNew("tag_pattern"); err != nil {
			return err
		}
	}

	return nil
}
