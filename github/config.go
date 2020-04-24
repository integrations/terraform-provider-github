package github

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/google/go-github/v29/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type Config struct {
	Token        string
	Organization string
	BaseURL      string
	Insecure     bool
	Individual   bool
	Anonymous    bool
}

type Organization struct {
	name         string
	id           int64
	v3client     *github.Client
	v4client     *githubv4.Client
	StopContext  context.Context
	isEnterprise bool
}

// Clients configures and returns a fully initialized GithubClient and Githubv4Client
func (c *Config) Clients() (interface{}, error) {
	var org Organization
	var ts oauth2.TokenSource
	var tc *http.Client

	ctx := context.Background()

	if c.Insecure {
		insecureClient := insecureHttpClient()
		ctx = context.WithValue(ctx, oauth2.HTTPClient, insecureClient)
	}

	// Either Organization needs to be set, or Individual needs to be true
	if c.Organization != "" && c.Individual {
		return nil, fmt.Errorf("If `individual` is true, `organization` cannot be set.")
	}
	if c.Organization == "" && !c.Individual {
		return nil, fmt.Errorf("If `individual` is false, `organization` is required.")
	}

	// Either run as anonymous, or run with a Token
	if c.Token != "" && c.Anonymous {
		return nil, fmt.Errorf("If `anonymous` is true, `token` cannot be set.")
	}
	if c.Token == "" && !c.Anonymous {
		return nil, fmt.Errorf("If `anonymous` is false, `token` is required.")
	}

	if !c.Anonymous {
		ts = oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: c.Token},
		)
	}

	tc = oauth2.NewClient(ctx, ts)

	if c.Anonymous {
		tc.Transport = http.DefaultTransport
	} else {
		tc.Transport = NewEtagTransport(tc.Transport)
	}

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

	org.v3client = v3client
	org.v4client = v4client

	if c.Individual {
		org.name = ""
	} else {
		org.name = c.Organization

		remoteOrg, _, err := org.v3client.Organizations.Get(ctx, org.name)
		if err != nil {
			return nil, err
		}
		org.id = remoteOrg.GetID()
		org.isEnterprise = strings.EqualFold(remoteOrg.GetPlan().GetName(), "enterprise")
	}

	return &org, nil
}

func insecureHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
