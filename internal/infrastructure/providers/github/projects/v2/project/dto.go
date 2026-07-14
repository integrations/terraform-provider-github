package project

import "github.com/shurcooL/githubv4"

type node struct {
	ID               githubv4.String
	Number           githubv4.Int
	Title            githubv4.String
	ShortDescription githubv4.String
	Readme           githubv4.String
	Public           githubv4.Boolean
	Closed           githubv4.Boolean
	URL              githubv4.URI
	Owner            struct {
		Typename     githubv4.String                 `graphql:"__typename"`
		Organization struct{ Login githubv4.String } `graphql:"... on Organization"`
		User         struct{ Login githubv4.String } `graphql:"... on User"`
	}
}
