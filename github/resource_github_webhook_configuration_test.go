package github

import (
	"encoding/json"
	"maps"
	"testing"

	gh "github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestWebhookObjectsOmitWriteOnlySecret(t *testing.T) {
	t.Parallel()

	for name, resource := range map[string]struct {
		schema map[string]*schema.Schema
		object func(*schema.ResourceData) *gh.Hook
		extra  map[string]any
	}{
		"repository": {
			schema: resourceGithubRepositoryWebhook().Schema,
			object: resourceGithubRepositoryWebhookObject,
			extra:  map[string]any{"repository": "example"},
		},
		"organization": {
			schema: resourceGithubOrganizationWebhook().Schema,
			object: resourceGithubOrganizationWebhookObject,
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			raw := map[string]any{
				"events": []any{"push"},
				"configuration": []any{map[string]any{
					"url":               "https://example.test/hook",
					"secret_wo_version": 1,
				}},
			}
			maps.Copy(raw, resource.extra)
			d := schema.TestResourceDataRaw(t, resource.schema, raw)
			hook := resource.object(d)
			if hook.Config != nil && hook.Config.Secret != nil {
				t.Fatal("ordinary webhook expansion included a secret")
			}

			payload, err := json.Marshal(hook)
			if err != nil {
				t.Fatalf("marshalling webhook request: %v", err)
			}
			var sanitized struct {
				Config map[string]json.RawMessage `json:"config"`
			}
			if err := json.Unmarshal(payload, &sanitized); err != nil {
				t.Fatalf("parsing webhook request: %v", err)
			}
			if _, exists := sanitized.Config["secret"]; exists {
				t.Fatal("serialized webhook request included a secret field")
			}
		})
	}
}

func TestInterfaceFromWebhookConfigPreservingState(t *testing.T) {
	t.Parallel()

	masked := "********"
	remote := &gh.HookConfig{
		URL:         new("https://example.test/hook"),
		ContentType: new("json"),
		InsecureSSL: new("0"),
		Secret:      &masked,
	}

	for name, test := range map[string]struct {
		configuration map[string]any
		wantSecret    string
		wantVersion   int
	}{
		"legacy secret": {
			configuration: map[string]any{"url": "https://example.test/hook", "secret": "legacy-secret"},
			wantSecret:    "legacy-secret",
		},
		"write-only secret": {
			configuration: map[string]any{"url": "https://example.test/hook", "secret_wo_version": 7},
			wantVersion:   7,
		},
		"import": {
			configuration: map[string]any{"url": "https://example.test/hook"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			d := schema.TestResourceDataRaw(t, resourceGithubRepositoryWebhook().Schema, map[string]any{
				"repository":    "example",
				"events":        []any{"push"},
				"configuration": []any{test.configuration},
			})
			flattened := interfaceFromWebhookConfigPreservingState(remote, d)[0].(map[string]any)
			if got, ok := flattened["secret"]; test.wantSecret == "" && ok {
				t.Fatalf("masked remote secret was flattened into state: value length %d", len(got.(string)))
			} else if test.wantSecret != "" && got != test.wantSecret {
				t.Fatalf("legacy secret was not preserved")
			}
			if got, ok := flattened["secret_wo_version"]; test.wantVersion == 0 && ok {
				t.Fatalf("unexpected write-only version in state: %v", got)
			} else if test.wantVersion != 0 && got != test.wantVersion {
				t.Fatalf("write-only version = %v, want %d", got, test.wantVersion)
			}
			if _, ok := flattened["secret_wo"]; ok {
				t.Fatal("write-only secret was flattened into state")
			}
		})
	}
}

func TestInterfaceFromNilWebhookConfig(t *testing.T) {
	t.Parallel()

	flattened := interfaceFromWebhookConfig(nil)
	if len(flattened) != 1 || flattened[0] == nil {
		t.Fatalf("nil webhook configuration was not handled safely: %#v", flattened)
	}
}
