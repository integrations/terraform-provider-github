package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\code-scanning\analyses\{analysis_id}
type ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderDeleteQueryParameters deletes a specified code scanning analysis from a repository.You can delete one analysis at a time.To delete a series of analyses, start with the most recent analysis and work backwards.Conceptually, the process is similar to the undo function in a text editor.When you list the analyses for a repository,one or more will be identified as deletable in the response:```"deletable": true```An analysis is deletable when it's the most recent in a set of analyses.Typically, a repository will have multiple sets of analysesfor each enabled code scanning tool,where a set is determined by a unique combination of analysis values:* `ref`* `tool`* `category`If you attempt to delete an analysis that is not the most recent in a set,you'll get a 400 response with the message:```Analysis specified is not deletable.```The response from a successful `DELETE` operation provides you withtwo alternative URLs for deleting the next analysis in the set:`next_analysis_url` and `confirm_delete_url`.Use the `next_analysis_url` URL if you want to avoid accidentally deleting the final analysisin a set. This is a useful option if you want to preserve at least one analysisfor the specified tool in your repository.Use the `confirm_delete_url` URL if you are content to remove all analyses for a tool.When you delete the last analysis in a set, the value of `next_analysis_url` and `confirm_delete_url`in the 200 response is `null`.As an example of the deletion process,let's imagine that you added a workflow that configured a particular code scanning toolto analyze the code in a repository. This tool has added 15 analyses:10 on the default branch, and another 5 on a topic branch.You therefore have two separate sets of analyses for this tool.You've now decided that you want to remove all of the analyses for the tool.To do this you must make 15 separate deletion requests.To start, you must find an analysis that's identified as deletable.Each set of analyses always has one that's identified as deletable.Having found the deletable analysis for one of the two sets,delete this analysis and then continue deleting the next analysis in the set until they're all deleted.Then repeat the process for the second set.The procedure therefore consists of a nested loop:**Outer loop**:* List the analyses for the repository, filtered by tool.* Parse this list to find a deletable analysis. If found:  **Inner loop**:  * Delete the identified analysis.  * Parse the response for the value of `confirm_delete_url` and, if found, use this in the next iteration.The above process assumes that you want to remove all trace of the tool's analyses from the GitHub user interface, for the specified repository, and it therefore uses the `confirm_delete_url` value. Alternatively, you could use the `next_analysis_url` value, which would leave the last analysis in each set undeleted to avoid removing a tool's analysis entirely.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
type ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderDeleteQueryParameters struct {
    // Allow deletion if the specified analysis is the last in a set. If you attempt to delete the final analysis in a set without setting this parameter to `true`, you'll get a 400 response with the message: `Analysis is last of its type and deletion may result in the loss of historical alert data. Please specify confirm_delete.`
    Confirm_delete *string `uriparametername:"confirm_delete"`
}
// NewItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderInternal instantiates a new ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder and sets the default values.
func NewItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) {
    m := &ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/code-scanning/analyses/{analysis_id}{?confirm_delete*}", pathParameters),
    }
    return m
}
// NewItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder instantiates a new ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder and sets the default values.
func NewItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a specified code scanning analysis from a repository.You can delete one analysis at a time.To delete a series of analyses, start with the most recent analysis and work backwards.Conceptually, the process is similar to the undo function in a text editor.When you list the analyses for a repository,one or more will be identified as deletable in the response:```"deletable": true```An analysis is deletable when it's the most recent in a set of analyses.Typically, a repository will have multiple sets of analysesfor each enabled code scanning tool,where a set is determined by a unique combination of analysis values:* `ref`* `tool`* `category`If you attempt to delete an analysis that is not the most recent in a set,you'll get a 400 response with the message:```Analysis specified is not deletable.```The response from a successful `DELETE` operation provides you withtwo alternative URLs for deleting the next analysis in the set:`next_analysis_url` and `confirm_delete_url`.Use the `next_analysis_url` URL if you want to avoid accidentally deleting the final analysisin a set. This is a useful option if you want to preserve at least one analysisfor the specified tool in your repository.Use the `confirm_delete_url` URL if you are content to remove all analyses for a tool.When you delete the last analysis in a set, the value of `next_analysis_url` and `confirm_delete_url`in the 200 response is `null`.As an example of the deletion process,let's imagine that you added a workflow that configured a particular code scanning toolto analyze the code in a repository. This tool has added 15 analyses:10 on the default branch, and another 5 on a topic branch.You therefore have two separate sets of analyses for this tool.You've now decided that you want to remove all of the analyses for the tool.To do this you must make 15 separate deletion requests.To start, you must find an analysis that's identified as deletable.Each set of analyses always has one that's identified as deletable.Having found the deletable analysis for one of the two sets,delete this analysis and then continue deleting the next analysis in the set until they're all deleted.Then repeat the process for the second set.The procedure therefore consists of a nested loop:**Outer loop**:* List the analyses for the repository, filtered by tool.* Parse this list to find a deletable analysis. If found:  **Inner loop**:  * Delete the identified analysis.  * Parse the response for the value of `confirm_delete_url` and, if found, use this in the next iteration.The above process assumes that you want to remove all trace of the tool's analyses from the GitHub user interface, for the specified repository, and it therefore uses the `confirm_delete_url` value. Alternatively, you could use the `next_analysis_url` value, which would leave the last analysis in each set undeleted to avoid removing a tool's analysis entirely.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a CodeScanningAnalysisDeletionable when successful
// returns a BasicError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a CodeScanningAnalysisDeletion503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#delete-a-code-scanning-analysis-from-a-repository
func (m *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderDeleteQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAnalysisDeletionable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningAnalysisDeletion503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningAnalysisDeletionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAnalysisDeletionable), nil
}
// Get gets a specified code scanning analysis for a repository.The default JSON response contains fields that describe the analysis.This includes the Git reference and commit SHA to which the analysis relates,the datetime of the analysis, the name of the code scanning tool,and the number of alerts.The `rules_count` field in the default response give the number of rulesthat were run in the analysis.For very old analyses this data is not available,and `0` is returned in this field.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/sarif+json`**: Instead of returning a summary of the analysis, this endpoint returns a subset of the analysis data that was uploaded. The data is formatted as [SARIF version 2.1.0](https://docs.oasis-open.org/sarif/sarif/v2.1.0/cs01/sarif-v2.1.0-cs01.html). It also returns additional data such as the `github/alertNumber` and `github/alertUrl` properties.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a CodeScanningAnalysisable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a CodeScanningAnalysis503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/code-scanning/code-scanning#get-a-code-scanning-analysis-for-a-repository
func (m *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAnalysisable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningAnalysis503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCodeScanningAnalysisFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningAnalysisable), nil
}
// ToDeleteRequestInformation deletes a specified code scanning analysis from a repository.You can delete one analysis at a time.To delete a series of analyses, start with the most recent analysis and work backwards.Conceptually, the process is similar to the undo function in a text editor.When you list the analyses for a repository,one or more will be identified as deletable in the response:```"deletable": true```An analysis is deletable when it's the most recent in a set of analyses.Typically, a repository will have multiple sets of analysesfor each enabled code scanning tool,where a set is determined by a unique combination of analysis values:* `ref`* `tool`* `category`If you attempt to delete an analysis that is not the most recent in a set,you'll get a 400 response with the message:```Analysis specified is not deletable.```The response from a successful `DELETE` operation provides you withtwo alternative URLs for deleting the next analysis in the set:`next_analysis_url` and `confirm_delete_url`.Use the `next_analysis_url` URL if you want to avoid accidentally deleting the final analysisin a set. This is a useful option if you want to preserve at least one analysisfor the specified tool in your repository.Use the `confirm_delete_url` URL if you are content to remove all analyses for a tool.When you delete the last analysis in a set, the value of `next_analysis_url` and `confirm_delete_url`in the 200 response is `null`.As an example of the deletion process,let's imagine that you added a workflow that configured a particular code scanning toolto analyze the code in a repository. This tool has added 15 analyses:10 on the default branch, and another 5 on a topic branch.You therefore have two separate sets of analyses for this tool.You've now decided that you want to remove all of the analyses for the tool.To do this you must make 15 separate deletion requests.To start, you must find an analysis that's identified as deletable.Each set of analyses always has one that's identified as deletable.Having found the deletable analysis for one of the two sets,delete this analysis and then continue deleting the next analysis in the set until they're all deleted.Then repeat the process for the second set.The procedure therefore consists of a nested loop:**Outer loop**:* List the analyses for the repository, filtered by tool.* Parse this list to find a deletable analysis. If found:  **Inner loop**:  * Delete the identified analysis.  * Parse the response for the value of `confirm_delete_url` and, if found, use this in the next iteration.The above process assumes that you want to remove all trace of the tool's analyses from the GitHub user interface, for the specified repository, and it therefore uses the `confirm_delete_url` value. Alternatively, you could use the `next_analysis_url` value, which would leave the last analysis in each set undeleted to avoid removing a tool's analysis entirely.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a *RequestInformation when successful
func (m *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilderDeleteQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation gets a specified code scanning analysis for a repository.The default JSON response contains fields that describe the analysis.This includes the Git reference and commit SHA to which the analysis relates,the datetime of the analysis, the name of the code scanning tool,and the number of alerts.The `rules_count` field in the default response give the number of rulesthat were run in the analysis.For very old analyses this data is not available,and `0` is returned in this field.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/sarif+json`**: Instead of returning a summary of the analysis, this endpoint returns a subset of the analysis data that was uploaded. The data is formatted as [SARIF version 2.1.0](https://docs.oasis-open.org/sarif/sarif/v2.1.0/cs01/sarif-v2.1.0-cs01.html). It also returns additional data such as the `github/alertNumber` and `github/alertUrl` properties.OAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint with private or public repositories, or the `public_repo` scope to use this endpoint with only public repositories.
// returns a *RequestInformation when successful
func (m *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder when successful
func (m *ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder) {
    return NewItemItemCodeScanningAnalysesWithAnalysis_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
