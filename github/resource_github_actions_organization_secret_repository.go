package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsOrganizationSecretRepository() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validateSecretNameFunc,
				Description:      "Name of the existing secret.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The repository ID that can access the organization secret.",
			},
		},

		CreateContext: resourceGithubActionsOrganizationSecretRepositoryCreate,
		ReadContext:   resourceGithubActionsOrganizationSecretRepositoryRead,
		DeleteContext: resourceGithubActionsOrganizationSecretRepositoryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsOrganizationSecretRepositoryImport,
		},
	}
}

func resourceGithubActionsOrganizationSecretRepositoryCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName, _ := d.Get("secret_name").(string)
	repoIDInt, _ := d.Get("repository_id").(int)
	repoID := int64(repoIDInt)

	if _, err := client.Actions.AddSelectedRepoToOrgSecret(ctx, owner, secretName, repoID); err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(secretName, strconv.Itoa(repoIDInt))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	return nil
}

func resourceGithubActionsOrganizationSecretRepositoryRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName, _ := d.Get("secret_name").(string)
	repoIDInt, _ := d.Get("repository_id").(int)
	repoID := int64(repoIDInt)

	for repo, err := range client.Actions.ListSelectedReposForOrgSecretIter(ctx, owner, secretName, &github.ListOptions{PerPage: maxPerPage}) {
		if err != nil {
			return diag.FromErr(err)
		}

		if repo.GetID() == repoID {
			return nil
		}
	}

	tflog.Info(ctx, "Removing secret repository association from state because it no longer exists in GitHub", map[string]any{"secret_name": secretName, "repository_id": repoIDInt})
	d.SetId("")

	return nil
}

func resourceGithubActionsOrganizationSecretRepositoryDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	secretName, _ := d.Get("secret_name").(string)
	repoIDInt, _ := d.Get("repository_id").(int)
	repoID := int64(repoIDInt)

	if _, err := client.Actions.RemoveSelectedRepoFromOrgSecret(ctx, owner, secretName, repoID); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsOrganizationSecretRepositoryImport(ctx context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
	secretName, repoIDStr, err := parseID2(d.Id())
	if err != nil {
		return nil, err
	}

	repoID, err := strconv.Atoi(repoIDStr)
	if err != nil {
		return nil, err
	}

	if err := d.Set("secret_name", secretName); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", repoID); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
