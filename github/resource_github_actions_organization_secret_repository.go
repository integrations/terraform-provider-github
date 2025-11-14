package github

import (
	"context"
	"log"
	"strconv"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubActionsOrganizationSecretRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubActionsOrganizationSecretRepositoryCreate,
		Read:   resourceGithubActionsOrganizationSecretRepositoryRead,
		Delete: resourceGithubActionsOrganizationSecretRepositoryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"secret_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Name of the existing secret.",
				ValidateDiagFunc: validateSecretNameFunc,
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The repository ID that can access the organization secret.",
			},
		},
	}
}

func resourceGithubActionsOrganizationSecretRepositoryCreate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	repositoryID := d.Get("repository_id").(int)
	secretName := d.Get("secret_name").(string)

	repoIDInt64 := int64(repositoryID)
	repository := &github.Repository{
		ID: &repoIDInt64,
	}

	_, err = client.Actions.AddSelectedRepoToOrgSecret(ctx, owner, secretName, repository)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(secretName, strconv.Itoa(repositoryID)))
	return resourceGithubActionsOrganizationSecretRepositoryRead(d, meta)
}

func resourceGithubActionsOrganizationSecretRepositoryRead(d *schema.ResourceData, meta any) error {
	owner := meta.(*Owner).name

	err := checkOrganization(meta)
	if err != nil {
		return err
	}
	client := meta.(*Owner).v3client

	secretName, repositoryIDString, err := parseTwoPartID(d.Id(), "secret_name", "repository_id")
	if err != nil {
		return err
	}

	repositoryID, err := strconv.ParseInt(repositoryIDString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(repositoryIDString, err)
	}

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	opt := &github.ListOptions{
		PerPage: maxPerPage,
	}
	for {
		repos, resp, err := client.Actions.ListSelectedReposForOrgSecret(ctx, owner, secretName, opt)
		if err != nil {
			return err
		}

		for _, repo := range repos.Repositories {
			if repo.GetID() == repositoryID {
				if err = d.Set("secret_name", secretName); err != nil {
					return err
				}
				if err = d.Set("repository_id", repositoryID); err != nil {
					return err
				}
				return nil
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	log.Printf("[INFO] Removing secret repository association %s from state because it no longer exists in GitHub",
		d.Id())
	d.SetId("")

	return nil
}

func resourceGithubActionsOrganizationSecretRepositoryDelete(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	secretName := d.Get("secret_name").(string)
	repositoryID := d.Get("repository_id").(int)

	repoIDInt64 := int64(repositoryID)
	repository := &github.Repository{
		ID: &repoIDInt64,
	}
	_, err = client.Actions.RemoveSelectedRepoFromOrgSecret(ctx, owner, secretName, repository)
	if err != nil {
		return err
	}

	return nil
}
