package github

import "github.com/shurcooL/githubv4"

type TeamsQuery struct {
	Organization struct {
		Teams struct {
			Nodes []struct {
				ID          githubv4.String
				Slug        githubv4.String
				Name        githubv4.String
				Description githubv4.String
				Privacy     githubv4.String
				Members     struct {
					Nodes []struct {
						Login githubv4.String
					}
				}
			}
			PageInfo PageInfo
		} `graphql:"teams(first:$first, after:$cursor)"`
	} `graphql:"organization(login:$login)"`
}
