package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemIssuesItemLabelsPostRequestBodyMember2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The labels property
    labels []ItemItemIssuesItemLabelsPostRequestBodyMember2_labelsable
}
// NewItemItemIssuesItemLabelsPostRequestBodyMember2 instantiates a new ItemItemIssuesItemLabelsPostRequestBodyMember2 and sets the default values.
func NewItemItemIssuesItemLabelsPostRequestBodyMember2()(*ItemItemIssuesItemLabelsPostRequestBodyMember2) {
    m := &ItemItemIssuesItemLabelsPostRequestBodyMember2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemIssuesItemLabelsPostRequestBodyMember2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemIssuesItemLabelsPostRequestBodyMember2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemIssuesItemLabelsPostRequestBodyMember2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemIssuesItemLabelsPostRequestBodyMember2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemIssuesItemLabelsPostRequestBodyMember2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemItemIssuesItemLabelsPostRequestBodyMember2_labelsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemItemIssuesItemLabelsPostRequestBodyMember2_labelsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ItemItemIssuesItemLabelsPostRequestBodyMember2_labelsable)
                }
            }
            m.SetLabels(res)
        }
        return nil
    }
    return res
}
// GetLabels gets the labels property value. The labels property
// returns a []ItemItemIssuesItemLabelsPostRequestBodyMember2_labelsable when successful
func (m *ItemItemIssuesItemLabelsPostRequestBodyMember2) GetLabels()([]ItemItemIssuesItemLabelsPostRequestBodyMember2_labelsable) {
    return m.labels
}
// Serialize serializes information the current object
func (m *ItemItemIssuesItemLabelsPostRequestBodyMember2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemItemIssuesItemLabelsPostRequestBodyMember2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLabels sets the labels property value. The labels property
func (m *ItemItemIssuesItemLabelsPostRequestBodyMember2) SetLabels(value []ItemItemIssuesItemLabelsPostRequestBodyMember2_labelsable)() {
    m.labels = value
}
type ItemItemIssuesItemLabelsPostRequestBodyMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabels()([]ItemItemIssuesItemLabelsPostRequestBodyMember2_labelsable)
    SetLabels(value []ItemItemIssuesItemLabelsPostRequestBodyMember2_labelsable)()
}
