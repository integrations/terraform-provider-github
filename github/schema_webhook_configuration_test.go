package github

import (
	"context"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestWebhookWriteOnlySecretSchema(t *testing.T) {
	t.Parallel()

	for name, resource := range map[string]*schema.Resource{
		"repository":   resourceGithubRepositoryWebhook(),
		"organization": resourceGithubOrganizationWebhook(),
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			configuration := resource.Schema["configuration"].Elem.(*schema.Resource).Schema
			secret := configuration["secret"]
			secretWO := configuration["secret_wo"]
			version := configuration["secret_wo_version"]

			if !secretWO.Optional || !secretWO.Sensitive || !secretWO.WriteOnly {
				t.Fatalf("secret_wo must be optional, sensitive, and write-only: %#v", secretWO)
			}
			if secretWO.Computed || secretWO.ForceNew || secretWO.Default != nil || secretWO.DefaultFunc != nil {
				t.Fatalf("secret_wo must not be computed, ForceNew, or defaulted: %#v", secretWO)
			}
			if version.Type != schema.TypeInt || !version.Optional || version.Sensitive || version.WriteOnly || version.ForceNew {
				t.Fatalf("secret_wo_version must be an ordinary optional integer: %#v", version)
			}
			if _, errors := version.ValidateFunc(0, "configuration.0.secret_wo_version"); len(errors) == 0 {
				t.Fatal("secret_wo_version must reject zero")
			}
			if len(secret.ConflictsWith) != 1 || secret.ConflictsWith[0] != "configuration.0.secret_wo" {
				t.Fatalf("legacy secret conflict is not configured correctly: %#v", secret.ConflictsWith)
			}
			if err := schema.InternalMap(resource.Schema).InternalValidate(nil); err != nil {
				t.Fatalf("resource schema failed internal validation: %v", err)
			}
		})
	}
}

func TestPreferWriteOnlyWebhookSecretValidator(t *testing.T) {
	t.Parallel()

	rawConfig := cty.ObjectVal(map[string]cty.Value{
		"configuration": cty.ListVal([]cty.Value{cty.ObjectVal(map[string]cty.Value{
			"secret":    cty.StringVal("legacy-secret"),
			"secret_wo": cty.NullVal(cty.String),
		})}),
	})
	request := schema.ValidateResourceConfigFuncRequest{
		WriteOnlyAttributesAllowed: true,
		RawConfig:                  rawConfig,
	}
	response := new(schema.ValidateResourceConfigFuncResponse)
	preferWriteOnlyWebhookSecretValidator()(context.Background(), request, response)

	if len(response.Diagnostics) != 1 || response.Diagnostics[0].Severity != diag.Warning {
		t.Fatalf("preference diagnostics count = %d, want one warning", len(response.Diagnostics))
	}
	wantPath := cty.GetAttrPath("configuration").IndexInt(0).GetAttr("secret")
	if !response.Diagnostics[0].AttributePath.Equals(wantPath) {
		t.Fatalf("warning path does not identify configuration[0].secret")
	}
}

func TestWebhookWriteOnlySecretConfigurationValidation(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		configuration map[string]any
		wantError     bool
	}{
		"write-only pair": {
			configuration: map[string]any{"url": "https://example.test", "secret_wo": "canary", "secret_wo_version": 1},
		},
		"legacy secret": {
			configuration: map[string]any{"url": "https://example.test", "secret": "legacy"},
		},
		"secret without version": {
			configuration: map[string]any{"url": "https://example.test", "secret_wo": "canary"},
			wantError:     true,
		},
		"version without secret": {
			configuration: map[string]any{"url": "https://example.test", "secret_wo_version": 1},
			wantError:     true,
		},
		"legacy conflict": {
			configuration: map[string]any{"url": "https://example.test", "secret": "legacy", "secret_wo": "canary", "secret_wo_version": 1},
			wantError:     true,
		},
		"zero version": {
			configuration: map[string]any{"url": "https://example.test", "secret_wo": "canary", "secret_wo_version": 0},
			wantError:     true,
		},
		"negative version": {
			configuration: map[string]any{"url": "https://example.test", "secret_wo": "canary", "secret_wo_version": -1},
			wantError:     true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			config := terraform.NewResourceConfigRaw(map[string]any{
				"repository":    "example",
				"events":        []any{"push"},
				"configuration": []any{test.configuration},
			})
			diags := schema.InternalMap(resourceGithubRepositoryWebhook().Schema).Validate(config)
			if diags.HasError() != test.wantError {
				t.Fatalf("validation error = %t, want %t; diagnostic count: %d", diags.HasError(), test.wantError, len(diags))
			}
		})
	}
}
