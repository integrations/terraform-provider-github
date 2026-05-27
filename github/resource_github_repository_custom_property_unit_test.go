package github

import (
	"strings"
	"testing"

	"github.com/google/go-github/v85/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TestResourceGithubRepositoryCustomPropertyCreate_EmptyValueRejectsCleanly
// pins the panic guard for #3358: when property_value resolves to an empty
// list at apply time (which can happen despite the schema's MinItems=1 when
// the value comes from a dynamic block whose source is empty), the Create
// path used to index propertyValue[0] and crash the provider with
// "index out of range [0] with length 0". The fix returns a clear error
// for the single-value property types instead. This unit test reaches that
// branch without going through the GitHub API by passing a nil v3client
// in the meta — the guard fires before any client call.
func TestResourceGithubRepositoryCustomPropertyCreate_EmptyValueRejectsCleanly(t *testing.T) {
	singleValueTypes := []github.PropertyValueType{
		github.PropertyValueTypeString,
		github.PropertyValueTypeSingleSelect,
		github.PropertyValueTypeURL,
		github.PropertyValueTypeTrueFalse,
	}

	res := resourceGithubRepositoryCustomProperty()

	for _, pt := range singleValueTypes {
		t.Run(string(pt), func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, res.Schema, map[string]any{
				"repository":     "some-repo",
				"property_name":  "test-property",
				"property_type":  string(pt),
				"property_value": []any{},
			})

			meta := &Owner{name: "test-owner"}
			err := resourceGithubRepositoryCustomPropertyCreate(d, meta)

			if err == nil {
				t.Fatalf("expected error for empty property_value with type %q, got nil", pt)
			}
			if !strings.Contains(err.Error(), "property_value") {
				t.Errorf("expected error to mention property_value, got: %v", err)
			}
		})
	}
}
