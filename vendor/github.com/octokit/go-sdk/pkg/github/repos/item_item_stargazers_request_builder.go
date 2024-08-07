package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemStargazersRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\stargazers
type ItemItemStargazersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemStargazersRequestBuilderGetQueryParameters lists the people that have starred the repository.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.star+json`**: Includes a timestamp of when the star was created.
type ItemItemStargazersRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
// StargazersGetResponse composed type wrapper for classes []ItemItemStargazersSimpleUserable, []ItemItemStargazersStargazerable
type StargazersGetResponse struct {
    // Composed type representation for type []ItemItemStargazersSimpleUserable
    itemItemStargazersSimpleUser []ItemItemStargazersSimpleUserable
    // Composed type representation for type []ItemItemStargazersStargazerable
    itemItemStargazersStargazer []ItemItemStargazersStargazerable
}
// NewStargazersGetResponse instantiates a new StargazersGetResponse and sets the default values.
func NewStargazersGetResponse()(*StargazersGetResponse) {
    m := &StargazersGetResponse{
    }
    return m
}
// CreateStargazersGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStargazersGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewStargazersGetResponse()
    if parseNode != nil {
        if val, err := parseNode.GetCollectionOfObjectValues(CreateItemItemStargazersSimpleUserFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            cast := make([]ItemItemStargazersSimpleUserable, len(val))
            for i, v := range val {
                if v != nil {
                    cast[i] = v.(ItemItemStargazersSimpleUserable)
                }
            }
            result.SetItemItemStargazersSimpleUser(cast)
        } else if val, err := parseNode.GetCollectionOfObjectValues(CreateItemItemStargazersStargazerFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            cast := make([]ItemItemStargazersStargazerable, len(val))
            for i, v := range val {
                if v != nil {
                    cast[i] = v.(ItemItemStargazersStargazerable)
                }
            }
            result.SetItemItemStargazersStargazer(cast)
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StargazersGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *StargazersGetResponse) GetIsComposedType()(bool) {
    return true
}
// GetItemItemStargazersSimpleUser gets the ItemItemStargazersSimpleUser property value. Composed type representation for type []ItemItemStargazersSimpleUserable
// returns a []ItemItemStargazersSimpleUserable when successful
func (m *StargazersGetResponse) GetItemItemStargazersSimpleUser()([]ItemItemStargazersSimpleUserable) {
    return m.itemItemStargazersSimpleUser
}
// GetItemItemStargazersStargazer gets the ItemItemStargazersStargazer property value. Composed type representation for type []ItemItemStargazersStargazerable
// returns a []ItemItemStargazersStargazerable when successful
func (m *StargazersGetResponse) GetItemItemStargazersStargazer()([]ItemItemStargazersStargazerable) {
    return m.itemItemStargazersStargazer
}
// Serialize serializes information the current object
func (m *StargazersGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItemItemStargazersSimpleUser() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItemItemStargazersSimpleUser()))
        for i, v := range m.GetItemItemStargazersSimpleUser() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("", cast)
        if err != nil {
            return err
        }
    } else if m.GetItemItemStargazersStargazer() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItemItemStargazersStargazer()))
        for i, v := range m.GetItemItemStargazersStargazer() {
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
// SetItemItemStargazersSimpleUser sets the ItemItemStargazersSimpleUser property value. Composed type representation for type []ItemItemStargazersSimpleUserable
func (m *StargazersGetResponse) SetItemItemStargazersSimpleUser(value []ItemItemStargazersSimpleUserable)() {
    m.itemItemStargazersSimpleUser = value
}
// SetItemItemStargazersStargazer sets the ItemItemStargazersStargazer property value. Composed type representation for type []ItemItemStargazersStargazerable
func (m *StargazersGetResponse) SetItemItemStargazersStargazer(value []ItemItemStargazersStargazerable)() {
    m.itemItemStargazersStargazer = value
}
type StargazersGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemItemStargazersSimpleUser()([]ItemItemStargazersSimpleUserable)
    GetItemItemStargazersStargazer()([]ItemItemStargazersStargazerable)
    SetItemItemStargazersSimpleUser(value []ItemItemStargazersSimpleUserable)()
    SetItemItemStargazersStargazer(value []ItemItemStargazersStargazerable)()
}
// NewItemItemStargazersRequestBuilderInternal instantiates a new ItemItemStargazersRequestBuilder and sets the default values.
func NewItemItemStargazersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStargazersRequestBuilder) {
    m := &ItemItemStargazersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/stargazers{?page*,per_page*}", pathParameters),
    }
    return m
}
// NewItemItemStargazersRequestBuilder instantiates a new ItemItemStargazersRequestBuilder and sets the default values.
func NewItemItemStargazersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemStargazersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemStargazersRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the people that have starred the repository.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.star+json`**: Includes a timestamp of when the star was created.
// returns a StargazersGetResponseable when successful
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/activity/starring#list-stargazers
func (m *ItemItemStargazersRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemStargazersRequestBuilderGetQueryParameters])(StargazersGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateStargazersGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(StargazersGetResponseable), nil
}
// ToGetRequestInformation lists the people that have starred the repository.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.star+json`**: Includes a timestamp of when the star was created.
// returns a *RequestInformation when successful
func (m *ItemItemStargazersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemStargazersRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemStargazersRequestBuilder when successful
func (m *ItemItemStargazersRequestBuilder) WithUrl(rawUrl string)(*ItemItemStargazersRequestBuilder) {
    return NewItemItemStargazersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
