package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Environment_protection_rulesMember2_reviewers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The reviewer property
    reviewer Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewerable
    // The type of reviewer.
    typeEscaped *DeploymentReviewerType
}
// Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewer composed type wrapper for classes SimpleUserable, Teamable
type Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewer struct {
    // Composed type representation for type SimpleUserable
    simpleUser SimpleUserable
    // Composed type representation for type Teamable
    team Teamable
}
// NewEnvironment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewer instantiates a new Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewer and sets the default values.
func NewEnvironment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewer()(*Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewer) {
    m := &Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewer{
    }
    return m
}
// CreateEnvironment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnvironment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewEnvironment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewer()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateSimpleUserFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(SimpleUserable); ok {
                result.SetSimpleUser(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateTeamFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(Teamable); ok {
                result.SetTeam(cast)
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewer) GetIsComposedType()(bool) {
    return true
}
// GetSimpleUser gets the simpleUser property value. Composed type representation for type SimpleUserable
// returns a SimpleUserable when successful
func (m *Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewer) GetSimpleUser()(SimpleUserable) {
    return m.simpleUser
}
// GetTeam gets the team property value. Composed type representation for type Teamable
// returns a Teamable when successful
func (m *Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewer) GetTeam()(Teamable) {
    return m.team
}
// Serialize serializes information the current object
func (m *Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSimpleUser() != nil {
        err := writer.WriteObjectValue("", m.GetSimpleUser())
        if err != nil {
            return err
        }
    } else if m.GetTeam() != nil {
        err := writer.WriteObjectValue("", m.GetTeam())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSimpleUser sets the simpleUser property value. Composed type representation for type SimpleUserable
func (m *Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewer) SetSimpleUser(value SimpleUserable)() {
    m.simpleUser = value
}
// SetTeam sets the team property value. Composed type representation for type Teamable
func (m *Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewer) SetTeam(value Teamable)() {
    m.team = value
}
type Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSimpleUser()(SimpleUserable)
    GetTeam()(Teamable)
    SetSimpleUser(value SimpleUserable)()
    SetTeam(value Teamable)()
}
// NewEnvironment_protection_rulesMember2_reviewers instantiates a new Environment_protection_rulesMember2_reviewers and sets the default values.
func NewEnvironment_protection_rulesMember2_reviewers()(*Environment_protection_rulesMember2_reviewers) {
    m := &Environment_protection_rulesMember2_reviewers{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnvironment_protection_rulesMember2_reviewersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnvironment_protection_rulesMember2_reviewersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnvironment_protection_rulesMember2_reviewers(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Environment_protection_rulesMember2_reviewers) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Environment_protection_rulesMember2_reviewers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["reviewer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnvironment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewer(val.(Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewerable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeploymentReviewerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*DeploymentReviewerType))
        }
        return nil
    }
    return res
}
// GetReviewer gets the reviewer property value. The reviewer property
// returns a Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewerable when successful
func (m *Environment_protection_rulesMember2_reviewers) GetReviewer()(Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewerable) {
    return m.reviewer
}
// GetTypeEscaped gets the type property value. The type of reviewer.
// returns a *DeploymentReviewerType when successful
func (m *Environment_protection_rulesMember2_reviewers) GetTypeEscaped()(*DeploymentReviewerType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Environment_protection_rulesMember2_reviewers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("reviewer", m.GetReviewer())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *Environment_protection_rulesMember2_reviewers) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetReviewer sets the reviewer property value. The reviewer property
func (m *Environment_protection_rulesMember2_reviewers) SetReviewer(value Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewerable)() {
    m.reviewer = value
}
// SetTypeEscaped sets the type property value. The type of reviewer.
func (m *Environment_protection_rulesMember2_reviewers) SetTypeEscaped(value *DeploymentReviewerType)() {
    m.typeEscaped = value
}
type Environment_protection_rulesMember2_reviewersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReviewer()(Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewerable)
    GetTypeEscaped()(*DeploymentReviewerType)
    SetReviewer(value Environment_protection_rulesMember2_reviewers_Environment_protection_rulesMember2_reviewers_reviewerable)()
    SetTypeEscaped(value *DeploymentReviewerType)()
}
