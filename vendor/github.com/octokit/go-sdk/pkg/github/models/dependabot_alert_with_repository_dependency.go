package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependabotAlertWithRepository_dependency details for the vulnerable dependency.
type DependabotAlertWithRepository_dependency struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The full path to the dependency manifest file, relative to the root of the repository.
    manifest_path *string
    // Details for the vulnerable package.
    packageEscaped DependabotAlertPackageable
    // The execution scope of the vulnerable dependency.
    scope *DependabotAlertWithRepository_dependency_scope
}
// NewDependabotAlertWithRepository_dependency instantiates a new DependabotAlertWithRepository_dependency and sets the default values.
func NewDependabotAlertWithRepository_dependency()(*DependabotAlertWithRepository_dependency) {
    m := &DependabotAlertWithRepository_dependency{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDependabotAlertWithRepository_dependencyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependabotAlertWithRepository_dependencyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependabotAlertWithRepository_dependency(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DependabotAlertWithRepository_dependency) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependabotAlertWithRepository_dependency) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetEnumValue(ParseDependabotAlertWithRepository_dependency_scope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(*DependabotAlertWithRepository_dependency_scope))
        }
        return nil
    }
    return res
}
// GetManifestPath gets the manifest_path property value. The full path to the dependency manifest file, relative to the root of the repository.
// returns a *string when successful
func (m *DependabotAlertWithRepository_dependency) GetManifestPath()(*string) {
    return m.manifest_path
}
// GetPackageEscaped gets the package property value. Details for the vulnerable package.
// returns a DependabotAlertPackageable when successful
func (m *DependabotAlertWithRepository_dependency) GetPackageEscaped()(DependabotAlertPackageable) {
    return m.packageEscaped
}
// GetScope gets the scope property value. The execution scope of the vulnerable dependency.
// returns a *DependabotAlertWithRepository_dependency_scope when successful
func (m *DependabotAlertWithRepository_dependency) GetScope()(*DependabotAlertWithRepository_dependency_scope) {
    return m.scope
}
// Serialize serializes information the current object
func (m *DependabotAlertWithRepository_dependency) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DependabotAlertWithRepository_dependency) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetManifestPath sets the manifest_path property value. The full path to the dependency manifest file, relative to the root of the repository.
func (m *DependabotAlertWithRepository_dependency) SetManifestPath(value *string)() {
    m.manifest_path = value
}
// SetPackageEscaped sets the package property value. Details for the vulnerable package.
func (m *DependabotAlertWithRepository_dependency) SetPackageEscaped(value DependabotAlertPackageable)() {
    m.packageEscaped = value
}
// SetScope sets the scope property value. The execution scope of the vulnerable dependency.
func (m *DependabotAlertWithRepository_dependency) SetScope(value *DependabotAlertWithRepository_dependency_scope)() {
    m.scope = value
}
type DependabotAlertWithRepository_dependencyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetManifestPath()(*string)
    GetPackageEscaped()(DependabotAlertPackageable)
    GetScope()(*DependabotAlertWithRepository_dependency_scope)
    SetManifestPath(value *string)()
    SetPackageEscaped(value DependabotAlertPackageable)()
    SetScope(value *DependabotAlertWithRepository_dependency_scope)()
}
