package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\branches\{branch}\protection\restrictions\apps
type ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppsDeleteRequestBody composed type wrapper for classes ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able, string
type AppsDeleteRequestBody struct {
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able
    appsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able
    // Composed type representation for type string
    appsDeleteRequestBodyString *string
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able
    itemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able
    // Composed type representation for type string
    string *string
}
// NewAppsDeleteRequestBody instantiates a new AppsDeleteRequestBody and sets the default values.
func NewAppsDeleteRequestBody()(*AppsDeleteRequestBody) {
    m := &AppsDeleteRequestBody{
    }
    return m
}
// CreateAppsDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppsDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewAppsDeleteRequestBody()
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
        result.SetAppsDeleteRequestBodyString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able when successful
func (m *AppsDeleteRequestBody) GetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able) {
    return m.appsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1
}
// GetAppsDeleteRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *AppsDeleteRequestBody) GetAppsDeleteRequestBodyString()(*string) {
    return m.appsDeleteRequestBodyString
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AppsDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *AppsDeleteRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able when successful
func (m *AppsDeleteRequestBody) GetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able) {
    return m.itemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *AppsDeleteRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *AppsDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetAppsDeleteRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetAppsDeleteRequestBodyString())
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
// SetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able
func (m *AppsDeleteRequestBody) SetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able)() {
    m.appsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 = value
}
// SetAppsDeleteRequestBodyString sets the string property value. Composed type representation for type string
func (m *AppsDeleteRequestBody) SetAppsDeleteRequestBodyString(value *string)() {
    m.appsDeleteRequestBodyString = value
}
// SetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able
func (m *AppsDeleteRequestBody) SetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able)() {
    m.itemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *AppsDeleteRequestBody) SetString(value *string)() {
    m.string = value
}
// AppsPostRequestBody composed type wrapper for classes ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able, string
type AppsPostRequestBody struct {
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able
    appsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able
    // Composed type representation for type string
    appsPostRequestBodyString *string
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able
    itemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able
    // Composed type representation for type string
    string *string
}
// NewAppsPostRequestBody instantiates a new AppsPostRequestBody and sets the default values.
func NewAppsPostRequestBody()(*AppsPostRequestBody) {
    m := &AppsPostRequestBody{
    }
    return m
}
// CreateAppsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewAppsPostRequestBody()
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
        result.SetAppsPostRequestBodyString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able when successful
func (m *AppsPostRequestBody) GetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able) {
    return m.appsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1
}
// GetAppsPostRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *AppsPostRequestBody) GetAppsPostRequestBodyString()(*string) {
    return m.appsPostRequestBodyString
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AppsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *AppsPostRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able when successful
func (m *AppsPostRequestBody) GetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able) {
    return m.itemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *AppsPostRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *AppsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetAppsPostRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetAppsPostRequestBodyString())
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
// SetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able
func (m *AppsPostRequestBody) SetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able)() {
    m.appsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 = value
}
// SetAppsPostRequestBodyString sets the string property value. Composed type representation for type string
func (m *AppsPostRequestBody) SetAppsPostRequestBodyString(value *string)() {
    m.appsPostRequestBodyString = value
}
// SetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able
func (m *AppsPostRequestBody) SetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able)() {
    m.itemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *AppsPostRequestBody) SetString(value *string)() {
    m.string = value
}
// AppsPutRequestBody composed type wrapper for classes ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able, string
type AppsPutRequestBody struct {
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able
    appsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able
    // Composed type representation for type string
    appsPutRequestBodyString *string
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able
    itemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able
    // Composed type representation for type string
    string *string
}
// NewAppsPutRequestBody instantiates a new AppsPutRequestBody and sets the default values.
func NewAppsPutRequestBody()(*AppsPutRequestBody) {
    m := &AppsPutRequestBody{
    }
    return m
}
// CreateAppsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewAppsPutRequestBody()
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
        result.SetAppsPutRequestBodyString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able when successful
func (m *AppsPutRequestBody) GetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able) {
    return m.appsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1
}
// GetAppsPutRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *AppsPutRequestBody) GetAppsPutRequestBodyString()(*string) {
    return m.appsPutRequestBodyString
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AppsPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *AppsPutRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able when successful
func (m *AppsPutRequestBody) GetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able) {
    return m.itemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *AppsPutRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *AppsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetAppsPutRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetAppsPutRequestBodyString())
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
// SetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able
func (m *AppsPutRequestBody) SetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able)() {
    m.appsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 = value
}
// SetAppsPutRequestBodyString sets the string property value. Composed type representation for type string
func (m *AppsPutRequestBody) SetAppsPutRequestBodyString(value *string)() {
    m.appsPutRequestBodyString = value
}
// SetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able
func (m *AppsPutRequestBody) SetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able)() {
    m.itemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *AppsPutRequestBody) SetString(value *string)() {
    m.string = value
}
type AppsDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able)
    GetAppsDeleteRequestBodyString()(*string)
    GetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able)
    GetString()(*string)
    SetAppsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able)()
    SetAppsDeleteRequestBodyString(value *string)()
    SetItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsDeleteRequestBodyMember1able)()
    SetString(value *string)()
}
type AppsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able)
    GetAppsPostRequestBodyString()(*string)
    GetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able)
    GetString()(*string)
    SetAppsPostRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able)()
    SetAppsPostRequestBodyString(value *string)()
    SetItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPostRequestBodyMember1able)()
    SetString(value *string)()
}
type AppsPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able)
    GetAppsPutRequestBodyString()(*string)
    GetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able)
    GetString()(*string)
    SetAppsPutRequestBodyItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able)()
    SetAppsPutRequestBodyString(value *string)()
    SetItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsAppsPutRequestBodyMember1able)()
    SetString(value *string)()
}
// NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderInternal instantiates a new ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) {
    m := &ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/branches/{branch}/protection/restrictions/apps", pathParameters),
    }
    return m
}
// NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder instantiates a new ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Removes the ability of an app to push to this branch. Only GitHub Apps that are installed on the repository and that have been granted write access to the repository contents can be added as authorized actors on a protected branch.
// returns a []Integrationable when successful
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#remove-app-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) Delete(ctx context.Context, body AppsDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Integrationable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateIntegrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Integrationable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Integrationable)
        }
    }
    return val, nil
}
// Get protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists the GitHub Apps that have push access to this branch. Only GitHub Apps that are installed on the repository and that have been granted write access to the repository contents can be added as authorized actors on a protected branch.
// returns a []Integrationable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#get-apps-with-access-to-the-protected-branch
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Integrationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateIntegrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Integrationable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Integrationable)
        }
    }
    return val, nil
}
// Post protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Grants the specified apps push access for this branch. Only GitHub Apps that are installed on the repository and that have been granted write access to the repository contents can be added as authorized actors on a protected branch.
// returns a []Integrationable when successful
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#add-app-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) Post(ctx context.Context, body AppsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Integrationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateIntegrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Integrationable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Integrationable)
        }
    }
    return val, nil
}
// Put protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Replaces the list of apps that have push access to this branch. This removes all apps that previously had push access and grants push access to the new list of apps. Only GitHub Apps that are installed on the repository and that have been granted write access to the repository contents can be added as authorized actors on a protected branch.
// returns a []Integrationable when successful
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#set-app-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) Put(ctx context.Context, body AppsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Integrationable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateIntegrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Integrationable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Integrationable)
        }
    }
    return val, nil
}
// ToDeleteRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Removes the ability of an app to push to this branch. Only GitHub Apps that are installed on the repository and that have been granted write access to the repository contents can be added as authorized actors on a protected branch.
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body AppsDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToGetRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists the GitHub Apps that have push access to this branch. Only GitHub Apps that are installed on the repository and that have been granted write access to the repository contents can be added as authorized actors on a protected branch.
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Grants the specified apps push access for this branch. Only GitHub Apps that are installed on the repository and that have been granted write access to the repository contents can be added as authorized actors on a protected branch.
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, body AppsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToPutRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Replaces the list of apps that have push access to this branch. This removes all apps that previously had push access and grants push access to the new list of apps. Only GitHub Apps that are installed on the repository and that have been granted write access to the repository contents can be added as authorized actors on a protected branch.
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) ToPutRequestInformation(ctx context.Context, body AppsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder when successful
func (m *ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) WithUrl(rawUrl string)(*ItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder) {
    return NewItemItemBranchesItemProtectionRestrictionsAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
