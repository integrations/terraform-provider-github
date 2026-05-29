package github

import (
	"testing"

	"github.com/shurcooL/githubv4"
)

func TestProvider(t *testing.T) {
	t.Run("validate", func(t *testing.T) {
		if err := NewProvider()().InternalValidate(); err != nil {
			t.Fatalf("err: %s", err)
		}
	})
}

func Test_configureProviderMeta(t *testing.T) {
	baseURL, _, err := getBaseURL(DotComAPIURL)
	if err != nil {
		t.Fatalf("failed to parse test base URL: %s", err.Error())
	}

	for _, tt := range []struct {
		name         string
		legacyClient bool
	}{
		{
			name:         "client",
			legacyClient: false,
		},
		{
			name:         "legacy_client",
			legacyClient: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("anonymous", func(t *testing.T) {
				config := &Config{
					BaseURL:      baseURL,
					LegacyClient: tt.legacyClient,
				}

				meta, err := configureProviderMeta(t.Context(), config)
				if err != nil {
					t.Fatalf("failed to return meta without error: %s", err.Error())
				}

				t.Run("rest_client", func(t *testing.T) {
					_, _, err = meta.v3client.Meta.Get(t.Context())
					if err != nil {
						t.Fatalf("failed to validate returned client without error: %s", err.Error())
					}
				})
			})

			t.Run("authenticated", func(t *testing.T) {
				skipUnauthenticated(t)

				config := &Config{
					GraphQLAPIPath: "graphql",
					BaseURL:        baseURL,
					Owner:          testAccConf.owner,
					Token:          testAccConf.token,
					LegacyClient:   tt.legacyClient,
				}

				meta, err := configureProviderMeta(t.Context(), config)
				if err != nil {
					t.Fatalf("failed to return meta without error: %s", err.Error())
				}

				t.Run("rest_client", func(t *testing.T) {
					if _, _, err = meta.v3client.Meta.Get(t.Context()); err != nil {
						t.Fatalf("failed to validate returned client without error: %s", err.Error())
					}
				})

				t.Run("graphql_client", func(t *testing.T) {
					client := meta.v4client
					var query struct {
						Meta struct {
							GitHubServicesSha githubv4.String
						}
					}
					if err := client.Query(t.Context(), &query, nil); err != nil {
						t.Fatalf("failed to validate returned client without error: %s", err.Error())
					}
				})
			})
		})
	}
}

func Test_ghCLIHostFromAPIHost(t *testing.T) {
	testCases := []struct {
		name         string
		host         string
		expectedHost string
	}{
		{
			name:         "dotcom API host is mapped to dotcom host",
			host:         "api.github.com",
			expectedHost: "github.com",
		},
		{
			name:         "ghec API host has api. prefix stripped",
			host:         "api.my-enterprise.ghe.com",
			expectedHost: "my-enterprise.ghe.com",
		},
		{
			name:         "ghec API host with numbers has api. prefix stripped",
			host:         "api.customer-123.ghe.com",
			expectedHost: "customer-123.ghe.com",
		},
		{
			name:         "ghes host is passed through unchanged",
			host:         "github.example.com",
			expectedHost: "github.example.com",
		},
		{
			name:         "ghes host with port is passed through unchanged",
			host:         "github.example.com:8443",
			expectedHost: "github.example.com:8443",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ghCLIHostFromAPIHost(tc.host)
			if got != tc.expectedHost {
				t.Errorf("ghCLIHostFromAPIHost(%q) = %q, want %q", tc.host, got, tc.expectedHost)
			}
		})
	}
}
