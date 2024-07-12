package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecurityAndAnalysis struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The advanced_security property
    advanced_security SecurityAndAnalysis_advanced_securityable
    // Enable or disable Dependabot security updates for the repository.
    dependabot_security_updates SecurityAndAnalysis_dependabot_security_updatesable
    // The secret_scanning property
    secret_scanning SecurityAndAnalysis_secret_scanningable
    // The secret_scanning_push_protection property
    secret_scanning_push_protection SecurityAndAnalysis_secret_scanning_push_protectionable
}
// NewSecurityAndAnalysis instantiates a new SecurityAndAnalysis and sets the default values.
func NewSecurityAndAnalysis()(*SecurityAndAnalysis) {
    m := &SecurityAndAnalysis{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecurityAndAnalysisFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecurityAndAnalysisFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityAndAnalysis(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecurityAndAnalysis) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdvancedSecurity gets the advanced_security property value. The advanced_security property
// returns a SecurityAndAnalysis_advanced_securityable when successful
func (m *SecurityAndAnalysis) GetAdvancedSecurity()(SecurityAndAnalysis_advanced_securityable) {
    return m.advanced_security
}
// GetDependabotSecurityUpdates gets the dependabot_security_updates property value. Enable or disable Dependabot security updates for the repository.
// returns a SecurityAndAnalysis_dependabot_security_updatesable when successful
func (m *SecurityAndAnalysis) GetDependabotSecurityUpdates()(SecurityAndAnalysis_dependabot_security_updatesable) {
    return m.dependabot_security_updates
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecurityAndAnalysis) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["advanced_security"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityAndAnalysis_advanced_securityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedSecurity(val.(SecurityAndAnalysis_advanced_securityable))
        }
        return nil
    }
    res["dependabot_security_updates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityAndAnalysis_dependabot_security_updatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependabotSecurityUpdates(val.(SecurityAndAnalysis_dependabot_security_updatesable))
        }
        return nil
    }
    res["secret_scanning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityAndAnalysis_secret_scanningFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanning(val.(SecurityAndAnalysis_secret_scanningable))
        }
        return nil
    }
    res["secret_scanning_push_protection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityAndAnalysis_secret_scanning_push_protectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningPushProtection(val.(SecurityAndAnalysis_secret_scanning_push_protectionable))
        }
        return nil
    }
    return res
}
// GetSecretScanning gets the secret_scanning property value. The secret_scanning property
// returns a SecurityAndAnalysis_secret_scanningable when successful
func (m *SecurityAndAnalysis) GetSecretScanning()(SecurityAndAnalysis_secret_scanningable) {
    return m.secret_scanning
}
// GetSecretScanningPushProtection gets the secret_scanning_push_protection property value. The secret_scanning_push_protection property
// returns a SecurityAndAnalysis_secret_scanning_push_protectionable when successful
func (m *SecurityAndAnalysis) GetSecretScanningPushProtection()(SecurityAndAnalysis_secret_scanning_push_protectionable) {
    return m.secret_scanning_push_protection
}
// Serialize serializes information the current object
func (m *SecurityAndAnalysis) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("advanced_security", m.GetAdvancedSecurity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("dependabot_security_updates", m.GetDependabotSecurityUpdates())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("secret_scanning", m.GetSecretScanning())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("secret_scanning_push_protection", m.GetSecretScanningPushProtection())
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
func (m *SecurityAndAnalysis) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdvancedSecurity sets the advanced_security property value. The advanced_security property
func (m *SecurityAndAnalysis) SetAdvancedSecurity(value SecurityAndAnalysis_advanced_securityable)() {
    m.advanced_security = value
}
// SetDependabotSecurityUpdates sets the dependabot_security_updates property value. Enable or disable Dependabot security updates for the repository.
func (m *SecurityAndAnalysis) SetDependabotSecurityUpdates(value SecurityAndAnalysis_dependabot_security_updatesable)() {
    m.dependabot_security_updates = value
}
// SetSecretScanning sets the secret_scanning property value. The secret_scanning property
func (m *SecurityAndAnalysis) SetSecretScanning(value SecurityAndAnalysis_secret_scanningable)() {
    m.secret_scanning = value
}
// SetSecretScanningPushProtection sets the secret_scanning_push_protection property value. The secret_scanning_push_protection property
func (m *SecurityAndAnalysis) SetSecretScanningPushProtection(value SecurityAndAnalysis_secret_scanning_push_protectionable)() {
    m.secret_scanning_push_protection = value
}
type SecurityAndAnalysisable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedSecurity()(SecurityAndAnalysis_advanced_securityable)
    GetDependabotSecurityUpdates()(SecurityAndAnalysis_dependabot_security_updatesable)
    GetSecretScanning()(SecurityAndAnalysis_secret_scanningable)
    GetSecretScanningPushProtection()(SecurityAndAnalysis_secret_scanning_push_protectionable)
    SetAdvancedSecurity(value SecurityAndAnalysis_advanced_securityable)()
    SetDependabotSecurityUpdates(value SecurityAndAnalysis_dependabot_security_updatesable)()
    SetSecretScanning(value SecurityAndAnalysis_secret_scanningable)()
    SetSecretScanningPushProtection(value SecurityAndAnalysis_secret_scanning_push_protectionable)()
}
