package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemTokenScopedPostRequestBody struct {
    // The access token used to authenticate to the GitHub API.
    access_token *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The permissions granted to the user access token.
    permissions i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AppPermissionsable
    // The list of repository names to scope the user access token to. `repositories` may not be specified if `repository_ids` is specified.
    repositories []string
    // The list of repository IDs to scope the user access token to. `repository_ids` may not be specified if `repositories` is specified.
    repository_ids []int32
    // The name of the user or organization to scope the user access token to. **Required** unless `target_id` is specified.
    target *string
    // The ID of the user or organization to scope the user access token to. **Required** unless `target` is specified.
    target_id *int32
}
// NewItemTokenScopedPostRequestBody instantiates a new ItemTokenScopedPostRequestBody and sets the default values.
func NewItemTokenScopedPostRequestBody()(*ItemTokenScopedPostRequestBody) {
    m := &ItemTokenScopedPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemTokenScopedPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTokenScopedPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTokenScopedPostRequestBody(), nil
}
// GetAccessToken gets the access_token property value. The access token used to authenticate to the GitHub API.
// returns a *string when successful
func (m *ItemTokenScopedPostRequestBody) GetAccessToken()(*string) {
    return m.access_token
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemTokenScopedPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemTokenScopedPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access_token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessToken(val)
        }
        return nil
    }
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateAppPermissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AppPermissionsable))
        }
        return nil
    }
    res["repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRepositories(res)
        }
        return nil
    }
    res["repository_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetRepositoryIds(res)
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    res["target_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetId(val)
        }
        return nil
    }
    return res
}
// GetPermissions gets the permissions property value. The permissions granted to the user access token.
// returns a AppPermissionsable when successful
func (m *ItemTokenScopedPostRequestBody) GetPermissions()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AppPermissionsable) {
    return m.permissions
}
// GetRepositories gets the repositories property value. The list of repository names to scope the user access token to. `repositories` may not be specified if `repository_ids` is specified.
// returns a []string when successful
func (m *ItemTokenScopedPostRequestBody) GetRepositories()([]string) {
    return m.repositories
}
// GetRepositoryIds gets the repository_ids property value. The list of repository IDs to scope the user access token to. `repository_ids` may not be specified if `repositories` is specified.
// returns a []int32 when successful
func (m *ItemTokenScopedPostRequestBody) GetRepositoryIds()([]int32) {
    return m.repository_ids
}
// GetTarget gets the target property value. The name of the user or organization to scope the user access token to. **Required** unless `target_id` is specified.
// returns a *string when successful
func (m *ItemTokenScopedPostRequestBody) GetTarget()(*string) {
    return m.target
}
// GetTargetId gets the target_id property value. The ID of the user or organization to scope the user access token to. **Required** unless `target` is specified.
// returns a *int32 when successful
func (m *ItemTokenScopedPostRequestBody) GetTargetId()(*int32) {
    return m.target_id
}
// Serialize serializes information the current object
func (m *ItemTokenScopedPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("access_token", m.GetAccessToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("permissions", m.GetPermissions())
        if err != nil {
            return err
        }
    }
    if m.GetRepositories() != nil {
        err := writer.WriteCollectionOfStringValues("repositories", m.GetRepositories())
        if err != nil {
            return err
        }
    }
    if m.GetRepositoryIds() != nil {
        err := writer.WriteCollectionOfInt32Values("repository_ids", m.GetRepositoryIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("target_id", m.GetTargetId())
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
// SetAccessToken sets the access_token property value. The access token used to authenticate to the GitHub API.
func (m *ItemTokenScopedPostRequestBody) SetAccessToken(value *string)() {
    m.access_token = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemTokenScopedPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPermissions sets the permissions property value. The permissions granted to the user access token.
func (m *ItemTokenScopedPostRequestBody) SetPermissions(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AppPermissionsable)() {
    m.permissions = value
}
// SetRepositories sets the repositories property value. The list of repository names to scope the user access token to. `repositories` may not be specified if `repository_ids` is specified.
func (m *ItemTokenScopedPostRequestBody) SetRepositories(value []string)() {
    m.repositories = value
}
// SetRepositoryIds sets the repository_ids property value. The list of repository IDs to scope the user access token to. `repository_ids` may not be specified if `repositories` is specified.
func (m *ItemTokenScopedPostRequestBody) SetRepositoryIds(value []int32)() {
    m.repository_ids = value
}
// SetTarget sets the target property value. The name of the user or organization to scope the user access token to. **Required** unless `target_id` is specified.
func (m *ItemTokenScopedPostRequestBody) SetTarget(value *string)() {
    m.target = value
}
// SetTargetId sets the target_id property value. The ID of the user or organization to scope the user access token to. **Required** unless `target` is specified.
func (m *ItemTokenScopedPostRequestBody) SetTargetId(value *int32)() {
    m.target_id = value
}
type ItemTokenScopedPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessToken()(*string)
    GetPermissions()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AppPermissionsable)
    GetRepositories()([]string)
    GetRepositoryIds()([]int32)
    GetTarget()(*string)
    GetTargetId()(*int32)
    SetAccessToken(value *string)()
    SetPermissions(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.AppPermissionsable)()
    SetRepositories(value []string)()
    SetRepositoryIds(value []int32)()
    SetTarget(value *string)()
    SetTargetId(value *int32)()
}
