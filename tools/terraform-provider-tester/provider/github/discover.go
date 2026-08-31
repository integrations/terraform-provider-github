package github

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/github/terraform-provider-tester/provider"
	gogithub "github.com/google/go-github/v88/github"
)

// discoverWithClient is the testable core of Discover.
// apiBase must end with "/" (or be empty for GitHub.com).
// token is used for the raw GraphQL POST authorization header.

const maxDiscoverPages = 100

func discoverWithClient(
	ctx context.Context,
	c *gogithub.Client,
	_ credInfo,
	apiBase, token string,
	opts provider.DiscoverOpts,
) (provider.DiscoveryResult, error) {
	var res provider.DiscoveryResult

	// List orgs via REST GET /user/orgs with pagination.
	listOpts := &gogithub.ListOptions{PerPage: 100}
	var lastOrgsNextPage int
	for page := 0; page < maxDiscoverPages; page++ {
		orgs, resp, err := c.Organizations.List(ctx, "", listOpts)
		if err != nil {
			res.Notes = append(res.Notes, "orgs unverified: "+err.Error())
			lastOrgsNextPage = 0
			break
		}
		for _, o := range orgs {
			res.Orgs = append(res.Orgs, o.GetLogin())
		}
		lastOrgsNextPage = resp.NextPage
		if resp.NextPage == 0 {
			break
		}
		listOpts.Page = resp.NextPage
	}
	if lastOrgsNextPage != 0 {
		res.Notes = append(res.Notes, "orgs may be incomplete: reached page limit")
	}

	// Enumerate enterprises via raw GraphQL POST.
	gqlBase := apiBase
	if gqlBase == "" {
		gqlBase = "https://api.github.com/"
	}
	gqlURL := strings.TrimSuffix(gqlBase, "/") + "/graphql"
	body, _ := json.Marshal(map[string]string{
		"query": "query{viewer{enterprises(first:100){nodes{slug name}}}}",
	})
	req, reqErr := http.NewRequestWithContext(ctx, http.MethodPost, gqlURL, bytes.NewReader(body))
	if reqErr != nil {
		res.Notes = append(res.Notes, "enterprises unverified: "+reqErr.Error())
	} else {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "bearer "+token)
		httpResp, httpErr := (&http.Client{Timeout: 30 * time.Second}).Do(req)
		if httpErr != nil {
			res.Notes = append(res.Notes, "enterprises unverified: "+httpErr.Error())
		} else {
			defer httpResp.Body.Close() //nolint:errcheck
			if httpResp.StatusCode != http.StatusOK {
				res.Notes = append(res.Notes, fmt.Sprintf("enterprises unverified: HTTP %d", httpResp.StatusCode))
			} else {
				raw, readErr := io.ReadAll(httpResp.Body)
				if readErr != nil {
					res.Notes = append(res.Notes, "enterprises unverified: "+readErr.Error())
				} else {
					var gqlResp struct {
						Data struct {
							Viewer struct {
								Enterprises struct {
									Nodes []struct {
										Slug string `json:"slug"`
										Name string `json:"name"`
									} `json:"nodes"`
								} `json:"enterprises"`
							} `json:"viewer"`
						} `json:"data"`
						Errors []struct {
							Message string `json:"message"`
						} `json:"errors"`
					}
					if jsonErr := json.Unmarshal(raw, &gqlResp); jsonErr != nil {
						res.Notes = append(res.Notes, "enterprises unverified: "+jsonErr.Error())
					} else {
						if len(gqlResp.Errors) > 0 {
							res.Notes = append(res.Notes, "enterprises unverified: "+gqlResp.Errors[0].Message)
						}
						for _, n := range gqlResp.Data.Viewer.Enterprises.Nodes {
							res.Enterprises = append(res.Enterprises, provider.DiscoveredEnterprise{
								Slug: n.Slug,
								Name: n.Name,
							})
						}
					}
				}
			}
		}
	}

	// List template repos for the given org (only when Org is specified).
	if opts.Org != "" {
		repoOpts := &gogithub.RepositoryListByOrgOptions{
			ListOptions: gogithub.ListOptions{PerPage: 100},
		}
		var lastReposNextPage int
		for page := 0; page < maxDiscoverPages; page++ {
			repos, resp, err := c.Repositories.ListByOrg(ctx, opts.Org, repoOpts)
			if err != nil {
				res.Notes = append(res.Notes, "template repos unverified: "+err.Error())
				lastReposNextPage = 0
				break
			}
			for _, r := range repos {
				if r.GetIsTemplate() {
					res.TemplateRepos = append(res.TemplateRepos, r.GetName())
				}
			}
			lastReposNextPage = resp.NextPage
			if resp.NextPage == 0 {
				break
			}
			repoOpts.Page = resp.NextPage
		}
		if lastReposNextPage != 0 {
			res.Notes = append(res.Notes, "template repos may be incomplete: reached page limit")
		}
	}

	return res, nil
}
