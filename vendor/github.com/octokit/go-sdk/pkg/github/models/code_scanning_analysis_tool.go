package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeScanningAnalysisTool struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The GUID of the tool used to generate the code scanning analysis, if provided in the uploaded SARIF data.
    guid *string
    // The name of the tool used to generate the code scanning analysis.
    name *string
    // The version of the tool used to generate the code scanning analysis.
    version *string
}
// NewCodeScanningAnalysisTool instantiates a new CodeScanningAnalysisTool and sets the default values.
func NewCodeScanningAnalysisTool()(*CodeScanningAnalysisTool) {
    m := &CodeScanningAnalysisTool{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningAnalysisToolFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningAnalysisToolFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningAnalysisTool(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningAnalysisTool) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningAnalysisTool) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["guid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuid(val)
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
    return res
}
// GetGuid gets the guid property value. The GUID of the tool used to generate the code scanning analysis, if provided in the uploaded SARIF data.
// returns a *string when successful
func (m *CodeScanningAnalysisTool) GetGuid()(*string) {
    return m.guid
}
// GetName gets the name property value. The name of the tool used to generate the code scanning analysis.
// returns a *string when successful
func (m *CodeScanningAnalysisTool) GetName()(*string) {
    return m.name
}
// GetVersion gets the version property value. The version of the tool used to generate the code scanning analysis.
// returns a *string when successful
func (m *CodeScanningAnalysisTool) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *CodeScanningAnalysisTool) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("guid", m.GetGuid())
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
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *CodeScanningAnalysisTool) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGuid sets the guid property value. The GUID of the tool used to generate the code scanning analysis, if provided in the uploaded SARIF data.
func (m *CodeScanningAnalysisTool) SetGuid(value *string)() {
    m.guid = value
}
// SetName sets the name property value. The name of the tool used to generate the code scanning analysis.
func (m *CodeScanningAnalysisTool) SetName(value *string)() {
    m.name = value
}
// SetVersion sets the version property value. The version of the tool used to generate the code scanning analysis.
func (m *CodeScanningAnalysisTool) SetVersion(value *string)() {
    m.version = value
}
type CodeScanningAnalysisToolable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGuid()(*string)
    GetName()(*string)
    GetVersion()(*string)
    SetGuid(value *string)()
    SetName(value *string)()
    SetVersion(value *string)()
}
