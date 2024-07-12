package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemBranchesItemProtectionPutRequestBody_required_status_checks require status checks to pass before merging. Set to `null` to disable.
type ItemItemBranchesItemProtectionPutRequestBody_required_status_checks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The list of status checks to require in order to merge into this branch.
    checks []ItemItemBranchesItemProtectionPutRequestBody_required_status_checks_checksable
    // **Deprecated**: The list of status checks to require in order to merge into this branch. If any of these checks have recently been set by a particular GitHub App, they will be required to come from that app in future for the branch to merge. Use `checks` instead of `contexts` for more fine-grained control.
    // Deprecated: 
    contexts []string
    // Require branches to be up to date before merging.
    strict *bool
}
// NewItemItemBranchesItemProtectionPutRequestBody_required_status_checks instantiates a new ItemItemBranchesItemProtectionPutRequestBody_required_status_checks and sets the default values.
func NewItemItemBranchesItemProtectionPutRequestBody_required_status_checks()(*ItemItemBranchesItemProtectionPutRequestBody_required_status_checks) {
    m := &ItemItemBranchesItemProtectionPutRequestBody_required_status_checks{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemBranchesItemProtectionPutRequestBody_required_status_checksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemBranchesItemProtectionPutRequestBody_required_status_checksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemBranchesItemProtectionPutRequestBody_required_status_checks(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_status_checks) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChecks gets the checks property value. The list of status checks to require in order to merge into this branch.
// returns a []ItemItemBranchesItemProtectionPutRequestBody_required_status_checks_checksable when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_status_checks) GetChecks()([]ItemItemBranchesItemProtectionPutRequestBody_required_status_checks_checksable) {
    return m.checks
}
// GetContexts gets the contexts property value. **Deprecated**: The list of status checks to require in order to merge into this branch. If any of these checks have recently been set by a particular GitHub App, they will be required to come from that app in future for the branch to merge. Use `checks` instead of `contexts` for more fine-grained control.
// Deprecated: 
// returns a []string when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_status_checks) GetContexts()([]string) {
    return m.contexts
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_status_checks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["checks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemItemBranchesItemProtectionPutRequestBody_required_status_checks_checksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemItemBranchesItemProtectionPutRequestBody_required_status_checks_checksable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ItemItemBranchesItemProtectionPutRequestBody_required_status_checks_checksable)
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
    return res
}
// GetStrict gets the strict property value. Require branches to be up to date before merging.
// returns a *bool when successful
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_status_checks) GetStrict()(*bool) {
    return m.strict
}
// Serialize serializes information the current object
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_status_checks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteBoolValue("strict", m.GetStrict())
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
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_status_checks) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChecks sets the checks property value. The list of status checks to require in order to merge into this branch.
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_status_checks) SetChecks(value []ItemItemBranchesItemProtectionPutRequestBody_required_status_checks_checksable)() {
    m.checks = value
}
// SetContexts sets the contexts property value. **Deprecated**: The list of status checks to require in order to merge into this branch. If any of these checks have recently been set by a particular GitHub App, they will be required to come from that app in future for the branch to merge. Use `checks` instead of `contexts` for more fine-grained control.
// Deprecated: 
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_status_checks) SetContexts(value []string)() {
    m.contexts = value
}
// SetStrict sets the strict property value. Require branches to be up to date before merging.
func (m *ItemItemBranchesItemProtectionPutRequestBody_required_status_checks) SetStrict(value *bool)() {
    m.strict = value
}
type ItemItemBranchesItemProtectionPutRequestBody_required_status_checksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChecks()([]ItemItemBranchesItemProtectionPutRequestBody_required_status_checks_checksable)
    GetContexts()([]string)
    GetStrict()(*bool)
    SetChecks(value []ItemItemBranchesItemProtectionPutRequestBody_required_status_checks_checksable)()
    SetContexts(value []string)()
    SetStrict(value *bool)()
}
