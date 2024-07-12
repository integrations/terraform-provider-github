package reactions
import (
    "errors"
)
// The [reaction type](https://docs.github.com/rest/reactions/reactions#about-reactions) to add to the team discussion.
type ReactionsPostRequestBody_content int

const (
    PLUS_1_REACTIONSPOSTREQUESTBODY_CONTENT ReactionsPostRequestBody_content = iota
    MINUS_1_REACTIONSPOSTREQUESTBODY_CONTENT
    LAUGH_REACTIONSPOSTREQUESTBODY_CONTENT
    CONFUSED_REACTIONSPOSTREQUESTBODY_CONTENT
    HEART_REACTIONSPOSTREQUESTBODY_CONTENT
    HOORAY_REACTIONSPOSTREQUESTBODY_CONTENT
    ROCKET_REACTIONSPOSTREQUESTBODY_CONTENT
    EYES_REACTIONSPOSTREQUESTBODY_CONTENT
)

func (i ReactionsPostRequestBody_content) String() string {
    return []string{"+1", "-1", "laugh", "confused", "heart", "hooray", "rocket", "eyes"}[i]
}
func ParseReactionsPostRequestBody_content(v string) (any, error) {
    result := PLUS_1_REACTIONSPOSTREQUESTBODY_CONTENT
    switch v {
        case "+1":
            result = PLUS_1_REACTIONSPOSTREQUESTBODY_CONTENT
        case "-1":
            result = MINUS_1_REACTIONSPOSTREQUESTBODY_CONTENT
        case "laugh":
            result = LAUGH_REACTIONSPOSTREQUESTBODY_CONTENT
        case "confused":
            result = CONFUSED_REACTIONSPOSTREQUESTBODY_CONTENT
        case "heart":
            result = HEART_REACTIONSPOSTREQUESTBODY_CONTENT
        case "hooray":
            result = HOORAY_REACTIONSPOSTREQUESTBODY_CONTENT
        case "rocket":
            result = ROCKET_REACTIONSPOSTREQUESTBODY_CONTENT
        case "eyes":
            result = EYES_REACTIONSPOSTREQUESTBODY_CONTENT
        default:
            return 0, errors.New("Unknown ReactionsPostRequestBody_content value: " + v)
    }
    return &result, nil
}
func SerializeReactionsPostRequestBody_content(values []ReactionsPostRequestBody_content) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReactionsPostRequestBody_content) isMultiValue() bool {
    return false
}
