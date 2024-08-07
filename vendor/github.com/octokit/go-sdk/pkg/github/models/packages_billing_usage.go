package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackagesBillingUsage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Free storage space (GB) for GitHub Packages.
    included_gigabytes_bandwidth *int32
    // Sum of the free and paid storage space (GB) for GitHuub Packages.
    total_gigabytes_bandwidth_used *int32
    // Total paid storage space (GB) for GitHuub Packages.
    total_paid_gigabytes_bandwidth_used *int32
}
// NewPackagesBillingUsage instantiates a new PackagesBillingUsage and sets the default values.
func NewPackagesBillingUsage()(*PackagesBillingUsage) {
    m := &PackagesBillingUsage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePackagesBillingUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackagesBillingUsageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPackagesBillingUsage(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PackagesBillingUsage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackagesBillingUsage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["included_gigabytes_bandwidth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludedGigabytesBandwidth(val)
        }
        return nil
    }
    res["total_gigabytes_bandwidth_used"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalGigabytesBandwidthUsed(val)
        }
        return nil
    }
    res["total_paid_gigabytes_bandwidth_used"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalPaidGigabytesBandwidthUsed(val)
        }
        return nil
    }
    return res
}
// GetIncludedGigabytesBandwidth gets the included_gigabytes_bandwidth property value. Free storage space (GB) for GitHub Packages.
// returns a *int32 when successful
func (m *PackagesBillingUsage) GetIncludedGigabytesBandwidth()(*int32) {
    return m.included_gigabytes_bandwidth
}
// GetTotalGigabytesBandwidthUsed gets the total_gigabytes_bandwidth_used property value. Sum of the free and paid storage space (GB) for GitHuub Packages.
// returns a *int32 when successful
func (m *PackagesBillingUsage) GetTotalGigabytesBandwidthUsed()(*int32) {
    return m.total_gigabytes_bandwidth_used
}
// GetTotalPaidGigabytesBandwidthUsed gets the total_paid_gigabytes_bandwidth_used property value. Total paid storage space (GB) for GitHuub Packages.
// returns a *int32 when successful
func (m *PackagesBillingUsage) GetTotalPaidGigabytesBandwidthUsed()(*int32) {
    return m.total_paid_gigabytes_bandwidth_used
}
// Serialize serializes information the current object
func (m *PackagesBillingUsage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("included_gigabytes_bandwidth", m.GetIncludedGigabytesBandwidth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_gigabytes_bandwidth_used", m.GetTotalGigabytesBandwidthUsed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_paid_gigabytes_bandwidth_used", m.GetTotalPaidGigabytesBandwidthUsed())
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
func (m *PackagesBillingUsage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIncludedGigabytesBandwidth sets the included_gigabytes_bandwidth property value. Free storage space (GB) for GitHub Packages.
func (m *PackagesBillingUsage) SetIncludedGigabytesBandwidth(value *int32)() {
    m.included_gigabytes_bandwidth = value
}
// SetTotalGigabytesBandwidthUsed sets the total_gigabytes_bandwidth_used property value. Sum of the free and paid storage space (GB) for GitHuub Packages.
func (m *PackagesBillingUsage) SetTotalGigabytesBandwidthUsed(value *int32)() {
    m.total_gigabytes_bandwidth_used = value
}
// SetTotalPaidGigabytesBandwidthUsed sets the total_paid_gigabytes_bandwidth_used property value. Total paid storage space (GB) for GitHuub Packages.
func (m *PackagesBillingUsage) SetTotalPaidGigabytesBandwidthUsed(value *int32)() {
    m.total_paid_gigabytes_bandwidth_used = value
}
type PackagesBillingUsageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIncludedGigabytesBandwidth()(*int32)
    GetTotalGigabytesBandwidthUsed()(*int32)
    GetTotalPaidGigabytesBandwidthUsed()(*int32)
    SetIncludedGigabytesBandwidth(value *int32)()
    SetTotalGigabytesBandwidthUsed(value *int32)()
    SetTotalPaidGigabytesBandwidthUsed(value *int32)()
}
