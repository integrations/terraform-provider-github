package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeScanningVariantAnalysis_scanned_repositories struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The new status of the CodeQL variant analysis repository task.
    analysis_status *CodeScanningVariantAnalysisStatus
    // The size of the artifact. This is only available for successful analyses.
    artifact_size_in_bytes *int32
    // The reason of the failure of this repo task. This is only available if the repository task has failed.
    failure_message *string
    // Repository Identifier
    repository CodeScanningVariantAnalysisRepositoryable
    // The number of results in the case of a successful analysis. This is only available for successful analyses.
    result_count *int32
}
// NewCodeScanningVariantAnalysis_scanned_repositories instantiates a new CodeScanningVariantAnalysis_scanned_repositories and sets the default values.
func NewCodeScanningVariantAnalysis_scanned_repositories()(*CodeScanningVariantAnalysis_scanned_repositories) {
    m := &CodeScanningVariantAnalysis_scanned_repositories{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningVariantAnalysis_scanned_repositoriesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningVariantAnalysis_scanned_repositoriesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningVariantAnalysis_scanned_repositories(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningVariantAnalysis_scanned_repositories) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAnalysisStatus gets the analysis_status property value. The new status of the CodeQL variant analysis repository task.
// returns a *CodeScanningVariantAnalysisStatus when successful
func (m *CodeScanningVariantAnalysis_scanned_repositories) GetAnalysisStatus()(*CodeScanningVariantAnalysisStatus) {
    return m.analysis_status
}
// GetArtifactSizeInBytes gets the artifact_size_in_bytes property value. The size of the artifact. This is only available for successful analyses.
// returns a *int32 when successful
func (m *CodeScanningVariantAnalysis_scanned_repositories) GetArtifactSizeInBytes()(*int32) {
    return m.artifact_size_in_bytes
}
// GetFailureMessage gets the failure_message property value. The reason of the failure of this repo task. This is only available if the repository task has failed.
// returns a *string when successful
func (m *CodeScanningVariantAnalysis_scanned_repositories) GetFailureMessage()(*string) {
    return m.failure_message
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningVariantAnalysis_scanned_repositories) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["analysis_status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningVariantAnalysisStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnalysisStatus(val.(*CodeScanningVariantAnalysisStatus))
        }
        return nil
    }
    res["artifact_size_in_bytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArtifactSizeInBytes(val)
        }
        return nil
    }
    res["failure_message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureMessage(val)
        }
        return nil
    }
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeScanningVariantAnalysisRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(CodeScanningVariantAnalysisRepositoryable))
        }
        return nil
    }
    res["result_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultCount(val)
        }
        return nil
    }
    return res
}
// GetRepository gets the repository property value. Repository Identifier
// returns a CodeScanningVariantAnalysisRepositoryable when successful
func (m *CodeScanningVariantAnalysis_scanned_repositories) GetRepository()(CodeScanningVariantAnalysisRepositoryable) {
    return m.repository
}
// GetResultCount gets the result_count property value. The number of results in the case of a successful analysis. This is only available for successful analyses.
// returns a *int32 when successful
func (m *CodeScanningVariantAnalysis_scanned_repositories) GetResultCount()(*int32) {
    return m.result_count
}
// Serialize serializes information the current object
func (m *CodeScanningVariantAnalysis_scanned_repositories) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAnalysisStatus() != nil {
        cast := (*m.GetAnalysisStatus()).String()
        err := writer.WriteStringValue("analysis_status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("artifact_size_in_bytes", m.GetArtifactSizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("failure_message", m.GetFailureMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repository", m.GetRepository())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("result_count", m.GetResultCount())
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
func (m *CodeScanningVariantAnalysis_scanned_repositories) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAnalysisStatus sets the analysis_status property value. The new status of the CodeQL variant analysis repository task.
func (m *CodeScanningVariantAnalysis_scanned_repositories) SetAnalysisStatus(value *CodeScanningVariantAnalysisStatus)() {
    m.analysis_status = value
}
// SetArtifactSizeInBytes sets the artifact_size_in_bytes property value. The size of the artifact. This is only available for successful analyses.
func (m *CodeScanningVariantAnalysis_scanned_repositories) SetArtifactSizeInBytes(value *int32)() {
    m.artifact_size_in_bytes = value
}
// SetFailureMessage sets the failure_message property value. The reason of the failure of this repo task. This is only available if the repository task has failed.
func (m *CodeScanningVariantAnalysis_scanned_repositories) SetFailureMessage(value *string)() {
    m.failure_message = value
}
// SetRepository sets the repository property value. Repository Identifier
func (m *CodeScanningVariantAnalysis_scanned_repositories) SetRepository(value CodeScanningVariantAnalysisRepositoryable)() {
    m.repository = value
}
// SetResultCount sets the result_count property value. The number of results in the case of a successful analysis. This is only available for successful analyses.
func (m *CodeScanningVariantAnalysis_scanned_repositories) SetResultCount(value *int32)() {
    m.result_count = value
}
type CodeScanningVariantAnalysis_scanned_repositoriesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalysisStatus()(*CodeScanningVariantAnalysisStatus)
    GetArtifactSizeInBytes()(*int32)
    GetFailureMessage()(*string)
    GetRepository()(CodeScanningVariantAnalysisRepositoryable)
    GetResultCount()(*int32)
    SetAnalysisStatus(value *CodeScanningVariantAnalysisStatus)()
    SetArtifactSizeInBytes(value *int32)()
    SetFailureMessage(value *string)()
    SetRepository(value CodeScanningVariantAnalysisRepositoryable)()
    SetResultCount(value *int32)()
}
