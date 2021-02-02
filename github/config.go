package github

import (
	"context"
	"net/http"
	"net/url"
	"path"

	"github.com/google/go-github/v32/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Config struct {
	Token        string
	Owner        string
	Organization string
	BaseURL      string
	Insecure     bool
	Individual   bool
	Anonymous    bool
}

type Owner struct {
	name           string
	id             int64
	v3client       *github.Client
	v4client       *githubv4.Client
	StopContext    context.Context
	IsOrganization bool
}

func RateLimitedHTTPClient(client *http.Client) *http.Client {

	client.Transport = NewEtagTransport(client.Transport)
	client.Transport = NewRateLimitTransport(client.Transport)
	client.Transport = logging.NewTransport("Github", client.Transport)

	return client

}

func (c *Config) AuthenticatedHTTPClient() *http.Client {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.Token},
	)
	client := oauth2.NewClient(ctx, ts)

	return RateLimitedHTTPClient(client)

}

func (c *Config) AnonymousHTTPClient() *http.Client {
	client := &http.Client{Transport: &http.Transport{}}
	return RateLimitedHTTPClient(client)
}

func (c *Config) NewGraphQLClient(client *http.Client) (*githubv4.Client, error) {

	uv4, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	if uv4.String() != "https://api.github.com/" {
		uv4.Path = path.Join(uv4.Path, "api/graphql/")
	} else {
		uv4.Path = path.Join(uv4.Path, "graphql")
	}

	return githubv4.NewEnterpriseClient(uv4.String(), client), nil

}

func (c *Config) NewRESTClient(client *http.Client) (*github.Client, error) {

	uv3, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, err
	}

	if uv3.String() != "https://api.github.com/" {
		uv3.Path = uv3.Path + "api/v3/"
	}

	v3client, err := github.NewEnterpriseClient(uv3.String(), "", client)
	if err != nil {
		return nil, err
	}

	v3client.BaseURL = uv3

	return v3client, nil

}

func (c *Config) ConfigureOwner(owner *Owner) (*Owner, error) {

	ctx := context.Background()
	owner.name = c.Owner
	if owner.name == "" {
		// Discover authenticated user
		user, _, err := owner.v3client.Users.Get(ctx, "")
		if err != nil {
			return nil, err
		}
		owner.name = user.GetLogin()
	} else {
		remoteOrg, _, err := owner.v3client.Organizations.Get(ctx, owner.name)
		if err == nil {
			if remoteOrg != nil {
				owner.id = remoteOrg.GetID()
				owner.IsOrganization = true
			}
		}
	}

	return owner, nil
}

// Meta returns the meta parameter that is passed into subsequent resources
// https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/helper/schema#ConfigureFunc
func (c *Config) Meta() (interface{}, error) {

	var client *http.Client
	if c.Anonymous {
		client = c.AnonymousHTTPClient()
	} else {
		client = c.AuthenticatedHTTPClient()
	}

	v3client, err := c.NewRESTClient(client)
	if err != nil {
		return nil, err
	}

	v4client, err := c.NewGraphQLClient(client)
	if err != nil {
		return nil, err
	}

	var owner Owner
	owner.v4client = v4client
	owner.v3client = v3client

	if c.Anonymous {
		return &owner, nil
	} else {
		return c.ConfigureOwner(&owner)
	}

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
	if owner.name == "" {
		// Discover authenticated user
		user, _, err := owner.v3client.Users.Get(ctx, "")
		if err != nil {
			return nil, err
		}
		owner.name = user.GetLogin()
	} else {
		remoteOrg, _, err := owner.v3client.Organizations.Get(ctx, owner.name)
		if err == nil {
			if remoteOrg != nil {
				owner.id = remoteOrg.GetID()
				owner.IsOrganization = true
			}
		} else {
			return nil, err
		}
	}
	return &owner, nil
}
