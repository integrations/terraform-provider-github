package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemCheckRunsWithCheck_run_ItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\check-runs\{check_run_id}
type ItemItemCheckRunsWithCheck_run_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WithCheck_run_PatchRequestBody composed type wrapper for classes ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able, ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able
type WithCheck_run_PatchRequestBody struct {
    // Composed type representation for type ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able
    itemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1 ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able
    // Composed type representation for type ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able
    itemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2 ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able
    // Composed type representation for type ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able
    withCheck_run_PatchRequestBodyItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1 ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able
    // Composed type representation for type ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able
    withCheck_run_PatchRequestBodyItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2 ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able
}
// NewWithCheck_run_PatchRequestBody instantiates a new WithCheck_run_PatchRequestBody and sets the default values.
func NewWithCheck_run_PatchRequestBody()(*WithCheck_run_PatchRequestBody) {
    m := &WithCheck_run_PatchRequestBody{
    }
    return m
}
// CreateWithCheck_run_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWithCheck_run_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWithCheck_run_PatchRequestBody()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able); ok {
                result.SetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able); ok {
                result.SetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able); ok {
                result.SetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2FromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able); ok {
                result.SetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2(cast)
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WithCheck_run_PatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *WithCheck_run_PatchRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1 gets the ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1 property value. Composed type representation for type ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able
// returns a ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able when successful
func (m *WithCheck_run_PatchRequestBody) GetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1()(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able) {
    return m.itemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1
}
// GetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2 gets the ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2 property value. Composed type representation for type ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able
// returns a ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able when successful
func (m *WithCheck_run_PatchRequestBody) GetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2()(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able) {
    return m.itemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2
}
// GetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1 gets the ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1 property value. Composed type representation for type ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able
// returns a ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able when successful
func (m *WithCheck_run_PatchRequestBody) GetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1()(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able) {
    return m.withCheck_run_PatchRequestBodyItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1
}
// GetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2 gets the ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2 property value. Composed type representation for type ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able
// returns a ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able when successful
func (m *WithCheck_run_PatchRequestBody) GetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2()(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able) {
    return m.withCheck_run_PatchRequestBodyItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2
}
// Serialize serializes information the current object
func (m *WithCheck_run_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2())
        if err != nil {
            return err
        }
    } else if m.GetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1 sets the ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1 property value. Composed type representation for type ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able
func (m *WithCheck_run_PatchRequestBody) SetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1(value ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able)() {
    m.itemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1 = value
}
// SetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2 sets the ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2 property value. Composed type representation for type ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able
func (m *WithCheck_run_PatchRequestBody) SetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2(value ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able)() {
    m.itemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2 = value
}
// SetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1 sets the ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1 property value. Composed type representation for type ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able
func (m *WithCheck_run_PatchRequestBody) SetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1(value ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able)() {
    m.withCheck_run_PatchRequestBodyItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1 = value
}
// SetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2 sets the ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2 property value. Composed type representation for type ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able
func (m *WithCheck_run_PatchRequestBody) SetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2(value ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able)() {
    m.withCheck_run_PatchRequestBodyItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2 = value
}
type WithCheck_run_PatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1()(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able)
    GetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2()(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able)
    GetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1()(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able)
    GetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2()(ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able)
    SetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1(value ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able)()
    SetItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2(value ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able)()
    SetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember1(value ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember1able)()
    SetWithCheckRunPatchRequestBodyItemItemCheckRunsItemWithCheckRunPatchRequestBodyMember2(value ItemItemCheckRunsItemWithCheck_run_PatchRequestBodyMember2able)()
}
// Annotations the annotations property
// returns a *ItemItemCheckRunsItemAnnotationsRequestBuilder when successful
func (m *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) Annotations()(*ItemItemCheckRunsItemAnnotationsRequestBuilder) {
    return NewItemItemCheckRunsItemAnnotationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCheckRunsWithCheck_run_ItemRequestBuilderInternal instantiates a new ItemItemCheckRunsWithCheck_run_ItemRequestBuilder and sets the default values.
func NewItemItemCheckRunsWithCheck_run_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) {
    m := &ItemItemCheckRunsWithCheck_run_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/check-runs/{check_run_id}", pathParameters),
    }
    return m
}
// NewItemItemCheckRunsWithCheck_run_ItemRequestBuilder instantiates a new ItemItemCheckRunsWithCheck_run_ItemRequestBuilder and sets the default values.
func NewItemItemCheckRunsWithCheck_run_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCheckRunsWithCheck_run_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets a single check run using its `id`.**Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint on a private repository.
// returns a CheckRunable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/checks/runs#get-a-check-run
func (m *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CheckRunable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCheckRunFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CheckRunable), nil
}
// Patch updates a check run for a specific commit in a repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.OAuth apps and personal access tokens (classic) cannot use this endpoint.
// returns a CheckRunable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/checks/runs#update-a-check-run
func (m *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) Patch(ctx context.Context, body WithCheck_run_PatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CheckRunable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCheckRunFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CheckRunable), nil
}
// Rerequest the rerequest property
// returns a *ItemItemCheckRunsItemRerequestRequestBuilder when successful
func (m *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) Rerequest()(*ItemItemCheckRunsItemRerequestRequestBuilder) {
    return NewItemItemCheckRunsItemRerequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation gets a single check run using its `id`.**Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint on a private repository.
// returns a *RequestInformation when successful
func (m *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation updates a check run for a specific commit in a repository.**Note:** The endpoints to manage checks only look for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.OAuth apps and personal access tokens (classic) cannot use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body WithCheck_run_PatchRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder when successful
func (m *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) {
    return NewItemItemCheckRunsWithCheck_run_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
