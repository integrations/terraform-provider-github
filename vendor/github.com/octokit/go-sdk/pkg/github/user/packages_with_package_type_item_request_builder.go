package user

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PackagesWithPackage_typeItemRequestBuilder builds and executes requests for operations under \user\packages\{package_type}
type PackagesWithPackage_typeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByPackage_name gets an item from the github.com/octokit/go-sdk/pkg/github.user.packages.item.item collection
// returns a *PackagesItemWithPackage_nameItemRequestBuilder when successful
func (m *PackagesWithPackage_typeItemRequestBuilder) ByPackage_name(package_name string)(*PackagesItemWithPackage_nameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if package_name != "" {
        urlTplParams["package_name"] = package_name
    }
    return NewPackagesItemWithPackage_nameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPackagesWithPackage_typeItemRequestBuilderInternal instantiates a new PackagesWithPackage_typeItemRequestBuilder and sets the default values.
func NewPackagesWithPackage_typeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PackagesWithPackage_typeItemRequestBuilder) {
    m := &PackagesWithPackage_typeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/packages/{package_type}", pathParameters),
    }
    return m
}
// NewPackagesWithPackage_typeItemRequestBuilder instantiates a new PackagesWithPackage_typeItemRequestBuilder and sets the default values.
func NewPackagesWithPackage_typeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PackagesWithPackage_typeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPackagesWithPackage_typeItemRequestBuilderInternal(urlParams, requestAdapter)
}
