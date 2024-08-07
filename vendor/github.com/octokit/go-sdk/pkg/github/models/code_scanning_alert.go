package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeScanningAlert struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    dismissed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    dismissed_by NullableSimpleUserable
    // The dismissal comment associated with the dismissal of the alert.
    dismissed_comment *string
    // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
    dismissed_reason *CodeScanningAlertDismissedReason
    // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    fixed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The GitHub URL of the alert resource.
    html_url *string
    // The REST API URL for fetching the list of instances for an alert.
    instances_url *string
    // The most_recent_instance property
    most_recent_instance CodeScanningAlertInstanceable
    // The security alert number.
    number *int32
    // The rule property
    rule CodeScanningAlertRuleable
    // State of a code scanning alert.
    state *CodeScanningAlertState
    // The tool property
    tool CodeScanningAnalysisToolable
    // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The REST API URL of the alert resource.
    url *string
}
// NewCodeScanningAlert instantiates a new CodeScanningAlert and sets the default values.
func NewCodeScanningAlert()(*CodeScanningAlert) {
    m := &CodeScanningAlert{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningAlertFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningAlertFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningAlert(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningAlert) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *CodeScanningAlert) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDismissedAt gets the dismissed_at property value. The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *CodeScanningAlert) GetDismissedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dismissed_at
}
// GetDismissedBy gets the dismissed_by property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *CodeScanningAlert) GetDismissedBy()(NullableSimpleUserable) {
    return m.dismissed_by
}
// GetDismissedComment gets the dismissed_comment property value. The dismissal comment associated with the dismissal of the alert.
// returns a *string when successful
func (m *CodeScanningAlert) GetDismissedComment()(*string) {
    return m.dismissed_comment
}
// GetDismissedReason gets the dismissed_reason property value. **Required when the state is dismissed.** The reason for dismissing or closing the alert.
// returns a *CodeScanningAlertDismissedReason when successful
func (m *CodeScanningAlert) GetDismissedReason()(*CodeScanningAlertDismissedReason) {
    return m.dismissed_reason
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningAlert) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["dismissed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissedAt(val)
        }
        return nil
    }
    res["dismissed_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissedBy(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["dismissed_comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissedComment(val)
        }
        return nil
    }
    res["dismissed_reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningAlertDismissedReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissedReason(val.(*CodeScanningAlertDismissedReason))
        }
        return nil
    }
    res["fixed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFixedAt(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
        }
        return nil
    }
    res["instances_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstancesUrl(val)
        }
        return nil
    }
    res["most_recent_instance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeScanningAlertInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMostRecentInstance(val.(CodeScanningAlertInstanceable))
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["rule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeScanningAlertRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRule(val.(CodeScanningAlertRuleable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningAlertState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*CodeScanningAlertState))
        }
        return nil
    }
    res["tool"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeScanningAnalysisToolFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTool(val.(CodeScanningAnalysisToolable))
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetFixedAt gets the fixed_at property value. The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *CodeScanningAlert) GetFixedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.fixed_at
}
// GetHtmlUrl gets the html_url property value. The GitHub URL of the alert resource.
// returns a *string when successful
func (m *CodeScanningAlert) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetInstancesUrl gets the instances_url property value. The REST API URL for fetching the list of instances for an alert.
// returns a *string when successful
func (m *CodeScanningAlert) GetInstancesUrl()(*string) {
    return m.instances_url
}
// GetMostRecentInstance gets the most_recent_instance property value. The most_recent_instance property
// returns a CodeScanningAlertInstanceable when successful
func (m *CodeScanningAlert) GetMostRecentInstance()(CodeScanningAlertInstanceable) {
    return m.most_recent_instance
}
// GetNumber gets the number property value. The security alert number.
// returns a *int32 when successful
func (m *CodeScanningAlert) GetNumber()(*int32) {
    return m.number
}
// GetRule gets the rule property value. The rule property
// returns a CodeScanningAlertRuleable when successful
func (m *CodeScanningAlert) GetRule()(CodeScanningAlertRuleable) {
    return m.rule
}
// GetState gets the state property value. State of a code scanning alert.
// returns a *CodeScanningAlertState when successful
func (m *CodeScanningAlert) GetState()(*CodeScanningAlertState) {
    return m.state
}
// GetTool gets the tool property value. The tool property
// returns a CodeScanningAnalysisToolable when successful
func (m *CodeScanningAlert) GetTool()(CodeScanningAnalysisToolable) {
    return m.tool
}
// GetUpdatedAt gets the updated_at property value. The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *CodeScanningAlert) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The REST API URL of the alert resource.
// returns a *string when successful
func (m *CodeScanningAlert) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *CodeScanningAlert) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("dismissed_by", m.GetDismissedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dismissed_comment", m.GetDismissedComment())
        if err != nil {
            return err
        }
    }
    if m.GetDismissedReason() != nil {
        cast := (*m.GetDismissedReason()).String()
        err := writer.WriteStringValue("dismissed_reason", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("most_recent_instance", m.GetMostRecentInstance())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("rule", m.GetRule())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("tool", m.GetTool())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodeScanningAlert) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *CodeScanningAlert) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDismissedAt sets the dismissed_at property value. The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *CodeScanningAlert) SetDismissedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dismissed_at = value
}
// SetDismissedBy sets the dismissed_by property value. A GitHub user.
func (m *CodeScanningAlert) SetDismissedBy(value NullableSimpleUserable)() {
    m.dismissed_by = value
}
// SetDismissedComment sets the dismissed_comment property value. The dismissal comment associated with the dismissal of the alert.
func (m *CodeScanningAlert) SetDismissedComment(value *string)() {
    m.dismissed_comment = value
}
// SetDismissedReason sets the dismissed_reason property value. **Required when the state is dismissed.** The reason for dismissing or closing the alert.
func (m *CodeScanningAlert) SetDismissedReason(value *CodeScanningAlertDismissedReason)() {
    m.dismissed_reason = value
}
// SetFixedAt sets the fixed_at property value. The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *CodeScanningAlert) SetFixedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.fixed_at = value
}
// SetHtmlUrl sets the html_url property value. The GitHub URL of the alert resource.
func (m *CodeScanningAlert) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetInstancesUrl sets the instances_url property value. The REST API URL for fetching the list of instances for an alert.
func (m *CodeScanningAlert) SetInstancesUrl(value *string)() {
    m.instances_url = value
}
// SetMostRecentInstance sets the most_recent_instance property value. The most_recent_instance property
func (m *CodeScanningAlert) SetMostRecentInstance(value CodeScanningAlertInstanceable)() {
    m.most_recent_instance = value
}
// SetNumber sets the number property value. The security alert number.
func (m *CodeScanningAlert) SetNumber(value *int32)() {
    m.number = value
}
// SetRule sets the rule property value. The rule property
func (m *CodeScanningAlert) SetRule(value CodeScanningAlertRuleable)() {
    m.rule = value
}
// SetState sets the state property value. State of a code scanning alert.
func (m *CodeScanningAlert) SetState(value *CodeScanningAlertState)() {
    m.state = value
}
// SetTool sets the tool property value. The tool property
func (m *CodeScanningAlert) SetTool(value CodeScanningAnalysisToolable)() {
    m.tool = value
}
// SetUpdatedAt sets the updated_at property value. The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *CodeScanningAlert) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The REST API URL of the alert resource.
func (m *CodeScanningAlert) SetUrl(value *string)() {
    m.url = value
}
type CodeScanningAlertable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDismissedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDismissedBy()(NullableSimpleUserable)
    GetDismissedComment()(*string)
    GetDismissedReason()(*CodeScanningAlertDismissedReason)
    GetFixedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHtmlUrl()(*string)
    GetInstancesUrl()(*string)
    GetMostRecentInstance()(CodeScanningAlertInstanceable)
    GetNumber()(*int32)
    GetRule()(CodeScanningAlertRuleable)
    GetState()(*CodeScanningAlertState)
    GetTool()(CodeScanningAnalysisToolable)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDismissedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDismissedBy(value NullableSimpleUserable)()
    SetDismissedComment(value *string)()
    SetDismissedReason(value *CodeScanningAlertDismissedReason)()
    SetFixedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHtmlUrl(value *string)()
    SetInstancesUrl(value *string)()
    SetMostRecentInstance(value CodeScanningAlertInstanceable)()
    SetNumber(value *int32)()
    SetRule(value CodeScanningAlertRuleable)()
    SetState(value *CodeScanningAlertState)()
    SetTool(value CodeScanningAnalysisToolable)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
