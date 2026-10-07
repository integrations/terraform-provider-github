package github

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubRepositoryDeploymentBranchPolicy() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "This resource is deprecated in favour of the github_repository_environment_deployment_policy resource.",

		CreateContext: resourceGithubRepositoryDeploymentBranchPolicyCreate,
		ReadContext:   resourceGithubRepositoryDeploymentBranchPolicyRead,
		UpdateContext: resourceGithubRepositoryDeploymentBranchPolicyUpdate,
		DeleteContext: resourceGithubRepositoryDeploymentBranchPolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryDeploymentBranchPolicyImport,
		},

		CustomizeDiff: diffETag,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The GitHub repository name.",
			},
			"environment_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The target environment name.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the branch",
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An etag representing the Branch object.",
			},
		},
	}
}

func resourceGithubRepositoryDeploymentBranchPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	if err := d.Set("etag", nil); err != nil {
		return diag.FromErr(err)
	}

	repoName := d.Get("repository").(string)
	environmentName := d.Get("environment_name").(string)
	name := d.Get("name").(string)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	_, _, err = client.Repositories.UpdateDeploymentBranchPolicy(ctx, owner, repoName, environmentName, int64(id), github.UpdateDeploymentBranchPolicyRequest{Name: name})
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubRepositoryDeploymentBranchPolicyRead(ctx, d, meta)
}

func resourceGithubRepositoryDeploymentBranchPolicyCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxId, d.Id())
	}

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	environmentName := d.Get("environment_name").(string)
	name := d.Get("name").(string)

	policy, _, err := client.Repositories.CreateDeploymentBranchPolicy(ctx, owner, repoName, environmentName, github.CreateDeploymentBranchPolicyRequest{Name: name, Type: new("branch")})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(*policy.ID, 10))

	return resourceGithubRepositoryDeploymentBranchPolicyRead(ctx, d, meta)
}

func resourceGithubRepositoryDeploymentBranchPolicyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ctx = context.WithValue(ctx, ctxId, d.Id())
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxEtag, d.Get("etag").(string))
	}

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	environmentName := d.Get("environment_name").(string)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	policy, resp, err := client.Repositories.GetDeploymentBranchPolicy(ctx, owner, repoName, environmentName, int64(id))
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				tflog.Info(ctx, "Removing deployment branch policy from state because it no longer exists in GitHub", map[string]any{"repository": repoName, "environment_name": environmentName})
				d.SetId("")
				return nil
			}
		}
		return diag.FromErr(err)
	}

	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("repository", repoName); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("environment_name", environmentName); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("name", policy.Name); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryDeploymentBranchPolicyDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ctx = context.WithValue(ctx, ctxId, d.Id())

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	environmentName := d.Get("environment_name").(string)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.Repositories.DeleteDeploymentBranchPolicy(ctx, owner, repoName, environmentName, int64(id))
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubRepositoryDeploymentBranchPolicyImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	repoName, environmentName, id, err := parseID3(d.Id())
	if err != nil {
		return nil, err
	}

	d.SetId(id)
	if err = d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err = d.Set("environment_name", environmentName); err != nil {
		return nil, err
	}

	if diags := resourceGithubRepositoryDeploymentBranchPolicyRead(ctx, d, meta); diags.HasError() {
		return nil, errors.New(diags[0].Summary)
	}

	return []*schema.ResourceData{d}, nil
}
