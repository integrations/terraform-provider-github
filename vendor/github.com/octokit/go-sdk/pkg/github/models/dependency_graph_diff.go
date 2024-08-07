package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DependencyGraphDiff struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The change_type property
    change_type *DependencyGraphDiff_change_type
    // The ecosystem property
    ecosystem *string
    // The license property
    license *string
    // The manifest property
    manifest *string
    // The name property
    name *string
    // The package_url property
    package_url *string
    // Where the dependency is utilized. `development` means that the dependency is only utilized in the development environment. `runtime` means that the dependency is utilized at runtime and in the development environment.
    scope *DependencyGraphDiff_scope
    // The source_repository_url property
    source_repository_url *string
    // The version property
    version *string
    // The vulnerabilities property
    vulnerabilities []DependencyGraphDiff_vulnerabilitiesable
}
// NewDependencyGraphDiff instantiates a new DependencyGraphDiff and sets the default values.
func NewDependencyGraphDiff()(*DependencyGraphDiff) {
    m := &DependencyGraphDiff{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDependencyGraphDiffFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependencyGraphDiffFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependencyGraphDiff(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DependencyGraphDiff) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChangeType gets the change_type property value. The change_type property
// returns a *DependencyGraphDiff_change_type when successful
func (m *DependencyGraphDiff) GetChangeType()(*DependencyGraphDiff_change_type) {
    return m.change_type
}
// GetEcosystem gets the ecosystem property value. The ecosystem property
// returns a *string when successful
func (m *DependencyGraphDiff) GetEcosystem()(*string) {
    return m.ecosystem
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependencyGraphDiff) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["change_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDependencyGraphDiff_change_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeType(val.(*DependencyGraphDiff_change_type))
        }
        return nil
    }
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
    res["license"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicense(val)
        }
        return nil
    }
    res["manifest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManifest(val)
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
    res["package_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageUrl(val)
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDependencyGraphDiff_scope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(*DependencyGraphDiff_scope))
        }
        return nil
    }
    res["source_repository_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceRepositoryUrl(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    res["vulnerabilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDependencyGraphDiff_vulnerabilitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DependencyGraphDiff_vulnerabilitiesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DependencyGraphDiff_vulnerabilitiesable)
                }
            }
            m.SetVulnerabilities(res)
        }
        return nil
    }
    return res
}
// GetLicense gets the license property value. The license property
// returns a *string when successful
func (m *DependencyGraphDiff) GetLicense()(*string) {
    return m.license
}
// GetManifest gets the manifest property value. The manifest property
// returns a *string when successful
func (m *DependencyGraphDiff) GetManifest()(*string) {
    return m.manifest
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *DependencyGraphDiff) GetName()(*string) {
    return m.name
}
// GetPackageUrl gets the package_url property value. The package_url property
// returns a *string when successful
func (m *DependencyGraphDiff) GetPackageUrl()(*string) {
    return m.package_url
}
// GetScope gets the scope property value. Where the dependency is utilized. `development` means that the dependency is only utilized in the development environment. `runtime` means that the dependency is utilized at runtime and in the development environment.
// returns a *DependencyGraphDiff_scope when successful
func (m *DependencyGraphDiff) GetScope()(*DependencyGraphDiff_scope) {
    return m.scope
}
// GetSourceRepositoryUrl gets the source_repository_url property value. The source_repository_url property
// returns a *string when successful
func (m *DependencyGraphDiff) GetSourceRepositoryUrl()(*string) {
    return m.source_repository_url
}
// GetVersion gets the version property value. The version property
// returns a *string when successful
func (m *DependencyGraphDiff) GetVersion()(*string) {
    return m.version
}
// GetVulnerabilities gets the vulnerabilities property value. The vulnerabilities property
// returns a []DependencyGraphDiff_vulnerabilitiesable when successful
func (m *DependencyGraphDiff) GetVulnerabilities()([]DependencyGraphDiff_vulnerabilitiesable) {
    return m.vulnerabilities
}
// Serialize serializes information the current object
func (m *DependencyGraphDiff) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetChangeType() != nil {
        cast := (*m.GetChangeType()).String()
        err := writer.WriteStringValue("change_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ecosystem", m.GetEcosystem())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("license", m.GetLicense())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("manifest", m.GetManifest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("package_url", m.GetPackageUrl())
        if err != nil {
            return err
        }
    }
    if m.GetScope() != nil {
        cast := (*m.GetScope()).String()
        err := writer.WriteStringValue("scope", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("source_repository_url", m.GetSourceRepositoryUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    if m.GetVulnerabilities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVulnerabilities()))
        for i, v := range m.GetVulnerabilities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("vulnerabilities", cast)
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
func (m *DependencyGraphDiff) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChangeType sets the change_type property value. The change_type property
func (m *DependencyGraphDiff) SetChangeType(value *DependencyGraphDiff_change_type)() {
    m.change_type = value
}
// SetEcosystem sets the ecosystem property value. The ecosystem property
func (m *DependencyGraphDiff) SetEcosystem(value *string)() {
    m.ecosystem = value
}
// SetLicense sets the license property value. The license property
func (m *DependencyGraphDiff) SetLicense(value *string)() {
    m.license = value
}
// SetManifest sets the manifest property value. The manifest property
func (m *DependencyGraphDiff) SetManifest(value *string)() {
    m.manifest = value
}
// SetName sets the name property value. The name property
func (m *DependencyGraphDiff) SetName(value *string)() {
    m.name = value
}
// SetPackageUrl sets the package_url property value. The package_url property
func (m *DependencyGraphDiff) SetPackageUrl(value *string)() {
    m.package_url = value
}
// SetScope sets the scope property value. Where the dependency is utilized. `development` means that the dependency is only utilized in the development environment. `runtime` means that the dependency is utilized at runtime and in the development environment.
func (m *DependencyGraphDiff) SetScope(value *DependencyGraphDiff_scope)() {
    m.scope = value
}
// SetSourceRepositoryUrl sets the source_repository_url property value. The source_repository_url property
func (m *DependencyGraphDiff) SetSourceRepositoryUrl(value *string)() {
    m.source_repository_url = value
}
// SetVersion sets the version property value. The version property
func (m *DependencyGraphDiff) SetVersion(value *string)() {
    m.version = value
}
// SetVulnerabilities sets the vulnerabilities property value. The vulnerabilities property
func (m *DependencyGraphDiff) SetVulnerabilities(value []DependencyGraphDiff_vulnerabilitiesable)() {
    m.vulnerabilities = value
}
type DependencyGraphDiffable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChangeType()(*DependencyGraphDiff_change_type)
    GetEcosystem()(*string)
    GetLicense()(*string)
    GetManifest()(*string)
    GetName()(*string)
    GetPackageUrl()(*string)
    GetScope()(*DependencyGraphDiff_scope)
    GetSourceRepositoryUrl()(*string)
    GetVersion()(*string)
    GetVulnerabilities()([]DependencyGraphDiff_vulnerabilitiesable)
    SetChangeType(value *DependencyGraphDiff_change_type)()
    SetEcosystem(value *string)()
    SetLicense(value *string)()
    SetManifest(value *string)()
    SetName(value *string)()
    SetPackageUrl(value *string)()
    SetScope(value *DependencyGraphDiff_scope)()
    SetSourceRepositoryUrl(value *string)()
    SetVersion(value *string)()
    SetVulnerabilities(value []DependencyGraphDiff_vulnerabilitiesable)()
}
