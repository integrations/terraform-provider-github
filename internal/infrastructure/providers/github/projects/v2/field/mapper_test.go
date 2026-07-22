package field

import (
	"testing"
	"time"
)

func TestResultFromNodeMapsIterationConfiguration(t *testing.T) {
	value := node{Typename: "ProjectV2IterationField"}
	value.Iteration.ID = "PVTF_1"
	value.Iteration.Configuration.Duration = 14
	value.Iteration.Configuration.Iterations = []iterationValue{{ID: "I_2", StartDate: "2026-07-13", Duration: 14}}
	result, err := resultFromNode(value)
	if err != nil {
		t.Fatalf("mapping field: %v", err)
	}
	if result.IterationConfiguration == nil || result.IterationConfiguration.Iterations[0].StartDate.Format(time.DateOnly) != "2026-07-13" {
		t.Fatalf("unexpected field mapping: %#v", result)
	}
}
