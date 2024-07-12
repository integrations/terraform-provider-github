package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Page the configuration for GitHub Pages for a repository.
type Page struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The process in which the Page will be built.
    build_type *Page_build_type
    // The Pages site's custom domain
    cname *string
    // Whether the Page has a custom 404 page.
    custom_404 *bool
    // The web address the Page can be accessed from.
    html_url *string
    // The https_certificate property
    https_certificate PagesHttpsCertificateable
    // Whether https is enabled on the domain
    https_enforced *bool
    // The timestamp when a pending domain becomes unverified.
    pending_domain_unverified_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The state if the domain is verified
    protected_domain_state *Page_protected_domain_state
    // Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.
    public *bool
    // The source property
    source PagesSourceHashable
    // The status of the most recent build of the Page.
    status *Page_status
    // The API address for accessing this Page resource.
    url *string
}
// NewPage instantiates a new Page and sets the default values.
func NewPage()(*Page) {
    m := &Page{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPage(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Page) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBuildType gets the build_type property value. The process in which the Page will be built.
// returns a *Page_build_type when successful
func (m *Page) GetBuildType()(*Page_build_type) {
    return m.build_type
}
// GetCname gets the cname property value. The Pages site's custom domain
// returns a *string when successful
func (m *Page) GetCname()(*string) {
    return m.cname
}
// GetCustom404 gets the custom_404 property value. Whether the Page has a custom 404 page.
// returns a *bool when successful
func (m *Page) GetCustom404()(*bool) {
    return m.custom_404
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Page) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["build_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePage_build_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuildType(val.(*Page_build_type))
        }
        return nil
    }
    res["cname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCname(val)
        }
        return nil
    }
    res["custom_404"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustom404(val)
        }
        return nil
    }
    res["html_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtmlUrl(val)
        }
        return nil
    }
    res["https_certificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePagesHttpsCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHttpsCertificate(val.(PagesHttpsCertificateable))
        }
        return nil
    }
    res["https_enforced"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHttpsEnforced(val)
        }
        return nil
    }
    res["pending_domain_unverified_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingDomainUnverifiedAt(val)
        }
        return nil
    }
    res["protected_domain_state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePage_protected_domain_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectedDomainState(val.(*Page_protected_domain_state))
        }
        return nil
    }
    res["public"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublic(val)
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePagesSourceHashFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(PagesSourceHashable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePage_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*Page_status))
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetHtmlUrl gets the html_url property value. The web address the Page can be accessed from.
// returns a *string when successful
func (m *Page) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetHttpsCertificate gets the https_certificate property value. The https_certificate property
// returns a PagesHttpsCertificateable when successful
func (m *Page) GetHttpsCertificate()(PagesHttpsCertificateable) {
    return m.https_certificate
}
// GetHttpsEnforced gets the https_enforced property value. Whether https is enabled on the domain
// returns a *bool when successful
func (m *Page) GetHttpsEnforced()(*bool) {
    return m.https_enforced
}
// GetPendingDomainUnverifiedAt gets the pending_domain_unverified_at property value. The timestamp when a pending domain becomes unverified.
// returns a *Time when successful
func (m *Page) GetPendingDomainUnverifiedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.pending_domain_unverified_at
}
// GetProtectedDomainState gets the protected_domain_state property value. The state if the domain is verified
// returns a *Page_protected_domain_state when successful
func (m *Page) GetProtectedDomainState()(*Page_protected_domain_state) {
    return m.protected_domain_state
}
// GetPublic gets the public property value. Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.
// returns a *bool when successful
func (m *Page) GetPublic()(*bool) {
    return m.public
}
// GetSource gets the source property value. The source property
// returns a PagesSourceHashable when successful
func (m *Page) GetSource()(PagesSourceHashable) {
    return m.source
}
// GetStatus gets the status property value. The status of the most recent build of the Page.
// returns a *Page_status when successful
func (m *Page) GetStatus()(*Page_status) {
    return m.status
}
// GetUrl gets the url property value. The API address for accessing this Page resource.
// returns a *string when successful
func (m *Page) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *Page) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBuildType() != nil {
        cast := (*m.GetBuildType()).String()
        err := writer.WriteStringValue("build_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cname", m.GetCname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("custom_404", m.GetCustom404())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("https_certificate", m.GetHttpsCertificate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("https_enforced", m.GetHttpsEnforced())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("pending_domain_unverified_at", m.GetPendingDomainUnverifiedAt())
        if err != nil {
            return err
        }
    }
    if m.GetProtectedDomainState() != nil {
        cast := (*m.GetProtectedDomainState()).String()
        err := writer.WriteStringValue("protected_domain_state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("public", m.GetPublic())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *Page) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBuildType sets the build_type property value. The process in which the Page will be built.
func (m *Page) SetBuildType(value *Page_build_type)() {
    m.build_type = value
}
// SetCname sets the cname property value. The Pages site's custom domain
func (m *Page) SetCname(value *string)() {
    m.cname = value
}
// SetCustom404 sets the custom_404 property value. Whether the Page has a custom 404 page.
func (m *Page) SetCustom404(value *bool)() {
    m.custom_404 = value
}
// SetHtmlUrl sets the html_url property value. The web address the Page can be accessed from.
func (m *Page) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetHttpsCertificate sets the https_certificate property value. The https_certificate property
func (m *Page) SetHttpsCertificate(value PagesHttpsCertificateable)() {
    m.https_certificate = value
}
// SetHttpsEnforced sets the https_enforced property value. Whether https is enabled on the domain
func (m *Page) SetHttpsEnforced(value *bool)() {
    m.https_enforced = value
}
// SetPendingDomainUnverifiedAt sets the pending_domain_unverified_at property value. The timestamp when a pending domain becomes unverified.
func (m *Page) SetPendingDomainUnverifiedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.pending_domain_unverified_at = value
}
// SetProtectedDomainState sets the protected_domain_state property value. The state if the domain is verified
func (m *Page) SetProtectedDomainState(value *Page_protected_domain_state)() {
    m.protected_domain_state = value
}
// SetPublic sets the public property value. Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.
func (m *Page) SetPublic(value *bool)() {
    m.public = value
}
// SetSource sets the source property value. The source property
func (m *Page) SetSource(value PagesSourceHashable)() {
    m.source = value
}
// SetStatus sets the status property value. The status of the most recent build of the Page.
func (m *Page) SetStatus(value *Page_status)() {
    m.status = value
}
// SetUrl sets the url property value. The API address for accessing this Page resource.
func (m *Page) SetUrl(value *string)() {
    m.url = value
}
type Pageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBuildType()(*Page_build_type)
    GetCname()(*string)
    GetCustom404()(*bool)
    GetHtmlUrl()(*string)
    GetHttpsCertificate()(PagesHttpsCertificateable)
    GetHttpsEnforced()(*bool)
    GetPendingDomainUnverifiedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProtectedDomainState()(*Page_protected_domain_state)
    GetPublic()(*bool)
    GetSource()(PagesSourceHashable)
    GetStatus()(*Page_status)
    GetUrl()(*string)
    SetBuildType(value *Page_build_type)()
    SetCname(value *string)()
    SetCustom404(value *bool)()
    SetHtmlUrl(value *string)()
    SetHttpsCertificate(value PagesHttpsCertificateable)()
    SetHttpsEnforced(value *bool)()
    SetPendingDomainUnverifiedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProtectedDomainState(value *Page_protected_domain_state)()
    SetPublic(value *bool)()
    SetSource(value PagesSourceHashable)()
    SetStatus(value *Page_status)()
    SetUrl(value *string)()
}
