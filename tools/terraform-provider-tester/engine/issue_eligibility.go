package engine

import (
	"encoding/hex"
	"strings"
)

func validFingerprintPair(full, short string) bool {
	const prefix = "sha256:"
	if !strings.HasPrefix(full, prefix) || len(full) != len(prefix)+64 {
		return false
	}
	digest := strings.TrimPrefix(full, prefix)
	if len(short) != 16 || short != digest[:16] {
		return false
	}
	_, err := hex.DecodeString(digest)
	return err == nil
}

func EligibleForIssueFiling(f PersistFailure) bool {
	switch f.Classification {
	case ClassificationReal, ClassificationRealUnstable:
		return validFingerprintPair(f.Fingerprint, f.ShortFingerprint) &&
			f.KnownIssue == 0 && f.IssueAction == ""
	default:
		return false
	}
}
