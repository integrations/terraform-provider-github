package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectedBranchRequiredStatusCheck protected Branch Required Status Check
type ProtectedBranchRequiredStatusCheck struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The checks property
    checks []ProtectedBranchRequiredStatusCheck_checksable
    // The contexts property
    contexts []string
    // The contexts_url property
    contexts_url *string
    // The enforcement_level property
    enforcement_level *string
    // The strict property
    strict *bool
    // The url property
    url *string
}
// NewProtectedBranchRequiredStatusCheck instantiates a new ProtectedBranchRequiredStatusCheck and sets the default values.
func NewProtectedBranchRequiredStatusCheck()(*ProtectedBranchRequiredStatusCheck) {
    m := &ProtectedBranchRequiredStatusCheck{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProtectedBranchRequiredStatusCheckFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProtectedBranchRequiredStatusCheckFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectedBranchRequiredStatusCheck(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProtectedBranchRequiredStatusCheck) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChecks gets the checks property value. The checks property
// returns a []ProtectedBranchRequiredStatusCheck_checksable when successful
func (m *ProtectedBranchRequiredStatusCheck) GetChecks()([]ProtectedBranchRequiredStatusCheck_checksable) {
    return m.checks
}
// GetContexts gets the contexts property value. The contexts property
// returns a []string when successful
func (m *ProtectedBranchRequiredStatusCheck) GetContexts()([]string) {
    return m.contexts
}
// GetContextsUrl gets the contexts_url property value. The contexts_url property
// returns a *string when successful
func (m *ProtectedBranchRequiredStatusCheck) GetContextsUrl()(*string) {
    return m.contexts_url
}
// GetEnforcementLevel gets the enforcement_level property value. The enforcement_level property
// returns a *string when successful
func (m *ProtectedBranchRequiredStatusCheck) GetEnforcementLevel()(*string) {
    return m.enforcement_level
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProtectedBranchRequiredStatusCheck) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["checks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProtectedBranchRequiredStatusCheck_checksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProtectedBranchRequiredStatusCheck_checksable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProtectedBranchRequiredStatusCheck_checksable)
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
    res["enforcement_level"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforcementLevel(val)
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
func (m *ProtectedBranchRequiredStatusCheck) GetStrict()(*bool) {
    return m.strict
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *ProtectedBranchRequiredStatusCheck) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ProtectedBranchRequiredStatusCheck) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("enforcement_level", m.GetEnforcementLevel())
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
func (m *ProtectedBranchRequiredStatusCheck) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChecks sets the checks property value. The checks property
func (m *ProtectedBranchRequiredStatusCheck) SetChecks(value []ProtectedBranchRequiredStatusCheck_checksable)() {
    m.checks = value
}
// SetContexts sets the contexts property value. The contexts property
func (m *ProtectedBranchRequiredStatusCheck) SetContexts(value []string)() {
    m.contexts = value
}
// SetContextsUrl sets the contexts_url property value. The contexts_url property
func (m *ProtectedBranchRequiredStatusCheck) SetContextsUrl(value *string)() {
    m.contexts_url = value
}
// SetEnforcementLevel sets the enforcement_level property value. The enforcement_level property
func (m *ProtectedBranchRequiredStatusCheck) SetEnforcementLevel(value *string)() {
    m.enforcement_level = value
}
// SetStrict sets the strict property value. The strict property
func (m *ProtectedBranchRequiredStatusCheck) SetStrict(value *bool)() {
    m.strict = value
}
// SetUrl sets the url property value. The url property
func (m *ProtectedBranchRequiredStatusCheck) SetUrl(value *string)() {
    m.url = value
}
type ProtectedBranchRequiredStatusCheckable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChecks()([]ProtectedBranchRequiredStatusCheck_checksable)
    GetContexts()([]string)
    GetContextsUrl()(*string)
    GetEnforcementLevel()(*string)
    GetStrict()(*bool)
    GetUrl()(*string)
    SetChecks(value []ProtectedBranchRequiredStatusCheck_checksable)()
    SetContexts(value []string)()
    SetContextsUrl(value *string)()
    SetEnforcementLevel(value *string)()
    SetStrict(value *bool)()
    SetUrl(value *string)()
}
