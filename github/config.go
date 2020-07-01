package github

import (
	"context"
	"net/http"
	"net/url"
	"path"

	"github.com/google/go-github/v31/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Config struct {
	Token   string
	Owner   string
	BaseURL string
}

type Owner struct {
	name           string
	id             int64
	v3client       *github.Client
	v4client       *githubv4.Client
	StopContext    context.Context
	IsOrganization bool
}

// Clients configures and returns a fully initialized GithubClient and Githubv4Client
func (c *Config) Clients() (interface{}, error) {
	var owner Owner
	var ts oauth2.TokenSource
	var tc *http.Client

	ctx := context.Background()

	ts = oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.Token},
	)
	tc = oauth2.NewClient(ctx, ts)
	tc.Transport = NewEtagTransport(tc.Transport)
	tc.Transport = NewRateLimitTransport(tc.Transport)
	tc.Transport = logging.NewTransport("Github", tc.Transport)

	// Create GraphQL Client
	uv4, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}
	uv4.Path = path.Join(uv4.Path, "graphql")
	v4client := githubv4.NewEnterpriseClient(uv4.String(), tc)

	// Create Rest Client
	uv3, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}
	if uv3.String() != "https://api.github.com/" {
		uv3.Path = uv3.Path + "v3/"
	}
	v3client, err := github.NewEnterpriseClient(uv3.String(), "", tc)
	if err != nil {
		return nil, err
	}
	v3client.BaseURL = uv3

	owner.v3client = v3client
	owner.v4client = v4client

	owner.name = c.Owner
	if owner.name != "" {
		remoteOrg, _, err := owner.v3client.Organizations.Get(ctx, owner.name)
		if err == nil {
			if remoteOrg != nil {
				owner.id = remoteOrg.GetID()
				owner.IsOrganization = true
			}
		}
	}
	return &owner, nil
}
