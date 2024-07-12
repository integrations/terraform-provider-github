package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemGeneratePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A short description of the new repository.
    description *string
    // Set to `true` to include the directory structure and files from all branches in the template repository, and not just the default branch. Default: `false`.
    include_all_branches *bool
    // The name of the new repository.
    name *string
    // The organization or person who will own the new repository. To create a new repository in an organization, the authenticated user must be a member of the specified organization.
    owner *string
    // Either `true` to create a new private repository or `false` to create a new public one.
    private *bool
}
// NewItemItemGeneratePostRequestBody instantiates a new ItemItemGeneratePostRequestBody and sets the default values.
func NewItemItemGeneratePostRequestBody()(*ItemItemGeneratePostRequestBody) {
    m := &ItemItemGeneratePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemGeneratePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemGeneratePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemGeneratePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemGeneratePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. A short description of the new repository.
// returns a *string when successful
func (m *ItemItemGeneratePostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemGeneratePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["include_all_branches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeAllBranches(val)
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
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
        }
        return nil
    }
    res["private"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivate(val)
        }
        return nil
    }
    return res
}
// GetIncludeAllBranches gets the include_all_branches property value. Set to `true` to include the directory structure and files from all branches in the template repository, and not just the default branch. Default: `false`.
// returns a *bool when successful
func (m *ItemItemGeneratePostRequestBody) GetIncludeAllBranches()(*bool) {
    return m.include_all_branches
}
// GetName gets the name property value. The name of the new repository.
// returns a *string when successful
func (m *ItemItemGeneratePostRequestBody) GetName()(*string) {
    return m.name
}
// GetOwner gets the owner property value. The organization or person who will own the new repository. To create a new repository in an organization, the authenticated user must be a member of the specified organization.
// returns a *string when successful
func (m *ItemItemGeneratePostRequestBody) GetOwner()(*string) {
    return m.owner
}
// GetPrivate gets the private property value. Either `true` to create a new private repository or `false` to create a new public one.
// returns a *bool when successful
func (m *ItemItemGeneratePostRequestBody) GetPrivate()(*bool) {
    return m.private
}
// Serialize serializes information the current object
func (m *ItemItemGeneratePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("include_all_branches", m.GetIncludeAllBranches())
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
        err := writer.WriteStringValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("private", m.GetPrivate())
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
func (m *ItemItemGeneratePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. A short description of the new repository.
func (m *ItemItemGeneratePostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetIncludeAllBranches sets the include_all_branches property value. Set to `true` to include the directory structure and files from all branches in the template repository, and not just the default branch. Default: `false`.
func (m *ItemItemGeneratePostRequestBody) SetIncludeAllBranches(value *bool)() {
    m.include_all_branches = value
}
// SetName sets the name property value. The name of the new repository.
func (m *ItemItemGeneratePostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetOwner sets the owner property value. The organization or person who will own the new repository. To create a new repository in an organization, the authenticated user must be a member of the specified organization.
func (m *ItemItemGeneratePostRequestBody) SetOwner(value *string)() {
    m.owner = value
}
// SetPrivate sets the private property value. Either `true` to create a new private repository or `false` to create a new public one.
func (m *ItemItemGeneratePostRequestBody) SetPrivate(value *bool)() {
    m.private = value
}
type ItemItemGeneratePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetIncludeAllBranches()(*bool)
    GetName()(*string)
    GetOwner()(*string)
    GetPrivate()(*bool)
    SetDescription(value *string)()
    SetIncludeAllBranches(value *bool)()
    SetName(value *string)()
    SetOwner(value *string)()
    SetPrivate(value *bool)()
}
