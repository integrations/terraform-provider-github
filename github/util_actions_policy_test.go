package github

import (
	"encoding/json"
	"testing"

	"github.com/google/go-github/v93/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestExpandActionsPolicyRules(t *testing.T) {
	s := actionsPolicyCommonSchema()

	t.Run("events and actors", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, s, map[string]any{
			"name":           "some-policy",
			"enforcement":    "active",
			"allowed_events": []any{"pull_request_target"},
			"allowed_actor":  []any{map[string]any{"id": 5, "type": "RepositoryRole"}},
		})

		got, err := json.Marshal(expandActionsPolicyRules(d))
		if err != nil {
			t.Fatal(err)
		}
		want := `[{"type":"restrict_actions_actors","parameters":{"allowed_actors":[{"id":5,"type":"RepositoryRole"}]}},{"type":"restrict_action_events","parameters":{"allowed_events":["pull_request_target"]}}]`
		if string(got) != want {
			t.Errorf("got  %s\nwant %s", got, want)
		}
	})

	t.Run("no rules serializes as empty list", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, s, map[string]any{
			"name":        "some-policy",
			"enforcement": "disabled",
		})

		got, err := json.Marshal(github.UpdateActionsPolicyRequest{Rules: expandActionsPolicyRules(d)})
		if err != nil {
			t.Fatal(err)
		}
		if want := `{"rules":[]}`; string(got) != want {
			t.Errorf("got  %s\nwant %s", got, want)
		}
	})
}

func TestExpandActionsPolicyWorkflowPath(t *testing.T) {
	s := actionsPolicyCommonSchema()

	t.Run("set", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, s, map[string]any{
			"workflow_path": []any{map[string]any{
				"include": []any{".github/workflows/a.yml"},
				"exclude": []any{},
			}},
		})

		for _, isUpdate := range []bool{false, true} {
			got, err := json.Marshal(expandActionsPolicyWorkflowPath(d, isUpdate))
			if err != nil {
				t.Fatal(err)
			}
			if want := `{"include":[".github/workflows/a.yml"],"exclude":[]}`; string(got) != want {
				t.Errorf("isUpdate=%t: got %s, want %s", isUpdate, got, want)
			}
		}
	})

	t.Run("unset on create is omitted", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, s, map[string]any{})
		if got := expandActionsPolicyWorkflowPath(d, false); got != nil {
			t.Errorf("got %+v, want nil", got)
		}
	})

	t.Run("unset on update targets all workflows", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, s, map[string]any{})
		got := expandActionsPolicyWorkflowPath(d, true)
		if got == nil || len(got.Include) != 1 || got.Include[0] != actionsPolicyAllWorkflows {
			t.Errorf("got %+v, want include [%q]", got, actionsPolicyAllWorkflows)
		}
	})
}

func TestFlattenActionsPolicyCommon(t *testing.T) {
	d := schema.TestResourceDataRaw(t, actionsPolicyCommonSchema(), map[string]any{})

	var policy github.ActionsPolicy
	body := `{"id":42,"name":"some-policy","target":"actions","enforcement":"active","conditions":{"workflow_path":{"include":["~ALL"],"exclude":[]}},"rules":[{"type":"restrict_action_events","parameters":{"allowed_events":["pull_request","pull_request_target"]}},{"type":"restrict_actions_actors","parameters":{"allowed_actors":[{"id":5,"type":"RepositoryRole"}]}}]}`
	if err := json.Unmarshal([]byte(body), &policy); err != nil {
		t.Fatal(err)
	}

	if err := flattenActionsPolicyCommon(d, &policy); err != nil {
		t.Fatal(err)
	}

	if got := d.Get("policy_id").(int); got != 42 {
		t.Errorf("policy_id = %d, want 42", got)
	}
	if got := d.Get("enforcement").(string); got != "active" {
		t.Errorf("enforcement = %q, want active", got)
	}
	if got := d.Get("workflow_path").([]any); len(got) != 0 {
		t.Errorf("workflow_path = %v, want empty for ~ALL", got)
	}
	if got := d.Get("allowed_events").(*schema.Set); got.Len() != 2 || !got.Contains("pull_request_target") {
		t.Errorf("allowed_events = %v", got.List())
	}
	if got := d.Get("allowed_actor").(*schema.Set); got.Len() != 1 {
		t.Errorf("allowed_actor = %v, want 1 actor", got.List())
	}
}

func TestFlattenActionsPolicyWorkflowPath(t *testing.T) {
	got := flattenActionsPolicyWorkflowPath(&github.ActionsPolicyConditions{
		WorkflowPath: &github.ActionsPolicyWorkflowPathConditionParameters{
			Include: []string{".github/workflows/a.yml"},
			Exclude: []string{},
		},
	})
	if len(got) != 1 {
		t.Fatalf("got %v, want 1 element", got)
	}
	if inc := got[0].(map[string]any)["include"].([]any); len(inc) != 1 || inc[0] != ".github/workflows/a.yml" {
		t.Errorf("include = %v", inc)
	}

	if got := flattenActionsPolicyWorkflowPath(nil); got != nil {
		t.Errorf("nil conditions: got %v, want nil", got)
	}
}
