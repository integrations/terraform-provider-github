package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeScanningDefaultSetup503Error struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ApiError
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The code property
    code *string
    // The documentation_url property
    documentation_url *string
    // The message property
    message *string
}
// NewCodeScanningDefaultSetup503Error instantiates a new CodeScanningDefaultSetup503Error and sets the default values.
func NewCodeScanningDefaultSetup503Error()(*CodeScanningDefaultSetup503Error) {
    m := &CodeScanningDefaultSetup503Error{
        ApiError: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewApiError(),
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningDefaultSetup503ErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningDefaultSetup503ErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningDefaultSetup503Error(), nil
}
// Error the primary error message.
// returns a string when successful
func (m *CodeScanningDefaultSetup503Error) Error()(string) {
    return m.ApiError.Error()
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningDefaultSetup503Error) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. The code property
// returns a *string when successful
func (m *CodeScanningDefaultSetup503Error) GetCode()(*string) {
    return m.code
}
// GetDocumentationUrl gets the documentation_url property value. The documentation_url property
// returns a *string when successful
func (m *CodeScanningDefaultSetup503Error) GetDocumentationUrl()(*string) {
    return m.documentation_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningDefaultSetup503Error) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["documentation_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentationUrl(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *CodeScanningDefaultSetup503Error) GetMessage()(*string) {
    return m.message
}
// Serialize serializes information the current object
func (m *CodeScanningDefaultSetup503Error) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("documentation_url", m.GetDocumentationUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
func (m *CodeScanningDefaultSetup503Error) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. The code property
func (m *CodeScanningDefaultSetup503Error) SetCode(value *string)() {
    m.code = value
}
// SetDocumentationUrl sets the documentation_url property value. The documentation_url property
func (m *CodeScanningDefaultSetup503Error) SetDocumentationUrl(value *string)() {
    m.documentation_url = value
}
// SetMessage sets the message property value. The message property
func (m *CodeScanningDefaultSetup503Error) SetMessage(value *string)() {
    m.message = value
}
type CodeScanningDefaultSetup503Errorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetDocumentationUrl()(*string)
    GetMessage()(*string)
    SetCode(value *string)()
    SetDocumentationUrl(value *string)()
    SetMessage(value *string)()
}
