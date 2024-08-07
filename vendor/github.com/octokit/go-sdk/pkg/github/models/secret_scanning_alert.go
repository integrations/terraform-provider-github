package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecretScanningAlert struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The GitHub URL of the alert resource.
    html_url *string
    // The REST API URL of the code locations for this alert.
    locations_url *string
    // The security alert number.
    number *int32
    // Whether push protection was bypassed for the detected secret.
    push_protection_bypassed *bool
    // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    push_protection_bypassed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    push_protection_bypassed_by NullableSimpleUserable
    // **Required when the `state` is `resolved`.** The reason for resolving the alert.
    resolution *SecretScanningAlertResolution
    // An optional comment to resolve an alert.
    resolution_comment *string
    // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    resolved_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    resolved_by NullableSimpleUserable
    // The secret that was detected.
    secret *string
    // The type of secret that secret scanning detected.
    secret_type *string
    // User-friendly name for the detected secret, matching the `secret_type`.For a list of built-in patterns, see "[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)."
    secret_type_display_name *string
    // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
    state *SecretScanningAlertState
    // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The REST API URL of the alert resource.
    url *string
    // The token status as of the latest validity check.
    validity *SecretScanningAlert_validity
}
// NewSecretScanningAlert instantiates a new SecretScanningAlert and sets the default values.
func NewSecretScanningAlert()(*SecretScanningAlert) {
    m := &SecretScanningAlert{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningAlertFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningAlertFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningAlert(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningAlert) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the created_at property value. The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *SecretScanningAlert) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningAlert) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["locations_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationsUrl(val)
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
    res["push_protection_bypassed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPushProtectionBypassed(val)
        }
        return nil
    }
    res["push_protection_bypassed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPushProtectionBypassedAt(val)
        }
        return nil
    }
    res["push_protection_bypassed_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPushProtectionBypassedBy(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["resolution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecretScanningAlertResolution)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolution(val.(*SecretScanningAlertResolution))
        }
        return nil
    }
    res["resolution_comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolutionComment(val)
        }
        return nil
    }
    res["resolved_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolvedAt(val)
        }
        return nil
    }
    res["resolved_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolvedBy(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["secret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecret(val)
        }
        return nil
    }
    res["secret_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretType(val)
        }
        return nil
    }
    res["secret_type_display_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretTypeDisplayName(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecretScanningAlertState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*SecretScanningAlertState))
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
    res["validity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecretScanningAlert_validity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidity(val.(*SecretScanningAlert_validity))
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The GitHub URL of the alert resource.
// returns a *string when successful
func (m *SecretScanningAlert) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetLocationsUrl gets the locations_url property value. The REST API URL of the code locations for this alert.
// returns a *string when successful
func (m *SecretScanningAlert) GetLocationsUrl()(*string) {
    return m.locations_url
}
// GetNumber gets the number property value. The security alert number.
// returns a *int32 when successful
func (m *SecretScanningAlert) GetNumber()(*int32) {
    return m.number
}
// GetPushProtectionBypassed gets the push_protection_bypassed property value. Whether push protection was bypassed for the detected secret.
// returns a *bool when successful
func (m *SecretScanningAlert) GetPushProtectionBypassed()(*bool) {
    return m.push_protection_bypassed
}
// GetPushProtectionBypassedAt gets the push_protection_bypassed_at property value. The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *SecretScanningAlert) GetPushProtectionBypassedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.push_protection_bypassed_at
}
// GetPushProtectionBypassedBy gets the push_protection_bypassed_by property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *SecretScanningAlert) GetPushProtectionBypassedBy()(NullableSimpleUserable) {
    return m.push_protection_bypassed_by
}
// GetResolution gets the resolution property value. **Required when the `state` is `resolved`.** The reason for resolving the alert.
// returns a *SecretScanningAlertResolution when successful
func (m *SecretScanningAlert) GetResolution()(*SecretScanningAlertResolution) {
    return m.resolution
}
// GetResolutionComment gets the resolution_comment property value. An optional comment to resolve an alert.
// returns a *string when successful
func (m *SecretScanningAlert) GetResolutionComment()(*string) {
    return m.resolution_comment
}
// GetResolvedAt gets the resolved_at property value. The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *SecretScanningAlert) GetResolvedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.resolved_at
}
// GetResolvedBy gets the resolved_by property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *SecretScanningAlert) GetResolvedBy()(NullableSimpleUserable) {
    return m.resolved_by
}
// GetSecret gets the secret property value. The secret that was detected.
// returns a *string when successful
func (m *SecretScanningAlert) GetSecret()(*string) {
    return m.secret
}
// GetSecretType gets the secret_type property value. The type of secret that secret scanning detected.
// returns a *string when successful
func (m *SecretScanningAlert) GetSecretType()(*string) {
    return m.secret_type
}
// GetSecretTypeDisplayName gets the secret_type_display_name property value. User-friendly name for the detected secret, matching the `secret_type`.For a list of built-in patterns, see "[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)."
// returns a *string when successful
func (m *SecretScanningAlert) GetSecretTypeDisplayName()(*string) {
    return m.secret_type_display_name
}
// GetState gets the state property value. Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
// returns a *SecretScanningAlertState when successful
func (m *SecretScanningAlert) GetState()(*SecretScanningAlertState) {
    return m.state
}
// GetUpdatedAt gets the updated_at property value. The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *SecretScanningAlert) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The REST API URL of the alert resource.
// returns a *string when successful
func (m *SecretScanningAlert) GetUrl()(*string) {
    return m.url
}
// GetValidity gets the validity property value. The token status as of the latest validity check.
// returns a *SecretScanningAlert_validity when successful
func (m *SecretScanningAlert) GetValidity()(*SecretScanningAlert_validity) {
    return m.validity
}
// Serialize serializes information the current object
func (m *SecretScanningAlert) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("locations_url", m.GetLocationsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("push_protection_bypassed", m.GetPushProtectionBypassed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("push_protection_bypassed_at", m.GetPushProtectionBypassedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("push_protection_bypassed_by", m.GetPushProtectionBypassedBy())
        if err != nil {
            return err
        }
    }
    if m.GetResolution() != nil {
        cast := (*m.GetResolution()).String()
        err := writer.WriteStringValue("resolution", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resolution_comment", m.GetResolutionComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("resolved_at", m.GetResolvedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resolved_by", m.GetResolvedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secret", m.GetSecret())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secret_type", m.GetSecretType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secret_type_display_name", m.GetSecretTypeDisplayName())
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
    if m.GetValidity() != nil {
        cast := (*m.GetValidity()).String()
        err := writer.WriteStringValue("validity", &cast)
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
func (m *SecretScanningAlert) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the created_at property value. The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *SecretScanningAlert) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetHtmlUrl sets the html_url property value. The GitHub URL of the alert resource.
func (m *SecretScanningAlert) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetLocationsUrl sets the locations_url property value. The REST API URL of the code locations for this alert.
func (m *SecretScanningAlert) SetLocationsUrl(value *string)() {
    m.locations_url = value
}
// SetNumber sets the number property value. The security alert number.
func (m *SecretScanningAlert) SetNumber(value *int32)() {
    m.number = value
}
// SetPushProtectionBypassed sets the push_protection_bypassed property value. Whether push protection was bypassed for the detected secret.
func (m *SecretScanningAlert) SetPushProtectionBypassed(value *bool)() {
    m.push_protection_bypassed = value
}
// SetPushProtectionBypassedAt sets the push_protection_bypassed_at property value. The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *SecretScanningAlert) SetPushProtectionBypassedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.push_protection_bypassed_at = value
}
// SetPushProtectionBypassedBy sets the push_protection_bypassed_by property value. A GitHub user.
func (m *SecretScanningAlert) SetPushProtectionBypassedBy(value NullableSimpleUserable)() {
    m.push_protection_bypassed_by = value
}
// SetResolution sets the resolution property value. **Required when the `state` is `resolved`.** The reason for resolving the alert.
func (m *SecretScanningAlert) SetResolution(value *SecretScanningAlertResolution)() {
    m.resolution = value
}
// SetResolutionComment sets the resolution_comment property value. An optional comment to resolve an alert.
func (m *SecretScanningAlert) SetResolutionComment(value *string)() {
    m.resolution_comment = value
}
// SetResolvedAt sets the resolved_at property value. The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *SecretScanningAlert) SetResolvedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.resolved_at = value
}
// SetResolvedBy sets the resolved_by property value. A GitHub user.
func (m *SecretScanningAlert) SetResolvedBy(value NullableSimpleUserable)() {
    m.resolved_by = value
}
// SetSecret sets the secret property value. The secret that was detected.
func (m *SecretScanningAlert) SetSecret(value *string)() {
    m.secret = value
}
// SetSecretType sets the secret_type property value. The type of secret that secret scanning detected.
func (m *SecretScanningAlert) SetSecretType(value *string)() {
    m.secret_type = value
}
// SetSecretTypeDisplayName sets the secret_type_display_name property value. User-friendly name for the detected secret, matching the `secret_type`.For a list of built-in patterns, see "[Secret scanning patterns](https://docs.github.com/code-security/secret-scanning/secret-scanning-patterns#supported-secrets-for-advanced-security)."
func (m *SecretScanningAlert) SetSecretTypeDisplayName(value *string)() {
    m.secret_type_display_name = value
}
// SetState sets the state property value. Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
func (m *SecretScanningAlert) SetState(value *SecretScanningAlertState)() {
    m.state = value
}
// SetUpdatedAt sets the updated_at property value. The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *SecretScanningAlert) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The REST API URL of the alert resource.
func (m *SecretScanningAlert) SetUrl(value *string)() {
    m.url = value
}
// SetValidity sets the validity property value. The token status as of the latest validity check.
func (m *SecretScanningAlert) SetValidity(value *SecretScanningAlert_validity)() {
    m.validity = value
}
type SecretScanningAlertable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHtmlUrl()(*string)
    GetLocationsUrl()(*string)
    GetNumber()(*int32)
    GetPushProtectionBypassed()(*bool)
    GetPushProtectionBypassedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPushProtectionBypassedBy()(NullableSimpleUserable)
    GetResolution()(*SecretScanningAlertResolution)
    GetResolutionComment()(*string)
    GetResolvedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetResolvedBy()(NullableSimpleUserable)
    GetSecret()(*string)
    GetSecretType()(*string)
    GetSecretTypeDisplayName()(*string)
    GetState()(*SecretScanningAlertState)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetValidity()(*SecretScanningAlert_validity)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHtmlUrl(value *string)()
    SetLocationsUrl(value *string)()
    SetNumber(value *int32)()
    SetPushProtectionBypassed(value *bool)()
    SetPushProtectionBypassedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPushProtectionBypassedBy(value NullableSimpleUserable)()
    SetResolution(value *SecretScanningAlertResolution)()
    SetResolutionComment(value *string)()
    SetResolvedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetResolvedBy(value NullableSimpleUserable)()
    SetSecret(value *string)()
    SetSecretType(value *string)()
    SetSecretTypeDisplayName(value *string)()
    SetState(value *SecretScanningAlertState)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetValidity(value *SecretScanningAlert_validity)()
}
