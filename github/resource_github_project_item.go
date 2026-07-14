package github

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/integrations/terraform-provider-github/v6/internal/application/projects"
	itemapplication "github.com/integrations/terraform-provider-github/v6/internal/application/projects/item"
	itemusecases "github.com/integrations/terraform-provider-github/v6/internal/application/projects/item/use-cases"
	itemgithub "github.com/integrations/terraform-provider-github/v6/internal/infrastructure/providers/github/projects/v2/item"
)

func resourceGithubProjectItem() *schema.Resource {
	return &schema.Resource{
		Description:   "Adds an issue or pull request to a GitHub Projects V2 project.",
		CreateContext: resourceGithubProjectItemCreate,
		ReadContext:   resourceGithubProjectItemRead,
		UpdateContext: resourceGithubProjectItemUpdate,
		DeleteContext: resourceGithubProjectItemDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Node ID of the Projects V2 project.",
			},
			"content_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Node ID of the issue or pull request.",
			},
			"archived": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether the project item is archived.",
			},
		},
	}
}

func resourceGithubProjectItemCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	item, err := itemusecases.NewAdd(itemgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, itemapplication.AddInput{
		ProjectID: projectV2Get[string](d, "project_id"), ContentID: projectV2Get[string](d, "content_id"), Archived: projectV2Get[bool](d, "archived"),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(item.ID)
	if err := setProjectV2ItemState(d, item); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubProjectItemRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	item, err := itemusecases.NewGet(itemgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, d.Id())
	if errors.Is(err, projects.ErrNotFound) {
		tflog.Info(ctx, "Removing project item from state because it no longer exists in GitHub", map[string]any{"id": d.Id()})
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	if err := setProjectV2ItemState(d, item); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubProjectItemUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	gateway := itemgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)
	var (
		item itemapplication.Result
		err  error
	)
	if projectV2Get[bool](d, "archived") {
		item, err = itemusecases.NewArchive(gateway).Run(ctx, projectV2Get[string](d, "project_id"), d.Id())
	} else {
		item, err = itemusecases.NewRestore(gateway).Run(ctx, projectV2Get[string](d, "project_id"), d.Id())
	}
	if err != nil {
		return diag.FromErr(err)
	}
	if err := setProjectV2ItemState(d, item); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubProjectItemDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := itemusecases.NewRemove(itemgithub.NewGateway(projectV2OwnerMetadata(meta).v4client)).Run(ctx, projectV2Get[string](d, "project_id"), d.Id())
	if err != nil && !errors.Is(err, projects.ErrNotFound) {
		return diag.FromErr(err)
	}
	return nil
}

func setProjectV2ItemState(d *schema.ResourceData, item itemapplication.Result) error {
	for key, value := range map[string]any{"project_id": item.ProjectID, "content_id": item.ContentID, "archived": item.Archived} {
		if err := d.Set(key, value); err != nil {
			return fmt.Errorf("setting %s: %w", key, err)
		}
	}
	return nil
}
