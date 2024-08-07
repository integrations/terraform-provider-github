package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryRuleParamsCodeScanningTool a tool that must provide code scanning results for this rule to pass.
type RepositoryRuleParamsCodeScanningTool struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The severity level at which code scanning results that raise alerts block a reference update. For more information on alert severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
    alerts_threshold *RepositoryRuleParamsCodeScanningTool_alerts_threshold
    // The severity level at which code scanning results that raise security alerts block a reference update. For more information on security severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
    security_alerts_threshold *RepositoryRuleParamsCodeScanningTool_security_alerts_threshold
    // The name of a code scanning tool
    tool *string
}
// NewRepositoryRuleParamsCodeScanningTool instantiates a new RepositoryRuleParamsCodeScanningTool and sets the default values.
func NewRepositoryRuleParamsCodeScanningTool()(*RepositoryRuleParamsCodeScanningTool) {
    m := &RepositoryRuleParamsCodeScanningTool{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleParamsCodeScanningToolFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleParamsCodeScanningToolFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleParamsCodeScanningTool(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleParamsCodeScanningTool) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAlertsThreshold gets the alerts_threshold property value. The severity level at which code scanning results that raise alerts block a reference update. For more information on alert severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
// returns a *RepositoryRuleParamsCodeScanningTool_alerts_threshold when successful
func (m *RepositoryRuleParamsCodeScanningTool) GetAlertsThreshold()(*RepositoryRuleParamsCodeScanningTool_alerts_threshold) {
    return m.alerts_threshold
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleParamsCodeScanningTool) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alerts_threshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleParamsCodeScanningTool_alerts_threshold)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertsThreshold(val.(*RepositoryRuleParamsCodeScanningTool_alerts_threshold))
        }
        return nil
    }
    res["security_alerts_threshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryRuleParamsCodeScanningTool_security_alerts_threshold)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityAlertsThreshold(val.(*RepositoryRuleParamsCodeScanningTool_security_alerts_threshold))
        }
        return nil
    }
    res["tool"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTool(val)
        }
        return nil
    }
    return res
}
// GetSecurityAlertsThreshold gets the security_alerts_threshold property value. The severity level at which code scanning results that raise security alerts block a reference update. For more information on security severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
// returns a *RepositoryRuleParamsCodeScanningTool_security_alerts_threshold when successful
func (m *RepositoryRuleParamsCodeScanningTool) GetSecurityAlertsThreshold()(*RepositoryRuleParamsCodeScanningTool_security_alerts_threshold) {
    return m.security_alerts_threshold
}
// GetTool gets the tool property value. The name of a code scanning tool
// returns a *string when successful
func (m *RepositoryRuleParamsCodeScanningTool) GetTool()(*string) {
    return m.tool
}
// Serialize serializes information the current object
func (m *RepositoryRuleParamsCodeScanningTool) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAlertsThreshold() != nil {
        cast := (*m.GetAlertsThreshold()).String()
        err := writer.WriteStringValue("alerts_threshold", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecurityAlertsThreshold() != nil {
        cast := (*m.GetSecurityAlertsThreshold()).String()
        err := writer.WriteStringValue("security_alerts_threshold", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tool", m.GetTool())
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
func (m *RepositoryRuleParamsCodeScanningTool) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAlertsThreshold sets the alerts_threshold property value. The severity level at which code scanning results that raise alerts block a reference update. For more information on alert severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
func (m *RepositoryRuleParamsCodeScanningTool) SetAlertsThreshold(value *RepositoryRuleParamsCodeScanningTool_alerts_threshold)() {
    m.alerts_threshold = value
}
// SetSecurityAlertsThreshold sets the security_alerts_threshold property value. The severity level at which code scanning results that raise security alerts block a reference update. For more information on security severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
func (m *RepositoryRuleParamsCodeScanningTool) SetSecurityAlertsThreshold(value *RepositoryRuleParamsCodeScanningTool_security_alerts_threshold)() {
    m.security_alerts_threshold = value
}
// SetTool sets the tool property value. The name of a code scanning tool
func (m *RepositoryRuleParamsCodeScanningTool) SetTool(value *string)() {
    m.tool = value
}
type RepositoryRuleParamsCodeScanningToolable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertsThreshold()(*RepositoryRuleParamsCodeScanningTool_alerts_threshold)
    GetSecurityAlertsThreshold()(*RepositoryRuleParamsCodeScanningTool_security_alerts_threshold)
    GetTool()(*string)
    SetAlertsThreshold(value *RepositoryRuleParamsCodeScanningTool_alerts_threshold)()
    SetSecurityAlertsThreshold(value *RepositoryRuleParamsCodeScanningTool_security_alerts_threshold)()
    SetTool(value *string)()
}
