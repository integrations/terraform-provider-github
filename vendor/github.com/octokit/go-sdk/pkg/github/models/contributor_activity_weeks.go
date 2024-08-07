package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContributorActivity_weeks struct {
    // The a property
    a *int32
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The c property
    c *int32
    // The d property
    d *int32
    // The w property
    w *int32
}
// NewContributorActivity_weeks instantiates a new ContributorActivity_weeks and sets the default values.
func NewContributorActivity_weeks()(*ContributorActivity_weeks) {
    m := &ContributorActivity_weeks{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContributorActivity_weeksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContributorActivity_weeksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContributorActivity_weeks(), nil
}
// GetA gets the a property value. The a property
// returns a *int32 when successful
func (m *ContributorActivity_weeks) GetA()(*int32) {
    return m.a
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContributorActivity_weeks) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetC gets the c property value. The c property
// returns a *int32 when successful
func (m *ContributorActivity_weeks) GetC()(*int32) {
    return m.c
}
// GetD gets the d property value. The d property
// returns a *int32 when successful
func (m *ContributorActivity_weeks) GetD()(*int32) {
    return m.d
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContributorActivity_weeks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["a"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetA(val)
        }
        return nil
    }
    res["c"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetC(val)
        }
        return nil
    }
    res["d"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetD(val)
        }
        return nil
    }
    res["w"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetW(val)
        }
        return nil
    }
    return res
}
// GetW gets the w property value. The w property
// returns a *int32 when successful
func (m *ContributorActivity_weeks) GetW()(*int32) {
    return m.w
}
// Serialize serializes information the current object
func (m *ContributorActivity_weeks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("a", m.GetA())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("c", m.GetC())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("d", m.GetD())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("w", m.GetW())
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
// SetA sets the a property value. The a property
func (m *ContributorActivity_weeks) SetA(value *int32)() {
    m.a = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContributorActivity_weeks) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetC sets the c property value. The c property
func (m *ContributorActivity_weeks) SetC(value *int32)() {
    m.c = value
}
// SetD sets the d property value. The d property
func (m *ContributorActivity_weeks) SetD(value *int32)() {
    m.d = value
}
// SetW sets the w property value. The w property
func (m *ContributorActivity_weeks) SetW(value *int32)() {
    m.w = value
}
type ContributorActivity_weeksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetA()(*int32)
    GetC()(*int32)
    GetD()(*int32)
    GetW()(*int32)
    SetA(value *int32)()
    SetC(value *int32)()
    SetD(value *int32)()
    SetW(value *int32)()
}
