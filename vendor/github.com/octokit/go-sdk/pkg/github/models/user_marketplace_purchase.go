package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserMarketplacePurchase user Marketplace Purchase
type UserMarketplacePurchase struct {
    // The account property
    account MarketplaceAccountable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The billing_cycle property
    billing_cycle *string
    // The free_trial_ends_on property
    free_trial_ends_on *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The next_billing_date property
    next_billing_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The on_free_trial property
    on_free_trial *bool
    // Marketplace Listing Plan
    plan MarketplaceListingPlanable
    // The unit_count property
    unit_count *int32
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewUserMarketplacePurchase instantiates a new UserMarketplacePurchase and sets the default values.
func NewUserMarketplacePurchase()(*UserMarketplacePurchase) {
    m := &UserMarketplacePurchase{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserMarketplacePurchaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserMarketplacePurchaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserMarketplacePurchase(), nil
}
// GetAccount gets the account property value. The account property
// returns a MarketplaceAccountable when successful
func (m *UserMarketplacePurchase) GetAccount()(MarketplaceAccountable) {
    return m.account
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UserMarketplacePurchase) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBillingCycle gets the billing_cycle property value. The billing_cycle property
// returns a *string when successful
func (m *UserMarketplacePurchase) GetBillingCycle()(*string) {
    return m.billing_cycle
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserMarketplacePurchase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMarketplaceAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccount(val.(MarketplaceAccountable))
        }
        return nil
    }
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
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFreeTrialEndsOn(val)
        }
        return nil
    }
    res["next_billing_date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
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
        val, err := n.GetTimeValue()
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
// returns a *Time when successful
func (m *UserMarketplacePurchase) GetFreeTrialEndsOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.free_trial_ends_on
}
// GetNextBillingDate gets the next_billing_date property value. The next_billing_date property
// returns a *Time when successful
func (m *UserMarketplacePurchase) GetNextBillingDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.next_billing_date
}
// GetOnFreeTrial gets the on_free_trial property value. The on_free_trial property
// returns a *bool when successful
func (m *UserMarketplacePurchase) GetOnFreeTrial()(*bool) {
    return m.on_free_trial
}
// GetPlan gets the plan property value. Marketplace Listing Plan
// returns a MarketplaceListingPlanable when successful
func (m *UserMarketplacePurchase) GetPlan()(MarketplaceListingPlanable) {
    return m.plan
}
// GetUnitCount gets the unit_count property value. The unit_count property
// returns a *int32 when successful
func (m *UserMarketplacePurchase) GetUnitCount()(*int32) {
    return m.unit_count
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *UserMarketplacePurchase) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *UserMarketplacePurchase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("account", m.GetAccount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("billing_cycle", m.GetBillingCycle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("free_trial_ends_on", m.GetFreeTrialEndsOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("next_billing_date", m.GetNextBillingDate())
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
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
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
// SetAccount sets the account property value. The account property
func (m *UserMarketplacePurchase) SetAccount(value MarketplaceAccountable)() {
    m.account = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserMarketplacePurchase) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBillingCycle sets the billing_cycle property value. The billing_cycle property
func (m *UserMarketplacePurchase) SetBillingCycle(value *string)() {
    m.billing_cycle = value
}
// SetFreeTrialEndsOn sets the free_trial_ends_on property value. The free_trial_ends_on property
func (m *UserMarketplacePurchase) SetFreeTrialEndsOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.free_trial_ends_on = value
}
// SetNextBillingDate sets the next_billing_date property value. The next_billing_date property
func (m *UserMarketplacePurchase) SetNextBillingDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.next_billing_date = value
}
// SetOnFreeTrial sets the on_free_trial property value. The on_free_trial property
func (m *UserMarketplacePurchase) SetOnFreeTrial(value *bool)() {
    m.on_free_trial = value
}
// SetPlan sets the plan property value. Marketplace Listing Plan
func (m *UserMarketplacePurchase) SetPlan(value MarketplaceListingPlanable)() {
    m.plan = value
}
// SetUnitCount sets the unit_count property value. The unit_count property
func (m *UserMarketplacePurchase) SetUnitCount(value *int32)() {
    m.unit_count = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *UserMarketplacePurchase) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
type UserMarketplacePurchaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()(MarketplaceAccountable)
    GetBillingCycle()(*string)
    GetFreeTrialEndsOn()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNextBillingDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnFreeTrial()(*bool)
    GetPlan()(MarketplaceListingPlanable)
    GetUnitCount()(*int32)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccount(value MarketplaceAccountable)()
    SetBillingCycle(value *string)()
    SetFreeTrialEndsOn(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNextBillingDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnFreeTrial(value *bool)()
    SetPlan(value MarketplaceListingPlanable)()
    SetUnitCount(value *int32)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
