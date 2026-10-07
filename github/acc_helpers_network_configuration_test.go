package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func testRunnerGroupNetworking(t *testing.T, runnerGroup *schema.Resource, scope string, attributes map[string]any) {
	t.Helper()

	path := scope + "/actions/runner-groups"
	networkID := ""
	writes := 0
	reads := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == path,
			r.Method == http.MethodPatch && r.URL.Path == path+"/42":
			var body struct {
				NetworkConfigurationID string `json:"network_configuration_id"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("decode runner group request: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			writes++
			expected := fmt.Sprintf("network-%d", writes)
			if body.NetworkConfigurationID != expected {
				t.Errorf("network_configuration_id in %s = %q, want %q", r.Method, body.NetworkConfigurationID, expected)
			}
			networkID = body.NetworkConfigurationID
		case r.Method == http.MethodGet && r.URL.Path == path+"/42":
			reads++
		case r.Method == http.MethodGet && (r.URL.Path == path+"/42/repositories" || r.URL.Path == path+"/42/organizations"):
			fmt.Fprint(w, `{"repositories":[],"organizations":[]}`)
			return
		case r.Method == http.MethodPut && (r.URL.Path == path+"/42/repositories" || r.URL.Path == path+"/42/organizations"):
			w.WriteHeader(http.StatusNoContent)
			return
		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(map[string]any{
			"id":                       42,
			"name":                     "test-group",
			"visibility":               "all",
			"network_configuration_id": networkID,
		}); err != nil {
			t.Errorf("encode runner group response: %v", err)
		}
	}))
	defer server.Close()

	meta := &Owner{
		name:           "test-org",
		IsOrganization: true,
		maxPerPage:     100,
		v3client:       mustCreateTestGitHubClient(t, server.URL),
	}
	d := schema.TestResourceDataRaw(t, runnerGroup.Schema, attributes)
	if diags := runnerGroup.CreateContext(t.Context(), d, meta); diags.HasError() {
		t.Fatalf("create runner group: %v", diags)
	}
	if d.Id() != "42" || d.Get("network_configuration_id") != "network-1" {
		t.Fatalf("unexpected create state: ID=%q, network_configuration_id=%v", d.Id(), d.Get("network_configuration_id"))
	}
	if reads != 0 {
		t.Fatalf("create issued %d read-after-write requests", reads)
	}

	if err := d.Set("network_configuration_id", "network-2"); err != nil {
		t.Fatal(err)
	}
	if diags := runnerGroup.UpdateContext(t.Context(), d, meta); diags.HasError() {
		t.Fatalf("update runner group: %v", diags)
	}
	if d.Id() != "42" || d.Get("network_configuration_id") != "network-2" {
		t.Fatalf("unexpected update state: ID=%q, network_configuration_id=%v", d.Id(), d.Get("network_configuration_id"))
	}
	if writes != 2 {
		t.Errorf("write requests = %d, want 2", writes)
	}
	if reads != 0 {
		t.Fatalf("update issued %d read-after-write requests", reads)
	}

	importAttributes := maps.Clone(attributes)
	delete(importAttributes, "network_configuration_id")
	imported := schema.TestResourceDataRaw(t, runnerGroup.Schema, importAttributes)
	imported.SetId("42")
	if diags := runnerGroup.ReadContext(t.Context(), imported, meta); diags.HasError() {
		t.Fatalf("read unconfigured network assignment: %v", diags)
	}
	if got := imported.Get("network_configuration_id"); got != "network-2" {
		t.Errorf("imported network_configuration_id = %v, want network-2", got)
	}
}

func testRunnerGroupNetworkingLifecycle(t *testing.T, runnerGroup func() *schema.Resource, resourceType, scope, enterpriseSlug string) {
	t.Helper()

	path := scope + "/actions/runner-groups"
	var mu sync.Mutex
	groups := map[string]map[string]any{}
	nextID := 41
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		w.Header().Set("Content-Type", "application/json")

		parts := strings.Split(strings.TrimPrefix(r.URL.Path, path+"/"), "/")
		id := parts[0]
		group := groups[id]
		switch {
		case r.Method == http.MethodPost && r.URL.Path == path:
			nextID++
			id = strconv.Itoa(nextID)
			group = map[string]any{
				"id":                       nextID,
				"network_configuration_id": nil,
				"default":                  false,
				"inherited":                false,
				"runners_url":              path + "/" + id + "/runners",
			}
			groups[id] = group
		case group == nil:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, `{"message":"Not Found"}`)
			return
		case len(parts) == 2 && (parts[1] == "repositories" || parts[1] == "organizations"):
			switch r.Method {
			case http.MethodGet:
				fmt.Fprint(w, `{"repositories":[],"organizations":[]}`)
			case http.MethodPut:
				w.WriteHeader(http.StatusNoContent)
			default:
				t.Errorf("unexpected access request: %s %s", r.Method, r.URL.Path)
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
			return
		case r.Method == http.MethodDelete:
			delete(groups, id)
			w.WriteHeader(http.StatusNoContent)
			return
		case r.Method != http.MethodGet && r.Method != http.MethodPatch:
			t.Errorf("unexpected runner group request: %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if r.Method == http.MethodPost || r.Method == http.MethodPatch {
			var request map[string]any
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				t.Errorf("decode runner group request: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if request["network_configuration_id"] == "" {
				t.Errorf("provider sent an empty-string network_configuration_id in a %s request", r.Method)
			}
			maps.Copy(group, request)
		}
		if err := json.NewEncoder(w).Encode(group); err != nil {
			t.Errorf("encode runner group response: %v", err)
		}
	}))
	t.Cleanup(server.Close)

	meta := &Owner{
		name:           "test-org",
		IsOrganization: true,
		maxPerPage:     100,
		v3client:       mustCreateTestGitHubClient(t, server.URL),
	}
	enterpriseConfig := ""
	importPrefix := ""
	if enterpriseSlug != "" {
		enterpriseConfig = fmt.Sprintf("enterprise_slug = %q", enterpriseSlug)
		importPrefix = enterpriseSlug + "/"
	}
	config := fmt.Sprintf(`
		resource %q "test" {
		  %s
		  name       = "test-group"
		  visibility = "all"
		  %%s
		}
	`, resourceType, enterpriseConfig)
	resourceName := resourceType + ".test"
	unchangedID := statecheck.CompareValue(compare.ValuesSame())
	replacedID := statecheck.CompareValue(compare.ValuesDiffer())
	attachedConfig := fmt.Sprintf(config, `network_configuration_id = "network-1"`)

	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"github": func() (*schema.Provider, error) {
				provider := &schema.Provider{
					ResourcesMap: map[string]*schema.Resource{resourceType: runnerGroup()},
					ConfigureContextFunc: func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
						return meta, nil
					},
				}
				return provider, provider.InternalValidate()
			},
		},
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(config, ""),
				ConfigStateChecks: []statecheck.StateCheck{
					unchangedID.AddStateValue(resourceName, tfjsonpath.New("id")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_configuration_id"), knownvalue.StringExact("")),
				},
			},
			{
				Config: attachedConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					unchangedID.AddStateValue(resourceName, tfjsonpath.New("id")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_configuration_id"), knownvalue.StringExact("network-1")),
				},
			},
			{
				ResourceName:        resourceName,
				ImportState:         true,
				ImportStateVerify:   true,
				ImportStateIdPrefix: importPrefix,
			},
			{
				Config: fmt.Sprintf(config, ""),
				ConfigStateChecks: []statecheck.StateCheck{
					unchangedID.AddStateValue(resourceName, tfjsonpath.New("id")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_configuration_id"), knownvalue.StringExact("network-1")),
				},
			},
			{
				Config: fmt.Sprintf(config, `network_configuration_id = null`),
				ConfigStateChecks: []statecheck.StateCheck{
					unchangedID.AddStateValue(resourceName, tfjsonpath.New("id")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_configuration_id"), knownvalue.StringExact("network-1")),
				},
			},
			{
				PreConfig: func() {
					mu.Lock()
					defer mu.Unlock()
					for _, group := range groups {
						group["network_configuration_id"] = "network-external"
					}
				},
				Config: strings.Replace(fmt.Sprintf(config, ""), `"test-group"`, `"renamed-group"`, 1),
				ConfigStateChecks: []statecheck.StateCheck{
					unchangedID.AddStateValue(resourceName, tfjsonpath.New("id")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_configuration_id"), knownvalue.StringExact("network-external")),
				},
			},
			{
				Config: fmt.Sprintf(config, `network_configuration_id = "network-2"`),
				ConfigStateChecks: []statecheck.StateCheck{
					unchangedID.AddStateValue(resourceName, tfjsonpath.New("id")),
					replacedID.AddStateValue(resourceName, tfjsonpath.New("id")),
				},
			},
			{
				Config: fmt.Sprintf(config, `network_configuration_id = ""`),
				ConfigStateChecks: []statecheck.StateCheck{
					replacedID.AddStateValue(resourceName, tfjsonpath.New("id")),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_configuration_id"), knownvalue.StringExact("")),
				},
			},
		},
	})
	mu.Lock()
	defer mu.Unlock()
	if len(groups) != 0 {
		t.Errorf("runner groups remain after destroy: %v", groups)
	}
}

func testRunnerGroupContextCancellation(t *testing.T, runnerGroup *schema.Resource, attributes map[string]any) {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("cancelled CRUD context reached the API: %s %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	t.Cleanup(server.Close)

	meta := &Owner{
		name:           "test-org",
		IsOrganization: true,
		v3client:       mustCreateTestGitHubClient(t, server.URL),
	}
	for name, handler := range map[string]func(context.Context, *schema.ResourceData, any) diag.Diagnostics{
		"create": runnerGroup.CreateContext,
		"read":   runnerGroup.ReadContext,
		"update": runnerGroup.UpdateContext,
		"delete": runnerGroup.DeleteContext,
	} {
		t.Run(name, func(t *testing.T) {
			if handler == nil {
				t.Fatal("context-aware CRUD handler is not registered")
			}
			d := schema.TestResourceDataRaw(t, runnerGroup.Schema, attributes)
			d.SetId("42")
			ctx, cancel := context.WithCancel(t.Context())
			cancel()

			diags := handler(ctx, d, meta)
			if !diags.HasError() || !strings.Contains(fmt.Sprint(diags), context.Canceled.Error()) {
				t.Fatalf("expected context cancellation diagnostic, got %v", diags)
			}
		})
	}
}

func mustCreateTestOrganizationNetworkConfiguration(t *testing.T) *github.NetworkConfiguration {
	t.Helper()
	skipUnlessHasPaidOrgs(t)

	name := fmt.Sprintf("%snetwork-%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
	configuration, _, err := testAccConf.meta.v3client.Organizations.CreateNetworkConfiguration(t.Context(), testAccConf.meta.name, github.NetworkConfigurationRequest{
		Name:               new(name),
		ComputeService:     new(github.ComputeServiceActions),
		NetworkSettingsIDs: []string{testAccOrganizationNetworkConfigurationID(t)},
	})
	if err != nil {
		t.Fatalf("failed to create test organization network configuration: %v", err)
	}

	t.Cleanup(func() {
		_, err := testAccConf.meta.v3client.Organizations.DeleteNetworkConfigurations(context.Background(), testAccConf.meta.name, configuration.GetID())
		if err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
				return
			}
			t.Errorf("failed to delete test organization network configuration %s: %v", configuration.GetID(), err)
		}
	})

	return configuration
}

func mustCreateTestEnterpriseNetworkConfiguration(t *testing.T) *github.NetworkConfiguration {
	t.Helper()
	skipUnlessEnterprise(t)

	name := fmt.Sprintf("%snetwork-%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
	configuration, _, err := testAccConf.meta.v3client.Enterprise.CreateEnterpriseNetworkConfiguration(t.Context(), testAccConf.enterpriseSlug, github.NetworkConfigurationRequest{
		Name:               new(name),
		ComputeService:     new(github.ComputeServiceActions),
		NetworkSettingsIDs: []string{testAccEnterpriseNetworkConfigurationID(t)},
	})
	if err != nil {
		t.Fatalf("failed to create test enterprise network configuration: %v", err)
	}

	t.Cleanup(func() {
		_, err := testAccConf.meta.v3client.Enterprise.DeleteEnterpriseNetworkConfiguration(context.Background(), testAccConf.enterpriseSlug, configuration.GetID())
		if err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
				return
			}
			t.Errorf("failed to delete test enterprise network configuration %s: %v", configuration.GetID(), err)
		}
	})

	return configuration
}
