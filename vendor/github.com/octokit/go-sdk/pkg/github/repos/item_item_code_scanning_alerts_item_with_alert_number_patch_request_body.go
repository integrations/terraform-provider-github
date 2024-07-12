package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The dismissal comment associated with the dismissal of the alert.
    dismissed_comment *string
    // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
    dismissed_reason *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAlertDismissedReason
    // Sets the state of the code scanning alert. You must provide `dismissed_reason` when you set the state to `dismissed`.
    state *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAlertSetState
}
// NewItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody instantiates a new ItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody and sets the default values.
func NewItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody()(*ItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody) {
    m := &ItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDismissedComment gets the dismissed_comment property value. The dismissal comment associated with the dismissal of the alert.
// returns a *string when successful
func (m *ItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody) GetDismissedComment()(*string) {
    return m.dismissed_comment
}
// GetDismissedReason gets the dismissed_reason property value. **Required when the state is dismissed.** The reason for dismissing or closing the alert.
// returns a *CodeScanningAlertDismissedReason when successful
func (m *ItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody) GetDismissedReason()(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAlertDismissedReason) {
    return m.dismissed_reason
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dismissed_comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissedComment(val)
        }
        return nil
    }
    res["dismissed_reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ParseCodeScanningAlertDismissedReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDismissedReason(val.(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAlertDismissedReason))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ParseCodeScanningAlertSetState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAlertSetState))
        }
        return nil
    }
    return res
}
// GetState gets the state property value. Sets the state of the code scanning alert. You must provide `dismissed_reason` when you set the state to `dismissed`.
// returns a *CodeScanningAlertSetState when successful
func (m *ItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody) GetState()(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAlertSetState) {
    return m.state
}
// Serialize serializes information the current object
func (m *ItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dismissed_comment", m.GetDismissedComment())
        if err != nil {
            return err
        }
    }
    if m.GetDismissedReason() != nil {
        cast := (*m.GetDismissedReason()).String()
        err := writer.WriteStringValue("dismissed_reason", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *ItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDismissedComment sets the dismissed_comment property value. The dismissal comment associated with the dismissal of the alert.
func (m *ItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody) SetDismissedComment(value *string)() {
    m.dismissed_comment = value
}
// SetDismissedReason sets the dismissed_reason property value. **Required when the state is dismissed.** The reason for dismissing or closing the alert.
func (m *ItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody) SetDismissedReason(value *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAlertDismissedReason)() {
    m.dismissed_reason = value
}
// SetState sets the state property value. Sets the state of the code scanning alert. You must provide `dismissed_reason` when you set the state to `dismissed`.
func (m *ItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBody) SetState(value *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAlertSetState)() {
    m.state = value
}
type ItemItemCodeScanningAlertsItemWithAlert_numberPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDismissedComment()(*string)
    GetDismissedReason()(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAlertDismissedReason)
    GetState()(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAlertSetState)
    SetDismissedComment(value *string)()
    SetDismissedReason(value *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAlertDismissedReason)()
    SetState(value *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAlertSetState)()
}
