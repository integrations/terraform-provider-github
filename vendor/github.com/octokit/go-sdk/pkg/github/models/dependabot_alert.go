package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependabotAlert a Dependabot alert.
type DependabotAlert struct {
    // The time that the alert was auto-dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    auto_dismissed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Details for the vulnerable dependency.
    dependency DependabotAlert_dependencyable
    // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    dismissed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    dismissed_by NullableSimpleUserable
    // An optional comment associated with the alert's dismissal.
    dismissed_comment *string
    // The reason that the alert was dismissed.
    dismissed_reason *DependabotAlert_dismissed_reason
    // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    fixed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The GitHub URL of the alert resource.
    html_url *string
    // The security alert number.
    number *int32
    // Details for the GitHub Security Advisory.
    security_advisory DependabotAlertSecurityAdvisoryable
    // Details pertaining to one vulnerable version range for the advisory.
    security_vulnerability DependabotAlertSecurityVulnerabilityable
    // The state of the Dependabot alert.
    state *DependabotAlert_state
    // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The REST API URL of the alert resource.
    url *string
}
// NewDependabotAlert instantiates a new DependabotAlert and sets the default values.
func NewDependabotAlert()(*DependabotAlert) {
    m := &DependabotAlert{
    }
    return m
}
// CreateDependabotAlertFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependabotAlertFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependabotAlert(), nil
}
// GetAutoDismissedAt gets the auto_dismissed_at property value. The time that the alert was auto-dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *DependabotAlert) GetAutoDismissedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.auto_dismissed_at
}
// GetCreatedAt gets the created_at property value. The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *DependabotAlert) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDependency gets the dependency property value. Details for the vulnerable dependency.
// returns a DependabotAlert_dependencyable when successful
func (m *DependabotAlert) GetDependency()(DependabotAlert_dependencyable) {
    return m.dependency
}
// GetDismissedAt gets the dismissed_at property value. The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *DependabotAlert) GetDismissedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dismissed_at
}
// GetDismissedBy gets the dismissed_by property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *DependabotAlert) GetDismissedBy()(NullableSimpleUserable) {
    return m.dismissed_by
}
// GetDismissedComment gets the dismissed_comment property value. An optional comment associated with the alert's dismissal.
// returns a *string when successful
func (m *DependabotAlert) GetDismissedComment()(*string) {
    return m.dismissed_comment
}
// GetDismissedReason gets the dismissed_reason property value. The reason that the alert was dismissed.
// returns a *DependabotAlert_dismissed_reason when successful
func (m *DependabotAlert) GetDismissedReason()(*DependabotAlert_dismissed_reason) {
    return m.dismissed_reason
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependabotAlert) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["auto_dismissed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoDismissedAt(val)
        }
        return nil
    }
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
    res["dependency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDependabotAlert_dependencyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependency(val.(DependabotAlert_dependencyable))
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
        val, err := n.GetEnumValue(ParseDependabotAlert_dismissed_reason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissedReason(val.(*DependabotAlert_dismissed_reason))
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
    res["security_advisory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDependabotAlertSecurityAdvisoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityAdvisory(val.(DependabotAlertSecurityAdvisoryable))
        }
        return nil
    }
    res["security_vulnerability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDependabotAlertSecurityVulnerabilityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityVulnerability(val.(DependabotAlertSecurityVulnerabilityable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDependabotAlert_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*DependabotAlert_state))
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
func (m *DependabotAlert) GetFixedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.fixed_at
}
// GetHtmlUrl gets the html_url property value. The GitHub URL of the alert resource.
// returns a *string when successful
func (m *DependabotAlert) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetNumber gets the number property value. The security alert number.
// returns a *int32 when successful
func (m *DependabotAlert) GetNumber()(*int32) {
    return m.number
}
// GetSecurityAdvisory gets the security_advisory property value. Details for the GitHub Security Advisory.
// returns a DependabotAlertSecurityAdvisoryable when successful
func (m *DependabotAlert) GetSecurityAdvisory()(DependabotAlertSecurityAdvisoryable) {
    return m.security_advisory
}
// GetSecurityVulnerability gets the security_vulnerability property value. Details pertaining to one vulnerable version range for the advisory.
// returns a DependabotAlertSecurityVulnerabilityable when successful
func (m *DependabotAlert) GetSecurityVulnerability()(DependabotAlertSecurityVulnerabilityable) {
    return m.security_vulnerability
}
// GetState gets the state property value. The state of the Dependabot alert.
// returns a *DependabotAlert_state when successful
func (m *DependabotAlert) GetState()(*DependabotAlert_state) {
    return m.state
}
// GetUpdatedAt gets the updated_at property value. The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *DependabotAlert) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The REST API URL of the alert resource.
// returns a *string when successful
func (m *DependabotAlert) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *DependabotAlert) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetAutoDismissedAt sets the auto_dismissed_at property value. The time that the alert was auto-dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *DependabotAlert) SetAutoDismissedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.auto_dismissed_at = value
}
// SetCreatedAt sets the created_at property value. The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *DependabotAlert) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDependency sets the dependency property value. Details for the vulnerable dependency.
func (m *DependabotAlert) SetDependency(value DependabotAlert_dependencyable)() {
    m.dependency = value
}
// SetDismissedAt sets the dismissed_at property value. The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *DependabotAlert) SetDismissedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dismissed_at = value
}
// SetDismissedBy sets the dismissed_by property value. A GitHub user.
func (m *DependabotAlert) SetDismissedBy(value NullableSimpleUserable)() {
    m.dismissed_by = value
}
// SetDismissedComment sets the dismissed_comment property value. An optional comment associated with the alert's dismissal.
func (m *DependabotAlert) SetDismissedComment(value *string)() {
    m.dismissed_comment = value
}
// SetDismissedReason sets the dismissed_reason property value. The reason that the alert was dismissed.
func (m *DependabotAlert) SetDismissedReason(value *DependabotAlert_dismissed_reason)() {
    m.dismissed_reason = value
}
// SetFixedAt sets the fixed_at property value. The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *DependabotAlert) SetFixedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.fixed_at = value
}
// SetHtmlUrl sets the html_url property value. The GitHub URL of the alert resource.
func (m *DependabotAlert) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetNumber sets the number property value. The security alert number.
func (m *DependabotAlert) SetNumber(value *int32)() {
    m.number = value
}
// SetSecurityAdvisory sets the security_advisory property value. Details for the GitHub Security Advisory.
func (m *DependabotAlert) SetSecurityAdvisory(value DependabotAlertSecurityAdvisoryable)() {
    m.security_advisory = value
}
// SetSecurityVulnerability sets the security_vulnerability property value. Details pertaining to one vulnerable version range for the advisory.
func (m *DependabotAlert) SetSecurityVulnerability(value DependabotAlertSecurityVulnerabilityable)() {
    m.security_vulnerability = value
}
// SetState sets the state property value. The state of the Dependabot alert.
func (m *DependabotAlert) SetState(value *DependabotAlert_state)() {
    m.state = value
}
// SetUpdatedAt sets the updated_at property value. The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *DependabotAlert) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The REST API URL of the alert resource.
func (m *DependabotAlert) SetUrl(value *string)() {
    m.url = value
}
type DependabotAlertable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutoDismissedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDependency()(DependabotAlert_dependencyable)
    GetDismissedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDismissedBy()(NullableSimpleUserable)
    GetDismissedComment()(*string)
    GetDismissedReason()(*DependabotAlert_dismissed_reason)
    GetFixedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHtmlUrl()(*string)
    GetNumber()(*int32)
    GetSecurityAdvisory()(DependabotAlertSecurityAdvisoryable)
    GetSecurityVulnerability()(DependabotAlertSecurityVulnerabilityable)
    GetState()(*DependabotAlert_state)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetAutoDismissedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDependency(value DependabotAlert_dependencyable)()
    SetDismissedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDismissedBy(value NullableSimpleUserable)()
    SetDismissedComment(value *string)()
    SetDismissedReason(value *DependabotAlert_dismissed_reason)()
    SetFixedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHtmlUrl(value *string)()
    SetNumber(value *int32)()
    SetSecurityAdvisory(value DependabotAlertSecurityAdvisoryable)()
    SetSecurityVulnerability(value DependabotAlertSecurityVulnerabilityable)()
    SetState(value *DependabotAlert_state)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
