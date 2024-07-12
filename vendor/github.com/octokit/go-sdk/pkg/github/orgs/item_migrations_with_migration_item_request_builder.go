package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    i399c3da064b81c83d30565b554a09039c3bc9dc0590affc1306cb54f75cd8e0d "github.com/octokit/go-sdk/pkg/github/orgs/item/migrations/item"
)

// ItemMigrationsWithMigration_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\migrations\{migration_id}
type ItemMigrationsWithMigration_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMigrationsWithMigration_ItemRequestBuilderGetQueryParameters fetches the status of a migration.The `state` of a migration can be one of the following values:*   `pending`, which means the migration hasn't started yet.*   `exporting`, which means the migration is in progress.*   `exported`, which means the migration finished successfully.*   `failed`, which means the migration failed.
type ItemMigrationsWithMigration_ItemRequestBuilderGetQueryParameters struct {
    // Exclude attributes from the API response to improve performance
    Exclude []i399c3da064b81c83d30565b554a09039c3bc9dc0590affc1306cb54f75cd8e0d.GetExcludeQueryParameterType `uriparametername:"exclude"`
}
// Archive the archive property
// returns a *ItemMigrationsItemArchiveRequestBuilder when successful
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Archive()(*ItemMigrationsItemArchiveRequestBuilder) {
    return NewItemMigrationsItemArchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMigrationsWithMigration_ItemRequestBuilderInternal instantiates a new ItemMigrationsWithMigration_ItemRequestBuilder and sets the default values.
func NewItemMigrationsWithMigration_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsWithMigration_ItemRequestBuilder) {
    m := &ItemMigrationsWithMigration_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/migrations/{migration_id}{?exclude*}", pathParameters),
    }
    return m
}
// NewItemMigrationsWithMigration_ItemRequestBuilder instantiates a new ItemMigrationsWithMigration_ItemRequestBuilder and sets the default values.
func NewItemMigrationsWithMigration_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMigrationsWithMigration_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMigrationsWithMigration_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get fetches the status of a migration.The `state` of a migration can be one of the following values:*   `pending`, which means the migration hasn't started yet.*   `exporting`, which means the migration is in progress.*   `exported`, which means the migration finished successfully.*   `failed`, which means the migration failed.
// returns a Migrationable when successful
// returns a BasicError error when the service returns a 404 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/migrations/orgs#get-an-organization-migration-status
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemMigrationsWithMigration_ItemRequestBuilderGetQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Migrationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateMigrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Migrationable), nil
}
// Repos the repos property
// returns a *ItemMigrationsItemReposRequestBuilder when successful
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Repos()(*ItemMigrationsItemReposRequestBuilder) {
    return NewItemMigrationsItemReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repositories the repositories property
// returns a *ItemMigrationsItemRepositoriesRequestBuilder when successful
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) Repositories()(*ItemMigrationsItemRepositoriesRequestBuilder) {
    return NewItemMigrationsItemRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation fetches the status of a migration.The `state` of a migration can be one of the following values:*   `pending`, which means the migration hasn't started yet.*   `exporting`, which means the migration is in progress.*   `exported`, which means the migration finished successfully.*   `failed`, which means the migration failed.
// returns a *RequestInformation when successful
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemMigrationsWithMigration_ItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemMigrationsWithMigration_ItemRequestBuilder when successful
func (m *ItemMigrationsWithMigration_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemMigrationsWithMigration_ItemRequestBuilder) {
    return NewItemMigrationsWithMigration_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
