package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Environment_protection_rulesMember2 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id property
    id *int32
    // The node_id property
    node_id *string
    // Whether deployments to this environment can be approved by the user who created the deployment.
    prevent_self_review *bool
    // The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
    reviewers []Environment_protection_rulesMember2_reviewersable
    // The type property
    typeEscaped *string
}
// NewEnvironment_protection_rulesMember2 instantiates a new Environment_protection_rulesMember2 and sets the default values.
func NewEnvironment_protection_rulesMember2()(*Environment_protection_rulesMember2) {
    m := &Environment_protection_rulesMember2{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnvironment_protection_rulesMember2FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnvironment_protection_rulesMember2FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnvironment_protection_rulesMember2(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Environment_protection_rulesMember2) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Environment_protection_rulesMember2) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
        }
        return nil
    }
    res["prevent_self_review"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreventSelfReview(val)
        }
        return nil
    }
    res["reviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEnvironment_protection_rulesMember2_reviewersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Environment_protection_rulesMember2_reviewersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Environment_protection_rulesMember2_reviewersable)
                }
            }
            m.SetReviewers(res)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *Environment_protection_rulesMember2) GetId()(*int32) {
    return m.id
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *Environment_protection_rulesMember2) GetNodeId()(*string) {
    return m.node_id
}
// GetPreventSelfReview gets the prevent_self_review property value. Whether deployments to this environment can be approved by the user who created the deployment.
// returns a *bool when successful
func (m *Environment_protection_rulesMember2) GetPreventSelfReview()(*bool) {
    return m.prevent_self_review
}
// GetReviewers gets the reviewers property value. The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
// returns a []Environment_protection_rulesMember2_reviewersable when successful
func (m *Environment_protection_rulesMember2) GetReviewers()([]Environment_protection_rulesMember2_reviewersable) {
    return m.reviewers
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *Environment_protection_rulesMember2) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Environment_protection_rulesMember2) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("prevent_self_review", m.GetPreventSelfReview())
        if err != nil {
            return err
        }
    }
    if m.GetReviewers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("reviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *Environment_protection_rulesMember2) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The id property
func (m *Environment_protection_rulesMember2) SetId(value *int32)() {
    m.id = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *Environment_protection_rulesMember2) SetNodeId(value *string)() {
    m.node_id = value
}
// SetPreventSelfReview sets the prevent_self_review property value. Whether deployments to this environment can be approved by the user who created the deployment.
func (m *Environment_protection_rulesMember2) SetPreventSelfReview(value *bool)() {
    m.prevent_self_review = value
}
// SetReviewers sets the reviewers property value. The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
func (m *Environment_protection_rulesMember2) SetReviewers(value []Environment_protection_rulesMember2_reviewersable)() {
    m.reviewers = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *Environment_protection_rulesMember2) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type Environment_protection_rulesMember2able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int32)
    GetNodeId()(*string)
    GetPreventSelfReview()(*bool)
    GetReviewers()([]Environment_protection_rulesMember2_reviewersable)
    GetTypeEscaped()(*string)
    SetId(value *int32)()
    SetNodeId(value *string)()
    SetPreventSelfReview(value *bool)()
    SetReviewers(value []Environment_protection_rulesMember2_reviewersable)()
    SetTypeEscaped(value *string)()
}
