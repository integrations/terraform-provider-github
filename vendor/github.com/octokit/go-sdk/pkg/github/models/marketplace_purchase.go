package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MarketplacePurchase marketplace Purchase
type MarketplacePurchase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The email property
    email *string
    // The id property
    id *int32
    // The login property
    login *string
    // The marketplace_pending_change property
    marketplace_pending_change MarketplacePurchase_marketplace_pending_changeable
    // The marketplace_purchase property
    marketplace_purchase MarketplacePurchase_marketplace_purchaseable
    // The organization_billing_email property
    organization_billing_email *string
    // The type property
    typeEscaped *string
    // The url property
    url *string
}
// NewMarketplacePurchase instantiates a new MarketplacePurchase and sets the default values.
func NewMarketplacePurchase()(*MarketplacePurchase) {
    m := &MarketplacePurchase{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMarketplacePurchaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMarketplacePurchaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMarketplacePurchase(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MarketplacePurchase) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *MarketplacePurchase) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MarketplacePurchase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
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
    res["login"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogin(val)
        }
        return nil
    }
    res["marketplace_pending_change"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMarketplacePurchase_marketplace_pending_changeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMarketplacePendingChange(val.(MarketplacePurchase_marketplace_pending_changeable))
        }
        return nil
    }
    res["marketplace_purchase"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMarketplacePurchase_marketplace_purchaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMarketplacePurchase(val.(MarketplacePurchase_marketplace_purchaseable))
        }
        return nil
    }
    res["organization_billing_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationBillingEmail(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
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
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *MarketplacePurchase) GetId()(*int32) {
    return m.id
}
// GetLogin gets the login property value. The login property
// returns a *string when successful
func (m *MarketplacePurchase) GetLogin()(*string) {
    return m.login
}
// GetMarketplacePendingChange gets the marketplace_pending_change property value. The marketplace_pending_change property
// returns a MarketplacePurchase_marketplace_pending_changeable when successful
func (m *MarketplacePurchase) GetMarketplacePendingChange()(MarketplacePurchase_marketplace_pending_changeable) {
    return m.marketplace_pending_change
}
// GetMarketplacePurchase gets the marketplace_purchase property value. The marketplace_purchase property
// returns a MarketplacePurchase_marketplace_purchaseable when successful
func (m *MarketplacePurchase) GetMarketplacePurchase()(MarketplacePurchase_marketplace_purchaseable) {
    return m.marketplace_purchase
}
// GetOrganizationBillingEmail gets the organization_billing_email property value. The organization_billing_email property
// returns a *string when successful
func (m *MarketplacePurchase) GetOrganizationBillingEmail()(*string) {
    return m.organization_billing_email
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *MarketplacePurchase) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *MarketplacePurchase) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *MarketplacePurchase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
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
        err := writer.WriteStringValue("login", m.GetLogin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("marketplace_pending_change", m.GetMarketplacePendingChange())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("marketplace_purchase", m.GetMarketplacePurchase())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("organization_billing_email", m.GetOrganizationBillingEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *MarketplacePurchase) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmail sets the email property value. The email property
func (m *MarketplacePurchase) SetEmail(value *string)() {
    m.email = value
}
// SetId sets the id property value. The id property
func (m *MarketplacePurchase) SetId(value *int32)() {
    m.id = value
}
// SetLogin sets the login property value. The login property
func (m *MarketplacePurchase) SetLogin(value *string)() {
    m.login = value
}
// SetMarketplacePendingChange sets the marketplace_pending_change property value. The marketplace_pending_change property
func (m *MarketplacePurchase) SetMarketplacePendingChange(value MarketplacePurchase_marketplace_pending_changeable)() {
    m.marketplace_pending_change = value
}
// SetMarketplacePurchase sets the marketplace_purchase property value. The marketplace_purchase property
func (m *MarketplacePurchase) SetMarketplacePurchase(value MarketplacePurchase_marketplace_purchaseable)() {
    m.marketplace_purchase = value
}
// SetOrganizationBillingEmail sets the organization_billing_email property value. The organization_billing_email property
func (m *MarketplacePurchase) SetOrganizationBillingEmail(value *string)() {
    m.organization_billing_email = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *MarketplacePurchase) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUrl sets the url property value. The url property
func (m *MarketplacePurchase) SetUrl(value *string)() {
    m.url = value
}
type MarketplacePurchaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetId()(*int32)
    GetLogin()(*string)
    GetMarketplacePendingChange()(MarketplacePurchase_marketplace_pending_changeable)
    GetMarketplacePurchase()(MarketplacePurchase_marketplace_purchaseable)
    GetOrganizationBillingEmail()(*string)
    GetTypeEscaped()(*string)
    GetUrl()(*string)
    SetEmail(value *string)()
    SetId(value *int32)()
    SetLogin(value *string)()
    SetMarketplacePendingChange(value MarketplacePurchase_marketplace_pending_changeable)()
    SetMarketplacePurchase(value MarketplacePurchase_marketplace_purchaseable)()
    SetOrganizationBillingEmail(value *string)()
    SetTypeEscaped(value *string)()
    SetUrl(value *string)()
}
