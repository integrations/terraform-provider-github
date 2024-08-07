package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DependencyGraphSpdxSbom_sbom struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The creationInfo property
    creationInfo DependencyGraphSpdxSbom_sbom_creationInfoable
    // The license under which the SPDX document is licensed.
    dataLicense *string
    // The name of the repository that the SPDX document describes.
    documentDescribes []string
    // The namespace for the SPDX document.
    documentNamespace *string
    // The name of the SPDX document.
    name *string
    // The packages property
    packages []DependencyGraphSpdxSbom_sbom_packagesable
    // The SPDX identifier for the SPDX document.
    sPDXID *string
    // The version of the SPDX specification that this document conforms to.
    spdxVersion *string
}
// NewDependencyGraphSpdxSbom_sbom instantiates a new DependencyGraphSpdxSbom_sbom and sets the default values.
func NewDependencyGraphSpdxSbom_sbom()(*DependencyGraphSpdxSbom_sbom) {
    m := &DependencyGraphSpdxSbom_sbom{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDependencyGraphSpdxSbom_sbomFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependencyGraphSpdxSbom_sbomFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependencyGraphSpdxSbom_sbom(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DependencyGraphSpdxSbom_sbom) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreationInfo gets the creationInfo property value. The creationInfo property
// returns a DependencyGraphSpdxSbom_sbom_creationInfoable when successful
func (m *DependencyGraphSpdxSbom_sbom) GetCreationInfo()(DependencyGraphSpdxSbom_sbom_creationInfoable) {
    return m.creationInfo
}
// GetDataLicense gets the dataLicense property value. The license under which the SPDX document is licensed.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom) GetDataLicense()(*string) {
    return m.dataLicense
}
// GetDocumentDescribes gets the documentDescribes property value. The name of the repository that the SPDX document describes.
// returns a []string when successful
func (m *DependencyGraphSpdxSbom_sbom) GetDocumentDescribes()([]string) {
    return m.documentDescribes
}
// GetDocumentNamespace gets the documentNamespace property value. The namespace for the SPDX document.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom) GetDocumentNamespace()(*string) {
    return m.documentNamespace
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependencyGraphSpdxSbom_sbom) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["creationInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDependencyGraphSpdxSbom_sbom_creationInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationInfo(val.(DependencyGraphSpdxSbom_sbom_creationInfoable))
        }
        return nil
    }
    res["dataLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataLicense(val)
        }
        return nil
    }
    res["documentDescribes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetDocumentDescribes(res)
        }
        return nil
    }
    res["documentNamespace"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentNamespace(val)
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
    res["packages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDependencyGraphSpdxSbom_sbom_packagesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DependencyGraphSpdxSbom_sbom_packagesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DependencyGraphSpdxSbom_sbom_packagesable)
                }
            }
            m.SetPackages(res)
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
    res["spdxVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpdxVersion(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the SPDX document.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom) GetName()(*string) {
    return m.name
}
// GetPackages gets the packages property value. The packages property
// returns a []DependencyGraphSpdxSbom_sbom_packagesable when successful
func (m *DependencyGraphSpdxSbom_sbom) GetPackages()([]DependencyGraphSpdxSbom_sbom_packagesable) {
    return m.packages
}
// GetSPDXID gets the SPDXID property value. The SPDX identifier for the SPDX document.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom) GetSPDXID()(*string) {
    return m.sPDXID
}
// GetSpdxVersion gets the spdxVersion property value. The version of the SPDX specification that this document conforms to.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom) GetSpdxVersion()(*string) {
    return m.spdxVersion
}
// Serialize serializes information the current object
func (m *DependencyGraphSpdxSbom_sbom) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("creationInfo", m.GetCreationInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dataLicense", m.GetDataLicense())
        if err != nil {
            return err
        }
    }
    if m.GetDocumentDescribes() != nil {
        err := writer.WriteCollectionOfStringValues("documentDescribes", m.GetDocumentDescribes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("documentNamespace", m.GetDocumentNamespace())
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
    if m.GetPackages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPackages()))
        for i, v := range m.GetPackages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("packages", cast)
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
        err := writer.WriteStringValue("spdxVersion", m.GetSpdxVersion())
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
func (m *DependencyGraphSpdxSbom_sbom) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreationInfo sets the creationInfo property value. The creationInfo property
func (m *DependencyGraphSpdxSbom_sbom) SetCreationInfo(value DependencyGraphSpdxSbom_sbom_creationInfoable)() {
    m.creationInfo = value
}
// SetDataLicense sets the dataLicense property value. The license under which the SPDX document is licensed.
func (m *DependencyGraphSpdxSbom_sbom) SetDataLicense(value *string)() {
    m.dataLicense = value
}
// SetDocumentDescribes sets the documentDescribes property value. The name of the repository that the SPDX document describes.
func (m *DependencyGraphSpdxSbom_sbom) SetDocumentDescribes(value []string)() {
    m.documentDescribes = value
}
// SetDocumentNamespace sets the documentNamespace property value. The namespace for the SPDX document.
func (m *DependencyGraphSpdxSbom_sbom) SetDocumentNamespace(value *string)() {
    m.documentNamespace = value
}
// SetName sets the name property value. The name of the SPDX document.
func (m *DependencyGraphSpdxSbom_sbom) SetName(value *string)() {
    m.name = value
}
// SetPackages sets the packages property value. The packages property
func (m *DependencyGraphSpdxSbom_sbom) SetPackages(value []DependencyGraphSpdxSbom_sbom_packagesable)() {
    m.packages = value
}
// SetSPDXID sets the SPDXID property value. The SPDX identifier for the SPDX document.
func (m *DependencyGraphSpdxSbom_sbom) SetSPDXID(value *string)() {
    m.sPDXID = value
}
// SetSpdxVersion sets the spdxVersion property value. The version of the SPDX specification that this document conforms to.
func (m *DependencyGraphSpdxSbom_sbom) SetSpdxVersion(value *string)() {
    m.spdxVersion = value
}
type DependencyGraphSpdxSbom_sbomable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreationInfo()(DependencyGraphSpdxSbom_sbom_creationInfoable)
    GetDataLicense()(*string)
    GetDocumentDescribes()([]string)
    GetDocumentNamespace()(*string)
    GetName()(*string)
    GetPackages()([]DependencyGraphSpdxSbom_sbom_packagesable)
    GetSPDXID()(*string)
    GetSpdxVersion()(*string)
    SetCreationInfo(value DependencyGraphSpdxSbom_sbom_creationInfoable)()
    SetDataLicense(value *string)()
    SetDocumentDescribes(value []string)()
    SetDocumentNamespace(value *string)()
    SetName(value *string)()
    SetPackages(value []DependencyGraphSpdxSbom_sbom_packagesable)()
    SetSPDXID(value *string)()
    SetSpdxVersion(value *string)()
}
