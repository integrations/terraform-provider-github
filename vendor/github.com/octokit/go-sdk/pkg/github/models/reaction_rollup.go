package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ReactionRollup struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The confused property
    confused *int32
    // The eyes property
    eyes *int32
    // The heart property
    heart *int32
    // The hooray property
    hooray *int32
    // The laugh property
    laugh *int32
    // The minus_1 property
    minus_1 *int32
    // The plus_1 property
    plus_1 *int32
    // The rocket property
    rocket *int32
    // The total_count property
    total_count *int32
    // The url property
    url *string
}
// NewReactionRollup instantiates a new ReactionRollup and sets the default values.
func NewReactionRollup()(*ReactionRollup) {
    m := &ReactionRollup{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReactionRollupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReactionRollupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReactionRollup(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReactionRollup) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfused gets the confused property value. The confused property
// returns a *int32 when successful
func (m *ReactionRollup) GetConfused()(*int32) {
    return m.confused
}
// GetEyes gets the eyes property value. The eyes property
// returns a *int32 when successful
func (m *ReactionRollup) GetEyes()(*int32) {
    return m.eyes
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReactionRollup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["confused"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfused(val)
        }
        return nil
    }
    res["eyes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEyes(val)
        }
        return nil
    }
    res["heart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeart(val)
        }
        return nil
    }
    res["hooray"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHooray(val)
        }
        return nil
    }
    res["laugh"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLaugh(val)
        }
        return nil
    }
    res["-1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinus1(val)
        }
        return nil
    }
    res["+1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlus1(val)
        }
        return nil
    }
    res["rocket"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRocket(val)
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
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetHeart gets the heart property value. The heart property
// returns a *int32 when successful
func (m *ReactionRollup) GetHeart()(*int32) {
    return m.heart
}
// GetHooray gets the hooray property value. The hooray property
// returns a *int32 when successful
func (m *ReactionRollup) GetHooray()(*int32) {
    return m.hooray
}
// GetLaugh gets the laugh property value. The laugh property
// returns a *int32 when successful
func (m *ReactionRollup) GetLaugh()(*int32) {
    return m.laugh
}
// GetMinus1 gets the -1 property value. The minus_1 property
// returns a *int32 when successful
func (m *ReactionRollup) GetMinus1()(*int32) {
    return m.minus_1
}
// GetPlus1 gets the +1 property value. The plus_1 property
// returns a *int32 when successful
func (m *ReactionRollup) GetPlus1()(*int32) {
    return m.plus_1
}
// GetRocket gets the rocket property value. The rocket property
// returns a *int32 when successful
func (m *ReactionRollup) GetRocket()(*int32) {
    return m.rocket
}
// GetTotalCount gets the total_count property value. The total_count property
// returns a *int32 when successful
func (m *ReactionRollup) GetTotalCount()(*int32) {
    return m.total_count
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *ReactionRollup) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ReactionRollup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("confused", m.GetConfused())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("eyes", m.GetEyes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("heart", m.GetHeart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("hooray", m.GetHooray())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("laugh", m.GetLaugh())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("-1", m.GetMinus1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("+1", m.GetPlus1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("rocket", m.GetRocket())
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
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *ReactionRollup) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfused sets the confused property value. The confused property
func (m *ReactionRollup) SetConfused(value *int32)() {
    m.confused = value
}
// SetEyes sets the eyes property value. The eyes property
func (m *ReactionRollup) SetEyes(value *int32)() {
    m.eyes = value
}
// SetHeart sets the heart property value. The heart property
func (m *ReactionRollup) SetHeart(value *int32)() {
    m.heart = value
}
// SetHooray sets the hooray property value. The hooray property
func (m *ReactionRollup) SetHooray(value *int32)() {
    m.hooray = value
}
// SetLaugh sets the laugh property value. The laugh property
func (m *ReactionRollup) SetLaugh(value *int32)() {
    m.laugh = value
}
// SetMinus1 sets the -1 property value. The minus_1 property
func (m *ReactionRollup) SetMinus1(value *int32)() {
    m.minus_1 = value
}
// SetPlus1 sets the +1 property value. The plus_1 property
func (m *ReactionRollup) SetPlus1(value *int32)() {
    m.plus_1 = value
}
// SetRocket sets the rocket property value. The rocket property
func (m *ReactionRollup) SetRocket(value *int32)() {
    m.rocket = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ReactionRollup) SetTotalCount(value *int32)() {
    m.total_count = value
}
// SetUrl sets the url property value. The url property
func (m *ReactionRollup) SetUrl(value *string)() {
    m.url = value
}
type ReactionRollupable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfused()(*int32)
    GetEyes()(*int32)
    GetHeart()(*int32)
    GetHooray()(*int32)
    GetLaugh()(*int32)
    GetMinus1()(*int32)
    GetPlus1()(*int32)
    GetRocket()(*int32)
    GetTotalCount()(*int32)
    GetUrl()(*string)
    SetConfused(value *int32)()
    SetEyes(value *int32)()
    SetHeart(value *int32)()
    SetHooray(value *int32)()
    SetLaugh(value *int32)()
    SetMinus1(value *int32)()
    SetPlus1(value *int32)()
    SetRocket(value *int32)()
    SetTotalCount(value *int32)()
    SetUrl(value *string)()
}
