package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeScanningVariantAnalysis_skipped_repositories_not_found_repos struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The total number of repositories that were skipped for this reason.
    repository_count *int32
    // A list of full repository names that were skipped. This list may not include all repositories that were skipped.
    repository_full_names []string
}
// NewCodeScanningVariantAnalysis_skipped_repositories_not_found_repos instantiates a new CodeScanningVariantAnalysis_skipped_repositories_not_found_repos and sets the default values.
func NewCodeScanningVariantAnalysis_skipped_repositories_not_found_repos()(*CodeScanningVariantAnalysis_skipped_repositories_not_found_repos) {
    m := &CodeScanningVariantAnalysis_skipped_repositories_not_found_repos{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningVariantAnalysis_skipped_repositories_not_found_reposFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningVariantAnalysis_skipped_repositories_not_found_reposFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningVariantAnalysis_skipped_repositories_not_found_repos(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningVariantAnalysis_skipped_repositories_not_found_repos) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningVariantAnalysis_skipped_repositories_not_found_repos) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["repository_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryCount(val)
        }
        return nil
    }
    res["repository_full_names"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRepositoryFullNames(res)
        }
        return nil
    }
    return res
}
// GetRepositoryCount gets the repository_count property value. The total number of repositories that were skipped for this reason.
// returns a *int32 when successful
func (m *CodeScanningVariantAnalysis_skipped_repositories_not_found_repos) GetRepositoryCount()(*int32) {
    return m.repository_count
}
// GetRepositoryFullNames gets the repository_full_names property value. A list of full repository names that were skipped. This list may not include all repositories that were skipped.
// returns a []string when successful
func (m *CodeScanningVariantAnalysis_skipped_repositories_not_found_repos) GetRepositoryFullNames()([]string) {
    return m.repository_full_names
}
// Serialize serializes information the current object
func (m *CodeScanningVariantAnalysis_skipped_repositories_not_found_repos) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("repository_count", m.GetRepositoryCount())
        if err != nil {
            return err
        }
    }
    if m.GetRepositoryFullNames() != nil {
        err := writer.WriteCollectionOfStringValues("repository_full_names", m.GetRepositoryFullNames())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodeScanningVariantAnalysis_skipped_repositories_not_found_repos) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRepositoryCount sets the repository_count property value. The total number of repositories that were skipped for this reason.
func (m *CodeScanningVariantAnalysis_skipped_repositories_not_found_repos) SetRepositoryCount(value *int32)() {
    m.repository_count = value
}
// SetRepositoryFullNames sets the repository_full_names property value. A list of full repository names that were skipped. This list may not include all repositories that were skipped.
func (m *CodeScanningVariantAnalysis_skipped_repositories_not_found_repos) SetRepositoryFullNames(value []string)() {
    m.repository_full_names = value
}
type CodeScanningVariantAnalysis_skipped_repositories_not_found_reposable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRepositoryCount()(*int32)
    GetRepositoryFullNames()([]string)
    SetRepositoryCount(value *int32)()
    SetRepositoryFullNames(value []string)()
}
