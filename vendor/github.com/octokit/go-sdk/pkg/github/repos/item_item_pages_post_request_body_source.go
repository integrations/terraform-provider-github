package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemPagesPostRequestBody_source the source branch and directory used to publish your Pages site.
type ItemItemPagesPostRequestBody_source struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The repository branch used to publish your site's source files.
    branch *string
}
// NewItemItemPagesPostRequestBody_source instantiates a new ItemItemPagesPostRequestBody_source and sets the default values.
func NewItemItemPagesPostRequestBody_source()(*ItemItemPagesPostRequestBody_source) {
    m := &ItemItemPagesPostRequestBody_source{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemPagesPostRequestBody_sourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemPagesPostRequestBody_sourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPagesPostRequestBody_source(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemPagesPostRequestBody_source) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBranch gets the branch property value. The repository branch used to publish your site's source files.
// returns a *string when successful
func (m *ItemItemPagesPostRequestBody_source) GetBranch()(*string) {
    return m.branch
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemPagesPostRequestBody_source) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranch(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemPagesPostRequestBody_source) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("branch", m.GetBranch())
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
func (m *ItemItemPagesPostRequestBody_source) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBranch sets the branch property value. The repository branch used to publish your site's source files.
func (m *ItemItemPagesPostRequestBody_source) SetBranch(value *string)() {
    m.branch = value
}
type ItemItemPagesPostRequestBody_sourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBranch()(*string)
    SetBranch(value *string)()
}
