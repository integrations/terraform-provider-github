package github

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestEnterpriseAppAuthGuard(t *testing.T) {
	provider := NewProvider("test", "test")()
	if err := provider.InternalValidate(); err != nil {
		t.Fatal(err)
	}
	data := schema.TestResourceDataRaw(t, provider.Schema, map[string]any{
		"auth_mode": "app", "token": "test-pat",
		"app_auth": []any{map[string]any{"id": "1", "installation_id": "2", "pem_file": "invalid-test-key"}},
		"enterprise_app_auth": []any{map[string]any{
			"enterprise_slug": "test-enterprise", "installer_app_id": "3", "installer_installation_id": "4", "installer_pem_file": "invalid-test-key",
			"automation_app_id": "5", "automation_client_id": "test-client", "automation_pem_file": "invalid-test-key",
		}},
	})
	_, diagnostics := configureProvider("test", "test")(context.Background(), data)
	if !diagnostics.HasError() || diagnostics[0].Summary != "enterprise_app_auth requires auth_mode = app, complete creation app_auth credentials, and no token attribute or token environment variable" {
		t.Fatalf("expected pre-network App-only guard, got %v", diagnostics)
	}
}

func TestEnterpriseAppConfigureWithoutOwnerDiscovery(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	key := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requests++
		writer.Header().Set("Content-Type", "application/json")
		var identity struct {
			Issuer string `json:"iss"`
		}
		parts := strings.Split(strings.TrimPrefix(request.Header.Get("Authorization"), "Bearer "), ".")
		if len(parts) == 3 {
			payload, err := base64.RawURLEncoding.DecodeString(parts[1])
			if err != nil {
				t.Error(err)
			}
			if err := json.Unmarshal(payload, &identity); err != nil {
				t.Error(err)
			}
		}
		switch request.URL.Path {
		case "/app":
			if identity.Issuer == "21" {
				fmt.Fprint(writer, `{"id":21,"client_id":"automation-client","permissions":{"organization_administration":"write","members":"read"}}`)
			} else if identity.Issuer == "installer-client" {
				fmt.Fprint(writer, `{"id":11,"client_id":"installer-client"}`)
			} else if identity.Issuer == "1" {
				fmt.Fprint(writer, `{"id":1,"client_id":"creation-client"}`)
			} else {
				t.Errorf("unexpected JWT issuer %q", identity.Issuer)
				writer.WriteHeader(http.StatusForbidden)
			}
		case "/app/installations/2":
			fmt.Fprint(writer, `{"id":2,"app_id":1,"target_type":"Enterprise","enterprise":{"slug":"test-enterprise"}}`)
		case "/app/installations/12":
			fmt.Fprint(writer, `{"id":12,"app_id":11,"target_type":"Enterprise","account":{"slug":"test-enterprise"},"permissions":{"enterprise_organization_installations":"write"}}`)
		case "/app/installations/2/access_tokens":
			fmt.Fprint(writer, `{"token":"mock-creation-token","expires_at":"2099-01-01T00:00:00Z"}`)
		case "/graphql":
			if request.Header.Get("Authorization") != "Bearer mock-creation-token" {
				t.Errorf("wrong enterprise auth")
			}
			fmt.Fprint(writer, `{"data":{"enterprise":{"id":"E_test"}}}`)
		default:
			t.Errorf("unexpected request %s %s", request.Method, request.URL.Path)
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}))
	defer server.Close()
	baseURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	owner, err := configureProviderMeta(t.Context(), "test", &Config{
		BaseURL: baseURL, Owner: "test-enterprise", AppID: new("1"), AppInstallationID: new("2"), AppPEM: key,
		EnterpriseApp: &enterpriseAppConfig{slug: "test-enterprise", installerID: "installer-client", installerInstallationID: "12", installerPEM: string(key), automationID: "21", automationClientID: "automation-client", automationPEM: string(key)},
	})
	if err != nil {
		t.Fatal(err)
	}
	if owner.enterpriseApp == nil || owner.enterpriseApp.enterpriseID != "E_test" || owner.IsOrganization || requests != 7 {
		t.Fatalf("wrong enterprise metadata or request count %d", requests)
	}
}

func TestEnterpriseAppAuthEnvironmentDefaults(t *testing.T) {
	for name, value := range map[string]string{
		"INSTALLER_APP_CLIENT_ID": "installer-client", "INSTALLER_APP_INSTALLATION_ID": "12", "INSTALLER_APP_PRIVATE_KEY": "installer-test-key",
		"ORG_APP_ID": "21", "ORG_APP_CLIENT_ID": "automation-client", "ORG_APP_PRIVATE_KEY": "automation-test-key",
	} {
		t.Setenv(name, value)
	}
	data := schema.TestResourceDataRaw(t, NewProvider("test", "test")().Schema, map[string]any{"enterprise_app_auth": []any{map[string]any{"enterprise_slug": "test-enterprise"}}})
	config, err := parseEnterpriseAppAuth(data)
	if err != nil {
		t.Fatal(err)
	}
	if config.installerID != "installer-client" || config.installerInstallationID != "12" || config.automationID != "21" || config.automationClientID != "automation-client" {
		t.Fatal("environment defaults not loaded")
	}
}

func TestEnterpriseAppDefaultOwnerLookupStillFails(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requests++
		if request.URL.Path != "/users/test-enterprise" {
			t.Errorf("unexpected default-mode request: %s", request.URL.Path)
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprint(writer, `{"message":"Not Found"}`)
	}))
	defer server.Close()
	baseURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	_, err = configureProviderMeta(t.Context(), "test", &Config{BaseURL: baseURL, Owner: "test-enterprise", Token: "mock-pat"})
	if err == nil || !strings.Contains(err.Error(), "failed to lookup owner") || requests != 1 {
		t.Fatalf("default owner 404 was bypassed: requests=%d err=%v", requests, err)
	}
}

func TestEnterpriseAppTokenEnvironmentGuard(t *testing.T) {
	t.Setenv("ENTERPRISE_TEST_PAT", "mock-pat")
	provider := NewProvider("test", "test")()
	data := schema.TestResourceDataRaw(t, provider.Schema, map[string]any{
		"auth_mode": "app", "token_env_name": "ENTERPRISE_TEST_PAT",
		"app_auth": []any{map[string]any{"id": "1", "installation_id": "2", "pem_file": "invalid-test-key"}},
		"enterprise_app_auth": []any{map[string]any{
			"enterprise_slug": "test-enterprise", "installer_app_id": "3", "installer_installation_id": "4", "installer_pem_file": "invalid-test-key",
			"automation_app_id": "5", "automation_client_id": "test-client", "automation_pem_file": "invalid-test-key",
		}},
	})
	_, diagnostics := configureProvider("test", "test")(t.Context(), data)
	if !diagnostics.HasError() || !strings.Contains(diagnostics[0].Summary, "no token attribute or token environment") {
		t.Fatalf("token environment was not rejected before network access: %v", diagnostics)
	}
}
