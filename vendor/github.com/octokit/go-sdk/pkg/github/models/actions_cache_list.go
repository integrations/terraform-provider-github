package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActionsCacheList repository actions caches
type ActionsCacheList struct {
    // Array of caches
    actions_caches []ActionsCacheList_actions_cachesable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Total number of caches
    total_count *int32
}
// NewActionsCacheList instantiates a new ActionsCacheList and sets the default values.
func NewActionsCacheList()(*ActionsCacheList) {
    m := &ActionsCacheList{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsCacheListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsCacheListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsCacheList(), nil
}
// GetActionsCaches gets the actions_caches property value. Array of caches
// returns a []ActionsCacheList_actions_cachesable when successful
func (m *ActionsCacheList) GetActionsCaches()([]ActionsCacheList_actions_cachesable) {
    return m.actions_caches
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsCacheList) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsCacheList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actions_caches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateActionsCacheList_actions_cachesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ActionsCacheList_actions_cachesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ActionsCacheList_actions_cachesable)
                }
            }
            m.SetActionsCaches(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetTotalCount gets the total_count property value. Total number of caches
// returns a *int32 when successful
func (m *ActionsCacheList) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ActionsCacheList) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActionsCaches() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActionsCaches()))
        for i, v := range m.GetActionsCaches() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("actions_caches", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
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
// SetActionsCaches sets the actions_caches property value. Array of caches
func (m *ActionsCacheList) SetActionsCaches(value []ActionsCacheList_actions_cachesable)() {
    m.actions_caches = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActionsCacheList) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTotalCount sets the total_count property value. Total number of caches
func (m *ActionsCacheList) SetTotalCount(value *int32)() {
    m.total_count = value
}
type ActionsCacheListable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionsCaches()([]ActionsCacheList_actions_cachesable)
    GetTotalCount()(*int32)
    SetActionsCaches(value []ActionsCacheList_actions_cachesable)()
    SetTotalCount(value *int32)()
}
