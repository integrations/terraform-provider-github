package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryRulesetConditions_ref_name struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.
    exclude []string
    // Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.
    include []string
}
// NewRepositoryRulesetConditions_ref_name instantiates a new RepositoryRulesetConditions_ref_name and sets the default values.
func NewRepositoryRulesetConditions_ref_name()(*RepositoryRulesetConditions_ref_name) {
    m := &RepositoryRulesetConditions_ref_name{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryRulesetConditions_ref_nameFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryRulesetConditions_ref_nameFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryRulesetConditions_ref_name(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryRulesetConditions_ref_name) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExclude gets the exclude property value. Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.
// returns a []string when successful
func (m *RepositoryRulesetConditions_ref_name) GetExclude()([]string) {
    return m.exclude
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryRulesetConditions_ref_name) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["exclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetExclude(res)
        }
        return nil
    }
    res["include"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetInclude(res)
        }
        return nil
    }
    return res
}
// GetInclude gets the include property value. Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.
// returns a []string when successful
func (m *RepositoryRulesetConditions_ref_name) GetInclude()([]string) {
    return m.include
}
// Serialize serializes information the current object
func (m *RepositoryRulesetConditions_ref_name) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExclude() != nil {
        err := writer.WriteCollectionOfStringValues("exclude", m.GetExclude())
        if err != nil {
            return err
        }
    }
    if m.GetInclude() != nil {
        err := writer.WriteCollectionOfStringValues("include", m.GetInclude())
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
func (m *RepositoryRulesetConditions_ref_name) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExclude sets the exclude property value. Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.
func (m *RepositoryRulesetConditions_ref_name) SetExclude(value []string)() {
    m.exclude = value
}
// SetInclude sets the include property value. Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.
func (m *RepositoryRulesetConditions_ref_name) SetInclude(value []string)() {
    m.include = value
}
type RepositoryRulesetConditions_ref_nameable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExclude()([]string)
    GetInclude()([]string)
    SetExclude(value []string)()
    SetInclude(value []string)()
}
