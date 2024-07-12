package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InteractionLimit limit interactions to a specific type of user for a specified duration
type InteractionLimit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The duration of the interaction restriction. Default: `one_day`.
    expiry *InteractionExpiry
    // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
    limit *InteractionGroup
}
// NewInteractionLimit instantiates a new InteractionLimit and sets the default values.
func NewInteractionLimit()(*InteractionLimit) {
    m := &InteractionLimit{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInteractionLimitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInteractionLimitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInteractionLimit(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *InteractionLimit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExpiry gets the expiry property value. The duration of the interaction restriction. Default: `one_day`.
// returns a *InteractionExpiry when successful
func (m *InteractionLimit) GetExpiry()(*InteractionExpiry) {
    return m.expiry
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InteractionLimit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expiry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInteractionExpiry)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiry(val.(*InteractionExpiry))
        }
        return nil
    }
    res["limit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInteractionGroup)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLimit(val.(*InteractionGroup))
        }
        return nil
    }
    return res
}
// GetLimit gets the limit property value. The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
// returns a *InteractionGroup when successful
func (m *InteractionLimit) GetLimit()(*InteractionGroup) {
    return m.limit
}
// Serialize serializes information the current object
func (m *InteractionLimit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExpiry() != nil {
        cast := (*m.GetExpiry()).String()
        err := writer.WriteStringValue("expiry", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLimit() != nil {
        cast := (*m.GetLimit()).String()
        err := writer.WriteStringValue("limit", &cast)
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
func (m *InteractionLimit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExpiry sets the expiry property value. The duration of the interaction restriction. Default: `one_day`.
func (m *InteractionLimit) SetExpiry(value *InteractionExpiry)() {
    m.expiry = value
}
// SetLimit sets the limit property value. The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
func (m *InteractionLimit) SetLimit(value *InteractionGroup)() {
    m.limit = value
}
type InteractionLimitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpiry()(*InteractionExpiry)
    GetLimit()(*InteractionGroup)
    SetExpiry(value *InteractionExpiry)()
    SetLimit(value *InteractionGroup)()
}
