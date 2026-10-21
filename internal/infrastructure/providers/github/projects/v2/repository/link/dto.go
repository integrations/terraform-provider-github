package link

import "github.com/shurcooL/githubv4"

type repositoryNode struct {
	ID         githubv4.String
	DatabaseID githubv4.Int
	Name       githubv4.String
	Owner      struct{ Login githubv4.String }
}

type pageInfo struct {
	EndCursor   githubv4.String
	HasNextPage githubv4.Boolean
}
