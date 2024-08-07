package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PublicUser_plan struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The collaborators property
    collaborators *int32
    // The name property
    name *string
    // The private_repos property
    private_repos *int32
    // The space property
    space *int32
}
// NewPublicUser_plan instantiates a new PublicUser_plan and sets the default values.
func NewPublicUser_plan()(*PublicUser_plan) {
    m := &PublicUser_plan{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePublicUser_planFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePublicUser_planFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPublicUser_plan(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PublicUser_plan) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCollaborators gets the collaborators property value. The collaborators property
// returns a *int32 when successful
func (m *PublicUser_plan) GetCollaborators()(*int32) {
    return m.collaborators
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PublicUser_plan) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["collaborators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollaborators(val)
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
    res["private_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateRepos(val)
        }
        return nil
    }
    res["space"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpace(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *PublicUser_plan) GetName()(*string) {
    return m.name
}
// GetPrivateRepos gets the private_repos property value. The private_repos property
// returns a *int32 when successful
func (m *PublicUser_plan) GetPrivateRepos()(*int32) {
    return m.private_repos
}
// GetSpace gets the space property value. The space property
// returns a *int32 when successful
func (m *PublicUser_plan) GetSpace()(*int32) {
    return m.space
}
// Serialize serializes information the current object
func (m *PublicUser_plan) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("collaborators", m.GetCollaborators())
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
    {
        err := writer.WriteInt32Value("private_repos", m.GetPrivateRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("space", m.GetSpace())
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
func (m *PublicUser_plan) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCollaborators sets the collaborators property value. The collaborators property
func (m *PublicUser_plan) SetCollaborators(value *int32)() {
    m.collaborators = value
}
// SetName sets the name property value. The name property
func (m *PublicUser_plan) SetName(value *string)() {
    m.name = value
}
// SetPrivateRepos sets the private_repos property value. The private_repos property
func (m *PublicUser_plan) SetPrivateRepos(value *int32)() {
    m.private_repos = value
}
// SetSpace sets the space property value. The space property
func (m *PublicUser_plan) SetSpace(value *int32)() {
    m.space = value
}
type PublicUser_planable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCollaborators()(*int32)
    GetName()(*string)
    GetPrivateRepos()(*int32)
    GetSpace()(*int32)
    SetCollaborators(value *int32)()
    SetName(value *string)()
    SetPrivateRepos(value *int32)()
    SetSpace(value *int32)()
}
