package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DependencyGraphDiff_vulnerabilities struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The advisory_ghsa_id property
    advisory_ghsa_id *string
    // The advisory_summary property
    advisory_summary *string
    // The advisory_url property
    advisory_url *string
    // The severity property
    severity *string
}
// NewDependencyGraphDiff_vulnerabilities instantiates a new DependencyGraphDiff_vulnerabilities and sets the default values.
func NewDependencyGraphDiff_vulnerabilities()(*DependencyGraphDiff_vulnerabilities) {
    m := &DependencyGraphDiff_vulnerabilities{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDependencyGraphDiff_vulnerabilitiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependencyGraphDiff_vulnerabilitiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependencyGraphDiff_vulnerabilities(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DependencyGraphDiff_vulnerabilities) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdvisoryGhsaId gets the advisory_ghsa_id property value. The advisory_ghsa_id property
// returns a *string when successful
func (m *DependencyGraphDiff_vulnerabilities) GetAdvisoryGhsaId()(*string) {
    return m.advisory_ghsa_id
}
// GetAdvisorySummary gets the advisory_summary property value. The advisory_summary property
// returns a *string when successful
func (m *DependencyGraphDiff_vulnerabilities) GetAdvisorySummary()(*string) {
    return m.advisory_summary
}
// GetAdvisoryUrl gets the advisory_url property value. The advisory_url property
// returns a *string when successful
func (m *DependencyGraphDiff_vulnerabilities) GetAdvisoryUrl()(*string) {
    return m.advisory_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependencyGraphDiff_vulnerabilities) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["advisory_ghsa_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvisoryGhsaId(val)
        }
        return nil
    }
    res["advisory_summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvisorySummary(val)
        }
        return nil
    }
    res["advisory_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvisoryUrl(val)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val)
        }
        return nil
    }
    return res
}
// GetSeverity gets the severity property value. The severity property
// returns a *string when successful
func (m *DependencyGraphDiff_vulnerabilities) GetSeverity()(*string) {
    return m.severity
}
// Serialize serializes information the current object
func (m *DependencyGraphDiff_vulnerabilities) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("advisory_ghsa_id", m.GetAdvisoryGhsaId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("advisory_summary", m.GetAdvisorySummary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("advisory_url", m.GetAdvisoryUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("severity", m.GetSeverity())
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
func (m *DependencyGraphDiff_vulnerabilities) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdvisoryGhsaId sets the advisory_ghsa_id property value. The advisory_ghsa_id property
func (m *DependencyGraphDiff_vulnerabilities) SetAdvisoryGhsaId(value *string)() {
    m.advisory_ghsa_id = value
}
// SetAdvisorySummary sets the advisory_summary property value. The advisory_summary property
func (m *DependencyGraphDiff_vulnerabilities) SetAdvisorySummary(value *string)() {
    m.advisory_summary = value
}
// SetAdvisoryUrl sets the advisory_url property value. The advisory_url property
func (m *DependencyGraphDiff_vulnerabilities) SetAdvisoryUrl(value *string)() {
    m.advisory_url = value
}
// SetSeverity sets the severity property value. The severity property
func (m *DependencyGraphDiff_vulnerabilities) SetSeverity(value *string)() {
    m.severity = value
}
type DependencyGraphDiff_vulnerabilitiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvisoryGhsaId()(*string)
    GetAdvisorySummary()(*string)
    GetAdvisoryUrl()(*string)
    GetSeverity()(*string)
    SetAdvisoryGhsaId(value *string)()
    SetAdvisorySummary(value *string)()
    SetAdvisoryUrl(value *string)()
    SetSeverity(value *string)()
}
