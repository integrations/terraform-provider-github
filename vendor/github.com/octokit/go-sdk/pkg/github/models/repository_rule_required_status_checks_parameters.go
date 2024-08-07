package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryRuleRequiredStatusChecks_parameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Status checks that are required.
    required_status_checks []RepositoryRuleParamsStatusCheckConfigurationable
    // Whether pull requests targeting a matching branch must be tested with the latest code. This setting will not take effect unless at least one status check is enabled.
    strict_required_status_checks_policy *bool
}
// NewRepositoryRuleRequiredStatusChecks_parameters instantiates a new RepositoryRuleRequiredStatusChecks_parameters and sets the default values.
func NewRepositoryRuleRequiredStatusChecks_parameters()(*RepositoryRuleRequiredStatusChecks_parameters) {
    m := &RepositoryRuleRequiredStatusChecks_parameters{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRuleRequiredStatusChecks_parametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRuleRequiredStatusChecks_parametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRuleRequiredStatusChecks_parameters(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRuleRequiredStatusChecks_parameters) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRuleRequiredStatusChecks_parameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["required_status_checks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryRuleParamsStatusCheckConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RepositoryRuleParamsStatusCheckConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RepositoryRuleParamsStatusCheckConfigurationable)
                }
            }
            m.SetRequiredStatusChecks(res)
        }
        return nil
    }
    res["strict_required_status_checks_policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStrictRequiredStatusChecksPolicy(val)
        }
        return nil
    }
    return res
}
// GetRequiredStatusChecks gets the required_status_checks property value. Status checks that are required.
// returns a []RepositoryRuleParamsStatusCheckConfigurationable when successful
func (m *RepositoryRuleRequiredStatusChecks_parameters) GetRequiredStatusChecks()([]RepositoryRuleParamsStatusCheckConfigurationable) {
    return m.required_status_checks
}
// GetStrictRequiredStatusChecksPolicy gets the strict_required_status_checks_policy property value. Whether pull requests targeting a matching branch must be tested with the latest code. This setting will not take effect unless at least one status check is enabled.
// returns a *bool when successful
func (m *RepositoryRuleRequiredStatusChecks_parameters) GetStrictRequiredStatusChecksPolicy()(*bool) {
    return m.strict_required_status_checks_policy
}
// Serialize serializes information the current object
func (m *RepositoryRuleRequiredStatusChecks_parameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRequiredStatusChecks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRequiredStatusChecks()))
        for i, v := range m.GetRequiredStatusChecks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("required_status_checks", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("strict_required_status_checks_policy", m.GetStrictRequiredStatusChecksPolicy())
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
func (m *RepositoryRuleRequiredStatusChecks_parameters) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRequiredStatusChecks sets the required_status_checks property value. Status checks that are required.
func (m *RepositoryRuleRequiredStatusChecks_parameters) SetRequiredStatusChecks(value []RepositoryRuleParamsStatusCheckConfigurationable)() {
    m.required_status_checks = value
}
// SetStrictRequiredStatusChecksPolicy sets the strict_required_status_checks_policy property value. Whether pull requests targeting a matching branch must be tested with the latest code. This setting will not take effect unless at least one status check is enabled.
func (m *RepositoryRuleRequiredStatusChecks_parameters) SetStrictRequiredStatusChecksPolicy(value *bool)() {
    m.strict_required_status_checks_policy = value
}
type RepositoryRuleRequiredStatusChecks_parametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRequiredStatusChecks()([]RepositoryRuleParamsStatusCheckConfigurationable)
    GetStrictRequiredStatusChecksPolicy()(*bool)
    SetRequiredStatusChecks(value []RepositoryRuleParamsStatusCheckConfigurationable)()
    SetStrictRequiredStatusChecksPolicy(value *bool)()
}
