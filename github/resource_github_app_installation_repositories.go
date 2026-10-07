package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubAppInstallationRepositories() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubAppInstallationRepositoriesCreateOrUpdate,
		ReadContext:   resourceGithubAppInstallationRepositoriesRead,
		UpdateContext: resourceGithubAppInstallationRepositoriesCreateOrUpdate,
		DeleteContext: resourceGithubAppInstallationRepositoriesDelete,
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
			"selected_repositories": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Required:    true,
				Description: "A list of repository names to install the app on.",
			},
		},
	}
}

func resourceGithubAppInstallationRepositoriesCreateOrUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	installationIDString := d.Get("installation_id").(string)
	selectedRepositories := d.Get("selected_repositories")

	ctx = context.WithValue(ctx, ctxId, installationIDString)

	selectedRepositoryNames := []string{}

	names := selectedRepositories.(*schema.Set).List()
	for _, name := range names {
		selectedRepositoryNames = append(selectedRepositoryNames, name.(string))
	}

	currentReposNameIDs, instID, err := getAllAccessibleRepos(ctx, meta, installationIDString)
	if err != nil {
		return diag.FromErr(err)
	}

	// Add repos that are not in the current state on GitHub
	for _, repoName := range selectedRepositoryNames {
		if _, ok := currentReposNameIDs[repoName]; ok {
			// If it already exists, remove it from the map so we can delete all that are left at the end
			delete(currentReposNameIDs, repoName)
		} else {
			repo, _, err := client.Repositories.Get(ctx, owner, repoName)
			if err != nil {
				return diag.FromErr(err)
			}
			repoID := repo.GetID()
			tflog.Debug(ctx, "Adding repository to app installation", map[string]any{"repository": repoName, "repository_id": repoID, "installation_id": instID})
			_, _, err = client.Apps.AddRepository(ctx, instID, repoID)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	// Remove repositories that existed on GitHub but not selectedRepositories
	// There is a github limitation that means we can't remove the last repository from an installation.
	// Therefore, we skip the first and delete the rest. The app will then need to be uninstalled via the GUI
	// as there is no current API endpoint for [un]installation. Ensure there is at least one repository remaining.
	if len(selectedRepositoryNames) >= 1 {
		for repoName, repoID := range currentReposNameIDs {
			tflog.Debug(ctx, "Removing repository from app installation", map[string]any{"repository": repoName, "repository_id": repoID, "installation_id": instID})
			_, err = client.Apps.RemoveRepository(ctx, instID, repoID)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	d.SetId(installationIDString)
	return resourceGithubAppInstallationRepositoriesRead(ctx, d, meta)
}

func resourceGithubAppInstallationRepositoriesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	installationIDString := d.Id()

	reposNameIDs, _, err := getAllAccessibleRepos(ctx, meta, installationIDString)
	if err != nil {
		return diag.FromErr(err)
	}

	repoNames := []string{}
	for name := range reposNameIDs {
		repoNames = append(repoNames, name)
	}

	if len(reposNameIDs) > 0 {
		if err = d.Set("installation_id", installationIDString); err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("selected_repositories", repoNames); err != nil {
			return diag.FromErr(err)
		}
		return nil
	}
	tflog.Info(ctx, "Removing app installation repository association from state because it no longer exists in GitHub", map[string]any{"resource_id": d.Id()})
	d.SetId("")
	return nil
}

func resourceGithubAppInstallationRepositoriesDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	installationIDString := d.Get("installation_id").(string)

	reposNameIDs, instID, err := getAllAccessibleRepos(ctx, meta, installationIDString)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.v3client
	ctx = context.WithValue(ctx, ctxId, installationIDString)

	// There is a github limitation that means we can't remove the last repository from an installation.
	// Therefore, we skip the first and delete the rest. The app will then need to be uninstalled via the GUI
	// as there is no current API endpoint for [un]installation.
	first := true
	for repoName, repoID := range reposNameIDs {
		if first {
			first = false
			tflog.Warn(ctx, "Cannot remove the last repository from an app installation due to API limitations. Manually uninstall the app to remove.", map[string]any{"repository": repoName, "repository_id": repoID, "installation_id": instID})
			continue
		} else {
			_, err = client.Apps.RemoveRepository(ctx, instID, repoID)
			tflog.Debug(ctx, "Removing repository from app installation", map[string]any{"repository": repoName, "repository_id": repoID, "installation_id": instID})
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}
	return nil
}

func getAllAccessibleRepos(ctx context.Context, meta *Owner, idString string) (map[string]int64, int64, error) {
	installationID, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return nil, 0, unconvertibleIdErr(idString, err)
	}

	ctx = context.WithValue(ctx, ctxId, idString)
	opt := &github.ListOptions{PerPage: meta.maxPerPage}
	client := meta.v3client

	allRepos := make(map[string]int64)

	for {
		repos, resp, err := client.Apps.ListUserRepos(ctx, installationID, opt)
		if err != nil {
			return nil, 0, err
		}
		for _, r := range repos.Repositories {
			allRepos[r.GetName()] = r.GetID()
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return allRepos, installationID, nil
}
