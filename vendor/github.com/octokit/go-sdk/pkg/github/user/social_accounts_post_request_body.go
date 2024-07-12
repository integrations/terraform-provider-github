package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Social_accountsPostRequestBody struct {
    // Full URLs for the social media profiles to add.
    account_urls []string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewSocial_accountsPostRequestBody instantiates a new Social_accountsPostRequestBody and sets the default values.
func NewSocial_accountsPostRequestBody()(*Social_accountsPostRequestBody) {
    m := &Social_accountsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSocial_accountsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSocial_accountsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSocial_accountsPostRequestBody(), nil
}
// GetAccountUrls gets the account_urls property value. Full URLs for the social media profiles to add.
// returns a []string when successful
func (m *Social_accountsPostRequestBody) GetAccountUrls()([]string) {
    return m.account_urls
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Social_accountsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Social_accountsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["account_urls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAccountUrls(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Social_accountsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccountUrls() != nil {
        err := writer.WriteCollectionOfStringValues("account_urls", m.GetAccountUrls())
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
// SetAccountUrls sets the account_urls property value. Full URLs for the social media profiles to add.
func (m *Social_accountsPostRequestBody) SetAccountUrls(value []string)() {
    m.account_urls = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Social_accountsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type Social_accountsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountUrls()([]string)
    SetAccountUrls(value []string)()
}
