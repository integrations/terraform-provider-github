package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryAdvisoryUpdate_vulnerabilities struct {
    // The name of the package affected by the vulnerability.
    packageEscaped RepositoryAdvisoryUpdate_vulnerabilities_packageable
    // The package version(s) that resolve the vulnerability.
    patched_versions *string
    // The functions in the package that are affected.
    vulnerable_functions []string
    // The range of the package versions affected by the vulnerability.
    vulnerable_version_range *string
}
// NewRepositoryAdvisoryUpdate_vulnerabilities instantiates a new RepositoryAdvisoryUpdate_vulnerabilities and sets the default values.
func NewRepositoryAdvisoryUpdate_vulnerabilities()(*RepositoryAdvisoryUpdate_vulnerabilities) {
    m := &RepositoryAdvisoryUpdate_vulnerabilities{
    }
    return m
}
// CreateRepositoryAdvisoryUpdate_vulnerabilitiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryAdvisoryUpdate_vulnerabilitiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryAdvisoryUpdate_vulnerabilities(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryAdvisoryUpdate_vulnerabilities) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["package"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryAdvisoryUpdate_vulnerabilities_packageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageEscaped(val.(RepositoryAdvisoryUpdate_vulnerabilities_packageable))
        }
        return nil
    }
    res["patched_versions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPatchedVersions(val)
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
// GetPackageEscaped gets the package property value. The name of the package affected by the vulnerability.
// returns a RepositoryAdvisoryUpdate_vulnerabilities_packageable when successful
func (m *RepositoryAdvisoryUpdate_vulnerabilities) GetPackageEscaped()(RepositoryAdvisoryUpdate_vulnerabilities_packageable) {
    return m.packageEscaped
}
// GetPatchedVersions gets the patched_versions property value. The package version(s) that resolve the vulnerability.
// returns a *string when successful
func (m *RepositoryAdvisoryUpdate_vulnerabilities) GetPatchedVersions()(*string) {
    return m.patched_versions
}
// GetVulnerableFunctions gets the vulnerable_functions property value. The functions in the package that are affected.
// returns a []string when successful
func (m *RepositoryAdvisoryUpdate_vulnerabilities) GetVulnerableFunctions()([]string) {
    return m.vulnerable_functions
}
// GetVulnerableVersionRange gets the vulnerable_version_range property value. The range of the package versions affected by the vulnerability.
// returns a *string when successful
func (m *RepositoryAdvisoryUpdate_vulnerabilities) GetVulnerableVersionRange()(*string) {
    return m.vulnerable_version_range
}
// Serialize serializes information the current object
func (m *RepositoryAdvisoryUpdate_vulnerabilities) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("package", m.GetPackageEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("patched_versions", m.GetPatchedVersions())
        if err != nil {
            return err
        }
    }
    if m.GetVulnerableFunctions() != nil {
        err := writer.WriteCollectionOfStringValues("vulnerable_functions", m.GetVulnerableFunctions())
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
    return nil
}
// SetPackageEscaped sets the package property value. The name of the package affected by the vulnerability.
func (m *RepositoryAdvisoryUpdate_vulnerabilities) SetPackageEscaped(value RepositoryAdvisoryUpdate_vulnerabilities_packageable)() {
    m.packageEscaped = value
}
// SetPatchedVersions sets the patched_versions property value. The package version(s) that resolve the vulnerability.
func (m *RepositoryAdvisoryUpdate_vulnerabilities) SetPatchedVersions(value *string)() {
    m.patched_versions = value
}
// SetVulnerableFunctions sets the vulnerable_functions property value. The functions in the package that are affected.
func (m *RepositoryAdvisoryUpdate_vulnerabilities) SetVulnerableFunctions(value []string)() {
    m.vulnerable_functions = value
}
// SetVulnerableVersionRange sets the vulnerable_version_range property value. The range of the package versions affected by the vulnerability.
func (m *RepositoryAdvisoryUpdate_vulnerabilities) SetVulnerableVersionRange(value *string)() {
    m.vulnerable_version_range = value
}
type RepositoryAdvisoryUpdate_vulnerabilitiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPackageEscaped()(RepositoryAdvisoryUpdate_vulnerabilities_packageable)
    GetPatchedVersions()(*string)
    GetVulnerableFunctions()([]string)
    GetVulnerableVersionRange()(*string)
    SetPackageEscaped(value RepositoryAdvisoryUpdate_vulnerabilities_packageable)()
    SetPatchedVersions(value *string)()
    SetVulnerableFunctions(value []string)()
    SetVulnerableVersionRange(value *string)()
}
