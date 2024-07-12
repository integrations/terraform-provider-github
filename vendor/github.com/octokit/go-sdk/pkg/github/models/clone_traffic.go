package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloneTraffic clone Traffic
type CloneTraffic struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The clones property
    clones []Trafficable
    // The count property
    count *int32
    // The uniques property
    uniques *int32
}
// NewCloneTraffic instantiates a new CloneTraffic and sets the default values.
func NewCloneTraffic()(*CloneTraffic) {
    m := &CloneTraffic{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloneTrafficFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloneTrafficFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloneTraffic(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CloneTraffic) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClones gets the clones property value. The clones property
// returns a []Trafficable when successful
func (m *CloneTraffic) GetClones()([]Trafficable) {
    return m.clones
}
// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *CloneTraffic) GetCount()(*int32) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloneTraffic) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetClones(res)
        }
        return nil
    }
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
    return res
}
// GetUniques gets the uniques property value. The uniques property
// returns a *int32 when successful
func (m *CloneTraffic) GetUniques()(*int32) {
    return m.uniques
}
// Serialize serializes information the current object
func (m *CloneTraffic) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetClones() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClones()))
        for i, v := range m.GetClones() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("clones", cast)
        if err != nil {
            return err
        }
    }
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloneTraffic) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClones sets the clones property value. The clones property
func (m *CloneTraffic) SetClones(value []Trafficable)() {
    m.clones = value
}
// SetCount sets the count property value. The count property
func (m *CloneTraffic) SetCount(value *int32)() {
    m.count = value
}
// SetUniques sets the uniques property value. The uniques property
func (m *CloneTraffic) SetUniques(value *int32)() {
    m.uniques = value
}
type CloneTrafficable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClones()([]Trafficable)
    GetCount()(*int32)
    GetUniques()(*int32)
    SetClones(value []Trafficable)()
    SetCount(value *int32)()
    SetUniques(value *int32)()
}
