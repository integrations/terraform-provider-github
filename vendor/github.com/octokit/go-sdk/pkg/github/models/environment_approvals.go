package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnvironmentApprovals an entry in the reviews log for environment deployments
type EnvironmentApprovals struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The comment submitted with the deployment review
    comment *string
    // The list of environments that were approved or rejected
    environments []EnvironmentApprovals_environmentsable
    // Whether deployment to the environment(s) was approved or rejected or pending (with comments)
    state *EnvironmentApprovals_state
    // A GitHub user.
    user SimpleUserable
}
// NewEnvironmentApprovals instantiates a new EnvironmentApprovals and sets the default values.
func NewEnvironmentApprovals()(*EnvironmentApprovals) {
    m := &EnvironmentApprovals{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEnvironmentApprovalsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEnvironmentApprovalsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnvironmentApprovals(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EnvironmentApprovals) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The comment submitted with the deployment review
// returns a *string when successful
func (m *EnvironmentApprovals) GetComment()(*string) {
    return m.comment
}
// GetEnvironments gets the environments property value. The list of environments that were approved or rejected
// returns a []EnvironmentApprovals_environmentsable when successful
func (m *EnvironmentApprovals) GetEnvironments()([]EnvironmentApprovals_environmentsable) {
    return m.environments
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EnvironmentApprovals) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["environments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEnvironmentApprovals_environmentsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EnvironmentApprovals_environmentsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EnvironmentApprovals_environmentsable)
                }
            }
            m.SetEnvironments(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnvironmentApprovals_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*EnvironmentApprovals_state))
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(SimpleUserable))
        }
        return nil
    }
    return res
}
// GetState gets the state property value. Whether deployment to the environment(s) was approved or rejected or pending (with comments)
// returns a *EnvironmentApprovals_state when successful
func (m *EnvironmentApprovals) GetState()(*EnvironmentApprovals_state) {
    return m.state
}
// GetUser gets the user property value. A GitHub user.
// returns a SimpleUserable when successful
func (m *EnvironmentApprovals) GetUser()(SimpleUserable) {
    return m.user
}
// Serialize serializes information the current object
func (m *EnvironmentApprovals) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    if m.GetEnvironments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEnvironments()))
        for i, v := range m.GetEnvironments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("environments", cast)
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
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *EnvironmentApprovals) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The comment submitted with the deployment review
func (m *EnvironmentApprovals) SetComment(value *string)() {
    m.comment = value
}
// SetEnvironments sets the environments property value. The list of environments that were approved or rejected
func (m *EnvironmentApprovals) SetEnvironments(value []EnvironmentApprovals_environmentsable)() {
    m.environments = value
}
// SetState sets the state property value. Whether deployment to the environment(s) was approved or rejected or pending (with comments)
func (m *EnvironmentApprovals) SetState(value *EnvironmentApprovals_state)() {
    m.state = value
}
// SetUser sets the user property value. A GitHub user.
func (m *EnvironmentApprovals) SetUser(value SimpleUserable)() {
    m.user = value
}
type EnvironmentApprovalsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComment()(*string)
    GetEnvironments()([]EnvironmentApprovals_environmentsable)
    GetState()(*EnvironmentApprovals_state)
    GetUser()(SimpleUserable)
    SetComment(value *string)()
    SetEnvironments(value []EnvironmentApprovals_environmentsable)()
    SetState(value *EnvironmentApprovals_state)()
    SetUser(value SimpleUserable)()
}
