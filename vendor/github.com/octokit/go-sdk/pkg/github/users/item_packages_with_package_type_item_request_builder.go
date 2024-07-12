package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemPackagesWithPackage_typeItemRequestBuilder builds and executes requests for operations under \users\{username}\packages\{package_type}
type ItemPackagesWithPackage_typeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByPackage_name gets an item from the github.com/octokit/go-sdk/pkg/github.users.item.packages.item.item collection
// returns a *ItemPackagesItemWithPackage_nameItemRequestBuilder when successful
func (m *ItemPackagesWithPackage_typeItemRequestBuilder) ByPackage_name(package_name string)(*ItemPackagesItemWithPackage_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if package_name != "" {
        urlTplParams["package_name"] = package_name
    }
    return NewItemPackagesItemWithPackage_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPackagesWithPackage_typeItemRequestBuilderInternal instantiates a new ItemPackagesWithPackage_typeItemRequestBuilder and sets the default values.
func NewItemPackagesWithPackage_typeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPackagesWithPackage_typeItemRequestBuilder) {
    m := &ItemPackagesWithPackage_typeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{username}/packages/{package_type}", pathParameters),
    }
    return m
}
// NewItemPackagesWithPackage_typeItemRequestBuilder instantiates a new ItemPackagesWithPackage_typeItemRequestBuilder and sets the default values.
func NewItemPackagesWithPackage_typeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPackagesWithPackage_typeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPackagesWithPackage_typeItemRequestBuilderInternal(urlParams, requestAdapter)
}
