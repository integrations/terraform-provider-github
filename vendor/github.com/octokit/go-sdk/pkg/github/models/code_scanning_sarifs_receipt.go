package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeScanningSarifsReceipt struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An identifier for the upload.
    id *string
    // The REST API URL for checking the status of the upload.
    url *string
}
// NewCodeScanningSarifsReceipt instantiates a new CodeScanningSarifsReceipt and sets the default values.
func NewCodeScanningSarifsReceipt()(*CodeScanningSarifsReceipt) {
    m := &CodeScanningSarifsReceipt{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningSarifsReceiptFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningSarifsReceiptFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningSarifsReceipt(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningSarifsReceipt) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningSarifsReceipt) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. An identifier for the upload.
// returns a *string when successful
func (m *CodeScanningSarifsReceipt) GetId()(*string) {
    return m.id
}
// GetUrl gets the url property value. The REST API URL for checking the status of the upload.
// returns a *string when successful
func (m *CodeScanningSarifsReceipt) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *CodeScanningSarifsReceipt) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *CodeScanningSarifsReceipt) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. An identifier for the upload.
func (m *CodeScanningSarifsReceipt) SetId(value *string)() {
    m.id = value
}
// SetUrl sets the url property value. The REST API URL for checking the status of the upload.
func (m *CodeScanningSarifsReceipt) SetUrl(value *string)() {
    m.url = value
}
type CodeScanningSarifsReceiptable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*string)
    GetUrl()(*string)
    SetId(value *string)()
    SetUrl(value *string)()
}
