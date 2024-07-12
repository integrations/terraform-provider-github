package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the status checks
    contexts []string
}
// NewItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1 instantiates a new ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1 and sets the default values.
func NewItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1()(*ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1) {
    m := &ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContexts gets the contexts property value. The name of the status checks
// returns a []string when successful
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1) GetContexts()([]string) {
    return m.contexts
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contexts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetContexts(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContexts() != nil {
        err := writer.WriteCollectionOfStringValues("contexts", m.GetContexts())
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
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContexts sets the contexts property value. The name of the status checks
func (m *ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1) SetContexts(value []string)() {
    m.contexts = value
}
type ItemItemBranchesItemProtectionRequired_status_checksContextsPostRequestBodyMember1able interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContexts()([]string)
    SetContexts(value []string)()
}
