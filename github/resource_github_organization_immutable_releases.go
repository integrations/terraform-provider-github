package github

import (
	"context"
	"errors"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubOrganizationImmutableReleases() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubOrganizationImmutableReleasesCreateOrUpdate,
		Read:   resourceGithubOrganizationImmutableReleasesRead,
		Update: resourceGithubOrganizationImmutableReleasesCreateOrUpdate,
		Delete: resourceGithubOrganizationImmutableReleasesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"enforced_repositories": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The policy that controls which repositories in the organization have immutable releases enforced. Can be one of: 'all', 'none', or 'selected'.",
				ValidateDiagFunc: toDiagFunc(validation.StringInSlice([]string{"all", "none", "selected"}, false), "enforced_repositories"),
			},
			"selected_repository_ids": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "An array of repository IDs for which immutable releases enforcement should be applied. Only valid when 'enforced_repositories' is set to 'selected'.",
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func resourceGithubOrganizationImmutableReleasesCreateOrUpdate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()
	if !d.IsNewResource() {
		ctx = context.WithValue(ctx, ctxId, d.Id())
	}

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	enforcedRepositories := d.Get("enforced_repositories").(string)

	policy := github.ImmutableReleasePolicy{
		EnforcedRepositories: &enforcedRepositories,
	}

	if enforcedRepositories == "selected" {
		repoIDs, err := expandSelectedRepositoryIDs(d)
		if err != nil {
			return err
		}
		policy.SelectedRepositoryIDs = repoIDs
	}

	_, err = client.Organizations.UpdateImmutableReleasesSettings(ctx, orgName, policy)
	if err != nil {
		return err
	}

	d.SetId(orgName)
	return resourceGithubOrganizationImmutableReleasesRead(d, meta)
}

func resourceGithubOrganizationImmutableReleasesRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	settings, _, err := client.Organizations.GetImmutableReleasesSettings(ctx, d.Id())
	if err != nil {
		return err
	}

	if err = d.Set("enforced_repositories", settings.GetEnforcedRepositories()); err != nil {
		return err
	}

	if settings.GetEnforcedRepositories() == "selected" {
		opts := github.ListOptions{PerPage: 100, Page: 1}
		var repoIDs []int64

		for {
			repos, resp, err := client.Organizations.ListImmutableReleaseRepositories(ctx, d.Id(), &opts)
			if err != nil {
				return err
			}
			for _, repo := range repos.Repositories {
				repoIDs = append(repoIDs, repo.GetID())
			}

			if resp.NextPage == 0 {
				break
			}
			opts.Page = resp.NextPage
		}

		if err = d.Set("selected_repository_ids", repoIDs); err != nil {
			return err
		}
	} else {
		if err = d.Set("selected_repository_ids", []int64{}); err != nil {
			return err
		}
	}

	return nil
}

func resourceGithubOrganizationImmutableReleasesDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	none := "none"
	_, err = client.Organizations.UpdateImmutableReleasesSettings(ctx, orgName, github.ImmutableReleasePolicy{
		EnforcedRepositories: &none,
	})
	if err != nil {
		return err
	}

	return nil
}

func expandSelectedRepositoryIDs(d *schema.ResourceData) ([]int64, error) {
	raw, ok := d.GetOk("selected_repository_ids")
	if !ok {
		return nil, errors.New("selected_repository_ids must be specified when enforced_repositories is set to 'selected'")
	}

	set := raw.(*schema.Set)
	ids := make([]int64, 0, set.Len())
	for _, v := range set.List() {
		ids = append(ids, int64(v.(int)))
	}

	return ids, nil
}
