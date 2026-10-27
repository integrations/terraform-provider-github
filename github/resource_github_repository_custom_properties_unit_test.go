package github

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sync"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// fakeCustomPropertiesAPI is a minimal in-memory GitHub API for the repository custom properties endpoints.
type fakeCustomPropertiesAPI struct {
	mu          sync.Mutex
	repoExists  bool
	values      map[string]any
	patchCalls  int
	lastPatched []map[string]any
}

func newFakeCustomPropertiesAPI(t *testing.T, repoExists bool, values map[string]any) (*fakeCustomPropertiesAPI, *Owner) {
	t.Helper()

	api := &fakeCustomPropertiesAPI{repoExists: repoExists, values: values}
	if api.values == nil {
		api.values = map[string]any{}
	}

	ts := httptest.NewServer(http.HandlerFunc(api.serve))
	t.Cleanup(ts.Close)

	base := ts.URL + "/"
	client, err := github.NewClient(github.WithURLs(&base, &base))
	if err != nil {
		t.Fatal(err)
	}

	return api, &Owner{name: "o", IsOrganization: true, v3client: client}
}

func (f *fakeCustomPropertiesAPI) serve(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")

	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/orgs/o/properties/schema":
		_, _ = w.Write([]byte(`[
			{"property_name":"env","value_type":"single_select","allowed_values":["production","staging"]},
			{"property_name":"team","value_type":"string"},
			{"property_name":"langs","value_type":"multi_select","allowed_values":["go","ts"]}
		]`))
	case r.Method == http.MethodGet && r.URL.Path == "/repos/o/r":
		if !f.repoExists {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"message":"Not Found"}`))
			return
		}
		_, _ = w.Write([]byte(`{"id":42,"name":"r"}`))
	case r.Method == http.MethodGet && r.URL.Path == "/repos/o/r/properties/values":
		out := make([]map[string]any, 0, len(f.values))
		for name, v := range f.values {
			out = append(out, map[string]any{"property_name": name, "value": v})
		}
		_ = json.NewEncoder(w).Encode(out)
	case r.Method == http.MethodPatch && r.URL.Path == "/repos/o/r/properties/values":
		var body struct {
			Properties []map[string]any `json:"properties"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		f.patchCalls++
		f.lastPatched = body.Properties
		for _, p := range body.Properties {
			name := p["property_name"].(string)
			if p["value"] == nil {
				delete(f.values, name)
			} else {
				f.values[name] = p["value"]
			}
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message":"Not Found"}`))
	}
}

func (f *fakeCustomPropertiesAPI) snapshot() (map[string]any, int) {
	f.mu.Lock()
	defer f.mu.Unlock()

	values := make(map[string]any, len(f.values))
	for k, v := range f.values {
		values[k] = v
	}
	return values, f.patchCalls
}

func customPropertiesUnitProviderFactories(owner *Owner) map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"github": func() (*schema.Provider, error) {
			return &schema.Provider{
				ResourcesMap: map[string]*schema.Resource{
					"github_repository_custom_properties": resourceGithubRepositoryCustomProperties(),
				},
				ConfigureContextFunc: func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
					return owner, nil
				},
			}, nil
		},
	}
}

func TestUnitGithubRepositoryCustomProperties(t *testing.T) {
	t.Run("clears a property whose block is removed from config", func(t *testing.T) {
		api, owner := newFakeCustomPropertiesAPI(t, true, nil)

		resource.UnitTest(t, resource.TestCase{
			ProviderFactories: customPropertiesUnitProviderFactories(owner),
			Steps: []resource.TestStep{
				{
					Config: `
						resource "github_repository_custom_properties" "test" {
							repository = "r"
							property {
								name  = "env"
								value = ["production"]
							}
							property {
								name  = "team"
								value = ["platform"]
							}
						}
					`,
				},
				{
					Config: `
						resource "github_repository_custom_properties" "test" {
							repository = "r"
							property {
								name  = "env"
								value = ["production"]
							}
						}
					`,
					Check: func(*terraform.State) error {
						values, _ := api.snapshot()
						if _, ok := values["team"]; ok {
							t.Errorf("expected property %q to be cleared on GitHub, got values %v", "team", values)
						}
						if values["env"] != "production" {
							t.Errorf("expected property %q to be kept, got values %v", "env", values)
						}
						return nil
					},
				},
			},
		})
	})

	t.Run("rejects multiple values for a single-valued property type", func(t *testing.T) {
		api, owner := newFakeCustomPropertiesAPI(t, true, nil)

		resource.UnitTest(t, resource.TestCase{
			ProviderFactories: customPropertiesUnitProviderFactories(owner),
			Steps: []resource.TestStep{
				{
					Config: `
						resource "github_repository_custom_properties" "test" {
							repository = "r"
							property {
								name  = "team"
								value = ["a", "b"]
							}
						}
					`,
					ExpectError: regexp.MustCompile(`requires exactly one value, got 2`),
				},
			},
		})

		if _, patches := api.snapshot(); patches != 0 {
			t.Errorf("expected no writes to GitHub, got %d", patches)
		}
	})

	t.Run("rejects duplicate property names", func(t *testing.T) {
		_, owner := newFakeCustomPropertiesAPI(t, true, nil)

		resource.UnitTest(t, resource.TestCase{
			ProviderFactories: customPropertiesUnitProviderFactories(owner),
			Steps: []resource.TestStep{
				{
					Config: `
						resource "github_repository_custom_properties" "test" {
							repository = "r"
							property {
								name  = "team"
								value = ["a"]
							}
							property {
								name  = "team"
								value = ["b"]
							}
						}
					`,
					PlanOnly:    true,
					ExpectError: regexp.MustCompile(`declared in more than one property block`),
				},
			},
		})
	})

	t.Run("does not write properties when the repository lookup fails", func(t *testing.T) {
		api, owner := newFakeCustomPropertiesAPI(t, false, nil)

		resource.UnitTest(t, resource.TestCase{
			ProviderFactories: customPropertiesUnitProviderFactories(owner),
			Steps: []resource.TestStep{
				{
					Config: `
						resource "github_repository_custom_properties" "test" {
							repository = "r"
							property {
								name  = "team"
								value = ["a"]
							}
						}
					`,
					ExpectError: regexp.MustCompile(`Not Found`),
				},
			},
		})

		if _, patches := api.snapshot(); patches != 0 {
			t.Errorf("expected no writes to GitHub, got %d", patches)
		}
	})

	t.Run("import of a repository without property values fails with a clear error", func(t *testing.T) {
		_, owner := newFakeCustomPropertiesAPI(t, true, nil)

		resource.UnitTest(t, resource.TestCase{
			ProviderFactories: customPropertiesUnitProviderFactories(owner),
			Steps: []resource.TestStep{
				{
					Config:        `resource "github_repository_custom_properties" "test" {` + "\n" + `repository = "r"` + "\n" + `property {` + "\n" + `name = "team"` + "\n" + `value = ["a"]` + "\n" + `}` + "\n" + `}`,
					ResourceName:  "github_repository_custom_properties.test",
					ImportState:   true,
					ImportStateId: "r",
					ExpectError:   regexp.MustCompile(`nothing to import`),
				},
			},
		})
	})
}
