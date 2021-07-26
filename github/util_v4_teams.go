package github

import "github.com/shurcooL/githubv4"

type TeamsQuery struct {
	Organization struct {
		ID    githubv4.String
		Teams struct {
			Nodes []struct {
				ID          githubv4.String
				DatabaseID  githubv4.Int
				Slug        githubv4.String
				Name        githubv4.String
				Description githubv4.String
				Privacy     githubv4.String
				Members     struct {
					Nodes []struct {
						Login githubv4.String
					}
				}
				Repositories struct {
					Nodes []struct {
						Name githubv4.String
					}
				}
			}
			PageInfo PageInfo
		} `graphql:"teams(first:$first, after:$cursor, rootTeamsOnly:$rootTeamsOnly)"`
	} `graphql:"organization(login:$login)"`
}
