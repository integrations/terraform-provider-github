package models
import (
    "errors"
)
// The notification setting the team has set
type TeamFull_notification_setting int

const (
    NOTIFICATIONS_ENABLED_TEAMFULL_NOTIFICATION_SETTING TeamFull_notification_setting = iota
    NOTIFICATIONS_DISABLED_TEAMFULL_NOTIFICATION_SETTING
)

func (i TeamFull_notification_setting) String() string {
    return []string{"notifications_enabled", "notifications_disabled"}[i]
}
func ParseTeamFull_notification_setting(v string) (any, error) {
    result := NOTIFICATIONS_ENABLED_TEAMFULL_NOTIFICATION_SETTING
    switch v {
        case "notifications_enabled":
            result = NOTIFICATIONS_ENABLED_TEAMFULL_NOTIFICATION_SETTING
        case "notifications_disabled":
            result = NOTIFICATIONS_DISABLED_TEAMFULL_NOTIFICATION_SETTING
        default:
            return 0, errors.New("Unknown TeamFull_notification_setting value: " + v)
    }
    return &result, nil
}
func SerializeTeamFull_notification_setting(values []TeamFull_notification_setting) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TeamFull_notification_setting) isMultiValue() bool {
    return false
}
