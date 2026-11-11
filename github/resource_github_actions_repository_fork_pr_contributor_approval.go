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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsRepositoryForkPRContributorApproval() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsRepositoryForkPRContributorApprovalCreate,
		ReadContext:   resourceGithubActionsRepositoryForkPRContributorApprovalRead,
		UpdateContext: resourceGithubActionsRepositoryForkPRContributorApprovalUpdate,
		DeleteContext: resourceGithubActionsRepositoryForkPRContributorApprovalDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsRepositoryForkPRContributorApprovalImport,
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The GitHub repository.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 100)),
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the GitHub repository.",
			},
			"approval_policy": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The policy controlling which fork PR contributors need maintainer approval before their workflows can run. Possible values are 'first_time_contributors_new_to_github', 'first_time_contributors', or 'all_external_contributors'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
					"first_time_contributors_new_to_github",
					"first_time_contributors",
					"all_external_contributors",
				}, false)),
			},
		},

		CustomizeDiff: diffRepository,
	}
}

func resourceGithubActionsRepositoryForkPRContributorApprovalCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	repoName, _ := d.Get("repository").(string)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	approvalPolicy, _ := d.Get("approval_policy").(string)
	if err := updateRepositoryForkPRContributorApproval(ctx, client, owner, repoName, approvalPolicy); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(repo.GetID(), 10))
	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubActionsRepositoryForkPRContributorApprovalRead(ctx, d, m)
}

func resourceGithubActionsRepositoryForkPRContributorApprovalRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	repoName, _ := d.Get("repository").(string)

	policy, _, err := client.Actions.GetForkPRContributorApprovalPermissions(ctx, owner, repoName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing fork PR contributor approval from state because the repository no longer exists", map[string]any{"owner": owner, "repository": repoName})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := d.Set("approval_policy", policy.ApprovalPolicy); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsRepositoryForkPRContributorApprovalUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	repoName, _ := d.Get("repository").(string)

	approvalPolicy, _ := d.Get("approval_policy").(string)
	if err := updateRepositoryForkPRContributorApproval(ctx, client, owner, repoName, approvalPolicy); err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubActionsRepositoryForkPRContributorApprovalRead(ctx, d, m)
}

func resourceGithubActionsRepositoryForkPRContributorApprovalDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	repoName, _ := d.Get("repository").(string)

	// The API has no "off" state for this policy, so deleting the resource
	// restores GitHub's default.
	err := updateRepositoryForkPRContributorApproval(ctx, client, owner, repoName, "first_time_contributors")
	if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
		return nil
	}

	return diag.FromErr(handleArchivedRepoDelete(err, "fork PR contributor approval", d.Id(), owner, repoName))
}

func resourceGithubActionsRepositoryForkPRContributorApprovalImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name
	repoName := d.Id()

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}

	d.SetId(strconv.FormatInt(repo.GetID(), 10))
	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func updateRepositoryForkPRContributorApproval(ctx context.Context, client *github.Client, owner, repoName, approvalPolicy string) error {
	_, err := client.Actions.UpdateForkPRContributorApprovalPermissions(ctx, owner, repoName, github.ContributorApprovalPermissions{
		ApprovalPolicy: approvalPolicy,
	})
	return err
}
