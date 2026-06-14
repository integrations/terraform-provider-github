package github

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseAppInstallation() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubEnterpriseAppInstallationCreate,
		Read:   resourceGithubEnterpriseAppInstallationRead,
		Update: resourceGithubEnterpriseAppInstallationUpdate,
		Delete: resourceGithubEnterpriseAppInstallationDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGithubEnterpriseAppInstallationImport,
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
				Description: "The Client ID of the GitHub App to install.",
			},
			"repository_selection": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Which repositories the app can access. One of 'all', 'selected', or 'none'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "selected", "none"}, false)),
			},
			"repositories": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "Repository names the installation should have access to. Only used when repository_selection is 'selected'.",
			},
			"installation_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the GitHub App installation.",
			},
			"app_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the GitHub App.",
			},
			"app_slug": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL-friendly name of the GitHub App.",
			},
		},
	}
}

func resourceGithubEnterpriseAppInstallationCreate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	enterprise := d.Get("enterprise_slug").(string)
	org := d.Get("organization").(string)
	selection := d.Get("repository_selection").(string)

	req := github.InstallAppRequest{
		ClientID:            d.Get("client_id").(string),
		RepositorySelection: selection,
	}
	if selection == "selected" {
		req.Repositories = expandStringList(d.Get("repositories").(*schema.Set).List())
	}

	installation, _, err := client.Enterprise.InstallApp(ctx, enterprise, org, req)
	if err != nil {
		return fmt.Errorf("error installing GitHub App on %s/%s: %w", enterprise, org, err)
	}

	d.SetId(buildThreePartID(enterprise, org, strconv.FormatInt(installation.GetID(), 10)))
	return resourceGithubEnterpriseAppInstallationRead(d, meta)
}

func resourceGithubEnterpriseAppInstallationRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	enterprise, org, idStr, err := parseID3(d.Id())
	if err != nil {
		return err
	}
	installationID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return unconvertibleIdErr(idStr, err)
	}

	installation, err := findEnterpriseAppInstallation(ctx, client, enterprise, org, installationID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotFound {
			log.Printf("[INFO] Removing enterprise app installation %s from state because it no longer exists in GitHub", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}
	if installation == nil {
		log.Printf("[INFO] Removing enterprise app installation %s from state because it no longer exists in GitHub", d.Id())
		d.SetId("")
		return nil
	}

	if err = d.Set("enterprise_slug", enterprise); err != nil {
		return err
	}
	if err = d.Set("organization", org); err != nil {
		return err
	}
	if err = d.Set("installation_id", installation.GetID()); err != nil {
		return err
	}
	if err = d.Set("app_id", installation.GetAppID()); err != nil {
		return err
	}
	if err = d.Set("app_slug", installation.GetAppSlug()); err != nil {
		return err
	}
	if v := installation.GetRepositorySelection(); v != "" {
		if err = d.Set("repository_selection", v); err != nil {
			return err
		}
	}

	if installation.GetRepositorySelection() == "selected" {
		repos, err := listEnterpriseAppInstallationRepositories(ctx, client, enterprise, org, installationID)
		if err != nil {
			return err
		}
		names := make([]string, 0, len(repos))
		for _, r := range repos {
			names = append(names, r.Name)
		}
		if err = d.Set("repositories", names); err != nil {
			return err
		}
	} else {
		if err = d.Set("repositories", []string{}); err != nil {
			return err
		}
	}

	return nil
}

func resourceGithubEnterpriseAppInstallationUpdate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	enterprise, org, idStr, err := parseID3(d.Id())
	if err != nil {
		return err
	}
	installationID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return unconvertibleIdErr(idStr, err)
	}

	selection := d.Get("repository_selection").(string)

	if d.HasChange("repository_selection") {
		opts := github.UpdateAppInstallationRepositoriesRequest{
			RepositorySelection: &selection,
		}
		if selection == "selected" {
			opts.Repositories = expandStringList(d.Get("repositories").(*schema.Set).List())
		}
		if _, _, err := client.Enterprise.UpdateAppInstallationRepositories(ctx, enterprise, org, installationID, opts); err != nil {
			return fmt.Errorf("error updating repository_selection for installation %d: %w", installationID, err)
		}
	} else if selection == "selected" && d.HasChange("repositories") {
		oldVal, newVal := d.GetChange("repositories")
		oldSet := oldVal.(*schema.Set)
		newSet := newVal.(*schema.Set)

		toAdd := expandStringList(newSet.Difference(oldSet).List())
		toRemove := expandStringList(oldSet.Difference(newSet).List())

		if len(toAdd) > 0 {
			if _, _, err := client.Enterprise.AddRepositoriesToAppInstallation(ctx, enterprise, org, installationID, github.AppInstallationRepositoriesRequest{Repositories: toAdd}); err != nil {
				return fmt.Errorf("error adding repositories to installation %d: %w", installationID, err)
			}
		}
		if len(toRemove) > 0 {
			if _, _, err := client.Enterprise.RemoveRepositoriesFromAppInstallation(ctx, enterprise, org, installationID, github.AppInstallationRepositoriesRequest{Repositories: toRemove}); err != nil {
				return fmt.Errorf("error removing repositories from installation %d: %w", installationID, err)
			}
		}
	}

	return resourceGithubEnterpriseAppInstallationRead(d, meta)
}

func resourceGithubEnterpriseAppInstallationDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	enterprise, org, idStr, err := parseID3(d.Id())
	if err != nil {
		return err
	}
	installationID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return unconvertibleIdErr(idStr, err)
	}

	_, err = client.Enterprise.UninstallApp(ctx, enterprise, org, installationID)
	return err
}

func resourceGithubEnterpriseAppInstallationImport(d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	if _, _, _, err := parseID3(d.Id()); err != nil {
		return nil, fmt.Errorf("invalid ID specified: supplied ID must be written as <enterprise_slug>:<organization>:<installation_id>")
	}
	if err := resourceGithubEnterpriseAppInstallationRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

// findEnterpriseAppInstallation walks the enterprise app installations on an organization
// to locate the installation matching the given ID. The REST API does not provide a direct
// GET endpoint for a single enterprise-owned org installation, so a list-and-filter is used.
func findEnterpriseAppInstallation(ctx context.Context, client *github.Client, enterprise, org string, installationID int64) (*github.Installation, error) {
	opts := &github.ListOptions{PerPage: maxPerPage}
	for {
		installations, resp, err := client.Enterprise.ListAppInstallations(ctx, enterprise, org, opts)
		if err != nil {
			return nil, err
		}
		for _, inst := range installations {
			if inst.GetID() == installationID {
				return inst, nil
			}
		}
		if resp.NextPage == 0 {
			return nil, nil
		}
		opts.Page = resp.NextPage
	}
}

func listEnterpriseAppInstallationRepositories(ctx context.Context, client *github.Client, enterprise, org string, installationID int64) ([]*github.AccessibleRepository, error) {
	var all []*github.AccessibleRepository
	opts := &github.ListOptions{PerPage: maxPerPage}
	for {
		repos, resp, err := client.Enterprise.ListRepositoriesForOrgAppInstallation(ctx, enterprise, org, installationID, opts)
		if err != nil {
			return nil, err
		}
		all = append(all, repos...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return all, nil
}
