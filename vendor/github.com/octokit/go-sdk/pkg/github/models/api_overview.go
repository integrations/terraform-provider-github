package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApiOverview api Overview
type ApiOverview struct {
    // The actions property
    actions []string
    // The actions_macos property
    actions_macos []string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The api property
    api []string
    // The dependabot property
    dependabot []string
    // The domains property
    domains ApiOverview_domainsable
    // The git property
    git []string
    // The github_enterprise_importer property
    github_enterprise_importer []string
    // The hooks property
    hooks []string
    // The importer property
    importer []string
    // The packages property
    packages []string
    // The pages property
    pages []string
    // The ssh_key_fingerprints property
    ssh_key_fingerprints ApiOverview_ssh_key_fingerprintsable
    // The ssh_keys property
    ssh_keys []string
    // The verifiable_password_authentication property
    verifiable_password_authentication *bool
    // The web property
    web []string
}
// NewApiOverview instantiates a new ApiOverview and sets the default values.
func NewApiOverview()(*ApiOverview) {
    m := &ApiOverview{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiOverview(), nil
}
// GetActions gets the actions property value. The actions property
// returns a []string when successful
func (m *ApiOverview) GetActions()([]string) {
    return m.actions
}
// GetActionsMacos gets the actions_macos property value. The actions_macos property
// returns a []string when successful
func (m *ApiOverview) GetActionsMacos()([]string) {
    return m.actions_macos
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApiOverview) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApi gets the api property value. The api property
// returns a []string when successful
func (m *ApiOverview) GetApi()([]string) {
    return m.api
}
// GetDependabot gets the dependabot property value. The dependabot property
// returns a []string when successful
func (m *ApiOverview) GetDependabot()([]string) {
    return m.dependabot
}
// GetDomains gets the domains property value. The domains property
// returns a ApiOverview_domainsable when successful
func (m *ApiOverview) GetDomains()(ApiOverview_domainsable) {
    return m.domains
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetActions(res)
        }
        return nil
    }
    res["actions_macos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetActionsMacos(res)
        }
        return nil
    }
    res["api"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetApi(res)
        }
        return nil
    }
    res["dependabot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetDependabot(res)
        }
        return nil
    }
    res["domains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApiOverview_domainsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomains(val.(ApiOverview_domainsable))
        }
        return nil
    }
    res["git"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetGit(res)
        }
        return nil
    }
    res["github_enterprise_importer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetGithubEnterpriseImporter(res)
        }
        return nil
    }
    res["hooks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetHooks(res)
        }
        return nil
    }
    res["importer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetImporter(res)
        }
        return nil
    }
    res["packages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPackages(res)
        }
        return nil
    }
    res["pages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPages(res)
        }
        return nil
    }
    res["ssh_key_fingerprints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApiOverview_ssh_key_fingerprintsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSshKeyFingerprints(val.(ApiOverview_ssh_key_fingerprintsable))
        }
        return nil
    }
    res["ssh_keys"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSshKeys(res)
        }
        return nil
    }
    res["verifiable_password_authentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiablePasswordAuthentication(val)
        }
        return nil
    }
    res["web"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetWeb(res)
        }
        return nil
    }
    return res
}
// GetGit gets the git property value. The git property
// returns a []string when successful
func (m *ApiOverview) GetGit()([]string) {
    return m.git
}
// GetGithubEnterpriseImporter gets the github_enterprise_importer property value. The github_enterprise_importer property
// returns a []string when successful
func (m *ApiOverview) GetGithubEnterpriseImporter()([]string) {
    return m.github_enterprise_importer
}
// GetHooks gets the hooks property value. The hooks property
// returns a []string when successful
func (m *ApiOverview) GetHooks()([]string) {
    return m.hooks
}
// GetImporter gets the importer property value. The importer property
// returns a []string when successful
func (m *ApiOverview) GetImporter()([]string) {
    return m.importer
}
// GetPackages gets the packages property value. The packages property
// returns a []string when successful
func (m *ApiOverview) GetPackages()([]string) {
    return m.packages
}
// GetPages gets the pages property value. The pages property
// returns a []string when successful
func (m *ApiOverview) GetPages()([]string) {
    return m.pages
}
// GetSshKeyFingerprints gets the ssh_key_fingerprints property value. The ssh_key_fingerprints property
// returns a ApiOverview_ssh_key_fingerprintsable when successful
func (m *ApiOverview) GetSshKeyFingerprints()(ApiOverview_ssh_key_fingerprintsable) {
    return m.ssh_key_fingerprints
}
// GetSshKeys gets the ssh_keys property value. The ssh_keys property
// returns a []string when successful
func (m *ApiOverview) GetSshKeys()([]string) {
    return m.ssh_keys
}
// GetVerifiablePasswordAuthentication gets the verifiable_password_authentication property value. The verifiable_password_authentication property
// returns a *bool when successful
func (m *ApiOverview) GetVerifiablePasswordAuthentication()(*bool) {
    return m.verifiable_password_authentication
}
// GetWeb gets the web property value. The web property
// returns a []string when successful
func (m *ApiOverview) GetWeb()([]string) {
    return m.web
}
// Serialize serializes information the current object
func (m *ApiOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActions() != nil {
        err := writer.WriteCollectionOfStringValues("actions", m.GetActions())
        if err != nil {
            return err
        }
    }
    if m.GetActionsMacos() != nil {
        err := writer.WriteCollectionOfStringValues("actions_macos", m.GetActionsMacos())
        if err != nil {
            return err
        }
    }
    if m.GetApi() != nil {
        err := writer.WriteCollectionOfStringValues("api", m.GetApi())
        if err != nil {
            return err
        }
    }
    if m.GetDependabot() != nil {
        err := writer.WriteCollectionOfStringValues("dependabot", m.GetDependabot())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("domains", m.GetDomains())
        if err != nil {
            return err
        }
    }
    if m.GetGit() != nil {
        err := writer.WriteCollectionOfStringValues("git", m.GetGit())
        if err != nil {
            return err
        }
    }
    if m.GetGithubEnterpriseImporter() != nil {
        err := writer.WriteCollectionOfStringValues("github_enterprise_importer", m.GetGithubEnterpriseImporter())
        if err != nil {
            return err
        }
    }
    if m.GetHooks() != nil {
        err := writer.WriteCollectionOfStringValues("hooks", m.GetHooks())
        if err != nil {
            return err
        }
    }
    if m.GetImporter() != nil {
        err := writer.WriteCollectionOfStringValues("importer", m.GetImporter())
        if err != nil {
            return err
        }
    }
    if m.GetPackages() != nil {
        err := writer.WriteCollectionOfStringValues("packages", m.GetPackages())
        if err != nil {
            return err
        }
    }
    if m.GetPages() != nil {
        err := writer.WriteCollectionOfStringValues("pages", m.GetPages())
        if err != nil {
            return err
        }
    }
    if m.GetSshKeys() != nil {
        err := writer.WriteCollectionOfStringValues("ssh_keys", m.GetSshKeys())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("ssh_key_fingerprints", m.GetSshKeyFingerprints())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("verifiable_password_authentication", m.GetVerifiablePasswordAuthentication())
        if err != nil {
            return err
        }
    }
    if m.GetWeb() != nil {
        err := writer.WriteCollectionOfStringValues("web", m.GetWeb())
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
// SetActions sets the actions property value. The actions property
func (m *ApiOverview) SetActions(value []string)() {
    m.actions = value
}
// SetActionsMacos sets the actions_macos property value. The actions_macos property
func (m *ApiOverview) SetActionsMacos(value []string)() {
    m.actions_macos = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApiOverview) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApi sets the api property value. The api property
func (m *ApiOverview) SetApi(value []string)() {
    m.api = value
}
// SetDependabot sets the dependabot property value. The dependabot property
func (m *ApiOverview) SetDependabot(value []string)() {
    m.dependabot = value
}
// SetDomains sets the domains property value. The domains property
func (m *ApiOverview) SetDomains(value ApiOverview_domainsable)() {
    m.domains = value
}
// SetGit sets the git property value. The git property
func (m *ApiOverview) SetGit(value []string)() {
    m.git = value
}
// SetGithubEnterpriseImporter sets the github_enterprise_importer property value. The github_enterprise_importer property
func (m *ApiOverview) SetGithubEnterpriseImporter(value []string)() {
    m.github_enterprise_importer = value
}
// SetHooks sets the hooks property value. The hooks property
func (m *ApiOverview) SetHooks(value []string)() {
    m.hooks = value
}
// SetImporter sets the importer property value. The importer property
func (m *ApiOverview) SetImporter(value []string)() {
    m.importer = value
}
// SetPackages sets the packages property value. The packages property
func (m *ApiOverview) SetPackages(value []string)() {
    m.packages = value
}
// SetPages sets the pages property value. The pages property
func (m *ApiOverview) SetPages(value []string)() {
    m.pages = value
}
// SetSshKeyFingerprints sets the ssh_key_fingerprints property value. The ssh_key_fingerprints property
func (m *ApiOverview) SetSshKeyFingerprints(value ApiOverview_ssh_key_fingerprintsable)() {
    m.ssh_key_fingerprints = value
}
// SetSshKeys sets the ssh_keys property value. The ssh_keys property
func (m *ApiOverview) SetSshKeys(value []string)() {
    m.ssh_keys = value
}
// SetVerifiablePasswordAuthentication sets the verifiable_password_authentication property value. The verifiable_password_authentication property
func (m *ApiOverview) SetVerifiablePasswordAuthentication(value *bool)() {
    m.verifiable_password_authentication = value
}
// SetWeb sets the web property value. The web property
func (m *ApiOverview) SetWeb(value []string)() {
    m.web = value
}
type ApiOverviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]string)
    GetActionsMacos()([]string)
    GetApi()([]string)
    GetDependabot()([]string)
    GetDomains()(ApiOverview_domainsable)
    GetGit()([]string)
    GetGithubEnterpriseImporter()([]string)
    GetHooks()([]string)
    GetImporter()([]string)
    GetPackages()([]string)
    GetPages()([]string)
    GetSshKeyFingerprints()(ApiOverview_ssh_key_fingerprintsable)
    GetSshKeys()([]string)
    GetVerifiablePasswordAuthentication()(*bool)
    GetWeb()([]string)
    SetActions(value []string)()
    SetActionsMacos(value []string)()
    SetApi(value []string)()
    SetDependabot(value []string)()
    SetDomains(value ApiOverview_domainsable)()
    SetGit(value []string)()
    SetGithubEnterpriseImporter(value []string)()
    SetHooks(value []string)()
    SetImporter(value []string)()
    SetPackages(value []string)()
    SetPages(value []string)()
    SetSshKeyFingerprints(value ApiOverview_ssh_key_fingerprintsable)()
    SetSshKeys(value []string)()
    SetVerifiablePasswordAuthentication(value *bool)()
    SetWeb(value []string)()
}
