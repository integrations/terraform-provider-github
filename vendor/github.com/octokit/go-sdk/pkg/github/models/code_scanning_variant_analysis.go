package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningVariantAnalysis a run of a CodeQL query against one or more repositories.
type CodeScanningVariantAnalysis struct {
    // The GitHub Actions workflow run used to execute this variant analysis. This is only available if the workflow run has started.
    actions_workflow_run_id *int32
    // A GitHub user.
    actor SimpleUserable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The date and time at which the variant analysis was completed, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ. Will be null if the variant analysis has not yet completed or this information is not available.
    completed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub repository.
    controller_repo SimpleRepositoryable
    // The date and time at which the variant analysis was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The reason for a failure of the variant analysis. This is only available if the variant analysis has failed.
    failure_reason *CodeScanningVariantAnalysis_failure_reason
    // The ID of the variant analysis.
    id *int32
    // The language targeted by the CodeQL query
    query_language *CodeScanningVariantAnalysisLanguage
    // The download url for the query pack.
    query_pack_url *string
    // The scanned_repositories property
    scanned_repositories []CodeScanningVariantAnalysis_scanned_repositoriesable
    // Information about repositories that were skipped from processing. This information is only available to the user that initiated the variant analysis.
    skipped_repositories CodeScanningVariantAnalysis_skipped_repositoriesable
    // The status property
    status *CodeScanningVariantAnalysisStatus
    // The date and time at which the variant analysis was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewCodeScanningVariantAnalysis instantiates a new CodeScanningVariantAnalysis and sets the default values.
func NewCodeScanningVariantAnalysis()(*CodeScanningVariantAnalysis) {
    m := &CodeScanningVariantAnalysis{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningVariantAnalysisFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningVariantAnalysisFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningVariantAnalysis(), nil
}
// GetActionsWorkflowRunId gets the actions_workflow_run_id property value. The GitHub Actions workflow run used to execute this variant analysis. This is only available if the workflow run has started.
// returns a *int32 when successful
func (m *CodeScanningVariantAnalysis) GetActionsWorkflowRunId()(*int32) {
    return m.actions_workflow_run_id
}
// GetActor gets the actor property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *CodeScanningVariantAnalysis) GetActor()(SimpleUserable) {
    return m.actor
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningVariantAnalysis) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCompletedAt gets the completed_at property value. The date and time at which the variant analysis was completed, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ. Will be null if the variant analysis has not yet completed or this information is not available.
// returns a *Time when successful
func (m *CodeScanningVariantAnalysis) GetCompletedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completed_at
}
// GetControllerRepo gets the controller_repo property value. A GitHub repository.
// returns a SimpleRepositoryable when successful
func (m *CodeScanningVariantAnalysis) GetControllerRepo()(SimpleRepositoryable) {
    return m.controller_repo
}
// GetCreatedAt gets the created_at property value. The date and time at which the variant analysis was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
// returns a *Time when successful
func (m *CodeScanningVariantAnalysis) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetFailureReason gets the failure_reason property value. The reason for a failure of the variant analysis. This is only available if the variant analysis has failed.
// returns a *CodeScanningVariantAnalysis_failure_reason when successful
func (m *CodeScanningVariantAnalysis) GetFailureReason()(*CodeScanningVariantAnalysis_failure_reason) {
    return m.failure_reason
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningVariantAnalysis) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actions_workflow_run_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionsWorkflowRunId(val)
        }
        return nil
    }
    res["actor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val.(SimpleUserable))
        }
        return nil
    }
    res["completed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedAt(val)
        }
        return nil
    }
    res["controller_repo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControllerRepo(val.(SimpleRepositoryable))
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["failure_reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningVariantAnalysis_failure_reason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureReason(val.(*CodeScanningVariantAnalysis_failure_reason))
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["query_language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningVariantAnalysisLanguage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryLanguage(val.(*CodeScanningVariantAnalysisLanguage))
        }
        return nil
    }
    res["query_pack_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryPackUrl(val)
        }
        return nil
    }
    res["scanned_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCodeScanningVariantAnalysis_scanned_repositoriesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CodeScanningVariantAnalysis_scanned_repositoriesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CodeScanningVariantAnalysis_scanned_repositoriesable)
                }
            }
            m.SetScannedRepositories(res)
        }
        return nil
    }
    res["skipped_repositories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeScanningVariantAnalysis_skipped_repositoriesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkippedRepositories(val.(CodeScanningVariantAnalysis_skipped_repositoriesable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningVariantAnalysisStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CodeScanningVariantAnalysisStatus))
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The ID of the variant analysis.
// returns a *int32 when successful
func (m *CodeScanningVariantAnalysis) GetId()(*int32) {
    return m.id
}
// GetQueryLanguage gets the query_language property value. The language targeted by the CodeQL query
// returns a *CodeScanningVariantAnalysisLanguage when successful
func (m *CodeScanningVariantAnalysis) GetQueryLanguage()(*CodeScanningVariantAnalysisLanguage) {
    return m.query_language
}
// GetQueryPackUrl gets the query_pack_url property value. The download url for the query pack.
// returns a *string when successful
func (m *CodeScanningVariantAnalysis) GetQueryPackUrl()(*string) {
    return m.query_pack_url
}
// GetScannedRepositories gets the scanned_repositories property value. The scanned_repositories property
// returns a []CodeScanningVariantAnalysis_scanned_repositoriesable when successful
func (m *CodeScanningVariantAnalysis) GetScannedRepositories()([]CodeScanningVariantAnalysis_scanned_repositoriesable) {
    return m.scanned_repositories
}
// GetSkippedRepositories gets the skipped_repositories property value. Information about repositories that were skipped from processing. This information is only available to the user that initiated the variant analysis.
// returns a CodeScanningVariantAnalysis_skipped_repositoriesable when successful
func (m *CodeScanningVariantAnalysis) GetSkippedRepositories()(CodeScanningVariantAnalysis_skipped_repositoriesable) {
    return m.skipped_repositories
}
// GetStatus gets the status property value. The status property
// returns a *CodeScanningVariantAnalysisStatus when successful
func (m *CodeScanningVariantAnalysis) GetStatus()(*CodeScanningVariantAnalysisStatus) {
    return m.status
}
// GetUpdatedAt gets the updated_at property value. The date and time at which the variant analysis was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
// returns a *Time when successful
func (m *CodeScanningVariantAnalysis) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *CodeScanningVariantAnalysis) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("actions_workflow_run_id", m.GetActionsWorkflowRunId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("actor", m.GetActor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("completed_at", m.GetCompletedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("controller_repo", m.GetControllerRepo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    if m.GetFailureReason() != nil {
        cast := (*m.GetFailureReason()).String()
        err := writer.WriteStringValue("failure_reason", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetQueryLanguage() != nil {
        cast := (*m.GetQueryLanguage()).String()
        err := writer.WriteStringValue("query_language", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("query_pack_url", m.GetQueryPackUrl())
        if err != nil {
            return err
        }
    }
    if m.GetScannedRepositories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetScannedRepositories()))
        for i, v := range m.GetScannedRepositories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("scanned_repositories", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("skipped_repositories", m.GetSkippedRepositories())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
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
// SetActionsWorkflowRunId sets the actions_workflow_run_id property value. The GitHub Actions workflow run used to execute this variant analysis. This is only available if the workflow run has started.
func (m *CodeScanningVariantAnalysis) SetActionsWorkflowRunId(value *int32)() {
    m.actions_workflow_run_id = value
}
// SetActor sets the actor property value. A GitHub user.
func (m *CodeScanningVariantAnalysis) SetActor(value SimpleUserable)() {
    m.actor = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodeScanningVariantAnalysis) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCompletedAt sets the completed_at property value. The date and time at which the variant analysis was completed, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ. Will be null if the variant analysis has not yet completed or this information is not available.
func (m *CodeScanningVariantAnalysis) SetCompletedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completed_at = value
}
// SetControllerRepo sets the controller_repo property value. A GitHub repository.
func (m *CodeScanningVariantAnalysis) SetControllerRepo(value SimpleRepositoryable)() {
    m.controller_repo = value
}
// SetCreatedAt sets the created_at property value. The date and time at which the variant analysis was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
func (m *CodeScanningVariantAnalysis) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetFailureReason sets the failure_reason property value. The reason for a failure of the variant analysis. This is only available if the variant analysis has failed.
func (m *CodeScanningVariantAnalysis) SetFailureReason(value *CodeScanningVariantAnalysis_failure_reason)() {
    m.failure_reason = value
}
// SetId sets the id property value. The ID of the variant analysis.
func (m *CodeScanningVariantAnalysis) SetId(value *int32)() {
    m.id = value
}
// SetQueryLanguage sets the query_language property value. The language targeted by the CodeQL query
func (m *CodeScanningVariantAnalysis) SetQueryLanguage(value *CodeScanningVariantAnalysisLanguage)() {
    m.query_language = value
}
// SetQueryPackUrl sets the query_pack_url property value. The download url for the query pack.
func (m *CodeScanningVariantAnalysis) SetQueryPackUrl(value *string)() {
    m.query_pack_url = value
}
// SetScannedRepositories sets the scanned_repositories property value. The scanned_repositories property
func (m *CodeScanningVariantAnalysis) SetScannedRepositories(value []CodeScanningVariantAnalysis_scanned_repositoriesable)() {
    m.scanned_repositories = value
}
// SetSkippedRepositories sets the skipped_repositories property value. Information about repositories that were skipped from processing. This information is only available to the user that initiated the variant analysis.
func (m *CodeScanningVariantAnalysis) SetSkippedRepositories(value CodeScanningVariantAnalysis_skipped_repositoriesable)() {
    m.skipped_repositories = value
}
// SetStatus sets the status property value. The status property
func (m *CodeScanningVariantAnalysis) SetStatus(value *CodeScanningVariantAnalysisStatus)() {
    m.status = value
}
// SetUpdatedAt sets the updated_at property value. The date and time at which the variant analysis was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
func (m *CodeScanningVariantAnalysis) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
type CodeScanningVariantAnalysisable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionsWorkflowRunId()(*int32)
    GetActor()(SimpleUserable)
    GetCompletedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetControllerRepo()(SimpleRepositoryable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFailureReason()(*CodeScanningVariantAnalysis_failure_reason)
    GetId()(*int32)
    GetQueryLanguage()(*CodeScanningVariantAnalysisLanguage)
    GetQueryPackUrl()(*string)
    GetScannedRepositories()([]CodeScanningVariantAnalysis_scanned_repositoriesable)
    GetSkippedRepositories()(CodeScanningVariantAnalysis_skipped_repositoriesable)
    GetStatus()(*CodeScanningVariantAnalysisStatus)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetActionsWorkflowRunId(value *int32)()
    SetActor(value SimpleUserable)()
    SetCompletedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetControllerRepo(value SimpleRepositoryable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFailureReason(value *CodeScanningVariantAnalysis_failure_reason)()
    SetId(value *int32)()
    SetQueryLanguage(value *CodeScanningVariantAnalysisLanguage)()
    SetQueryPackUrl(value *string)()
    SetScannedRepositories(value []CodeScanningVariantAnalysis_scanned_repositoriesable)()
    SetSkippedRepositories(value CodeScanningVariantAnalysis_skipped_repositoriesable)()
    SetStatus(value *CodeScanningVariantAnalysisStatus)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
