package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependencyGraphSpdxSbom a schema for the SPDX JSON format returned by the Dependency Graph.
type DependencyGraphSpdxSbom struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The sbom property
    sbom DependencyGraphSpdxSbom_sbomable
}
// NewDependencyGraphSpdxSbom instantiates a new DependencyGraphSpdxSbom and sets the default values.
func NewDependencyGraphSpdxSbom()(*DependencyGraphSpdxSbom) {
    m := &DependencyGraphSpdxSbom{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDependencyGraphSpdxSbomFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependencyGraphSpdxSbomFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependencyGraphSpdxSbom(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DependencyGraphSpdxSbom) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependencyGraphSpdxSbom) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sbom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDependencyGraphSpdxSbom_sbomFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSbom(val.(DependencyGraphSpdxSbom_sbomable))
        }
        return nil
    }
    return res
}
// GetSbom gets the sbom property value. The sbom property
// returns a DependencyGraphSpdxSbom_sbomable when successful
func (m *DependencyGraphSpdxSbom) GetSbom()(DependencyGraphSpdxSbom_sbomable) {
    return m.sbom
}
// Serialize serializes information the current object
func (m *DependencyGraphSpdxSbom) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("sbom", m.GetSbom())
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
func (m *DependencyGraphSpdxSbom) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSbom sets the sbom property value. The sbom property
func (m *DependencyGraphSpdxSbom) SetSbom(value DependencyGraphSpdxSbom_sbomable)() {
    m.sbom = value
}
type DependencyGraphSpdxSbomable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSbom()(DependencyGraphSpdxSbom_sbomable)
    SetSbom(value DependencyGraphSpdxSbom_sbomable)()
}
