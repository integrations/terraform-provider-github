package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependabotAlertSecurityAdvisory_references a link to additional advisory information.
type DependabotAlertSecurityAdvisory_references struct {
    // The URL of the reference.
    url *string
}
// NewDependabotAlertSecurityAdvisory_references instantiates a new DependabotAlertSecurityAdvisory_references and sets the default values.
func NewDependabotAlertSecurityAdvisory_references()(*DependabotAlertSecurityAdvisory_references) {
    m := &DependabotAlertSecurityAdvisory_references{
    }
    return m
}
// CreateDependabotAlertSecurityAdvisory_referencesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependabotAlertSecurityAdvisory_referencesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependabotAlertSecurityAdvisory_references(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependabotAlertSecurityAdvisory_references) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
// GetUrl gets the url property value. The URL of the reference.
// returns a *string when successful
func (m *DependabotAlertSecurityAdvisory_references) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *DependabotAlertSecurityAdvisory_references) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// SetUrl sets the url property value. The URL of the reference.
func (m *DependabotAlertSecurityAdvisory_references) SetUrl(value *string)() {
    m.url = value
}
type DependabotAlertSecurityAdvisory_referencesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUrl()(*string)
    SetUrl(value *string)()
}
