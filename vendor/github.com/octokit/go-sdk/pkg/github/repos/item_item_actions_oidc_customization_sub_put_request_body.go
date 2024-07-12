package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsOidcCustomizationSubPutRequestBody actions OIDC subject customization for a repository
type ItemItemActionsOidcCustomizationSubPutRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
    include_claim_keys []string
    // Whether to use the default template or not. If `true`, the `include_claim_keys` field is ignored.
    use_default *bool
}
// NewItemItemActionsOidcCustomizationSubPutRequestBody instantiates a new ItemItemActionsOidcCustomizationSubPutRequestBody and sets the default values.
func NewItemItemActionsOidcCustomizationSubPutRequestBody()(*ItemItemActionsOidcCustomizationSubPutRequestBody) {
    m := &ItemItemActionsOidcCustomizationSubPutRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemActionsOidcCustomizationSubPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemActionsOidcCustomizationSubPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsOidcCustomizationSubPutRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemActionsOidcCustomizationSubPutRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemActionsOidcCustomizationSubPutRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["include_claim_keys"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetIncludeClaimKeys(res)
        }
        return nil
    }
    res["use_default"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseDefault(val)
        }
        return nil
    }
    return res
}
// GetIncludeClaimKeys gets the include_claim_keys property value. Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
// returns a []string when successful
func (m *ItemItemActionsOidcCustomizationSubPutRequestBody) GetIncludeClaimKeys()([]string) {
    return m.include_claim_keys
}
// GetUseDefault gets the use_default property value. Whether to use the default template or not. If `true`, the `include_claim_keys` field is ignored.
// returns a *bool when successful
func (m *ItemItemActionsOidcCustomizationSubPutRequestBody) GetUseDefault()(*bool) {
    return m.use_default
}
// Serialize serializes information the current object
func (m *ItemItemActionsOidcCustomizationSubPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIncludeClaimKeys() != nil {
        err := writer.WriteCollectionOfStringValues("include_claim_keys", m.GetIncludeClaimKeys())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("use_default", m.GetUseDefault())
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
func (m *ItemItemActionsOidcCustomizationSubPutRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIncludeClaimKeys sets the include_claim_keys property value. Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
func (m *ItemItemActionsOidcCustomizationSubPutRequestBody) SetIncludeClaimKeys(value []string)() {
    m.include_claim_keys = value
}
// SetUseDefault sets the use_default property value. Whether to use the default template or not. If `true`, the `include_claim_keys` field is ignored.
func (m *ItemItemActionsOidcCustomizationSubPutRequestBody) SetUseDefault(value *bool)() {
    m.use_default = value
}
type ItemItemActionsOidcCustomizationSubPutRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIncludeClaimKeys()([]string)
    GetUseDefault()(*bool)
    SetIncludeClaimKeys(value []string)()
    SetUseDefault(value *bool)()
}
