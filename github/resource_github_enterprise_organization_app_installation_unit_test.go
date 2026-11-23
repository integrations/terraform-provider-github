package github

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// terraformUnknown is the value the SDK reads as "known after apply" in a raw
// resource configuration.
const terraformUnknown = "74D93920-ED26-11E3-AC10-0800200C9A66"

// testEnterpriseOrganizationAppInstallationState returns the state of an
// existing installation with the given repository selection.
func testEnterpriseOrganizationAppInstallationState(t *testing.T, selection string, repositories []any) *terraform.InstanceState {
	t.Helper()

	d := schema.TestResourceDataRaw(t, resourceGithubEnterpriseOrganizationAppInstallation().Schema, map[string]any{
		"enterprise_slug":       "ent",
		"organization":          "org",
		"client_id":             "Iv1.abc",
		"repository_selection":  selection,
		"selected_repositories": repositories,
	})
	d.SetId("ent:org:Iv1.abc")
	if err := d.Set("installation_id", "7"); err != nil {
		t.Fatal(err)
	}
	if err := d.Set("app_slug", "app"); err != nil {
		t.Fatal(err)
	}

	return d.State()
}

func TestGithubEnterpriseOrganizationAppInstallationDiff(t *testing.T) {
	t.Parallel()

	state := func(selection string, repos ...any) *terraform.InstanceState {
		return testEnterpriseOrganizationAppInstallationState(t, selection, repos)
	}
	config := func(selection string, repos ...any) map[string]any {
		c := map[string]any{
			"enterprise_slug":      "ent",
			"organization":         "org",
			"client_id":            "Iv1.abc",
			"repository_selection": selection,
		}
		if len(repos) > 0 {
			c["selected_repositories"] = repos
		}
		return c
	}

	for _, tt := range []struct {
		name        string
		state       *terraform.InstanceState
		config      map[string]any
		wantErr     bool
		wantNoDiff  bool
		wantReplace bool
	}{
		{name: "unchanged none", state: state("none"), config: config("none"), wantNoDiff: true},
		{name: "unchanged all", state: state("all"), config: config("all"), wantNoDiff: true},
		{name: "all to selected updates in place", state: state("all"), config: config("selected", "a")},
		{name: "selected to all updates in place", state: state("selected", "a"), config: config("all")},
		{name: "selected repositories change in place", state: state("selected", "a"), config: config("selected", "b")},
		{name: "all to none replaces", state: state("all"), config: config("none"), wantReplace: true},
		{name: "none to all replaces", state: state("none"), config: config("all"), wantReplace: true},
		{name: "unknown selection replaces", state: state("all"), config: config(terraformUnknown), wantReplace: true},
		{name: "selected without repositories", state: state("all"), config: config("selected"), wantErr: true},
		{name: "repositories without selected", state: state("all"), config: config("all", "a"), wantErr: true},
		{name: "new installation", config: config("none")},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			diff, err := resourceGithubEnterpriseOrganizationAppInstallation().Diff(t.Context(), tt.state, terraform.NewResourceConfigRaw(tt.config), nil)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected an error, got none")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if tt.wantNoDiff {
				if diff != nil && len(diff.Attributes) > 0 {
					t.Fatalf("expected no diff, got: %v", diff.Attributes)
				}
				return
			}
			if diff == nil {
				t.Fatal("expected a diff, got none")
			}
			if tt.state != nil && diff.RequiresNew() != tt.wantReplace {
				t.Fatalf("expected replacement to be %t, got %t", tt.wantReplace, diff.RequiresNew())
			}
		})
	}
}

func TestGithubEnterpriseOrganizationAppInstallationBatching(t *testing.T) {
	t.Parallel()

	repositories := func(prefix string, n int) []any {
		names := make([]any, n)
		for i := range names {
			names[i] = prefix + strconv.Itoa(i)
		}
		return names
	}

	type request struct {
		method       string
		path         string
		repositories int
	}
	newServer := func(t *testing.T) (*Owner, *[]request) {
		t.Helper()

		var requests []request
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body, _ := io.ReadAll(r.Body)
			var payload struct {
				Repositories []string `json:"repositories"`
			}
			_ = json.Unmarshal(body, &payload)
			requests = append(requests, request{method: r.Method, path: r.URL.Path, repositories: len(payload.Repositories)})

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			switch r.Method {
			case http.MethodPost:
				w.WriteHeader(http.StatusCreated)
				_, _ = w.Write([]byte(`{"id": 7, "app_slug": "app"}`))
			default:
				_, _ = w.Write([]byte(`[]`))
			}
		}))
		t.Cleanup(ts.Close)

		return &Owner{name: "org", v3client: mustCreateTestGitHubClient(t, ts.URL)}, &requests
	}

	const installations = "/enterprises/ent/apps/organizations/org/installations"

	t.Run("create grants repositories beyond the first request separately", func(t *testing.T) {
		t.Parallel()

		meta, requests := newServer(t)
		d := schema.TestResourceDataRaw(t, resourceGithubEnterpriseOrganizationAppInstallation().Schema, map[string]any{
			"enterprise_slug":       "ent",
			"organization":          "org",
			"client_id":             "Iv1.abc",
			"repository_selection":  "selected",
			"selected_repositories": repositories("repo-", 51),
		})

		if diags := resourceGithubEnterpriseOrganizationAppInstallationCreate(t.Context(), d, meta); diags.HasError() {
			t.Fatalf("unexpected error: %v", diags)
		}

		want := []request{
			{method: http.MethodPost, path: installations, repositories: 50},
			{method: http.MethodPatch, path: installations + "/7/repositories/add", repositories: 1},
		}
		if len(*requests) != len(want) {
			t.Fatalf("expected %d requests, got %d: %v", len(want), len(*requests), *requests)
		}
		for i, w := range want {
			if (*requests)[i] != w {
				t.Fatalf("request %d: expected %v, got %v", i, w, (*requests)[i])
			}
		}
		if d.Id() != "ent:org:Iv1.abc" || d.Get("installation_id") != "7" {
			t.Fatalf("unexpected state: id %q, installation_id %q", d.Id(), d.Get("installation_id"))
		}
	})

	t.Run("update adds before it removes, in requests the API accepts", func(t *testing.T) {
		t.Parallel()

		meta, requests := newServer(t)
		r := resourceGithubEnterpriseOrganizationAppInstallation()
		state := testEnterpriseOrganizationAppInstallationState(t, "selected", repositories("old-", 51))
		diff, err := r.Diff(t.Context(), state, terraform.NewResourceConfigRaw(map[string]any{
			"enterprise_slug":       "ent",
			"organization":          "org",
			"client_id":             "Iv1.abc",
			"repository_selection":  "selected",
			"selected_repositories": repositories("new-", 51),
		}), meta)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		d, err := schema.InternalMap(r.Schema).Data(state, diff)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if diags := resourceGithubEnterpriseOrganizationAppInstallationUpdate(t.Context(), d, meta); diags.HasError() {
			t.Fatalf("unexpected error: %v", diags)
		}

		want := []request{
			{method: http.MethodPatch, path: installations + "/7/repositories/add", repositories: 50},
			{method: http.MethodPatch, path: installations + "/7/repositories/add", repositories: 1},
			{method: http.MethodPatch, path: installations + "/7/repositories/remove", repositories: 50},
			{method: http.MethodPatch, path: installations + "/7/repositories/remove", repositories: 1},
		}
		if len(*requests) != len(want) {
			t.Fatalf("expected %d requests, got %d: %v", len(want), len(*requests), *requests)
		}
		for i, w := range want {
			if (*requests)[i] != w {
				t.Fatalf("request %d: expected %v, got %v", i, w, (*requests)[i])
			}
		}
	})
}

func TestGithubEnterpriseOrganizationAppInstallationImport(t *testing.T) {
	t.Parallel()

	r := resourceGithubEnterpriseOrganizationAppInstallation()

	t.Run("sets what read needs from the ID", func(t *testing.T) {
		t.Parallel()

		d := r.Data(nil)
		d.SetId("ent:org:Iv1.abc")

		got, err := resourceGithubEnterpriseOrganizationAppInstallationImport(t.Context(), d, nil)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if len(got) != 1 {
			t.Fatalf("expected one resource, got %d", len(got))
		}
		for key, want := range map[string]string{"enterprise_slug": "ent", "organization": "org", "client_id": "Iv1.abc"} {
			if value := got[0].Get(key); value != want {
				t.Fatalf("expected %s to be %q, got %v", key, want, value)
			}
		}
	})

	t.Run("refuses an ID that is not three parts", func(t *testing.T) {
		t.Parallel()

		d := r.Data(nil)
		d.SetId("ent:org")

		if _, err := resourceGithubEnterpriseOrganizationAppInstallationImport(t.Context(), d, nil); err == nil {
			t.Fatal("expected an error, got none")
		}
	})
}
