package markdown
import (
    "errors"
)
// The rendering mode.
type MarkdownPostRequestBody_mode int

const (
    MARKDOWN_MARKDOWNPOSTREQUESTBODY_MODE MarkdownPostRequestBody_mode = iota
    GFM_MARKDOWNPOSTREQUESTBODY_MODE
)

func (i MarkdownPostRequestBody_mode) String() string {
    return []string{"markdown", "gfm"}[i]
}
func ParseMarkdownPostRequestBody_mode(v string) (any, error) {
    result := MARKDOWN_MARKDOWNPOSTREQUESTBODY_MODE
    switch v {
        case "markdown":
            result = MARKDOWN_MARKDOWNPOSTREQUESTBODY_MODE
        case "gfm":
            result = GFM_MARKDOWNPOSTREQUESTBODY_MODE
        default:
            return 0, errors.New("Unknown MarkdownPostRequestBody_mode value: " + v)
    }
    return &result, nil
}
func SerializeMarkdownPostRequestBody_mode(values []MarkdownPostRequestBody_mode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MarkdownPostRequestBody_mode) isMultiValue() bool {
    return false
}
