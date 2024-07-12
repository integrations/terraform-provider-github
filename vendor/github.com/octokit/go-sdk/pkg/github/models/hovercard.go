package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Hovercard hovercard
type Hovercard struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The contexts property
    contexts []Hovercard_contextsable
}
// NewHovercard instantiates a new Hovercard and sets the default values.
func NewHovercard()(*Hovercard) {
    m := &Hovercard{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHovercardFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHovercardFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHovercard(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Hovercard) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContexts gets the contexts property value. The contexts property
// returns a []Hovercard_contextsable when successful
func (m *Hovercard) GetContexts()([]Hovercard_contextsable) {
    return m.contexts
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Hovercard) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contexts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHovercard_contextsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Hovercard_contextsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Hovercard_contextsable)
                }
            }
            m.SetContexts(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Hovercard) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContexts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContexts()))
        for i, v := range m.GetContexts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("contexts", cast)
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
func (m *Hovercard) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContexts sets the contexts property value. The contexts property
func (m *Hovercard) SetContexts(value []Hovercard_contextsable)() {
    m.contexts = value
}
type Hovercardable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContexts()([]Hovercard_contextsable)
    SetContexts(value []Hovercard_contextsable)()
}
