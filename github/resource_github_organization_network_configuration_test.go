package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdkterraform "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubOrganizationNetworkConfiguration(t *testing.T) {
	t.Run("create, import, and update in place", func(t *testing.T) {
		networkSettingsID := testAccOrganizationNetworkConfigurationID(t)

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "github_organization_network_configuration.test"
		configurationName := fmt.Sprintf("%snetwork-config-%s", testResourcePrefix, randomID)
		sameID := statecheck.CompareValue(compare.ValuesSame())
		createdOn := knownvalue.StringFunc(func(value string) error {
			if value == "" {
				return nil
			}
			_, err := time.Parse(time.RFC3339, value)
			return err
		})

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubOrganizationNetworkConfigurationDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccOrganizationNetworkConfigurationConfig(configurationName, "actions", networkSettingsID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringRegexp(regexp.MustCompile(`^\S+$`))),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), createdOn),
						sameID.AddStateValue(resourceName, tfjsonpath.New("id")),
					},
				},
				{
					ResourceName:      resourceName,
					ImportState:       true,
					ImportStateVerify: true,
				},
				{
					Config: testAccOrganizationNetworkConfigurationConfig(configurationName+"-updated", "none", networkSettingsID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), createdOn),
						sameID.AddStateValue(resourceName, tfjsonpath.New("id")),
					},
				},
				{
					ResourceName:      resourceName,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}

func TestNetworkConfigurationSchema(t *testing.T) {
	for scope, r := range map[string]*schema.Resource{
		"organization": resourceGithubOrganizationNetworkConfiguration(),
		"enterprise":   resourceGithubEnterpriseNetworkConfiguration(),
	} {
		t.Run(scope, func(t *testing.T) {
			if err := r.InternalValidate(nil, true); err != nil {
				t.Fatal(err)
			}
			p := &schema.Provider{ResourcesMap: map[string]*schema.Resource{"github_network_configuration": r}}

			for _, tc := range []struct {
				name      string
				attribute string
				value     any
				wantError bool
			}{
				{name: "short name", attribute: "name", value: "a"},
				{name: "valid name characters", attribute: "name", value: "Network_123.test-1"},
				{name: "maximum name length", attribute: "name", value: strings.Repeat("a", 100)},
				{name: "empty name", attribute: "name", value: "", wantError: true},
				{name: "long name", attribute: "name", value: strings.Repeat("a", 101), wantError: true},
				{name: "invalid name characters", attribute: "name", value: "network config", wantError: true},
				{name: "none compute service", attribute: "compute_service", value: "none"},
				{name: "actions compute service", attribute: "compute_service", value: "actions"},
				{name: "invalid compute service", attribute: "compute_service", value: "codespaces", wantError: true},
				{name: "case sensitive compute service", attribute: "compute_service", value: "Actions", wantError: true},
				{name: "one network settings ID", attribute: "network_settings_ids", value: []any{"NS_123"}},
				{name: "no network settings IDs", attribute: "network_settings_ids", value: []any{}, wantError: true},
				{name: "multiple network settings IDs", attribute: "network_settings_ids", value: []any{"NS_123", "NS_456"}, wantError: true},
				{name: "empty network settings ID", attribute: "network_settings_ids", value: []any{""}, wantError: true},
				{name: "blank network settings ID", attribute: "network_settings_ids", value: []any{" \t"}, wantError: true},
			} {
				t.Run(tc.name, func(t *testing.T) {
					config := map[string]any{"name": "network", "network_settings_ids": []any{"NS_123"}}
					if scope == "enterprise" {
						config["enterprise_slug"] = "my-enterprise"
					}
					config[tc.attribute] = tc.value
					diags := p.ValidateResource("github_network_configuration", sdkterraform.NewResourceConfigRaw(config))
					if diags.HasError() != tc.wantError {
						t.Fatalf("validation diagnostics = %v, want error = %t", diags, tc.wantError)
					}
				})
			}

			d := schema.TestResourceDataRaw(t, r.Schema, nil)
			if got := d.Get("compute_service"); got != "none" {
				t.Errorf("default compute_service = %v, want none", got)
			}
		})
	}
}

func TestNetworkSettingsScopeError(t *testing.T) {
	for _, scope := range []string{"organization", "enterprise"} {
		t.Run(scope, func(t *testing.T) {
			apiErr := &github.ErrorResponse{Response: &http.Response{StatusCode: http.StatusUnprocessableEntity}}
			err := networkSettingsScopeError(apiErr, scope)
			if !errors.Is(err, apiErr) || !strings.Contains(err.Error(), "same "+scope) {
				t.Fatalf("expected wrapped error with %s scope guidance, got %v", scope, err)
			}
		})
	}

	for _, err := range []error{
		errors.New("transport failure"),
		&github.ErrorResponse{Response: &http.Response{StatusCode: http.StatusForbidden}},
	} {
		if got := networkSettingsScopeError(err, "organization"); !errors.Is(got, err) {
			t.Errorf("non-422 error was changed: %v", got)
		}
	}
}

func TestNetworkConfigurationState(t *testing.T) {
	for scope, r := range map[string]*schema.Resource{
		"organization": resourceGithubOrganizationNetworkConfiguration(),
		"enterprise":   resourceGithubEnterpriseNetworkConfiguration(),
	} {
		for _, service := range []github.ComputeService{"actions", "none"} {
			t.Run(scope+"/"+string(service), func(t *testing.T) {
				d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
					"name":                 "stale",
					"compute_service":      "actions",
					"network_settings_ids": []any{"NS_OLD"},
				})
				d.SetId("NC_123")
				createdOn := time.Date(2026, time.September, 1, 12, 0, 0, 0, time.UTC)
				configuration := &github.NetworkConfiguration{
					ID:                 new("NC_123"),
					Name:               new("network-updated"),
					ComputeService:     new(service),
					NetworkSettingsIDs: []string{"NS_NEW"},
					CreatedOn:          &github.Timestamp{Time: createdOn},
				}
				if err := setNetworkConfigurationState(d, configuration); err != nil {
					t.Fatal(err)
				}
				for key, want := range map[string]string{
					"name":            configuration.GetName(),
					"compute_service": string(service),
					"created_on":      createdOn.Format(time.RFC3339),
				} {
					if got := d.Get(key); got != want {
						t.Errorf("%s = %v, want %s", key, got, want)
					}
				}
				configuredIDs, _ := d.Get("network_settings_ids").([]any)
				ids := expandStringList(configuredIDs)
				if !slices.Equal(ids, configuration.NetworkSettingsIDs) {
					t.Errorf("network_settings_ids = %v, want %v", ids, configuration.NetworkSettingsIDs)
				}
				if d.Id() != "NC_123" {
					t.Errorf("refresh changed the resource ID to %q", d.Id())
				}
				configuration.ComputeService = nil
				configuration.CreatedOn = nil
				if err := setNetworkConfigurationState(d, configuration); err != nil {
					t.Fatal(err)
				}
				if got := d.Get("compute_service"); got != string(service) {
					t.Errorf("refresh changed configured compute_service to %v when GitHub omitted it", got)
				}
				if got := d.Get("created_on"); got != "" {
					t.Errorf("refresh retained stale created_on = %v when GitHub returned no timestamp", got)
				}
			})
		}
	}
}

func TestGithubOrganizationNetworkConfigurationImport(t *testing.T) {
	r := resourceGithubOrganizationNetworkConfiguration()
	d := schema.TestResourceDataRaw(t, r.Schema, nil)
	d.SetId("NC_123")
	states, err := r.Importer.StateContext(context.Background(), d, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(states) != 1 || states[0].Id() != "NC_123" {
		t.Fatalf("import did not preserve the network configuration ID: %v", states)
	}
}

func TestNetworkConfigurationCRUD(t *testing.T) {
	for _, scope := range []struct {
		name string
		path string
		r    *schema.Resource
	}{
		{
			name: "organization",
			path: "/orgs/test-org/settings/network-configurations",
			r:    resourceGithubOrganizationNetworkConfiguration(),
		},
		{
			name: "enterprise",
			path: "/enterprises/test-enterprise/network-configurations",
			r:    resourceGithubEnterpriseNetworkConfiguration(),
		},
	} {
		t.Run(scope.name, func(t *testing.T) {
			for _, operation := range []struct {
				name            string
				method          string
				success         int
				ignoredStatuses []int
				run             func(context.Context, *schema.ResourceData, any) diag.Diagnostics
			}{
				{name: "create", method: http.MethodPost, success: http.StatusCreated, run: scope.r.CreateContext},
				{name: "read", method: http.MethodGet, success: http.StatusOK, ignoredStatuses: []int{http.StatusNotModified, http.StatusNotFound}, run: scope.r.ReadContext},
				{name: "update", method: http.MethodPatch, success: http.StatusOK, run: scope.r.UpdateContext},
				{name: "delete", method: http.MethodDelete, success: http.StatusNoContent, ignoredStatuses: []int{http.StatusNotFound}, run: scope.r.DeleteContext},
			} {
				for _, status := range []int{operation.success, http.StatusNotModified, http.StatusNotFound, http.StatusForbidden, http.StatusUnprocessableEntity, http.StatusInternalServerError} {
					t.Run(fmt.Sprintf("%s/%d", operation.name, status), func(t *testing.T) {
						config := map[string]any{
							"name":                 "configured",
							"compute_service":      "actions",
							"network_settings_ids": []any{"NS_CONFIGURED"},
						}
						if scope.name == "enterprise" {
							config["enterprise_slug"] = "test-enterprise"
						}
						d := schema.TestResourceDataRaw(t, scope.r.Schema, config)
						if operation.name != "create" {
							d.SetId("NC_123")
							if err := d.Set("created_on", "2026-09-01T00:00:00Z"); err != nil {
								t.Fatal(err)
							}
						}
						before := d.State()
						calls := 0
						client, err := github.NewClient(github.WithTransport(localRoundTripper{handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
							calls++
							wantPath := scope.path
							if operation.name != "create" {
								wantPath += "/NC_123"
							}
							if req.Method != operation.method || req.URL.Path != wantPath {
								t.Errorf("request = %s %s, want %s %s", req.Method, req.URL.Path, operation.method, wantPath)
							}
							if operation.name == "create" || operation.name == "update" {
								var body map[string]any
								if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
									t.Fatal(err)
								}
								want := map[string]any{
									"name":                 "configured",
									"compute_service":      "actions",
									"network_settings_ids": []any{"NS_CONFIGURED"},
								}
								if !reflect.DeepEqual(body, want) {
									t.Errorf("request body = %v, want %v", body, want)
								}
							}
							w.Header().Set("Content-Type", "application/json")
							w.WriteHeader(status)
							if status == http.StatusNotModified || status == http.StatusNoContent {
								return
							}
							if status != operation.success {
								mustWrite(w, `{"message":"request rejected"}`)
								return
							}
							mustWrite(w, `{"id":"NC_123","name":"returned","compute_service":"none","network_settings_ids":["NS_RETURNED"],"created_on":"2026-09-02T12:00:00Z"}`)
						})}))
						if err != nil {
							t.Fatal(err)
						}
						meta := &Owner{name: "test-org", IsOrganization: true, v3client: client}
						diags := operation.run(context.Background(), d, meta)
						wantError := status != operation.success && !slices.Contains(operation.ignoredStatuses, status)
						if diags.HasError() != wantError {
							t.Fatalf("diagnostics = %v, want error = %t", diags, wantError)
						}
						if calls != 1 {
							t.Errorf("made %d API calls, want exactly one", calls)
						}
						if wantError && !strings.Contains(fmt.Sprint(diags), fmt.Sprint(status)) {
							t.Errorf("diagnostics lost API status %d: %v", status, diags)
						}
						if status == http.StatusUnprocessableEntity && (operation.name == "create" || operation.name == "update") &&
							!strings.Contains(fmt.Sprint(diags), "same "+scope.name) {
							t.Errorf("missing network settings scope guidance: %v", diags)
						}

						if operation.name == "read" && status == http.StatusNotFound {
							if d.Id() != "" {
								t.Errorf("404 read retained resource ID %q", d.Id())
							}
							return
						}
						if status != operation.success || operation.name == "delete" {
							if after := d.State(); !reflect.DeepEqual(after, before) {
								t.Errorf("unsuccessful request or delete changed state: before %v, after %v", before, after)
							}
							return
						}
						if d.Id() != "NC_123" {
							t.Errorf("resource ID = %q, want NC_123", d.Id())
						}
						for key, want := range map[string]any{
							"name":                 "returned",
							"compute_service":      "none",
							"network_settings_ids": []any{"NS_RETURNED"},
							"created_on":           "2026-09-02T12:00:00Z",
						} {
							if got := d.Get(key); !reflect.DeepEqual(got, want) {
								t.Errorf("%s = %v, want %v", key, got, want)
							}
						}
						if scope.name == "enterprise" && d.Get("enterprise_slug") != "test-enterprise" {
							t.Errorf("enterprise_slug = %v, want test-enterprise", d.Get("enterprise_slug"))
						}
					})
				}
			}
		})
	}
}

func testAccCheckGithubOrganizationNetworkConfigurationDestroy(s *terraform.State) error {
	client := testAccConf.meta.v3client
	orgName := testAccConf.meta.name

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_organization_network_configuration" {
			continue
		}

		_, _, err := client.Organizations.GetNetworkConfiguration(context.Background(), orgName, rs.Primary.ID)
		if err != nil {
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
				continue
			}

			return err
		}

		return fmt.Errorf("organization network configuration still exists: %s", rs.Primary.ID)
	}

	return nil
}

func testAccOrganizationNetworkConfigurationID(t *testing.T) string {
	t.Helper()

	networkSettingsID := os.Getenv("GITHUB_TEST_NETWORK_SETTINGS_ID")
	if networkSettingsID == "" {
		t.Skip("GITHUB_TEST_NETWORK_SETTINGS_ID not set")
	}

	return networkSettingsID
}

func testAccOrganizationNetworkConfigurationConfig(name, computeService, networkSettingsID string) string {
	return fmt.Sprintf(`
resource "github_organization_network_configuration" "test" {
  name                 = %q
  compute_service      = %q
  network_settings_ids = [%q]
}
`, name, computeService, networkSettingsID)
}
