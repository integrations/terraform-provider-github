package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeScanningSarifsStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The REST API URL for getting the analyses associated with the upload.
    analyses_url *string
    // Any errors that ocurred during processing of the delivery.
    errors []string
    // `pending` files have not yet been processed, while `complete` means results from the SARIF have been stored. `failed` files have either not been processed at all, or could only be partially processed.
    processing_status *CodeScanningSarifsStatus_processing_status
}
// NewCodeScanningSarifsStatus instantiates a new CodeScanningSarifsStatus and sets the default values.
func NewCodeScanningSarifsStatus()(*CodeScanningSarifsStatus) {
    m := &CodeScanningSarifsStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningSarifsStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningSarifsStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningSarifsStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningSarifsStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAnalysesUrl gets the analyses_url property value. The REST API URL for getting the analyses associated with the upload.
// returns a *string when successful
func (m *CodeScanningSarifsStatus) GetAnalysesUrl()(*string) {
    return m.analyses_url
}
// GetErrors gets the errors property value. Any errors that ocurred during processing of the delivery.
// returns a []string when successful
func (m *CodeScanningSarifsStatus) GetErrors()([]string) {
    return m.errors
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningSarifsStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["analyses_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnalysesUrl(val)
        }
        return nil
    }
    res["errors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetErrors(res)
        }
        return nil
    }
    res["processing_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningSarifsStatus_processing_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingStatus(val.(*CodeScanningSarifsStatus_processing_status))
        }
        return nil
    }
    return res
}
// GetProcessingStatus gets the processing_status property value. `pending` files have not yet been processed, while `complete` means results from the SARIF have been stored. `failed` files have either not been processed at all, or could only be partially processed.
// returns a *CodeScanningSarifsStatus_processing_status when successful
func (m *CodeScanningSarifsStatus) GetProcessingStatus()(*CodeScanningSarifsStatus_processing_status) {
    return m.processing_status
}
// Serialize serializes information the current object
func (m *CodeScanningSarifsStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetProcessingStatus() != nil {
        cast := (*m.GetProcessingStatus()).String()
        err := writer.WriteStringValue("processing_status", &cast)
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
func (m *CodeScanningSarifsStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAnalysesUrl sets the analyses_url property value. The REST API URL for getting the analyses associated with the upload.
func (m *CodeScanningSarifsStatus) SetAnalysesUrl(value *string)() {
    m.analyses_url = value
}
// SetErrors sets the errors property value. Any errors that ocurred during processing of the delivery.
func (m *CodeScanningSarifsStatus) SetErrors(value []string)() {
    m.errors = value
}
// SetProcessingStatus sets the processing_status property value. `pending` files have not yet been processed, while `complete` means results from the SARIF have been stored. `failed` files have either not been processed at all, or could only be partially processed.
func (m *CodeScanningSarifsStatus) SetProcessingStatus(value *CodeScanningSarifsStatus_processing_status)() {
    m.processing_status = value
}
type CodeScanningSarifsStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalysesUrl()(*string)
    GetErrors()([]string)
    GetProcessingStatus()(*CodeScanningSarifsStatus_processing_status)
    SetAnalysesUrl(value *string)()
    SetErrors(value []string)()
    SetProcessingStatus(value *CodeScanningSarifsStatus_processing_status)()
}
