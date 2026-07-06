package github

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestProvider(t *testing.T) {
	t.Parallel()

	t.Run("validate", func(t *testing.T) {
		t.Parallel()

		if err := NewProvider("test", "none")().InternalValidate(); err != nil {
			t.Fatalf("err: %s", err)
		}
	})
}

func Test_configureProviderMeta(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name        string
		installResp *string
		userResp    *string
		orgResp     *string
		conf        *Config
		wantName    string
		wantIsOrg   bool
		wantOrgId   int64
		wantErr     string
	}{
		{
			name: "anonymous",
			conf: &Config{},
		},
		{
			name:        "app_auth_organization",
			installResp: new(`{"id": 999999}`),
			orgResp:     new(`{"id": 123456}`),
			conf: &Config{
				AppID:             new("111111"),
				AppInstallationID: new("999999"),
				AppPEM:            mustNewPEM(t),
				Owner:             "test-org",
			},
			wantName:  "test-org",
			wantIsOrg: true,
			wantOrgId: 123456,
		},
		{
			name:        "app_auth_user",
			installResp: new(`{"id": 999999}`),
			conf: &Config{
				AppID:             new("111111"),
				AppInstallationID: new("999999"),
				AppPEM:            mustNewPEM(t),
				Owner:             "test-user",
			},
			wantName: "test-user",
		},
		{
			name:    "token_auth_organization",
			orgResp: new(`{"id": 123456}`),
			conf: &Config{
				Owner: "test-org",
				Token: "test-token",
			},
			wantName:  "test-org",
			wantIsOrg: true,
			wantOrgId: 123456,
		},
		{
			name: "token_auth_user",
			conf: &Config{
				Owner: "test-user",
				Token: "test-token",
			},
			wantName: "test-user",
		},
		{
			name: "errors_on_missing_owner",
			conf: &Config{
				Token: "test-token",
			},
			wantErr: "owner must be set when authenticating using the new client implementation",
		},
		{
			name: "legacy_client_anonymous",
			conf: &Config{
				LegacyClient: true,
			},
		},
		{
			name:        "legacy_client_app_auth_organization",
			installResp: new(`{"id": 999999}`),
			orgResp:     new(`{"id": 123456}`),
			conf: &Config{
				LegacyClient:      true,
				AppID:             new("111111"),
				AppInstallationID: new("999999"),
				AppPEM:            mustNewPEM(t),
				Owner:             "test-org",
			},
			wantName:  "test-org",
			wantIsOrg: true,
			wantOrgId: 123456,
		},
		{
			name:        "legacy_client_app_auth_user",
			installResp: new(`{"id": 999999}`),
			conf: &Config{
				LegacyClient:      true,
				AppID:             new("111111"),
				AppInstallationID: new("999999"),
				AppPEM:            mustNewPEM(t),
				Owner:             "test-user",
			},
			wantName: "test-user",
		},
		{
			name: "legacy_client_token_auth_user",
			conf: &Config{
				LegacyClient: true,
				Owner:        "test-user",
				Token:        "test-token",
			},
			wantName: "test-user",
		},
		{
			name:     "legacy_client_token_auth_no_owner",
			userResp: new(`{"login": "test-user"}`),
			conf: &Config{
				LegacyClient: true,
				Token:        "test-token",
			},
			wantName: "test-user",
		},
		{
			name: "legacy_client_token_auth_no_owner_found",
			conf: &Config{
				LegacyClient: true,
				Token:        "test-token",
			},
			wantErr: "owner cannot be found by token",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if regexp.MustCompile(`/access_tokens$`).MatchString(r.URL.Path) {
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(`{"token": "test-token", "expires_at": "2024-12-31T23:59:59Z"}`))
					return
				}

				if regexp.MustCompile(`/installation$`).MatchString(r.URL.Path) {
					if tt.installResp == nil {
						w.WriteHeader(http.StatusNotFound)
						return
					}

					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(*tt.installResp))
					return
				}

				if regexp.MustCompile(`/user$`).MatchString(r.URL.Path) {
					if tt.userResp == nil {
						w.WriteHeader(http.StatusNotFound)
						return
					}

					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(*tt.userResp))
					return
				}

				if regexp.MustCompile(`/orgs/[^/]+$`).MatchString(r.URL.Path) {
					if tt.orgResp == nil {
						w.WriteHeader(http.StatusNotFound)
						return
					}

					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(*tt.orgResp))
					return
				}

				w.WriteHeader(http.StatusNotFound)
			}))
			t.Cleanup(ts.Close)

			tt.conf.BaseURL = mustNewURL(t, ts.URL)

			meta, err := configureProviderMeta(t.Context(), "test", tt.conf)
			if err != nil {
				if tt.wantErr == "" {
					t.Fatalf("unexpected error: %v", err)
				}

				if !regexp.MustCompile(regexp.QuoteMeta(tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("expected error to match %q, got %v", tt.wantErr, err)
				}

				return
			}

			if tt.wantErr != "" {
				t.Fatalf("expected error %q, got nil", tt.wantErr)
			}

			if meta.name != tt.wantName {
				t.Errorf("expected owner name to be %q, got %q", tt.wantName, meta.name)
			}

			if meta.IsOrganization != tt.wantIsOrg {
				t.Errorf("expected IsOrganization to be %v, got %v", tt.wantIsOrg, meta.IsOrganization)
			}

			if meta.id != tt.wantOrgId {
				t.Errorf("expected owner id to be %d, got %d", tt.wantOrgId, meta.id)
			}

			if meta.v3client == nil {
				t.Errorf("expected rest client to be non-nil")
			}

			if tt.conf.Owner != "" && meta.v4client == nil {
				t.Errorf("expected graphql client to be non-nil")
			}
		})
	}
}

func Test_ghCLIHostFromAPIHost(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			got := ghCLIHostFromAPIHost(tc.host)
			if got != tc.expectedHost {
				t.Errorf("ghCLIHostFromAPIHost(%q) = %q, want %q", tc.host, got, tc.expectedHost)
			}
		})
	}
}
