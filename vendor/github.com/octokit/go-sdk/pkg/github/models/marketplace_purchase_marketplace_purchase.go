package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MarketplacePurchase_marketplace_purchase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The billing_cycle property
    billing_cycle *string
    // The free_trial_ends_on property
    free_trial_ends_on *string
    // The is_installed property
    is_installed *bool
    // The next_billing_date property
    next_billing_date *string
    // The on_free_trial property
    on_free_trial *bool
    // Marketplace Listing Plan
    plan MarketplaceListingPlanable
    // The unit_count property
    unit_count *int32
    // The updated_at property
    updated_at *string
}
// NewMarketplacePurchase_marketplace_purchase instantiates a new MarketplacePurchase_marketplace_purchase and sets the default values.
func NewMarketplacePurchase_marketplace_purchase()(*MarketplacePurchase_marketplace_purchase) {
    m := &MarketplacePurchase_marketplace_purchase{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMarketplacePurchase_marketplace_purchaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMarketplacePurchase_marketplace_purchaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMarketplacePurchase_marketplace_purchase(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MarketplacePurchase_marketplace_purchase) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBillingCycle gets the billing_cycle property value. The billing_cycle property
// returns a *string when successful
func (m *MarketplacePurchase_marketplace_purchase) GetBillingCycle()(*string) {
    return m.billing_cycle
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MarketplacePurchase_marketplace_purchase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["billing_cycle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBillingCycle(val)
        }
        return nil
    }
    res["free_trial_ends_on"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFreeTrialEndsOn(val)
        }
        return nil
    }
    res["is_installed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInstalled(val)
        }
        return nil
    }
    res["next_billing_date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextBillingDate(val)
        }
        return nil
    }
    res["on_free_trial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnFreeTrial(val)
        }
        return nil
    }
    res["plan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMarketplaceListingPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlan(val.(MarketplaceListingPlanable))
        }
        return nil
    }
    res["unit_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnitCount(val)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetFreeTrialEndsOn gets the free_trial_ends_on property value. The free_trial_ends_on property
// returns a *string when successful
func (m *MarketplacePurchase_marketplace_purchase) GetFreeTrialEndsOn()(*string) {
    return m.free_trial_ends_on
}
// GetIsInstalled gets the is_installed property value. The is_installed property
// returns a *bool when successful
func (m *MarketplacePurchase_marketplace_purchase) GetIsInstalled()(*bool) {
    return m.is_installed
}
// GetNextBillingDate gets the next_billing_date property value. The next_billing_date property
// returns a *string when successful
func (m *MarketplacePurchase_marketplace_purchase) GetNextBillingDate()(*string) {
    return m.next_billing_date
}
// GetOnFreeTrial gets the on_free_trial property value. The on_free_trial property
// returns a *bool when successful
func (m *MarketplacePurchase_marketplace_purchase) GetOnFreeTrial()(*bool) {
    return m.on_free_trial
}
// GetPlan gets the plan property value. Marketplace Listing Plan
// returns a MarketplaceListingPlanable when successful
func (m *MarketplacePurchase_marketplace_purchase) GetPlan()(MarketplaceListingPlanable) {
    return m.plan
}
// GetUnitCount gets the unit_count property value. The unit_count property
// returns a *int32 when successful
func (m *MarketplacePurchase_marketplace_purchase) GetUnitCount()(*int32) {
    return m.unit_count
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *string when successful
func (m *MarketplacePurchase_marketplace_purchase) GetUpdatedAt()(*string) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *MarketplacePurchase_marketplace_purchase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("billing_cycle", m.GetBillingCycle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("free_trial_ends_on", m.GetFreeTrialEndsOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_installed", m.GetIsInstalled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("next_billing_date", m.GetNextBillingDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("on_free_trial", m.GetOnFreeTrial())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("plan", m.GetPlan())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unit_count", m.GetUnitCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("updated_at", m.GetUpdatedAt())
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
func (m *MarketplacePurchase_marketplace_purchase) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBillingCycle sets the billing_cycle property value. The billing_cycle property
func (m *MarketplacePurchase_marketplace_purchase) SetBillingCycle(value *string)() {
    m.billing_cycle = value
}
// SetFreeTrialEndsOn sets the free_trial_ends_on property value. The free_trial_ends_on property
func (m *MarketplacePurchase_marketplace_purchase) SetFreeTrialEndsOn(value *string)() {
    m.free_trial_ends_on = value
}
// SetIsInstalled sets the is_installed property value. The is_installed property
func (m *MarketplacePurchase_marketplace_purchase) SetIsInstalled(value *bool)() {
    m.is_installed = value
}
// SetNextBillingDate sets the next_billing_date property value. The next_billing_date property
func (m *MarketplacePurchase_marketplace_purchase) SetNextBillingDate(value *string)() {
    m.next_billing_date = value
}
// SetOnFreeTrial sets the on_free_trial property value. The on_free_trial property
func (m *MarketplacePurchase_marketplace_purchase) SetOnFreeTrial(value *bool)() {
    m.on_free_trial = value
}
// SetPlan sets the plan property value. Marketplace Listing Plan
func (m *MarketplacePurchase_marketplace_purchase) SetPlan(value MarketplaceListingPlanable)() {
    m.plan = value
}
// SetUnitCount sets the unit_count property value. The unit_count property
func (m *MarketplacePurchase_marketplace_purchase) SetUnitCount(value *int32)() {
    m.unit_count = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *MarketplacePurchase_marketplace_purchase) SetUpdatedAt(value *string)() {
    m.updated_at = value
}
type MarketplacePurchase_marketplace_purchaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBillingCycle()(*string)
    GetFreeTrialEndsOn()(*string)
    GetIsInstalled()(*bool)
    GetNextBillingDate()(*string)
    GetOnFreeTrial()(*bool)
    GetPlan()(MarketplaceListingPlanable)
    GetUnitCount()(*int32)
    GetUpdatedAt()(*string)
    SetBillingCycle(value *string)()
    SetFreeTrialEndsOn(value *string)()
    SetIsInstalled(value *bool)()
    SetNextBillingDate(value *string)()
    SetOnFreeTrial(value *bool)()
    SetPlan(value MarketplaceListingPlanable)()
    SetUnitCount(value *int32)()
    SetUpdatedAt(value *string)()
}
