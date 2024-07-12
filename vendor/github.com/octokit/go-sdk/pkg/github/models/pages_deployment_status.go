package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PagesDeploymentStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The current status of the deployment.
    status *PagesDeploymentStatus_status
}
// NewPagesDeploymentStatus instantiates a new PagesDeploymentStatus and sets the default values.
func NewPagesDeploymentStatus()(*PagesDeploymentStatus) {
    m := &PagesDeploymentStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePagesDeploymentStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePagesDeploymentStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPagesDeploymentStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PagesDeploymentStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PagesDeploymentStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePagesDeploymentStatus_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*PagesDeploymentStatus_status))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The current status of the deployment.
// returns a *PagesDeploymentStatus_status when successful
func (m *PagesDeploymentStatus) GetStatus()(*PagesDeploymentStatus_status) {
    return m.status
}
// Serialize serializes information the current object
func (m *PagesDeploymentStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PagesDeploymentStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetStatus sets the status property value. The current status of the deployment.
func (m *PagesDeploymentStatus) SetStatus(value *PagesDeploymentStatus_status)() {
    m.status = value
}
type PagesDeploymentStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetStatus()(*PagesDeploymentStatus_status)
    SetStatus(value *PagesDeploymentStatus_status)()
}
