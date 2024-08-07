package orgs

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemCodeSecurityConfigurationsItemAttachPostRequestBody struct {
    // An array of repository IDs to attach the configuration to. You can only provide a list of repository ids when the `scope` is set to `selected`.
    selected_repository_ids []int32
}
// NewItemCodeSecurityConfigurationsItemAttachPostRequestBody instantiates a new ItemCodeSecurityConfigurationsItemAttachPostRequestBody and sets the default values.
func NewItemCodeSecurityConfigurationsItemAttachPostRequestBody()(*ItemCodeSecurityConfigurationsItemAttachPostRequestBody) {
    m := &ItemCodeSecurityConfigurationsItemAttachPostRequestBody{
    }
    return m
}
// CreateItemCodeSecurityConfigurationsItemAttachPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCodeSecurityConfigurationsItemAttachPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCodeSecurityConfigurationsItemAttachPostRequestBody(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemCodeSecurityConfigurationsItemAttachPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["selected_repository_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*int32))
                }
            }
            m.SetSelectedRepositoryIds(res)
        }
        return nil
    }
    return res
}
// GetSelectedRepositoryIds gets the selected_repository_ids property value. An array of repository IDs to attach the configuration to. You can only provide a list of repository ids when the `scope` is set to `selected`.
// returns a []int32 when successful
func (m *ItemCodeSecurityConfigurationsItemAttachPostRequestBody) GetSelectedRepositoryIds()([]int32) {
    return m.selected_repository_ids
}
// Serialize serializes information the current object
func (m *ItemCodeSecurityConfigurationsItemAttachPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSelectedRepositoryIds() != nil {
        err := writer.WriteCollectionOfInt32Values("selected_repository_ids", m.GetSelectedRepositoryIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSelectedRepositoryIds sets the selected_repository_ids property value. An array of repository IDs to attach the configuration to. You can only provide a list of repository ids when the `scope` is set to `selected`.
func (m *ItemCodeSecurityConfigurationsItemAttachPostRequestBody) SetSelectedRepositoryIds(value []int32)() {
    m.selected_repository_ids = value
}
type ItemCodeSecurityConfigurationsItemAttachPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSelectedRepositoryIds()([]int32)
    SetSelectedRepositoryIds(value []int32)()
}
