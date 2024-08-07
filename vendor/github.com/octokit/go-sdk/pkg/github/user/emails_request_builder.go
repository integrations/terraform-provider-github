package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// EmailsRequestBuilder builds and executes requests for operations under \user\emails
type EmailsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EmailsDeleteRequestBody composed type wrapper for classes EmailsDeleteRequestBodyMember1able, string
type EmailsDeleteRequestBody struct {
    // Composed type representation for type EmailsDeleteRequestBodyMember1able
    emailsDeleteRequestBodyEmailsDeleteRequestBodyMember1 EmailsDeleteRequestBodyMember1able
    // Composed type representation for type EmailsDeleteRequestBodyMember1able
    emailsDeleteRequestBodyMember1 EmailsDeleteRequestBodyMember1able
    // Composed type representation for type string
    emailsDeleteRequestBodyString *string
    // Composed type representation for type string
    string *string
}
// NewEmailsDeleteRequestBody instantiates a new EmailsDeleteRequestBody and sets the default values.
func NewEmailsDeleteRequestBody()(*EmailsDeleteRequestBody) {
    m := &EmailsDeleteRequestBody{
    }
    return m
}
// CreateEmailsDeleteRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsDeleteRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewEmailsDeleteRequestBody()
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
        result.SetEmailsDeleteRequestBodyString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetEmailsDeleteRequestBodyEmailsDeleteRequestBodyMember1 gets the EmailsDeleteRequestBodyMember1 property value. Composed type representation for type EmailsDeleteRequestBodyMember1able
// returns a EmailsDeleteRequestBodyMember1able when successful
func (m *EmailsDeleteRequestBody) GetEmailsDeleteRequestBodyEmailsDeleteRequestBodyMember1()(EmailsDeleteRequestBodyMember1able) {
    return m.emailsDeleteRequestBodyEmailsDeleteRequestBodyMember1
}
// GetEmailsDeleteRequestBodyMember1 gets the EmailsDeleteRequestBodyMember1 property value. Composed type representation for type EmailsDeleteRequestBodyMember1able
// returns a EmailsDeleteRequestBodyMember1able when successful
func (m *EmailsDeleteRequestBody) GetEmailsDeleteRequestBodyMember1()(EmailsDeleteRequestBodyMember1able) {
    return m.emailsDeleteRequestBodyMember1
}
// GetEmailsDeleteRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *EmailsDeleteRequestBody) GetEmailsDeleteRequestBodyString()(*string) {
    return m.emailsDeleteRequestBodyString
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailsDeleteRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *EmailsDeleteRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *EmailsDeleteRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *EmailsDeleteRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEmailsDeleteRequestBodyEmailsDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetEmailsDeleteRequestBodyEmailsDeleteRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetEmailsDeleteRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetEmailsDeleteRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetEmailsDeleteRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetEmailsDeleteRequestBodyString())
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
// SetEmailsDeleteRequestBodyEmailsDeleteRequestBodyMember1 sets the EmailsDeleteRequestBodyMember1 property value. Composed type representation for type EmailsDeleteRequestBodyMember1able
func (m *EmailsDeleteRequestBody) SetEmailsDeleteRequestBodyEmailsDeleteRequestBodyMember1(value EmailsDeleteRequestBodyMember1able)() {
    m.emailsDeleteRequestBodyEmailsDeleteRequestBodyMember1 = value
}
// SetEmailsDeleteRequestBodyMember1 sets the EmailsDeleteRequestBodyMember1 property value. Composed type representation for type EmailsDeleteRequestBodyMember1able
func (m *EmailsDeleteRequestBody) SetEmailsDeleteRequestBodyMember1(value EmailsDeleteRequestBodyMember1able)() {
    m.emailsDeleteRequestBodyMember1 = value
}
// SetEmailsDeleteRequestBodyString sets the string property value. Composed type representation for type string
func (m *EmailsDeleteRequestBody) SetEmailsDeleteRequestBodyString(value *string)() {
    m.emailsDeleteRequestBodyString = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *EmailsDeleteRequestBody) SetString(value *string)() {
    m.string = value
}
// EmailsPostRequestBody composed type wrapper for classes EmailsPostRequestBodyMember1able, string
type EmailsPostRequestBody struct {
    // Composed type representation for type EmailsPostRequestBodyMember1able
    emailsPostRequestBodyEmailsPostRequestBodyMember1 EmailsPostRequestBodyMember1able
    // Composed type representation for type EmailsPostRequestBodyMember1able
    emailsPostRequestBodyMember1 EmailsPostRequestBodyMember1able
    // Composed type representation for type string
    emailsPostRequestBodyString *string
    // Composed type representation for type string
    string *string
}
// NewEmailsPostRequestBody instantiates a new EmailsPostRequestBody and sets the default values.
func NewEmailsPostRequestBody()(*EmailsPostRequestBody) {
    m := &EmailsPostRequestBody{
    }
    return m
}
// CreateEmailsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewEmailsPostRequestBody()
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
        result.SetEmailsPostRequestBodyString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetEmailsPostRequestBodyEmailsPostRequestBodyMember1 gets the EmailsPostRequestBodyMember1 property value. Composed type representation for type EmailsPostRequestBodyMember1able
// returns a EmailsPostRequestBodyMember1able when successful
func (m *EmailsPostRequestBody) GetEmailsPostRequestBodyEmailsPostRequestBodyMember1()(EmailsPostRequestBodyMember1able) {
    return m.emailsPostRequestBodyEmailsPostRequestBodyMember1
}
// GetEmailsPostRequestBodyMember1 gets the EmailsPostRequestBodyMember1 property value. Composed type representation for type EmailsPostRequestBodyMember1able
// returns a EmailsPostRequestBodyMember1able when successful
func (m *EmailsPostRequestBody) GetEmailsPostRequestBodyMember1()(EmailsPostRequestBodyMember1able) {
    return m.emailsPostRequestBodyMember1
}
// GetEmailsPostRequestBodyString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *EmailsPostRequestBody) GetEmailsPostRequestBodyString()(*string) {
    return m.emailsPostRequestBodyString
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *EmailsPostRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *EmailsPostRequestBody) GetString()(*string) {
    return m.string
}
// Serialize serializes information the current object
func (m *EmailsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEmailsPostRequestBodyEmailsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetEmailsPostRequestBodyEmailsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetEmailsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetEmailsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetEmailsPostRequestBodyString() != nil {
        err := writer.WriteStringValue("", m.GetEmailsPostRequestBodyString())
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
// SetEmailsPostRequestBodyEmailsPostRequestBodyMember1 sets the EmailsPostRequestBodyMember1 property value. Composed type representation for type EmailsPostRequestBodyMember1able
func (m *EmailsPostRequestBody) SetEmailsPostRequestBodyEmailsPostRequestBodyMember1(value EmailsPostRequestBodyMember1able)() {
    m.emailsPostRequestBodyEmailsPostRequestBodyMember1 = value
}
// SetEmailsPostRequestBodyMember1 sets the EmailsPostRequestBodyMember1 property value. Composed type representation for type EmailsPostRequestBodyMember1able
func (m *EmailsPostRequestBody) SetEmailsPostRequestBodyMember1(value EmailsPostRequestBodyMember1able)() {
    m.emailsPostRequestBodyMember1 = value
}
// SetEmailsPostRequestBodyString sets the string property value. Composed type representation for type string
func (m *EmailsPostRequestBody) SetEmailsPostRequestBodyString(value *string)() {
    m.emailsPostRequestBodyString = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *EmailsPostRequestBody) SetString(value *string)() {
    m.string = value
}
// EmailsRequestBuilderGetQueryParameters lists all of your email addresses, and specifies which one is visibleto the public.OAuth app tokens and personal access tokens (classic) need the `user:email` scope to use this endpoint.
type EmailsRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
type EmailsDeleteRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmailsDeleteRequestBodyEmailsDeleteRequestBodyMember1()(EmailsDeleteRequestBodyMember1able)
    GetEmailsDeleteRequestBodyMember1()(EmailsDeleteRequestBodyMember1able)
    GetEmailsDeleteRequestBodyString()(*string)
    GetString()(*string)
    SetEmailsDeleteRequestBodyEmailsDeleteRequestBodyMember1(value EmailsDeleteRequestBodyMember1able)()
    SetEmailsDeleteRequestBodyMember1(value EmailsDeleteRequestBodyMember1able)()
    SetEmailsDeleteRequestBodyString(value *string)()
    SetString(value *string)()
}
type EmailsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmailsPostRequestBodyEmailsPostRequestBodyMember1()(EmailsPostRequestBodyMember1able)
    GetEmailsPostRequestBodyMember1()(EmailsPostRequestBodyMember1able)
    GetEmailsPostRequestBodyString()(*string)
    GetString()(*string)
    SetEmailsPostRequestBodyEmailsPostRequestBodyMember1(value EmailsPostRequestBodyMember1able)()
    SetEmailsPostRequestBodyMember1(value EmailsPostRequestBodyMember1able)()
    SetEmailsPostRequestBodyString(value *string)()
    SetString(value *string)()
}
// NewEmailsRequestBuilderInternal instantiates a new EmailsRequestBuilder and sets the default values.
func NewEmailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmailsRequestBuilder) {
    m := &EmailsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/emails{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewEmailsRequestBuilder instantiates a new EmailsRequestBuilder and sets the default values.
func NewEmailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmailsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete oAuth app tokens and personal access tokens (classic) need the `user` scope to use this endpoint.
// returns a BasicError error when the service returns a 401 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/emails#delete-an-email-address-for-the-authenticated-user
func (m *EmailsRequestBuilder) Delete(ctx context.Context, body EmailsDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get lists all of your email addresses, and specifies which one is visibleto the public.OAuth app tokens and personal access tokens (classic) need the `user:email` scope to use this endpoint.
// returns a []Emailable when successful
// returns a BasicError error when the service returns a 401 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/emails#list-email-addresses-for-the-authenticated-user
func (m *EmailsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[EmailsRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Emailable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateEmailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Emailable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Emailable)
        }
    }
    return val, nil
}
// Post oAuth app tokens and personal access tokens (classic) need the `user` scope to use this endpoint.
// returns a []Emailable when successful
// returns a BasicError error when the service returns a 401 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/users/emails#add-an-email-address-for-the-authenticated-user
func (m *EmailsRequestBuilder) Post(ctx context.Context, body EmailsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Emailable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateEmailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Emailable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Emailable)
        }
    }
    return val, nil
}
// ToDeleteRequestInformation oAuth app tokens and personal access tokens (classic) need the `user` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *EmailsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body EmailsDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToGetRequestInformation lists all of your email addresses, and specifies which one is visibleto the public.OAuth app tokens and personal access tokens (classic) need the `user:email` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *EmailsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[EmailsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation oAuth app tokens and personal access tokens (classic) need the `user` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *EmailsRequestBuilder) ToPostRequestInformation(ctx context.Context, body EmailsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EmailsRequestBuilder when successful
func (m *EmailsRequestBuilder) WithUrl(rawUrl string)(*EmailsRequestBuilder) {
    return NewEmailsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
