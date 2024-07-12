package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeScanningOrganizationAlertItems struct {
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
    // A GitHub repository.
    repository SimpleRepositoryable
    // The rule property
    rule CodeScanningAlertRuleSummaryable
    // State of a code scanning alert.
    state *CodeScanningAlertState
    // The tool property
    tool CodeScanningAnalysisToolable
    // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The REST API URL of the alert resource.
    url *string
}
// NewCodeScanningOrganizationAlertItems instantiates a new CodeScanningOrganizationAlertItems and sets the default values.
func NewCodeScanningOrganizationAlertItems()(*CodeScanningOrganizationAlertItems) {
    m := &CodeScanningOrganizationAlertItems{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningOrganizationAlertItemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningOrganizationAlertItemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningOrganizationAlertItems(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningOrganizationAlertItems) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *CodeScanningOrganizationAlertItems) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDismissedAt gets the dismissed_at property value. The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *CodeScanningOrganizationAlertItems) GetDismissedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dismissed_at
}
// GetDismissedBy gets the dismissed_by property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *CodeScanningOrganizationAlertItems) GetDismissedBy()(NullableSimpleUserable) {
    return m.dismissed_by
}
// GetDismissedComment gets the dismissed_comment property value. The dismissal comment associated with the dismissal of the alert.
// returns a *string when successful
func (m *CodeScanningOrganizationAlertItems) GetDismissedComment()(*string) {
    return m.dismissed_comment
}
// GetDismissedReason gets the dismissed_reason property value. **Required when the state is dismissed.** The reason for dismissing or closing the alert.
// returns a *CodeScanningAlertDismissedReason when successful
func (m *CodeScanningOrganizationAlertItems) GetDismissedReason()(*CodeScanningAlertDismissedReason) {
    return m.dismissed_reason
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningOrganizationAlertItems) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(SimpleRepositoryable))
        }
        return nil
    }
    res["rule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeScanningAlertRuleSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRule(val.(CodeScanningAlertRuleSummaryable))
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
func (m *CodeScanningOrganizationAlertItems) GetFixedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.fixed_at
}
// GetHtmlUrl gets the html_url property value. The GitHub URL of the alert resource.
// returns a *string when successful
func (m *CodeScanningOrganizationAlertItems) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetInstancesUrl gets the instances_url property value. The REST API URL for fetching the list of instances for an alert.
// returns a *string when successful
func (m *CodeScanningOrganizationAlertItems) GetInstancesUrl()(*string) {
    return m.instances_url
}
// GetMostRecentInstance gets the most_recent_instance property value. The most_recent_instance property
// returns a CodeScanningAlertInstanceable when successful
func (m *CodeScanningOrganizationAlertItems) GetMostRecentInstance()(CodeScanningAlertInstanceable) {
    return m.most_recent_instance
}
// GetNumber gets the number property value. The security alert number.
// returns a *int32 when successful
func (m *CodeScanningOrganizationAlertItems) GetNumber()(*int32) {
    return m.number
}
// GetRepository gets the repository property value. A GitHub repository.
// returns a SimpleRepositoryable when successful
func (m *CodeScanningOrganizationAlertItems) GetRepository()(SimpleRepositoryable) {
    return m.repository
}
// GetRule gets the rule property value. The rule property
// returns a CodeScanningAlertRuleSummaryable when successful
func (m *CodeScanningOrganizationAlertItems) GetRule()(CodeScanningAlertRuleSummaryable) {
    return m.rule
}
// GetState gets the state property value. State of a code scanning alert.
// returns a *CodeScanningAlertState when successful
func (m *CodeScanningOrganizationAlertItems) GetState()(*CodeScanningAlertState) {
    return m.state
}
// GetTool gets the tool property value. The tool property
// returns a CodeScanningAnalysisToolable when successful
func (m *CodeScanningOrganizationAlertItems) GetTool()(CodeScanningAnalysisToolable) {
    return m.tool
}
// GetUpdatedAt gets the updated_at property value. The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *CodeScanningOrganizationAlertItems) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The REST API URL of the alert resource.
// returns a *string when successful
func (m *CodeScanningOrganizationAlertItems) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *CodeScanningOrganizationAlertItems) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("repository", m.GetRepository())
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
func (m *CodeScanningOrganizationAlertItems) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *CodeScanningOrganizationAlertItems) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDismissedAt sets the dismissed_at property value. The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *CodeScanningOrganizationAlertItems) SetDismissedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dismissed_at = value
}
// SetDismissedBy sets the dismissed_by property value. A GitHub user.
func (m *CodeScanningOrganizationAlertItems) SetDismissedBy(value NullableSimpleUserable)() {
    m.dismissed_by = value
}
// SetDismissedComment sets the dismissed_comment property value. The dismissal comment associated with the dismissal of the alert.
func (m *CodeScanningOrganizationAlertItems) SetDismissedComment(value *string)() {
    m.dismissed_comment = value
}
// SetDismissedReason sets the dismissed_reason property value. **Required when the state is dismissed.** The reason for dismissing or closing the alert.
func (m *CodeScanningOrganizationAlertItems) SetDismissedReason(value *CodeScanningAlertDismissedReason)() {
    m.dismissed_reason = value
}
// SetFixedAt sets the fixed_at property value. The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *CodeScanningOrganizationAlertItems) SetFixedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.fixed_at = value
}
// SetHtmlUrl sets the html_url property value. The GitHub URL of the alert resource.
func (m *CodeScanningOrganizationAlertItems) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetInstancesUrl sets the instances_url property value. The REST API URL for fetching the list of instances for an alert.
func (m *CodeScanningOrganizationAlertItems) SetInstancesUrl(value *string)() {
    m.instances_url = value
}
// SetMostRecentInstance sets the most_recent_instance property value. The most_recent_instance property
func (m *CodeScanningOrganizationAlertItems) SetMostRecentInstance(value CodeScanningAlertInstanceable)() {
    m.most_recent_instance = value
}
// SetNumber sets the number property value. The security alert number.
func (m *CodeScanningOrganizationAlertItems) SetNumber(value *int32)() {
    m.number = value
}
// SetRepository sets the repository property value. A GitHub repository.
func (m *CodeScanningOrganizationAlertItems) SetRepository(value SimpleRepositoryable)() {
    m.repository = value
}
// SetRule sets the rule property value. The rule property
func (m *CodeScanningOrganizationAlertItems) SetRule(value CodeScanningAlertRuleSummaryable)() {
    m.rule = value
}
// SetState sets the state property value. State of a code scanning alert.
func (m *CodeScanningOrganizationAlertItems) SetState(value *CodeScanningAlertState)() {
    m.state = value
}
// SetTool sets the tool property value. The tool property
func (m *CodeScanningOrganizationAlertItems) SetTool(value CodeScanningAnalysisToolable)() {
    m.tool = value
}
// SetUpdatedAt sets the updated_at property value. The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *CodeScanningOrganizationAlertItems) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The REST API URL of the alert resource.
func (m *CodeScanningOrganizationAlertItems) SetUrl(value *string)() {
    m.url = value
}
type CodeScanningOrganizationAlertItemsable interface {
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
    GetRepository()(SimpleRepositoryable)
    GetRule()(CodeScanningAlertRuleSummaryable)
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
    SetRepository(value SimpleRepositoryable)()
    SetRule(value CodeScanningAlertRuleSummaryable)()
    SetState(value *CodeScanningAlertState)()
    SetTool(value CodeScanningAnalysisToolable)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
