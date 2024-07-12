package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PagesHttpsCertificate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description property
    description *string
    // Array of the domain set and its alternate name (if it is configured)
    domains []string
    // The expires_at property
    expires_at *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The state property
    state *PagesHttpsCertificate_state
}
// NewPagesHttpsCertificate instantiates a new PagesHttpsCertificate and sets the default values.
func NewPagesHttpsCertificate()(*PagesHttpsCertificate) {
    m := &PagesHttpsCertificate{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePagesHttpsCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePagesHttpsCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPagesHttpsCertificate(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PagesHttpsCertificate) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *PagesHttpsCertificate) GetDescription()(*string) {
    return m.description
}
// GetDomains gets the domains property value. Array of the domain set and its alternate name (if it is configured)
// returns a []string when successful
func (m *PagesHttpsCertificate) GetDomains()([]string) {
    return m.domains
}
// GetExpiresAt gets the expires_at property value. The expires_at property
// returns a *DateOnly when successful
func (m *PagesHttpsCertificate) GetExpiresAt()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.expires_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PagesHttpsCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["domains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDomains(res)
        }
        return nil
    }
    res["expires_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiresAt(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePagesHttpsCertificate_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*PagesHttpsCertificate_state))
        }
        return nil
    }
    return res
}
// GetState gets the state property value. The state property
// returns a *PagesHttpsCertificate_state when successful
func (m *PagesHttpsCertificate) GetState()(*PagesHttpsCertificate_state) {
    return m.state
}
// Serialize serializes information the current object
func (m *PagesHttpsCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDomains() != nil {
        err := writer.WriteCollectionOfStringValues("domains", m.GetDomains())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteDateOnlyValue("expires_at", m.GetExpiresAt())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *PagesHttpsCertificate) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description property
func (m *PagesHttpsCertificate) SetDescription(value *string)() {
    m.description = value
}
// SetDomains sets the domains property value. Array of the domain set and its alternate name (if it is configured)
func (m *PagesHttpsCertificate) SetDomains(value []string)() {
    m.domains = value
}
// SetExpiresAt sets the expires_at property value. The expires_at property
func (m *PagesHttpsCertificate) SetExpiresAt(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.expires_at = value
}
// SetState sets the state property value. The state property
func (m *PagesHttpsCertificate) SetState(value *PagesHttpsCertificate_state)() {
    m.state = value
}
type PagesHttpsCertificateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDomains()([]string)
    GetExpiresAt()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetState()(*PagesHttpsCertificate_state)
    SetDescription(value *string)()
    SetDomains(value []string)()
    SetExpiresAt(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetState(value *PagesHttpsCertificate_state)()
}
