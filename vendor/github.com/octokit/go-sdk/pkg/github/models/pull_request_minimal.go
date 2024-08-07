package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PullRequestMinimal struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The base property
    base PullRequestMinimal_baseable
    // The head property
    head PullRequestMinimal_headable
    // The id property
    id *int64
    // The number property
    number *int32
    // The url property
    url *string
}
// NewPullRequestMinimal instantiates a new PullRequestMinimal and sets the default values.
func NewPullRequestMinimal()(*PullRequestMinimal) {
    m := &PullRequestMinimal{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePullRequestMinimalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePullRequestMinimalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPullRequestMinimal(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PullRequestMinimal) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBase gets the base property value. The base property
// returns a PullRequestMinimal_baseable when successful
func (m *PullRequestMinimal) GetBase()(PullRequestMinimal_baseable) {
    return m.base
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PullRequestMinimal) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["base"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePullRequestMinimal_baseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBase(val.(PullRequestMinimal_baseable))
        }
        return nil
    }
    res["head"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePullRequestMinimal_headFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHead(val.(PullRequestMinimal_headable))
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
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
// GetHead gets the head property value. The head property
// returns a PullRequestMinimal_headable when successful
func (m *PullRequestMinimal) GetHead()(PullRequestMinimal_headable) {
    return m.head
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *PullRequestMinimal) GetId()(*int64) {
    return m.id
}
// GetNumber gets the number property value. The number property
// returns a *int32 when successful
func (m *PullRequestMinimal) GetNumber()(*int32) {
    return m.number
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *PullRequestMinimal) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *PullRequestMinimal) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("base", m.GetBase())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("head", m.GetHead())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("number", m.GetNumber())
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
func (m *PullRequestMinimal) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBase sets the base property value. The base property
func (m *PullRequestMinimal) SetBase(value PullRequestMinimal_baseable)() {
    m.base = value
}
// SetHead sets the head property value. The head property
func (m *PullRequestMinimal) SetHead(value PullRequestMinimal_headable)() {
    m.head = value
}
// SetId sets the id property value. The id property
func (m *PullRequestMinimal) SetId(value *int64)() {
    m.id = value
}
// SetNumber sets the number property value. The number property
func (m *PullRequestMinimal) SetNumber(value *int32)() {
    m.number = value
}
// SetUrl sets the url property value. The url property
func (m *PullRequestMinimal) SetUrl(value *string)() {
    m.url = value
}
type PullRequestMinimalable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBase()(PullRequestMinimal_baseable)
    GetHead()(PullRequestMinimal_headable)
    GetId()(*int64)
    GetNumber()(*int32)
    GetUrl()(*string)
    SetBase(value PullRequestMinimal_baseable)()
    SetHead(value PullRequestMinimal_headable)()
    SetId(value *int64)()
    SetNumber(value *int32)()
    SetUrl(value *string)()
}
