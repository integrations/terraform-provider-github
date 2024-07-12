package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeScanningAlertRuleSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A short description of the rule used to detect the alert.
    description *string
    // A unique identifier for the rule used to detect the alert.
    id *string
    // The name of the rule used to detect the alert.
    name *string
    // The security severity of the alert.
    security_severity_level *CodeScanningAlertRuleSummary_security_severity_level
    // The severity of the alert.
    severity *CodeScanningAlertRuleSummary_severity
    // A set of tags applicable for the rule.
    tags []string
}
// NewCodeScanningAlertRuleSummary instantiates a new CodeScanningAlertRuleSummary and sets the default values.
func NewCodeScanningAlertRuleSummary()(*CodeScanningAlertRuleSummary) {
    m := &CodeScanningAlertRuleSummary{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningAlertRuleSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningAlertRuleSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningAlertRuleSummary(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningAlertRuleSummary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. A short description of the rule used to detect the alert.
// returns a *string when successful
func (m *CodeScanningAlertRuleSummary) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningAlertRuleSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["security_severity_level"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningAlertRuleSummary_security_severity_level)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecuritySeverityLevel(val.(*CodeScanningAlertRuleSummary_security_severity_level))
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningAlertRuleSummary_severity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*CodeScanningAlertRuleSummary_severity))
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetTags(res)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. A unique identifier for the rule used to detect the alert.
// returns a *string when successful
func (m *CodeScanningAlertRuleSummary) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. The name of the rule used to detect the alert.
// returns a *string when successful
func (m *CodeScanningAlertRuleSummary) GetName()(*string) {
    return m.name
}
// GetSecuritySeverityLevel gets the security_severity_level property value. The security severity of the alert.
// returns a *CodeScanningAlertRuleSummary_security_severity_level when successful
func (m *CodeScanningAlertRuleSummary) GetSecuritySeverityLevel()(*CodeScanningAlertRuleSummary_security_severity_level) {
    return m.security_severity_level
}
// GetSeverity gets the severity property value. The severity of the alert.
// returns a *CodeScanningAlertRuleSummary_severity when successful
func (m *CodeScanningAlertRuleSummary) GetSeverity()(*CodeScanningAlertRuleSummary_severity) {
    return m.severity
}
// GetTags gets the tags property value. A set of tags applicable for the rule.
// returns a []string when successful
func (m *CodeScanningAlertRuleSummary) GetTags()([]string) {
    return m.tags
}
// Serialize serializes information the current object
func (m *CodeScanningAlertRuleSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetSecuritySeverityLevel() != nil {
        cast := (*m.GetSecuritySeverityLevel()).String()
        err := writer.WriteStringValue("security_severity_level", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err := writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err := writer.WriteCollectionOfStringValues("tags", m.GetTags())
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
func (m *CodeScanningAlertRuleSummary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. A short description of the rule used to detect the alert.
func (m *CodeScanningAlertRuleSummary) SetDescription(value *string)() {
    m.description = value
}
// SetId sets the id property value. A unique identifier for the rule used to detect the alert.
func (m *CodeScanningAlertRuleSummary) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. The name of the rule used to detect the alert.
func (m *CodeScanningAlertRuleSummary) SetName(value *string)() {
    m.name = value
}
// SetSecuritySeverityLevel sets the security_severity_level property value. The security severity of the alert.
func (m *CodeScanningAlertRuleSummary) SetSecuritySeverityLevel(value *CodeScanningAlertRuleSummary_security_severity_level)() {
    m.security_severity_level = value
}
// SetSeverity sets the severity property value. The severity of the alert.
func (m *CodeScanningAlertRuleSummary) SetSeverity(value *CodeScanningAlertRuleSummary_severity)() {
    m.severity = value
}
// SetTags sets the tags property value. A set of tags applicable for the rule.
func (m *CodeScanningAlertRuleSummary) SetTags(value []string)() {
    m.tags = value
}
type CodeScanningAlertRuleSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetId()(*string)
    GetName()(*string)
    GetSecuritySeverityLevel()(*CodeScanningAlertRuleSummary_security_severity_level)
    GetSeverity()(*CodeScanningAlertRuleSummary_severity)
    GetTags()([]string)
    SetDescription(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetSecuritySeverityLevel(value *CodeScanningAlertRuleSummary_security_severity_level)()
    SetSeverity(value *CodeScanningAlertRuleSummary_severity)()
    SetTags(value []string)()
}
