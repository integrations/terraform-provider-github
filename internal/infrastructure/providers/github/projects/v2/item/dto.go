package item

import "github.com/shurcooL/githubv4"

type node struct {
	ID         githubv4.String
	IsArchived githubv4.Boolean
	Project    struct{ ID githubv4.String }
	Content    struct {
		Issue       struct{ ID githubv4.String } `graphql:"... on Issue"`
		PullRequest struct{ ID githubv4.String } `graphql:"... on PullRequest"`
		DraftIssue  struct{ ID githubv4.String } `graphql:"... on DraftIssue"`
	}
}
