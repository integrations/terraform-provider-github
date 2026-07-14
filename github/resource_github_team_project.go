package github

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects"
	linkapplication "github.com/integrations/terraform-provider-github/v6/internal/application/projects/team/link"
	linkusecases "github.com/integrations/terraform-provider-github/v6/internal/application/projects/team/link/use-cases"
	linkgithub "github.com/integrations/terraform-provider-github/v6/internal/infrastructure/providers/github/projects/v2/team/link"
)

func resourceGithubTeamProject() *schema.Resource {
	return &schema.Resource{
		Description:   "Links a team to a GitHub Projects V2 project.",
		CreateContext: resourceGithubTeamProjectCreate,
		ReadContext:   resourceGithubTeamProjectRead,
		DeleteContext: resourceGithubTeamProjectDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubTeamProjectImport,
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Node ID of the Projects V2 project.",
			},
			"organization": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Organization login. Defaults to the provider owner.",
			},
			"team_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Team slug.",
			},
		},
	}
}

func resourceGithubTeamProjectCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	link, err := linkusecases.NewAttach(linkgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, linkapplication.AttachInput{
		ProjectID: projectV2Get[string](d, "project_id"), Organization: projectV2TeamOrganization(d, meta), Slug: projectV2Get[string](d, "team_slug"),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	id, err := buildID(link.ProjectID, link.TeamID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)
	if err := setProjectV2TeamState(d, link); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubTeamProjectRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	projectID, teamID, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	link, err := linkusecases.NewGet(linkgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, projectID, teamID)
	if errors.Is(err, projects.ErrNotFound) {
		tflog.Info(ctx, "Removing team project link from state because it no longer exists in GitHub", map[string]any{"id": d.Id()})
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	if err := setProjectV2TeamState(d, link); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubTeamProjectDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	projectID, teamID, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	err = linkusecases.NewDetach(linkgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, projectID, teamID)
	if err != nil && !errors.Is(err, projects.ErrNotFound) {
		return diag.FromErr(err)
	}
	return nil
}

func setProjectV2TeamState(d *schema.ResourceData, link linkapplication.Result) error {
	for key, value := range map[string]any{"project_id": link.ProjectID, "organization": link.Organization, "team_slug": link.Slug} {
		if err := d.Set(key, value); err != nil {
			return fmt.Errorf("setting %s: %w", key, err)
		}
	}
	return nil
}

func resourceGithubTeamProjectImport(_ context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
	if _, _, err := parseID2(d.Id()); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func projectV2TeamOrganization(d *schema.ResourceData, meta any) string {
	if organization := projectV2Get[string](d, "organization"); organization != "" {
		return organization
	}
	return projectV2OwnerMetadata(meta).name
}
