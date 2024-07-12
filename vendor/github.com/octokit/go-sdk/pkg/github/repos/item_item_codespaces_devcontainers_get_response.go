package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemCodespacesDevcontainersGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The devcontainers property
    devcontainers []ItemItemCodespacesDevcontainersGetResponse_devcontainersable
    // The total_count property
    total_count *int32
}
// NewItemItemCodespacesDevcontainersGetResponse instantiates a new ItemItemCodespacesDevcontainersGetResponse and sets the default values.
func NewItemItemCodespacesDevcontainersGetResponse()(*ItemItemCodespacesDevcontainersGetResponse) {
    m := &ItemItemCodespacesDevcontainersGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemCodespacesDevcontainersGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemCodespacesDevcontainersGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCodespacesDevcontainersGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemCodespacesDevcontainersGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDevcontainers gets the devcontainers property value. The devcontainers property
// returns a []ItemItemCodespacesDevcontainersGetResponse_devcontainersable when successful
func (m *ItemItemCodespacesDevcontainersGetResponse) GetDevcontainers()([]ItemItemCodespacesDevcontainersGetResponse_devcontainersable) {
    return m.devcontainers
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemCodespacesDevcontainersGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["devcontainers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemItemCodespacesDevcontainersGetResponse_devcontainersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemItemCodespacesDevcontainersGetResponse_devcontainersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ItemItemCodespacesDevcontainersGetResponse_devcontainersable)
                }
            }
            m.SetDevcontainers(res)
        }
        return nil
    }
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    return res
}
// GetTotalCount gets the total_count property value. The total_count property
// returns a *int32 when successful
func (m *ItemItemCodespacesDevcontainersGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// Serialize serializes information the current object
func (m *ItemItemCodespacesDevcontainersGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDevcontainers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDevcontainers()))
        for i, v := range m.GetDevcontainers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("devcontainers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
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
func (m *ItemItemCodespacesDevcontainersGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDevcontainers sets the devcontainers property value. The devcontainers property
func (m *ItemItemCodespacesDevcontainersGetResponse) SetDevcontainers(value []ItemItemCodespacesDevcontainersGetResponse_devcontainersable)() {
    m.devcontainers = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemItemCodespacesDevcontainersGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
type ItemItemCodespacesDevcontainersGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDevcontainers()([]ItemItemCodespacesDevcontainersGetResponse_devcontainersable)
    GetTotalCount()(*int32)
    SetDevcontainers(value []ItemItemCodespacesDevcontainersGetResponse_devcontainersable)()
    SetTotalCount(value *int32)()
}
