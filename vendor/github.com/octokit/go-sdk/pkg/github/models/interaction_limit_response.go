package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InteractionLimitResponse interaction limit settings.
type InteractionLimitResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The expires_at property
    expires_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
    limit *InteractionGroup
    // The origin property
    origin *string
}
// NewInteractionLimitResponse instantiates a new InteractionLimitResponse and sets the default values.
func NewInteractionLimitResponse()(*InteractionLimitResponse) {
    m := &InteractionLimitResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInteractionLimitResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInteractionLimitResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInteractionLimitResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *InteractionLimitResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExpiresAt gets the expires_at property value. The expires_at property
// returns a *Time when successful
func (m *InteractionLimitResponse) GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expires_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InteractionLimitResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expires_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiresAt(val)
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
    res["origin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrigin(val)
        }
        return nil
    }
    return res
}
// GetLimit gets the limit property value. The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
// returns a *InteractionGroup when successful
func (m *InteractionLimitResponse) GetLimit()(*InteractionGroup) {
    return m.limit
}
// GetOrigin gets the origin property value. The origin property
// returns a *string when successful
func (m *InteractionLimitResponse) GetOrigin()(*string) {
    return m.origin
}
// Serialize serializes information the current object
func (m *InteractionLimitResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("expires_at", m.GetExpiresAt())
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
        err := writer.WriteStringValue("origin", m.GetOrigin())
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
func (m *InteractionLimitResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExpiresAt sets the expires_at property value. The expires_at property
func (m *InteractionLimitResponse) SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expires_at = value
}
// SetLimit sets the limit property value. The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
func (m *InteractionLimitResponse) SetLimit(value *InteractionGroup)() {
    m.limit = value
}
// SetOrigin sets the origin property value. The origin property
func (m *InteractionLimitResponse) SetOrigin(value *string)() {
    m.origin = value
}
type InteractionLimitResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLimit()(*InteractionGroup)
    GetOrigin()(*string)
    SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLimit(value *InteractionGroup)()
    SetOrigin(value *string)()
}
