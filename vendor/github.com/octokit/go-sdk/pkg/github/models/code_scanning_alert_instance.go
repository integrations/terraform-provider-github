package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CodeScanningAlertInstance struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
    analysis_key *string
    // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
    category *string
    // Classifications that have been applied to the file that triggered the alert.For example identifying it as documentation, or a generated file.
    classifications []CodeScanningAlertClassification
    // The commit_sha property
    commit_sha *string
    // Identifies the variable values associated with the environment in which the analysis that generated this alert instance was performed, such as the language that was analyzed.
    environment *string
    // The html_url property
    html_url *string
    // Describe a region within a file for the alert.
    location CodeScanningAlertLocationable
    // The message property
    message CodeScanningAlertInstance_messageable
    // The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`,`refs/heads/<branch name>` or simply `<branch name>`.
    ref *string
    // State of a code scanning alert.
    state *CodeScanningAlertState
}
// NewCodeScanningAlertInstance instantiates a new CodeScanningAlertInstance and sets the default values.
func NewCodeScanningAlertInstance()(*CodeScanningAlertInstance) {
    m := &CodeScanningAlertInstance{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningAlertInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningAlertInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningAlertInstance(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningAlertInstance) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAnalysisKey gets the analysis_key property value. Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
// returns a *string when successful
func (m *CodeScanningAlertInstance) GetAnalysisKey()(*string) {
    return m.analysis_key
}
// GetCategory gets the category property value. Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
// returns a *string when successful
func (m *CodeScanningAlertInstance) GetCategory()(*string) {
    return m.category
}
// GetClassifications gets the classifications property value. Classifications that have been applied to the file that triggered the alert.For example identifying it as documentation, or a generated file.
// returns a []CodeScanningAlertClassification when successful
func (m *CodeScanningAlertInstance) GetClassifications()([]CodeScanningAlertClassification) {
    return m.classifications
}
// GetCommitSha gets the commit_sha property value. The commit_sha property
// returns a *string when successful
func (m *CodeScanningAlertInstance) GetCommitSha()(*string) {
    return m.commit_sha
}
// GetEnvironment gets the environment property value. Identifies the variable values associated with the environment in which the analysis that generated this alert instance was performed, such as the language that was analyzed.
// returns a *string when successful
func (m *CodeScanningAlertInstance) GetEnvironment()(*string) {
    return m.environment
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningAlertInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["classifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseCodeScanningAlertClassification)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CodeScanningAlertClassification, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*CodeScanningAlertClassification))
                }
            }
            m.SetClassifications(res)
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
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeScanningAlertLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(CodeScanningAlertLocationable))
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCodeScanningAlertInstance_messageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val.(CodeScanningAlertInstance_messageable))
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningAlertState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*CodeScanningAlertState))
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *CodeScanningAlertInstance) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetLocation gets the location property value. Describe a region within a file for the alert.
// returns a CodeScanningAlertLocationable when successful
func (m *CodeScanningAlertInstance) GetLocation()(CodeScanningAlertLocationable) {
    return m.location
}
// GetMessage gets the message property value. The message property
// returns a CodeScanningAlertInstance_messageable when successful
func (m *CodeScanningAlertInstance) GetMessage()(CodeScanningAlertInstance_messageable) {
    return m.message
}
// GetRef gets the ref property value. The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`,`refs/heads/<branch name>` or simply `<branch name>`.
// returns a *string when successful
func (m *CodeScanningAlertInstance) GetRef()(*string) {
    return m.ref
}
// GetState gets the state property value. State of a code scanning alert.
// returns a *CodeScanningAlertState when successful
func (m *CodeScanningAlertInstance) GetState()(*CodeScanningAlertState) {
    return m.state
}
// Serialize serializes information the current object
func (m *CodeScanningAlertInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetClassifications() != nil {
        err := writer.WriteCollectionOfStringValues("classifications", SerializeCodeScanningAlertClassification(m.GetClassifications()))
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
        err := writer.WriteStringValue("environment", m.GetEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("message", m.GetMessage())
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
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *CodeScanningAlertInstance) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAnalysisKey sets the analysis_key property value. Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
func (m *CodeScanningAlertInstance) SetAnalysisKey(value *string)() {
    m.analysis_key = value
}
// SetCategory sets the category property value. Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
func (m *CodeScanningAlertInstance) SetCategory(value *string)() {
    m.category = value
}
// SetClassifications sets the classifications property value. Classifications that have been applied to the file that triggered the alert.For example identifying it as documentation, or a generated file.
func (m *CodeScanningAlertInstance) SetClassifications(value []CodeScanningAlertClassification)() {
    m.classifications = value
}
// SetCommitSha sets the commit_sha property value. The commit_sha property
func (m *CodeScanningAlertInstance) SetCommitSha(value *string)() {
    m.commit_sha = value
}
// SetEnvironment sets the environment property value. Identifies the variable values associated with the environment in which the analysis that generated this alert instance was performed, such as the language that was analyzed.
func (m *CodeScanningAlertInstance) SetEnvironment(value *string)() {
    m.environment = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *CodeScanningAlertInstance) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetLocation sets the location property value. Describe a region within a file for the alert.
func (m *CodeScanningAlertInstance) SetLocation(value CodeScanningAlertLocationable)() {
    m.location = value
}
// SetMessage sets the message property value. The message property
func (m *CodeScanningAlertInstance) SetMessage(value CodeScanningAlertInstance_messageable)() {
    m.message = value
}
// SetRef sets the ref property value. The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`,`refs/heads/<branch name>` or simply `<branch name>`.
func (m *CodeScanningAlertInstance) SetRef(value *string)() {
    m.ref = value
}
// SetState sets the state property value. State of a code scanning alert.
func (m *CodeScanningAlertInstance) SetState(value *CodeScanningAlertState)() {
    m.state = value
}
type CodeScanningAlertInstanceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalysisKey()(*string)
    GetCategory()(*string)
    GetClassifications()([]CodeScanningAlertClassification)
    GetCommitSha()(*string)
    GetEnvironment()(*string)
    GetHtmlUrl()(*string)
    GetLocation()(CodeScanningAlertLocationable)
    GetMessage()(CodeScanningAlertInstance_messageable)
    GetRef()(*string)
    GetState()(*CodeScanningAlertState)
    SetAnalysisKey(value *string)()
    SetCategory(value *string)()
    SetClassifications(value []CodeScanningAlertClassification)()
    SetCommitSha(value *string)()
    SetEnvironment(value *string)()
    SetHtmlUrl(value *string)()
    SetLocation(value CodeScanningAlertLocationable)()
    SetMessage(value CodeScanningAlertInstance_messageable)()
    SetRef(value *string)()
    SetState(value *CodeScanningAlertState)()
}
