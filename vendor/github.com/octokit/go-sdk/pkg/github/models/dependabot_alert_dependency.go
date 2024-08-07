package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependabotAlert_dependency details for the vulnerable dependency.
type DependabotAlert_dependency struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The full path to the dependency manifest file, relative to the root of the repository.
    manifest_path *string
    // Details for the vulnerable package.
    packageEscaped DependabotAlertPackageable
    // The execution scope of the vulnerable dependency.
    scope *DependabotAlert_dependency_scope
}
// NewDependabotAlert_dependency instantiates a new DependabotAlert_dependency and sets the default values.
func NewDependabotAlert_dependency()(*DependabotAlert_dependency) {
    m := &DependabotAlert_dependency{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDependabotAlert_dependencyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependabotAlert_dependencyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependabotAlert_dependency(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DependabotAlert_dependency) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependabotAlert_dependency) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["manifest_path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManifestPath(val)
        }
        return nil
    }
    res["package"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDependabotAlertPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageEscaped(val.(DependabotAlertPackageable))
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDependabotAlert_dependency_scope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(*DependabotAlert_dependency_scope))
        }
        return nil
    }
    return res
}
// GetManifestPath gets the manifest_path property value. The full path to the dependency manifest file, relative to the root of the repository.
// returns a *string when successful
func (m *DependabotAlert_dependency) GetManifestPath()(*string) {
    return m.manifest_path
}
// GetPackageEscaped gets the package property value. Details for the vulnerable package.
// returns a DependabotAlertPackageable when successful
func (m *DependabotAlert_dependency) GetPackageEscaped()(DependabotAlertPackageable) {
    return m.packageEscaped
}
// GetScope gets the scope property value. The execution scope of the vulnerable dependency.
// returns a *DependabotAlert_dependency_scope when successful
func (m *DependabotAlert_dependency) GetScope()(*DependabotAlert_dependency_scope) {
    return m.scope
}
// Serialize serializes information the current object
func (m *DependabotAlert_dependency) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DependabotAlert_dependency) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetManifestPath sets the manifest_path property value. The full path to the dependency manifest file, relative to the root of the repository.
func (m *DependabotAlert_dependency) SetManifestPath(value *string)() {
    m.manifest_path = value
}
// SetPackageEscaped sets the package property value. Details for the vulnerable package.
func (m *DependabotAlert_dependency) SetPackageEscaped(value DependabotAlertPackageable)() {
    m.packageEscaped = value
}
// SetScope sets the scope property value. The execution scope of the vulnerable dependency.
func (m *DependabotAlert_dependency) SetScope(value *DependabotAlert_dependency_scope)() {
    m.scope = value
}
type DependabotAlert_dependencyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetManifestPath()(*string)
    GetPackageEscaped()(DependabotAlertPackageable)
    GetScope()(*DependabotAlert_dependency_scope)
    SetManifestPath(value *string)()
    SetPackageEscaped(value DependabotAlertPackageable)()
    SetScope(value *DependabotAlert_dependency_scope)()
}
