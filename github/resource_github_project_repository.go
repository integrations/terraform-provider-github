package github

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects"
	linkapplication "github.com/integrations/terraform-provider-github/v6/internal/application/projects/repository/link"
	linkusecases "github.com/integrations/terraform-provider-github/v6/internal/application/projects/repository/link/use-cases"
	linkgithub "github.com/integrations/terraform-provider-github/v6/internal/infrastructure/providers/github/projects/v2/repository/link"
)

func resourceGithubProjectRepository() *schema.Resource {
	return &schema.Resource{
		Description:   "Links a repository to a GitHub Projects V2 project.",
		CreateContext: resourceGithubProjectRepositoryCreate,
		ReadContext:   resourceGithubProjectRepositoryRead,
		DeleteContext: resourceGithubProjectRepositoryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubProjectRepositoryImport,
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Node ID of the Projects V2 project.",
			},
			"repository_owner": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Repository owner. Defaults to the provider owner.",
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Repository name.",
			},
		},
	}
}

func resourceGithubProjectRepositoryCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	link, err := linkusecases.NewAttach(linkgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, linkapplication.AttachInput{
		ProjectID: projectV2Get[string](d, "project_id"), Owner: projectV2RepositoryOwner(d, meta), Name: projectV2Get[string](d, "repository"),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	id, err := buildID(link.ProjectID, link.RepositoryID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)
	if err := setProjectV2RepositoryState(d, link); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubProjectRepositoryRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	projectID, repositoryID, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	link, err := linkusecases.NewGet(linkgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, projectID, repositoryID)
	if errors.Is(err, projects.ErrNotFound) {
		tflog.Info(ctx, "Removing project repository link from state because it no longer exists in GitHub", map[string]any{"id": d.Id()})
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	if err := setProjectV2RepositoryState(d, link); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubProjectRepositoryDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	projectID, repositoryID, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	err = linkusecases.NewDetach(linkgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, projectID, repositoryID)
	if err != nil && !errors.Is(err, projects.ErrNotFound) {
		return diag.FromErr(err)
	}
	return nil
}

func setProjectV2RepositoryState(d *schema.ResourceData, link linkapplication.Result) error {
	for key, value := range map[string]any{"project_id": link.ProjectID, "repository_owner": link.Owner, "repository": link.Name} {
		if err := d.Set(key, value); err != nil {
			return fmt.Errorf("setting %s: %w", key, err)
		}
	}
	return nil
}

func resourceGithubProjectRepositoryImport(_ context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
	if _, _, err := parseID2(d.Id()); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func projectV2RepositoryOwner(d *schema.ResourceData, meta any) string {
	if owner := projectV2Get[string](d, "repository_owner"); owner != "" {
		return owner
	}
	return projectV2OwnerMetadata(meta).name
}
