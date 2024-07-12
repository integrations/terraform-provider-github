package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GpgKey_subkeys struct {
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
    created_at *string
    // The emails property
    emails []GpgKey_subkeys_emailsable
    // The expires_at property
    expires_at *string
    // The id property
    id *int64
    // The key_id property
    key_id *string
    // The primary_key_id property
    primary_key_id *int32
    // The public_key property
    public_key *string
    // The raw_key property
    raw_key *string
    // The revoked property
    revoked *bool
    // The subkeys property
    subkeys i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable
}
// NewGpgKey_subkeys instantiates a new GpgKey_subkeys and sets the default values.
func NewGpgKey_subkeys()(*GpgKey_subkeys) {
    m := &GpgKey_subkeys{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGpgKey_subkeysFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGpgKey_subkeysFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGpgKey_subkeys(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GpgKey_subkeys) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCanCertify gets the can_certify property value. The can_certify property
// returns a *bool when successful
func (m *GpgKey_subkeys) GetCanCertify()(*bool) {
    return m.can_certify
}
// GetCanEncryptComms gets the can_encrypt_comms property value. The can_encrypt_comms property
// returns a *bool when successful
func (m *GpgKey_subkeys) GetCanEncryptComms()(*bool) {
    return m.can_encrypt_comms
}
// GetCanEncryptStorage gets the can_encrypt_storage property value. The can_encrypt_storage property
// returns a *bool when successful
func (m *GpgKey_subkeys) GetCanEncryptStorage()(*bool) {
    return m.can_encrypt_storage
}
// GetCanSign gets the can_sign property value. The can_sign property
// returns a *bool when successful
func (m *GpgKey_subkeys) GetCanSign()(*bool) {
    return m.can_sign
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *string when successful
func (m *GpgKey_subkeys) GetCreatedAt()(*string) {
    return m.created_at
}
// GetEmails gets the emails property value. The emails property
// returns a []GpgKey_subkeys_emailsable when successful
func (m *GpgKey_subkeys) GetEmails()([]GpgKey_subkeys_emailsable) {
    return m.emails
}
// GetExpiresAt gets the expires_at property value. The expires_at property
// returns a *string when successful
func (m *GpgKey_subkeys) GetExpiresAt()(*string) {
    return m.expires_at
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GpgKey_subkeys) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["emails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGpgKey_subkeys_emailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GpgKey_subkeys_emailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GpgKey_subkeys_emailsable)
                }
            }
            m.SetEmails(res)
        }
        return nil
    }
    res["expires_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
        val, err := n.GetObjectValue(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.CreateUntypedNodeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubkeys(val.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *GpgKey_subkeys) GetId()(*int64) {
    return m.id
}
// GetKeyId gets the key_id property value. The key_id property
// returns a *string when successful
func (m *GpgKey_subkeys) GetKeyId()(*string) {
    return m.key_id
}
// GetPrimaryKeyId gets the primary_key_id property value. The primary_key_id property
// returns a *int32 when successful
func (m *GpgKey_subkeys) GetPrimaryKeyId()(*int32) {
    return m.primary_key_id
}
// GetPublicKey gets the public_key property value. The public_key property
// returns a *string when successful
func (m *GpgKey_subkeys) GetPublicKey()(*string) {
    return m.public_key
}
// GetRawKey gets the raw_key property value. The raw_key property
// returns a *string when successful
func (m *GpgKey_subkeys) GetRawKey()(*string) {
    return m.raw_key
}
// GetRevoked gets the revoked property value. The revoked property
// returns a *bool when successful
func (m *GpgKey_subkeys) GetRevoked()(*bool) {
    return m.revoked
}
// GetSubkeys gets the subkeys property value. The subkeys property
// returns a UntypedNodeable when successful
func (m *GpgKey_subkeys) GetSubkeys()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable) {
    return m.subkeys
}
// Serialize serializes information the current object
func (m *GpgKey_subkeys) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("created_at", m.GetCreatedAt())
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
        err := writer.WriteStringValue("expires_at", m.GetExpiresAt())
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
    {
        err := writer.WriteObjectValue("subkeys", m.GetSubkeys())
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
func (m *GpgKey_subkeys) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCanCertify sets the can_certify property value. The can_certify property
func (m *GpgKey_subkeys) SetCanCertify(value *bool)() {
    m.can_certify = value
}
// SetCanEncryptComms sets the can_encrypt_comms property value. The can_encrypt_comms property
func (m *GpgKey_subkeys) SetCanEncryptComms(value *bool)() {
    m.can_encrypt_comms = value
}
// SetCanEncryptStorage sets the can_encrypt_storage property value. The can_encrypt_storage property
func (m *GpgKey_subkeys) SetCanEncryptStorage(value *bool)() {
    m.can_encrypt_storage = value
}
// SetCanSign sets the can_sign property value. The can_sign property
func (m *GpgKey_subkeys) SetCanSign(value *bool)() {
    m.can_sign = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *GpgKey_subkeys) SetCreatedAt(value *string)() {
    m.created_at = value
}
// SetEmails sets the emails property value. The emails property
func (m *GpgKey_subkeys) SetEmails(value []GpgKey_subkeys_emailsable)() {
    m.emails = value
}
// SetExpiresAt sets the expires_at property value. The expires_at property
func (m *GpgKey_subkeys) SetExpiresAt(value *string)() {
    m.expires_at = value
}
// SetId sets the id property value. The id property
func (m *GpgKey_subkeys) SetId(value *int64)() {
    m.id = value
}
// SetKeyId sets the key_id property value. The key_id property
func (m *GpgKey_subkeys) SetKeyId(value *string)() {
    m.key_id = value
}
// SetPrimaryKeyId sets the primary_key_id property value. The primary_key_id property
func (m *GpgKey_subkeys) SetPrimaryKeyId(value *int32)() {
    m.primary_key_id = value
}
// SetPublicKey sets the public_key property value. The public_key property
func (m *GpgKey_subkeys) SetPublicKey(value *string)() {
    m.public_key = value
}
// SetRawKey sets the raw_key property value. The raw_key property
func (m *GpgKey_subkeys) SetRawKey(value *string)() {
    m.raw_key = value
}
// SetRevoked sets the revoked property value. The revoked property
func (m *GpgKey_subkeys) SetRevoked(value *bool)() {
    m.revoked = value
}
// SetSubkeys sets the subkeys property value. The subkeys property
func (m *GpgKey_subkeys) SetSubkeys(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)() {
    m.subkeys = value
}
type GpgKey_subkeysable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCanCertify()(*bool)
    GetCanEncryptComms()(*bool)
    GetCanEncryptStorage()(*bool)
    GetCanSign()(*bool)
    GetCreatedAt()(*string)
    GetEmails()([]GpgKey_subkeys_emailsable)
    GetExpiresAt()(*string)
    GetId()(*int64)
    GetKeyId()(*string)
    GetPrimaryKeyId()(*int32)
    GetPublicKey()(*string)
    GetRawKey()(*string)
    GetRevoked()(*bool)
    GetSubkeys()(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)
    SetCanCertify(value *bool)()
    SetCanEncryptComms(value *bool)()
    SetCanEncryptStorage(value *bool)()
    SetCanSign(value *bool)()
    SetCreatedAt(value *string)()
    SetEmails(value []GpgKey_subkeys_emailsable)()
    SetExpiresAt(value *string)()
    SetId(value *int64)()
    SetKeyId(value *string)()
    SetPrimaryKeyId(value *int32)()
    SetPublicKey(value *string)()
    SetRawKey(value *string)()
    SetRevoked(value *bool)()
    SetSubkeys(value i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.UntypedNodeable)()
}
