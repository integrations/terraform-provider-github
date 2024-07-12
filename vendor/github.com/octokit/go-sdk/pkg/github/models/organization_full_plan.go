package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrganizationFull_plan struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The filled_seats property
    filled_seats *int32
    // The name property
    name *string
    // The private_repos property
    private_repos *int32
    // The seats property
    seats *int32
    // The space property
    space *int32
}
// NewOrganizationFull_plan instantiates a new OrganizationFull_plan and sets the default values.
func NewOrganizationFull_plan()(*OrganizationFull_plan) {
    m := &OrganizationFull_plan{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrganizationFull_planFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationFull_planFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationFull_plan(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrganizationFull_plan) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationFull_plan) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["filled_seats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilledSeats(val)
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
    res["seats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeats(val)
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
// GetFilledSeats gets the filled_seats property value. The filled_seats property
// returns a *int32 when successful
func (m *OrganizationFull_plan) GetFilledSeats()(*int32) {
    return m.filled_seats
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *OrganizationFull_plan) GetName()(*string) {
    return m.name
}
// GetPrivateRepos gets the private_repos property value. The private_repos property
// returns a *int32 when successful
func (m *OrganizationFull_plan) GetPrivateRepos()(*int32) {
    return m.private_repos
}
// GetSeats gets the seats property value. The seats property
// returns a *int32 when successful
func (m *OrganizationFull_plan) GetSeats()(*int32) {
    return m.seats
}
// GetSpace gets the space property value. The space property
// returns a *int32 when successful
func (m *OrganizationFull_plan) GetSpace()(*int32) {
    return m.space
}
// Serialize serializes information the current object
func (m *OrganizationFull_plan) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("filled_seats", m.GetFilledSeats())
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
        err := writer.WriteInt32Value("seats", m.GetSeats())
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
func (m *OrganizationFull_plan) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFilledSeats sets the filled_seats property value. The filled_seats property
func (m *OrganizationFull_plan) SetFilledSeats(value *int32)() {
    m.filled_seats = value
}
// SetName sets the name property value. The name property
func (m *OrganizationFull_plan) SetName(value *string)() {
    m.name = value
}
// SetPrivateRepos sets the private_repos property value. The private_repos property
func (m *OrganizationFull_plan) SetPrivateRepos(value *int32)() {
    m.private_repos = value
}
// SetSeats sets the seats property value. The seats property
func (m *OrganizationFull_plan) SetSeats(value *int32)() {
    m.seats = value
}
// SetSpace sets the space property value. The space property
func (m *OrganizationFull_plan) SetSpace(value *int32)() {
    m.space = value
}
type OrganizationFull_planable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFilledSeats()(*int32)
    GetName()(*string)
    GetPrivateRepos()(*int32)
    GetSeats()(*int32)
    GetSpace()(*int32)
    SetFilledSeats(value *int32)()
    SetName(value *string)()
    SetPrivateRepos(value *int32)()
    SetSeats(value *int32)()
    SetSpace(value *int32)()
}
