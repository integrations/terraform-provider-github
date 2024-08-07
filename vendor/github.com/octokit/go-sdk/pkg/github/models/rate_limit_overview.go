package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RateLimitOverview rate Limit Overview
type RateLimitOverview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The rate property
    rate RateLimitable
    // The resources property
    resources RateLimitOverview_resourcesable
}
// NewRateLimitOverview instantiates a new RateLimitOverview and sets the default values.
func NewRateLimitOverview()(*RateLimitOverview) {
    m := &RateLimitOverview{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRateLimitOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRateLimitOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRateLimitOverview(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RateLimitOverview) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RateLimitOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["rate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRateLimitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRate(val.(RateLimitable))
        }
        return nil
    }
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRateLimitOverview_resourcesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResources(val.(RateLimitOverview_resourcesable))
        }
        return nil
    }
    return res
}
// GetRate gets the rate property value. The rate property
// returns a RateLimitable when successful
func (m *RateLimitOverview) GetRate()(RateLimitable) {
    return m.rate
}
// GetResources gets the resources property value. The resources property
// returns a RateLimitOverview_resourcesable when successful
func (m *RateLimitOverview) GetResources()(RateLimitOverview_resourcesable) {
    return m.resources
}
// Serialize serializes information the current object
func (m *RateLimitOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("rate", m.GetRate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resources", m.GetResources())
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
func (m *RateLimitOverview) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRate sets the rate property value. The rate property
func (m *RateLimitOverview) SetRate(value RateLimitable)() {
    m.rate = value
}
// SetResources sets the resources property value. The resources property
func (m *RateLimitOverview) SetResources(value RateLimitOverview_resourcesable)() {
    m.resources = value
}
type RateLimitOverviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRate()(RateLimitable)
    GetResources()(RateLimitOverview_resourcesable)
    SetRate(value RateLimitable)()
    SetResources(value RateLimitOverview_resourcesable)()
}
