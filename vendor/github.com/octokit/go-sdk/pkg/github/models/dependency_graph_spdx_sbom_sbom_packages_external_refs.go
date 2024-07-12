package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DependencyGraphSpdxSbom_sbom_packages_externalRefs struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The category of reference to an external resource this reference refers to.
    referenceCategory *string
    // A locator for the particular external resource this reference refers to.
    referenceLocator *string
    // The category of reference to an external resource this reference refers to.
    referenceType *string
}
// NewDependencyGraphSpdxSbom_sbom_packages_externalRefs instantiates a new DependencyGraphSpdxSbom_sbom_packages_externalRefs and sets the default values.
func NewDependencyGraphSpdxSbom_sbom_packages_externalRefs()(*DependencyGraphSpdxSbom_sbom_packages_externalRefs) {
    m := &DependencyGraphSpdxSbom_sbom_packages_externalRefs{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDependencyGraphSpdxSbom_sbom_packages_externalRefsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependencyGraphSpdxSbom_sbom_packages_externalRefsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependencyGraphSpdxSbom_sbom_packages_externalRefs(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DependencyGraphSpdxSbom_sbom_packages_externalRefs) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependencyGraphSpdxSbom_sbom_packages_externalRefs) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["referenceCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceCategory(val)
        }
        return nil
    }
    res["referenceLocator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceLocator(val)
        }
        return nil
    }
    res["referenceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceType(val)
        }
        return nil
    }
    return res
}
// GetReferenceCategory gets the referenceCategory property value. The category of reference to an external resource this reference refers to.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom_packages_externalRefs) GetReferenceCategory()(*string) {
    return m.referenceCategory
}
// GetReferenceLocator gets the referenceLocator property value. A locator for the particular external resource this reference refers to.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom_packages_externalRefs) GetReferenceLocator()(*string) {
    return m.referenceLocator
}
// GetReferenceType gets the referenceType property value. The category of reference to an external resource this reference refers to.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom_packages_externalRefs) GetReferenceType()(*string) {
    return m.referenceType
}
// Serialize serializes information the current object
func (m *DependencyGraphSpdxSbom_sbom_packages_externalRefs) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("referenceCategory", m.GetReferenceCategory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("referenceLocator", m.GetReferenceLocator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("referenceType", m.GetReferenceType())
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
func (m *DependencyGraphSpdxSbom_sbom_packages_externalRefs) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReferenceCategory sets the referenceCategory property value. The category of reference to an external resource this reference refers to.
func (m *DependencyGraphSpdxSbom_sbom_packages_externalRefs) SetReferenceCategory(value *string)() {
    m.referenceCategory = value
}
// SetReferenceLocator sets the referenceLocator property value. A locator for the particular external resource this reference refers to.
func (m *DependencyGraphSpdxSbom_sbom_packages_externalRefs) SetReferenceLocator(value *string)() {
    m.referenceLocator = value
}
// SetReferenceType sets the referenceType property value. The category of reference to an external resource this reference refers to.
func (m *DependencyGraphSpdxSbom_sbom_packages_externalRefs) SetReferenceType(value *string)() {
    m.referenceType = value
}
type DependencyGraphSpdxSbom_sbom_packages_externalRefsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReferenceCategory()(*string)
    GetReferenceLocator()(*string)
    GetReferenceType()(*string)
    SetReferenceCategory(value *string)()
    SetReferenceLocator(value *string)()
    SetReferenceType(value *string)()
}
