package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemIssuesItemLabelsPutRequestBodyMember2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The labels property
    labels []ItemItemIssuesItemLabelsPutRequestBodyMember2_labelsable
}
// NewItemItemIssuesItemLabelsPutRequestBodyMember2 instantiates a new ItemItemIssuesItemLabelsPutRequestBodyMember2 and sets the default values.
func NewItemItemIssuesItemLabelsPutRequestBodyMember2()(*ItemItemIssuesItemLabelsPutRequestBodyMember2) {
    m := &ItemItemIssuesItemLabelsPutRequestBodyMember2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemIssuesItemLabelsPutRequestBodyMember2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemIssuesItemLabelsPutRequestBodyMember2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemIssuesItemLabelsPutRequestBodyMember2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemIssuesItemLabelsPutRequestBodyMember2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemIssuesItemLabelsPutRequestBodyMember2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemItemIssuesItemLabelsPutRequestBodyMember2_labelsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemItemIssuesItemLabelsPutRequestBodyMember2_labelsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ItemItemIssuesItemLabelsPutRequestBodyMember2_labelsable)
                }
            }
            m.SetLabels(res)
        }
        return nil
    }
    return res
}
// GetLabels gets the labels property value. The labels property
// returns a []ItemItemIssuesItemLabelsPutRequestBodyMember2_labelsable when successful
func (m *ItemItemIssuesItemLabelsPutRequestBodyMember2) GetLabels()([]ItemItemIssuesItemLabelsPutRequestBodyMember2_labelsable) {
    return m.labels
}
// Serialize serializes information the current object
func (m *ItemItemIssuesItemLabelsPutRequestBodyMember2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLabels()))
        for i, v := range m.GetLabels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("labels", cast)
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
func (m *ItemItemIssuesItemLabelsPutRequestBodyMember2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLabels sets the labels property value. The labels property
func (m *ItemItemIssuesItemLabelsPutRequestBodyMember2) SetLabels(value []ItemItemIssuesItemLabelsPutRequestBodyMember2_labelsable)() {
    m.labels = value
}
type ItemItemIssuesItemLabelsPutRequestBodyMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabels()([]ItemItemIssuesItemLabelsPutRequestBodyMember2_labelsable)
    SetLabels(value []ItemItemIssuesItemLabelsPutRequestBodyMember2_labelsable)()
}
