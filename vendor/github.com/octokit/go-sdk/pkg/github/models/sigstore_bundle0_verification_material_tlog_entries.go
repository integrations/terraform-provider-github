package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SigstoreBundle0_verificationMaterial_tlogEntries struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The canonicalizedBody property
    canonicalizedBody *string
    // The inclusionPromise property
    inclusionPromise SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromiseable
    // The inclusionProof property
    inclusionProof *string
    // The integratedTime property
    integratedTime *string
    // The kindVersion property
    kindVersion SigstoreBundle0_verificationMaterial_tlogEntries_kindVersionable
    // The logId property
    logId SigstoreBundle0_verificationMaterial_tlogEntries_logIdable
    // The logIndex property
    logIndex *string
}
// NewSigstoreBundle0_verificationMaterial_tlogEntries instantiates a new SigstoreBundle0_verificationMaterial_tlogEntries and sets the default values.
func NewSigstoreBundle0_verificationMaterial_tlogEntries()(*SigstoreBundle0_verificationMaterial_tlogEntries) {
    m := &SigstoreBundle0_verificationMaterial_tlogEntries{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSigstoreBundle0_verificationMaterial_tlogEntriesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSigstoreBundle0_verificationMaterial_tlogEntriesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSigstoreBundle0_verificationMaterial_tlogEntries(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCanonicalizedBody gets the canonicalizedBody property value. The canonicalizedBody property
// returns a *string when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) GetCanonicalizedBody()(*string) {
    return m.canonicalizedBody
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["canonicalizedBody"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanonicalizedBody(val)
        }
        return nil
    }
    res["inclusionPromise"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromiseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInclusionPromise(val.(SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromiseable))
        }
        return nil
    }
    res["inclusionProof"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInclusionProof(val)
        }
        return nil
    }
    res["integratedTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegratedTime(val)
        }
        return nil
    }
    res["kindVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSigstoreBundle0_verificationMaterial_tlogEntries_kindVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKindVersion(val.(SigstoreBundle0_verificationMaterial_tlogEntries_kindVersionable))
        }
        return nil
    }
    res["logId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSigstoreBundle0_verificationMaterial_tlogEntries_logIdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogId(val.(SigstoreBundle0_verificationMaterial_tlogEntries_logIdable))
        }
        return nil
    }
    res["logIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogIndex(val)
        }
        return nil
    }
    return res
}
// GetInclusionPromise gets the inclusionPromise property value. The inclusionPromise property
// returns a SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromiseable when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) GetInclusionPromise()(SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromiseable) {
    return m.inclusionPromise
}
// GetInclusionProof gets the inclusionProof property value. The inclusionProof property
// returns a *string when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) GetInclusionProof()(*string) {
    return m.inclusionProof
}
// GetIntegratedTime gets the integratedTime property value. The integratedTime property
// returns a *string when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) GetIntegratedTime()(*string) {
    return m.integratedTime
}
// GetKindVersion gets the kindVersion property value. The kindVersion property
// returns a SigstoreBundle0_verificationMaterial_tlogEntries_kindVersionable when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) GetKindVersion()(SigstoreBundle0_verificationMaterial_tlogEntries_kindVersionable) {
    return m.kindVersion
}
// GetLogId gets the logId property value. The logId property
// returns a SigstoreBundle0_verificationMaterial_tlogEntries_logIdable when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) GetLogId()(SigstoreBundle0_verificationMaterial_tlogEntries_logIdable) {
    return m.logId
}
// GetLogIndex gets the logIndex property value. The logIndex property
// returns a *string when successful
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) GetLogIndex()(*string) {
    return m.logIndex
}
// Serialize serializes information the current object
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("canonicalizedBody", m.GetCanonicalizedBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("inclusionPromise", m.GetInclusionPromise())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("inclusionProof", m.GetInclusionProof())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("integratedTime", m.GetIntegratedTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("kindVersion", m.GetKindVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("logId", m.GetLogId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("logIndex", m.GetLogIndex())
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
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCanonicalizedBody sets the canonicalizedBody property value. The canonicalizedBody property
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) SetCanonicalizedBody(value *string)() {
    m.canonicalizedBody = value
}
// SetInclusionPromise sets the inclusionPromise property value. The inclusionPromise property
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) SetInclusionPromise(value SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromiseable)() {
    m.inclusionPromise = value
}
// SetInclusionProof sets the inclusionProof property value. The inclusionProof property
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) SetInclusionProof(value *string)() {
    m.inclusionProof = value
}
// SetIntegratedTime sets the integratedTime property value. The integratedTime property
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) SetIntegratedTime(value *string)() {
    m.integratedTime = value
}
// SetKindVersion sets the kindVersion property value. The kindVersion property
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) SetKindVersion(value SigstoreBundle0_verificationMaterial_tlogEntries_kindVersionable)() {
    m.kindVersion = value
}
// SetLogId sets the logId property value. The logId property
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) SetLogId(value SigstoreBundle0_verificationMaterial_tlogEntries_logIdable)() {
    m.logId = value
}
// SetLogIndex sets the logIndex property value. The logIndex property
func (m *SigstoreBundle0_verificationMaterial_tlogEntries) SetLogIndex(value *string)() {
    m.logIndex = value
}
type SigstoreBundle0_verificationMaterial_tlogEntriesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCanonicalizedBody()(*string)
    GetInclusionPromise()(SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromiseable)
    GetInclusionProof()(*string)
    GetIntegratedTime()(*string)
    GetKindVersion()(SigstoreBundle0_verificationMaterial_tlogEntries_kindVersionable)
    GetLogId()(SigstoreBundle0_verificationMaterial_tlogEntries_logIdable)
    GetLogIndex()(*string)
    SetCanonicalizedBody(value *string)()
    SetInclusionPromise(value SigstoreBundle0_verificationMaterial_tlogEntries_inclusionPromiseable)()
    SetInclusionProof(value *string)()
    SetIntegratedTime(value *string)()
    SetKindVersion(value SigstoreBundle0_verificationMaterial_tlogEntries_kindVersionable)()
    SetLogId(value SigstoreBundle0_verificationMaterial_tlogEntries_logIdable)()
    SetLogIndex(value *string)()
}
