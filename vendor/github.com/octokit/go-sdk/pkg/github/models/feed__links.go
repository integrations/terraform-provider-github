package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Feed__links struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Hypermedia Link with Type
    current_user LinkWithTypeable
    // Hypermedia Link with Type
    current_user_actor LinkWithTypeable
    // Hypermedia Link with Type
    current_user_organization LinkWithTypeable
    // The current_user_organizations property
    current_user_organizations []LinkWithTypeable
    // Hypermedia Link with Type
    current_user_public LinkWithTypeable
    // Hypermedia Link with Type
    repository_discussions LinkWithTypeable
    // Hypermedia Link with Type
    repository_discussions_category LinkWithTypeable
    // Hypermedia Link with Type
    security_advisories LinkWithTypeable
    // Hypermedia Link with Type
    timeline LinkWithTypeable
    // Hypermedia Link with Type
    user LinkWithTypeable
}
// NewFeed__links instantiates a new Feed__links and sets the default values.
func NewFeed__links()(*Feed__links) {
    m := &Feed__links{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFeed__linksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFeed__linksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFeed__links(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Feed__links) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCurrentUser gets the current_user property value. Hypermedia Link with Type
// returns a LinkWithTypeable when successful
func (m *Feed__links) GetCurrentUser()(LinkWithTypeable) {
    return m.current_user
}
// GetCurrentUserActor gets the current_user_actor property value. Hypermedia Link with Type
// returns a LinkWithTypeable when successful
func (m *Feed__links) GetCurrentUserActor()(LinkWithTypeable) {
    return m.current_user_actor
}
// GetCurrentUserOrganization gets the current_user_organization property value. Hypermedia Link with Type
// returns a LinkWithTypeable when successful
func (m *Feed__links) GetCurrentUserOrganization()(LinkWithTypeable) {
    return m.current_user_organization
}
// GetCurrentUserOrganizations gets the current_user_organizations property value. The current_user_organizations property
// returns a []LinkWithTypeable when successful
func (m *Feed__links) GetCurrentUserOrganizations()([]LinkWithTypeable) {
    return m.current_user_organizations
}
// GetCurrentUserPublic gets the current_user_public property value. Hypermedia Link with Type
// returns a LinkWithTypeable when successful
func (m *Feed__links) GetCurrentUserPublic()(LinkWithTypeable) {
    return m.current_user_public
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Feed__links) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["current_user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkWithTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUser(val.(LinkWithTypeable))
        }
        return nil
    }
    res["current_user_actor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkWithTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUserActor(val.(LinkWithTypeable))
        }
        return nil
    }
    res["current_user_organization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkWithTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUserOrganization(val.(LinkWithTypeable))
        }
        return nil
    }
    res["current_user_organizations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLinkWithTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LinkWithTypeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LinkWithTypeable)
                }
            }
            m.SetCurrentUserOrganizations(res)
        }
        return nil
    }
    res["current_user_public"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkWithTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUserPublic(val.(LinkWithTypeable))
        }
        return nil
    }
    res["repository_discussions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkWithTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryDiscussions(val.(LinkWithTypeable))
        }
        return nil
    }
    res["repository_discussions_category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkWithTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryDiscussionsCategory(val.(LinkWithTypeable))
        }
        return nil
    }
    res["security_advisories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkWithTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityAdvisories(val.(LinkWithTypeable))
        }
        return nil
    }
    res["timeline"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkWithTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeline(val.(LinkWithTypeable))
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkWithTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(LinkWithTypeable))
        }
        return nil
    }
    return res
}
// GetRepositoryDiscussions gets the repository_discussions property value. Hypermedia Link with Type
// returns a LinkWithTypeable when successful
func (m *Feed__links) GetRepositoryDiscussions()(LinkWithTypeable) {
    return m.repository_discussions
}
// GetRepositoryDiscussionsCategory gets the repository_discussions_category property value. Hypermedia Link with Type
// returns a LinkWithTypeable when successful
func (m *Feed__links) GetRepositoryDiscussionsCategory()(LinkWithTypeable) {
    return m.repository_discussions_category
}
// GetSecurityAdvisories gets the security_advisories property value. Hypermedia Link with Type
// returns a LinkWithTypeable when successful
func (m *Feed__links) GetSecurityAdvisories()(LinkWithTypeable) {
    return m.security_advisories
}
// GetTimeline gets the timeline property value. Hypermedia Link with Type
// returns a LinkWithTypeable when successful
func (m *Feed__links) GetTimeline()(LinkWithTypeable) {
    return m.timeline
}
// GetUser gets the user property value. Hypermedia Link with Type
// returns a LinkWithTypeable when successful
func (m *Feed__links) GetUser()(LinkWithTypeable) {
    return m.user
}
// Serialize serializes information the current object
func (m *Feed__links) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("current_user", m.GetCurrentUser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("current_user_actor", m.GetCurrentUserActor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("current_user_organization", m.GetCurrentUserOrganization())
        if err != nil {
            return err
        }
    }
    if m.GetCurrentUserOrganizations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCurrentUserOrganizations()))
        for i, v := range m.GetCurrentUserOrganizations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("current_user_organizations", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("current_user_public", m.GetCurrentUserPublic())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repository_discussions", m.GetRepositoryDiscussions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repository_discussions_category", m.GetRepositoryDiscussionsCategory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("security_advisories", m.GetSecurityAdvisories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("timeline", m.GetTimeline())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *Feed__links) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCurrentUser sets the current_user property value. Hypermedia Link with Type
func (m *Feed__links) SetCurrentUser(value LinkWithTypeable)() {
    m.current_user = value
}
// SetCurrentUserActor sets the current_user_actor property value. Hypermedia Link with Type
func (m *Feed__links) SetCurrentUserActor(value LinkWithTypeable)() {
    m.current_user_actor = value
}
// SetCurrentUserOrganization sets the current_user_organization property value. Hypermedia Link with Type
func (m *Feed__links) SetCurrentUserOrganization(value LinkWithTypeable)() {
    m.current_user_organization = value
}
// SetCurrentUserOrganizations sets the current_user_organizations property value. The current_user_organizations property
func (m *Feed__links) SetCurrentUserOrganizations(value []LinkWithTypeable)() {
    m.current_user_organizations = value
}
// SetCurrentUserPublic sets the current_user_public property value. Hypermedia Link with Type
func (m *Feed__links) SetCurrentUserPublic(value LinkWithTypeable)() {
    m.current_user_public = value
}
// SetRepositoryDiscussions sets the repository_discussions property value. Hypermedia Link with Type
func (m *Feed__links) SetRepositoryDiscussions(value LinkWithTypeable)() {
    m.repository_discussions = value
}
// SetRepositoryDiscussionsCategory sets the repository_discussions_category property value. Hypermedia Link with Type
func (m *Feed__links) SetRepositoryDiscussionsCategory(value LinkWithTypeable)() {
    m.repository_discussions_category = value
}
// SetSecurityAdvisories sets the security_advisories property value. Hypermedia Link with Type
func (m *Feed__links) SetSecurityAdvisories(value LinkWithTypeable)() {
    m.security_advisories = value
}
// SetTimeline sets the timeline property value. Hypermedia Link with Type
func (m *Feed__links) SetTimeline(value LinkWithTypeable)() {
    m.timeline = value
}
// SetUser sets the user property value. Hypermedia Link with Type
func (m *Feed__links) SetUser(value LinkWithTypeable)() {
    m.user = value
}
type Feed__linksable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentUser()(LinkWithTypeable)
    GetCurrentUserActor()(LinkWithTypeable)
    GetCurrentUserOrganization()(LinkWithTypeable)
    GetCurrentUserOrganizations()([]LinkWithTypeable)
    GetCurrentUserPublic()(LinkWithTypeable)
    GetRepositoryDiscussions()(LinkWithTypeable)
    GetRepositoryDiscussionsCategory()(LinkWithTypeable)
    GetSecurityAdvisories()(LinkWithTypeable)
    GetTimeline()(LinkWithTypeable)
    GetUser()(LinkWithTypeable)
    SetCurrentUser(value LinkWithTypeable)()
    SetCurrentUserActor(value LinkWithTypeable)()
    SetCurrentUserOrganization(value LinkWithTypeable)()
    SetCurrentUserOrganizations(value []LinkWithTypeable)()
    SetCurrentUserPublic(value LinkWithTypeable)()
    SetRepositoryDiscussions(value LinkWithTypeable)()
    SetRepositoryDiscussionsCategory(value LinkWithTypeable)()
    SetSecurityAdvisories(value LinkWithTypeable)()
    SetTimeline(value LinkWithTypeable)()
    SetUser(value LinkWithTypeable)()
}
