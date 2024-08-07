package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReferrerTraffic referrer Traffic
type ReferrerTraffic struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The count property
    count *int32
    // The referrer property
    referrer *string
    // The uniques property
    uniques *int32
}
// NewReferrerTraffic instantiates a new ReferrerTraffic and sets the default values.
func NewReferrerTraffic()(*ReferrerTraffic) {
    m := &ReferrerTraffic{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReferrerTrafficFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReferrerTrafficFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReferrerTraffic(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReferrerTraffic) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *ReferrerTraffic) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReferrerTraffic) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
        }
        return nil
    }
    res["referrer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferrer(val)
        }
        return nil
    }
    res["uniques"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniques(val)
        }
        return nil
    }
    return res
}
// GetReferrer gets the referrer property value. The referrer property
// returns a *string when successful
func (m *ReferrerTraffic) GetReferrer()(*string) {
    return m.referrer
}
// GetUniques gets the uniques property value. The uniques property
// returns a *int32 when successful
func (m *ReferrerTraffic) GetUniques()(*int32) {
    return m.uniques
}
// Serialize serializes information the current object
func (m *ReferrerTraffic) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("count", m.GetCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("referrer", m.GetReferrer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("uniques", m.GetUniques())
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
func (m *ReferrerTraffic) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. The count property
func (m *ReferrerTraffic) SetCount(value *int32)() {
    m.count = value
}
// SetReferrer sets the referrer property value. The referrer property
func (m *ReferrerTraffic) SetReferrer(value *string)() {
    m.referrer = value
}
// SetUniques sets the uniques property value. The uniques property
func (m *ReferrerTraffic) SetUniques(value *int32)() {
    m.uniques = value
}
type ReferrerTrafficable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int32)
    GetReferrer()(*string)
    GetUniques()(*int32)
    SetCount(value *int32)()
    SetReferrer(value *string)()
    SetUniques(value *int32)()
}
