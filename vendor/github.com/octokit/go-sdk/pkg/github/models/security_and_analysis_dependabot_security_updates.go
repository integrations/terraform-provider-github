package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityAndAnalysis_dependabot_security_updates enable or disable Dependabot security updates for the repository.
type SecurityAndAnalysis_dependabot_security_updates struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enablement status of Dependabot security updates for the repository.
    status *SecurityAndAnalysis_dependabot_security_updates_status
}
// NewSecurityAndAnalysis_dependabot_security_updates instantiates a new SecurityAndAnalysis_dependabot_security_updates and sets the default values.
func NewSecurityAndAnalysis_dependabot_security_updates()(*SecurityAndAnalysis_dependabot_security_updates) {
    m := &SecurityAndAnalysis_dependabot_security_updates{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecurityAndAnalysis_dependabot_security_updatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecurityAndAnalysis_dependabot_security_updatesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityAndAnalysis_dependabot_security_updates(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecurityAndAnalysis_dependabot_security_updates) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecurityAndAnalysis_dependabot_security_updates) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityAndAnalysis_dependabot_security_updates_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SecurityAndAnalysis_dependabot_security_updates_status))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The enablement status of Dependabot security updates for the repository.
// returns a *SecurityAndAnalysis_dependabot_security_updates_status when successful
func (m *SecurityAndAnalysis_dependabot_security_updates) GetStatus()(*SecurityAndAnalysis_dependabot_security_updates_status) {
    return m.status
}
// Serialize serializes information the current object
func (m *SecurityAndAnalysis_dependabot_security_updates) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *SecurityAndAnalysis_dependabot_security_updates) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetStatus sets the status property value. The enablement status of Dependabot security updates for the repository.
func (m *SecurityAndAnalysis_dependabot_security_updates) SetStatus(value *SecurityAndAnalysis_dependabot_security_updates_status)() {
    m.status = value
}
type SecurityAndAnalysis_dependabot_security_updatesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetStatus()(*SecurityAndAnalysis_dependabot_security_updates_status)
    SetStatus(value *SecurityAndAnalysis_dependabot_security_updates_status)()
}
