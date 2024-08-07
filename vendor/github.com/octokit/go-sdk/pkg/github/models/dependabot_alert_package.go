package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependabotAlertPackage details for the vulnerable package.
type DependabotAlertPackage struct {
    // The package's language or package management ecosystem.
    ecosystem *string
    // The unique package name within its ecosystem.
    name *string
}
// NewDependabotAlertPackage instantiates a new DependabotAlertPackage and sets the default values.
func NewDependabotAlertPackage()(*DependabotAlertPackage) {
    m := &DependabotAlertPackage{
    }
    return m
}
// CreateDependabotAlertPackageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependabotAlertPackageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependabotAlertPackage(), nil
}
// GetEcosystem gets the ecosystem property value. The package's language or package management ecosystem.
// returns a *string when successful
func (m *DependabotAlertPackage) GetEcosystem()(*string) {
    return m.ecosystem
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependabotAlertPackage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ecosystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEcosystem(val)
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
// GetName gets the name property value. The unique package name within its ecosystem.
// returns a *string when successful
func (m *DependabotAlertPackage) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *DependabotAlertPackage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// SetEcosystem sets the ecosystem property value. The package's language or package management ecosystem.
func (m *DependabotAlertPackage) SetEcosystem(value *string)() {
    m.ecosystem = value
}
// SetName sets the name property value. The unique package name within its ecosystem.
func (m *DependabotAlertPackage) SetName(value *string)() {
    m.name = value
}
type DependabotAlertPackageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEcosystem()(*string)
    GetName()(*string)
    SetEcosystem(value *string)()
    SetName(value *string)()
}
