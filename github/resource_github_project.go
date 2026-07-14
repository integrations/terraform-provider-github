package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/shurcooL/githubv4"
)

func resourceGithubProject() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages a GitHub Projects V2 project for an organization or user.",
		CreateContext: resourceGithubProjectCreate,
		ReadContext:   resourceGithubProjectRead,
		UpdateContext: resourceGithubProjectUpdate,
		DeleteContext: resourceGithubProjectDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"owner_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          projectV2OwnerOrganization,
				ForceNew:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{projectV2OwnerOrganization, projectV2OwnerUser}, false)),
				Description:      "Type of project owner. Must be `organization` or `user`.",
			},
			"owner": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Login of the project owner. Defaults to the provider owner.",
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
	client := meta.(*Owner).v4client
	ownerType := d.Get("owner_type").(string)
	owner := projectV2OwnerLogin(d, meta)
	ownerID, err := projectV2OwnerID(ctx, client, ownerType, owner)
	if err != nil {
		return diag.FromErr(err)
	}

	var mutation struct {
		CreateProjectV2 struct {
			Project projectV2Node `graphql:"projectV2"`
		} `graphql:"createProjectV2(input: $input)"`
	}
	input := githubv4.CreateProjectV2Input{
		OwnerID: ownerID,
		Title:   githubv4.String(d.Get("title").(string)),
	}
	if err := client.Mutate(ctx, &mutation, input, nil); err != nil {
		return diag.FromErr(err)
	}

	id := string(mutation.CreateProjectV2.Project.ID)
	if id == "" {
		return diag.Errorf("GitHub returned a Projects V2 project without an ID")
	}
	d.SetId(id)
	if diags := resourceGithubProjectUpdate(ctx, d, meta); diags.HasError() {
		return diags
	}
	return resourceGithubProjectRead(ctx, d, meta)
}

func resourceGithubProjectRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	project, err := queryProjectV2(ctx, meta.(*Owner).v4client, d.Id())
	if isProjectV2NotFound(err) {
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
	client := meta.(*Owner).v4client
	title := githubv4.String(d.Get("title").(string))
	shortDescription := githubv4.String(d.Get("short_description").(string))
	readme := githubv4.String(d.Get("readme").(string))
	public := githubv4.Boolean(d.Get("public").(bool))
	closed := githubv4.Boolean(d.Get("closed").(bool))
	input := githubv4.UpdateProjectV2Input{
		ProjectID:        githubv4.ID(d.Id()),
		Title:            &title,
		ShortDescription: &shortDescription,
		Readme:           &readme,
		Public:           &public,
		Closed:           &closed,
	}
	var mutation struct {
		UpdateProjectV2 struct {
			Project projectV2Node `graphql:"projectV2"`
		} `graphql:"updateProjectV2(input: $input)"`
	}
	if err := client.Mutate(ctx, &mutation, input, nil); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubProjectDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var mutation struct {
		DeleteProjectV2 struct {
			ClientMutationID githubv4.String
		} `graphql:"deleteProjectV2(input: $input)"`
	}
	input := githubv4.DeleteProjectV2Input{ProjectID: githubv4.ID(d.Id())}
	err := meta.(*Owner).v4client.Mutate(ctx, &mutation, input, nil)
	if err != nil && !isProjectV2NotFound(err) {
		return diag.FromErr(err)
	}
	return nil
}
