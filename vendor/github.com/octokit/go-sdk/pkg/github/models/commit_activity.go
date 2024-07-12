package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommitActivity commit Activity
type CommitActivity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The days property
    days []int32
    // The total property
    total *int32
    // The week property
    week *int32
}
// NewCommitActivity instantiates a new CommitActivity and sets the default values.
func NewCommitActivity()(*CommitActivity) {
    m := &CommitActivity{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCommitActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCommitActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommitActivity(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CommitActivity) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDays gets the days property value. The days property
// returns a []int32 when successful
func (m *CommitActivity) GetDays()([]int32) {
    return m.days
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CommitActivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["days"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetDays(res)
        }
        return nil
    }
    res["total"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    res["week"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeek(val)
        }
        return nil
    }
    return res
}
// GetTotal gets the total property value. The total property
// returns a *int32 when successful
func (m *CommitActivity) GetTotal()(*int32) {
    return m.total
}
// GetWeek gets the week property value. The week property
// returns a *int32 when successful
func (m *CommitActivity) GetWeek()(*int32) {
    return m.week
}
// Serialize serializes information the current object
func (m *CommitActivity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDays() != nil {
        err := writer.WriteCollectionOfInt32Values("days", m.GetDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total", m.GetTotal())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("week", m.GetWeek())
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
func (m *CommitActivity) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDays sets the days property value. The days property
func (m *CommitActivity) SetDays(value []int32)() {
    m.days = value
}
// SetTotal sets the total property value. The total property
func (m *CommitActivity) SetTotal(value *int32)() {
    m.total = value
}
// SetWeek sets the week property value. The week property
func (m *CommitActivity) SetWeek(value *int32)() {
    m.week = value
}
type CommitActivityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDays()([]int32)
    GetTotal()(*int32)
    GetWeek()(*int32)
    SetDays(value []int32)()
    SetTotal(value *int32)()
    SetWeek(value *int32)()
}
