package github

import (
	"context"
	"log"
	"strconv"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsOrganizationVariableRepository() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"variable_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validateSecretNameFunc,
				Description:      "Name of the existing variable.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The repository ID that can access the organization variable.",
			},
		},

		CreateContext: resourceGithubActionsOrganizationVariableRepositoryCreate,
		ReadContext:   resourceGithubActionsOrganizationVariableRepositoryRead,
		DeleteContext: resourceGithubActionsOrganizationVariableRepositoryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsOrganizationVariableRepositoryImport,
		},
	}
}

func resourceGithubActionsOrganizationVariableRepositoryCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	variableName := d.Get("variable_name").(string)
	repoID := d.Get("repository_id").(int)

	repository := &github.Repository{
		ID: github.Ptr(int64(repoID)),
	}

	_, err := client.Actions.AddSelectedRepoToOrgVariable(ctx, owner, variableName, repository)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(variableName, strconv.Itoa(repoID))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	return nil
}

func resourceGithubActionsOrganizationVariableRepositoryRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	variableName := d.Get("variable_name").(string)
	repoID := int64(d.Get("repository_id").(int))

	opt := &github.ListOptions{
		PerPage: maxPerPage,
	}

	for {
		repos, resp, err := client.Actions.ListSelectedReposForOrgVariable(ctx, owner, variableName, opt)
		if err != nil {
			return diag.FromErr(err)
		}

		for _, repo := range repos.Repositories {
			if repo.GetID() == repoID {
				return nil
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	log.Printf("[INFO] Removing variable repository association %s from state because it no longer exists in GitHub", d.Id())
	d.SetId("")

	return nil
}

func resourceGithubActionsOrganizationVariableRepositoryDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}

	meta := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	variableName := d.Get("variable_name").(string)
	repoID := d.Get("repository_id").(int)

	repository := &github.Repository{
		ID: github.Ptr(int64(repoID)),
	}
	_, err := client.Actions.RemoveSelectedRepoFromOrgVariable(ctx, owner, variableName, repository)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsOrganizationVariableRepositoryImport(ctx context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
	variableName, repoIDStr, err := parseID2(d.Id())
	if err != nil {
		return nil, err
	}

	repoID, err := strconv.Atoi(repoIDStr)
	if err != nil {
		return nil, err
	}

	if err := d.Set("variable_name", variableName); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", repoID); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
