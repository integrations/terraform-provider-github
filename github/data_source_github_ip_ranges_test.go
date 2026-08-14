package github

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestGithubIpRangesDataSourceRead(t *testing.T) {
	t.Parallel()

	t.Run("gives up on a stalled metadata request once the deadline expires", func(t *testing.T) {
		t.Parallel()

		serverDelay := 5 * time.Second
		readDeadline := 200 * time.Millisecond

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			select {
			case <-r.Context().Done():
			case <-time.After(serverDelay):
				mustWrite(w, `{}`)
			}
		}))
		defer ts.Close()

		// StopContext mirrors the provider configuration and never expires, so the
		// read has to use the context it is called with to respect the deadline.
		meta := &Owner{
			v3client:    mustCreateTestGitHubClient(t, ts.URL),
			StopContext: context.Background(),
		}

		ctx, cancel := context.WithTimeout(t.Context(), readDeadline)
		defer cancel()

		start := time.Now()
		diags := dataSourceGithubIpRangesRead(ctx, dataSourceGithubIpRanges().TestResourceData(), meta)
		elapsed := time.Since(start)

		if !diags.HasError() {
			t.Fatal("expected an error when the read deadline expires before the response arrives")
		}
		if elapsed >= serverDelay {
			t.Fatalf("read waited %s for the metadata response instead of honoring the %s deadline", elapsed, readDeadline)
		}
	})

	t.Run("populates IP ranges when the response arrives before the deadline", func(t *testing.T) {
		t.Parallel()

		ts := githubApiMock([]*mockResponse{
			mustGetTestMockResponse(t, "/meta", http.StatusOK, &github.APIMeta{
				Hooks: []string{"192.0.2.0/24", "2001:db8::/32"},
			}),
		})
		defer ts.Close()

		meta := &Owner{v3client: mustCreateTestGitHubClient(t, ts.URL)}

		ctx, cancel := context.WithTimeout(t.Context(), time.Minute)
		defer cancel()

		d := dataSourceGithubIpRanges().TestResourceData()
		diags := dataSourceGithubIpRangesRead(ctx, d, meta)
		if diags.HasError() {
			t.Fatalf("unexpected error: %v", diags)
		}

		if got, want := d.Get("hooks_ipv4").([]any), "192.0.2.0/24"; len(got) != 1 || got[0] != want {
			t.Errorf("expected hooks_ipv4 to be [%s], got %v", want, got)
		}
		if got, want := d.Get("hooks_ipv6").([]any), "2001:db8::/32"; len(got) != 1 || got[0] != want {
			t.Errorf("expected hooks_ipv6 to be [%s], got %v", want, got)
		}
	})
}

func TestAccGithubIpRangesDataSource(t *testing.T) {
	t.Parallel()

	t.Run("reads IP ranges with a configured read timeout", func(t *testing.T) {
		t.Parallel()

		config := `
			data "github_ip_ranges" "test" {
			  timeouts {
			    read = "2m"
			  }
			}
		`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("actions_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("actions_ipv6"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("reads IP ranges without error", func(t *testing.T) {
		t.Parallel()

		config := `data "github_ip_ranges" "test" {}`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("actions_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("actions_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("actions_macos_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("actions_macos_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("actions_macos"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("actions"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("api_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("api_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("api"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("dependabot_ipv4"), knownvalue.Null()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("dependabot_ipv6"), knownvalue.Null()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("dependabot"), knownvalue.Null()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("git_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("git_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("git"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("github_enterprise_importer_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("github_enterprise_importer_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("github_enterprise_importer"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("hooks_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("hooks_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("hooks"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("importer_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("importer_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("importer"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("packages_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("packages_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("packages"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("pages_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("pages_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("pages"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("web_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("web_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("web"), knownvalue.NotNull()),
					},
				},
			},
		})
	})
}
