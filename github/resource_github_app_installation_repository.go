package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubAppInstallationRepository() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubAppInstallationRepositoryCreate,
		ReadContext:   resourceGithubAppInstallationRepositoryRead,
		DeleteContext: resourceGithubAppInstallationRepositoryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"installation_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The GitHub app installation id.",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The repository to install the app on.",
			},
			"repo_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceGithubAppInstallationRepositoryCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	installationIDString := d.Get("installation_id").(string)
	installationID, err := strconv.ParseInt(installationIDString, 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(installationIDString, err))
	}

	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)
	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	repoID := repo.GetID()

	_, _, err = client.Apps.AddRepository(ctx, installationID, repoID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildTwoPartID(installationIDString, repoName))
	return resourceGithubAppInstallationRepositoryRead(ctx, d, meta)
}

func resourceGithubAppInstallationRepositoryRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	installationIDString, repoName, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	installationID, err := strconv.ParseInt(installationIDString, 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(installationIDString, err))
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())
	opt := &github.ListOptions{PerPage: meta.maxPerPage}

	for {
		repos, resp, err := client.Apps.ListUserRepos(ctx, installationID, opt)
		if err != nil {
			return diag.FromErr(err)
		}

		for _, r := range repos.Repositories {
			if r.GetName() == repoName {
				if err = d.Set("installation_id", installationIDString); err != nil {
					return diag.FromErr(err)
				}
				if err = d.Set("repository", repoName); err != nil {
					return diag.FromErr(err)
				}
				if err = d.Set("repo_id", r.GetID()); err != nil {
					return diag.FromErr(err)
				}
				return nil
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	tflog.Info(ctx, "Removing app installation repository association from state because it no longer exists in GitHub", map[string]any{"resource_id": d.Id()})
	d.SetId("")
	return nil
}

func resourceGithubAppInstallationRepositoryDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	installationIDString := d.Get("installation_id").(string)
	installationID, err := strconv.ParseInt(installationIDString, 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(installationIDString, err))
	}

	client := meta.v3client

	repoID := d.Get("repo_id").(int)

	_, err = client.Apps.RemoveRepository(ctx, installationID, int64(repoID))
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
