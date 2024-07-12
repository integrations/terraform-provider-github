package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependabotAlertSecurityAdvisory_cwes a CWE weakness assigned to the advisory.
type DependabotAlertSecurityAdvisory_cwes struct {
    // The unique CWE ID.
    cwe_id *string
    // The short, plain text name of the CWE.
    name *string
}
// NewDependabotAlertSecurityAdvisory_cwes instantiates a new DependabotAlertSecurityAdvisory_cwes and sets the default values.
func NewDependabotAlertSecurityAdvisory_cwes()(*DependabotAlertSecurityAdvisory_cwes) {
    m := &DependabotAlertSecurityAdvisory_cwes{
    }
    return m
}
// CreateDependabotAlertSecurityAdvisory_cwesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependabotAlertSecurityAdvisory_cwesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependabotAlertSecurityAdvisory_cwes(), nil
}
// GetCweId gets the cwe_id property value. The unique CWE ID.
// returns a *string when successful
func (m *DependabotAlertSecurityAdvisory_cwes) GetCweId()(*string) {
    return m.cwe_id
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependabotAlertSecurityAdvisory_cwes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cwe_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCweId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The short, plain text name of the CWE.
// returns a *string when successful
func (m *DependabotAlertSecurityAdvisory_cwes) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *DependabotAlertSecurityAdvisory_cwes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// SetCweId sets the cwe_id property value. The unique CWE ID.
func (m *DependabotAlertSecurityAdvisory_cwes) SetCweId(value *string)() {
    m.cwe_id = value
}
// SetName sets the name property value. The short, plain text name of the CWE.
func (m *DependabotAlertSecurityAdvisory_cwes) SetName(value *string)() {
    m.name = value
}
type DependabotAlertSecurityAdvisory_cwesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCweId()(*string)
    GetName()(*string)
    SetCweId(value *string)()
    SetName(value *string)()
}
