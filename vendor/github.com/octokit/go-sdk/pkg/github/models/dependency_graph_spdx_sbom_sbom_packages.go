package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DependencyGraphSpdxSbom_sbom_packages struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The copyright holders of the package, and any dates present with those notices, if available.
    copyrightText *string
    // The location where the package can be downloaded,or NOASSERTION if this has not been determined.
    downloadLocation *string
    // The externalRefs property
    externalRefs []DependencyGraphSpdxSbom_sbom_packages_externalRefsable
    // Whether the package's file content has been subjected toanalysis during the creation of the SPDX document.
    filesAnalyzed *bool
    // The license of the package as determined while creating the SPDX document.
    licenseConcluded *string
    // The license of the package as declared by its author, or NOASSERTION if this informationwas not available when the SPDX document was created.
    licenseDeclared *string
    // The name of the package.
    name *string
    // A unique SPDX identifier for the package.
    sPDXID *string
    // The distribution source of this package, or NOASSERTION if this was not determined.
    supplier *string
    // The version of the package. If the package does not have an exact version specified,a version range is given.
    versionInfo *string
}
// NewDependencyGraphSpdxSbom_sbom_packages instantiates a new DependencyGraphSpdxSbom_sbom_packages and sets the default values.
func NewDependencyGraphSpdxSbom_sbom_packages()(*DependencyGraphSpdxSbom_sbom_packages) {
    m := &DependencyGraphSpdxSbom_sbom_packages{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDependencyGraphSpdxSbom_sbom_packagesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependencyGraphSpdxSbom_sbom_packagesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependencyGraphSpdxSbom_sbom_packages(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DependencyGraphSpdxSbom_sbom_packages) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCopyrightText gets the copyrightText property value. The copyright holders of the package, and any dates present with those notices, if available.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom_packages) GetCopyrightText()(*string) {
    return m.copyrightText
}
// GetDownloadLocation gets the downloadLocation property value. The location where the package can be downloaded,or NOASSERTION if this has not been determined.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom_packages) GetDownloadLocation()(*string) {
    return m.downloadLocation
}
// GetExternalRefs gets the externalRefs property value. The externalRefs property
// returns a []DependencyGraphSpdxSbom_sbom_packages_externalRefsable when successful
func (m *DependencyGraphSpdxSbom_sbom_packages) GetExternalRefs()([]DependencyGraphSpdxSbom_sbom_packages_externalRefsable) {
    return m.externalRefs
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependencyGraphSpdxSbom_sbom_packages) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["copyrightText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopyrightText(val)
        }
        return nil
    }
    res["downloadLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadLocation(val)
        }
        return nil
    }
    res["externalRefs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDependencyGraphSpdxSbom_sbom_packages_externalRefsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DependencyGraphSpdxSbom_sbom_packages_externalRefsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DependencyGraphSpdxSbom_sbom_packages_externalRefsable)
                }
            }
            m.SetExternalRefs(res)
        }
        return nil
    }
    res["filesAnalyzed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilesAnalyzed(val)
        }
        return nil
    }
    res["licenseConcluded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseConcluded(val)
        }
        return nil
    }
    res["licenseDeclared"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseDeclared(val)
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
    res["SPDXID"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSPDXID(val)
        }
        return nil
    }
    res["supplier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupplier(val)
        }
        return nil
    }
    res["versionInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionInfo(val)
        }
        return nil
    }
    return res
}
// GetFilesAnalyzed gets the filesAnalyzed property value. Whether the package's file content has been subjected toanalysis during the creation of the SPDX document.
// returns a *bool when successful
func (m *DependencyGraphSpdxSbom_sbom_packages) GetFilesAnalyzed()(*bool) {
    return m.filesAnalyzed
}
// GetLicenseConcluded gets the licenseConcluded property value. The license of the package as determined while creating the SPDX document.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom_packages) GetLicenseConcluded()(*string) {
    return m.licenseConcluded
}
// GetLicenseDeclared gets the licenseDeclared property value. The license of the package as declared by its author, or NOASSERTION if this informationwas not available when the SPDX document was created.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom_packages) GetLicenseDeclared()(*string) {
    return m.licenseDeclared
}
// GetName gets the name property value. The name of the package.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom_packages) GetName()(*string) {
    return m.name
}
// GetSPDXID gets the SPDXID property value. A unique SPDX identifier for the package.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom_packages) GetSPDXID()(*string) {
    return m.sPDXID
}
// GetSupplier gets the supplier property value. The distribution source of this package, or NOASSERTION if this was not determined.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom_packages) GetSupplier()(*string) {
    return m.supplier
}
// GetVersionInfo gets the versionInfo property value. The version of the package. If the package does not have an exact version specified,a version range is given.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom_packages) GetVersionInfo()(*string) {
    return m.versionInfo
}
// Serialize serializes information the current object
func (m *DependencyGraphSpdxSbom_sbom_packages) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("copyrightText", m.GetCopyrightText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("downloadLocation", m.GetDownloadLocation())
        if err != nil {
            return err
        }
    }
    if m.GetExternalRefs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExternalRefs()))
        for i, v := range m.GetExternalRefs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("externalRefs", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("filesAnalyzed", m.GetFilesAnalyzed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("licenseConcluded", m.GetLicenseConcluded())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("licenseDeclared", m.GetLicenseDeclared())
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
        err := writer.WriteStringValue("SPDXID", m.GetSPDXID())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("supplier", m.GetSupplier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("versionInfo", m.GetVersionInfo())
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
func (m *DependencyGraphSpdxSbom_sbom_packages) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCopyrightText sets the copyrightText property value. The copyright holders of the package, and any dates present with those notices, if available.
func (m *DependencyGraphSpdxSbom_sbom_packages) SetCopyrightText(value *string)() {
    m.copyrightText = value
}
// SetDownloadLocation sets the downloadLocation property value. The location where the package can be downloaded,or NOASSERTION if this has not been determined.
func (m *DependencyGraphSpdxSbom_sbom_packages) SetDownloadLocation(value *string)() {
    m.downloadLocation = value
}
// SetExternalRefs sets the externalRefs property value. The externalRefs property
func (m *DependencyGraphSpdxSbom_sbom_packages) SetExternalRefs(value []DependencyGraphSpdxSbom_sbom_packages_externalRefsable)() {
    m.externalRefs = value
}
// SetFilesAnalyzed sets the filesAnalyzed property value. Whether the package's file content has been subjected toanalysis during the creation of the SPDX document.
func (m *DependencyGraphSpdxSbom_sbom_packages) SetFilesAnalyzed(value *bool)() {
    m.filesAnalyzed = value
}
// SetLicenseConcluded sets the licenseConcluded property value. The license of the package as determined while creating the SPDX document.
func (m *DependencyGraphSpdxSbom_sbom_packages) SetLicenseConcluded(value *string)() {
    m.licenseConcluded = value
}
// SetLicenseDeclared sets the licenseDeclared property value. The license of the package as declared by its author, or NOASSERTION if this informationwas not available when the SPDX document was created.
func (m *DependencyGraphSpdxSbom_sbom_packages) SetLicenseDeclared(value *string)() {
    m.licenseDeclared = value
}
// SetName sets the name property value. The name of the package.
func (m *DependencyGraphSpdxSbom_sbom_packages) SetName(value *string)() {
    m.name = value
}
// SetSPDXID sets the SPDXID property value. A unique SPDX identifier for the package.
func (m *DependencyGraphSpdxSbom_sbom_packages) SetSPDXID(value *string)() {
    m.sPDXID = value
}
// SetSupplier sets the supplier property value. The distribution source of this package, or NOASSERTION if this was not determined.
func (m *DependencyGraphSpdxSbom_sbom_packages) SetSupplier(value *string)() {
    m.supplier = value
}
// SetVersionInfo sets the versionInfo property value. The version of the package. If the package does not have an exact version specified,a version range is given.
func (m *DependencyGraphSpdxSbom_sbom_packages) SetVersionInfo(value *string)() {
    m.versionInfo = value
}
type DependencyGraphSpdxSbom_sbom_packagesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCopyrightText()(*string)
    GetDownloadLocation()(*string)
    GetExternalRefs()([]DependencyGraphSpdxSbom_sbom_packages_externalRefsable)
    GetFilesAnalyzed()(*bool)
    GetLicenseConcluded()(*string)
    GetLicenseDeclared()(*string)
    GetName()(*string)
    GetSPDXID()(*string)
    GetSupplier()(*string)
    GetVersionInfo()(*string)
    SetCopyrightText(value *string)()
    SetDownloadLocation(value *string)()
    SetExternalRefs(value []DependencyGraphSpdxSbom_sbom_packages_externalRefsable)()
    SetFilesAnalyzed(value *bool)()
    SetLicenseConcluded(value *string)()
    SetLicenseDeclared(value *string)()
    SetName(value *string)()
    SetSPDXID(value *string)()
    SetSupplier(value *string)()
    SetVersionInfo(value *string)()
}
