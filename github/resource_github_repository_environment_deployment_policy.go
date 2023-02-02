package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v50/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func resourceGithubRepositoryEnvironmentDeploymentPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryEnvironmentDeploymentPolicyCreate,
		Read:   resourceGithubRepositoryEnvironmentDeploymentPolicyRead,
		Update: resourceGithubRepositoryEnvironmentDeploymentPolicyUpdate,
		Delete: resourceGithubRepositoryEnvironmentDeploymentPolicyDelete,
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"environment": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"branch_pattern": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
		},
	}

}

func resourceGithubRepositoryEnvironmentDeploymentPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	branchPattern := d.Get("branch_pattern").(string)
	escapedEnvName := url.QueryEscape(envName)

	createData := github.DeploymentBranchPolicyRequest{
		Name: github.String(branchPattern),
	}

	resultKey, _, err := client.Repositories.CreateDeploymentBranchPolicy(ctx, owner, repoName, escapedEnvName, &createData)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s/%d", repoName, envName, resultKey.GetID()))
	return resourceGithubRepositoryEnvironmentDeploymentPolicyRead(d, meta)
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	owner := meta.(*Owner).name
	branchPolicyId, err := parseDeploymentBranchPolicyId(d.Id())
	if err != nil {
		return err
	}
	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	escapedEnvName := url.QueryEscape(envName)

	branchPolicy, _, err := client.Repositories.GetDeploymentBranchPolicy(ctx, owner, repoName, escapedEnvName, branchPolicyId)
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

	d.Set("branch_pattern", branchPolicy.GetName())
	return nil
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	branchPattern := d.Get("branch_pattern").(string)
	escapedEnvName := url.QueryEscape(envName)
	branchPolicyId, err := parseDeploymentBranchPolicyId(d.Id())
	if err != nil {
		return err
	}
	updateData := github.DeploymentBranchPolicyRequest{
		Name: github.String(branchPattern),
	}

	ctx := context.Background()

	resultKey, _, err := client.Repositories.UpdateDeploymentBranchPolicy(ctx, owner, repoName, escapedEnvName, branchPolicyId, &updateData)
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s/%s/%d", repoName, envName, resultKey.GetID()))
	return resourceGithubRepositoryEnvironmentDeploymentPolicyRead(d, meta)
}

func resourceGithubRepositoryEnvironmentDeploymentPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	envName := d.Get("environment").(string)
	escapedEnvName := url.QueryEscape(envName)
	branchPolicyId, err := parseDeploymentBranchPolicyId(d.Id())
	if err != nil {
		return err
	}

	ctx := context.Background()
	_, err = client.Repositories.DeleteDeploymentBranchPolicy(ctx, owner, repoName, escapedEnvName, branchPolicyId)
	if err != nil {
		return err
	}

	return nil
}

func parseDeploymentBranchPolicyId(id string) (int64, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 3 {
		return -1, fmt.Errorf("ID not properly formatted: %s", id)
	}
	number, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		return -1, err
	}
	return number, nil
}
