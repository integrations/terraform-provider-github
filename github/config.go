package github

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-github/v31/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"
	"golang.org/x/oauth2"
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
	name        string
	id          int64
	client      *github.Client
	StopContext context.Context
}

// Client configures and returns a fully initialized GithubClient
func (c *Config) Client() (interface{}, error) {
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

	org.client = github.NewClient(tc)

	if c.BaseURL != "" {
		u, err := url.Parse(c.BaseURL)
		if err != nil {
			return nil, err
		}
		org.client.BaseURL = u
	}

	if c.Individual {
		org.name = ""
	} else {
		org.name = c.Organization

		remoteOrg, _, err := org.client.Organizations.Get(ctx, org.name)
		if err != nil {
			return nil, err
		}
		org.id = remoteOrg.GetID()
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
