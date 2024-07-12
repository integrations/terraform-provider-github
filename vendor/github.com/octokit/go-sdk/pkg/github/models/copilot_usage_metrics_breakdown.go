package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopilotUsageMetrics_breakdown breakdown of Copilot usage by editor for this language
type CopilotUsageMetrics_breakdown struct {
    // The number of Copilot suggestions accepted by users in the editor specified during the day specified.
    acceptances_count *int32
    // The number of users who were shown Copilot completion suggestions in the editor specified during the day specified.
    active_users *int32
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The editor in which Copilot suggestions were shown to users for the specified language.
    editor *string
    // The language in which Copilot suggestions were shown to users in the specified editor.
    language *string
    // The number of lines of code accepted by users in the editor specified during the day specified.
    lines_accepted *int32
    // The number of lines of code suggested by Copilot in the editor specified during the day specified.
    lines_suggested *int32
    // The number of Copilot suggestions shown to users in the editor specified during the day specified.
    suggestions_count *int32
}
// NewCopilotUsageMetrics_breakdown instantiates a new CopilotUsageMetrics_breakdown and sets the default values.
func NewCopilotUsageMetrics_breakdown()(*CopilotUsageMetrics_breakdown) {
    m := &CopilotUsageMetrics_breakdown{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopilotUsageMetrics_breakdownFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotUsageMetrics_breakdownFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotUsageMetrics_breakdown(), nil
}
// GetAcceptancesCount gets the acceptances_count property value. The number of Copilot suggestions accepted by users in the editor specified during the day specified.
// returns a *int32 when successful
func (m *CopilotUsageMetrics_breakdown) GetAcceptancesCount()(*int32) {
    return m.acceptances_count
}
// GetActiveUsers gets the active_users property value. The number of users who were shown Copilot completion suggestions in the editor specified during the day specified.
// returns a *int32 when successful
func (m *CopilotUsageMetrics_breakdown) GetActiveUsers()(*int32) {
    return m.active_users
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CopilotUsageMetrics_breakdown) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEditor gets the editor property value. The editor in which Copilot suggestions were shown to users for the specified language.
// returns a *string when successful
func (m *CopilotUsageMetrics_breakdown) GetEditor()(*string) {
    return m.editor
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotUsageMetrics_breakdown) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acceptances_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptancesCount(val)
        }
        return nil
    }
    res["active_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveUsers(val)
        }
        return nil
    }
    res["editor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEditor(val)
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    res["lines_accepted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinesAccepted(val)
        }
        return nil
    }
    res["lines_suggested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinesSuggested(val)
        }
        return nil
    }
    res["suggestions_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuggestionsCount(val)
        }
        return nil
    }
    return res
}
// GetLanguage gets the language property value. The language in which Copilot suggestions were shown to users in the specified editor.
// returns a *string when successful
func (m *CopilotUsageMetrics_breakdown) GetLanguage()(*string) {
    return m.language
}
// GetLinesAccepted gets the lines_accepted property value. The number of lines of code accepted by users in the editor specified during the day specified.
// returns a *int32 when successful
func (m *CopilotUsageMetrics_breakdown) GetLinesAccepted()(*int32) {
    return m.lines_accepted
}
// GetLinesSuggested gets the lines_suggested property value. The number of lines of code suggested by Copilot in the editor specified during the day specified.
// returns a *int32 when successful
func (m *CopilotUsageMetrics_breakdown) GetLinesSuggested()(*int32) {
    return m.lines_suggested
}
// GetSuggestionsCount gets the suggestions_count property value. The number of Copilot suggestions shown to users in the editor specified during the day specified.
// returns a *int32 when successful
func (m *CopilotUsageMetrics_breakdown) GetSuggestionsCount()(*int32) {
    return m.suggestions_count
}
// Serialize serializes information the current object
func (m *CopilotUsageMetrics_breakdown) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("acceptances_count", m.GetAcceptancesCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("active_users", m.GetActiveUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("editor", m.GetEditor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("lines_accepted", m.GetLinesAccepted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("lines_suggested", m.GetLinesSuggested())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("suggestions_count", m.GetSuggestionsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptancesCount sets the acceptances_count property value. The number of Copilot suggestions accepted by users in the editor specified during the day specified.
func (m *CopilotUsageMetrics_breakdown) SetAcceptancesCount(value *int32)() {
    m.acceptances_count = value
}
// SetActiveUsers sets the active_users property value. The number of users who were shown Copilot completion suggestions in the editor specified during the day specified.
func (m *CopilotUsageMetrics_breakdown) SetActiveUsers(value *int32)() {
    m.active_users = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopilotUsageMetrics_breakdown) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEditor sets the editor property value. The editor in which Copilot suggestions were shown to users for the specified language.
func (m *CopilotUsageMetrics_breakdown) SetEditor(value *string)() {
    m.editor = value
}
// SetLanguage sets the language property value. The language in which Copilot suggestions were shown to users in the specified editor.
func (m *CopilotUsageMetrics_breakdown) SetLanguage(value *string)() {
    m.language = value
}
// SetLinesAccepted sets the lines_accepted property value. The number of lines of code accepted by users in the editor specified during the day specified.
func (m *CopilotUsageMetrics_breakdown) SetLinesAccepted(value *int32)() {
    m.lines_accepted = value
}
// SetLinesSuggested sets the lines_suggested property value. The number of lines of code suggested by Copilot in the editor specified during the day specified.
func (m *CopilotUsageMetrics_breakdown) SetLinesSuggested(value *int32)() {
    m.lines_suggested = value
}
// SetSuggestionsCount sets the suggestions_count property value. The number of Copilot suggestions shown to users in the editor specified during the day specified.
func (m *CopilotUsageMetrics_breakdown) SetSuggestionsCount(value *int32)() {
    m.suggestions_count = value
}
type CopilotUsageMetrics_breakdownable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcceptancesCount()(*int32)
    GetActiveUsers()(*int32)
    GetEditor()(*string)
    GetLanguage()(*string)
    GetLinesAccepted()(*int32)
    GetLinesSuggested()(*int32)
    GetSuggestionsCount()(*int32)
    SetAcceptancesCount(value *int32)()
    SetActiveUsers(value *int32)()
    SetEditor(value *string)()
    SetLanguage(value *string)()
    SetLinesAccepted(value *int32)()
    SetLinesSuggested(value *int32)()
    SetSuggestionsCount(value *int32)()
}
