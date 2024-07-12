package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrgMembership_permissions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The can_create_repository property
    can_create_repository *bool
}
// NewOrgMembership_permissions instantiates a new OrgMembership_permissions and sets the default values.
func NewOrgMembership_permissions()(*OrgMembership_permissions) {
    m := &OrgMembership_permissions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrgMembership_permissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrgMembership_permissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrgMembership_permissions(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrgMembership_permissions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCanCreateRepository gets the can_create_repository property value. The can_create_repository property
// returns a *bool when successful
func (m *OrgMembership_permissions) GetCanCreateRepository()(*bool) {
    return m.can_create_repository
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrgMembership_permissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["can_create_repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanCreateRepository(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *OrgMembership_permissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("can_create_repository", m.GetCanCreateRepository())
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
func (m *OrgMembership_permissions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCanCreateRepository sets the can_create_repository property value. The can_create_repository property
func (m *OrgMembership_permissions) SetCanCreateRepository(value *bool)() {
    m.can_create_repository = value
}
type OrgMembership_permissionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCanCreateRepository()(*bool)
    SetCanCreateRepository(value *bool)()
}
