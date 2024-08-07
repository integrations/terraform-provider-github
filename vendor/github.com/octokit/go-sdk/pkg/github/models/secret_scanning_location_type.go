package models
import (
    "errors"
)
// The location type. Because secrets may be found in different types of resources (ie. code, comments, issues, pull requests, discussions), this field identifies the type of resource where the secret was found.
type SecretScanningLocation_type int

const (
    COMMIT_SECRETSCANNINGLOCATION_TYPE SecretScanningLocation_type = iota
    WIKI_COMMIT_SECRETSCANNINGLOCATION_TYPE
    ISSUE_TITLE_SECRETSCANNINGLOCATION_TYPE
    ISSUE_BODY_SECRETSCANNINGLOCATION_TYPE
    ISSUE_COMMENT_SECRETSCANNINGLOCATION_TYPE
    DISCUSSION_TITLE_SECRETSCANNINGLOCATION_TYPE
    DISCUSSION_BODY_SECRETSCANNINGLOCATION_TYPE
    DISCUSSION_COMMENT_SECRETSCANNINGLOCATION_TYPE
    PULL_REQUEST_TITLE_SECRETSCANNINGLOCATION_TYPE
    PULL_REQUEST_BODY_SECRETSCANNINGLOCATION_TYPE
    PULL_REQUEST_COMMENT_SECRETSCANNINGLOCATION_TYPE
    PULL_REQUEST_REVIEW_SECRETSCANNINGLOCATION_TYPE
    PULL_REQUEST_REVIEW_COMMENT_SECRETSCANNINGLOCATION_TYPE
)

func (i SecretScanningLocation_type) String() string {
    return []string{"commit", "wiki_commit", "issue_title", "issue_body", "issue_comment", "discussion_title", "discussion_body", "discussion_comment", "pull_request_title", "pull_request_body", "pull_request_comment", "pull_request_review", "pull_request_review_comment"}[i]
}
func ParseSecretScanningLocation_type(v string) (any, error) {
    result := COMMIT_SECRETSCANNINGLOCATION_TYPE
    switch v {
        case "commit":
            result = COMMIT_SECRETSCANNINGLOCATION_TYPE
        case "wiki_commit":
            result = WIKI_COMMIT_SECRETSCANNINGLOCATION_TYPE
        case "issue_title":
            result = ISSUE_TITLE_SECRETSCANNINGLOCATION_TYPE
        case "issue_body":
            result = ISSUE_BODY_SECRETSCANNINGLOCATION_TYPE
        case "issue_comment":
            result = ISSUE_COMMENT_SECRETSCANNINGLOCATION_TYPE
        case "discussion_title":
            result = DISCUSSION_TITLE_SECRETSCANNINGLOCATION_TYPE
        case "discussion_body":
            result = DISCUSSION_BODY_SECRETSCANNINGLOCATION_TYPE
        case "discussion_comment":
            result = DISCUSSION_COMMENT_SECRETSCANNINGLOCATION_TYPE
        case "pull_request_title":
            result = PULL_REQUEST_TITLE_SECRETSCANNINGLOCATION_TYPE
        case "pull_request_body":
            result = PULL_REQUEST_BODY_SECRETSCANNINGLOCATION_TYPE
        case "pull_request_comment":
            result = PULL_REQUEST_COMMENT_SECRETSCANNINGLOCATION_TYPE
        case "pull_request_review":
            result = PULL_REQUEST_REVIEW_SECRETSCANNINGLOCATION_TYPE
        case "pull_request_review_comment":
            result = PULL_REQUEST_REVIEW_COMMENT_SECRETSCANNINGLOCATION_TYPE
        default:
            return 0, errors.New("Unknown SecretScanningLocation_type value: " + v)
    }
    return &result, nil
}
func SerializeSecretScanningLocation_type(values []SecretScanningLocation_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecretScanningLocation_type) isMultiValue() bool {
    return false
}
