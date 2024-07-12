package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\branches\{branch}\protection\restrictions\teams
type ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamsDeleteRequestBody composed type wrapper for classes ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able, string
type TeamsDeleteRequestBody struct {
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able
    itemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able
    // Composed type representation for type string
    string *string
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able
    teamsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able
    // Composed type representation for type string
    teamsDeleteRequestBodyString *string
}
// NewTeamsDeleteRequestBody instantiates a new TeamsDeleteRequestBody and sets the default values.
func NewTeamsDeleteRequestBody()(*TeamsDeleteRequestBody) {
    m := &TeamsDeleteRequestBody{
    }
    return m
}
// CreateTeamsDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamsDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewTeamsDeleteRequestBody()
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
        result.SetTeamsDeleteRequestBodyString(val)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamsDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *TeamsDeleteRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able when successful
func (m *TeamsDeleteRequestBody) GetItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able) {
    return m.itemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *TeamsDeleteRequestBody) GetString()(*string) {
    return m.string
}
// GetTeamsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able when successful
func (m *TeamsDeleteRequestBody) GetTeamsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able) {
    return m.teamsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1
}
// GetTeamsDeleteRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *TeamsDeleteRequestBody) GetTeamsDeleteRequestBodyString()(*string) {
    return m.teamsDeleteRequestBodyString
}
// Serialize serializes information the current object
func (m *TeamsDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetTeamsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetTeamsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    } else if m.GetTeamsDeleteRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetTeamsDeleteRequestBodyString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able
func (m *TeamsDeleteRequestBody) SetItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able)() {
    m.itemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *TeamsDeleteRequestBody) SetString(value *string)() {
    m.string = value
}
// SetTeamsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able
func (m *TeamsDeleteRequestBody) SetTeamsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able)() {
    m.teamsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1 = value
}
// SetTeamsDeleteRequestBodyString sets the string property value. Composed type representation for type string
func (m *TeamsDeleteRequestBody) SetTeamsDeleteRequestBodyString(value *string)() {
    m.teamsDeleteRequestBodyString = value
}
// TeamsPostRequestBody composed type wrapper for classes ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able, string
type TeamsPostRequestBody struct {
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able
    itemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able
    // Composed type representation for type string
    string *string
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able
    teamsPostRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able
    // Composed type representation for type string
    teamsPostRequestBodyString *string
}
// NewTeamsPostRequestBody instantiates a new TeamsPostRequestBody and sets the default values.
func NewTeamsPostRequestBody()(*TeamsPostRequestBody) {
    m := &TeamsPostRequestBody{
    }
    return m
}
// CreateTeamsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewTeamsPostRequestBody()
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
        result.SetTeamsPostRequestBodyString(val)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *TeamsPostRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able when successful
func (m *TeamsPostRequestBody) GetItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able) {
    return m.itemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *TeamsPostRequestBody) GetString()(*string) {
    return m.string
}
// GetTeamsPostRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able when successful
func (m *TeamsPostRequestBody) GetTeamsPostRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able) {
    return m.teamsPostRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1
}
// GetTeamsPostRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *TeamsPostRequestBody) GetTeamsPostRequestBodyString()(*string) {
    return m.teamsPostRequestBodyString
}
// Serialize serializes information the current object
func (m *TeamsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetTeamsPostRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetTeamsPostRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    } else if m.GetTeamsPostRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetTeamsPostRequestBodyString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able
func (m *TeamsPostRequestBody) SetItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able)() {
    m.itemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *TeamsPostRequestBody) SetString(value *string)() {
    m.string = value
}
// SetTeamsPostRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able
func (m *TeamsPostRequestBody) SetTeamsPostRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able)() {
    m.teamsPostRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1 = value
}
// SetTeamsPostRequestBodyString sets the string property value. Composed type representation for type string
func (m *TeamsPostRequestBody) SetTeamsPostRequestBodyString(value *string)() {
    m.teamsPostRequestBodyString = value
}
// TeamsPutRequestBody composed type wrapper for classes ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able, string
type TeamsPutRequestBody struct {
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able
    itemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able
    // Composed type representation for type string
    string *string
    // Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able
    teamsPutRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1 ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able
    // Composed type representation for type string
    teamsPutRequestBodyString *string
}
// NewTeamsPutRequestBody instantiates a new TeamsPutRequestBody and sets the default values.
func NewTeamsPutRequestBody()(*TeamsPutRequestBody) {
    m := &TeamsPutRequestBody{
    }
    return m
}
// CreateTeamsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewTeamsPutRequestBody()
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
        result.SetTeamsPutRequestBodyString(val)
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamsPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *TeamsPutRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able when successful
func (m *TeamsPutRequestBody) GetItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able) {
    return m.itemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *TeamsPutRequestBody) GetString()(*string) {
    return m.string
}
// GetTeamsPutRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1 gets the ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able
// returns a ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able when successful
func (m *TeamsPutRequestBody) GetTeamsPutRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able) {
    return m.teamsPutRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1
}
// GetTeamsPutRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *TeamsPutRequestBody) GetTeamsPutRequestBodyString()(*string) {
    return m.teamsPutRequestBodyString
}
// Serialize serializes information the current object
func (m *TeamsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetTeamsPutRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetTeamsPutRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    } else if m.GetTeamsPutRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetTeamsPutRequestBodyString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able
func (m *TeamsPutRequestBody) SetItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able)() {
    m.itemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1 = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *TeamsPutRequestBody) SetString(value *string)() {
    m.string = value
}
// SetTeamsPutRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1 sets the ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1 property value. Composed type representation for type ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able
func (m *TeamsPutRequestBody) SetTeamsPutRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able)() {
    m.teamsPutRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1 = value
}
// SetTeamsPutRequestBodyString sets the string property value. Composed type representation for type string
func (m *TeamsPutRequestBody) SetTeamsPutRequestBodyString(value *string)() {
    m.teamsPutRequestBodyString = value
}
type TeamsDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able)
    GetString()(*string)
    GetTeamsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able)
    GetTeamsDeleteRequestBodyString()(*string)
    SetItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able)()
    SetString(value *string)()
    SetTeamsDeleteRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsDeleteRequestBodyMember1able)()
    SetTeamsDeleteRequestBodyString(value *string)()
}
type TeamsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able)
    GetString()(*string)
    GetTeamsPostRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able)
    GetTeamsPostRequestBodyString()(*string)
    SetItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able)()
    SetString(value *string)()
    SetTeamsPostRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsPostRequestBodyMember1able)()
    SetTeamsPostRequestBodyString(value *string)()
}
type TeamsPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able)
    GetString()(*string)
    GetTeamsPutRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1()(ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able)
    GetTeamsPutRequestBodyString()(*string)
    SetItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able)()
    SetString(value *string)()
    SetTeamsPutRequestBodyItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1(value ItemItemBranchesItemProtectionRestrictionsTeamsPutRequestBodyMember1able)()
    SetTeamsPutRequestBodyString(value *string)()
}
// NewItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderInternal instantiates a new ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) {
    m := &ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/branches/{branch}/protection/restrictions/teams", pathParameters),
    }
    return m
}
// NewItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder instantiates a new ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder and sets the default values.
func NewItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Removes the ability of a team to push to this branch. You can also remove push access for child teams.
// returns a []Teamable when successful
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#remove-team-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) Delete(ctx context.Context, body TeamsDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Teamable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Teamable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Teamable)
        }
    }
    return val, nil
}
// Get protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists the teams who have push access to this branch. The list includes child teams.
// returns a []Teamable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#get-teams-with-access-to-the-protected-branch
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Teamable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Teamable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Teamable)
        }
    }
    return val, nil
}
// Post protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Grants the specified teams push access for this branch. You can also give push access to child teams.
// returns a []Teamable when successful
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#add-team-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) Post(ctx context.Context, body TeamsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Teamable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Teamable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Teamable)
        }
    }
    return val, nil
}
// Put protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Replaces the list of teams that have push access to this branch. This removes all teams that previously had push access and grants push access to the new list of teams. Team restrictions include child teams.
// returns a []Teamable when successful
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/branches/branch-protection#set-team-access-restrictions
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) Put(ctx context.Context, body TeamsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Teamable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Teamable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Teamable)
        }
    }
    return val, nil
}
// ToDeleteRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Removes the ability of a team to push to this branch. You can also remove push access for child teams.
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body TeamsDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToGetRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Lists the teams who have push access to this branch. The list includes child teams.
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Grants the specified teams push access for this branch. You can also give push access to child teams.
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) ToPostRequestInformation(ctx context.Context, body TeamsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToPutRequestInformation protected branches are available in public repositories with GitHub Free and GitHub Free for organizations, and in public and private repositories with GitHub Pro, GitHub Team, GitHub Enterprise Cloud, and GitHub Enterprise Server. For more information, see [GitHub's products](https://docs.github.com/github/getting-started-with-github/githubs-products) in the GitHub Help documentation.Replaces the list of teams that have push access to this branch. This removes all teams that previously had push access and grants push access to the new list of teams. Team restrictions include child teams.
// returns a *RequestInformation when successful
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) ToPutRequestInformation(ctx context.Context, body TeamsPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder when successful
func (m *ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) WithUrl(rawUrl string)(*ItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder) {
    return NewItemItemBranchesItemProtectionRestrictionsTeamsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
