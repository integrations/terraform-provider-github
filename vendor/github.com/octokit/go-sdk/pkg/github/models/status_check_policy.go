package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StatusCheckPolicy status Check Policy
type StatusCheckPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The checks property
    checks []StatusCheckPolicy_checksable
    // The contexts property
    contexts []string
    // The contexts_url property
    contexts_url *string
    // The strict property
    strict *bool
    // The url property
    url *string
}
// NewStatusCheckPolicy instantiates a new StatusCheckPolicy and sets the default values.
func NewStatusCheckPolicy()(*StatusCheckPolicy) {
    m := &StatusCheckPolicy{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStatusCheckPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStatusCheckPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStatusCheckPolicy(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StatusCheckPolicy) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChecks gets the checks property value. The checks property
// returns a []StatusCheckPolicy_checksable when successful
func (m *StatusCheckPolicy) GetChecks()([]StatusCheckPolicy_checksable) {
    return m.checks
}
// GetContexts gets the contexts property value. The contexts property
// returns a []string when successful
func (m *StatusCheckPolicy) GetContexts()([]string) {
    return m.contexts
}
// GetContextsUrl gets the contexts_url property value. The contexts_url property
// returns a *string when successful
func (m *StatusCheckPolicy) GetContextsUrl()(*string) {
    return m.contexts_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StatusCheckPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["checks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateStatusCheckPolicy_checksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]StatusCheckPolicy_checksable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(StatusCheckPolicy_checksable)
                }
            }
            m.SetChecks(res)
        }
        return nil
    }
    res["contexts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetContexts(res)
        }
        return nil
    }
    res["contexts_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContextsUrl(val)
        }
        return nil
    }
    res["strict"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStrict(val)
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
    return res
}
// GetStrict gets the strict property value. The strict property
// returns a *bool when successful
func (m *StatusCheckPolicy) GetStrict()(*bool) {
    return m.strict
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *StatusCheckPolicy) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *StatusCheckPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetChecks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChecks()))
        for i, v := range m.GetChecks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("checks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetContexts() != nil {
        err := writer.WriteCollectionOfStringValues("contexts", m.GetContexts())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contexts_url", m.GetContextsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("strict", m.GetStrict())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *StatusCheckPolicy) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChecks sets the checks property value. The checks property
func (m *StatusCheckPolicy) SetChecks(value []StatusCheckPolicy_checksable)() {
    m.checks = value
}
// SetContexts sets the contexts property value. The contexts property
func (m *StatusCheckPolicy) SetContexts(value []string)() {
    m.contexts = value
}
// SetContextsUrl sets the contexts_url property value. The contexts_url property
func (m *StatusCheckPolicy) SetContextsUrl(value *string)() {
    m.contexts_url = value
}
// SetStrict sets the strict property value. The strict property
func (m *StatusCheckPolicy) SetStrict(value *bool)() {
    m.strict = value
}
// SetUrl sets the url property value. The url property
func (m *StatusCheckPolicy) SetUrl(value *string)() {
    m.url = value
}
type StatusCheckPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChecks()([]StatusCheckPolicy_checksable)
    GetContexts()([]string)
    GetContextsUrl()(*string)
    GetStrict()(*bool)
    GetUrl()(*string)
    SetChecks(value []StatusCheckPolicy_checksable)()
    SetContexts(value []string)()
    SetContextsUrl(value *string)()
    SetStrict(value *bool)()
    SetUrl(value *string)()
}
