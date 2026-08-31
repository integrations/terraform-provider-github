package provider

import (
	"encoding/json"
	"testing"
)

func TestStatusString(t *testing.T) {
	cases := map[Status]string{
		StatusPass: "pass", StatusFail: "fail", StatusSkip: "skip",
		StatusPanic: "panic", StatusTimeout: "timeout", StatusRunning: "running",
		StatusUnknown: "unknown",
	}
	for s, want := range cases {
		if got := s.String(); got != want {
			t.Errorf("Status(%d).String() = %q, want %q", s, got, want)
		}
	}
}

func TestParseStatusRoundTrip(t *testing.T) {
	for _, s := range []Status{StatusPass, StatusFail, StatusSkip, StatusPanic, StatusTimeout} {
		got, ok := ParseStatus(s.String())
		if !ok || got != s {
			t.Errorf("ParseStatus(%q) = %v,%v want %v,true", s.String(), got, ok, s)
		}
	}
	if _, ok := ParseStatus("bogus"); ok {
		t.Error("ParseStatus(bogus) should be !ok")
	}
}

func TestTestRequirementsJSONOmitsEmptySlices(t *testing.T) {
	data, err := json.Marshal(TestRequirements{})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(data), `{}`; got != want {
		t.Fatalf("JSON = %s, want %s", got, want)
	}
}

func TestSweepOptsExplicitEmptyResourcesIsDistinctFromNil(t *testing.T) {
	var zero SweepOpts
	if zero.Resources != nil {
		t.Fatalf("zero-value Resources = %v, want nil", zero.Resources)
	}

	explicitEmpty := SweepOpts{Resources: []Resource{}}
	if explicitEmpty.Resources == nil {
		t.Fatal("explicit empty Resources should stay non-nil")
	}
	if len(explicitEmpty.Resources) != 0 {
		t.Fatalf("explicit empty Resources len = %d, want 0", len(explicitEmpty.Resources))
	}
}

func TestPreflightReportOK(t *testing.T) {
	ok := PreflightReport{Checks: []Check{{Status: CheckOK}, {Status: CheckWarn}}}
	if !ok.OK() {
		t.Error("warn-only report should be OK")
	}
	bad := PreflightReport{Checks: []Check{{Status: CheckOK}, {Status: CheckFail}}}
	if bad.OK() {
		t.Error("report with a CheckFail must not be OK")
	}
}
