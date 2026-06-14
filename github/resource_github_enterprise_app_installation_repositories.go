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
)

func resourceGithubEnterpriseAppInstallationRepositories() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubEnterpriseAppInstallationRepositoriesCreateOrUpdate,
		Read:   resourceGithubEnterpriseAppInstallationRepositoriesRead,
		Update: resourceGithubEnterpriseAppInstallationRepositoriesCreateOrUpdate,
		Delete: resourceGithubEnterpriseAppInstallationRepositoriesDelete,
		Importer: &schema.ResourceImporter{
			State: resourceGithubEnterpriseAppInstallationRepositoriesImport,
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
				Description: "The login of the enterprise-owned organization the app is installed on.",
			},
			"installation_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the GitHub App installation.",
			},
			"selected_repositories": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "A set of repository names the installation should have access to.",
			},
		},
	}
}

func resourceGithubEnterpriseAppInstallationRepositoriesCreateOrUpdate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	enterprise := d.Get("enterprise_slug").(string)
	org := d.Get("organization").(string)
	installationIDString := d.Get("installation_id").(string)
	installationID, err := strconv.ParseInt(installationIDString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(installationIDString, err)
	}

	desired := stringSetFromAny(d.Get("selected_repositories").(*schema.Set).List())

	current, err := listEnterpriseAppInstallationRepositories(ctx, client, enterprise, org, installationID)
	if err != nil {
		return err
	}
	currentSet := make(map[string]struct{}, len(current))
	for _, r := range current {
		currentSet[r.Name] = struct{}{}
	}

	var toAdd, toRemove []string
	for name := range desired {
		if _, ok := currentSet[name]; !ok {
			toAdd = append(toAdd, name)
		}
	}
	for name := range currentSet {
		if _, ok := desired[name]; !ok {
			toRemove = append(toRemove, name)
		}
	}

	if len(toAdd) > 0 {
		log.Printf("[DEBUG] Adding %d repositories to enterprise app installation %d", len(toAdd), installationID)
		if _, _, err := client.Enterprise.AddRepositoriesToAppInstallation(ctx, enterprise, org, installationID, github.AppInstallationRepositoriesRequest{Repositories: toAdd}); err != nil {
			return fmt.Errorf("error adding repositories to enterprise app installation %d: %w", installationID, err)
		}
	}
	if len(toRemove) > 0 {
		log.Printf("[DEBUG] Removing %d repositories from enterprise app installation %d", len(toRemove), installationID)
		if _, _, err := client.Enterprise.RemoveRepositoriesFromAppInstallation(ctx, enterprise, org, installationID, github.AppInstallationRepositoriesRequest{Repositories: toRemove}); err != nil {
			return fmt.Errorf("error removing repositories from enterprise app installation %d: %w", installationID, err)
		}
	}

	d.SetId(buildThreePartID(enterprise, org, installationIDString))
	return resourceGithubEnterpriseAppInstallationRepositoriesRead(d, meta)
}

func resourceGithubEnterpriseAppInstallationRepositoriesRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	enterprise, org, installationIDString, err := parseID3(d.Id())
	if err != nil {
		return err
	}
	installationID, err := strconv.ParseInt(installationIDString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(installationIDString, err)
	}

	repos, err := listEnterpriseAppInstallationRepositories(ctx, client, enterprise, org, installationID)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotFound {
			log.Printf("[INFO] Removing enterprise app installation repositories %s from state because the installation no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return err
	}

	names := make([]string, 0, len(repos))
	for _, r := range repos {
		names = append(names, r.Name)
	}

	if err = d.Set("enterprise_slug", enterprise); err != nil {
		return err
	}
	if err = d.Set("organization", org); err != nil {
		return err
	}
	if err = d.Set("installation_id", installationIDString); err != nil {
		return err
	}
	if err = d.Set("selected_repositories", names); err != nil {
		return err
	}
	return nil
}

func resourceGithubEnterpriseAppInstallationRepositoriesDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	enterprise, org, installationIDString, err := parseID3(d.Id())
	if err != nil {
		return err
	}
	installationID, err := strconv.ParseInt(installationIDString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(installationIDString, err)
	}

	current, err := listEnterpriseAppInstallationRepositories(ctx, client, enterprise, org, installationID)
	if err != nil {
		return err
	}
	if len(current) == 0 {
		return nil
	}

	names := make([]string, 0, len(current))
	for _, r := range current {
		names = append(names, r.Name)
	}

	_, _, err = client.Enterprise.RemoveRepositoriesFromAppInstallation(ctx, enterprise, org, installationID, github.AppInstallationRepositoriesRequest{Repositories: names})
	return err
}

func resourceGithubEnterpriseAppInstallationRepositoriesImport(d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	if _, _, _, err := parseID3(d.Id()); err != nil {
		return nil, fmt.Errorf("invalid ID specified: supplied ID must be written as <enterprise_slug>:<organization>:<installation_id>")
	}
	if err := resourceGithubEnterpriseAppInstallationRepositoriesRead(d, meta); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func stringSetFromAny(vs []any) map[string]struct{} {
	out := make(map[string]struct{}, len(vs))
	for _, v := range vs {
		if s, ok := v.(string); ok && s != "" {
			out[s] = struct{}{}
		}
	}
	return out
}
