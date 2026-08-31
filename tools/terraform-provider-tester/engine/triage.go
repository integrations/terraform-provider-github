package engine

import "strings"

const (
	ClassificationFlakeConfirmed  = "flake-confirmed"
	ClassificationFlakeHistorical = "flake-historical"
	ClassificationReal            = "real"
	ClassificationRealUnstable    = "real-unstable"
)

// TriageAttempt records one in-memory attempt without raw output.
type TriageAttempt struct {
	Number    int
	Status    string
	Signature Signature
}

// TriageInput contains the safe data needed to classify one failure.
type TriageInput struct {
	Mode            string
	Signature       Signature
	Attempts        []TriageAttempt
	PreviousHistory []string
	FailN           int
	LastM           int
	LogPath         string
}

func ClassifyFailure(in TriageInput) PersistFailure {
	if in.FailN == 0 {
		in.FailN = 1
	}
	if in.LastM == 0 {
		in.LastM = 10
	}
	attempts := in.Attempts
	if len(attempts) == 0 {
		attempts = []TriageAttempt{{Number: 1, Status: in.Signature.Status, Signature: in.Signature}}
	}

	classification := ClassificationReal
	reasons := []string{}
	selected := in.Signature
	if selected.Fingerprint == "" {
		selected = firstAttemptSignature(attempts)
	}

	if passedAfterFailure(attempts) {
		classification = ClassificationFlakeConfirmed
		reasons = append(reasons, "later retry passed")
		selected = firstFailedSignature(attempts, selected)
	} else if failedFingerprintsDiffer(attempts) {
		classification = ClassificationRealUnstable
		reasons = append(reasons, "failed attempts had different fingerprints")
		selected = lastFailedSignature(attempts, selected)
	} else {
		history := append([]string{}, in.PreviousHistory...)
		if len(attempts) > 0 {
			history = append(history, attempts[0].Status)
		}
		if IsInfraRetryableClass(selected.Class) && Flaky(history, in.FailN, in.LastM) {
			classification = ClassificationFlakeHistorical
			reasons = append(reasons, "history window has mixed pass and fail")
		} else if !selected.Retryable {
			reasons = append(reasons, "signature is not retryable")
		} else {
			reasons = append(reasons, "all attempts failed")
		}
	}

	return PersistFailure{
		Package: selected.Package, Test: selected.Test, Sub: selected.Sub, Status: selected.Status,
		Fingerprint: selected.Fingerprint, ShortFingerprint: selected.ShortFingerprint,
		Class: selected.Class, Canonical: selected.Canonical, Retryable: selected.Retryable,
		Classification: classification, Attempts: len(attempts), Mode: in.Mode, Reasons: reasons, LogPath: in.LogPath,
	}
}

func passedAfterFailure(attempts []TriageAttempt) bool {
	failed := false
	for _, a := range attempts {
		if isFailureStatusString(a.Status) {
			failed = true
			continue
		}
		if failed && a.Status == "pass" {
			return true
		}
	}
	return false
}

func failedFingerprintsDiffer(attempts []TriageAttempt) bool {
	seen := ""
	for _, a := range attempts {
		if !isFailureStatusString(a.Status) || a.Signature.Fingerprint == "" {
			continue
		}
		if seen == "" {
			seen = a.Signature.Fingerprint
			continue
		}
		if a.Signature.Fingerprint != seen {
			return true
		}
	}
	return false
}

func firstAttemptSignature(attempts []TriageAttempt) Signature {
	for _, a := range attempts {
		if a.Signature.Fingerprint != "" {
			return a.Signature
		}
	}
	return Signature{}
}

func firstFailedSignature(attempts []TriageAttempt, fallback Signature) Signature {
	for _, a := range attempts {
		if isFailureStatusString(a.Status) && a.Signature.Fingerprint != "" {
			return a.Signature
		}
	}
	return fallback
}

func lastFailedSignature(attempts []TriageAttempt, fallback Signature) Signature {
	for i := len(attempts) - 1; i >= 0; i-- {
		if isFailureStatusString(attempts[i].Status) && attempts[i].Signature.Fingerprint != "" {
			return attempts[i].Signature
		}
	}
	return fallback
}

func isFailureStatusString(status string) bool {
	switch strings.ToLower(status) {
	case "fail", "panic", "timeout", "build", "pre-run":
		return true
	default:
		return false
	}
}
