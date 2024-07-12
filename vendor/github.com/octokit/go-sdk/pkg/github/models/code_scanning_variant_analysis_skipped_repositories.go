package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningVariantAnalysis_skipped_repositories information about repositories that were skipped from processing. This information is only available to the user that initiated the variant analysis.
type CodeScanningVariantAnalysis_skipped_repositories struct {
    // The access_mismatch_repos property
    access_mismatch_repos CodeScanningVariantAnalysisSkippedRepoGroupable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The no_codeql_db_repos property
    no_codeql_db_repos CodeScanningVariantAnalysisSkippedRepoGroupable
    // The not_found_repos property
    not_found_repos CodeScanningVariantAnalysis_skipped_repositories_not_found_reposable
    // The over_limit_repos property
    over_limit_repos CodeScanningVariantAnalysisSkippedRepoGroupable
}
// NewCodeScanningVariantAnalysis_skipped_repositories instantiates a new CodeScanningVariantAnalysis_skipped_repositories and sets the default values.
func NewCodeScanningVariantAnalysis_skipped_repositories()(*CodeScanningVariantAnalysis_skipped_repositories) {
    m := &CodeScanningVariantAnalysis_skipped_repositories{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningVariantAnalysis_skipped_repositoriesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningVariantAnalysis_skipped_repositoriesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningVariantAnalysis_skipped_repositories(), nil
}
// GetAccessMismatchRepos gets the access_mismatch_repos property value. The access_mismatch_repos property
// returns a CodeScanningVariantAnalysisSkippedRepoGroupable when successful
func (m *CodeScanningVariantAnalysis_skipped_repositories) GetAccessMismatchRepos()(CodeScanningVariantAnalysisSkippedRepoGroupable) {
    return m.access_mismatch_repos
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningVariantAnalysis_skipped_repositories) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningVariantAnalysis_skipped_repositories) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access_mismatch_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeScanningVariantAnalysisSkippedRepoGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessMismatchRepos(val.(CodeScanningVariantAnalysisSkippedRepoGroupable))
        }
        return nil
    }
    res["no_codeql_db_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeScanningVariantAnalysisSkippedRepoGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNoCodeqlDbRepos(val.(CodeScanningVariantAnalysisSkippedRepoGroupable))
        }
        return nil
    }
    res["not_found_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeScanningVariantAnalysis_skipped_repositories_not_found_reposFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotFoundRepos(val.(CodeScanningVariantAnalysis_skipped_repositories_not_found_reposable))
        }
        return nil
    }
    res["over_limit_repos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeScanningVariantAnalysisSkippedRepoGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverLimitRepos(val.(CodeScanningVariantAnalysisSkippedRepoGroupable))
        }
        return nil
    }
    return res
}
// GetNoCodeqlDbRepos gets the no_codeql_db_repos property value. The no_codeql_db_repos property
// returns a CodeScanningVariantAnalysisSkippedRepoGroupable when successful
func (m *CodeScanningVariantAnalysis_skipped_repositories) GetNoCodeqlDbRepos()(CodeScanningVariantAnalysisSkippedRepoGroupable) {
    return m.no_codeql_db_repos
}
// GetNotFoundRepos gets the not_found_repos property value. The not_found_repos property
// returns a CodeScanningVariantAnalysis_skipped_repositories_not_found_reposable when successful
func (m *CodeScanningVariantAnalysis_skipped_repositories) GetNotFoundRepos()(CodeScanningVariantAnalysis_skipped_repositories_not_found_reposable) {
    return m.not_found_repos
}
// GetOverLimitRepos gets the over_limit_repos property value. The over_limit_repos property
// returns a CodeScanningVariantAnalysisSkippedRepoGroupable when successful
func (m *CodeScanningVariantAnalysis_skipped_repositories) GetOverLimitRepos()(CodeScanningVariantAnalysisSkippedRepoGroupable) {
    return m.over_limit_repos
}
// Serialize serializes information the current object
func (m *CodeScanningVariantAnalysis_skipped_repositories) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("access_mismatch_repos", m.GetAccessMismatchRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("not_found_repos", m.GetNotFoundRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("no_codeql_db_repos", m.GetNoCodeqlDbRepos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("over_limit_repos", m.GetOverLimitRepos())
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
// SetAccessMismatchRepos sets the access_mismatch_repos property value. The access_mismatch_repos property
func (m *CodeScanningVariantAnalysis_skipped_repositories) SetAccessMismatchRepos(value CodeScanningVariantAnalysisSkippedRepoGroupable)() {
    m.access_mismatch_repos = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodeScanningVariantAnalysis_skipped_repositories) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNoCodeqlDbRepos sets the no_codeql_db_repos property value. The no_codeql_db_repos property
func (m *CodeScanningVariantAnalysis_skipped_repositories) SetNoCodeqlDbRepos(value CodeScanningVariantAnalysisSkippedRepoGroupable)() {
    m.no_codeql_db_repos = value
}
// SetNotFoundRepos sets the not_found_repos property value. The not_found_repos property
func (m *CodeScanningVariantAnalysis_skipped_repositories) SetNotFoundRepos(value CodeScanningVariantAnalysis_skipped_repositories_not_found_reposable)() {
    m.not_found_repos = value
}
// SetOverLimitRepos sets the over_limit_repos property value. The over_limit_repos property
func (m *CodeScanningVariantAnalysis_skipped_repositories) SetOverLimitRepos(value CodeScanningVariantAnalysisSkippedRepoGroupable)() {
    m.over_limit_repos = value
}
type CodeScanningVariantAnalysis_skipped_repositoriesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessMismatchRepos()(CodeScanningVariantAnalysisSkippedRepoGroupable)
    GetNoCodeqlDbRepos()(CodeScanningVariantAnalysisSkippedRepoGroupable)
    GetNotFoundRepos()(CodeScanningVariantAnalysis_skipped_repositories_not_found_reposable)
    GetOverLimitRepos()(CodeScanningVariantAnalysisSkippedRepoGroupable)
    SetAccessMismatchRepos(value CodeScanningVariantAnalysisSkippedRepoGroupable)()
    SetNoCodeqlDbRepos(value CodeScanningVariantAnalysisSkippedRepoGroupable)()
    SetNotFoundRepos(value CodeScanningVariantAnalysis_skipped_repositories_not_found_reposable)()
    SetOverLimitRepos(value CodeScanningVariantAnalysisSkippedRepoGroupable)()
}
