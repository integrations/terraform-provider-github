package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GpgKey a unique encryption key
type GpgKey struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The can_certify property
    can_certify *bool
    // The can_encrypt_comms property
    can_encrypt_comms *bool
    // The can_encrypt_storage property
    can_encrypt_storage *bool
    // The can_sign property
    can_sign *bool
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The emails property
    emails []GpgKey_emailsable
    // The expires_at property
    expires_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The id property
    id *int64
    // The key_id property
    key_id *string
    // The name property
    name *string
    // The primary_key_id property
    primary_key_id *int32
    // The public_key property
    public_key *string
    // The raw_key property
    raw_key *string
    // The revoked property
    revoked *bool
    // The subkeys property
    subkeys []GpgKey_subkeysable
}
// NewGpgKey instantiates a new GpgKey and sets the default values.
func NewGpgKey()(*GpgKey) {
    m := &GpgKey{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGpgKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGpgKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGpgKey(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GpgKey) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCanCertify gets the can_certify property value. The can_certify property
// returns a *bool when successful
func (m *GpgKey) GetCanCertify()(*bool) {
    return m.can_certify
}
// GetCanEncryptComms gets the can_encrypt_comms property value. The can_encrypt_comms property
// returns a *bool when successful
func (m *GpgKey) GetCanEncryptComms()(*bool) {
    return m.can_encrypt_comms
}
// GetCanEncryptStorage gets the can_encrypt_storage property value. The can_encrypt_storage property
// returns a *bool when successful
func (m *GpgKey) GetCanEncryptStorage()(*bool) {
    return m.can_encrypt_storage
}
// GetCanSign gets the can_sign property value. The can_sign property
// returns a *bool when successful
func (m *GpgKey) GetCanSign()(*bool) {
    return m.can_sign
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *GpgKey) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetEmails gets the emails property value. The emails property
// returns a []GpgKey_emailsable when successful
func (m *GpgKey) GetEmails()([]GpgKey_emailsable) {
    return m.emails
}
// GetExpiresAt gets the expires_at property value. The expires_at property
// returns a *Time when successful
func (m *GpgKey) GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expires_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GpgKey) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["can_certify"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanCertify(val)
        }
        return nil
    }
    res["can_encrypt_comms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanEncryptComms(val)
        }
        return nil
    }
    res["can_encrypt_storage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanEncryptStorage(val)
        }
        return nil
    }
    res["can_sign"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanSign(val)
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["emails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGpgKey_emailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GpgKey_emailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GpgKey_emailsable)
                }
            }
            m.SetEmails(res)
        }
        return nil
    }
    res["expires_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiresAt(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["key_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
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
    res["primary_key_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryKeyId(val)
        }
        return nil
    }
    res["public_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicKey(val)
        }
        return nil
    }
    res["raw_key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRawKey(val)
        }
        return nil
    }
    res["revoked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRevoked(val)
        }
        return nil
    }
    res["subkeys"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGpgKey_subkeysFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GpgKey_subkeysable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GpgKey_subkeysable)
                }
            }
            m.SetSubkeys(res)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *GpgKey) GetId()(*int64) {
    return m.id
}
// GetKeyId gets the key_id property value. The key_id property
// returns a *string when successful
func (m *GpgKey) GetKeyId()(*string) {
    return m.key_id
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *GpgKey) GetName()(*string) {
    return m.name
}
// GetPrimaryKeyId gets the primary_key_id property value. The primary_key_id property
// returns a *int32 when successful
func (m *GpgKey) GetPrimaryKeyId()(*int32) {
    return m.primary_key_id
}
// GetPublicKey gets the public_key property value. The public_key property
// returns a *string when successful
func (m *GpgKey) GetPublicKey()(*string) {
    return m.public_key
}
// GetRawKey gets the raw_key property value. The raw_key property
// returns a *string when successful
func (m *GpgKey) GetRawKey()(*string) {
    return m.raw_key
}
// GetRevoked gets the revoked property value. The revoked property
// returns a *bool when successful
func (m *GpgKey) GetRevoked()(*bool) {
    return m.revoked
}
// GetSubkeys gets the subkeys property value. The subkeys property
// returns a []GpgKey_subkeysable when successful
func (m *GpgKey) GetSubkeys()([]GpgKey_subkeysable) {
    return m.subkeys
}
// Serialize serializes information the current object
func (m *GpgKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("can_certify", m.GetCanCertify())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("can_encrypt_comms", m.GetCanEncryptComms())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("can_encrypt_storage", m.GetCanEncryptStorage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("can_sign", m.GetCanSign())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    if m.GetEmails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmails()))
        for i, v := range m.GetEmails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("emails", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("expires_at", m.GetExpiresAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("key_id", m.GetKeyId())
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
        err := writer.WriteInt32Value("primary_key_id", m.GetPrimaryKeyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("public_key", m.GetPublicKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("raw_key", m.GetRawKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("revoked", m.GetRevoked())
        if err != nil {
            return err
        }
    }
    if m.GetSubkeys() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubkeys()))
        for i, v := range m.GetSubkeys() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("subkeys", cast)
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
func (m *GpgKey) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCanCertify sets the can_certify property value. The can_certify property
func (m *GpgKey) SetCanCertify(value *bool)() {
    m.can_certify = value
}
// SetCanEncryptComms sets the can_encrypt_comms property value. The can_encrypt_comms property
func (m *GpgKey) SetCanEncryptComms(value *bool)() {
    m.can_encrypt_comms = value
}
// SetCanEncryptStorage sets the can_encrypt_storage property value. The can_encrypt_storage property
func (m *GpgKey) SetCanEncryptStorage(value *bool)() {
    m.can_encrypt_storage = value
}
// SetCanSign sets the can_sign property value. The can_sign property
func (m *GpgKey) SetCanSign(value *bool)() {
    m.can_sign = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *GpgKey) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetEmails sets the emails property value. The emails property
func (m *GpgKey) SetEmails(value []GpgKey_emailsable)() {
    m.emails = value
}
// SetExpiresAt sets the expires_at property value. The expires_at property
func (m *GpgKey) SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expires_at = value
}
// SetId sets the id property value. The id property
func (m *GpgKey) SetId(value *int64)() {
    m.id = value
}
// SetKeyId sets the key_id property value. The key_id property
func (m *GpgKey) SetKeyId(value *string)() {
    m.key_id = value
}
// SetName sets the name property value. The name property
func (m *GpgKey) SetName(value *string)() {
    m.name = value
}
// SetPrimaryKeyId sets the primary_key_id property value. The primary_key_id property
func (m *GpgKey) SetPrimaryKeyId(value *int32)() {
    m.primary_key_id = value
}
// SetPublicKey sets the public_key property value. The public_key property
func (m *GpgKey) SetPublicKey(value *string)() {
    m.public_key = value
}
// SetRawKey sets the raw_key property value. The raw_key property
func (m *GpgKey) SetRawKey(value *string)() {
    m.raw_key = value
}
// SetRevoked sets the revoked property value. The revoked property
func (m *GpgKey) SetRevoked(value *bool)() {
    m.revoked = value
}
// SetSubkeys sets the subkeys property value. The subkeys property
func (m *GpgKey) SetSubkeys(value []GpgKey_subkeysable)() {
    m.subkeys = value
}
type GpgKeyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCanCertify()(*bool)
    GetCanEncryptComms()(*bool)
    GetCanEncryptStorage()(*bool)
    GetCanSign()(*bool)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEmails()([]GpgKey_emailsable)
    GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*int64)
    GetKeyId()(*string)
    GetName()(*string)
    GetPrimaryKeyId()(*int32)
    GetPublicKey()(*string)
    GetRawKey()(*string)
    GetRevoked()(*bool)
    GetSubkeys()([]GpgKey_subkeysable)
    SetCanCertify(value *bool)()
    SetCanEncryptComms(value *bool)()
    SetCanEncryptStorage(value *bool)()
    SetCanSign(value *bool)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEmails(value []GpgKey_emailsable)()
    SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *int64)()
    SetKeyId(value *string)()
    SetName(value *string)()
    SetPrimaryKeyId(value *int32)()
    SetPublicKey(value *string)()
    SetRawKey(value *string)()
    SetRevoked(value *bool)()
    SetSubkeys(value []GpgKey_subkeysable)()
}
