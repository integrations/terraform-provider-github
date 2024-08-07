package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemItemEnvironmentsItemDeployment_protection_rulesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The custom_deployment_protection_rules property
    custom_deployment_protection_rules []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentProtectionRuleable
    // The number of enabled custom deployment protection rules for this environment
    total_count *int32
}
// NewItemItemEnvironmentsItemDeployment_protection_rulesGetResponse instantiates a new ItemItemEnvironmentsItemDeployment_protection_rulesGetResponse and sets the default values.
func NewItemItemEnvironmentsItemDeployment_protection_rulesGetResponse()(*ItemItemEnvironmentsItemDeployment_protection_rulesGetResponse) {
    m := &ItemItemEnvironmentsItemDeployment_protection_rulesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemEnvironmentsItemDeployment_protection_rulesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemEnvironmentsItemDeployment_protection_rulesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemEnvironmentsItemDeployment_protection_rulesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomDeploymentProtectionRules gets the custom_deployment_protection_rules property value. The custom_deployment_protection_rules property
// returns a []DeploymentProtectionRuleable when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesGetResponse) GetCustomDeploymentProtectionRules()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentProtectionRuleable) {
    return m.custom_deployment_protection_rules
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["custom_deployment_protection_rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateDeploymentProtectionRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentProtectionRuleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentProtectionRuleable)
                }
            }
            m.SetCustomDeploymentProtectionRules(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetTotalCount gets the total_count property value. The number of enabled custom deployment protection rules for this environment
// returns a *int32 when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCustomDeploymentProtectionRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomDeploymentProtectionRules()))
        for i, v := range m.GetCustomDeploymentProtectionRules() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("custom_deployment_protection_rules", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
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
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomDeploymentProtectionRules sets the custom_deployment_protection_rules property value. The custom_deployment_protection_rules property
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesGetResponse) SetCustomDeploymentProtectionRules(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentProtectionRuleable)() {
    m.custom_deployment_protection_rules = value
}
// SetTotalCount sets the total_count property value. The number of enabled custom deployment protection rules for this environment
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
type ItemItemEnvironmentsItemDeployment_protection_rulesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomDeploymentProtectionRules()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentProtectionRuleable)
    GetTotalCount()(*int32)
    SetCustomDeploymentProtectionRules(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.DeploymentProtectionRuleable)()
    SetTotalCount(value *int32)()
}
