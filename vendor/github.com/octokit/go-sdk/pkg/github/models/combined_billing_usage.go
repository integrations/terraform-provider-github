package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CombinedBillingUsage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Numbers of days left in billing cycle.
    days_left_in_billing_cycle *int32
    // Estimated storage space (GB) used in billing cycle.
    estimated_paid_storage_for_month *int32
    // Estimated sum of free and paid storage space (GB) used in billing cycle.
    estimated_storage_for_month *int32
}
// NewCombinedBillingUsage instantiates a new CombinedBillingUsage and sets the default values.
func NewCombinedBillingUsage()(*CombinedBillingUsage) {
    m := &CombinedBillingUsage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCombinedBillingUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCombinedBillingUsageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCombinedBillingUsage(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CombinedBillingUsage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDaysLeftInBillingCycle gets the days_left_in_billing_cycle property value. Numbers of days left in billing cycle.
// returns a *int32 when successful
func (m *CombinedBillingUsage) GetDaysLeftInBillingCycle()(*int32) {
    return m.days_left_in_billing_cycle
}
// GetEstimatedPaidStorageForMonth gets the estimated_paid_storage_for_month property value. Estimated storage space (GB) used in billing cycle.
// returns a *int32 when successful
func (m *CombinedBillingUsage) GetEstimatedPaidStorageForMonth()(*int32) {
    return m.estimated_paid_storage_for_month
}
// GetEstimatedStorageForMonth gets the estimated_storage_for_month property value. Estimated sum of free and paid storage space (GB) used in billing cycle.
// returns a *int32 when successful
func (m *CombinedBillingUsage) GetEstimatedStorageForMonth()(*int32) {
    return m.estimated_storage_for_month
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CombinedBillingUsage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["days_left_in_billing_cycle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDaysLeftInBillingCycle(val)
        }
        return nil
    }
    res["estimated_paid_storage_for_month"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEstimatedPaidStorageForMonth(val)
        }
        return nil
    }
    res["estimated_storage_for_month"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEstimatedStorageForMonth(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CombinedBillingUsage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("days_left_in_billing_cycle", m.GetDaysLeftInBillingCycle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("estimated_paid_storage_for_month", m.GetEstimatedPaidStorageForMonth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("estimated_storage_for_month", m.GetEstimatedStorageForMonth())
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
func (m *CombinedBillingUsage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDaysLeftInBillingCycle sets the days_left_in_billing_cycle property value. Numbers of days left in billing cycle.
func (m *CombinedBillingUsage) SetDaysLeftInBillingCycle(value *int32)() {
    m.days_left_in_billing_cycle = value
}
// SetEstimatedPaidStorageForMonth sets the estimated_paid_storage_for_month property value. Estimated storage space (GB) used in billing cycle.
func (m *CombinedBillingUsage) SetEstimatedPaidStorageForMonth(value *int32)() {
    m.estimated_paid_storage_for_month = value
}
// SetEstimatedStorageForMonth sets the estimated_storage_for_month property value. Estimated sum of free and paid storage space (GB) used in billing cycle.
func (m *CombinedBillingUsage) SetEstimatedStorageForMonth(value *int32)() {
    m.estimated_storage_for_month = value
}
type CombinedBillingUsageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDaysLeftInBillingCycle()(*int32)
    GetEstimatedPaidStorageForMonth()(*int32)
    GetEstimatedStorageForMonth()(*int32)
    SetDaysLeftInBillingCycle(value *int32)()
    SetEstimatedPaidStorageForMonth(value *int32)()
    SetEstimatedStorageForMonth(value *int32)()
}
