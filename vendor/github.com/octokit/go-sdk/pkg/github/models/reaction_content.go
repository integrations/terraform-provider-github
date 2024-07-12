package models
import (
    "errors"
)
// The reaction to use
type Reaction_content int

const (
    PLUS_1_REACTION_CONTENT Reaction_content = iota
    MINUS_1_REACTION_CONTENT
    LAUGH_REACTION_CONTENT
    CONFUSED_REACTION_CONTENT
    HEART_REACTION_CONTENT
    HOORAY_REACTION_CONTENT
    ROCKET_REACTION_CONTENT
    EYES_REACTION_CONTENT
)

func (i Reaction_content) String() string {
    return []string{"+1", "-1", "laugh", "confused", "heart", "hooray", "rocket", "eyes"}[i]
}
func ParseReaction_content(v string) (any, error) {
    result := PLUS_1_REACTION_CONTENT
    switch v {
        case "+1":
            result = PLUS_1_REACTION_CONTENT
        case "-1":
            result = MINUS_1_REACTION_CONTENT
        case "laugh":
            result = LAUGH_REACTION_CONTENT
        case "confused":
            result = CONFUSED_REACTION_CONTENT
        case "heart":
            result = HEART_REACTION_CONTENT
        case "hooray":
            result = HOORAY_REACTION_CONTENT
        case "rocket":
            result = ROCKET_REACTION_CONTENT
        case "eyes":
            result = EYES_REACTION_CONTENT
        default:
            return 0, errors.New("Unknown Reaction_content value: " + v)
    }
    return &result, nil
}
func SerializeReaction_content(values []Reaction_content) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Reaction_content) isMultiValue() bool {
    return false
}
