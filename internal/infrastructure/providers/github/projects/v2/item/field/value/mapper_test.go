package value

import (
	"testing"

	application "github.com/integrations/terraform-provider-github/v6/internal/application/projects/item/field/value"
)

func TestResultFromNodePreservesZero(t *testing.T) {
	value := node{Typename: "ProjectV2ItemFieldNumberValue"}
	result, err := resultFromNode(value)
	if err != nil {
		t.Fatalf("mapping field value: %v", err)
	}
	if result.Kind != application.KindNumber || result.Number != 0 {
		t.Fatalf("unexpected field value: %#v", result)
	}
}
