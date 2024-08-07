package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PagesHealthCheck pages Health Check Status
type PagesHealthCheck struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The alt_domain property
    alt_domain PagesHealthCheck_alt_domainable
    // The domain property
    domain PagesHealthCheck_domainable
}
// NewPagesHealthCheck instantiates a new PagesHealthCheck and sets the default values.
func NewPagesHealthCheck()(*PagesHealthCheck) {
    m := &PagesHealthCheck{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePagesHealthCheckFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePagesHealthCheckFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPagesHealthCheck(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PagesHealthCheck) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAltDomain gets the alt_domain property value. The alt_domain property
// returns a PagesHealthCheck_alt_domainable when successful
func (m *PagesHealthCheck) GetAltDomain()(PagesHealthCheck_alt_domainable) {
    return m.alt_domain
}
// GetDomain gets the domain property value. The domain property
// returns a PagesHealthCheck_domainable when successful
func (m *PagesHealthCheck) GetDomain()(PagesHealthCheck_domainable) {
    return m.domain
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PagesHealthCheck) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alt_domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePagesHealthCheck_alt_domainFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAltDomain(val.(PagesHealthCheck_alt_domainable))
        }
        return nil
    }
    res["domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePagesHealthCheck_domainFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomain(val.(PagesHealthCheck_domainable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PagesHealthCheck) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("alt_domain", m.GetAltDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("domain", m.GetDomain())
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
func (m *PagesHealthCheck) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAltDomain sets the alt_domain property value. The alt_domain property
func (m *PagesHealthCheck) SetAltDomain(value PagesHealthCheck_alt_domainable)() {
    m.alt_domain = value
}
// SetDomain sets the domain property value. The domain property
func (m *PagesHealthCheck) SetDomain(value PagesHealthCheck_domainable)() {
    m.domain = value
}
type PagesHealthCheckable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAltDomain()(PagesHealthCheck_alt_domainable)
    GetDomain()(PagesHealthCheck_domainable)
    SetAltDomain(value PagesHealthCheck_alt_domainable)()
    SetDomain(value PagesHealthCheck_domainable)()
}
