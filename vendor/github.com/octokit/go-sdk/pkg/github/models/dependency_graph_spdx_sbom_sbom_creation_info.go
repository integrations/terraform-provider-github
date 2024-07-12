package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DependencyGraphSpdxSbom_sbom_creationInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The date and time the SPDX document was created.
    created *string
    // The tools that were used to generate the SPDX document.
    creators []string
}
// NewDependencyGraphSpdxSbom_sbom_creationInfo instantiates a new DependencyGraphSpdxSbom_sbom_creationInfo and sets the default values.
func NewDependencyGraphSpdxSbom_sbom_creationInfo()(*DependencyGraphSpdxSbom_sbom_creationInfo) {
    m := &DependencyGraphSpdxSbom_sbom_creationInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDependencyGraphSpdxSbom_sbom_creationInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependencyGraphSpdxSbom_sbom_creationInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependencyGraphSpdxSbom_sbom_creationInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DependencyGraphSpdxSbom_sbom_creationInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreated gets the created property value. The date and time the SPDX document was created.
// returns a *string when successful
func (m *DependencyGraphSpdxSbom_sbom_creationInfo) GetCreated()(*string) {
    return m.created
}
// GetCreators gets the creators property value. The tools that were used to generate the SPDX document.
// returns a []string when successful
func (m *DependencyGraphSpdxSbom_sbom_creationInfo) GetCreators()([]string) {
    return m.creators
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependencyGraphSpdxSbom_sbom_creationInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["created"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreated(val)
        }
        return nil
    }
    res["creators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetCreators(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DependencyGraphSpdxSbom_sbom_creationInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("created", m.GetCreated())
        if err != nil {
            return err
        }
    }
    if m.GetCreators() != nil {
        err := writer.WriteCollectionOfStringValues("creators", m.GetCreators())
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
func (m *DependencyGraphSpdxSbom_sbom_creationInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreated sets the created property value. The date and time the SPDX document was created.
func (m *DependencyGraphSpdxSbom_sbom_creationInfo) SetCreated(value *string)() {
    m.created = value
}
// SetCreators sets the creators property value. The tools that were used to generate the SPDX document.
func (m *DependencyGraphSpdxSbom_sbom_creationInfo) SetCreators(value []string)() {
    m.creators = value
}
type DependencyGraphSpdxSbom_sbom_creationInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreated()(*string)
    GetCreators()([]string)
    SetCreated(value *string)()
    SetCreators(value []string)()
}
