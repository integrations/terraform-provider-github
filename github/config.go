package github

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Config struct {
	Token            string
	Owner            string
	BaseURL          *url.URL
	IsGHES           bool
	Insecure         bool
	WriteDelay       time.Duration
	ReadDelay        time.Duration
	RetryDelay       time.Duration
	RetryableErrors  map[int]bool
	MaxRetries       int
	ParallelRequests bool
}

type Owner struct {
	name           string
	id             int64
	v3client       *github.Client
	v4client       *githubv4.Client
	StopContext    context.Context
	IsOrganization bool
}

const (
	// DotComAPIURL is the base API URL for github.com.
	DotComAPIURL = "https://api.github.com/"
	// DotComHost is the hostname for github.com.
	DotComHost = "github.com"
	// DotComAPIHost is the API hostname for github.com.
	DotComAPIHost = "api.github.com"
	// GHESRESTAPISuffix is the rest api suffix for GitHub Enterprise Server.
	GHESRESTAPIPath = "api/v3/"
)

var (
	// GHECHostMatch is a regex to match GitHub Enterprise Cloud hosts.
	GHECHostMatch = regexp.MustCompile(`\.ghe\.com$`)
	// GHECAPIHostMatch is a regex to match GitHub Enterprise Cloud API hosts.
	GHECAPIHostMatch = regexp.MustCompile(`^api\.[a-zA-Z0-9-]+\.ghe\.com$`)
)

func RateLimitedHTTPClient(client *http.Client, writeDelay, readDelay, retryDelay time.Duration, parallelRequests bool, retryableErrors map[int]bool, maxRetries int) *http.Client {
	client.Transport = NewEtagTransport(client.Transport)
	client.Transport = NewRateLimitTransport(client.Transport, WithWriteDelay(writeDelay), WithReadDelay(readDelay), WithParallelRequests(parallelRequests))
	client.Transport = logging.NewSubsystemLoggingHTTPTransport("GitHub", client.Transport)
	client.Transport = newPreviewHeaderInjectorTransport(map[string]string{
		// TODO: remove when Stone Crop preview is moved to general availability in the GraphQL API
		"Accept": "application/vnd.github.stone-crop-preview+json",
	}, client.Transport)

	if maxRetries > 0 {
		client.Transport = NewRetryTransport(client.Transport, WithRetryDelay(retryDelay), WithRetryableErrors(retryableErrors), WithMaxRetries(maxRetries))
	}

	return client
}

func (c *Config) AuthenticatedHTTPClient() *http.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.Token},
	)
	client := oauth2.NewClient(ctx, ts)

	return RateLimitedHTTPClient(client, c.WriteDelay, c.ReadDelay, c.RetryDelay, c.ParallelRequests, c.RetryableErrors, c.MaxRetries)
}

func (c *Config) Anonymous() bool {
	return c.Token == ""
}

func (c *Config) AnonymousHTTPClient() *http.Client {
	client := &http.Client{Transport: &http.Transport{}}
	return RateLimitedHTTPClient(client, c.WriteDelay, c.ReadDelay, c.RetryDelay, c.ParallelRequests, c.RetryableErrors, c.MaxRetries)
}

func (c *Config) NewGraphQLClient(client *http.Client) (*githubv4.Client, error) {
	var path string
	if c.IsGHES {
		path = "api/graphql"
	} else {
		path = "graphql"
	}

	return githubv4.NewEnterpriseClient(c.BaseURL.JoinPath(path).String(), client), nil
}

func (c *Config) NewRESTClient(client *http.Client) (*github.Client, error) {
	path := ""
	if c.IsGHES {
		path = GHESRESTAPIPath
	}

	v3client := github.NewClient(client)
	v3client.BaseURL = c.BaseURL.JoinPath(path)

	return v3client, nil
}

func (c *Config) ConfigureOwner(owner *Owner) (*Owner, error) {
	ctx := context.Background()
	owner.name = c.Owner
	if owner.name == "" {
		if c.Anonymous() {
			return owner, nil
		}
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
// https://godoc.org/github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema#ConfigureFunc
func (c *Config) Meta() (any, error) {
	var client *http.Client
	if c.Anonymous() {
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
	owner.StopContext = context.Background()

	_, err = c.ConfigureOwner(&owner)
	if err != nil {
		return &owner, err
	}
	return &owner, nil
}

type previewHeaderInjectorTransport struct {
	rt             http.RoundTripper
	previewHeaders map[string]string
}

func newPreviewHeaderInjectorTransport(headers map[string]string, rt http.RoundTripper) *previewHeaderInjectorTransport {
	return &previewHeaderInjectorTransport{
		rt:             rt,
		previewHeaders: headers,
	}
}

func (injector *previewHeaderInjectorTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for name, value := range injector.previewHeaders {
		header := req.Header.Get(name)
		if header == "" {
			header = value
		} else {
			header = strings.Join([]string{header, value}, ",")
		}
		req.Header.Set(name, header)
	}
	return injector.rt.RoundTrip(req)
}

// getBaseURL returns a correctly configured base URL and a bool as to if this is GitHub Enterprise Server.
func getBaseURL(s string) (*url.URL, bool, error) {
	if len(s) == 0 {
		s = DotComAPIURL
	}

	u, err := url.Parse(s)
	if err != nil {
		return nil, false, err
	}

	if !u.IsAbs() {
		return nil, false, fmt.Errorf("base url must be absolute")
	}

	u = u.JoinPath("/")

	switch {
	case u.Host == DotComAPIHost:
	case u.Host == DotComHost:
		u.Host = DotComAPIHost
	case GHECAPIHostMatch.MatchString(u.Host):
	case GHECHostMatch.MatchString(u.Host):
		u.Host = fmt.Sprintf("api.%s", u.Host)
	default:
		u.Path = strings.TrimSuffix(u.Path, GHESRESTAPIPath)
		return u, true, nil
	}

	if u.Scheme != "https" {
		return nil, false, fmt.Errorf("base url for github.com or ghe.com must use the https scheme")
	}

	if len(u.Path) > 1 {
		return nil, false, fmt.Errorf("base url for github.com or ghe.com must not contain a path, got %s", u.Path)
	}

	u.Path = "/"

	return u, false, nil
}
