package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

type projectV2ItemNode struct {
	ID         githubv4.String
	IsArchived githubv4.Boolean
	Project    struct {
		ID githubv4.String
	}
	Content struct {
		Issue struct {
			ID githubv4.String
		} `graphql:"... on Issue"`
		PullRequest struct {
			ID githubv4.String
		} `graphql:"... on PullRequest"`
		DraftIssue struct {
			ID githubv4.String
		} `graphql:"... on DraftIssue"`
	}
}

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
	var mutation struct {
		AddProjectV2ItemByID struct {
			Item projectV2ItemNode
		} `graphql:"addProjectV2ItemById(input: $input)"`
	}
	input := githubv4.AddProjectV2ItemByIdInput{
		ProjectID: githubv4.ID(d.Get("project_id").(string)),
		ContentID: githubv4.ID(d.Get("content_id").(string)),
	}
	if err := meta.(*Owner).v4client.Mutate(ctx, &mutation, input, nil); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(string(mutation.AddProjectV2ItemByID.Item.ID))
	if d.Get("archived").(bool) {
		return resourceGithubProjectItemUpdate(ctx, d, meta)
	}
	return resourceGithubProjectItemRead(ctx, d, meta)
}

func resourceGithubProjectItemRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var query struct {
		Node struct {
			Item projectV2ItemNode `graphql:"... on ProjectV2Item"`
		} `graphql:"node(id: $id)"`
	}
	err := meta.(*Owner).v4client.Query(ctx, &query, map[string]any{"id": githubv4.ID(d.Id())})
	if isProjectV2NotFound(err) {
		tflog.Info(ctx, "Removing project item from state because it no longer exists in GitHub", map[string]any{"id": d.Id()})
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	contentID := query.Node.Item.Content.Issue.ID
	if contentID == "" {
		contentID = query.Node.Item.Content.PullRequest.ID
	}
	if contentID == "" {
		contentID = query.Node.Item.Content.DraftIssue.ID
	}
	if contentID == "" {
		return diag.Errorf("project item %q has no supported content", d.Id())
	}
	for key, value := range map[string]any{
		"project_id": query.Node.Item.Project.ID,
		"content_id": contentID,
		"archived":   query.Node.Item.IsArchived,
	} {
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(fmt.Errorf("setting %s: %w", key, err))
		}
	}
	return nil
}

func resourceGithubProjectItemUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v4client
	if d.Get("archived").(bool) {
		var mutation struct {
			ArchiveProjectV2Item struct {
				Item projectV2ItemNode
			} `graphql:"archiveProjectV2Item(input: $input)"`
		}
		input := githubv4.ArchiveProjectV2ItemInput{ProjectID: githubv4.ID(d.Get("project_id").(string)), ItemID: githubv4.ID(d.Id())}
		if err := client.Mutate(ctx, &mutation, input, nil); err != nil {
			return diag.FromErr(err)
		}
	} else {
		var mutation struct {
			UnarchiveProjectV2Item struct {
				Item projectV2ItemNode
			} `graphql:"unarchiveProjectV2Item(input: $input)"`
		}
		input := githubv4.UnarchiveProjectV2ItemInput{ProjectID: githubv4.ID(d.Get("project_id").(string)), ItemID: githubv4.ID(d.Id())}
		if err := client.Mutate(ctx, &mutation, input, nil); err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceGithubProjectItemRead(ctx, d, meta)
}

func resourceGithubProjectItemDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var mutation struct {
		DeleteProjectV2Item struct {
			DeletedItemID githubv4.String
		} `graphql:"deleteProjectV2Item(input: $input)"`
	}
	input := githubv4.DeleteProjectV2ItemInput{ProjectID: githubv4.ID(d.Get("project_id").(string)), ItemID: githubv4.ID(d.Id())}
	err := meta.(*Owner).v4client.Mutate(ctx, &mutation, input, nil)
	if err != nil && !isProjectV2NotFound(err) {
		return diag.FromErr(err)
	}
	return nil
}
