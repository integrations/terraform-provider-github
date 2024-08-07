package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeScanningVariantAnalysisSkippedRepoGroup struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A list of repositories that were skipped. This list may not include all repositories that were skipped. This is only available when the repository was found and the user has access to it.
    repositories []CodeScanningVariantAnalysisRepositoryable
    // The total number of repositories that were skipped for this reason.
    repository_count *int32
}
// NewCodeScanningVariantAnalysisSkippedRepoGroup instantiates a new CodeScanningVariantAnalysisSkippedRepoGroup and sets the default values.
func NewCodeScanningVariantAnalysisSkippedRepoGroup()(*CodeScanningVariantAnalysisSkippedRepoGroup) {
    m := &CodeScanningVariantAnalysisSkippedRepoGroup{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningVariantAnalysisSkippedRepoGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningVariantAnalysisSkippedRepoGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningVariantAnalysisSkippedRepoGroup(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningVariantAnalysisSkippedRepoGroup) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningVariantAnalysisSkippedRepoGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCodeScanningVariantAnalysisRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CodeScanningVariantAnalysisRepositoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CodeScanningVariantAnalysisRepositoryable)
                }
            }
            m.SetRepositories(res)
        }
        return nil
    }
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
    return res
}
// GetRepositories gets the repositories property value. A list of repositories that were skipped. This list may not include all repositories that were skipped. This is only available when the repository was found and the user has access to it.
// returns a []CodeScanningVariantAnalysisRepositoryable when successful
func (m *CodeScanningVariantAnalysisSkippedRepoGroup) GetRepositories()([]CodeScanningVariantAnalysisRepositoryable) {
    return m.repositories
}
// GetRepositoryCount gets the repository_count property value. The total number of repositories that were skipped for this reason.
// returns a *int32 when successful
func (m *CodeScanningVariantAnalysisSkippedRepoGroup) GetRepositoryCount()(*int32) {
    return m.repository_count
}
// Serialize serializes information the current object
func (m *CodeScanningVariantAnalysisSkippedRepoGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRepositories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRepositories()))
        for i, v := range m.GetRepositories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("repositories", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("repository_count", m.GetRepositoryCount())
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
func (m *CodeScanningVariantAnalysisSkippedRepoGroup) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRepositories sets the repositories property value. A list of repositories that were skipped. This list may not include all repositories that were skipped. This is only available when the repository was found and the user has access to it.
func (m *CodeScanningVariantAnalysisSkippedRepoGroup) SetRepositories(value []CodeScanningVariantAnalysisRepositoryable)() {
    m.repositories = value
}
// SetRepositoryCount sets the repository_count property value. The total number of repositories that were skipped for this reason.
func (m *CodeScanningVariantAnalysisSkippedRepoGroup) SetRepositoryCount(value *int32)() {
    m.repository_count = value
}
type CodeScanningVariantAnalysisSkippedRepoGroupable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRepositories()([]CodeScanningVariantAnalysisRepositoryable)
    GetRepositoryCount()(*int32)
    SetRepositories(value []CodeScanningVariantAnalysisRepositoryable)()
    SetRepositoryCount(value *int32)()
}
