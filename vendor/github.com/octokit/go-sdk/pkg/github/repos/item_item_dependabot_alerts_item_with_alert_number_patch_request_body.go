package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemDependabotAlertsItemWithAlert_numberPatchRequestBody struct {
    // An optional comment associated with dismissing the alert.
    dismissed_comment *string
}
// NewItemItemDependabotAlertsItemWithAlert_numberPatchRequestBody instantiates a new ItemItemDependabotAlertsItemWithAlert_numberPatchRequestBody and sets the default values.
func NewItemItemDependabotAlertsItemWithAlert_numberPatchRequestBody()(*ItemItemDependabotAlertsItemWithAlert_numberPatchRequestBody) {
    m := &ItemItemDependabotAlertsItemWithAlert_numberPatchRequestBody{
    }
    return m
}
// CreateItemItemDependabotAlertsItemWithAlert_numberPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemDependabotAlertsItemWithAlert_numberPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemDependabotAlertsItemWithAlert_numberPatchRequestBody(), nil
}
// GetDismissedComment gets the dismissed_comment property value. An optional comment associated with dismissing the alert.
// returns a *string when successful
func (m *ItemItemDependabotAlertsItemWithAlert_numberPatchRequestBody) GetDismissedComment()(*string) {
    return m.dismissed_comment
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemDependabotAlertsItemWithAlert_numberPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// Serialize serializes information the current object
func (m *ItemItemDependabotAlertsItemWithAlert_numberPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dismissed_comment", m.GetDismissedComment())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDismissedComment sets the dismissed_comment property value. An optional comment associated with dismissing the alert.
func (m *ItemItemDependabotAlertsItemWithAlert_numberPatchRequestBody) SetDismissedComment(value *string)() {
    m.dismissed_comment = value
}
type ItemItemDependabotAlertsItemWithAlert_numberPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDismissedComment()(*string)
    SetDismissedComment(value *string)()
}
