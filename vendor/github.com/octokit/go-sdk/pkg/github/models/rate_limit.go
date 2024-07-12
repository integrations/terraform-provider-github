package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RateLimit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The limit property
    limit *int32
    // The remaining property
    remaining *int32
    // The reset property
    reset *int32
    // The used property
    used *int32
}
// NewRateLimit instantiates a new RateLimit and sets the default values.
func NewRateLimit()(*RateLimit) {
    m := &RateLimit{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRateLimitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRateLimitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRateLimit(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RateLimit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RateLimit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["limit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLimit(val)
        }
        return nil
    }
    res["remaining"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemaining(val)
        }
        return nil
    }
    res["reset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReset(val)
        }
        return nil
    }
    res["used"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsed(val)
        }
        return nil
    }
    return res
}
// GetLimit gets the limit property value. The limit property
// returns a *int32 when successful
func (m *RateLimit) GetLimit()(*int32) {
    return m.limit
}
// GetRemaining gets the remaining property value. The remaining property
// returns a *int32 when successful
func (m *RateLimit) GetRemaining()(*int32) {
    return m.remaining
}
// GetReset gets the reset property value. The reset property
// returns a *int32 when successful
func (m *RateLimit) GetReset()(*int32) {
    return m.reset
}
// GetUsed gets the used property value. The used property
// returns a *int32 when successful
func (m *RateLimit) GetUsed()(*int32) {
    return m.used
}
// Serialize serializes information the current object
func (m *RateLimit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("limit", m.GetLimit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("remaining", m.GetRemaining())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("reset", m.GetReset())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("used", m.GetUsed())
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
func (m *RateLimit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLimit sets the limit property value. The limit property
func (m *RateLimit) SetLimit(value *int32)() {
    m.limit = value
}
// SetRemaining sets the remaining property value. The remaining property
func (m *RateLimit) SetRemaining(value *int32)() {
    m.remaining = value
}
// SetReset sets the reset property value. The reset property
func (m *RateLimit) SetReset(value *int32)() {
    m.reset = value
}
// SetUsed sets the used property value. The used property
func (m *RateLimit) SetUsed(value *int32)() {
    m.used = value
}
type RateLimitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLimit()(*int32)
    GetRemaining()(*int32)
    GetReset()(*int32)
    GetUsed()(*int32)
    SetLimit(value *int32)()
    SetRemaining(value *int32)()
    SetReset(value *int32)()
    SetUsed(value *int32)()
}
