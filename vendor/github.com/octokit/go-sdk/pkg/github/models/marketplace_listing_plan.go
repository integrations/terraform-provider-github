package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MarketplaceListingPlan marketplace Listing Plan
type MarketplaceListingPlan struct {
    // The accounts_url property
    accounts_url *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The bullets property
    bullets []string
    // The description property
    description *string
    // The has_free_trial property
    has_free_trial *bool
    // The id property
    id *int32
    // The monthly_price_in_cents property
    monthly_price_in_cents *int32
    // The name property
    name *string
    // The number property
    number *int32
    // The price_model property
    price_model *MarketplaceListingPlan_price_model
    // The state property
    state *string
    // The unit_name property
    unit_name *string
    // The url property
    url *string
    // The yearly_price_in_cents property
    yearly_price_in_cents *int32
}
// NewMarketplaceListingPlan instantiates a new MarketplaceListingPlan and sets the default values.
func NewMarketplaceListingPlan()(*MarketplaceListingPlan) {
    m := &MarketplaceListingPlan{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMarketplaceListingPlanFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMarketplaceListingPlanFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMarketplaceListingPlan(), nil
}
// GetAccountsUrl gets the accounts_url property value. The accounts_url property
// returns a *string when successful
func (m *MarketplaceListingPlan) GetAccountsUrl()(*string) {
    return m.accounts_url
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MarketplaceListingPlan) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBullets gets the bullets property value. The bullets property
// returns a []string when successful
func (m *MarketplaceListingPlan) GetBullets()([]string) {
    return m.bullets
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *MarketplaceListingPlan) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MarketplaceListingPlan) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accounts_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountsUrl(val)
        }
        return nil
    }
    res["bullets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetBullets(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["has_free_trial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasFreeTrial(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["monthly_price_in_cents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonthlyPriceInCents(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["price_model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMarketplaceListingPlan_price_model)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriceModel(val.(*MarketplaceListingPlan_price_model))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["unit_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnitName(val)
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
    res["yearly_price_in_cents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYearlyPriceInCents(val)
        }
        return nil
    }
    return res
}
// GetHasFreeTrial gets the has_free_trial property value. The has_free_trial property
// returns a *bool when successful
func (m *MarketplaceListingPlan) GetHasFreeTrial()(*bool) {
    return m.has_free_trial
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *MarketplaceListingPlan) GetId()(*int32) {
    return m.id
}
// GetMonthlyPriceInCents gets the monthly_price_in_cents property value. The monthly_price_in_cents property
// returns a *int32 when successful
func (m *MarketplaceListingPlan) GetMonthlyPriceInCents()(*int32) {
    return m.monthly_price_in_cents
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *MarketplaceListingPlan) GetName()(*string) {
    return m.name
}
// GetNumber gets the number property value. The number property
// returns a *int32 when successful
func (m *MarketplaceListingPlan) GetNumber()(*int32) {
    return m.number
}
// GetPriceModel gets the price_model property value. The price_model property
// returns a *MarketplaceListingPlan_price_model when successful
func (m *MarketplaceListingPlan) GetPriceModel()(*MarketplaceListingPlan_price_model) {
    return m.price_model
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *MarketplaceListingPlan) GetState()(*string) {
    return m.state
}
// GetUnitName gets the unit_name property value. The unit_name property
// returns a *string when successful
func (m *MarketplaceListingPlan) GetUnitName()(*string) {
    return m.unit_name
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *MarketplaceListingPlan) GetUrl()(*string) {
    return m.url
}
// GetYearlyPriceInCents gets the yearly_price_in_cents property value. The yearly_price_in_cents property
// returns a *int32 when successful
func (m *MarketplaceListingPlan) GetYearlyPriceInCents()(*int32) {
    return m.yearly_price_in_cents
}
// Serialize serializes information the current object
func (m *MarketplaceListingPlan) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("accounts_url", m.GetAccountsUrl())
        if err != nil {
            return err
        }
    }
    if m.GetBullets() != nil {
        err := writer.WriteCollectionOfStringValues("bullets", m.GetBullets())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_free_trial", m.GetHasFreeTrial())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("monthly_price_in_cents", m.GetMonthlyPriceInCents())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    if m.GetPriceModel() != nil {
        cast := (*m.GetPriceModel()).String()
        err := writer.WriteStringValue("price_model", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("unit_name", m.GetUnitName())
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
        err := writer.WriteInt32Value("yearly_price_in_cents", m.GetYearlyPriceInCents())
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
// SetAccountsUrl sets the accounts_url property value. The accounts_url property
func (m *MarketplaceListingPlan) SetAccountsUrl(value *string)() {
    m.accounts_url = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MarketplaceListingPlan) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBullets sets the bullets property value. The bullets property
func (m *MarketplaceListingPlan) SetBullets(value []string)() {
    m.bullets = value
}
// SetDescription sets the description property value. The description property
func (m *MarketplaceListingPlan) SetDescription(value *string)() {
    m.description = value
}
// SetHasFreeTrial sets the has_free_trial property value. The has_free_trial property
func (m *MarketplaceListingPlan) SetHasFreeTrial(value *bool)() {
    m.has_free_trial = value
}
// SetId sets the id property value. The id property
func (m *MarketplaceListingPlan) SetId(value *int32)() {
    m.id = value
}
// SetMonthlyPriceInCents sets the monthly_price_in_cents property value. The monthly_price_in_cents property
func (m *MarketplaceListingPlan) SetMonthlyPriceInCents(value *int32)() {
    m.monthly_price_in_cents = value
}
// SetName sets the name property value. The name property
func (m *MarketplaceListingPlan) SetName(value *string)() {
    m.name = value
}
// SetNumber sets the number property value. The number property
func (m *MarketplaceListingPlan) SetNumber(value *int32)() {
    m.number = value
}
// SetPriceModel sets the price_model property value. The price_model property
func (m *MarketplaceListingPlan) SetPriceModel(value *MarketplaceListingPlan_price_model)() {
    m.price_model = value
}
// SetState sets the state property value. The state property
func (m *MarketplaceListingPlan) SetState(value *string)() {
    m.state = value
}
// SetUnitName sets the unit_name property value. The unit_name property
func (m *MarketplaceListingPlan) SetUnitName(value *string)() {
    m.unit_name = value
}
// SetUrl sets the url property value. The url property
func (m *MarketplaceListingPlan) SetUrl(value *string)() {
    m.url = value
}
// SetYearlyPriceInCents sets the yearly_price_in_cents property value. The yearly_price_in_cents property
func (m *MarketplaceListingPlan) SetYearlyPriceInCents(value *int32)() {
    m.yearly_price_in_cents = value
}
type MarketplaceListingPlanable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountsUrl()(*string)
    GetBullets()([]string)
    GetDescription()(*string)
    GetHasFreeTrial()(*bool)
    GetId()(*int32)
    GetMonthlyPriceInCents()(*int32)
    GetName()(*string)
    GetNumber()(*int32)
    GetPriceModel()(*MarketplaceListingPlan_price_model)
    GetState()(*string)
    GetUnitName()(*string)
    GetUrl()(*string)
    GetYearlyPriceInCents()(*int32)
    SetAccountsUrl(value *string)()
    SetBullets(value []string)()
    SetDescription(value *string)()
    SetHasFreeTrial(value *bool)()
    SetId(value *int32)()
    SetMonthlyPriceInCents(value *int32)()
    SetName(value *string)()
    SetNumber(value *int32)()
    SetPriceModel(value *MarketplaceListingPlan_price_model)()
    SetState(value *string)()
    SetUnitName(value *string)()
    SetUrl(value *string)()
    SetYearlyPriceInCents(value *int32)()
}
