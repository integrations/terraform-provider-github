package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryCollaboratorPermission repository Collaborator Permission
type RepositoryCollaboratorPermission struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The permission property
    permission *string
    // The role_name property
    role_name *string
    // Collaborator
    user NullableCollaboratorable
}
// NewRepositoryCollaboratorPermission instantiates a new RepositoryCollaboratorPermission and sets the default values.
func NewRepositoryCollaboratorPermission()(*RepositoryCollaboratorPermission) {
    m := &RepositoryCollaboratorPermission{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryCollaboratorPermissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryCollaboratorPermissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryCollaboratorPermission(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryCollaboratorPermission) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryCollaboratorPermission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["permission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermission(val)
        }
        return nil
    }
    res["role_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleName(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableCollaboratorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(NullableCollaboratorable))
        }
        return nil
    }
    return res
}
// GetPermission gets the permission property value. The permission property
// returns a *string when successful
func (m *RepositoryCollaboratorPermission) GetPermission()(*string) {
    return m.permission
}
// GetRoleName gets the role_name property value. The role_name property
// returns a *string when successful
func (m *RepositoryCollaboratorPermission) GetRoleName()(*string) {
    return m.role_name
}
// GetUser gets the user property value. Collaborator
// returns a NullableCollaboratorable when successful
func (m *RepositoryCollaboratorPermission) GetUser()(NullableCollaboratorable) {
    return m.user
}
// Serialize serializes information the current object
func (m *RepositoryCollaboratorPermission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("permission", m.GetPermission())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("role_name", m.GetRoleName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *RepositoryCollaboratorPermission) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPermission sets the permission property value. The permission property
func (m *RepositoryCollaboratorPermission) SetPermission(value *string)() {
    m.permission = value
}
// SetRoleName sets the role_name property value. The role_name property
func (m *RepositoryCollaboratorPermission) SetRoleName(value *string)() {
    m.role_name = value
}
// SetUser sets the user property value. Collaborator
func (m *RepositoryCollaboratorPermission) SetUser(value NullableCollaboratorable)() {
    m.user = value
}
type RepositoryCollaboratorPermissionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPermission()(*string)
    GetRoleName()(*string)
    GetUser()(NullableCollaboratorable)
    SetPermission(value *string)()
    SetRoleName(value *string)()
    SetUser(value NullableCollaboratorable)()
}
