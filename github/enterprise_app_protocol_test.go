package github

import (
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-cty/cty/msgpack"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestEnterpriseAppPartialCreateProtocolState(t *testing.T) {
	fixture := newEnterpriseLifecycleFixture(t)
	fixture.suspended = "2026-01-01T00:00:00Z"
	resource := resourceGithubEnterpriseOrganization()
	provider := &schema.Provider{ResourcesMap: map[string]*schema.Resource{"github_enterprise_organization": resource}}
	provider.SetMeta(fixture.owner)
	server := schema.NewGRPCProviderServer(provider)
	resourceType := resource.CoreConfigSchema().ImpliedType()
	attributes := map[string]cty.Value{}
	for name, attributeType := range resourceType.AttributeTypes() {
		attributes[name] = cty.NullVal(attributeType)
	}
	attributes["enterprise_id"] = cty.StringVal("E_test")
	attributes["name"] = cty.StringVal("test-org")
	attributes["billing_email"] = cty.StringVal("billing@example.test")
	attributes["admin_logins"] = cty.SetVal([]cty.Value{cty.StringVal("owner-one")})
	encode := func(value cty.Value) *tfprotov5.DynamicValue {
		encoded, err := msgpack.Marshal(value, resourceType)
		if err != nil {
			t.Fatal(err)
		}
		return &tfprotov5.DynamicValue{MsgPack: encoded}
	}
	config := encode(cty.ObjectVal(attributes))
	prior := encode(cty.NullVal(resourceType))
	plan, err := server.PlanResourceChange(t.Context(), &tfprotov5.PlanResourceChangeRequest{
		TypeName: "github_enterprise_organization", PriorState: prior, ProposedNewState: config, Config: config,
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, diagnostic := range plan.Diagnostics {
		if diagnostic.Severity == tfprotov5.DiagnosticSeverityError {
			t.Fatal(diagnostic)
		}
	}
	response, err := server.ApplyResourceChange(t.Context(), &tfprotov5.ApplyResourceChangeRequest{
		TypeName: "github_enterprise_organization", PriorState: prior, PlannedState: plan.PlannedState, Config: config, PlannedPrivate: plan.PlannedPrivate,
	})
	if err != nil {
		t.Fatal(err)
	}
	failed := false
	for _, diagnostic := range response.Diagnostics {
		if diagnostic.Severity == tfprotov5.DiagnosticSeverityError && strings.Contains(diagnostic.Summary+diagnostic.Detail, "ID retained") {
			failed = true
		}
	}
	if !failed || response.NewState == nil {
		t.Fatalf("missing recovery error/state: %v", response.Diagnostics)
	}
	state, err := msgpack.Unmarshal(response.NewState.MsgPack, resourceType)
	if err != nil {
		t.Fatal(err)
	}
	if state.GetAttr("id").AsString() != "O_test" || !state.GetAttr("database_id").RawEquals(cty.NumberIntVal(42)) {
		t.Fatalf("SDK discarded partial-create identity: %s", state.GoString())
	}
	if fixture.mutations != 1 || fixture.installPosts != 0 {
		t.Fatal("unexpected create/install count")
	}
}
