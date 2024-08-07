package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Feed feed
type Feed struct {
    // The _links property
    _links Feed__linksable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The current_user_actor_url property
    current_user_actor_url *string
    // The current_user_organization_url property
    current_user_organization_url *string
    // The current_user_organization_urls property
    current_user_organization_urls []string
    // The current_user_public_url property
    current_user_public_url *string
    // The current_user_url property
    current_user_url *string
    // A feed of discussions for a given repository and category.
    repository_discussions_category_url *string
    // A feed of discussions for a given repository.
    repository_discussions_url *string
    // The security_advisories_url property
    security_advisories_url *string
    // The timeline_url property
    timeline_url *string
    // The user_url property
    user_url *string
}
// NewFeed instantiates a new Feed and sets the default values.
func NewFeed()(*Feed) {
    m := &Feed{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFeedFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFeedFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFeed(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Feed) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCurrentUserActorUrl gets the current_user_actor_url property value. The current_user_actor_url property
// returns a *string when successful
func (m *Feed) GetCurrentUserActorUrl()(*string) {
    return m.current_user_actor_url
}
// GetCurrentUserOrganizationUrl gets the current_user_organization_url property value. The current_user_organization_url property
// returns a *string when successful
func (m *Feed) GetCurrentUserOrganizationUrl()(*string) {
    return m.current_user_organization_url
}
// GetCurrentUserOrganizationUrls gets the current_user_organization_urls property value. The current_user_organization_urls property
// returns a []string when successful
func (m *Feed) GetCurrentUserOrganizationUrls()([]string) {
    return m.current_user_organization_urls
}
// GetCurrentUserPublicUrl gets the current_user_public_url property value. The current_user_public_url property
// returns a *string when successful
func (m *Feed) GetCurrentUserPublicUrl()(*string) {
    return m.current_user_public_url
}
// GetCurrentUserUrl gets the current_user_url property value. The current_user_url property
// returns a *string when successful
func (m *Feed) GetCurrentUserUrl()(*string) {
    return m.current_user_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Feed) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["_links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFeed__linksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(Feed__linksable))
        }
        return nil
    }
    res["current_user_actor_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUserActorUrl(val)
        }
        return nil
    }
    res["current_user_organization_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUserOrganizationUrl(val)
        }
        return nil
    }
    res["current_user_organization_urls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetCurrentUserOrganizationUrls(res)
        }
        return nil
    }
    res["current_user_public_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUserPublicUrl(val)
        }
        return nil
    }
    res["current_user_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUserUrl(val)
        }
        return nil
    }
    res["repository_discussions_category_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryDiscussionsCategoryUrl(val)
        }
        return nil
    }
    res["repository_discussions_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryDiscussionsUrl(val)
        }
        return nil
    }
    res["security_advisories_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityAdvisoriesUrl(val)
        }
        return nil
    }
    res["timeline_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimelineUrl(val)
        }
        return nil
    }
    res["user_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserUrl(val)
        }
        return nil
    }
    return res
}
// GetLinks gets the _links property value. The _links property
// returns a Feed__linksable when successful
func (m *Feed) GetLinks()(Feed__linksable) {
    return m._links
}
// GetRepositoryDiscussionsCategoryUrl gets the repository_discussions_category_url property value. A feed of discussions for a given repository and category.
// returns a *string when successful
func (m *Feed) GetRepositoryDiscussionsCategoryUrl()(*string) {
    return m.repository_discussions_category_url
}
// GetRepositoryDiscussionsUrl gets the repository_discussions_url property value. A feed of discussions for a given repository.
// returns a *string when successful
func (m *Feed) GetRepositoryDiscussionsUrl()(*string) {
    return m.repository_discussions_url
}
// GetSecurityAdvisoriesUrl gets the security_advisories_url property value. The security_advisories_url property
// returns a *string when successful
func (m *Feed) GetSecurityAdvisoriesUrl()(*string) {
    return m.security_advisories_url
}
// GetTimelineUrl gets the timeline_url property value. The timeline_url property
// returns a *string when successful
func (m *Feed) GetTimelineUrl()(*string) {
    return m.timeline_url
}
// GetUserUrl gets the user_url property value. The user_url property
// returns a *string when successful
func (m *Feed) GetUserUrl()(*string) {
    return m.user_url
}
// Serialize serializes information the current object
func (m *Feed) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("current_user_actor_url", m.GetCurrentUserActorUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("current_user_organization_url", m.GetCurrentUserOrganizationUrl())
        if err != nil {
            return err
        }
    }
    if m.GetCurrentUserOrganizationUrls() != nil {
        err := writer.WriteCollectionOfStringValues("current_user_organization_urls", m.GetCurrentUserOrganizationUrls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("current_user_public_url", m.GetCurrentUserPublicUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("current_user_url", m.GetCurrentUserUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repository_discussions_category_url", m.GetRepositoryDiscussionsCategoryUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repository_discussions_url", m.GetRepositoryDiscussionsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("security_advisories_url", m.GetSecurityAdvisoriesUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeline_url", m.GetTimelineUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user_url", m.GetUserUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("_links", m.GetLinks())
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
func (m *Feed) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCurrentUserActorUrl sets the current_user_actor_url property value. The current_user_actor_url property
func (m *Feed) SetCurrentUserActorUrl(value *string)() {
    m.current_user_actor_url = value
}
// SetCurrentUserOrganizationUrl sets the current_user_organization_url property value. The current_user_organization_url property
func (m *Feed) SetCurrentUserOrganizationUrl(value *string)() {
    m.current_user_organization_url = value
}
// SetCurrentUserOrganizationUrls sets the current_user_organization_urls property value. The current_user_organization_urls property
func (m *Feed) SetCurrentUserOrganizationUrls(value []string)() {
    m.current_user_organization_urls = value
}
// SetCurrentUserPublicUrl sets the current_user_public_url property value. The current_user_public_url property
func (m *Feed) SetCurrentUserPublicUrl(value *string)() {
    m.current_user_public_url = value
}
// SetCurrentUserUrl sets the current_user_url property value. The current_user_url property
func (m *Feed) SetCurrentUserUrl(value *string)() {
    m.current_user_url = value
}
// SetLinks sets the _links property value. The _links property
func (m *Feed) SetLinks(value Feed__linksable)() {
    m._links = value
}
// SetRepositoryDiscussionsCategoryUrl sets the repository_discussions_category_url property value. A feed of discussions for a given repository and category.
func (m *Feed) SetRepositoryDiscussionsCategoryUrl(value *string)() {
    m.repository_discussions_category_url = value
}
// SetRepositoryDiscussionsUrl sets the repository_discussions_url property value. A feed of discussions for a given repository.
func (m *Feed) SetRepositoryDiscussionsUrl(value *string)() {
    m.repository_discussions_url = value
}
// SetSecurityAdvisoriesUrl sets the security_advisories_url property value. The security_advisories_url property
func (m *Feed) SetSecurityAdvisoriesUrl(value *string)() {
    m.security_advisories_url = value
}
// SetTimelineUrl sets the timeline_url property value. The timeline_url property
func (m *Feed) SetTimelineUrl(value *string)() {
    m.timeline_url = value
}
// SetUserUrl sets the user_url property value. The user_url property
func (m *Feed) SetUserUrl(value *string)() {
    m.user_url = value
}
type Feedable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentUserActorUrl()(*string)
    GetCurrentUserOrganizationUrl()(*string)
    GetCurrentUserOrganizationUrls()([]string)
    GetCurrentUserPublicUrl()(*string)
    GetCurrentUserUrl()(*string)
    GetLinks()(Feed__linksable)
    GetRepositoryDiscussionsCategoryUrl()(*string)
    GetRepositoryDiscussionsUrl()(*string)
    GetSecurityAdvisoriesUrl()(*string)
    GetTimelineUrl()(*string)
    GetUserUrl()(*string)
    SetCurrentUserActorUrl(value *string)()
    SetCurrentUserOrganizationUrl(value *string)()
    SetCurrentUserOrganizationUrls(value []string)()
    SetCurrentUserPublicUrl(value *string)()
    SetCurrentUserUrl(value *string)()
    SetLinks(value Feed__linksable)()
    SetRepositoryDiscussionsCategoryUrl(value *string)()
    SetRepositoryDiscussionsUrl(value *string)()
    SetSecurityAdvisoriesUrl(value *string)()
    SetTimelineUrl(value *string)()
    SetUserUrl(value *string)()
}
