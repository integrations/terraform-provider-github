package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

// ItemItemContentsWithPathItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\contents\{path}
type ItemItemContentsWithPathItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemContentsWithPathItemRequestBuilderGetQueryParameters gets the contents of a file or directory in a repository. Specify the file path or directory with the `path` parameter. If you omit the `path` parameter, you will receive the contents of the repository's root directory.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw file contents for files and symlinks.- **`application/vnd.github.html+json`**: Returns the file contents in HTML. Markup languages are rendered to HTML using GitHub's open-source [Markup library](https://github.com/github/markup).- **`application/vnd.github.object+json`**: Returns the contents in a consistent object format regardless of the content type. For example, instead of an array of objects for a directory, the response will be an object with an `entries` attribute containing the array of objects.If the content is a directory, the response will be an array of objects, one object for each item in the directory. When listing the contents of a directory, submodules have their "type" specified as "file". Logically, the value _should_ be "submodule". This behavior exists [for backwards compatibility purposes](https://git.io/v1YCW). In the next major version of the API, the type will be returned as "submodule".If the content is a symlink and the symlink's target is a normal file in the repository, then the API responds with the content of the file. Otherwise, the API responds with an object describing the symlink itself.If the content is a submodule, the `submodule_git_url` field identifies the location of the submodule repository, and the `sha` identifies a specific commit within the submodule repository. Git uses the given URL when cloning the submodule repository, and checks out the submodule at that specific commit. If the submodule repository is not hosted on github.com, the Git URLs (`git_url` and `_links["git"]`) and the github.com URLs (`html_url` and `_links["html"]`) will have null values.**Notes**:- To get a repository's contents recursively, you can [recursively get the tree](https://docs.github.com/rest/git/trees#get-a-tree).- This API has an upper limit of 1,000 files for a directory. If you need to retrievemore files, use the [Git Trees API](https://docs.github.com/rest/git/trees#get-a-tree).- Download URLs expire and are meant to be used just once. To ensure the download URL does not expire, please use the contents API to obtain a fresh download URL for each download.- If the requested file's size is:  - 1 MB or smaller: All features of this endpoint are supported.  - Between 1-100 MB: Only the `raw` or `object` custom media types are supported. Both will work as normal, except that when using the `object` media type, the `content` field will be an emptystring and the `encoding` field will be `"none"`. To get the contents of these larger files, use the `raw` media type.  - Greater than 100 MB: This endpoint is not supported.
type ItemItemContentsWithPathItemRequestBuilderGetQueryParameters struct {
    // The name of the commit/branch/tag. Default: the repository’s default branch.
    Ref *string `uriparametername:"ref"`
}
// WithPathGetResponse composed type wrapper for classes i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentFileable, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSubmoduleable, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSymlinkable, []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WithPathGetResponseMember1able
type WithPathGetResponse struct {
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentFileable
    contentFile i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentFileable
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSubmoduleable
    contentSubmodule i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSubmoduleable
    // Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSymlinkable
    contentSymlink i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSymlinkable
    // Composed type representation for type []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WithPathGetResponseMember1able
    withPathGetResponseMember1 []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WithPathGetResponseMember1able
}
// NewWithPathGetResponse instantiates a new WithPathGetResponse and sets the default values.
func NewWithPathGetResponse()(*WithPathGetResponse) {
    m := &WithPathGetResponse{
    }
    return m
}
// CreateWithPathGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWithPathGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWithPathGetResponse()
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
    if val, err := parseNode.GetCollectionOfObjectValues(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateWithPathGetResponseMember1FromDiscriminatorValue); val != nil {
        if err != nil {
            return nil, err
        }
        cast := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WithPathGetResponseMember1able, len(val))
        for i, v := range val {
            if v != nil {
                cast[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WithPathGetResponseMember1able)
            }
        }
        result.SetWithPathGetResponseMember1(cast)
    }
    return result, nil
}
// GetContentFile gets the contentFile property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentFileable
// returns a ContentFileable when successful
func (m *WithPathGetResponse) GetContentFile()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentFileable) {
    return m.contentFile
}
// GetContentSubmodule gets the contentSubmodule property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSubmoduleable
// returns a ContentSubmoduleable when successful
func (m *WithPathGetResponse) GetContentSubmodule()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSubmoduleable) {
    return m.contentSubmodule
}
// GetContentSymlink gets the contentSymlink property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSymlinkable
// returns a ContentSymlinkable when successful
func (m *WithPathGetResponse) GetContentSymlink()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSymlinkable) {
    return m.contentSymlink
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WithPathGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *WithPathGetResponse) GetIsComposedType()(bool) {
    return true
}
// GetWithPathGetResponseMember1 gets the WithPathGetResponseMember1 property value. Composed type representation for type []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WithPathGetResponseMember1able
// returns a []WithPathGetResponseMember1able when successful
func (m *WithPathGetResponse) GetWithPathGetResponseMember1()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WithPathGetResponseMember1able) {
    return m.withPathGetResponseMember1
}
// Serialize serializes information the current object
func (m *WithPathGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContentFile() != nil {
        err := writer.WriteObjectValue("", m.GetContentFile())
        if err != nil {
            return err
        }
    } else if m.GetContentSubmodule() != nil {
        err := writer.WriteObjectValue("", m.GetContentSubmodule())
        if err != nil {
            return err
        }
    } else if m.GetContentSymlink() != nil {
        err := writer.WriteObjectValue("", m.GetContentSymlink())
        if err != nil {
            return err
        }
    } else if m.GetWithPathGetResponseMember1() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWithPathGetResponseMember1()))
        for i, v := range m.GetWithPathGetResponseMember1() {
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
// SetContentFile sets the contentFile property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentFileable
func (m *WithPathGetResponse) SetContentFile(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentFileable)() {
    m.contentFile = value
}
// SetContentSubmodule sets the contentSubmodule property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSubmoduleable
func (m *WithPathGetResponse) SetContentSubmodule(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSubmoduleable)() {
    m.contentSubmodule = value
}
// SetContentSymlink sets the contentSymlink property value. Composed type representation for type i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSymlinkable
func (m *WithPathGetResponse) SetContentSymlink(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSymlinkable)() {
    m.contentSymlink = value
}
// SetWithPathGetResponseMember1 sets the WithPathGetResponseMember1 property value. Composed type representation for type []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WithPathGetResponseMember1able
func (m *WithPathGetResponse) SetWithPathGetResponseMember1(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WithPathGetResponseMember1able)() {
    m.withPathGetResponseMember1 = value
}
type WithPathGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentFile()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentFileable)
    GetContentSubmodule()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSubmoduleable)
    GetContentSymlink()(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSymlinkable)
    GetWithPathGetResponseMember1()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WithPathGetResponseMember1able)
    SetContentFile(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentFileable)()
    SetContentSubmodule(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSubmoduleable)()
    SetContentSymlink(value i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ContentSymlinkable)()
    SetWithPathGetResponseMember1(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WithPathGetResponseMember1able)()
}
// NewItemItemContentsWithPathItemRequestBuilderInternal instantiates a new ItemItemContentsWithPathItemRequestBuilder and sets the default values.
func NewItemItemContentsWithPathItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemContentsWithPathItemRequestBuilder) {
    m := &ItemItemContentsWithPathItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/contents/{path}{?ref*}", pathParameters),
    }
    return m
}
// NewItemItemContentsWithPathItemRequestBuilder instantiates a new ItemItemContentsWithPathItemRequestBuilder and sets the default values.
func NewItemItemContentsWithPathItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemContentsWithPathItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemContentsWithPathItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a file in a repository.You can provide an additional `committer` parameter, which is an object containing information about the committer. Or, you can provide an `author` parameter, which is an object containing information about the author.The `author` section is optional and is filled in with the `committer` information if omitted. If the `committer` information is omitted, the authenticated user's information is used.You must provide values for both `name` and `email`, whether you choose to use `author` or `committer`. Otherwise, you'll receive a `422` status code.**Note:** If you use this endpoint and the "[Create or update file contents](https://docs.github.com/rest/repos/contents/#create-or-update-file-contents)" endpoint in parallel, the concurrent requests will conflict and you will receive errors. You must use these endpoints serially instead.
// returns a FileCommitable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 409 status code
// returns a ValidationError error when the service returns a 422 status code
// returns a FileCommit503Error error when the service returns a 503 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/contents#delete-a-file
func (m *ItemItemContentsWithPathItemRequestBuilder) Delete(ctx context.Context, body ItemItemContentsItemWithPathDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.FileCommitable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "409": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
        "503": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateFileCommit503ErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateFileCommitFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.FileCommitable), nil
}
// Get gets the contents of a file or directory in a repository. Specify the file path or directory with the `path` parameter. If you omit the `path` parameter, you will receive the contents of the repository's root directory.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw file contents for files and symlinks.- **`application/vnd.github.html+json`**: Returns the file contents in HTML. Markup languages are rendered to HTML using GitHub's open-source [Markup library](https://github.com/github/markup).- **`application/vnd.github.object+json`**: Returns the contents in a consistent object format regardless of the content type. For example, instead of an array of objects for a directory, the response will be an object with an `entries` attribute containing the array of objects.If the content is a directory, the response will be an array of objects, one object for each item in the directory. When listing the contents of a directory, submodules have their "type" specified as "file". Logically, the value _should_ be "submodule". This behavior exists [for backwards compatibility purposes](https://git.io/v1YCW). In the next major version of the API, the type will be returned as "submodule".If the content is a symlink and the symlink's target is a normal file in the repository, then the API responds with the content of the file. Otherwise, the API responds with an object describing the symlink itself.If the content is a submodule, the `submodule_git_url` field identifies the location of the submodule repository, and the `sha` identifies a specific commit within the submodule repository. Git uses the given URL when cloning the submodule repository, and checks out the submodule at that specific commit. If the submodule repository is not hosted on github.com, the Git URLs (`git_url` and `_links["git"]`) and the github.com URLs (`html_url` and `_links["html"]`) will have null values.**Notes**:- To get a repository's contents recursively, you can [recursively get the tree](https://docs.github.com/rest/git/trees#get-a-tree).- This API has an upper limit of 1,000 files for a directory. If you need to retrievemore files, use the [Git Trees API](https://docs.github.com/rest/git/trees#get-a-tree).- Download URLs expire and are meant to be used just once. To ensure the download URL does not expire, please use the contents API to obtain a fresh download URL for each download.- If the requested file's size is:  - 1 MB or smaller: All features of this endpoint are supported.  - Between 1-100 MB: Only the `raw` or `object` custom media types are supported. Both will work as normal, except that when using the `object` media type, the `content` field will be an emptystring and the `encoding` field will be `"none"`. To get the contents of these larger files, use the `raw` media type.  - Greater than 100 MB: This endpoint is not supported.
// returns a WithPathGetResponseable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/contents#get-repository-content
func (m *ItemItemContentsWithPathItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemContentsWithPathItemRequestBuilderGetQueryParameters])(WithPathGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateWithPathGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(WithPathGetResponseable), nil
}
// Put creates a new file or replaces an existing file in a repository.**Note:** If you use this endpoint and the "[Delete a file](https://docs.github.com/rest/repos/contents/#delete-a-file)" endpoint in parallel, the concurrent requests will conflict and you will receive errors. You must use these endpoints serially instead.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint. The `workflow` scope is also required in order to modify files in the `.github/workflows` directory.
// returns a FileCommitable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 409 status code
// returns a ValidationError error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/repos/contents#create-or-update-file-contents
func (m *ItemItemContentsWithPathItemRequestBuilder) Put(ctx context.Context, body ItemItemContentsItemWithPathPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.FileCommitable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "409": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
        "422": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateValidationErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateFileCommitFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.FileCommitable), nil
}
// ToDeleteRequestInformation deletes a file in a repository.You can provide an additional `committer` parameter, which is an object containing information about the committer. Or, you can provide an `author` parameter, which is an object containing information about the author.The `author` section is optional and is filled in with the `committer` information if omitted. If the `committer` information is omitted, the authenticated user's information is used.You must provide values for both `name` and `email`, whether you choose to use `author` or `committer`. Otherwise, you'll receive a `422` status code.**Note:** If you use this endpoint and the "[Create or update file contents](https://docs.github.com/rest/repos/contents/#create-or-update-file-contents)" endpoint in parallel, the concurrent requests will conflict and you will receive errors. You must use these endpoints serially instead.
// returns a *RequestInformation when successful
func (m *ItemItemContentsWithPathItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body ItemItemContentsItemWithPathDeleteRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// ToGetRequestInformation gets the contents of a file or directory in a repository. Specify the file path or directory with the `path` parameter. If you omit the `path` parameter, you will receive the contents of the repository's root directory.This endpoint supports the following custom media types. For more information, see "[Media types](https://docs.github.com/rest/using-the-rest-api/getting-started-with-the-rest-api#media-types)."- **`application/vnd.github.raw+json`**: Returns the raw file contents for files and symlinks.- **`application/vnd.github.html+json`**: Returns the file contents in HTML. Markup languages are rendered to HTML using GitHub's open-source [Markup library](https://github.com/github/markup).- **`application/vnd.github.object+json`**: Returns the contents in a consistent object format regardless of the content type. For example, instead of an array of objects for a directory, the response will be an object with an `entries` attribute containing the array of objects.If the content is a directory, the response will be an array of objects, one object for each item in the directory. When listing the contents of a directory, submodules have their "type" specified as "file". Logically, the value _should_ be "submodule". This behavior exists [for backwards compatibility purposes](https://git.io/v1YCW). In the next major version of the API, the type will be returned as "submodule".If the content is a symlink and the symlink's target is a normal file in the repository, then the API responds with the content of the file. Otherwise, the API responds with an object describing the symlink itself.If the content is a submodule, the `submodule_git_url` field identifies the location of the submodule repository, and the `sha` identifies a specific commit within the submodule repository. Git uses the given URL when cloning the submodule repository, and checks out the submodule at that specific commit. If the submodule repository is not hosted on github.com, the Git URLs (`git_url` and `_links["git"]`) and the github.com URLs (`html_url` and `_links["html"]`) will have null values.**Notes**:- To get a repository's contents recursively, you can [recursively get the tree](https://docs.github.com/rest/git/trees#get-a-tree).- This API has an upper limit of 1,000 files for a directory. If you need to retrievemore files, use the [Git Trees API](https://docs.github.com/rest/git/trees#get-a-tree).- Download URLs expire and are meant to be used just once. To ensure the download URL does not expire, please use the contents API to obtain a fresh download URL for each download.- If the requested file's size is:  - 1 MB or smaller: All features of this endpoint are supported.  - Between 1-100 MB: Only the `raw` or `object` custom media types are supported. Both will work as normal, except that when using the `object` media type, the `content` field will be an emptystring and the `encoding` field will be `"none"`. To get the contents of these larger files, use the `raw` media type.  - Greater than 100 MB: This endpoint is not supported.
// returns a *RequestInformation when successful
func (m *ItemItemContentsWithPathItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemContentsWithPathItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPutRequestInformation creates a new file or replaces an existing file in a repository.**Note:** If you use this endpoint and the "[Delete a file](https://docs.github.com/rest/repos/contents/#delete-a-file)" endpoint in parallel, the concurrent requests will conflict and you will receive errors. You must use these endpoints serially instead.OAuth app tokens and personal access tokens (classic) need the `repo` scope to use this endpoint. The `workflow` scope is also required in order to modify files in the `.github/workflows` directory.
// returns a *RequestInformation when successful
func (m *ItemItemContentsWithPathItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body ItemItemContentsItemWithPathPutRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemContentsWithPathItemRequestBuilder when successful
func (m *ItemItemContentsWithPathItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemContentsWithPathItemRequestBuilder) {
    return NewItemItemContentsWithPathItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
