package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecurityAndAnalysis_secret_scanning_push_protection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The status property
    status *SecurityAndAnalysis_secret_scanning_push_protection_status
}
// NewSecurityAndAnalysis_secret_scanning_push_protection instantiates a new SecurityAndAnalysis_secret_scanning_push_protection and sets the default values.
func NewSecurityAndAnalysis_secret_scanning_push_protection()(*SecurityAndAnalysis_secret_scanning_push_protection) {
    m := &SecurityAndAnalysis_secret_scanning_push_protection{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecurityAndAnalysis_secret_scanning_push_protectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecurityAndAnalysis_secret_scanning_push_protectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityAndAnalysis_secret_scanning_push_protection(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecurityAndAnalysis_secret_scanning_push_protection) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecurityAndAnalysis_secret_scanning_push_protection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityAndAnalysis_secret_scanning_push_protection_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SecurityAndAnalysis_secret_scanning_push_protection_status))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The status property
// returns a *SecurityAndAnalysis_secret_scanning_push_protection_status when successful
func (m *SecurityAndAnalysis_secret_scanning_push_protection) GetStatus()(*SecurityAndAnalysis_secret_scanning_push_protection_status) {
    return m.status
}
// Serialize serializes information the current object
func (m *SecurityAndAnalysis_secret_scanning_push_protection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SecurityAndAnalysis_secret_scanning_push_protection) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetStatus sets the status property value. The status property
func (m *SecurityAndAnalysis_secret_scanning_push_protection) SetStatus(value *SecurityAndAnalysis_secret_scanning_push_protection_status)() {
    m.status = value
}
type SecurityAndAnalysis_secret_scanning_push_protectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetStatus()(*SecurityAndAnalysis_secret_scanning_push_protection_status)
    SetStatus(value *SecurityAndAnalysis_secret_scanning_push_protection_status)()
}
