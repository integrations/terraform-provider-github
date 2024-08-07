package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ActionsBillingUsage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The amount of free GitHub Actions minutes available.
    included_minutes *int32
    // The minutes_used_breakdown property
    minutes_used_breakdown ActionsBillingUsage_minutes_used_breakdownable
    // The sum of the free and paid GitHub Actions minutes used.
    total_minutes_used *int32
    // The total paid GitHub Actions minutes used.
    total_paid_minutes_used *int32
}
// NewActionsBillingUsage instantiates a new ActionsBillingUsage and sets the default values.
func NewActionsBillingUsage()(*ActionsBillingUsage) {
    m := &ActionsBillingUsage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateActionsBillingUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateActionsBillingUsageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionsBillingUsage(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ActionsBillingUsage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ActionsBillingUsage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["included_minutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludedMinutes(val)
        }
        return nil
    }
    res["minutes_used_breakdown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateActionsBillingUsage_minutes_used_breakdownFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinutesUsedBreakdown(val.(ActionsBillingUsage_minutes_used_breakdownable))
        }
        return nil
    }
    res["total_minutes_used"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalMinutesUsed(val)
        }
        return nil
    }
    res["total_paid_minutes_used"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalPaidMinutesUsed(val)
        }
        return nil
    }
    return res
}
// GetIncludedMinutes gets the included_minutes property value. The amount of free GitHub Actions minutes available.
// returns a *int32 when successful
func (m *ActionsBillingUsage) GetIncludedMinutes()(*int32) {
    return m.included_minutes
}
// GetMinutesUsedBreakdown gets the minutes_used_breakdown property value. The minutes_used_breakdown property
// returns a ActionsBillingUsage_minutes_used_breakdownable when successful
func (m *ActionsBillingUsage) GetMinutesUsedBreakdown()(ActionsBillingUsage_minutes_used_breakdownable) {
    return m.minutes_used_breakdown
}
// GetTotalMinutesUsed gets the total_minutes_used property value. The sum of the free and paid GitHub Actions minutes used.
// returns a *int32 when successful
func (m *ActionsBillingUsage) GetTotalMinutesUsed()(*int32) {
    return m.total_minutes_used
}
// GetTotalPaidMinutesUsed gets the total_paid_minutes_used property value. The total paid GitHub Actions minutes used.
// returns a *int32 when successful
func (m *ActionsBillingUsage) GetTotalPaidMinutesUsed()(*int32) {
    return m.total_paid_minutes_used
}
// Serialize serializes information the current object
func (m *ActionsBillingUsage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("included_minutes", m.GetIncludedMinutes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("minutes_used_breakdown", m.GetMinutesUsedBreakdown())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_minutes_used", m.GetTotalMinutesUsed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_paid_minutes_used", m.GetTotalPaidMinutesUsed())
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
func (m *ActionsBillingUsage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIncludedMinutes sets the included_minutes property value. The amount of free GitHub Actions minutes available.
func (m *ActionsBillingUsage) SetIncludedMinutes(value *int32)() {
    m.included_minutes = value
}
// SetMinutesUsedBreakdown sets the minutes_used_breakdown property value. The minutes_used_breakdown property
func (m *ActionsBillingUsage) SetMinutesUsedBreakdown(value ActionsBillingUsage_minutes_used_breakdownable)() {
    m.minutes_used_breakdown = value
}
// SetTotalMinutesUsed sets the total_minutes_used property value. The sum of the free and paid GitHub Actions minutes used.
func (m *ActionsBillingUsage) SetTotalMinutesUsed(value *int32)() {
    m.total_minutes_used = value
}
// SetTotalPaidMinutesUsed sets the total_paid_minutes_used property value. The total paid GitHub Actions minutes used.
func (m *ActionsBillingUsage) SetTotalPaidMinutesUsed(value *int32)() {
    m.total_paid_minutes_used = value
}
type ActionsBillingUsageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIncludedMinutes()(*int32)
    GetMinutesUsedBreakdown()(ActionsBillingUsage_minutes_used_breakdownable)
    GetTotalMinutesUsed()(*int32)
    GetTotalPaidMinutesUsed()(*int32)
    SetIncludedMinutes(value *int32)()
    SetMinutesUsedBreakdown(value ActionsBillingUsage_minutes_used_breakdownable)()
    SetTotalMinutesUsed(value *int32)()
    SetTotalPaidMinutesUsed(value *int32)()
}
