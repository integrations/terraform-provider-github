package github

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

const (
	projectV2OwnerOrganization = "organization"
	projectV2OwnerUser         = "user"
)

type projectV2Node struct {
	ID               githubv4.String
	Number           githubv4.Int
	Title            githubv4.String
	ShortDescription githubv4.String
	Readme           githubv4.String
	Public           githubv4.Boolean
	Closed           githubv4.Boolean
	URL              githubv4.URI
	Owner            struct {
		Typename     githubv4.String `graphql:"__typename"`
		Organization struct {
			Login githubv4.String
		} `graphql:"... on Organization"`
		User struct {
			Login githubv4.String
		} `graphql:"... on User"`
	}
}

func projectV2OwnerID(ctx context.Context, client *githubv4.Client, ownerType, login string) (githubv4.ID, error) {
	variables := map[string]any{"login": githubv4.String(login)}

	switch ownerType {
	case projectV2OwnerOrganization:
		var query struct {
			Organization struct {
				ID githubv4.String
			} `graphql:"organization(login: $login)"`
		}
		if err := client.Query(ctx, &query, variables); err != nil {
			return nil, fmt.Errorf("querying organization %q: %w", login, err)
		}
		return githubv4.ID(query.Organization.ID), nil
	case projectV2OwnerUser:
		var query struct {
			User struct {
				ID githubv4.String
			} `graphql:"user(login: $login)"`
		}
		if err := client.Query(ctx, &query, variables); err != nil {
			return nil, fmt.Errorf("querying user %q: %w", login, err)
		}
		return githubv4.ID(query.User.ID), nil
	default:
		return nil, fmt.Errorf("unsupported Projects V2 owner type %q", ownerType)
	}
}

func queryProjectV2(ctx context.Context, client *githubv4.Client, id string) (projectV2Node, error) {
	var query struct {
		Node struct {
			Project projectV2Node `graphql:"... on ProjectV2"`
		} `graphql:"node(id: $id)"`
	}

	err := client.Query(ctx, &query, map[string]any{"id": githubv4.ID(id)})
	return query.Node.Project, err
}

func projectV2Owner(project projectV2Node) (string, string) {
	if project.Owner.Typename == "Organization" {
		return projectV2OwnerOrganization, string(project.Owner.Organization.Login)
	}
	return projectV2OwnerUser, string(project.Owner.User.Login)
}

func setProjectV2State(d *schema.ResourceData, project projectV2Node) error {
	ownerType, owner := projectV2Owner(project)
	values := map[string]any{
		"owner_type":        ownerType,
		"owner":             owner,
		"number":            int(project.Number),
		"title":             string(project.Title),
		"short_description": string(project.ShortDescription),
		"readme":            string(project.Readme),
		"public":            bool(project.Public),
		"closed":            bool(project.Closed),
		"url":               project.URL.String(),
	}

	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			return fmt.Errorf("setting %s: %w", key, err)
		}
	}
	return nil
}

func projectV2OwnerLogin(d *schema.ResourceData, meta any) string {
	if owner := d.Get("owner").(string); owner != "" {
		return owner
	}
	return meta.(*Owner).name
}

func isProjectV2NotFound(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Could not resolve to a node with the global id")
}

func validateProjectV2Date(value any, key string) ([]string, []error) {
	if _, err := time.Parse(time.DateOnly, value.(string)); err != nil {
		return nil, []error{fmt.Errorf("%q must use YYYY-MM-DD format: %w", key, err)}
	}
	return nil, nil
}
