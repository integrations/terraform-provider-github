package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

func testRunnerGroupNetworking(t *testing.T, runnerGroup *schema.Resource, scope string, attributes map[string]any, create, update func(*schema.ResourceData, any) error) {
	t.Helper()

	path := scope + "/actions/runner-groups"
	networkID := ""
	writes := 0
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
	if err := create(d, meta); err != nil {
		t.Fatalf("create runner group: %v", err)
	}
	if d.Id() != "42" || d.Get("network_configuration_id") != "network-1" {
		t.Fatalf("unexpected create state: ID=%q, network_configuration_id=%v", d.Id(), d.Get("network_configuration_id"))
	}

	if err := d.Set("network_configuration_id", "network-2"); err != nil {
		t.Fatal(err)
	}
	if err := update(d, meta); err != nil {
		t.Fatalf("update runner group: %v", err)
	}
	if d.Id() != "42" || d.Get("network_configuration_id") != "network-2" {
		t.Fatalf("unexpected update state: ID=%q, network_configuration_id=%v", d.Id(), d.Get("network_configuration_id"))
	}
	if writes != 2 {
		t.Errorf("write requests = %d, want 2", writes)
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
