package github

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubAgentsOrganizationSecretRepository() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubAgentsOrganizationSecretRepositoryCreate,
		ReadContext:   resourceGithubAgentsOrganizationSecretRepositoryRead,
		DeleteContext: resourceGithubAgentsOrganizationSecretRepositoryDelete,
		Importer:      &schema.ResourceImporter{StateContext: resourceGithubAgentsOrganizationSecretRepositoryImport},
		Description:   "Manages one repository's access to an organization agent secret. Requires selected visibility.",
		Schema: map[string]*schema.Schema{
			"secret_name":   {Type: schema.TypeString, Required: true, ForceNew: true, ValidateDiagFunc: validateSecretNameFunc, Description: "Name of the existing organization agent secret."},
			"repository_id": {Type: schema.TypeInt, Required: true, ForceNew: true, ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)), Description: "ID of the repository allowed to access the secret."},
		},
	}
}

func resourceGithubAgentsOrganizationSecretRepositoryCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}
	meta, _ := m.(*Owner)
	name, _ := d.Get("secret_name").(string)
	id, _ := d.Get("repository_id").(int)
	if _, err := meta.v3client.Agents.AddSelectedRepoToOrgSecret(ctx, meta.name, name, int64(id)); err != nil {
		return diag.FromErr(err)
	}
	resourceID, err := buildID(name, strconv.Itoa(id))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(resourceID)
	return nil
}

func resourceGithubAgentsOrganizationSecretRepositoryRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}
	meta, _ := m.(*Owner)
	name, _ := d.Get("secret_name").(string)
	id, _ := d.Get("repository_id").(int)
	for repo, err := range meta.v3client.Agents.ListSelectedReposForOrgSecretIter(ctx, meta.name, name, &github.ListOptions{PerPage: meta.maxPerPage}) {
		if err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
				d.SetId("")
				return nil
			}
			return diag.FromErr(err)
		}
		if repo.GetID() == int64(id) {
			return nil
		}
	}
	d.SetId("")
	return nil
}

func resourceGithubAgentsOrganizationSecretRepositoryDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}
	meta, _ := m.(*Owner)
	name, _ := d.Get("secret_name").(string)
	id, _ := d.Get("repository_id").(int)
	_, err := meta.v3client.Agents.RemoveSelectedRepoFromOrgSecret(ctx, meta.name, name, int64(id))
	if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
		return nil
	}
	return diag.FromErr(err)
}

func resourceGithubAgentsOrganizationSecretRepositoryImport(_ context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	if err := checkOrganization(m); err != nil {
		return nil, err
	}
	name, rawID, err := parseID2(d.Id())
	if err != nil {
		return nil, err
	}
	id, err := strconv.Atoi(rawID)
	if err != nil {
		return nil, err
	}
	if err := d.Set("secret_name", name); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", id); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}
