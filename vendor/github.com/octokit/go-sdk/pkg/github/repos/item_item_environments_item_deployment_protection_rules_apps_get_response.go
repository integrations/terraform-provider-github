package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The available_custom_deployment_protection_rule_integrations property
    available_custom_deployment_protection_rule_integrations []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CustomDeploymentRuleAppable
    // The total number of custom deployment protection rule integrations available for this environment.
    total_count *int32
}
// NewItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse instantiates a new ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse and sets the default values.
func NewItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse()(*ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse) {
    m := &ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailableCustomDeploymentProtectionRuleIntegrations gets the available_custom_deployment_protection_rule_integrations property value. The available_custom_deployment_protection_rule_integrations property
// returns a []CustomDeploymentRuleAppable when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse) GetAvailableCustomDeploymentProtectionRuleIntegrations()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CustomDeploymentRuleAppable) {
    return m.available_custom_deployment_protection_rule_integrations
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["available_custom_deployment_protection_rule_integrations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateCustomDeploymentRuleAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CustomDeploymentRuleAppable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CustomDeploymentRuleAppable)
                }
            }
            m.SetAvailableCustomDeploymentProtectionRuleIntegrations(res)
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
// GetTotalCount gets the total_count property value. The total number of custom deployment protection rule integrations available for this environment.
// returns a *int32 when successful
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAvailableCustomDeploymentProtectionRuleIntegrations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAvailableCustomDeploymentProtectionRuleIntegrations()))
        for i, v := range m.GetAvailableCustomDeploymentProtectionRuleIntegrations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("available_custom_deployment_protection_rule_integrations", cast)
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
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailableCustomDeploymentProtectionRuleIntegrations sets the available_custom_deployment_protection_rule_integrations property value. The available_custom_deployment_protection_rule_integrations property
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse) SetAvailableCustomDeploymentProtectionRuleIntegrations(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CustomDeploymentRuleAppable)() {
    m.available_custom_deployment_protection_rule_integrations = value
}
// SetTotalCount sets the total_count property value. The total number of custom deployment protection rule integrations available for this environment.
func (m *ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
type ItemItemEnvironmentsItemDeployment_protection_rulesAppsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailableCustomDeploymentProtectionRuleIntegrations()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CustomDeploymentRuleAppable)
    GetTotalCount()(*int32)
    SetAvailableCustomDeploymentProtectionRuleIntegrations(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CustomDeploymentRuleAppable)()
    SetTotalCount(value *int32)()
}
