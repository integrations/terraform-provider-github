package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/shurcooL/githubv4"
)

const (
	REPOSITORY_COLLABORATORS = "collaborators"
	REPOSITORY_ID            = "repository_id"
)

const (
	USER_EMAIL         = "email"
	USER_IS_SITE_ADMIN = "is_site_admin"
	USER_LOGIN         = "login"
	USER_NAME          = "name"
	USER_PERMISSION    = "permission"
)

type User struct {
	Email       githubv4.String
	ID          githubv4.ID
	IsSiteAdmin githubv4.Boolean
	Login       githubv4.String
	Name        githubv4.String
}

type PageInfo struct {
	EndCursor   githubv4.String
	HasNextPage bool
}

func dataSourceGithubCollaborators() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			// Input
			REPOSITORY_ID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
			// Computed
			REPOSITORY_COLLABORATORS: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						USER_LOGIN: {
							Type:     schema.TypeString,
							Computed: true,
						},
						USER_EMAIL: {
							Type:     schema.TypeString,
							Computed: true,
						},
						USER_IS_SITE_ADMIN: {
							Type:     schema.TypeBool,
							Computed: true,
						},
						USER_NAME: {
							Type:     schema.TypeString,
							Computed: true,
						},
						USER_PERMISSION: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},

		Read: dataSourceGithubCollaboratorsRead,
	}
}

func dataSourceGithubCollaboratorsRead(d *schema.ResourceData, meta interface{}) error {
	var query struct {
		Node struct {
			Repository struct {
				Collaborators struct {
					Edges []struct {
						Node       User
						Permission githubv4.RepositoryPermission
					}
					PageInfo PageInfo
				} `graphql:"collaborators(first: $first, after: $cursor)"`
				ID githubv4.ID
			} `graphql:"... on Repository"`
		} `graphql:"node(id: $id)"`
	}
	variables := map[string]interface{}{
		"id":     githubv4.ID(d.Get(REPOSITORY_ID).(string)),
		"first":  githubv4.Int(100),
		"cursor": (*githubv4.String)(nil),
	}

	ctx := context.Background()
	client := meta.(*Organization).v4client

	var allEdges []struct {
		Node       User
		Permission githubv4.RepositoryPermission
	}
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return err
		}

		allEdges = append(allEdges, query.Node.Repository.Collaborators.Edges...)

		if !query.Node.Repository.Collaborators.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Node.Repository.Collaborators.PageInfo.EndCursor)
	}

	var allUsers []map[string]interface{}
	for _, u := range allEdges {
		user := make(map[string]interface{})
		user[USER_EMAIL] = string(u.Node.Email)
		user[USER_IS_SITE_ADMIN] = bool(u.Node.IsSiteAdmin)
		user[USER_LOGIN] = string(u.Node.Login)
		user[USER_NAME] = string(u.Node.Name)
		user[USER_PERMISSION] = string(u.Permission)
		allUsers = append(allUsers, user)
	}

	err := d.Set(REPOSITORY_COLLABORATORS, allUsers)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/collaborators", query.Node.Repository.ID))

	return nil
}
