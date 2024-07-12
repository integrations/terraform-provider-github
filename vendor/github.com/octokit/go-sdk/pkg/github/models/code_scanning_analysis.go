package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeScanningAnalysis struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
    analysis_key *string
    // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
    category *string
    // The SHA of the commit to which the analysis you are uploading relates.
    commit_sha *string
    // The time that the analysis was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The deletable property
    deletable *bool
    // Identifies the variable values associated with the environment in which this analysis was performed.
    environment *string
    // The error property
    error *string
    // Unique identifier for this analysis.
    id *int32
    // The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`,`refs/heads/<branch name>` or simply `<branch name>`.
    ref *string
    // The total number of results in the analysis.
    results_count *int32
    // The total number of rules used in the analysis.
    rules_count *int32
    // An identifier for the upload.
    sarif_id *string
    // The tool property
    tool CodeScanningAnalysisToolable
    // The REST API URL of the analysis resource.
    url *string
    // Warning generated when processing the analysis
    warning *string
}
// NewCodeScanningAnalysis instantiates a new CodeScanningAnalysis and sets the default values.
func NewCodeScanningAnalysis()(*CodeScanningAnalysis) {
    m := &CodeScanningAnalysis{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningAnalysisFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningAnalysisFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningAnalysis(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningAnalysis) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAnalysisKey gets the analysis_key property value. Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
// returns a *string when successful
func (m *CodeScanningAnalysis) GetAnalysisKey()(*string) {
    return m.analysis_key
}
// GetCategory gets the category property value. Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
// returns a *string when successful
func (m *CodeScanningAnalysis) GetCategory()(*string) {
    return m.category
}
// GetCommitSha gets the commit_sha property value. The SHA of the commit to which the analysis you are uploading relates.
// returns a *string when successful
func (m *CodeScanningAnalysis) GetCommitSha()(*string) {
    return m.commit_sha
}
// GetCreatedAt gets the created_at property value. The time that the analysis was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *CodeScanningAnalysis) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDeletable gets the deletable property value. The deletable property
// returns a *bool when successful
func (m *CodeScanningAnalysis) GetDeletable()(*bool) {
    return m.deletable
}
// GetEnvironment gets the environment property value. Identifies the variable values associated with the environment in which this analysis was performed.
// returns a *string when successful
func (m *CodeScanningAnalysis) GetEnvironment()(*string) {
    return m.environment
}
// GetError gets the error property value. The error property
// returns a *string when successful
func (m *CodeScanningAnalysis) GetError()(*string) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningAnalysis) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["analysis_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnalysisKey(val)
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val)
        }
        return nil
    }
    res["commit_sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitSha(val)
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
    res["deletable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletable(val)
        }
        return nil
    }
    res["environment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnvironment(val)
        }
        return nil
    }
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
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
    res["ref"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRef(val)
        }
        return nil
    }
    res["results_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultsCount(val)
        }
        return nil
    }
    res["rules_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRulesCount(val)
        }
        return nil
    }
    res["sarif_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSarifId(val)
        }
        return nil
    }
    res["tool"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeScanningAnalysisToolFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTool(val.(CodeScanningAnalysisToolable))
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    res["warning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWarning(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Unique identifier for this analysis.
// returns a *int32 when successful
func (m *CodeScanningAnalysis) GetId()(*int32) {
    return m.id
}
// GetRef gets the ref property value. The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`,`refs/heads/<branch name>` or simply `<branch name>`.
// returns a *string when successful
func (m *CodeScanningAnalysis) GetRef()(*string) {
    return m.ref
}
// GetResultsCount gets the results_count property value. The total number of results in the analysis.
// returns a *int32 when successful
func (m *CodeScanningAnalysis) GetResultsCount()(*int32) {
    return m.results_count
}
// GetRulesCount gets the rules_count property value. The total number of rules used in the analysis.
// returns a *int32 when successful
func (m *CodeScanningAnalysis) GetRulesCount()(*int32) {
    return m.rules_count
}
// GetSarifId gets the sarif_id property value. An identifier for the upload.
// returns a *string when successful
func (m *CodeScanningAnalysis) GetSarifId()(*string) {
    return m.sarif_id
}
// GetTool gets the tool property value. The tool property
// returns a CodeScanningAnalysisToolable when successful
func (m *CodeScanningAnalysis) GetTool()(CodeScanningAnalysisToolable) {
    return m.tool
}
// GetUrl gets the url property value. The REST API URL of the analysis resource.
// returns a *string when successful
func (m *CodeScanningAnalysis) GetUrl()(*string) {
    return m.url
}
// GetWarning gets the warning property value. Warning generated when processing the analysis
// returns a *string when successful
func (m *CodeScanningAnalysis) GetWarning()(*string) {
    return m.warning
}
// Serialize serializes information the current object
func (m *CodeScanningAnalysis) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("analysis_key", m.GetAnalysisKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commit_sha", m.GetCommitSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("deletable", m.GetDeletable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("environment", m.GetEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("error", m.GetError())
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
    {
        err := writer.WriteStringValue("ref", m.GetRef())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("results_count", m.GetResultsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("rules_count", m.GetRulesCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sarif_id", m.GetSarifId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("tool", m.GetTool())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("warning", m.GetWarning())
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
func (m *CodeScanningAnalysis) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAnalysisKey sets the analysis_key property value. Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
func (m *CodeScanningAnalysis) SetAnalysisKey(value *string)() {
    m.analysis_key = value
}
// SetCategory sets the category property value. Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
func (m *CodeScanningAnalysis) SetCategory(value *string)() {
    m.category = value
}
// SetCommitSha sets the commit_sha property value. The SHA of the commit to which the analysis you are uploading relates.
func (m *CodeScanningAnalysis) SetCommitSha(value *string)() {
    m.commit_sha = value
}
// SetCreatedAt sets the created_at property value. The time that the analysis was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *CodeScanningAnalysis) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDeletable sets the deletable property value. The deletable property
func (m *CodeScanningAnalysis) SetDeletable(value *bool)() {
    m.deletable = value
}
// SetEnvironment sets the environment property value. Identifies the variable values associated with the environment in which this analysis was performed.
func (m *CodeScanningAnalysis) SetEnvironment(value *string)() {
    m.environment = value
}
// SetError sets the error property value. The error property
func (m *CodeScanningAnalysis) SetError(value *string)() {
    m.error = value
}
// SetId sets the id property value. Unique identifier for this analysis.
func (m *CodeScanningAnalysis) SetId(value *int32)() {
    m.id = value
}
// SetRef sets the ref property value. The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`,`refs/heads/<branch name>` or simply `<branch name>`.
func (m *CodeScanningAnalysis) SetRef(value *string)() {
    m.ref = value
}
// SetResultsCount sets the results_count property value. The total number of results in the analysis.
func (m *CodeScanningAnalysis) SetResultsCount(value *int32)() {
    m.results_count = value
}
// SetRulesCount sets the rules_count property value. The total number of rules used in the analysis.
func (m *CodeScanningAnalysis) SetRulesCount(value *int32)() {
    m.rules_count = value
}
// SetSarifId sets the sarif_id property value. An identifier for the upload.
func (m *CodeScanningAnalysis) SetSarifId(value *string)() {
    m.sarif_id = value
}
// SetTool sets the tool property value. The tool property
func (m *CodeScanningAnalysis) SetTool(value CodeScanningAnalysisToolable)() {
    m.tool = value
}
// SetUrl sets the url property value. The REST API URL of the analysis resource.
func (m *CodeScanningAnalysis) SetUrl(value *string)() {
    m.url = value
}
// SetWarning sets the warning property value. Warning generated when processing the analysis
func (m *CodeScanningAnalysis) SetWarning(value *string)() {
    m.warning = value
}
type CodeScanningAnalysisable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalysisKey()(*string)
    GetCategory()(*string)
    GetCommitSha()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeletable()(*bool)
    GetEnvironment()(*string)
    GetError()(*string)
    GetId()(*int32)
    GetRef()(*string)
    GetResultsCount()(*int32)
    GetRulesCount()(*int32)
    GetSarifId()(*string)
    GetTool()(CodeScanningAnalysisToolable)
    GetUrl()(*string)
    GetWarning()(*string)
    SetAnalysisKey(value *string)()
    SetCategory(value *string)()
    SetCommitSha(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeletable(value *bool)()
    SetEnvironment(value *string)()
    SetError(value *string)()
    SetId(value *int32)()
    SetRef(value *string)()
    SetResultsCount(value *int32)()
    SetRulesCount(value *int32)()
    SetSarifId(value *string)()
    SetTool(value CodeScanningAnalysisToolable)()
    SetUrl(value *string)()
    SetWarning(value *string)()
}
