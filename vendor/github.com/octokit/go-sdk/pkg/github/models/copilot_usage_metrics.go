package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopilotUsageMetrics summary of Copilot usage.
type CopilotUsageMetrics struct {
    // Breakdown of Copilot code completions usage by language and editor
    breakdown []CopilotUsageMetrics_breakdownable
    // The date for which the usage metrics are reported, in `YYYY-MM-DD` format.
    day *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The total number of Copilot code completion suggestions accepted by users.
    total_acceptances_count *int32
    // The total number of users who interacted with Copilot Chat in the IDE during the day specified.
    total_active_chat_users *int32
    // The total number of users who were shown Copilot code completion suggestions during the day specified.
    total_active_users *int32
    // The total instances of users who accepted code suggested by Copilot Chat in the IDE (panel and inline).
    total_chat_acceptances *int32
    // The total number of chat turns (prompt and response pairs) sent between users and Copilot Chat in the IDE.
    total_chat_turns *int32
    // The total number of lines of code completions accepted by users.
    total_lines_accepted *int32
    // The total number of lines of code completions suggested by Copilot.
    total_lines_suggested *int32
    // The total number of Copilot code completion suggestions shown to users.
    total_suggestions_count *int32
}
// NewCopilotUsageMetrics instantiates a new CopilotUsageMetrics and sets the default values.
func NewCopilotUsageMetrics()(*CopilotUsageMetrics) {
    m := &CopilotUsageMetrics{
    }
    return m
}
// CreateCopilotUsageMetricsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotUsageMetricsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotUsageMetrics(), nil
}
// GetBreakdown gets the breakdown property value. Breakdown of Copilot code completions usage by language and editor
// returns a []CopilotUsageMetrics_breakdownable when successful
func (m *CopilotUsageMetrics) GetBreakdown()([]CopilotUsageMetrics_breakdownable) {
    return m.breakdown
}
// GetDay gets the day property value. The date for which the usage metrics are reported, in `YYYY-MM-DD` format.
// returns a *DateOnly when successful
func (m *CopilotUsageMetrics) GetDay()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.day
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotUsageMetrics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["breakdown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCopilotUsageMetrics_breakdownFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CopilotUsageMetrics_breakdownable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CopilotUsageMetrics_breakdownable)
                }
            }
            m.SetBreakdown(res)
        }
        return nil
    }
    res["day"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDay(val)
        }
        return nil
    }
    res["total_acceptances_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAcceptancesCount(val)
        }
        return nil
    }
    res["total_active_chat_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalActiveChatUsers(val)
        }
        return nil
    }
    res["total_active_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalActiveUsers(val)
        }
        return nil
    }
    res["total_chat_acceptances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalChatAcceptances(val)
        }
        return nil
    }
    res["total_chat_turns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalChatTurns(val)
        }
        return nil
    }
    res["total_lines_accepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalLinesAccepted(val)
        }
        return nil
    }
    res["total_lines_suggested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalLinesSuggested(val)
        }
        return nil
    }
    res["total_suggestions_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalSuggestionsCount(val)
        }
        return nil
    }
    return res
}
// GetTotalAcceptancesCount gets the total_acceptances_count property value. The total number of Copilot code completion suggestions accepted by users.
// returns a *int32 when successful
func (m *CopilotUsageMetrics) GetTotalAcceptancesCount()(*int32) {
    return m.total_acceptances_count
}
// GetTotalActiveChatUsers gets the total_active_chat_users property value. The total number of users who interacted with Copilot Chat in the IDE during the day specified.
// returns a *int32 when successful
func (m *CopilotUsageMetrics) GetTotalActiveChatUsers()(*int32) {
    return m.total_active_chat_users
}
// GetTotalActiveUsers gets the total_active_users property value. The total number of users who were shown Copilot code completion suggestions during the day specified.
// returns a *int32 when successful
func (m *CopilotUsageMetrics) GetTotalActiveUsers()(*int32) {
    return m.total_active_users
}
// GetTotalChatAcceptances gets the total_chat_acceptances property value. The total instances of users who accepted code suggested by Copilot Chat in the IDE (panel and inline).
// returns a *int32 when successful
func (m *CopilotUsageMetrics) GetTotalChatAcceptances()(*int32) {
    return m.total_chat_acceptances
}
// GetTotalChatTurns gets the total_chat_turns property value. The total number of chat turns (prompt and response pairs) sent between users and Copilot Chat in the IDE.
// returns a *int32 when successful
func (m *CopilotUsageMetrics) GetTotalChatTurns()(*int32) {
    return m.total_chat_turns
}
// GetTotalLinesAccepted gets the total_lines_accepted property value. The total number of lines of code completions accepted by users.
// returns a *int32 when successful
func (m *CopilotUsageMetrics) GetTotalLinesAccepted()(*int32) {
    return m.total_lines_accepted
}
// GetTotalLinesSuggested gets the total_lines_suggested property value. The total number of lines of code completions suggested by Copilot.
// returns a *int32 when successful
func (m *CopilotUsageMetrics) GetTotalLinesSuggested()(*int32) {
    return m.total_lines_suggested
}
// GetTotalSuggestionsCount gets the total_suggestions_count property value. The total number of Copilot code completion suggestions shown to users.
// returns a *int32 when successful
func (m *CopilotUsageMetrics) GetTotalSuggestionsCount()(*int32) {
    return m.total_suggestions_count
}
// Serialize serializes information the current object
func (m *CopilotUsageMetrics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBreakdown() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBreakdown()))
        for i, v := range m.GetBreakdown() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("breakdown", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteDateOnlyValue("day", m.GetDay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_acceptances_count", m.GetTotalAcceptancesCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_active_chat_users", m.GetTotalActiveChatUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_active_users", m.GetTotalActiveUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_chat_acceptances", m.GetTotalChatAcceptances())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_chat_turns", m.GetTotalChatTurns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_lines_accepted", m.GetTotalLinesAccepted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_lines_suggested", m.GetTotalLinesSuggested())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_suggestions_count", m.GetTotalSuggestionsCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBreakdown sets the breakdown property value. Breakdown of Copilot code completions usage by language and editor
func (m *CopilotUsageMetrics) SetBreakdown(value []CopilotUsageMetrics_breakdownable)() {
    m.breakdown = value
}
// SetDay sets the day property value. The date for which the usage metrics are reported, in `YYYY-MM-DD` format.
func (m *CopilotUsageMetrics) SetDay(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.day = value
}
// SetTotalAcceptancesCount sets the total_acceptances_count property value. The total number of Copilot code completion suggestions accepted by users.
func (m *CopilotUsageMetrics) SetTotalAcceptancesCount(value *int32)() {
    m.total_acceptances_count = value
}
// SetTotalActiveChatUsers sets the total_active_chat_users property value. The total number of users who interacted with Copilot Chat in the IDE during the day specified.
func (m *CopilotUsageMetrics) SetTotalActiveChatUsers(value *int32)() {
    m.total_active_chat_users = value
}
// SetTotalActiveUsers sets the total_active_users property value. The total number of users who were shown Copilot code completion suggestions during the day specified.
func (m *CopilotUsageMetrics) SetTotalActiveUsers(value *int32)() {
    m.total_active_users = value
}
// SetTotalChatAcceptances sets the total_chat_acceptances property value. The total instances of users who accepted code suggested by Copilot Chat in the IDE (panel and inline).
func (m *CopilotUsageMetrics) SetTotalChatAcceptances(value *int32)() {
    m.total_chat_acceptances = value
}
// SetTotalChatTurns sets the total_chat_turns property value. The total number of chat turns (prompt and response pairs) sent between users and Copilot Chat in the IDE.
func (m *CopilotUsageMetrics) SetTotalChatTurns(value *int32)() {
    m.total_chat_turns = value
}
// SetTotalLinesAccepted sets the total_lines_accepted property value. The total number of lines of code completions accepted by users.
func (m *CopilotUsageMetrics) SetTotalLinesAccepted(value *int32)() {
    m.total_lines_accepted = value
}
// SetTotalLinesSuggested sets the total_lines_suggested property value. The total number of lines of code completions suggested by Copilot.
func (m *CopilotUsageMetrics) SetTotalLinesSuggested(value *int32)() {
    m.total_lines_suggested = value
}
// SetTotalSuggestionsCount sets the total_suggestions_count property value. The total number of Copilot code completion suggestions shown to users.
func (m *CopilotUsageMetrics) SetTotalSuggestionsCount(value *int32)() {
    m.total_suggestions_count = value
}
type CopilotUsageMetricsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBreakdown()([]CopilotUsageMetrics_breakdownable)
    GetDay()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetTotalAcceptancesCount()(*int32)
    GetTotalActiveChatUsers()(*int32)
    GetTotalActiveUsers()(*int32)
    GetTotalChatAcceptances()(*int32)
    GetTotalChatTurns()(*int32)
    GetTotalLinesAccepted()(*int32)
    GetTotalLinesSuggested()(*int32)
    GetTotalSuggestionsCount()(*int32)
    SetBreakdown(value []CopilotUsageMetrics_breakdownable)()
    SetDay(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetTotalAcceptancesCount(value *int32)()
    SetTotalActiveChatUsers(value *int32)()
    SetTotalActiveUsers(value *int32)()
    SetTotalChatAcceptances(value *int32)()
    SetTotalChatTurns(value *int32)()
    SetTotalLinesAccepted(value *int32)()
    SetTotalLinesSuggested(value *int32)()
    SetTotalSuggestionsCount(value *int32)()
}
