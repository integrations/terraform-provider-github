package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeScanningVariantAnalysisRepoTask struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The new status of the CodeQL variant analysis repository task.
    analysis_status *CodeScanningVariantAnalysisStatus
    // The size of the artifact. This is only available for successful analyses.
    artifact_size_in_bytes *int32
    // The URL of the artifact. This is only available for successful analyses.
    artifact_url *string
    // The SHA of the commit the CodeQL database was built against. This is only available for successful analyses.
    database_commit_sha *string
    // The reason of the failure of this repo task. This is only available if the repository task has failed.
    failure_message *string
    // A GitHub repository.
    repository SimpleRepositoryable
    // The number of results in the case of a successful analysis. This is only available for successful analyses.
    result_count *int32
    // The source location prefix to use. This is only available for successful analyses.
    source_location_prefix *string
}
// NewCodeScanningVariantAnalysisRepoTask instantiates a new CodeScanningVariantAnalysisRepoTask and sets the default values.
func NewCodeScanningVariantAnalysisRepoTask()(*CodeScanningVariantAnalysisRepoTask) {
    m := &CodeScanningVariantAnalysisRepoTask{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningVariantAnalysisRepoTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningVariantAnalysisRepoTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningVariantAnalysisRepoTask(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningVariantAnalysisRepoTask) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAnalysisStatus gets the analysis_status property value. The new status of the CodeQL variant analysis repository task.
// returns a *CodeScanningVariantAnalysisStatus when successful
func (m *CodeScanningVariantAnalysisRepoTask) GetAnalysisStatus()(*CodeScanningVariantAnalysisStatus) {
    return m.analysis_status
}
// GetArtifactSizeInBytes gets the artifact_size_in_bytes property value. The size of the artifact. This is only available for successful analyses.
// returns a *int32 when successful
func (m *CodeScanningVariantAnalysisRepoTask) GetArtifactSizeInBytes()(*int32) {
    return m.artifact_size_in_bytes
}
// GetArtifactUrl gets the artifact_url property value. The URL of the artifact. This is only available for successful analyses.
// returns a *string when successful
func (m *CodeScanningVariantAnalysisRepoTask) GetArtifactUrl()(*string) {
    return m.artifact_url
}
// GetDatabaseCommitSha gets the database_commit_sha property value. The SHA of the commit the CodeQL database was built against. This is only available for successful analyses.
// returns a *string when successful
func (m *CodeScanningVariantAnalysisRepoTask) GetDatabaseCommitSha()(*string) {
    return m.database_commit_sha
}
// GetFailureMessage gets the failure_message property value. The reason of the failure of this repo task. This is only available if the repository task has failed.
// returns a *string when successful
func (m *CodeScanningVariantAnalysisRepoTask) GetFailureMessage()(*string) {
    return m.failure_message
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningVariantAnalysisRepoTask) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["artifact_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArtifactUrl(val)
        }
        return nil
    }
    res["database_commit_sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDatabaseCommitSha(val)
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
        val, err := n.GetObjectValue(CreateSimpleRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(SimpleRepositoryable))
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
    res["source_location_prefix"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceLocationPrefix(val)
        }
        return nil
    }
    return res
}
// GetRepository gets the repository property value. A GitHub repository.
// returns a SimpleRepositoryable when successful
func (m *CodeScanningVariantAnalysisRepoTask) GetRepository()(SimpleRepositoryable) {
    return m.repository
}
// GetResultCount gets the result_count property value. The number of results in the case of a successful analysis. This is only available for successful analyses.
// returns a *int32 when successful
func (m *CodeScanningVariantAnalysisRepoTask) GetResultCount()(*int32) {
    return m.result_count
}
// GetSourceLocationPrefix gets the source_location_prefix property value. The source location prefix to use. This is only available for successful analyses.
// returns a *string when successful
func (m *CodeScanningVariantAnalysisRepoTask) GetSourceLocationPrefix()(*string) {
    return m.source_location_prefix
}
// Serialize serializes information the current object
func (m *CodeScanningVariantAnalysisRepoTask) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("artifact_url", m.GetArtifactUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("database_commit_sha", m.GetDatabaseCommitSha())
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
        err := writer.WriteStringValue("source_location_prefix", m.GetSourceLocationPrefix())
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
func (m *CodeScanningVariantAnalysisRepoTask) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAnalysisStatus sets the analysis_status property value. The new status of the CodeQL variant analysis repository task.
func (m *CodeScanningVariantAnalysisRepoTask) SetAnalysisStatus(value *CodeScanningVariantAnalysisStatus)() {
    m.analysis_status = value
}
// SetArtifactSizeInBytes sets the artifact_size_in_bytes property value. The size of the artifact. This is only available for successful analyses.
func (m *CodeScanningVariantAnalysisRepoTask) SetArtifactSizeInBytes(value *int32)() {
    m.artifact_size_in_bytes = value
}
// SetArtifactUrl sets the artifact_url property value. The URL of the artifact. This is only available for successful analyses.
func (m *CodeScanningVariantAnalysisRepoTask) SetArtifactUrl(value *string)() {
    m.artifact_url = value
}
// SetDatabaseCommitSha sets the database_commit_sha property value. The SHA of the commit the CodeQL database was built against. This is only available for successful analyses.
func (m *CodeScanningVariantAnalysisRepoTask) SetDatabaseCommitSha(value *string)() {
    m.database_commit_sha = value
}
// SetFailureMessage sets the failure_message property value. The reason of the failure of this repo task. This is only available if the repository task has failed.
func (m *CodeScanningVariantAnalysisRepoTask) SetFailureMessage(value *string)() {
    m.failure_message = value
}
// SetRepository sets the repository property value. A GitHub repository.
func (m *CodeScanningVariantAnalysisRepoTask) SetRepository(value SimpleRepositoryable)() {
    m.repository = value
}
// SetResultCount sets the result_count property value. The number of results in the case of a successful analysis. This is only available for successful analyses.
func (m *CodeScanningVariantAnalysisRepoTask) SetResultCount(value *int32)() {
    m.result_count = value
}
// SetSourceLocationPrefix sets the source_location_prefix property value. The source location prefix to use. This is only available for successful analyses.
func (m *CodeScanningVariantAnalysisRepoTask) SetSourceLocationPrefix(value *string)() {
    m.source_location_prefix = value
}
type CodeScanningVariantAnalysisRepoTaskable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalysisStatus()(*CodeScanningVariantAnalysisStatus)
    GetArtifactSizeInBytes()(*int32)
    GetArtifactUrl()(*string)
    GetDatabaseCommitSha()(*string)
    GetFailureMessage()(*string)
    GetRepository()(SimpleRepositoryable)
    GetResultCount()(*int32)
    GetSourceLocationPrefix()(*string)
    SetAnalysisStatus(value *CodeScanningVariantAnalysisStatus)()
    SetArtifactSizeInBytes(value *int32)()
    SetArtifactUrl(value *string)()
    SetDatabaseCommitSha(value *string)()
    SetFailureMessage(value *string)()
    SetRepository(value SimpleRepositoryable)()
    SetResultCount(value *int32)()
    SetSourceLocationPrefix(value *string)()
}
