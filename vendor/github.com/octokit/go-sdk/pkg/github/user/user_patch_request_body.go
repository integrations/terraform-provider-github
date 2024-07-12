package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UserPatchRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The new short biography of the user.
    bio *string
    // The new blog URL of the user.
    blog *string
    // The new company of the user.
    company *string
    // The publicly visible email address of the user.
    email *string
    // The new hiring availability of the user.
    hireable *bool
    // The new location of the user.
    location *string
    // The new name of the user.
    name *string
    // The new Twitter username of the user.
    twitter_username *string
}
// NewUserPatchRequestBody instantiates a new UserPatchRequestBody and sets the default values.
func NewUserPatchRequestBody()(*UserPatchRequestBody) {
    m := &UserPatchRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserPatchRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UserPatchRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBio gets the bio property value. The new short biography of the user.
// returns a *string when successful
func (m *UserPatchRequestBody) GetBio()(*string) {
    return m.bio
}
// GetBlog gets the blog property value. The new blog URL of the user.
// returns a *string when successful
func (m *UserPatchRequestBody) GetBlog()(*string) {
    return m.blog
}
// GetCompany gets the company property value. The new company of the user.
// returns a *string when successful
func (m *UserPatchRequestBody) GetCompany()(*string) {
    return m.company
}
// GetEmail gets the email property value. The publicly visible email address of the user.
// returns a *string when successful
func (m *UserPatchRequestBody) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserPatchRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBio(val)
        }
        return nil
    }
    res["blog"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlog(val)
        }
        return nil
    }
    res["company"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompany(val)
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["hireable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHireable(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
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
    res["twitter_username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTwitterUsername(val)
        }
        return nil
    }
    return res
}
// GetHireable gets the hireable property value. The new hiring availability of the user.
// returns a *bool when successful
func (m *UserPatchRequestBody) GetHireable()(*bool) {
    return m.hireable
}
// GetLocation gets the location property value. The new location of the user.
// returns a *string when successful
func (m *UserPatchRequestBody) GetLocation()(*string) {
    return m.location
}
// GetName gets the name property value. The new name of the user.
// returns a *string when successful
func (m *UserPatchRequestBody) GetName()(*string) {
    return m.name
}
// GetTwitterUsername gets the twitter_username property value. The new Twitter username of the user.
// returns a *string when successful
func (m *UserPatchRequestBody) GetTwitterUsername()(*string) {
    return m.twitter_username
}
// Serialize serializes information the current object
func (m *UserPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("bio", m.GetBio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("blog", m.GetBlog())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("company", m.GetCompany())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hireable", m.GetHireable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("location", m.GetLocation())
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
        err := writer.WriteStringValue("twitter_username", m.GetTwitterUsername())
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
func (m *UserPatchRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBio sets the bio property value. The new short biography of the user.
func (m *UserPatchRequestBody) SetBio(value *string)() {
    m.bio = value
}
// SetBlog sets the blog property value. The new blog URL of the user.
func (m *UserPatchRequestBody) SetBlog(value *string)() {
    m.blog = value
}
// SetCompany sets the company property value. The new company of the user.
func (m *UserPatchRequestBody) SetCompany(value *string)() {
    m.company = value
}
// SetEmail sets the email property value. The publicly visible email address of the user.
func (m *UserPatchRequestBody) SetEmail(value *string)() {
    m.email = value
}
// SetHireable sets the hireable property value. The new hiring availability of the user.
func (m *UserPatchRequestBody) SetHireable(value *bool)() {
    m.hireable = value
}
// SetLocation sets the location property value. The new location of the user.
func (m *UserPatchRequestBody) SetLocation(value *string)() {
    m.location = value
}
// SetName sets the name property value. The new name of the user.
func (m *UserPatchRequestBody) SetName(value *string)() {
    m.name = value
}
// SetTwitterUsername sets the twitter_username property value. The new Twitter username of the user.
func (m *UserPatchRequestBody) SetTwitterUsername(value *string)() {
    m.twitter_username = value
}
type UserPatchRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBio()(*string)
    GetBlog()(*string)
    GetCompany()(*string)
    GetEmail()(*string)
    GetHireable()(*bool)
    GetLocation()(*string)
    GetName()(*string)
    GetTwitterUsername()(*string)
    SetBio(value *string)()
    SetBlog(value *string)()
    SetCompany(value *string)()
    SetEmail(value *string)()
    SetHireable(value *bool)()
    SetLocation(value *string)()
    SetName(value *string)()
    SetTwitterUsername(value *string)()
}
