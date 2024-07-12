package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\branches\{branch}\protection\restrictions\users
type ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UsersDeleteRequestBody composed type wrapper for classes ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able, string
type UsersDeleteRequestBody struct {
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able
    itemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able
    // Composed type representation for type string
    string *string
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able
    usersDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able
    // Composed type representation for type string
    usersDeleteRequestBodyString *string
}
// NewUsersDeleteRequestBody instantiates a new UsersDeleteRequestBody and sets the default values.
func NewUsersDeleteRequestBody()(*UsersDeleteRequestBody) {
    m := &UsersDeleteRequestBody{
    }
    return m
}
// CreateUsersDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUsersDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUsersDeleteRequestBody()
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
        result.SetString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetUsersDeleteRequestBodyString(val)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UsersDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UsersDeleteRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able when successful
func (m *UsersDeleteRequestBody) GetItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able) {
    return m.itemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *UsersDeleteRequestBody) GetString()(*string) {
    return m.string
}
// GetUsersDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able when successful
func (m *UsersDeleteRequestBody) GetUsersDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able) {
    return m.usersDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1
}
// GetUsersDeleteRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *UsersDeleteRequestBody) GetUsersDeleteRequestBodyString()(*string) {
    return m.usersDeleteRequestBodyString
}
// Serialize serializes information the current object
func (m *UsersDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetUsersDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetUsersDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    } else if m.GetUsersDeleteRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetUsersDeleteRequestBodyString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able
func (m *UsersDeleteRequestBody) SetItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able)() {
    m.itemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *UsersDeleteRequestBody) SetString(value *string)() {
    m.string = value
}
// SetUsersDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able
func (m *UsersDeleteRequestBody) SetUsersDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able)() {
    m.usersDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1 = value
}
// SetUsersDeleteRequestBodyString sets the string property value. Composed type representation for type string
func (m *UsersDeleteRequestBody) SetUsersDeleteRequestBodyString(value *string)() {
    m.usersDeleteRequestBodyString = value
}
// UsersPostRequestBody composed type wrapper for classes ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able, string
type UsersPostRequestBody struct {
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able
    itemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able
    // Composed type representation for type string
    string *string
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able
    usersPostRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able
    // Composed type representation for type string
    usersPostRequestBodyString *string
}
// NewUsersPostRequestBody instantiates a new UsersPostRequestBody and sets the default values.
func NewUsersPostRequestBody()(*UsersPostRequestBody) {
    m := &UsersPostRequestBody{
    }
    return m
}
// CreateUsersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUsersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUsersPostRequestBody()
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
        result.SetString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetUsersPostRequestBodyString(val)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UsersPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UsersPostRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able when successful
func (m *UsersPostRequestBody) GetItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able) {
    return m.itemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *UsersPostRequestBody) GetString()(*string) {
    return m.string
}
// GetUsersPostRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able when successful
func (m *UsersPostRequestBody) GetUsersPostRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able) {
    return m.usersPostRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1
}
// GetUsersPostRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *UsersPostRequestBody) GetUsersPostRequestBodyString()(*string) {
    return m.usersPostRequestBodyString
}
// Serialize serializes information the current object
func (m *UsersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetUsersPostRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetUsersPostRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    } else if m.GetUsersPostRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetUsersPostRequestBodyString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able
func (m *UsersPostRequestBody) SetItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able)() {
    m.itemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *UsersPostRequestBody) SetString(value *string)() {
    m.string = value
}
// SetUsersPostRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able
func (m *UsersPostRequestBody) SetUsersPostRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able)() {
    m.usersPostRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1 = value
}
// SetUsersPostRequestBodyString sets the string property value. Composed type representation for type string
func (m *UsersPostRequestBody) SetUsersPostRequestBodyString(value *string)() {
    m.usersPostRequestBodyString = value
}
// UsersPutRequestBody composed type wrapper for classes ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able, string
type UsersPutRequestBody struct {
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able
    itemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able
    // Composed type representation for type string
    string *string
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able
    usersPutRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able
    // Composed type representation for type string
    usersPutRequestBodyString *string
}
// NewUsersPutRequestBody instantiates a new UsersPutRequestBody and sets the default values.
func NewUsersPutRequestBody()(*UsersPutRequestBody) {
    m := &UsersPutRequestBody{
    }
    return m
}
// CreateUsersPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUsersPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUsersPutRequestBody()
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
        result.SetString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetUsersPutRequestBodyString(val)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UsersPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UsersPutRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able when successful
func (m *UsersPutRequestBody) GetItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able) {
    return m.itemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *UsersPutRequestBody) GetString()(*string) {
    return m.string
}
// GetUsersPutRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able when successful
func (m *UsersPutRequestBody) GetUsersPutRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able) {
    return m.usersPutRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1
}
// GetUsersPutRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *UsersPutRequestBody) GetUsersPutRequestBodyString()(*string) {
    return m.usersPutRequestBodyString
}
// Serialize serializes information the current object
func (m *UsersPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetUsersPutRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetUsersPutRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    } else if m.GetUsersPutRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetUsersPutRequestBodyString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able
func (m *UsersPutRequestBody) SetItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able)() {
    m.itemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *UsersPutRequestBody) SetString(value *string)() {
    m.string = value
}
// SetUsersPutRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able
func (m *UsersPutRequestBody) SetUsersPutRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able)() {
    m.usersPutRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1 = value
}
// SetUsersPutRequestBodyString sets the string property value. Composed type representation for type string
func (m *UsersPutRequestBody) SetUsersPutRequestBodyString(value *string)() {
    m.usersPutRequestBodyString = value
}
type UsersDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able)
    GetString()(*string)
    GetUsersDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able)
    GetUsersDeleteRequestBodyString()(*string)
    SetItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able)()
    SetString(value *string)()
    SetUsersDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersDeleteRequestBodyMember1able)()
    SetUsersDeleteRequestBodyString(value *string)()
}
type UsersPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able)
    GetString()(*string)
    GetUsersPostRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able)
    GetUsersPostRequestBodyString()(*string)
    SetItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able)()
    SetString(value *string)()
    SetUsersPostRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersPostRequestBodyMember1able)()
    SetUsersPostRequestBodyString(value *string)()
}
type UsersPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able)
    GetString()(*string)
    GetUsersPutRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able)
    GetUsersPutRequestBodyString()(*string)
    SetItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able)()
    SetString(value *string)()
    SetUsersPutRequestBodyItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsUsersPutRequestBodyMember1able)()
    SetUsersPutRequestBodyString(value *string)()
}
// NewItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderInternal instantiates a new ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) {
    m := &ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/branches/{branch}/protection/restrictions/users", pathParameters),
    }
    return m
}
// NewItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder instantiates a new ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBranchesItemProtectionRestrictionsUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Removes the ability of a user to push to this branch.| Type    | Description                                                                                                                                   || ------- | --------------------------------------------------------------------------------------------------------------------------------------------- || `array` | Usernames of the people who should no longer have push access. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
// returns a []SimpleUserable when successful
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#remove-user-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) Delete(ctx context.Context, body UsersDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SimpleUserable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateSimpleUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SimpleUserable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SimpleUserable)
        }
    }
    return val, nil
}
// Get protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists the people who have push access to this branch.
// returns a []SimpleUserable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#get-users-with-access-to-the-protected-branch
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SimpleUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateSimpleUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SimpleUserable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SimpleUserable)
        }
    }
    return val, nil
}
// Post protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Grants the specified people push access for this branch.| Type    | Description                                                                                                                   || ------- | ----------------------------------------------------------------------------------------------------------------------------- || `array` | Usernames for people who can have push access. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
// returns a []SimpleUserable when successful
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#add-user-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) Post(ctx context.Context, body UsersPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SimpleUserable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateSimpleUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SimpleUserable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SimpleUserable)
        }
    }
    return val, nil
}
// Put protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Replaces the list of people that have push access to this branch. This removes all people that previously had push access and grants push access to the new list of people.| Type    | Description                                                                                                                   || ------- | ----------------------------------------------------------------------------------------------------------------------------- || `array` | Usernames for people who can have push access. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
// returns a []SimpleUserable when successful
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#set-user-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) Put(ctx context.Context, body UsersPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SimpleUserable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateSimpleUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SimpleUserable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.SimpleUserable)
        }
    }
    return val, nil
}
// ToDeleteRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Removes the ability of a user to push to this branch.| Type    | Description                                                                                                                                   || ------- | --------------------------------------------------------------------------------------------------------------------------------------------- || `array` | Usernames of the people who should no longer have push access. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body UsersDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToGetRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists the people who have push access to this branch.
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Grants the specified people push access for this branch.| Type    | Description                                                                                                                   || ------- | ----------------------------------------------------------------------------------------------------------------------------- || `array` | Usernames for people who can have push access. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) ToPostRequestInformation(ctx context.Context, body UsersPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToPutRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Replaces the list of people that have push access to this branch. This removes all people that previously had push access and grants push access to the new list of people.| Type    | Description                                                                                                                   || ------- | ----------------------------------------------------------------------------------------------------------------------------- || `array` | Usernames for people who can have push access. **Note**: The list of users, apps, and teams in total is limited to 100 items. |
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) ToPutRequestInformation(ctx context.Context, body UsersPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder when successful
func (m *ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) WithUrl(rawUrl string)(*ItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder) {
    return NewItemItemBranchesItemProtectionRestrictionsUsersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
