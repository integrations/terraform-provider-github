package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody struct {
    // The language targeted by the CodeQL query
    language *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningVariantAnalysisLanguage
    // A Base64-encoded tarball containing a CodeQL query and all its dependencies
    query_pack *string
    // List of repository names (in the form `owner/repo-name`) to run the query against. Precisely one property from `repositories`, `repository_lists` and `repository_owners` is required.
    repositories []string
    // List of repository lists to run the query against. Precisely one property from `repositories`, `repository_lists` and `repository_owners` is required.
    repository_lists []string
    // List of organization or user names whose repositories the query should be run against. Precisely one property from `repositories`, `repository_lists` and `repository_owners` is required.
    repository_owners []string
}
// NewItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody instantiates a new ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody and sets the default values.
func NewItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody()(*ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody) {
    m := &ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody{
    }
    return m
}
// CreateItemItemCodeScanningCodeqlVariantAnalysesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemCodeScanningCodeqlVariantAnalysesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.ParseCodeScanningVariantAnalysisLanguage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val.(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningVariantAnalysisLanguage))
        }
        return nil
    }
    res["query_pack"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryPack(val)
        }
        return nil
    }
    res["repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRepositories(res)
        }
        return nil
    }
    res["repository_lists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRepositoryLists(res)
        }
        return nil
    }
    res["repository_owners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRepositoryOwners(res)
        }
        return nil
    }
    return res
}
// GetLanguage gets the language property value. The language targeted by the CodeQL query
// returns a *CodeScanningVariantAnalysisLanguage when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody) GetLanguage()(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningVariantAnalysisLanguage) {
    return m.language
}
// GetQueryPack gets the query_pack property value. A Base64-encoded tarball containing a CodeQL query and all its dependencies
// returns a *string when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody) GetQueryPack()(*string) {
    return m.query_pack
}
// GetRepositories gets the repositories property value. List of repository names (in the form `owner/repo-name`) to run the query against. Precisely one property from `repositories`, `repository_lists` and `repository_owners` is required.
// returns a []string when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody) GetRepositories()([]string) {
    return m.repositories
}
// GetRepositoryLists gets the repository_lists property value. List of repository lists to run the query against. Precisely one property from `repositories`, `repository_lists` and `repository_owners` is required.
// returns a []string when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody) GetRepositoryLists()([]string) {
    return m.repository_lists
}
// GetRepositoryOwners gets the repository_owners property value. List of organization or user names whose repositories the query should be run against. Precisely one property from `repositories`, `repository_lists` and `repository_owners` is required.
// returns a []string when successful
func (m *ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody) GetRepositoryOwners()([]string) {
    return m.repository_owners
}
// Serialize serializes information the current object
func (m *ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLanguage() != nil {
        cast := (*m.GetLanguage()).String()
        err := writer.WriteStringValue("language", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("query_pack", m.GetQueryPack())
        if err != nil {
            return err
        }
    }
    if m.GetRepositories() != nil {
        err := writer.WriteCollectionOfStringValues("repositories", m.GetRepositories())
        if err != nil {
            return err
        }
    }
    if m.GetRepositoryLists() != nil {
        err := writer.WriteCollectionOfStringValues("repository_lists", m.GetRepositoryLists())
        if err != nil {
            return err
        }
    }
    if m.GetRepositoryOwners() != nil {
        err := writer.WriteCollectionOfStringValues("repository_owners", m.GetRepositoryOwners())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLanguage sets the language property value. The language targeted by the CodeQL query
func (m *ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody) SetLanguage(value *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningVariantAnalysisLanguage)() {
    m.language = value
}
// SetQueryPack sets the query_pack property value. A Base64-encoded tarball containing a CodeQL query and all its dependencies
func (m *ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody) SetQueryPack(value *string)() {
    m.query_pack = value
}
// SetRepositories sets the repositories property value. List of repository names (in the form `owner/repo-name`) to run the query against. Precisely one property from `repositories`, `repository_lists` and `repository_owners` is required.
func (m *ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody) SetRepositories(value []string)() {
    m.repositories = value
}
// SetRepositoryLists sets the repository_lists property value. List of repository lists to run the query against. Precisely one property from `repositories`, `repository_lists` and `repository_owners` is required.
func (m *ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody) SetRepositoryLists(value []string)() {
    m.repository_lists = value
}
// SetRepositoryOwners sets the repository_owners property value. List of organization or user names whose repositories the query should be run against. Precisely one property from `repositories`, `repository_lists` and `repository_owners` is required.
func (m *ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBody) SetRepositoryOwners(value []string)() {
    m.repository_owners = value
}
type ItemItemCodeScanningCodeqlVariantAnalysesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLanguage()(*i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningVariantAnalysisLanguage)
    GetQueryPack()(*string)
    GetRepositories()([]string)
    GetRepositoryLists()([]string)
    GetRepositoryOwners()([]string)
    SetLanguage(value *i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CodeScanningVariantAnalysisLanguage)()
    SetQueryPack(value *string)()
    SetRepositories(value []string)()
    SetRepositoryLists(value []string)()
    SetRepositoryOwners(value []string)()
}
