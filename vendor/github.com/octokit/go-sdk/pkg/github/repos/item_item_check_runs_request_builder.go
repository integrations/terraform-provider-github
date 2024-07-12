package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemCheckRunsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\check-runs
type ItemItemCheckRunsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CheckRunsPostRequestBody composed type wrapper for classes ItemItemCheckRunsPostRequestBodyMember1able, ItemItemCheckRunsPostRequestBodyMember2able
type CheckRunsPostRequestBody struct {
    // Composed type representation for type ItemItemCheckRunsPostRequestBodyMember1able
    checkRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember1 ItemItemCheckRunsPostRequestBodyMember1able
    // Composed type representation for type ItemItemCheckRunsPostRequestBodyMember2able
    checkRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember2 ItemItemCheckRunsPostRequestBodyMember2able
    // Composed type representation for type ItemItemCheckRunsPostRequestBodyMember1able
    itemItemCheckRunsPostRequestBodyMember1 ItemItemCheckRunsPostRequestBodyMember1able
    // Composed type representation for type ItemItemCheckRunsPostRequestBodyMember2able
    itemItemCheckRunsPostRequestBodyMember2 ItemItemCheckRunsPostRequestBodyMember2able
}
// NewCheckRunsPostRequestBody instantiates a new CheckRunsPostRequestBody and sets the default values.
func NewCheckRunsPostRequestBody()(*CheckRunsPostRequestBody) {
    m := &CheckRunsPostRequestBody{
    }
    return m
}
// CreateCheckRunsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCheckRunsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewCheckRunsPostRequestBody()
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
// GetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember1 gets the ItemItemCheckRunsPostRequestBodyMember1 property value. Composed type representation for type ItemItemCheckRunsPostRequestBodyMember1able
// returns a ItemItemCheckRunsPostRequestBodyMember1able when successful
func (m *CheckRunsPostRequestBody) GetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember1()(ItemItemCheckRunsPostRequestBodyMember1able) {
    return m.checkRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember1
}
// GetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember2 gets the ItemItemCheckRunsPostRequestBodyMember2 property value. Composed type representation for type ItemItemCheckRunsPostRequestBodyMember2able
// returns a ItemItemCheckRunsPostRequestBodyMember2able when successful
func (m *CheckRunsPostRequestBody) GetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember2()(ItemItemCheckRunsPostRequestBodyMember2able) {
    return m.checkRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember2
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CheckRunsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *CheckRunsPostRequestBody) GetIsComposedType()(bool) {
    return true
}
// GetItemItemCheckRunsPostRequestBodyMember1 gets the ItemItemCheckRunsPostRequestBodyMember1 property value. Composed type representation for type ItemItemCheckRunsPostRequestBodyMember1able
// returns a ItemItemCheckRunsPostRequestBodyMember1able when successful
func (m *CheckRunsPostRequestBody) GetItemItemCheckRunsPostRequestBodyMember1()(ItemItemCheckRunsPostRequestBodyMember1able) {
    return m.itemItemCheckRunsPostRequestBodyMember1
}
// GetItemItemCheckRunsPostRequestBodyMember2 gets the ItemItemCheckRunsPostRequestBodyMember2 property value. Composed type representation for type ItemItemCheckRunsPostRequestBodyMember2able
// returns a ItemItemCheckRunsPostRequestBodyMember2able when successful
func (m *CheckRunsPostRequestBody) GetItemItemCheckRunsPostRequestBodyMember2()(ItemItemCheckRunsPostRequestBodyMember2able) {
    return m.itemItemCheckRunsPostRequestBodyMember2
}
// Serialize serializes information the current object
func (m *CheckRunsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember2())
        if err != nil {
            return err
        }
    } else if m.GetItemItemCheckRunsPostRequestBodyMember1() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemCheckRunsPostRequestBodyMember1())
        if err != nil {
            return err
        }
    } else if m.GetItemItemCheckRunsPostRequestBodyMember2() != nil {
        err := writer.WriteObjectValue("", m.GetItemItemCheckRunsPostRequestBodyMember2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember1 sets the ItemItemCheckRunsPostRequestBodyMember1 property value. Composed type representation for type ItemItemCheckRunsPostRequestBodyMember1able
func (m *CheckRunsPostRequestBody) SetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember1(value ItemItemCheckRunsPostRequestBodyMember1able)() {
    m.checkRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember1 = value
}
// SetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember2 sets the ItemItemCheckRunsPostRequestBodyMember2 property value. Composed type representation for type ItemItemCheckRunsPostRequestBodyMember2able
func (m *CheckRunsPostRequestBody) SetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember2(value ItemItemCheckRunsPostRequestBodyMember2able)() {
    m.checkRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember2 = value
}
// SetItemItemCheckRunsPostRequestBodyMember1 sets the ItemItemCheckRunsPostRequestBodyMember1 property value. Composed type representation for type ItemItemCheckRunsPostRequestBodyMember1able
func (m *CheckRunsPostRequestBody) SetItemItemCheckRunsPostRequestBodyMember1(value ItemItemCheckRunsPostRequestBodyMember1able)() {
    m.itemItemCheckRunsPostRequestBodyMember1 = value
}
// SetItemItemCheckRunsPostRequestBodyMember2 sets the ItemItemCheckRunsPostRequestBodyMember2 property value. Composed type representation for type ItemItemCheckRunsPostRequestBodyMember2able
func (m *CheckRunsPostRequestBody) SetItemItemCheckRunsPostRequestBodyMember2(value ItemItemCheckRunsPostRequestBodyMember2able)() {
    m.itemItemCheckRunsPostRequestBodyMember2 = value
}
type CheckRunsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember1()(ItemItemCheckRunsPostRequestBodyMember1able)
    GetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember2()(ItemItemCheckRunsPostRequestBodyMember2able)
    GetItemItemCheckRunsPostRequestBodyMember1()(ItemItemCheckRunsPostRequestBodyMember1able)
    GetItemItemCheckRunsPostRequestBodyMember2()(ItemItemCheckRunsPostRequestBodyMember2able)
    SetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember1(value ItemItemCheckRunsPostRequestBodyMember1able)()
    SetCheckRunsPostRequestBodyItemItemCheckRunsPostRequestBodyMember2(value ItemItemCheckRunsPostRequestBodyMember2able)()
    SetItemItemCheckRunsPostRequestBodyMember1(value ItemItemCheckRunsPostRequestBodyMember1able)()
    SetItemItemCheckRunsPostRequestBodyMember2(value ItemItemCheckRunsPostRequestBodyMember2able)()
}
// ByCheck_run_id gets an item from the github.com/octokit/go-sdk/pkg/github.repos.item.item.checkRuns.item collection
// returns a *ItemItemCheckRunsWithCheck_run_ItemRequestBuilder when successful
func (m *ItemItemCheckRunsRequestBuilder) ByCheck_run_id(check_run_id int32)(*ItemItemCheckRunsWithCheck_run_ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["check_run_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(check_run_id), 10)
    return NewItemItemCheckRunsWithCheck_run_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemCheckRunsRequestBuilderInternal instantiates a new ItemItemCheckRunsRequestBuilder and sets the default values.
func NewItemItemCheckRunsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckRunsRequestBuilder) {
    m := &ItemItemCheckRunsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/check-runs", pathParameters),
    }
    return m
}
// NewItemItemCheckRunsRequestBuilder instantiates a new ItemItemCheckRunsRequestBuilder and sets the default values.
func NewItemItemCheckRunsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCheckRunsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCheckRunsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post creates a new check run for a specific commit in a repository.To create a check run, you must use a GitHub App. OAuth apps and authenticated users are not able to create a check suite.In a check suite, GitHub limits the number of check runs with the same name to 1000. Once these check runs exceed 1000, GitHub will start to automatically delete older check runs.**Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.
// returns a CheckRunable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/checks/runs#create-a-check-run
func (m *ItemItemCheckRunsRequestBuilder) Post(ctx context.Context, body CheckRunsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CheckRunable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation creates a new check run for a specific commit in a repository.To create a check run, you must use a GitHub App. OAuth apps and authenticated users are not able to create a check suite.In a check suite, GitHub limits the number of check runs with the same name to 1000. Once these check runs exceed 1000, GitHub will start to automatically delete older check runs.**Note:** The Checks API only looks for pushes in the repository where the check suite or check run were created. Pushes to a branch in a forked repository are not detected and return an empty `pull_requests` array.
// returns a *RequestInformation when successful
func (m *ItemItemCheckRunsRequestBuilder) ToPostRequestInformation(ctx context.Context, body CheckRunsPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemCheckRunsRequestBuilder when successful
func (m *ItemItemCheckRunsRequestBuilder) WithUrl(rawUrl string)(*ItemItemCheckRunsRequestBuilder) {
    return NewItemItemCheckRunsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
