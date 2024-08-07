package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RuleSuites struct {
    // The number that identifies the user.
    actor_id *int32
    // The handle for the GitHub user account.
    actor_name *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The last commit sha in the push evaluation.
    after_sha *string
    // The first commit sha before the push evaluation.
    before_sha *string
    // The result of the rule evaluations for rules with the `active` and `evaluate` enforcement statuses, demonstrating whether rules would pass or fail if all rules in the rule suite were `active`.
    evaluation_result *RuleSuites_evaluation_result
    // The unique identifier of the rule insight.
    id *int32
    // The pushed_at property
    pushed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ref name that the evaluation ran on.
    ref *string
    // The ID of the repository associated with the rule evaluation.
    repository_id *int32
    // The name of the repository without the `.git` extension.
    repository_name *string
    // The result of the rule evaluations for rules with the `active` enforcement status.
    result *RuleSuites_result
}
// NewRuleSuites instantiates a new RuleSuites and sets the default values.
func NewRuleSuites()(*RuleSuites) {
    m := &RuleSuites{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRuleSuitesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRuleSuitesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRuleSuites(), nil
}
// GetActorId gets the actor_id property value. The number that identifies the user.
// returns a *int32 when successful
func (m *RuleSuites) GetActorId()(*int32) {
    return m.actor_id
}
// GetActorName gets the actor_name property value. The handle for the GitHub user account.
// returns a *string when successful
func (m *RuleSuites) GetActorName()(*string) {
    return m.actor_name
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RuleSuites) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAfterSha gets the after_sha property value. The last commit sha in the push evaluation.
// returns a *string when successful
func (m *RuleSuites) GetAfterSha()(*string) {
    return m.after_sha
}
// GetBeforeSha gets the before_sha property value. The first commit sha before the push evaluation.
// returns a *string when successful
func (m *RuleSuites) GetBeforeSha()(*string) {
    return m.before_sha
}
// GetEvaluationResult gets the evaluation_result property value. The result of the rule evaluations for rules with the `active` and `evaluate` enforcement statuses, demonstrating whether rules would pass or fail if all rules in the rule suite were `active`.
// returns a *RuleSuites_evaluation_result when successful
func (m *RuleSuites) GetEvaluationResult()(*RuleSuites_evaluation_result) {
    return m.evaluation_result
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RuleSuites) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actor_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActorId(val)
        }
        return nil
    }
    res["actor_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActorName(val)
        }
        return nil
    }
    res["after_sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAfterSha(val)
        }
        return nil
    }
    res["before_sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBeforeSha(val)
        }
        return nil
    }
    res["evaluation_result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRuleSuites_evaluation_result)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluationResult(val.(*RuleSuites_evaluation_result))
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
    res["pushed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPushedAt(val)
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
    res["repository_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryId(val)
        }
        return nil
    }
    res["repository_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryName(val)
        }
        return nil
    }
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRuleSuites_result)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(*RuleSuites_result))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The unique identifier of the rule insight.
// returns a *int32 when successful
func (m *RuleSuites) GetId()(*int32) {
    return m.id
}
// GetPushedAt gets the pushed_at property value. The pushed_at property
// returns a *Time when successful
func (m *RuleSuites) GetPushedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.pushed_at
}
// GetRef gets the ref property value. The ref name that the evaluation ran on.
// returns a *string when successful
func (m *RuleSuites) GetRef()(*string) {
    return m.ref
}
// GetRepositoryId gets the repository_id property value. The ID of the repository associated with the rule evaluation.
// returns a *int32 when successful
func (m *RuleSuites) GetRepositoryId()(*int32) {
    return m.repository_id
}
// GetRepositoryName gets the repository_name property value. The name of the repository without the `.git` extension.
// returns a *string when successful
func (m *RuleSuites) GetRepositoryName()(*string) {
    return m.repository_name
}
// GetResult gets the result property value. The result of the rule evaluations for rules with the `active` enforcement status.
// returns a *RuleSuites_result when successful
func (m *RuleSuites) GetResult()(*RuleSuites_result) {
    return m.result
}
// Serialize serializes information the current object
func (m *RuleSuites) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("actor_id", m.GetActorId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("actor_name", m.GetActorName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("after_sha", m.GetAfterSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("before_sha", m.GetBeforeSha())
        if err != nil {
            return err
        }
    }
    if m.GetEvaluationResult() != nil {
        cast := (*m.GetEvaluationResult()).String()
        err := writer.WriteStringValue("evaluation_result", &cast)
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
        err := writer.WriteTimeValue("pushed_at", m.GetPushedAt())
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
        err := writer.WriteInt32Value("repository_id", m.GetRepositoryId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repository_name", m.GetRepositoryName())
        if err != nil {
            return err
        }
    }
    if m.GetResult() != nil {
        cast := (*m.GetResult()).String()
        err := writer.WriteStringValue("result", &cast)
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
// SetActorId sets the actor_id property value. The number that identifies the user.
func (m *RuleSuites) SetActorId(value *int32)() {
    m.actor_id = value
}
// SetActorName sets the actor_name property value. The handle for the GitHub user account.
func (m *RuleSuites) SetActorName(value *string)() {
    m.actor_name = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RuleSuites) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAfterSha sets the after_sha property value. The last commit sha in the push evaluation.
func (m *RuleSuites) SetAfterSha(value *string)() {
    m.after_sha = value
}
// SetBeforeSha sets the before_sha property value. The first commit sha before the push evaluation.
func (m *RuleSuites) SetBeforeSha(value *string)() {
    m.before_sha = value
}
// SetEvaluationResult sets the evaluation_result property value. The result of the rule evaluations for rules with the `active` and `evaluate` enforcement statuses, demonstrating whether rules would pass or fail if all rules in the rule suite were `active`.
func (m *RuleSuites) SetEvaluationResult(value *RuleSuites_evaluation_result)() {
    m.evaluation_result = value
}
// SetId sets the id property value. The unique identifier of the rule insight.
func (m *RuleSuites) SetId(value *int32)() {
    m.id = value
}
// SetPushedAt sets the pushed_at property value. The pushed_at property
func (m *RuleSuites) SetPushedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.pushed_at = value
}
// SetRef sets the ref property value. The ref name that the evaluation ran on.
func (m *RuleSuites) SetRef(value *string)() {
    m.ref = value
}
// SetRepositoryId sets the repository_id property value. The ID of the repository associated with the rule evaluation.
func (m *RuleSuites) SetRepositoryId(value *int32)() {
    m.repository_id = value
}
// SetRepositoryName sets the repository_name property value. The name of the repository without the `.git` extension.
func (m *RuleSuites) SetRepositoryName(value *string)() {
    m.repository_name = value
}
// SetResult sets the result property value. The result of the rule evaluations for rules with the `active` enforcement status.
func (m *RuleSuites) SetResult(value *RuleSuites_result)() {
    m.result = value
}
type RuleSuitesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActorId()(*int32)
    GetActorName()(*string)
    GetAfterSha()(*string)
    GetBeforeSha()(*string)
    GetEvaluationResult()(*RuleSuites_evaluation_result)
    GetId()(*int32)
    GetPushedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRef()(*string)
    GetRepositoryId()(*int32)
    GetRepositoryName()(*string)
    GetResult()(*RuleSuites_result)
    SetActorId(value *int32)()
    SetActorName(value *string)()
    SetAfterSha(value *string)()
    SetBeforeSha(value *string)()
    SetEvaluationResult(value *RuleSuites_evaluation_result)()
    SetId(value *int32)()
    SetPushedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRef(value *string)()
    SetRepositoryId(value *int32)()
    SetRepositoryName(value *string)()
    SetResult(value *RuleSuites_result)()
}
