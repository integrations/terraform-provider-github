package projects

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    iaecb1d9ce8f039d73681b50025b6b1a10860720b981e5b38dba5e4bcaf58ea7c "github.com/octokit/go-sdk/pkg/github/projects/columns/item/cards"
)

// ColumnsItemCardsRequestBuilder builds and executes requests for operations under \projects\columns\{column_id}\cards
type ColumnsItemCardsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CardsPostRequestBody composed type wrapper for classes ColumnsItemCardsPostRequestBodyMember1able, ColumnsItemCardsPostRequestBodyMember2able
type CardsPostRequestBody struct {
    // Composed type representation for type ColumnsItemCardsPostRequestBodyMember1able
    cardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1 ColumnsItemCardsPostRequestBodyMember1able
    // Composed type representation for type ColumnsItemCardsPostRequestBodyMember2able
    cardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2 ColumnsItemCardsPostRequestBodyMember2able
    // Composed type representation for type ColumnsItemCardsPostRequestBodyMember1able
    columnsItemCardsPostRequestBodyMember1 ColumnsItemCardsPostRequestBodyMember1able
    // Composed type representation for type ColumnsItemCardsPostRequestBodyMember2able
    columnsItemCardsPostRequestBodyMember2 ColumnsItemCardsPostRequestBodyMember2able
}
// NewCardsPostRequestBody instantiates a new CardsPostRequestBody and sets the default values.
func NewCardsPostRequestBody()(*CardsPostRequestBody) {
    m := &CardsPostRequestBody{
    }
    return m
}
// CreateCardsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCardsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewCardsPostRequestBody()
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
// GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1 gets the ColumnsItemCardsPostRequestBodyMember1 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember1able
// returns a ColumnsItemCardsPostRequestBodyMember1able when successful
func (m *CardsPostRequestBody) GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1()(ColumnsItemCardsPostRequestBodyMember1able) {
    return m.cardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1
}
// GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2 gets the ColumnsItemCardsPostRequestBodyMember2 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember2able
// returns a ColumnsItemCardsPostRequestBodyMember2able when successful
func (m *CardsPostRequestBody) GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2()(ColumnsItemCardsPostRequestBodyMember2able) {
    return m.cardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2
}
// GetColumnsItemCardsPostRequestBodyMember1 gets the ColumnsItemCardsPostRequestBodyMember1 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember1able
// returns a ColumnsItemCardsPostRequestBodyMember1able when successful
func (m *CardsPostRequestBody) GetColumnsItemCardsPostRequestBodyMember1()(ColumnsItemCardsPostRequestBodyMember1able) {
    return m.columnsItemCardsPostRequestBodyMember1
}
// GetColumnsItemCardsPostRequestBodyMember2 gets the ColumnsItemCardsPostRequestBodyMember2 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember2able
// returns a ColumnsItemCardsPostRequestBodyMember2able when successful
func (m *CardsPostRequestBody) GetColumnsItemCardsPostRequestBodyMember2()(ColumnsItemCardsPostRequestBodyMember2able) {
    return m.columnsItemCardsPostRequestBodyMember2
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CardsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *CardsPostRequestBody) GetIsComposedType()(bool) {
    return true
}
// Serialize serializes information the current object
func (m *CardsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2())
        if err != nil {
            return err
        }
    } else if m.GetColumnsItemCardsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetColumnsItemCardsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetColumnsItemCardsPostRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetColumnsItemCardsPostRequestBodyMember2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1 sets the ColumnsItemCardsPostRequestBodyMember1 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember1able
func (m *CardsPostRequestBody) SetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1(value ColumnsItemCardsPostRequestBodyMember1able)() {
    m.cardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1 = value
}
// SetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2 sets the ColumnsItemCardsPostRequestBodyMember2 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember2able
func (m *CardsPostRequestBody) SetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2(value ColumnsItemCardsPostRequestBodyMember2able)() {
    m.cardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2 = value
}
// SetColumnsItemCardsPostRequestBodyMember1 sets the ColumnsItemCardsPostRequestBodyMember1 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember1able
func (m *CardsPostRequestBody) SetColumnsItemCardsPostRequestBodyMember1(value ColumnsItemCardsPostRequestBodyMember1able)() {
    m.columnsItemCardsPostRequestBodyMember1 = value
}
// SetColumnsItemCardsPostRequestBodyMember2 sets the ColumnsItemCardsPostRequestBodyMember2 property value. Composed type representation for type ColumnsItemCardsPostRequestBodyMember2able
func (m *CardsPostRequestBody) SetColumnsItemCardsPostRequestBodyMember2(value ColumnsItemCardsPostRequestBodyMember2able)() {
    m.columnsItemCardsPostRequestBodyMember2 = value
}
// ColumnsItemCardsRequestBuilderGetQueryParameters lists the project cards in a project.
type ColumnsItemCardsRequestBuilderGetQueryParameters struct {
    // Filters the project cards that are returned by the card's state.
    Archived_state *iaecb1d9ce8f039d73681b50025b6b1a10860720b981e5b38dba5e4bcaf58ea7c.GetArchived_stateQueryParameterType `uriparametername:"archived_state"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
}
type CardsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1()(ColumnsItemCardsPostRequestBodyMember1able)
    GetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2()(ColumnsItemCardsPostRequestBodyMember2able)
    GetColumnsItemCardsPostRequestBodyMember1()(ColumnsItemCardsPostRequestBodyMember1able)
    GetColumnsItemCardsPostRequestBodyMember2()(ColumnsItemCardsPostRequestBodyMember2able)
    SetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember1(value ColumnsItemCardsPostRequestBodyMember1able)()
    SetCardsPostRequestBodyColumnsItemCardsPostRequestBodyMember2(value ColumnsItemCardsPostRequestBodyMember2able)()
    SetColumnsItemCardsPostRequestBodyMember1(value ColumnsItemCardsPostRequestBodyMember1able)()
    SetColumnsItemCardsPostRequestBodyMember2(value ColumnsItemCardsPostRequestBodyMember2able)()
}
// NewColumnsItemCardsRequestBuilderInternal instantiates a new ColumnsItemCardsRequestBuilder and sets the default values.
func NewColumnsItemCardsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ColumnsItemCardsRequestBuilder) {
    m := &ColumnsItemCardsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/projects/columns/{column_id}/cards{?archived_state*,page*,per_page*}", pathParameters),
    }
    return m
}
// NewColumnsItemCardsRequestBuilder instantiates a new ColumnsItemCardsRequestBuilder and sets the default values.
func NewColumnsItemCardsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ColumnsItemCardsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewColumnsItemCardsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the project cards in a project.
// returns a []ProjectCardable when successful
// returns a BasicError error when the service returns a 401 status code
// returns a BasicError error when the service returns a 403 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/projects/cards#list-project-cards
func (m *ColumnsItemCardsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ColumnsItemCardsRequestBuilderGetQueryParameters])([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ProjectCardable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateProjectCardFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ProjectCardable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ProjectCardable)
        }
    }
    return val, nil
}
// Post create a project card
// returns a ProjectCardable when successful
// returns a BasicError error when the service returns a 401 status code
// returns a BasicError error when the service returns a 403 status code
// returns a ColumnsItemCardsProjectCard503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/projects/cards#create-a-project-card
func (m *ColumnsItemCardsRequestBuilder) Post(ctx context.Context, body CardsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ProjectCardable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "401": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "503": CreateColumnsItemCardsProjectCard503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateProjectCardFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ProjectCardable), nil
}
// ToGetRequestInformation lists the project cards in a project.
// returns a *RequestInformation when successful
func (m *ColumnsItemCardsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ColumnsItemCardsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// returns a *RequestInformation when successful
func (m *ColumnsItemCardsRequestBuilder) ToPostRequestInformation(ctx context.Context, body CardsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ColumnsItemCardsRequestBuilder when successful
func (m *ColumnsItemCardsRequestBuilder) WithUrl(rawUrl string)(*ColumnsItemCardsRequestBuilder) {
    return NewColumnsItemCardsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
