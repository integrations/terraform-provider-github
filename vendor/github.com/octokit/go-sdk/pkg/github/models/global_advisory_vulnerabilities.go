package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GlobalAdvisory_vulnerabilities struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The package version that resolve the vulnerability.
    first_patched_version *string
    // The name of the package affected by the vulnerability.
    packageEscaped GlobalAdvisory_vulnerabilities_packageable
    // The functions in the package that are affected by the vulnerability.
    vulnerable_functions []string
    // The range of the package versions affected by the vulnerability.
    vulnerable_version_range *string
}
// NewGlobalAdvisory_vulnerabilities instantiates a new GlobalAdvisory_vulnerabilities and sets the default values.
func NewGlobalAdvisory_vulnerabilities()(*GlobalAdvisory_vulnerabilities) {
    m := &GlobalAdvisory_vulnerabilities{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGlobalAdvisory_vulnerabilitiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGlobalAdvisory_vulnerabilitiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGlobalAdvisory_vulnerabilities(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GlobalAdvisory_vulnerabilities) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GlobalAdvisory_vulnerabilities) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["first_patched_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstPatchedVersion(val)
        }
        return nil
    }
    res["package"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGlobalAdvisory_vulnerabilities_packageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageEscaped(val.(GlobalAdvisory_vulnerabilities_packageable))
        }
        return nil
    }
    res["vulnerable_functions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetVulnerableFunctions(res)
        }
        return nil
    }
    res["vulnerable_version_range"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVulnerableVersionRange(val)
        }
        return nil
    }
    return res
}
// GetFirstPatchedVersion gets the first_patched_version property value. The package version that resolve the vulnerability.
// returns a *string when successful
func (m *GlobalAdvisory_vulnerabilities) GetFirstPatchedVersion()(*string) {
    return m.first_patched_version
}
// GetPackageEscaped gets the package property value. The name of the package affected by the vulnerability.
// returns a GlobalAdvisory_vulnerabilities_packageable when successful
func (m *GlobalAdvisory_vulnerabilities) GetPackageEscaped()(GlobalAdvisory_vulnerabilities_packageable) {
    return m.packageEscaped
}
// GetVulnerableFunctions gets the vulnerable_functions property value. The functions in the package that are affected by the vulnerability.
// returns a []string when successful
func (m *GlobalAdvisory_vulnerabilities) GetVulnerableFunctions()([]string) {
    return m.vulnerable_functions
}
// GetVulnerableVersionRange gets the vulnerable_version_range property value. The range of the package versions affected by the vulnerability.
// returns a *string when successful
func (m *GlobalAdvisory_vulnerabilities) GetVulnerableVersionRange()(*string) {
    return m.vulnerable_version_range
}
// Serialize serializes information the current object
func (m *GlobalAdvisory_vulnerabilities) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("first_patched_version", m.GetFirstPatchedVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("package", m.GetPackageEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vulnerable_version_range", m.GetVulnerableVersionRange())
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
func (m *GlobalAdvisory_vulnerabilities) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFirstPatchedVersion sets the first_patched_version property value. The package version that resolve the vulnerability.
func (m *GlobalAdvisory_vulnerabilities) SetFirstPatchedVersion(value *string)() {
    m.first_patched_version = value
}
// SetPackageEscaped sets the package property value. The name of the package affected by the vulnerability.
func (m *GlobalAdvisory_vulnerabilities) SetPackageEscaped(value GlobalAdvisory_vulnerabilities_packageable)() {
    m.packageEscaped = value
}
// SetVulnerableFunctions sets the vulnerable_functions property value. The functions in the package that are affected by the vulnerability.
func (m *GlobalAdvisory_vulnerabilities) SetVulnerableFunctions(value []string)() {
    m.vulnerable_functions = value
}
// SetVulnerableVersionRange sets the vulnerable_version_range property value. The range of the package versions affected by the vulnerability.
func (m *GlobalAdvisory_vulnerabilities) SetVulnerableVersionRange(value *string)() {
    m.vulnerable_version_range = value
}
type GlobalAdvisory_vulnerabilitiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFirstPatchedVersion()(*string)
    GetPackageEscaped()(GlobalAdvisory_vulnerabilities_packageable)
    GetVulnerableFunctions()([]string)
    GetVulnerableVersionRange()(*string)
    SetFirstPatchedVersion(value *string)()
    SetPackageEscaped(value GlobalAdvisory_vulnerabilities_packageable)()
    SetVulnerableFunctions(value []string)()
    SetVulnerableVersionRange(value *string)()
}
