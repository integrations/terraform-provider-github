package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

type projectV2RepositoryNode struct {
	ID            githubv4.String
	Name          githubv4.String
	NameWithOwner githubv4.String
	Owner         struct {
		Login githubv4.String
	}
}

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
	repository, err := queryProjectV2RepositoryByName(ctx, meta.(*Owner).v4client, projectV2RepositoryOwner(d, meta), d.Get("repository").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	var mutation struct {
		LinkProjectV2ToRepository struct {
			Repository projectV2RepositoryNode
		} `graphql:"linkProjectV2ToRepository(input: $input)"`
	}
	input := githubv4.LinkProjectV2ToRepositoryInput{ProjectID: githubv4.ID(d.Get("project_id").(string)), RepositoryID: githubv4.ID(repository.ID)}
	if err := meta.(*Owner).v4client.Mutate(ctx, &mutation, input, nil); err != nil {
		return diag.FromErr(err)
	}
	id, err := buildID(d.Get("project_id").(string), string(repository.ID))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)
	return resourceGithubProjectRepositoryRead(ctx, d, meta)
}

func resourceGithubProjectRepositoryRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	projectID, repositoryID, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	client := meta.(*Owner).v4client
	linked, err := projectV2HasRepository(ctx, client, projectID, repositoryID)
	if isProjectV2NotFound(err) || (err == nil && !linked) {
		tflog.Info(ctx, "Removing project repository link from state because it no longer exists in GitHub", map[string]any{"id": d.Id()})
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	repository, err := queryProjectV2RepositoryByID(ctx, client, repositoryID)
	if err != nil {
		return diag.FromErr(err)
	}
	for key, value := range map[string]any{"project_id": projectID, "repository_owner": repository.Owner.Login, "repository": repository.Name} {
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(fmt.Errorf("setting %s: %w", key, err))
		}
	}
	return nil
}

func resourceGithubProjectRepositoryDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	projectID, repositoryID, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	var mutation struct {
		UnlinkProjectV2FromRepository struct {
			Repository projectV2RepositoryNode
		} `graphql:"unlinkProjectV2FromRepository(input: $input)"`
	}
	input := githubv4.UnlinkProjectV2FromRepositoryInput{ProjectID: githubv4.ID(projectID), RepositoryID: githubv4.ID(repositoryID)}
	err = meta.(*Owner).v4client.Mutate(ctx, &mutation, input, nil)
	if err != nil && !isProjectV2NotFound(err) {
		return diag.FromErr(err)
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
	if owner := d.Get("repository_owner").(string); owner != "" {
		return owner
	}
	return meta.(*Owner).name
}

func queryProjectV2RepositoryByName(ctx context.Context, client *githubv4.Client, owner, name string) (projectV2RepositoryNode, error) {
	var query struct {
		Repository projectV2RepositoryNode `graphql:"repository(owner: $owner, name: $name)"`
	}
	err := client.Query(ctx, &query, map[string]any{"owner": githubv4.String(owner), "name": githubv4.String(name)})
	return query.Repository, err
}

func queryProjectV2RepositoryByID(ctx context.Context, client *githubv4.Client, id string) (projectV2RepositoryNode, error) {
	var query struct {
		Node struct {
			Repository projectV2RepositoryNode `graphql:"... on Repository"`
		} `graphql:"node(id: $id)"`
	}
	err := client.Query(ctx, &query, map[string]any{"id": githubv4.ID(id)})
	return query.Node.Repository, err
}

func projectV2HasRepository(ctx context.Context, client *githubv4.Client, projectID, repositoryID string) (bool, error) {
	var after *githubv4.String
	for {
		var query struct {
			Node struct {
				Project struct {
					Repositories struct {
						Nodes    []struct{ ID githubv4.String }
						PageInfo PageInfo
					} `graphql:"repositories(first: 100, after: $after)"`
				} `graphql:"... on ProjectV2"`
			} `graphql:"node(id: $id)"`
		}
		err := client.Query(ctx, &query, map[string]any{"id": githubv4.ID(projectID), "after": after})
		if err != nil {
			return false, err
		}
		for _, repository := range query.Node.Project.Repositories.Nodes {
			if string(repository.ID) == repositoryID {
				return true, nil
			}
		}
		if !bool(query.Node.Project.Repositories.PageInfo.HasNextPage) {
			return false, nil
		}
		cursor := query.Node.Project.Repositories.PageInfo.EndCursor
		after = &cursor
	}
}
