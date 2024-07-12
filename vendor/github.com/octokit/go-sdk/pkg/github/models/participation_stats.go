package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ParticipationStats struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The all property
    all []int32
    // The owner property
    owner []int32
}
// NewParticipationStats instantiates a new ParticipationStats and sets the default values.
func NewParticipationStats()(*ParticipationStats) {
    m := &ParticipationStats{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateParticipationStatsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateParticipationStatsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParticipationStats(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ParticipationStats) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAll gets the all property value. The all property
// returns a []int32 when successful
func (m *ParticipationStats) GetAll()([]int32) {
    return m.all
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ParticipationStats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["all"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetAll(res)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetOwner(res)
        }
        return nil
    }
    return res
}
// GetOwner gets the owner property value. The owner property
// returns a []int32 when successful
func (m *ParticipationStats) GetOwner()([]int32) {
    return m.owner
}
// Serialize serializes information the current object
func (m *ParticipationStats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAll() != nil {
        err := writer.WriteCollectionOfInt32Values("all", m.GetAll())
        if err != nil {
            return err
        }
    }
    if m.GetOwner() != nil {
        err := writer.WriteCollectionOfInt32Values("owner", m.GetOwner())
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
func (m *ParticipationStats) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAll sets the all property value. The all property
func (m *ParticipationStats) SetAll(value []int32)() {
    m.all = value
}
// SetOwner sets the owner property value. The owner property
func (m *ParticipationStats) SetOwner(value []int32)() {
    m.owner = value
}
type ParticipationStatsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAll()([]int32)
    GetOwner()([]int32)
    SetAll(value []int32)()
    SetOwner(value []int32)()
}
