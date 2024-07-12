package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\branches\{branch}\protection\required_status_checks\contexts
type ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ContextsDeleteRequestBody composed type wrapper for classes ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able, string
type ContextsDeleteRequestBody struct {
    // Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able
    contextsDeleteRequestBodyItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1 ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able
    // Composed type representation for type string
    contextsDeleteRequestBodyString *string
    // Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able
    itemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1 ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able
    // Composed type representation for type string
    string *string
}
// NewContextsDeleteRequestBody instantiates a new ContextsDeleteRequestBody and sets the default values.
func NewContextsDeleteRequestBody()(*ContextsDeleteRequestBody) {
    m := &ContextsDeleteRequestBody{
    }
    return m
}
// CreateContextsDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContextsDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewContextsDeleteRequestBody()
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
        result.SetContextsDeleteRequestBodyString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetContextsDeleteRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1 gets the ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able when successful
func (m *ContextsDeleteRequestBody) GetContextsDeleteRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able) {
    return m.contextsDeleteRequestBodyItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1
}
// GetContextsDeleteRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ContextsDeleteRequestBody) GetContextsDeleteRequestBodyString()(*string) {
    return m.contextsDeleteRequestBodyString
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContextsDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ContextsDeleteRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1 gets the ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able when successful
func (m *ContextsDeleteRequestBody) GetItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able) {
    return m.itemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ContextsDeleteRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ContextsDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContextsDeleteRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetContextsDeleteRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetContextsDeleteRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetContextsDeleteRequestBodyString())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContextsDeleteRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1 sets the ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able
func (m *ContextsDeleteRequestBody) SetContextsDeleteRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able)() {
    m.contextsDeleteRequestBodyItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1 = value
}
// SetContextsDeleteRequestBodyString sets the string property value. Composed type representation for type string
func (m *ContextsDeleteRequestBody) SetContextsDeleteRequestBodyString(value *string)() {
    m.contextsDeleteRequestBodyString = value
}
// SetItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1 sets the ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able
func (m *ContextsDeleteRequestBody) SetItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able)() {
    m.itemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ContextsDeleteRequestBody) SetString(value *string)() {
    m.string = value
}
// ContextsPostRequestBody composed type wrapper for classes ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able, string
type ContextsPostRequestBody struct {
    // Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able
    contextsPostRequestBodyItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1 ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able
    // Composed type representation for type string
    contextsPostRequestBodyString *string
    // Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able
    itemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1 ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able
    // Composed type representation for type string
    string *string
}
// NewContextsPostRequestBody instantiates a new ContextsPostRequestBody and sets the default values.
func NewContextsPostRequestBody()(*ContextsPostRequestBody) {
    m := &ContextsPostRequestBody{
    }
    return m
}
// CreateContextsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContextsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewContextsPostRequestBody()
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
        result.SetContextsPostRequestBodyString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetContextsPostRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1 gets the ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able when successful
func (m *ContextsPostRequestBody) GetContextsPostRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able) {
    return m.contextsPostRequestBodyItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1
}
// GetContextsPostRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ContextsPostRequestBody) GetContextsPostRequestBodyString()(*string) {
    return m.contextsPostRequestBodyString
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContextsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ContextsPostRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1 gets the ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able when successful
func (m *ContextsPostRequestBody) GetItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able) {
    return m.itemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ContextsPostRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ContextsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContextsPostRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetContextsPostRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetContextsPostRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetContextsPostRequestBodyString())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContextsPostRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1 sets the ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able
func (m *ContextsPostRequestBody) SetContextsPostRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able)() {
    m.contextsPostRequestBodyItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1 = value
}
// SetContextsPostRequestBodyString sets the string property value. Composed type representation for type string
func (m *ContextsPostRequestBody) SetContextsPostRequestBodyString(value *string)() {
    m.contextsPostRequestBodyString = value
}
// SetItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1 sets the ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able
func (m *ContextsPostRequestBody) SetItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able)() {
    m.itemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ContextsPostRequestBody) SetString(value *string)() {
    m.string = value
}
// ContextsPutRequestBody composed type wrapper for classes ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able, string
type ContextsPutRequestBody struct {
    // Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able
    contextsPutRequestBodyItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1 ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able
    // Composed type representation for type string
    contextsPutRequestBodyString *string
    // Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able
    itemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1 ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able
    // Composed type representation for type string
    string *string
}
// NewContextsPutRequestBody instantiates a new ContextsPutRequestBody and sets the default values.
func NewContextsPutRequestBody()(*ContextsPutRequestBody) {
    m := &ContextsPutRequestBody{
    }
    return m
}
// CreateContextsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContextsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewContextsPutRequestBody()
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
        result.SetContextsPutRequestBodyString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetContextsPutRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1 gets the ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able when successful
func (m *ContextsPutRequestBody) GetContextsPutRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able) {
    return m.contextsPutRequestBodyItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1
}
// GetContextsPutRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ContextsPutRequestBody) GetContextsPutRequestBodyString()(*string) {
    return m.contextsPutRequestBodyString
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContextsPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *ContextsPutRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1 gets the ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able when successful
func (m *ContextsPutRequestBody) GetItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able) {
    return m.itemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *ContextsPutRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *ContextsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContextsPutRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetContextsPutRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetContextsPutRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetContextsPutRequestBodyString())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContextsPutRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1 sets the ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able
func (m *ContextsPutRequestBody) SetContextsPutRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able)() {
    m.contextsPutRequestBodyItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1 = value
}
// SetContextsPutRequestBodyString sets the string property value. Composed type representation for type string
func (m *ContextsPutRequestBody) SetContextsPutRequestBodyString(value *string)() {
    m.contextsPutRequestBodyString = value
}
// SetItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1 sets the ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able
func (m *ContextsPutRequestBody) SetItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able)() {
    m.itemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *ContextsPutRequestBody) SetString(value *string)() {
    m.string = value
}
type ContextsDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContextsDeleteRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able)
    GetContextsDeleteRequestBodyString()(*string)
    GetItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able)
    GetString()(*string)
    SetContextsDeleteRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able)()
    SetContextsDeleteRequestBodyString(value *string)()
    SetItemItemBranchesItemProtectionRequiredStatusChecksContextsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsDeleteRequestBodyMember1able)()
    SetString(value *string)()
}
type ContextsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContextsPostRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able)
    GetContextsPostRequestBodyString()(*string)
    GetItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able)
    GetString()(*string)
    SetContextsPostRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able)()
    SetContextsPostRequestBodyString(value *string)()
    SetItemItemBranchesItemProtectionRequiredStatusChecksContextsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able)()
    SetString(value *string)()
}
type ContextsPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContextsPutRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able)
    GetContextsPutRequestBodyString()(*string)
    GetItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able)
    GetString()(*string)
    SetContextsPutRequestBodyItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able)()
    SetContextsPutRequestBodyString(value *string)()
    SetItemItemBranchesItemProtectionRequiredStatusChecksContextsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRequired_status_checksContextsPutRequestBodyMember1able)()
    SetString(value *string)()
}
// NewItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderInternal instantiates a new ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) {
    m := &ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/branches/{branch}/protection/required_status_checks/contexts", pathParameters),
    }
    return m
}
// NewItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder instantiates a new ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// returns a []string when successful
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#remove-status-check-contexts
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) Delete(ctx context.Context, body ContextsDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]string, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitiveCollection(ctx, requestInfo, "string", errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]string, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = *(v.(*string))
        }
    }
    return val, nil
}
// Get protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// returns a []string when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#get-all-status-check-contexts
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]string, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitiveCollection(ctx, requestInfo, "string", errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]string, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = *(v.(*string))
        }
    }
    return val, nil
}
// Post protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// returns a []string when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#add-status-check-contexts
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) Post(ctx context.Context, body ContextsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]string, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitiveCollection(ctx, requestInfo, "string", errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]string, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = *(v.(*string))
        }
    }
    return val, nil
}
// Put protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// returns a []string when successful
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#set-status-check-contexts
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) Put(ctx context.Context, body ContextsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]string, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitiveCollection(ctx, requestInfo, "string", errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]string, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = *(v.(*string))
        }
    }
    return val, nil
}
// ToDeleteRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body ContextsDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToGetRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ContextsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToPutRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) ToPutRequestInformation(ctx context.Context, body ContextsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder when successful
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) WithUrl(rawUrl string)(*ItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder) {
    return NewItemItemBranchesItemProtectionRequired_status_checksContextsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
