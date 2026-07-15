package github

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects"
	projectapplication "github.com/integrations/terraform-provider-github/v6/internal/application/projects/project"
	projectusecases "github.com/integrations/terraform-provider-github/v6/internal/application/projects/project/use-cases"
	projectgithub "github.com/integrations/terraform-provider-github/v6/internal/infrastructure/providers/github/projects/v2/project"
)

func resourceGithubProject() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages a GitHub Projects V2 project for an organization or user.",
		CreateContext: resourceGithubProjectCreate,
		ReadContext:   resourceGithubProjectRead,
		UpdateContext: resourceGithubProjectUpdate,
		DeleteContext: resourceGithubProjectDelete,
		CustomizeDiff: diffProjectV2Owner,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"owner_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{projectV2OwnerOrganization, projectV2OwnerUser}, false)),
				Description:      "Type of project owner. Must be `organization` or `user`.",
			},
			"owner": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Login of the project owner. Defaults to the provider owner.",
			},
			"owner_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Stable database ID of the project owner.",
			},
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Title of the project.",
			},
			"short_description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Short description of the project.",
			},
			"readme": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "README content of the project.",
			},
			"public": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether the project is public.",
			},
			"closed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether the project is closed.",
			},
			"number": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Project number scoped to the owner.",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the project.",
			},
		},
	}
}

func resourceGithubProjectCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	gateway := projectgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)
	project, err := projectusecases.NewCreate(gateway).Run(ctx, projectapplication.CreateInput{
		OwnerKind: projectapplication.OwnerKind(projectV2OwnerKind(d, meta)), Owner: projectV2OwnerLogin(d, meta),
		Title: projectV2Get[string](d, "title"), ShortDescription: projectV2Get[string](d, "short_description"), Readme: projectV2Get[string](d, "readme"),
		Public: projectV2Get[bool](d, "public"), Closed: projectV2Get[bool](d, "closed"),
	})
	if project.ID != "" {
		d.SetId(project.ID)
		if stateErr := setProjectV2State(d, project); stateErr != nil {
			if err != nil {
				return diag.FromErr(errors.Join(err, stateErr))
			}
			return diag.FromErr(stateErr)
		}
	}
	if err != nil {
		return diag.FromErr(err)
	}
	if project.ID == "" {
		return diag.Errorf("GitHub returned a Projects V2 project without an ID")
	}
	return nil
}

func resourceGithubProjectRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	project, err := projectusecases.NewGet(projectgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, d.Id())
	if errors.Is(err, projects.ErrNotFound) {
		tflog.Info(ctx, "Removing project from state because it no longer exists in GitHub", map[string]any{"id": d.Id()})
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	if err := setProjectV2State(d, project); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubProjectUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	update := projectapplication.UpdateInput{
		ID: d.Id(), Title: projectV2Get[string](d, "title"), ShortDescription: projectV2Get[string](d, "short_description"), Readme: projectV2Get[string](d, "readme"),
		Public: projectV2Get[bool](d, "public"), Closed: projectV2Get[bool](d, "closed"),
	}
	if err := projectusecases.NewUpdate(projectgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, update); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubProjectDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := projectusecases.NewDelete(projectgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, d.Id())
	if err != nil && !errors.Is(err, projects.ErrNotFound) {
		return diag.FromErr(err)
	}
	return nil
}

func setProjectV2State(d *schema.ResourceData, project projectapplication.Result) error {
	values := map[string]any{
		"owner_type": string(project.OwnerKind), "owner": project.Owner, "owner_id": project.OwnerID, "number": project.Number, "title": project.Title,
		"short_description": project.ShortDescription, "readme": project.Readme, "public": project.Public, "closed": project.Closed, "url": project.URL,
	}
	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			return fmt.Errorf("setting %s: %w", key, err)
		}
	}
	return nil
}

func projectV2OwnerLogin(d *schema.ResourceData, meta any) string {
	if owner := projectV2Get[string](d, "owner"); owner != "" {
		return owner
	}
	return projectV2OwnerMetadata(meta).name
}

func projectV2OwnerKind(d *schema.ResourceData, meta any) string {
	if kind := projectV2Get[string](d, "owner_type"); kind != "" {
		return kind
	}
	return projectV2DefaultOwnerKind(meta)
}
