package github

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubCodeowners() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			// Input
			REPOSITORY_ID: {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			CODEOWNERS_EXISTS: {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},

		Read: resourceGithubAppInitCodeownersRead,
	}
}

func resourceGithubAppInitCodeownersRead(d *schema.ResourceData, meta interface{}) error {
	var query struct {
		Node struct {
			Repository struct {
				Root struct {
					GitObject struct {
						ID githubv4.ID
					} `graphql:"... on Blob"`
				} `graphql:"root:object(expression: $rootExpression)"`
				Github struct {
					GitObject struct {
						ID githubv4.ID
					} `graphql:"... on Blob"`
				} `graphql:"github:object(expression: $githubExpression)"`
				Docs struct {
					GitObject struct {
						ID githubv4.ID
					} `graphql:"... on Blob"`
				} `graphql:"docs:object(expression: $docsExpression)"`
			} `graphql:"... on Repository"`
		} `graphql:"node(id: $id)"`
	}
	variables := map[string]interface{}{
		"id":               githubv4.ID(d.Get(REPOSITORY_ID).(string)),
		"rootExpression":   githubv4.String("master:CODEOWNERS"),
		"githubExpression": githubv4.String("master:.github/CODEOWNERS"),
		"docsExpression":   githubv4.String("master:docs/CODEOWNERS"),
	}

	ctx := context.Background()
	client := meta.(*Organization).v4client
	err := client.Query(ctx, &query, variables)
	if err != nil {
		return err
	}

	root := query.Node.Repository.Root.GitObject.ID
	github := query.Node.Repository.Github.GitObject.ID
	docs := query.Node.Repository.Docs.GitObject.ID
	if root != nil || github != nil || docs != nil {
		err = d.Set(CODEOWNERS_EXISTS, true)
		if err != nil {
			return err
		}
	} else {
		err = d.Set(CODEOWNERS_EXISTS, false)
		if err != nil {
			return err
		}
	}

	d.SetId(fmt.Sprintf("%s/codeowners", d.Get(REPOSITORY_ID).(string)))

	return nil
}
