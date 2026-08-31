package engine

import (
	"strings"
	"testing"
)

func TestEligibleForIssueFiling(t *testing.T) {
	validA := "sha256:" + strings.Repeat("a", 64)

	tests := []struct {
		name string
		in   PersistFailure
		want bool
	}{
		{
			name: "real with valid matching fingerprints",
			in:   PersistFailure{Classification: ClassificationReal, Fingerprint: validA, ShortFingerprint: strings.Repeat("a", 16)},
			want: true,
		},
		{
			name: "real unstable with valid matching fingerprints",
			in:   PersistFailure{Classification: ClassificationRealUnstable, Fingerprint: validA, ShortFingerprint: strings.Repeat("a", 16)},
			want: true,
		},
		{
			name: "flake is ineligible",
			in:   PersistFailure{Classification: ClassificationFlakeConfirmed, Fingerprint: validA, ShortFingerprint: strings.Repeat("a", 16)},
			want: false,
		},
		{
			name: "known issue is ineligible",
			in:   PersistFailure{Classification: ClassificationReal, Fingerprint: validA, ShortFingerprint: strings.Repeat("a", 16), KnownIssue: 42},
			want: false,
		},
		{
			name: "existing issue action is ineligible",
			in:   PersistFailure{Classification: ClassificationReal, Fingerprint: validA, ShortFingerprint: strings.Repeat("a", 16), IssueAction: "dedup"},
			want: false,
		},
		{
			name: "empty fingerprint is ineligible",
			in:   PersistFailure{Classification: ClassificationReal, Fingerprint: "", ShortFingerprint: ""},
			want: false,
		},
		{
			name: "truncated fingerprint is ineligible",
			in:   PersistFailure{Classification: ClassificationReal, Fingerprint: "sha256:" + strings.Repeat("a", 63), ShortFingerprint: strings.Repeat("a", 16)},
			want: false,
		},
		{
			name: "non-hex fingerprint is ineligible",
			in:   PersistFailure{Classification: ClassificationReal, Fingerprint: "sha256:" + strings.Repeat("g", 64), ShortFingerprint: strings.Repeat("g", 16)},
			want: false,
		},
		{
			name: "mismatched short fingerprint is ineligible",
			in:   PersistFailure{Classification: ClassificationReal, Fingerprint: validA, ShortFingerprint: strings.Repeat("b", 16)},
			want: false,
		},
		{
			name: "short fingerprint with wrong length is ineligible",
			in:   PersistFailure{Classification: ClassificationReal, Fingerprint: validA, ShortFingerprint: strings.Repeat("a", 15)},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EligibleForIssueFiling(tt.in); got != tt.want {
				t.Fatalf("EligibleForIssueFiling() = %v, want %v", got, tt.want)
			}
		})
	}
}
