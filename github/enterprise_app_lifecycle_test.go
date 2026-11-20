package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

type enterpriseLifecycleFixture struct {
	lifecycle           *enterpriseAppLifecycle
	owner               *Owner
	data                *schema.ResourceData
	billing             any
	orgNode             string
	inventoryID         int
	installationID      int
	suspended           any
	permission          string
	installerPermission string
	missing             bool
	listed              bool
	listingStatus       int
	listingPages        int
	listingMalformed    bool
	deny                bool
	installPosts        int
	profileWrites       int
	ownerPages          int
	inventoryPages      int
	mutations           int
}

func newEnterpriseLifecycleFixture(t *testing.T) *enterpriseLifecycleFixture {
	t.Helper()
	fixture := &enterpriseLifecycleFixture{billing: "billing@example.test", orgNode: "O_test", inventoryID: 42, installationID: 51, permission: "write", installerPermission: "write"}
	config := &enterpriseAppConfig{slug: "test-enterprise", installerID: "11", installerInstallationID: "12", automationID: "21", automationClientID: "test-client"}
	fixture.lifecycle = &enterpriseAppLifecycle{config: config, enterpriseID: "E_test", automationPermissions: map[string]string{"organization_administration": "write", "members": "read"}}
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		write := func(value any) {
			if err := json.NewEncoder(writer).Encode(value); err != nil {
				t.Error(err)
			}
		}
		switch {
		case strings.HasPrefix(request.URL.Path, "/users/"):
			t.Error("unexpected owner discovery")
			writer.WriteHeader(http.StatusInternalServerError)
		case request.URL.Path == "/app":
			if request.Header.Get("X-Test-Identity") == "installer" {
				write(map[string]any{"id": 11, "client_id": "installer-client"})
			} else {
				write(map[string]any{"id": 21, "client_id": "test-client", "permissions": map[string]string{"organization_administration": "write", "members": "read"}})
			}
		case request.URL.Path == "/app/installations/12":
			write(map[string]any{"id": 12, "app_id": 11, "target_type": "Enterprise", "account": map[string]any{"slug": "test-enterprise"}, "permissions": map[string]string{"enterprise_organization_installations": fixture.installerPermission}})
		case strings.Contains(request.URL.Path, "installable_organizations"):
			fixture.inventoryPages++
			if fixture.deny {
				writer.WriteHeader(http.StatusForbidden)
				write(map[string]string{"message": "denied"})
				return
			}
			if request.URL.Query().Get("page") == "1" {
				writer.Header().Set("Link", fmt.Sprintf("<http://%s%s?per_page=100&page=2>; rel=\"next\"", request.Host, request.URL.Path))
				write([]any{map[string]any{"id": 99, "login": "other-org"}})
			} else {
				write([]any{map[string]any{"id": fixture.inventoryID, "login": "test-org"}})
			}
		case request.URL.Path == "/orgs/test-org/installation":
			if fixture.missing {
				writer.WriteHeader(http.StatusNotFound)
				write(map[string]string{"message": "Not Found"})
				return
			}
			write(map[string]any{
				"id": fixture.installationID, "app_id": 21, "account": map[string]any{"id": 42, "login": "test-org"}, "suspended_at": fixture.suspended,
				"permissions": map[string]string{"organization_administration": fixture.permission, "members": "read"},
			})
		case request.URL.Path == "/enterprises/test-enterprise/apps/organizations/test-org/installations" && request.Method == http.MethodGet:
			fixture.listingPages++
			if fixture.listingStatus != 0 {
				writer.WriteHeader(fixture.listingStatus)
				write(map[string]string{"message": "listing unavailable"})
				return
			}
			if fixture.listingMalformed {
				write([]any{map[string]any{"id": 51}})
				return
			}
			if request.URL.Query().Get("page") == "1" {
				writer.Header().Set("Link", fmt.Sprintf("<http://%s%s?per_page=100&page=2>; rel=\"next\"", request.Host, request.URL.Path))
				write([]any{map[string]any{"id": 70, "client_id": "unrelated-client"}})
			} else if fixture.listed {
				write([]any{map[string]any{"id": 51, "client_id": "test-client", "permissions": map[string]string{"organization_administration": fixture.permission}}})
			} else {
				write([]any{})
			}
		case request.URL.Path == "/enterprises/test-enterprise/apps/organizations/test-org/installations" && request.Method == http.MethodPost:
			fixture.installPosts++
			var body map[string]any
			if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
				t.Error(err)
			}
			if body["client_id"] != "test-client" || body["repository_selection"] != "none" {
				t.Errorf("unsafe install payload: %v", body)
			}
			fixture.missing = false
			write(map[string]any{"id": 51})
		case request.URL.Path == "/orgs/test-org":
			if request.Method == http.MethodPatch {
				fixture.profileWrites++
			}
			write(map[string]any{"id": 42, "node_id": fixture.orgNode, "login": "test-org", "name": "Test Org", "description": "managed", "billing_email": fixture.billing})
		case request.Method == http.MethodPost && (request.URL.Path == "/graphql" || request.URL.Path == "/"):
			var body struct {
				Query     string
				Variables map[string]any
			}
			if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
				t.Error(err)
			}
			if strings.Contains(body.Query, "organizationBillingEmail") {
				t.Error("forbidden billing query")
			}
			if strings.Contains(body.Query, "createEnterpriseOrganization") {
				fixture.mutations++
				write(map[string]any{"data": map[string]any{"createEnterpriseOrganization": map[string]any{"organization": map[string]any{"id": "O_test", "databaseId": 42}}}})
				return
			}
			if strings.Contains(body.Query, "enterprise(") {
				write(map[string]any{"data": map[string]any{"enterprise": map[string]string{"id": "E_test"}}})
				return
			}
			fixture.ownerPages++
			next := body.Variables["cursor"] == nil
			login, role := "owner-one", "ADMIN"
			if !next {
				login = "owner-two"
			}
			write(map[string]any{"data": map[string]any{"organization": map[string]any{
				"id": "O_test", "databaseId": 42, "login": "test-org", "membersWithRole": map[string]any{
					"edges":    []any{map[string]any{"node": map[string]string{"login": login}, "role": role}, map[string]any{"node": map[string]string{"login": "regular-member"}, "role": "MEMBER"}},
					"pageInfo": map[string]any{"hasNextPage": next, "endCursor": "end"},
				},
			}}})
		default:
			t.Errorf("unexpected request: %s %s", request.Method, request.URL)
			writer.WriteHeader(http.StatusInternalServerError)
		}
	})
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	client := mustCreateTestGitHubClient(t, server.URL+"/")
	fixture.lifecycle.installer = client
	fixture.lifecycle.automationJWT = client
	fixture.lifecycle.installerJWT = mustCreateTestGitHubClient(t, server.URL+"/", github.WithTransport(enterpriseTestIdentityTransport{base: http.DefaultTransport}))
	graphql := githubv4.NewEnterpriseClient(server.URL+"/graphql", server.Client())
	fixture.lifecycle.orgClients = func(int64) (*github.Client, *githubv4.Client, error) { return client, graphql, nil }
	fixture.owner = &Owner{enterpriseApp: fixture.lifecycle, v3client: client, v4client: graphql}
	fixture.data = schema.TestResourceDataRaw(t, resourceGithubEnterpriseOrganization().Schema, map[string]any{
		"enterprise_id": "E_test", "name": "test-org", "billing_email": "old@example.test", "admin_logins": []any{"owner-one"}, "database_id": 42,
	})
	fixture.data.SetId("O_test")
	return fixture
}

type enterpriseTestIdentityTransport struct{ base http.RoundTripper }

func (transport enterpriseTestIdentityTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	copy := request.Clone(request.Context())
	copy.Header.Set("X-Test-Identity", "installer")
	return transport.base.RoundTrip(copy)
}

func TestEnterpriseAppLifecycleReadImport(t *testing.T) {
	fixture := newEnterpriseLifecycleFixture(t)
	if diagnostics := enterpriseOrganizationReadContext(t.Context(), fixture.data, fixture.owner); diagnostics.HasError() {
		t.Fatal(diagnostics)
	}
	if fixture.data.Get("billing_email") != "billing@example.test" || fixture.data.Get("automation_installation_id") != 51 {
		t.Fatal("REST billing/installation not populated")
	}
	owners := fixture.data.Get("admin_logins").(*schema.Set)
	if owners.Len() != 2 || owners.Contains("regular-member") || fixture.ownerPages != 2 || fixture.inventoryPages != 2 {
		t.Fatalf("pagination/owners filtering failed: %v", owners.List())
	}
	fixture.data.SetId("test-enterprise/test-org")
	if _, err := enterpriseOrganizationImportContext(t.Context(), fixture.data, fixture.owner); err != nil {
		t.Fatal(err)
	}
	if fixture.data.Id() != "O_test" || fixture.installPosts != 0 || fixture.profileWrites != 0 {
		t.Fatal("import identity or write isolation failed")
	}
}

func TestEnterpriseAppLifecycleReadRetainsState(t *testing.T) {
	for _, scenario := range []string{"billing-null", "inventory-denied", "wrong-node", "wrong-database", "suspended", "missing-installation", "missing-permission", "wrong-enterprise"} {
		t.Run(scenario, func(t *testing.T) {
			fixture := newEnterpriseLifecycleFixture(t)
			switch scenario {
			case "billing-null":
				fixture.billing = nil
			case "inventory-denied":
				fixture.deny = true
			case "wrong-node":
				fixture.orgNode = "O_wrong"
			case "wrong-database":
				fixture.inventoryID = 43
			case "suspended":
				fixture.suspended = "2026-01-01T00:00:00Z"
			case "missing-installation":
				fixture.missing = true
			case "missing-permission":
				fixture.permission = "read"
			case "wrong-enterprise":
				if err := fixture.data.Set("enterprise_id", "E_wrong"); err != nil {
					t.Fatal(err)
				}
			}
			if err := fixture.lifecycle.read(t.Context(), fixture.data); err == nil {
				t.Fatal("expected read failure")
			}
			if fixture.data.Id() != "O_test" || fixture.data.Get("billing_email") != "old@example.test" || fixture.installPosts != 0 {
				t.Fatal("failure changed identity/billing or installed App")
			}
		})
	}
}

func TestEnterpriseAppLifecycleBlockedOperations(t *testing.T) {
	fixture := newEnterpriseLifecycleFixture(t)
	if !enterpriseOrganizationDeleteContext(t.Context(), fixture.data, fixture.owner).HasError() {
		t.Fatal("delete must be blocked")
	}
	if !enterpriseOrganizationUpdateContext(t.Context(), fixture.data, fixture.owner).HasError() {
		t.Fatal("update must be blocked")
	}
	fixture.data.SetId("another-enterprise/test-org")
	if _, err := enterpriseOrganizationImportContext(t.Context(), fixture.data, fixture.owner); err == nil {
		t.Fatal("cross-enterprise import must fail")
	}
	if fixture.inventoryPages != 0 || fixture.installPosts != 0 {
		t.Fatal("blocked operation made API requests")
	}
}

func TestEnterpriseAppLifecycleBootstrap(t *testing.T) {
	for _, scenario := range []string{"existing", "absent", "suspended", "mismatch", "denied"} {
		t.Run(scenario, func(t *testing.T) {
			fixture := newEnterpriseLifecycleFixture(t)
			fixture.missing = scenario == "absent"
			if scenario == "suspended" {
				fixture.suspended = "2026-01-01T00:00:00Z"
			}
			if scenario == "mismatch" {
				fixture.inventoryID = 43
			}
			if scenario == "denied" {
				fixture.deny = true
			}
			err := fixture.lifecycle.bootstrap(t.Context(), fixture.data)
			failure := scenario == "suspended" || scenario == "mismatch" || scenario == "denied"
			if (err != nil) != failure {
				t.Fatalf("bootstrap error = %v", err)
			}
			expected := 0
			if scenario == "absent" {
				expected = 1
			}
			if fixture.installPosts != expected {
				t.Fatalf("installation POST count = %d, want %d", fixture.installPosts, expected)
			}
		})
	}
}

func TestEnterpriseAppLifecyclePartialCreate(t *testing.T) {
	fixture := newEnterpriseLifecycleFixture(t)
	fixture.suspended = "2026-01-01T00:00:00Z"
	fixture.data.SetId("")
	err := fixture.lifecycle.create(t.Context(), fixture.data, fixture.owner)
	if err == nil || !strings.Contains(err.Error(), "ID retained") {
		t.Fatalf("expected actionable partial-create error: %v", err)
	}
	if fixture.data.Id() != "O_test" || fixture.data.Get("database_id") != 42 || fixture.mutations != 1 || fixture.installPosts != 0 {
		t.Fatal("partial-create identity lost or suspended installation changed")
	}
}

func TestEnterpriseAppLifecycleCreate(t *testing.T) {
	for _, missing := range []bool{false, true} {
		t.Run(fmt.Sprintf("missing-installation-%t", missing), func(t *testing.T) {
			fixture := newEnterpriseLifecycleFixture(t)
			fixture.missing = missing
			fixture.data.SetId("")
			if err := fixture.lifecycle.create(t.Context(), fixture.data, fixture.owner); err != nil {
				t.Fatal(err)
			}
			if fixture.data.Id() != "O_test" || fixture.data.Get("database_id") != 42 || fixture.data.Get("automation_installation_id") != 51 || fixture.mutations != 1 {
				t.Fatal("create did not preserve complete organization/installation identity")
			}
			expectedPosts := 0
			if missing {
				expectedPosts = 1
			}
			if fixture.installPosts != expectedPosts {
				t.Fatalf("installation POSTs = %d, want %d", fixture.installPosts, expectedPosts)
			}
		})
	}
}

func TestEnterpriseAppLifecycleStaleInstallationPermissions(t *testing.T) {
	fixture := newEnterpriseLifecycleFixture(t)
	fixture.lifecycle.automationPermissions["organization_custom_roles"] = "write"
	err := fixture.lifecycle.bootstrap(t.Context(), fixture.data)
	if err == nil || !strings.Contains(err.Error(), "no installation upgrade was attempted") || fixture.installPosts != 0 {
		t.Fatalf("stale grants were not rejected without an upgrade: %v", err)
	}
}

func TestEnterpriseAppLifecyclePermissionsAndCancellation(t *testing.T) {
	if checkAutomationPermissions(map[string]string{"organization_administration": "write", "members": "read", "contents": "read"}) == nil {
		t.Fatal("repository permissions must fail")
	}
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	if propagationWait(ctx, 0) == nil {
		t.Fatal("propagation must respect cancellation")
	}
}

func TestEnterpriseAppLifecycleEnterpriseIdentity(t *testing.T) {
	for _, scenario := range []struct {
		name           string
		accountSlug    string
		enterpriseSlug string
		login          string
		targetType     string
		appID          int
		installationID int
		suspended      any
		issuer         string
		wantError      bool
	}{
		{name: "live-account-slug", accountSlug: "test-enterprise", targetType: "Enterprise", appID: 11, installationID: 12, issuer: "11"},
		{name: "case-insensitive-client-id", accountSlug: "TEST-ENTERPRISE", targetType: "Enterprise", appID: 11, installationID: 12, issuer: "installer-client"},
		{name: "enterprise-slug", enterpriseSlug: "test-enterprise", targetType: "Enterprise", appID: 11, installationID: 12, issuer: "11"},
		{name: "legacy-login", login: "test-enterprise", targetType: "Enterprise", appID: 11, installationID: 12, issuer: "11"},
		{name: "wrong-slug", accountSlug: "other-enterprise", targetType: "Enterprise", appID: 11, installationID: 12, issuer: "11", wantError: true},
		{name: "conflicting-legacy-identity", accountSlug: "other-enterprise", enterpriseSlug: "test-enterprise", login: "test-enterprise", targetType: "Enterprise", appID: 11, installationID: 12, issuer: "11", wantError: true},
		{name: "missing-slug", targetType: "Enterprise", appID: 11, installationID: 12, issuer: "11", wantError: true},
		{name: "wrong-target", accountSlug: "test-enterprise", targetType: "Organization", appID: 11, installationID: 12, issuer: "11", wantError: true},
		{name: "wrong-app", accountSlug: "test-enterprise", targetType: "Enterprise", appID: 22, installationID: 12, issuer: "11", wantError: true},
		{name: "wrong-installation", accountSlug: "test-enterprise", targetType: "Enterprise", appID: 11, installationID: 13, issuer: "11", wantError: true},
		{name: "suspended", accountSlug: "test-enterprise", targetType: "Enterprise", appID: 11, installationID: 12, suspended: "2026-01-01T00:00:00Z", issuer: "11", wantError: true},
		{name: "wrong-issuer", accountSlug: "test-enterprise", targetType: "Enterprise", appID: 11, installationID: 12, issuer: "22", wantError: true},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				writer.Header().Set("Content-Type", "application/json")
				var body any
				switch request.URL.Path {
				case "/app":
					body = map[string]any{"id": 11, "client_id": "installer-client"}
				case "/app/installations/12":
					body = map[string]any{"id": scenario.installationID, "app_id": scenario.appID, "target_type": scenario.targetType, "suspended_at": scenario.suspended,
						"account": map[string]string{"slug": scenario.accountSlug, "login": scenario.login}, "enterprise": map[string]string{"slug": scenario.enterpriseSlug}}
				default:
					t.Errorf("unexpected request: %s %s", request.Method, request.URL)
				}
				if err := json.NewEncoder(writer).Encode(body); err != nil {
					t.Error(err)
				}
			}))
			t.Cleanup(server.Close)
			lifecycle := &enterpriseAppLifecycle{config: &enterpriseAppConfig{slug: "test-enterprise"}}
			err := lifecycle.validateEnterpriseInstallation(t.Context(), mustCreateTestGitHubClient(t, server.URL+"/"), scenario.issuer, 12, "")
			if (err != nil) != scenario.wantError {
				t.Fatalf("identity validation error = %v, want error %t", err, scenario.wantError)
			}
		})
	}
}

func TestEnterpriseAppLifecycleInstallerPermissionBeforeCreate(t *testing.T) {
	for _, permission := range []string{"", "read", "none", "write"} {
		t.Run("grant-"+permission, func(t *testing.T) {
			fixture := newEnterpriseLifecycleFixture(t)
			fixture.installerPermission = permission
			fixture.data.SetId("")
			err := fixture.lifecycle.create(t.Context(), fixture.data, fixture.owner)
			if permission == "write" {
				if err != nil || fixture.mutations != 1 {
					t.Fatalf("write grant blocked creation: %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), "enterprise_organization_installations=write") || fixture.mutations != 0 || fixture.installPosts != 0 || fixture.data.Id() != "" {
				t.Fatalf("insufficient installer grant did not block before creation: %v", err)
			}
		})
	}
}

func TestEnterpriseAppLifecycleBootstrapListingGuard(t *testing.T) {
	for _, scenario := range []string{"listed", "listed-suspended", "listed-stale-grants", "denied", "not-found", "malformed"} {
		t.Run(scenario, func(t *testing.T) {
			fixture := newEnterpriseLifecycleFixture(t)
			fixture.missing = true
			fixture.listed = true
			switch scenario {
			case "listed-suspended":
				fixture.suspended = "2026-01-01T00:00:00Z"
			case "listed-stale-grants":
				fixture.permission = "read"
			case "denied":
				fixture.listingStatus = http.StatusForbidden
			case "not-found":
				fixture.listingStatus = http.StatusNotFound
			case "malformed":
				fixture.listingMalformed = true
			}
			err := fixture.lifecycle.bootstrap(t.Context(), fixture.data)
			if err == nil || fixture.installPosts != 0 || fixture.profileWrites != 0 || fixture.data.Id() != "O_test" {
				t.Fatalf("listing guard did not block duplicate installation without losing state: %v", err)
			}
			if strings.HasPrefix(scenario, "listed") && (fixture.listingPages != 2 || !strings.Contains(err.Error(), "no installation POST or upgrade was attempted")) {
				t.Fatalf("paginated client_id match not detected: pages=%d, error=%v", fixture.listingPages, err)
			}
		})
	}
}
