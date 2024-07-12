package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ViewTraffic view Traffic
type ViewTraffic struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The count property
    count *int32
    // The uniques property
    uniques *int32
    // The views property
    views []Trafficable
}
// NewViewTraffic instantiates a new ViewTraffic and sets the default values.
func NewViewTraffic()(*ViewTraffic) {
    m := &ViewTraffic{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateViewTrafficFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateViewTrafficFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewViewTraffic(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ViewTraffic) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *ViewTraffic) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ViewTraffic) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["views"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTrafficFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Trafficable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Trafficable)
                }
            }
            m.SetViews(res)
        }
        return nil
    }
    return res
}
// GetUniques gets the uniques property value. The uniques property
// returns a *int32 when successful
func (m *ViewTraffic) GetUniques()(*int32) {
    return m.uniques
}
// GetViews gets the views property value. The views property
// returns a []Trafficable when successful
func (m *ViewTraffic) GetViews()([]Trafficable) {
    return m.views
}
// Serialize serializes information the current object
func (m *ViewTraffic) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("count", m.GetCount())
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
    if m.GetViews() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetViews()))
        for i, v := range m.GetViews() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("views", cast)
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
func (m *ViewTraffic) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. The count property
func (m *ViewTraffic) SetCount(value *int32)() {
    m.count = value
}
// SetUniques sets the uniques property value. The uniques property
func (m *ViewTraffic) SetUniques(value *int32)() {
    m.uniques = value
}
// SetViews sets the views property value. The views property
func (m *ViewTraffic) SetViews(value []Trafficable)() {
    m.views = value
}
type ViewTrafficable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int32)
    GetUniques()(*int32)
    GetViews()([]Trafficable)
    SetCount(value *int32)()
    SetUniques(value *int32)()
    SetViews(value []Trafficable)()
}
