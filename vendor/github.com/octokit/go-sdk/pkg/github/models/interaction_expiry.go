package models
import (
    "errors"
)
// The duration of the interaction restriction. Default: `one_day`.
type InteractionExpiry int

const (
    ONE_DAY_INTERACTIONEXPIRY InteractionExpiry = iota
    THREE_DAYS_INTERACTIONEXPIRY
    ONE_WEEK_INTERACTIONEXPIRY
    ONE_MONTH_INTERACTIONEXPIRY
    SIX_MONTHS_INTERACTIONEXPIRY
)

func (i InteractionExpiry) String() string {
    return []string{"one_day", "three_days", "one_week", "one_month", "six_months"}[i]
}
func ParseInteractionExpiry(v string) (any, error) {
    result := ONE_DAY_INTERACTIONEXPIRY
    switch v {
        case "one_day":
            result = ONE_DAY_INTERACTIONEXPIRY
        case "three_days":
            result = THREE_DAYS_INTERACTIONEXPIRY
        case "one_week":
            result = ONE_WEEK_INTERACTIONEXPIRY
        case "one_month":
            result = ONE_MONTH_INTERACTIONEXPIRY
        case "six_months":
            result = SIX_MONTHS_INTERACTIONEXPIRY
        default:
            return 0, errors.New("Unknown InteractionExpiry value: " + v)
    }
    return &result, nil
}
func SerializeInteractionExpiry(values []InteractionExpiry) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i InteractionExpiry) isMultiValue() bool {
    return false
}
