package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// WithAccount_ItemRequestBuilder builds and executes requests for operations under \user\{account_id}
type WithAccount_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithAccount_GetResponse composed type wrapper for classes i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PrivateUserable, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PublicUserable
type WithAccount_GetResponse struct {
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PrivateUserable
    privateUser i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PrivateUserable
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PublicUserable
    publicUser i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PublicUserable
}
// NewWithAccount_GetResponse instantiates a new WithAccount_GetResponse and sets the default values.
func NewWithAccount_GetResponse()(*WithAccount_GetResponse) {
    m := &WithAccount_GetResponse{
    }
    return m
}
// CreateWithAccount_GetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWithAccount_GetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWithAccount_GetResponse()
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
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WithAccount_GetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *WithAccount_GetResponse) GetIsComposedType()(bool) {
    return true
}
// GetPrivateUser gets the privateUser property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PrivateUserable
// returns a PrivateUserable when successful
func (m *WithAccount_GetResponse) GetPrivateUser()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PrivateUserable) {
    return m.privateUser
}
// GetPublicUser gets the publicUser property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PublicUserable
// returns a PublicUserable when successful
func (m *WithAccount_GetResponse) GetPublicUser()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PublicUserable) {
    return m.publicUser
}
// Serialize serializes information the current object
func (m *WithAccount_GetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPrivateUser() != nil {
        err := writer.WriteObjectValue("", m.GetPrivateUser())
        if err != nil {
            return err
        }
    } else if m.GetPublicUser() != nil {
        err := writer.WriteObjectValue("", m.GetPublicUser())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPrivateUser sets the privateUser property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PrivateUserable
func (m *WithAccount_GetResponse) SetPrivateUser(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PrivateUserable)() {
    m.privateUser = value
}
// SetPublicUser sets the publicUser property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PublicUserable
func (m *WithAccount_GetResponse) SetPublicUser(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PublicUserable)() {
    m.publicUser = value
}
type WithAccount_GetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrivateUser()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PrivateUserable)
    GetPublicUser()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PublicUserable)
    SetPrivateUser(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PrivateUserable)()
    SetPublicUser(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.PublicUserable)()
}
// NewWithAccount_ItemRequestBuilderInternal instantiates a new WithAccount_ItemRequestBuilder and sets the default values.
func NewWithAccount_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithAccount_ItemRequestBuilder) {
    m := &WithAccount_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/{account_id}", pathParameters),
    }
    return m
}
// NewWithAccount_ItemRequestBuilder instantiates a new WithAccount_ItemRequestBuilder and sets the default values.
func NewWithAccount_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithAccount_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithAccount_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get provides publicly available information about someone with a GitHub account. This method takes their durable user `ID` instead of their `login`, which can change over time.The `email` key in the following response is the publicly visible email address from your GitHub [profile page](https://github.com/settings/profile). When setting up your profile, you can select a primary email address to be “public” which provides an email entry for this endpoint. If you do not set a public email address for `email`, then it will have a value of `null`. You only see publicly visible email addresses when authenticated with GitHub. For more information, see [Authentication](https://docs.github.com/rest/guides/getting-started-with-the-rest-api#authentication).The Emails API enables you to list all of your email addresses, and toggle a primary email to be visible publicly. For more information, see "[Emails API](https://docs.github.com/rest/users/emails)".
// returns a WithAccount_GetResponseable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/users#get-a-user-using-their-id
func (m *WithAccount_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(WithAccount_GetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateWithAccount_GetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(WithAccount_GetResponseable), nil
}
// ToGetRequestInformation provides publicly available information about someone with a GitHub account. This method takes their durable user `ID` instead of their `login`, which can change over time.The `email` key in the following response is the publicly visible email address from your GitHub [profile page](https://github.com/settings/profile). When setting up your profile, you can select a primary email address to be “public” which provides an email entry for this endpoint. If you do not set a public email address for `email`, then it will have a value of `null`. You only see publicly visible email addresses when authenticated with GitHub. For more information, see [Authentication](https://docs.github.com/rest/guides/getting-started-with-the-rest-api#authentication).The Emails API enables you to list all of your email addresses, and toggle a primary email to be visible publicly. For more information, see "[Emails API](https://docs.github.com/rest/users/emails)".
// returns a *RequestInformation when successful
func (m *WithAccount_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WithAccount_ItemRequestBuilder when successful
func (m *WithAccount_ItemRequestBuilder) WithUrl(rawUrl string)(*WithAccount_ItemRequestBuilder) {
    return NewWithAccount_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
