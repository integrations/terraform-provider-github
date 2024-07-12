package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemIssuesItemLabelsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\issues\{issue_number}\labels
type ItemItemIssuesItemLabelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemIssuesItemLabelsRequestBuilderGetQueryParameters lists all labels for an issue.
type ItemItemIssuesItemLabelsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// LabelsPostRequestBody composed type wrapper for classes ItemItemIssuesItemLabelsPostRequestBodyMember1able, ItemItemIssuesItemLabelsPostRequestBodyMember2able, string, []ItemItemIssuesItemLabelsPostRequestBodyMember3able
type LabelsPostRequestBody struct {
    // Composed type representation for type ItemItemIssuesItemLabelsPostRequestBodyMember1able
    itemItemIssuesItemLabelsPostRequestBodyMember1 ItemItemIssuesItemLabelsPostRequestBodyMember1able
    // Composed type representation for type ItemItemIssuesItemLabelsPostRequestBodyMember2able
    itemItemIssuesItemLabelsPostRequestBodyMember2 ItemItemIssuesItemLabelsPostRequestBodyMember2able
    // Composed type representation for type []ItemItemIssuesItemLabelsPostRequestBodyMember3able
    itemItemIssuesItemLabelsPostRequestBodyMember3 []ItemItemIssuesItemLabelsPostRequestBodyMember3able
    // Composed type representation for type ItemItemIssuesItemLabelsPostRequestBodyMember1able
    labelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember1 ItemItemIssuesItemLabelsPostRequestBodyMember1able
    // Composed type representation for type ItemItemIssuesItemLabelsPostRequestBodyMember2able
    labelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember2 ItemItemIssuesItemLabelsPostRequestBodyMember2able
    // Composed type representation for type []ItemItemIssuesItemLabelsPostRequestBodyMember3able
    labelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember3 []ItemItemIssuesItemLabelsPostRequestBodyMember3able
    // Composed type representation for type string
    labelsPostRequestBodyString *string
    // Composed type representation for type string
    string *string
}
// NewLabelsPostRequestBody instantiates a new LabelsPostRequestBody and sets the default values.
func NewLabelsPostRequestBody()(*LabelsPostRequestBody) {
    m := &LabelsPostRequestBody{
    }
    return m
}
// CreateLabelsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewLabelsPostRequestBody()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetLabelsPostRequestBodyString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    } else if val, err := parseNode.GetCollectionOfObjectValues(CreateItemItemIssuesItemLabelsPostRequestBodyMember3FromDiscriminatorValue); val != nil {
        if err != nil {
            return nil, err
        }
        cast := make([]ItemItemIssuesItemLabelsPostRequestBodyMember3able, len(val))
        for i, v := range val {
            if v != nil {
                cast[i] = v.(ItemItemIssuesItemLabelsPostRequestBodyMember3able)
            }
        }
        result.SetItemItemIssuesItemLabelsPostRequestBodyMember3(cast)
    } else if val, err := parseNode.GetCollectionOfObjectValues(CreateItemItemIssuesItemLabelsPostRequestBodyMember3FromDiscriminatorValue); val != nil {
        if err != nil {
            return nil, err
        }
        cast := make([]ItemItemIssuesItemLabelsPostRequestBodyMember3able, len(val))
        for i, v := range val {
            if v != nil {
                cast[i] = v.(ItemItemIssuesItemLabelsPostRequestBodyMember3able)
            }
        }
        result.SetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember3(cast)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LabelsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *LabelsPostRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemIssuesItemLabelsPostRequestBodyMember1 gets the ItemItemIssuesItemLabelsPostRequestBodyMember1 property value. Composed type representation for type ItemItemIssuesItemLabelsPostRequestBodyMember1able
// returns a ItemItemIssuesItemLabelsPostRequestBodyMember1able when successful
func (m *LabelsPostRequestBody) GetItemItemIssuesItemLabelsPostRequestBodyMember1()(ItemItemIssuesItemLabelsPostRequestBodyMember1able) {
    return m.itemItemIssuesItemLabelsPostRequestBodyMember1
}
// GetItemItemIssuesItemLabelsPostRequestBodyMember2 gets the ItemItemIssuesItemLabelsPostRequestBodyMember2 property value. Composed type representation for type ItemItemIssuesItemLabelsPostRequestBodyMember2able
// returns a ItemItemIssuesItemLabelsPostRequestBodyMember2able when successful
func (m *LabelsPostRequestBody) GetItemItemIssuesItemLabelsPostRequestBodyMember2()(ItemItemIssuesItemLabelsPostRequestBodyMember2able) {
    return m.itemItemIssuesItemLabelsPostRequestBodyMember2
}
// GetItemItemIssuesItemLabelsPostRequestBodyMember3 gets the ItemItemIssuesItemLabelsPostRequestBodyMember3 property value. Composed type representation for type []ItemItemIssuesItemLabelsPostRequestBodyMember3able
// returns a []ItemItemIssuesItemLabelsPostRequestBodyMember3able when successful
func (m *LabelsPostRequestBody) GetItemItemIssuesItemLabelsPostRequestBodyMember3()([]ItemItemIssuesItemLabelsPostRequestBodyMember3able) {
    return m.itemItemIssuesItemLabelsPostRequestBodyMember3
}
// GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember1 gets the ItemItemIssuesItemLabelsPostRequestBodyMember1 property value. Composed type representation for type ItemItemIssuesItemLabelsPostRequestBodyMember1able
// returns a ItemItemIssuesItemLabelsPostRequestBodyMember1able when successful
func (m *LabelsPostRequestBody) GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember1()(ItemItemIssuesItemLabelsPostRequestBodyMember1able) {
    return m.labelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember1
}
// GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember2 gets the ItemItemIssuesItemLabelsPostRequestBodyMember2 property value. Composed type representation for type ItemItemIssuesItemLabelsPostRequestBodyMember2able
// returns a ItemItemIssuesItemLabelsPostRequestBodyMember2able when successful
func (m *LabelsPostRequestBody) GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember2()(ItemItemIssuesItemLabelsPostRequestBodyMember2able) {
    return m.labelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember2
}
// GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember3 gets the ItemItemIssuesItemLabelsPostRequestBodyMember3 property value. Composed type representation for type []ItemItemIssuesItemLabelsPostRequestBodyMember3able
// returns a []ItemItemIssuesItemLabelsPostRequestBodyMember3able when successful
func (m *LabelsPostRequestBody) GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember3()([]ItemItemIssuesItemLabelsPostRequestBodyMember3able) {
    return m.labelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember3
}
// GetLabelsPostRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *LabelsPostRequestBody) GetLabelsPostRequestBodyString()(*string) {
    return m.labelsPostRequestBodyString
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *LabelsPostRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *LabelsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItemItemIssuesItemLabelsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemIssuesItemLabelsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetItemItemIssuesItemLabelsPostRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemIssuesItemLabelsPostRequestBodyMember2())
        if err != nil {
            return err
        }
    } else if m.GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember2())
        if err != nil {
            return err
        }
    } else if m.GetLabelsPostRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetLabelsPostRequestBodyString())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    } else if m.GetItemItemIssuesItemLabelsPostRequestBodyMember3() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItemItemIssuesItemLabelsPostRequestBodyMember3()))
        for i, v := range m.GetItemItemIssuesItemLabelsPostRequestBodyMember3() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("", cast)
        if err != nil {
            return err
        }
    } else if m.GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember3() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember3()))
        for i, v := range m.GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember3() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemItemIssuesItemLabelsPostRequestBodyMember1 sets the ItemItemIssuesItemLabelsPostRequestBodyMember1 property value. Composed type representation for type ItemItemIssuesItemLabelsPostRequestBodyMember1able
func (m *LabelsPostRequestBody) SetItemItemIssuesItemLabelsPostRequestBodyMember1(value ItemItemIssuesItemLabelsPostRequestBodyMember1able)() {
    m.itemItemIssuesItemLabelsPostRequestBodyMember1 = value
}
// SetItemItemIssuesItemLabelsPostRequestBodyMember2 sets the ItemItemIssuesItemLabelsPostRequestBodyMember2 property value. Composed type representation for type ItemItemIssuesItemLabelsPostRequestBodyMember2able
func (m *LabelsPostRequestBody) SetItemItemIssuesItemLabelsPostRequestBodyMember2(value ItemItemIssuesItemLabelsPostRequestBodyMember2able)() {
    m.itemItemIssuesItemLabelsPostRequestBodyMember2 = value
}
// SetItemItemIssuesItemLabelsPostRequestBodyMember3 sets the ItemItemIssuesItemLabelsPostRequestBodyMember3 property value. Composed type representation for type []ItemItemIssuesItemLabelsPostRequestBodyMember3able
func (m *LabelsPostRequestBody) SetItemItemIssuesItemLabelsPostRequestBodyMember3(value []ItemItemIssuesItemLabelsPostRequestBodyMember3able)() {
    m.itemItemIssuesItemLabelsPostRequestBodyMember3 = value
}
// SetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember1 sets the ItemItemIssuesItemLabelsPostRequestBodyMember1 property value. Composed type representation for type ItemItemIssuesItemLabelsPostRequestBodyMember1able
func (m *LabelsPostRequestBody) SetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember1(value ItemItemIssuesItemLabelsPostRequestBodyMember1able)() {
    m.labelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember1 = value
}
// SetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember2 sets the ItemItemIssuesItemLabelsPostRequestBodyMember2 property value. Composed type representation for type ItemItemIssuesItemLabelsPostRequestBodyMember2able
func (m *LabelsPostRequestBody) SetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember2(value ItemItemIssuesItemLabelsPostRequestBodyMember2able)() {
    m.labelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember2 = value
}
// SetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember3 sets the ItemItemIssuesItemLabelsPostRequestBodyMember3 property value. Composed type representation for type []ItemItemIssuesItemLabelsPostRequestBodyMember3able
func (m *LabelsPostRequestBody) SetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember3(value []ItemItemIssuesItemLabelsPostRequestBodyMember3able)() {
    m.labelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember3 = value
}
// SetLabelsPostRequestBodyString sets the string property value. Composed type representation for type string
func (m *LabelsPostRequestBody) SetLabelsPostRequestBodyString(value *string)() {
    m.labelsPostRequestBodyString = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *LabelsPostRequestBody) SetString(value *string)() {
    m.string = value
}
// LabelsPutRequestBody composed type wrapper for classes ItemItemIssuesItemLabelsPutRequestBodyMember1able, ItemItemIssuesItemLabelsPutRequestBodyMember2able, string, []ItemItemIssuesItemLabelsPutRequestBodyMember3able
type LabelsPutRequestBody struct {
    // Composed type representation for type ItemItemIssuesItemLabelsPutRequestBodyMember1able
    itemItemIssuesItemLabelsPutRequestBodyMember1 ItemItemIssuesItemLabelsPutRequestBodyMember1able
    // Composed type representation for type ItemItemIssuesItemLabelsPutRequestBodyMember2able
    itemItemIssuesItemLabelsPutRequestBodyMember2 ItemItemIssuesItemLabelsPutRequestBodyMember2able
    // Composed type representation for type []ItemItemIssuesItemLabelsPutRequestBodyMember3able
    itemItemIssuesItemLabelsPutRequestBodyMember3 []ItemItemIssuesItemLabelsPutRequestBodyMember3able
    // Composed type representation for type ItemItemIssuesItemLabelsPutRequestBodyMember1able
    labelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember1 ItemItemIssuesItemLabelsPutRequestBodyMember1able
    // Composed type representation for type ItemItemIssuesItemLabelsPutRequestBodyMember2able
    labelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember2 ItemItemIssuesItemLabelsPutRequestBodyMember2able
    // Composed type representation for type []ItemItemIssuesItemLabelsPutRequestBodyMember3able
    labelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember3 []ItemItemIssuesItemLabelsPutRequestBodyMember3able
    // Composed type representation for type string
    labelsPutRequestBodyString *string
    // Composed type representation for type string
    string *string
}
// NewLabelsPutRequestBody instantiates a new LabelsPutRequestBody and sets the default values.
func NewLabelsPutRequestBody()(*LabelsPutRequestBody) {
    m := &LabelsPutRequestBody{
    }
    return m
}
// CreateLabelsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLabelsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewLabelsPutRequestBody()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetLabelsPutRequestBodyString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    } else if val, err := parseNode.GetCollectionOfObjectValues(CreateItemItemIssuesItemLabelsPutRequestBodyMember3FromDiscriminatorValue); val != nil {
        if err != nil {
            return nil, err
        }
        cast := make([]ItemItemIssuesItemLabelsPutRequestBodyMember3able, len(val))
        for i, v := range val {
            if v != nil {
                cast[i] = v.(ItemItemIssuesItemLabelsPutRequestBodyMember3able)
            }
        }
        result.SetItemItemIssuesItemLabelsPutRequestBodyMember3(cast)
    } else if val, err := parseNode.GetCollectionOfObjectValues(CreateItemItemIssuesItemLabelsPutRequestBodyMember3FromDiscriminatorValue); val != nil {
        if err != nil {
            return nil, err
        }
        cast := make([]ItemItemIssuesItemLabelsPutRequestBodyMember3able, len(val))
        for i, v := range val {
            if v != nil {
                cast[i] = v.(ItemItemIssuesItemLabelsPutRequestBodyMember3able)
            }
        }
        result.SetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember3(cast)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LabelsPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *LabelsPutRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemIssuesItemLabelsPutRequestBodyMember1 gets the ItemItemIssuesItemLabelsPutRequestBodyMember1 property value. Composed type representation for type ItemItemIssuesItemLabelsPutRequestBodyMember1able
// returns a ItemItemIssuesItemLabelsPutRequestBodyMember1able when successful
func (m *LabelsPutRequestBody) GetItemItemIssuesItemLabelsPutRequestBodyMember1()(ItemItemIssuesItemLabelsPutRequestBodyMember1able) {
    return m.itemItemIssuesItemLabelsPutRequestBodyMember1
}
// GetItemItemIssuesItemLabelsPutRequestBodyMember2 gets the ItemItemIssuesItemLabelsPutRequestBodyMember2 property value. Composed type representation for type ItemItemIssuesItemLabelsPutRequestBodyMember2able
// returns a ItemItemIssuesItemLabelsPutRequestBodyMember2able when successful
func (m *LabelsPutRequestBody) GetItemItemIssuesItemLabelsPutRequestBodyMember2()(ItemItemIssuesItemLabelsPutRequestBodyMember2able) {
    return m.itemItemIssuesItemLabelsPutRequestBodyMember2
}
// GetItemItemIssuesItemLabelsPutRequestBodyMember3 gets the ItemItemIssuesItemLabelsPutRequestBodyMember3 property value. Composed type representation for type []ItemItemIssuesItemLabelsPutRequestBodyMember3able
// returns a []ItemItemIssuesItemLabelsPutRequestBodyMember3able when successful
func (m *LabelsPutRequestBody) GetItemItemIssuesItemLabelsPutRequestBodyMember3()([]ItemItemIssuesItemLabelsPutRequestBodyMember3able) {
    return m.itemItemIssuesItemLabelsPutRequestBodyMember3
}
// GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember1 gets the ItemItemIssuesItemLabelsPutRequestBodyMember1 property value. Composed type representation for type ItemItemIssuesItemLabelsPutRequestBodyMember1able
// returns a ItemItemIssuesItemLabelsPutRequestBodyMember1able when successful
func (m *LabelsPutRequestBody) GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember1()(ItemItemIssuesItemLabelsPutRequestBodyMember1able) {
    return m.labelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember1
}
// GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember2 gets the ItemItemIssuesItemLabelsPutRequestBodyMember2 property value. Composed type representation for type ItemItemIssuesItemLabelsPutRequestBodyMember2able
// returns a ItemItemIssuesItemLabelsPutRequestBodyMember2able when successful
func (m *LabelsPutRequestBody) GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember2()(ItemItemIssuesItemLabelsPutRequestBodyMember2able) {
    return m.labelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember2
}
// GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember3 gets the ItemItemIssuesItemLabelsPutRequestBodyMember3 property value. Composed type representation for type []ItemItemIssuesItemLabelsPutRequestBodyMember3able
// returns a []ItemItemIssuesItemLabelsPutRequestBodyMember3able when successful
func (m *LabelsPutRequestBody) GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember3()([]ItemItemIssuesItemLabelsPutRequestBodyMember3able) {
    return m.labelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember3
}
// GetLabelsPutRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *LabelsPutRequestBody) GetLabelsPutRequestBodyString()(*string) {
    return m.labelsPutRequestBodyString
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *LabelsPutRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *LabelsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItemItemIssuesItemLabelsPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemIssuesItemLabelsPutRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetItemItemIssuesItemLabelsPutRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemIssuesItemLabelsPutRequestBodyMember2())
        if err != nil {
            return err
        }
    } else if m.GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember2())
        if err != nil {
            return err
        }
    } else if m.GetLabelsPutRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetLabelsPutRequestBodyString())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    } else if m.GetItemItemIssuesItemLabelsPutRequestBodyMember3() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItemItemIssuesItemLabelsPutRequestBodyMember3()))
        for i, v := range m.GetItemItemIssuesItemLabelsPutRequestBodyMember3() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("", cast)
        if err != nil {
            return err
        }
    } else if m.GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember3() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember3()))
        for i, v := range m.GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember3() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemItemIssuesItemLabelsPutRequestBodyMember1 sets the ItemItemIssuesItemLabelsPutRequestBodyMember1 property value. Composed type representation for type ItemItemIssuesItemLabelsPutRequestBodyMember1able
func (m *LabelsPutRequestBody) SetItemItemIssuesItemLabelsPutRequestBodyMember1(value ItemItemIssuesItemLabelsPutRequestBodyMember1able)() {
    m.itemItemIssuesItemLabelsPutRequestBodyMember1 = value
}
// SetItemItemIssuesItemLabelsPutRequestBodyMember2 sets the ItemItemIssuesItemLabelsPutRequestBodyMember2 property value. Composed type representation for type ItemItemIssuesItemLabelsPutRequestBodyMember2able
func (m *LabelsPutRequestBody) SetItemItemIssuesItemLabelsPutRequestBodyMember2(value ItemItemIssuesItemLabelsPutRequestBodyMember2able)() {
    m.itemItemIssuesItemLabelsPutRequestBodyMember2 = value
}
// SetItemItemIssuesItemLabelsPutRequestBodyMember3 sets the ItemItemIssuesItemLabelsPutRequestBodyMember3 property value. Composed type representation for type []ItemItemIssuesItemLabelsPutRequestBodyMember3able
func (m *LabelsPutRequestBody) SetItemItemIssuesItemLabelsPutRequestBodyMember3(value []ItemItemIssuesItemLabelsPutRequestBodyMember3able)() {
    m.itemItemIssuesItemLabelsPutRequestBodyMember3 = value
}
// SetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember1 sets the ItemItemIssuesItemLabelsPutRequestBodyMember1 property value. Composed type representation for type ItemItemIssuesItemLabelsPutRequestBodyMember1able
func (m *LabelsPutRequestBody) SetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember1(value ItemItemIssuesItemLabelsPutRequestBodyMember1able)() {
    m.labelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember1 = value
}
// SetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember2 sets the ItemItemIssuesItemLabelsPutRequestBodyMember2 property value. Composed type representation for type ItemItemIssuesItemLabelsPutRequestBodyMember2able
func (m *LabelsPutRequestBody) SetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember2(value ItemItemIssuesItemLabelsPutRequestBodyMember2able)() {
    m.labelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember2 = value
}
// SetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember3 sets the ItemItemIssuesItemLabelsPutRequestBodyMember3 property value. Composed type representation for type []ItemItemIssuesItemLabelsPutRequestBodyMember3able
func (m *LabelsPutRequestBody) SetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember3(value []ItemItemIssuesItemLabelsPutRequestBodyMember3able)() {
    m.labelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember3 = value
}
// SetLabelsPutRequestBodyString sets the string property value. Composed type representation for type string
func (m *LabelsPutRequestBody) SetLabelsPutRequestBodyString(value *string)() {
    m.labelsPutRequestBodyString = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *LabelsPutRequestBody) SetString(value *string)() {
    m.string = value
}
type LabelsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemItemIssuesItemLabelsPostRequestBodyMember1()(ItemItemIssuesItemLabelsPostRequestBodyMember1able)
    GetItemItemIssuesItemLabelsPostRequestBodyMember2()(ItemItemIssuesItemLabelsPostRequestBodyMember2able)
    GetItemItemIssuesItemLabelsPostRequestBodyMember3()([]ItemItemIssuesItemLabelsPostRequestBodyMember3able)
    GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember1()(ItemItemIssuesItemLabelsPostRequestBodyMember1able)
    GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember2()(ItemItemIssuesItemLabelsPostRequestBodyMember2able)
    GetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember3()([]ItemItemIssuesItemLabelsPostRequestBodyMember3able)
    GetLabelsPostRequestBodyString()(*string)
    GetString()(*string)
    SetItemItemIssuesItemLabelsPostRequestBodyMember1(value ItemItemIssuesItemLabelsPostRequestBodyMember1able)()
    SetItemItemIssuesItemLabelsPostRequestBodyMember2(value ItemItemIssuesItemLabelsPostRequestBodyMember2able)()
    SetItemItemIssuesItemLabelsPostRequestBodyMember3(value []ItemItemIssuesItemLabelsPostRequestBodyMember3able)()
    SetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember1(value ItemItemIssuesItemLabelsPostRequestBodyMember1able)()
    SetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember2(value ItemItemIssuesItemLabelsPostRequestBodyMember2able)()
    SetLabelsPostRequestBodyItemItemIssuesItemLabelsPostRequestBodyMember3(value []ItemItemIssuesItemLabelsPostRequestBodyMember3able)()
    SetLabelsPostRequestBodyString(value *string)()
    SetString(value *string)()
}
type LabelsPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemItemIssuesItemLabelsPutRequestBodyMember1()(ItemItemIssuesItemLabelsPutRequestBodyMember1able)
    GetItemItemIssuesItemLabelsPutRequestBodyMember2()(ItemItemIssuesItemLabelsPutRequestBodyMember2able)
    GetItemItemIssuesItemLabelsPutRequestBodyMember3()([]ItemItemIssuesItemLabelsPutRequestBodyMember3able)
    GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember1()(ItemItemIssuesItemLabelsPutRequestBodyMember1able)
    GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember2()(ItemItemIssuesItemLabelsPutRequestBodyMember2able)
    GetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember3()([]ItemItemIssuesItemLabelsPutRequestBodyMember3able)
    GetLabelsPutRequestBodyString()(*string)
    GetString()(*string)
    SetItemItemIssuesItemLabelsPutRequestBodyMember1(value ItemItemIssuesItemLabelsPutRequestBodyMember1able)()
    SetItemItemIssuesItemLabelsPutRequestBodyMember2(value ItemItemIssuesItemLabelsPutRequestBodyMember2able)()
    SetItemItemIssuesItemLabelsPutRequestBodyMember3(value []ItemItemIssuesItemLabelsPutRequestBodyMember3able)()
    SetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember1(value ItemItemIssuesItemLabelsPutRequestBodyMember1able)()
    SetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember2(value ItemItemIssuesItemLabelsPutRequestBodyMember2able)()
    SetLabelsPutRequestBodyItemItemIssuesItemLabelsPutRequestBodyMember3(value []ItemItemIssuesItemLabelsPutRequestBodyMember3able)()
    SetLabelsPutRequestBodyString(value *string)()
    SetString(value *string)()
}
// ByName gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.issues.item.labels.item collection
// returns a *ItemItemIssuesItemLabelsWithNameItemRequestBuilder when successful
func (m *ItemItemIssuesItemLabelsRequestBuilder) ByName(name string)(*ItemItemIssuesItemLabelsWithNameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if name != "" {
        urlTplParams["name"] = name
    }
    return NewItemItemIssuesItemLabelsWithNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemIssuesItemLabelsRequestBuilderInternal instantiates a new ItemItemIssuesItemLabelsRequestBuilder and sets the default values.
func NewItemItemIssuesItemLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesItemLabelsRequestBuilder) {
    m := &ItemItemIssuesItemLabelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/issues/{issue_number}/labels{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemItemIssuesItemLabelsRequestBuilder instantiates a new ItemItemIssuesItemLabelsRequestBuilder and sets the default values.
func NewItemItemIssuesItemLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemIssuesItemLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemIssuesItemLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete removes all labels from an issue.
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 410 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/labels#remove-all-labels-from-an-issue
func (m *ItemItemIssuesItemLabelsRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "410": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get lists all labels for an issue.
// returns a []Labelable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 410 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/labels#list-labels-for-an-issue
func (m *ItemItemIssuesItemLabelsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemIssuesItemLabelsRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Labelable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "410": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Labelable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Labelable)
        }
    }
    return val, nil
}
// Post adds labels to an issue. If you provide an empty array of labels, all labels are removed from the issue. 
// returns a []Labelable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 410 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/labels#add-labels-to-an-issue
func (m *ItemItemIssuesItemLabelsRequestBuilder) Post(ctx context.Context, body LabelsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Labelable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "410": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Labelable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Labelable)
        }
    }
    return val, nil
}
// Put removes any previous labels and sets the new labels for an issue.
// returns a []Labelable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 410 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/issues/labels#set-labels-for-an-issue
func (m *ItemItemIssuesItemLabelsRequestBuilder) Put(ctx context.Context, body LabelsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Labelable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "410": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Labelable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Labelable)
        }
    }
    return val, nil
}
// ToDeleteRequestInformation removes all labels from an issue.
// returns a *RequestInformation when successful
func (m *ItemItemIssuesItemLabelsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation lists all labels for an issue.
// returns a *RequestInformation when successful
func (m *ItemItemIssuesItemLabelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemIssuesItemLabelsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation adds labels to an issue. If you provide an empty array of labels, all labels are removed from the issue. 
// returns a *RequestInformation when successful
func (m *ItemItemIssuesItemLabelsRequestBuilder) ToPostRequestInformation(ctx context.Context, body LabelsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToPutRequestInformation removes any previous labels and sets the new labels for an issue.
// returns a *RequestInformation when successful
func (m *ItemItemIssuesItemLabelsRequestBuilder) ToPutRequestInformation(ctx context.Context, body LabelsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemIssuesItemLabelsRequestBuilder when successful
func (m *ItemItemIssuesItemLabelsRequestBuilder) WithUrl(rawUrl string)(*ItemItemIssuesItemLabelsRequestBuilder) {
    return NewItemItemIssuesItemLabelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
