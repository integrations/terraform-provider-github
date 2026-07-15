package link

import "github.com/shurcooL/githubv4"

type teamNode struct {
	ID           githubv4.String
	DatabaseID   githubv4.Int
	Slug         githubv4.String
	Organization struct{ Login githubv4.String }
}

type pageInfo struct {
	EndCursor   githubv4.String
	HasNextPage githubv4.Boolean
}
