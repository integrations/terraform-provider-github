package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestTeamMembersCaseDiff(t *testing.T) {
	for _, role := range []string{"member", "maintainer"} {
		t.Run(role, func(t *testing.T) {
			resource := resourceGithubTeamMembers()
			data := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
				"team_slug": "test-team",
				"team_id":   "42",
				"members":   []any{map[string]any{"username": "Alice", "role": "member"}},
			})
			data.SetId("42")

			rawConfig, err := resource.CoreConfigSchema().CoerceValue(cty.ObjectVal(map[string]cty.Value{
				"team_slug": cty.StringVal("test-team"),
				"members": cty.SetVal([]cty.Value{cty.ObjectVal(map[string]cty.Value{
					"username": cty.StringVal("ALICE"),
					"role":     cty.StringVal(role),
				})}),
			}))
			if err != nil {
				t.Fatal(err)
			}
			state := data.State()
			state.RawConfig = rawConfig
			diff, err := resource.Diff(t.Context(), state, terraform.NewResourceConfigShimmed(rawConfig, resource.CoreConfigSchema()), &Owner{IsOrganization: true})
			if err != nil {
				t.Fatal(err)
			}

			membersChanged := false
			if diff != nil {
				if diff.RequiresNew() {
					t.Fatal("case or role changes must not replace team membership")
				}
				for key := range diff.Attributes {
					membersChanged = membersChanged || strings.HasPrefix(key, "members.")
				}
			}
			if wantChange := role != "member"; membersChanged != wantChange {
				t.Errorf("members changed = %t, want %t for role %q", membersChanged, wantChange, role)
			}
		})
	}
}

func TestNestedUsernameDiff(t *testing.T) {
	for _, scope := range []struct {
		name       string
		resource   *schema.Resource
		field      string
		roleField  string
		roles      [2]string
		configName string
	}{
		{"team members", resourceGithubTeamMembers(), "members", "role", [2]string{"member", "maintainer"}, "team_slug"},
		{"collaborators", resourceGithubRepositoryCollaborators(), "user", "permission", [2]string{"pull", "push"}, "repository"},
	} {
		t.Run(scope.name, func(t *testing.T) {
			for _, change := range []struct {
				name      string
				username  cty.Value
				duplicate bool
			}{
				{"distinct", cty.StringVal("bob"), false},
				{"duplicate", cty.StringVal("alice"), true},
				{"case duplicate", cty.StringVal("ALICE"), true},
				{"unknown", cty.UnknownVal(cty.String), false},
			} {
				t.Run(change.name, func(t *testing.T) {
					rawConfig, err := scope.resource.CoreConfigSchema().CoerceValue(cty.ObjectVal(map[string]cty.Value{
						scope.configName: cty.StringVal("test-resource"),
						scope.field: cty.SetVal([]cty.Value{
							cty.ObjectVal(map[string]cty.Value{"username": cty.StringVal("alice"), scope.roleField: cty.StringVal(scope.roles[0])}),
							cty.ObjectVal(map[string]cty.Value{"username": change.username, scope.roleField: cty.StringVal(scope.roles[1])}),
						}),
					}))
					if err != nil {
						t.Fatal(err)
					}
					_, err = scope.resource.Diff(t.Context(), &terraform.InstanceState{RawConfig: rawConfig}, terraform.NewResourceConfigShimmed(rawConfig, scope.resource.CoreConfigSchema()), &Owner{IsOrganization: true})
					if change.duplicate {
						want := fmt.Sprintf("duplicate user 'alice' found in '%s' collection", scope.field)
						if err == nil || !strings.Contains(err.Error(), want) {
							t.Errorf("duplicate validation error = %v, want %q", err, want)
						}
					} else if err != nil {
						t.Fatal(err)
					}
				})
			}
		})
	}
}

func TestRunnerGroupNetworkingDiff(t *testing.T) {
	for _, tc := range []struct {
		name     string
		resource *schema.Resource
	}{
		{"organization", resourceGithubActionsRunnerGroup()},
		{"enterprise", resourceGithubActionsEnterpriseRunnerGroup()},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, change := range []struct {
				name    string
				oldID   string
				newID   cty.Value
				omitted bool
				wantNew *string
			}{
				{name: "retained", oldID: "network-1", newID: cty.StringVal("network-1")},
				{name: "changed", oldID: "network-1", newID: cty.StringVal("network-2"), wantNew: new("network-2")},
				{name: "assigned", newID: cty.StringVal("network-1"), wantNew: new("network-1")},
				{name: "omitted", oldID: "network-1", omitted: true},
				{name: "null", oldID: "network-1", newID: cty.NullVal(cty.String)},
				{name: "unknown", oldID: "network-1", newID: cty.UnknownVal(cty.String)},
				{name: "unassigned", omitted: true},
				{name: "cleared", oldID: "network-1", newID: cty.StringVal(""), wantNew: new("")},
				{name: "already cleared", newID: cty.StringVal("")},
			} {
				t.Run(change.name, func(t *testing.T) {
					attributes := map[string]string{
						"name":                     "before",
						"visibility":               "all",
						"network_configuration_id": change.oldID,
						"etag":                     "previous-etag",
					}
					config := map[string]cty.Value{
						"name":       cty.StringVal("after"),
						"visibility": cty.StringVal("all"),
					}
					if !change.omitted {
						config["network_configuration_id"] = change.newID
					}
					if tc.name == "enterprise" {
						attributes["enterprise_slug"] = "test-enterprise"
						config["enterprise_slug"] = cty.StringVal("test-enterprise")
					}

					rawConfig, err := tc.resource.CoreConfigSchema().CoerceValue(cty.ObjectVal(config))
					if err != nil {
						t.Fatal(err)
					}
					diff, err := tc.resource.Diff(t.Context(), &terraform.InstanceState{
						ID:         "42",
						Attributes: attributes,
						RawConfig:  rawConfig,
					}, terraform.NewResourceConfigShimmed(rawConfig, tc.resource.CoreConfigSchema()), nil)
					if err != nil {
						t.Fatal(err)
					}
					if diff == nil || diff.Empty() {
						t.Fatal("expected a diff")
					}
					if diff.RequiresNew() {
						t.Error("network assignment changes must not replace the runner group")
					}
					if change.wantNew != nil {
						networkDiff := diff.Attributes["network_configuration_id"]
						if networkDiff == nil {
							t.Fatalf("expected network assignment diff to %q", *change.wantNew)
						}
						if networkDiff.New != *change.wantNew || networkDiff.NewComputed != (*change.wantNew == "") || networkDiff.NewRemoved {
							t.Errorf("network assignment diff = %+v, want value %q", networkDiff, *change.wantNew)
						}
					}
					if change.omitted || change.newID.IsNull() {
						if networkDiff, ok := diff.Attributes["network_configuration_id"]; ok {
							t.Errorf("unconfigured network assignment must be preserved, got diff %+v", networkDiff)
						}
					}
					if _, ok := diff.Attributes["etag"]; ok {
						t.Error("ETag diff must remain suppressed when networking diff logic is installed")
					}
				})
			}
		})
	}
}
