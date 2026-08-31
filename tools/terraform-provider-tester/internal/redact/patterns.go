package redact

import "regexp"

var (
	githubTokenPattern = regexp.MustCompile(`\b(?:gh[pousr]_[A-Za-z0-9_]{20,}|github_pat_[A-Za-z0-9_]{20,}_[A-Za-z0-9_]{20,})\b`)
	bearerJWTPattern   = regexp.MustCompile(`(?i)\bBearer\s+[A-Za-z0-9_-]{10,}\.[A-Za-z0-9_-]{10,}\.[A-Za-z0-9_-]{10,}\b`)
	authHeaderPattern  = regexp.MustCompile(`(?i)Authorization:\s*(?:token|bearer)\s+\S+`)
	pemBlockPattern    = regexp.MustCompile(`(?s)-----BEGIN [A-Z0-9 ]+-----\r?\n.*?\r?\n-----END [A-Z0-9 ]+-----`)
)

func redactPatterns(in string) string {
	out := pemBlockPattern.ReplaceAllString(in, marker)
	out = authHeaderPattern.ReplaceAllString(out, marker)
	out = bearerJWTPattern.ReplaceAllString(out, marker)
	out = githubTokenPattern.ReplaceAllString(out, marker)
	return out
}
