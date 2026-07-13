package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// maxInstallationRepositoriesPerRequest is the maximum number of repositories
// the enterprise organization-installations endpoints accept per request.
const maxInstallationRepositoriesPerRequest = 50

func resourceGithubEnterpriseAppInstallation() *schema.Resource {
	return &schema.Resource{
		Description: "Manage a GitHub App installation on an enterprise-owned organization. " +
			"This resource requires GitHub Enterprise Cloud or GitHub Enterprise Server 3.19+ and an authenticated user that is an enterprise owner.",
		CreateContext: resourceGithubEnterpriseAppInstallationCreate,
		ReadContext:   resourceGithubEnterpriseAppInstallationRead,
		UpdateContext: resourceGithubEnterpriseAppInstallationUpdate,
		DeleteContext: resourceGithubEnterpriseAppInstallationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise that owns the organization.",
			},
			"organization": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The login of the enterprise-owned organization to install the app on.",
			},
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The client ID of the GitHub App to install.",
			},
			"repository_selection": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validateValueFunc([]string{"all", "selected", "none"}),
				Description:      "The repositories the installation can access. Can be one of 'all', 'selected' or 'none'.",
			},
			"selected_repositories": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "The names of the repositories the installation can access when 'repository_selection' is 'selected'.",
			},
			"installation_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the installation.",
			},
			"app_slug": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The slug of the installed app.",
			},
		},

		CustomizeDiff: customdiff.All(
			// The API can only toggle an existing installation between 'all' and
			// 'selected'; transitions involving 'none' require a reinstall.
			customdiff.ForceNewIfChange("repository_selection", func(ctx context.Context, oldValue, newValue, meta any) bool {
				return oldValue.(string) == "none" || newValue.(string) == "none"
			}),
			func(ctx context.Context, d *schema.ResourceDiff, meta any) error {
				selection := d.Get("repository_selection").(string)
				repoCount := d.Get("selected_repositories").(*schema.Set).Len()
				if selection == "selected" && repoCount == 0 {
					return fmt.Errorf("'selected_repositories' must be set when 'repository_selection' is 'selected'")
				}
				if selection != "selected" && repoCount > 0 {
					return fmt.Errorf("'selected_repositories' can only be set when 'repository_selection' is 'selected'")
				}
				return nil
			},
		),
	}
}

func resourceGithubEnterpriseAppInstallationCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug := d.Get("enterprise_slug").(string)
	org := d.Get("organization").(string)
	clientID := d.Get("client_id").(string)

	repositories := expandStringList(d.Get("selected_repositories").(*schema.Set).List())

	// The install endpoint accepts at most maxInstallationRepositoriesPerRequest
	// repositories; any remainder is granted with follow-up requests.
	initial := repositories
	var remainder []string
	if len(repositories) > maxInstallationRepositoriesPerRequest {
		initial = repositories[:maxInstallationRepositoriesPerRequest]
		remainder = repositories[maxInstallationRepositoriesPerRequest:]
	}

	req := github.InstallAppRequest{
		ClientID:            clientID,
		RepositorySelection: d.Get("repository_selection").(string),
		Repositories:        initial,
	}

	tflog.Debug(ctx, "Installing app on enterprise-owned organization", map[string]any{
		"client_id":  clientID,
		"org":        org,
		"enterprise": enterpriseSlug,
	})
	installation, _, err := client.Enterprise.InstallApp(ctx, enterpriseSlug, org, req)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := addEnterpriseAppInstallationRepositories(ctx, client, enterpriseSlug, org, installation.GetID(), remainder); err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(enterpriseSlug, org, clientID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	return resourceGithubEnterpriseAppInstallationRead(ctx, d, meta)
}

func resourceGithubEnterpriseAppInstallationRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug, org, clientID, err := parseID3(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	installation, err := findEnterpriseAppInstallation(ctx, client, enterpriseSlug, org, clientID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing enterprise app installation from state because it no longer exists in GitHub", map[string]any{
				"id": d.Id(),
			})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	if installation == nil {
		tflog.Info(ctx, "Removing enterprise app installation from state because it no longer exists in GitHub", map[string]any{
			"id": d.Id(),
		})
		d.SetId("")
		return nil
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("organization", org); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("client_id", clientID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("repository_selection", installation.GetRepositorySelection()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("installation_id", strconv.FormatInt(installation.GetID(), 10)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("app_slug", installation.GetAppSlug()); err != nil {
		return diag.FromErr(err)
	}

	selectedRepositories := []string{}
	if installation.GetRepositorySelection() == "selected" {
		opts := &github.ListOptions{PerPage: maxPerPage}
		for {
			repos, resp, err := client.Enterprise.ListRepositoriesForOrgAppInstallation(ctx, enterpriseSlug, org, installation.GetID(), opts)
			if err != nil {
				return diag.FromErr(err)
			}
			for _, repo := range repos {
				selectedRepositories = append(selectedRepositories, repo.GetName())
			}
			if resp.NextPage == 0 {
				break
			}
			opts.Page = resp.NextPage
		}
	}
	if err := d.Set("selected_repositories", flattenStringList(selectedRepositories)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseAppInstallationUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug, org, _, err := parseID3(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	installationID, err := strconv.ParseInt(d.Get("installation_id").(string), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Get("installation_id").(string), err))
	}

	if d.HasChange("repository_selection") {
		selection := d.Get("repository_selection").(string)
		req := github.UpdateAppInstallationRepositoriesRequest{
			RepositorySelection: new(selection),
		}
		var remainder []string
		if selection == "selected" {
			repositories := expandStringList(d.Get("selected_repositories").(*schema.Set).List())
			// The toggle endpoint accepts at most
			// maxInstallationRepositoriesPerRequest repositories; any remainder
			// is granted with follow-up requests.
			req.Repositories = repositories
			if len(repositories) > maxInstallationRepositoriesPerRequest {
				req.Repositories = repositories[:maxInstallationRepositoriesPerRequest]
				remainder = repositories[maxInstallationRepositoriesPerRequest:]
			}
		}

		tflog.Debug(ctx, "Updating repository selection for enterprise app installation", map[string]any{
			"selection":       selection,
			"installation_id": installationID,
			"org":             org,
			"enterprise":      enterpriseSlug,
		})
		_, _, err := client.Enterprise.UpdateAppInstallationRepositories(ctx, enterpriseSlug, org, installationID, req)
		if err != nil {
			return diag.FromErr(err)
		}

		if err := addEnterpriseAppInstallationRepositories(ctx, client, enterpriseSlug, org, installationID, remainder); err != nil {
			return diag.FromErr(err)
		}
	} else if d.HasChange("selected_repositories") {
		oldRepos, newRepos := d.GetChange("selected_repositories")
		oldSet := oldRepos.(*schema.Set)
		newSet := newRepos.(*schema.Set)

		// Add before removing so the installation never has an empty selection.
		toAdd := expandStringList(newSet.Difference(oldSet).List())
		if err := addEnterpriseAppInstallationRepositories(ctx, client, enterpriseSlug, org, installationID, toAdd); err != nil {
			return diag.FromErr(err)
		}

		toRemove := expandStringList(oldSet.Difference(newSet).List())
		for chunk := range slices.Chunk(toRemove, maxInstallationRepositoriesPerRequest) {
			tflog.Debug(ctx, "Revoking enterprise app installation access to repositories", map[string]any{
				"installation_id": installationID,
				"org":             org,
				"enterprise":      enterpriseSlug,
				"repositories":    chunk,
			})
			_, _, err := client.Enterprise.RemoveRepositoriesFromAppInstallation(ctx, enterpriseSlug, org, installationID, github.AppInstallationRepositoriesRequest{Repositories: chunk})
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	return resourceGithubEnterpriseAppInstallationRead(ctx, d, meta)
}

func resourceGithubEnterpriseAppInstallationDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug, org, clientID, err := parseID3(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	installationID, err := strconv.ParseInt(d.Get("installation_id").(string), 10, 64)
	if err != nil {
		return diag.FromErr(unconvertibleIdErr(d.Get("installation_id").(string), err))
	}

	tflog.Debug(ctx, "Uninstalling app from enterprise-owned organization", map[string]any{
		"client_id":       clientID,
		"installation_id": installationID,
		"org":             org,
		"enterprise":      enterpriseSlug,
	})
	_, err = client.Enterprise.UninstallApp(ctx, enterpriseSlug, org, installationID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}

// addEnterpriseAppInstallationRepositories grants an installation access to
// the given repositories, in chunks the API accepts.
func addEnterpriseAppInstallationRepositories(ctx context.Context, client *github.Client, enterpriseSlug, org string, installationID int64, repositories []string) error {
	for chunk := range slices.Chunk(repositories, maxInstallationRepositoriesPerRequest) {
		tflog.Debug(ctx, "Granting enterprise app installation access to repositories", map[string]any{
			"installation_id": installationID,
			"org":             org,
			"enterprise":      enterpriseSlug,
			"repositories":    chunk,
		})
		_, _, err := client.Enterprise.AddRepositoriesToAppInstallation(ctx, enterpriseSlug, org, installationID, github.AppInstallationRepositoriesRequest{Repositories: chunk})
		if err != nil {
			return err
		}
	}

	return nil
}

// findEnterpriseAppInstallation returns the installation of the app with the
// given client ID on an enterprise-owned organization, or nil if the app is
// not installed.
func findEnterpriseAppInstallation(ctx context.Context, client *github.Client, enterpriseSlug, org, clientID string) (*github.Installation, error) {
	opts := &github.ListOptions{PerPage: maxPerPage}
	for {
		installations, resp, err := client.Enterprise.ListAppInstallations(ctx, enterpriseSlug, org, opts)
		if err != nil {
			return nil, err
		}
		for _, installation := range installations {
			if installation.GetClientID() == clientID {
				return installation, nil
			}
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return nil, nil
}
