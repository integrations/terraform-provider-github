package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MarketplacePurchase_marketplace_pending_change struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The effective_date property
    effective_date *string
    // The id property
    id *int32
    // The is_installed property
    is_installed *bool
    // Marketplace Listing Plan
    plan MarketplaceListingPlanable
    // The unit_count property
    unit_count *int32
}
// NewMarketplacePurchase_marketplace_pending_change instantiates a new MarketplacePurchase_marketplace_pending_change and sets the default values.
func NewMarketplacePurchase_marketplace_pending_change()(*MarketplacePurchase_marketplace_pending_change) {
    m := &MarketplacePurchase_marketplace_pending_change{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMarketplacePurchase_marketplace_pending_changeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMarketplacePurchase_marketplace_pending_changeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMarketplacePurchase_marketplace_pending_change(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MarketplacePurchase_marketplace_pending_change) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEffectiveDate gets the effective_date property value. The effective_date property
// returns a *string when successful
func (m *MarketplacePurchase_marketplace_pending_change) GetEffectiveDate()(*string) {
    return m.effective_date
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MarketplacePurchase_marketplace_pending_change) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["effective_date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEffectiveDate(val)
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
    return res
}
// GetId gets the id property value. The id property
// returns a *int32 when successful
func (m *MarketplacePurchase_marketplace_pending_change) GetId()(*int32) {
    return m.id
}
// GetIsInstalled gets the is_installed property value. The is_installed property
// returns a *bool when successful
func (m *MarketplacePurchase_marketplace_pending_change) GetIsInstalled()(*bool) {
    return m.is_installed
}
// GetPlan gets the plan property value. Marketplace Listing Plan
// returns a MarketplaceListingPlanable when successful
func (m *MarketplacePurchase_marketplace_pending_change) GetPlan()(MarketplaceListingPlanable) {
    return m.plan
}
// GetUnitCount gets the unit_count property value. The unit_count property
// returns a *int32 when successful
func (m *MarketplacePurchase_marketplace_pending_change) GetUnitCount()(*int32) {
    return m.unit_count
}
// Serialize serializes information the current object
func (m *MarketplacePurchase_marketplace_pending_change) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("effective_date", m.GetEffectiveDate())
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
        err := writer.WriteBoolValue("is_installed", m.GetIsInstalled())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MarketplacePurchase_marketplace_pending_change) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEffectiveDate sets the effective_date property value. The effective_date property
func (m *MarketplacePurchase_marketplace_pending_change) SetEffectiveDate(value *string)() {
    m.effective_date = value
}
// SetId sets the id property value. The id property
func (m *MarketplacePurchase_marketplace_pending_change) SetId(value *int32)() {
    m.id = value
}
// SetIsInstalled sets the is_installed property value. The is_installed property
func (m *MarketplacePurchase_marketplace_pending_change) SetIsInstalled(value *bool)() {
    m.is_installed = value
}
// SetPlan sets the plan property value. Marketplace Listing Plan
func (m *MarketplacePurchase_marketplace_pending_change) SetPlan(value MarketplaceListingPlanable)() {
    m.plan = value
}
// SetUnitCount sets the unit_count property value. The unit_count property
func (m *MarketplacePurchase_marketplace_pending_change) SetUnitCount(value *int32)() {
    m.unit_count = value
}
type MarketplacePurchase_marketplace_pending_changeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEffectiveDate()(*string)
    GetId()(*int32)
    GetIsInstalled()(*bool)
    GetPlan()(MarketplaceListingPlanable)
    GetUnitCount()(*int32)
    SetEffectiveDate(value *string)()
    SetId(value *int32)()
    SetIsInstalled(value *bool)()
    SetPlan(value MarketplaceListingPlanable)()
    SetUnitCount(value *int32)()
}
