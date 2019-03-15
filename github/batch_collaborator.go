package github

import (
	"context"
	"strings"
	"time"

	"github.com/google/go-github/github"
	batcher "github.com/paultyng/go-batcher"
)

type readCollaboratorParam struct {
	owner string
	repo  string
	name  string
}

func batchReadCollaborator(after time.Duration, repositories repositoriesService) func(ctx context.Context, owner, repo, name string) (*github.User, error) {
	b := batcher.New(after, func(params []interface{}) ([]interface{}, error) {
		// the size of results must match params
		results := make([]interface{}, len(params))

		type key struct {
			owner string
			repo  string
		}

		repos := map[key]bool{}
		for _, pi := range params {
			p := pi.(*readCollaboratorParam)
			k := key{strings.ToLower(p.owner), strings.ToLower(p.repo)}
			repos[k] = true
		}

		for k := range repos {
			err := listAllPages(func(opt github.ListOptions) (*github.Response, error) {
				page, resp, err := repositories.ListCollaborators(context.Background(), k.owner, k.repo, &github.ListCollaboratorsOptions{
					ListOptions: opt,
				})
				if err != nil {
					return nil, err
				}

				for _, h := range page {
					// match parameters with results in each page and set the result
					for i, pi := range params {
						p := pi.(*readCollaboratorParam)
						if k.owner == strings.ToLower(p.owner) &&
							k.repo == strings.ToLower(p.repo) &&
							strings.ToLower(h.GetName()) == p.name {
							results[i] = h
						}
					}
				}
				return resp, nil
			})
			if err != nil {
				return nil, err
			}
		}
		return results, nil
	})

	return func(ctx context.Context, owner, repo, name string) (*github.User, error) {
		result, err := b.Get(ctx, &readCollaboratorParam{
			owner: owner,
			repo:  repo,
			name:  name,
		})
		if err != nil {
			return nil, err
		}
		if result == nil {
			return nil, nil
		}
		r := result.(*github.User)
		return r, nil
	}
}
