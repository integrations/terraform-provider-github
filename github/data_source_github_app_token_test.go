package github

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAccGithubAppTokenDataSource(t *testing.T) {
	t.Run("creates a application token without error", func(t *testing.T) {
		expectedAccessToken := "W+2e/zjiMTweDAr2b35toCF+h29l7NW92rJIPvFrCJQK"

		owner := "test-owner"

		pemData, err := os.ReadFile(testGitHubAppPrivateKeyFile)
		if err != nil {
			t.Logf("Unexpected error: %s", err)
			t.Fail()
		}

		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri: fmt.Sprintf("/app/installations/%s/access_tokens", testGitHubAppInstallationID),
				ExpectedHeaders: map[string]string{
					"Accept": "application/vnd.github.v3+json",
				},
				ResponseBody: fmt.Sprintf(`{"token": "%s"}`, expectedAccessToken),
				StatusCode:   201,
			},
		})
		defer ts.Close()

		httpCl := http.DefaultClient
		httpCl.Transport = http.DefaultTransport

		client := github.NewClient(httpCl)
		u, _ := url.Parse(ts.URL + "/")
		client.BaseURL = u

		meta := &Owner{
			name:     owner,
			v3client: client,
		}

		testSchema := map[string]*schema.Schema{
			"app_id":          {Type: schema.TypeString},
			"installation_id": {Type: schema.TypeString},
			"pem_file":        {Type: schema.TypeString},
			"repositories":    {Type: schema.TypeList, Elem: &schema.Schema{Type: schema.TypeString}},
			"token":           {Type: schema.TypeString},
		}

		schema := schema.TestResourceDataRaw(t, testSchema, map[string]any{
			"app_id":          testGitHubAppID,
			"installation_id": testGitHubAppInstallationID,
			"pem_file":        string(pemData),
			"token":           "",
		})

		err = dataSourceGithubAppTokenRead(schema, meta)
		if err != nil {
			t.Logf("Unexpected error: %s", err)
			t.Fail()
		}

		if schema.Get("token") != expectedAccessToken {
			t.Logf("Expected %s, got %s", expectedAccessToken, schema.Get("token"))
			t.Fail()
		}
	})

	t.Run("creates a application token scoped to repositories", func(t *testing.T) {
		expectedAccessToken := "ghs_scoped_token_12345"

		owner := "test-owner"

		pemData, err := os.ReadFile(testGitHubAppPrivateKeyFile)
		if err != nil {
			t.Logf("Unexpected error: %s", err)
			t.Fail()
		}

		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri: fmt.Sprintf("/app/installations/%s/access_tokens", testGitHubAppInstallationID),
				ExpectedHeaders: map[string]string{
					"Accept":       "application/vnd.github.v3+json",
					"Content-Type": "application/json",
				},
				ExpectedBody:  []byte(`{"repositories":["repo1","repo2"]}`),
				ResponseBody: fmt.Sprintf(`{"token": "%s"}`, expectedAccessToken),
				StatusCode:   201,
			},
		})
		defer ts.Close()

		httpCl := http.DefaultClient
		httpCl.Transport = http.DefaultTransport

		client := github.NewClient(httpCl)
		u, _ := url.Parse(ts.URL + "/")
		client.BaseURL = u

		meta := &Owner{
			name:     owner,
			v3client: client,
		}

		testSchema := map[string]*schema.Schema{
			"app_id":          {Type: schema.TypeString},
			"installation_id": {Type: schema.TypeString},
			"pem_file":        {Type: schema.TypeString},
			"repositories":    {Type: schema.TypeList, Elem: &schema.Schema{Type: schema.TypeString}},
			"token":           {Type: schema.TypeString},
		}

		schema := schema.TestResourceDataRaw(t, testSchema, map[string]any{
			"app_id":          testGitHubAppID,
			"installation_id": testGitHubAppInstallationID,
			"pem_file":        string(pemData),
			"repositories":    []any{"repo1", "repo2"},
			"token":           "",
		})

		err = dataSourceGithubAppTokenRead(schema, meta)
		if err != nil {
			t.Logf("Unexpected error: %s", err)
			t.Fail()
		}

		if schema.Get("token") != expectedAccessToken {
			t.Logf("Expected %s, got %s", expectedAccessToken, schema.Get("token"))
			t.Fail()
		}
	})
}
