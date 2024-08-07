package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GlobalAdvisory a GitHub Security Advisory.
type GlobalAdvisory struct {
    // The users who contributed to the advisory.
    credits []GlobalAdvisory_creditsable
    // The Common Vulnerabilities and Exposures (CVE) ID.
    cve_id *string
    // The cvss property
    cvss GlobalAdvisory_cvssable
    // The cwes property
    cwes []GlobalAdvisory_cwesable
    // A detailed description of what the advisory entails.
    description *string
    // The GitHub Security Advisory ID.
    ghsa_id *string
    // The date and time of when the advisory was reviewed by GitHub, in ISO 8601 format.
    github_reviewed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The URL for the advisory.
    html_url *string
    // The identifiers property
    identifiers []GlobalAdvisory_identifiersable
    // The date and time when the advisory was published in the National Vulnerability Database, in ISO 8601 format.This field is only populated when the advisory is imported from the National Vulnerability Database.
    nvd_published_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time of when the advisory was published, in ISO 8601 format.
    published_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The references property
    references []string
    // The API URL for the repository advisory.
    repository_advisory_url *string
    // The severity of the advisory.
    severity *GlobalAdvisory_severity
    // The URL of the advisory's source code.
    source_code_location *string
    // A short summary of the advisory.
    summary *string
    // The type of advisory.
    typeEscaped *GlobalAdvisory_type
    // The date and time of when the advisory was last updated, in ISO 8601 format.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The API URL for the advisory.
    url *string
    // The products and respective version ranges affected by the advisory.
    vulnerabilities []Vulnerabilityable
    // The date and time of when the advisory was withdrawn, in ISO 8601 format.
    withdrawn_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewGlobalAdvisory instantiates a new GlobalAdvisory and sets the default values.
func NewGlobalAdvisory()(*GlobalAdvisory) {
    m := &GlobalAdvisory{
    }
    return m
}
// CreateGlobalAdvisoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGlobalAdvisoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGlobalAdvisory(), nil
}
// GetCredits gets the credits property value. The users who contributed to the advisory.
// returns a []GlobalAdvisory_creditsable when successful
func (m *GlobalAdvisory) GetCredits()([]GlobalAdvisory_creditsable) {
    return m.credits
}
// GetCveId gets the cve_id property value. The Common Vulnerabilities and Exposures (CVE) ID.
// returns a *string when successful
func (m *GlobalAdvisory) GetCveId()(*string) {
    return m.cve_id
}
// GetCvss gets the cvss property value. The cvss property
// returns a GlobalAdvisory_cvssable when successful
func (m *GlobalAdvisory) GetCvss()(GlobalAdvisory_cvssable) {
    return m.cvss
}
// GetCwes gets the cwes property value. The cwes property
// returns a []GlobalAdvisory_cwesable when successful
func (m *GlobalAdvisory) GetCwes()([]GlobalAdvisory_cwesable) {
    return m.cwes
}
// GetDescription gets the description property value. A detailed description of what the advisory entails.
// returns a *string when successful
func (m *GlobalAdvisory) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GlobalAdvisory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["credits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGlobalAdvisory_creditsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GlobalAdvisory_creditsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GlobalAdvisory_creditsable)
                }
            }
            m.SetCredits(res)
        }
        return nil
    }
    res["cve_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCveId(val)
        }
        return nil
    }
    res["cvss"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGlobalAdvisory_cvssFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCvss(val.(GlobalAdvisory_cvssable))
        }
        return nil
    }
    res["cwes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGlobalAdvisory_cwesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GlobalAdvisory_cwesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GlobalAdvisory_cwesable)
                }
            }
            m.SetCwes(res)
        }
        return nil
    }
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
    res["ghsa_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGhsaId(val)
        }
        return nil
    }
    res["github_reviewed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGithubReviewedAt(val)
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
    res["identifiers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGlobalAdvisory_identifiersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GlobalAdvisory_identifiersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GlobalAdvisory_identifiersable)
                }
            }
            m.SetIdentifiers(res)
        }
        return nil
    }
    res["nvd_published_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNvdPublishedAt(val)
        }
        return nil
    }
    res["published_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedAt(val)
        }
        return nil
    }
    res["references"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetReferences(res)
        }
        return nil
    }
    res["repository_advisory_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryAdvisoryUrl(val)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGlobalAdvisory_severity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*GlobalAdvisory_severity))
        }
        return nil
    }
    res["source_code_location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceCodeLocation(val)
        }
        return nil
    }
    res["summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGlobalAdvisory_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*GlobalAdvisory_type))
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
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
    res["vulnerabilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVulnerabilityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Vulnerabilityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Vulnerabilityable)
                }
            }
            m.SetVulnerabilities(res)
        }
        return nil
    }
    res["withdrawn_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWithdrawnAt(val)
        }
        return nil
    }
    return res
}
// GetGhsaId gets the ghsa_id property value. The GitHub Security Advisory ID.
// returns a *string when successful
func (m *GlobalAdvisory) GetGhsaId()(*string) {
    return m.ghsa_id
}
// GetGithubReviewedAt gets the github_reviewed_at property value. The date and time of when the advisory was reviewed by GitHub, in ISO 8601 format.
// returns a *Time when successful
func (m *GlobalAdvisory) GetGithubReviewedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.github_reviewed_at
}
// GetHtmlUrl gets the html_url property value. The URL for the advisory.
// returns a *string when successful
func (m *GlobalAdvisory) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetIdentifiers gets the identifiers property value. The identifiers property
// returns a []GlobalAdvisory_identifiersable when successful
func (m *GlobalAdvisory) GetIdentifiers()([]GlobalAdvisory_identifiersable) {
    return m.identifiers
}
// GetNvdPublishedAt gets the nvd_published_at property value. The date and time when the advisory was published in the National Vulnerability Database, in ISO 8601 format.This field is only populated when the advisory is imported from the National Vulnerability Database.
// returns a *Time when successful
func (m *GlobalAdvisory) GetNvdPublishedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.nvd_published_at
}
// GetPublishedAt gets the published_at property value. The date and time of when the advisory was published, in ISO 8601 format.
// returns a *Time when successful
func (m *GlobalAdvisory) GetPublishedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.published_at
}
// GetReferences gets the references property value. The references property
// returns a []string when successful
func (m *GlobalAdvisory) GetReferences()([]string) {
    return m.references
}
// GetRepositoryAdvisoryUrl gets the repository_advisory_url property value. The API URL for the repository advisory.
// returns a *string when successful
func (m *GlobalAdvisory) GetRepositoryAdvisoryUrl()(*string) {
    return m.repository_advisory_url
}
// GetSeverity gets the severity property value. The severity of the advisory.
// returns a *GlobalAdvisory_severity when successful
func (m *GlobalAdvisory) GetSeverity()(*GlobalAdvisory_severity) {
    return m.severity
}
// GetSourceCodeLocation gets the source_code_location property value. The URL of the advisory's source code.
// returns a *string when successful
func (m *GlobalAdvisory) GetSourceCodeLocation()(*string) {
    return m.source_code_location
}
// GetSummary gets the summary property value. A short summary of the advisory.
// returns a *string when successful
func (m *GlobalAdvisory) GetSummary()(*string) {
    return m.summary
}
// GetTypeEscaped gets the type property value. The type of advisory.
// returns a *GlobalAdvisory_type when successful
func (m *GlobalAdvisory) GetTypeEscaped()(*GlobalAdvisory_type) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updated_at property value. The date and time of when the advisory was last updated, in ISO 8601 format.
// returns a *Time when successful
func (m *GlobalAdvisory) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The API URL for the advisory.
// returns a *string when successful
func (m *GlobalAdvisory) GetUrl()(*string) {
    return m.url
}
// GetVulnerabilities gets the vulnerabilities property value. The products and respective version ranges affected by the advisory.
// returns a []Vulnerabilityable when successful
func (m *GlobalAdvisory) GetVulnerabilities()([]Vulnerabilityable) {
    return m.vulnerabilities
}
// GetWithdrawnAt gets the withdrawn_at property value. The date and time of when the advisory was withdrawn, in ISO 8601 format.
// returns a *Time when successful
func (m *GlobalAdvisory) GetWithdrawnAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.withdrawn_at
}
// Serialize serializes information the current object
func (m *GlobalAdvisory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cvss", m.GetCvss())
        if err != nil {
            return err
        }
    }
    if m.GetCwes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCwes()))
        for i, v := range m.GetCwes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("cwes", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetReferences() != nil {
        err := writer.WriteCollectionOfStringValues("references", m.GetReferences())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err := writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("source_code_location", m.GetSourceCodeLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("summary", m.GetSummary())
        if err != nil {
            return err
        }
    }
    if m.GetVulnerabilities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVulnerabilities()))
        for i, v := range m.GetVulnerabilities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("vulnerabilities", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCredits sets the credits property value. The users who contributed to the advisory.
func (m *GlobalAdvisory) SetCredits(value []GlobalAdvisory_creditsable)() {
    m.credits = value
}
// SetCveId sets the cve_id property value. The Common Vulnerabilities and Exposures (CVE) ID.
func (m *GlobalAdvisory) SetCveId(value *string)() {
    m.cve_id = value
}
// SetCvss sets the cvss property value. The cvss property
func (m *GlobalAdvisory) SetCvss(value GlobalAdvisory_cvssable)() {
    m.cvss = value
}
// SetCwes sets the cwes property value. The cwes property
func (m *GlobalAdvisory) SetCwes(value []GlobalAdvisory_cwesable)() {
    m.cwes = value
}
// SetDescription sets the description property value. A detailed description of what the advisory entails.
func (m *GlobalAdvisory) SetDescription(value *string)() {
    m.description = value
}
// SetGhsaId sets the ghsa_id property value. The GitHub Security Advisory ID.
func (m *GlobalAdvisory) SetGhsaId(value *string)() {
    m.ghsa_id = value
}
// SetGithubReviewedAt sets the github_reviewed_at property value. The date and time of when the advisory was reviewed by GitHub, in ISO 8601 format.
func (m *GlobalAdvisory) SetGithubReviewedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.github_reviewed_at = value
}
// SetHtmlUrl sets the html_url property value. The URL for the advisory.
func (m *GlobalAdvisory) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetIdentifiers sets the identifiers property value. The identifiers property
func (m *GlobalAdvisory) SetIdentifiers(value []GlobalAdvisory_identifiersable)() {
    m.identifiers = value
}
// SetNvdPublishedAt sets the nvd_published_at property value. The date and time when the advisory was published in the National Vulnerability Database, in ISO 8601 format.This field is only populated when the advisory is imported from the National Vulnerability Database.
func (m *GlobalAdvisory) SetNvdPublishedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.nvd_published_at = value
}
// SetPublishedAt sets the published_at property value. The date and time of when the advisory was published, in ISO 8601 format.
func (m *GlobalAdvisory) SetPublishedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.published_at = value
}
// SetReferences sets the references property value. The references property
func (m *GlobalAdvisory) SetReferences(value []string)() {
    m.references = value
}
// SetRepositoryAdvisoryUrl sets the repository_advisory_url property value. The API URL for the repository advisory.
func (m *GlobalAdvisory) SetRepositoryAdvisoryUrl(value *string)() {
    m.repository_advisory_url = value
}
// SetSeverity sets the severity property value. The severity of the advisory.
func (m *GlobalAdvisory) SetSeverity(value *GlobalAdvisory_severity)() {
    m.severity = value
}
// SetSourceCodeLocation sets the source_code_location property value. The URL of the advisory's source code.
func (m *GlobalAdvisory) SetSourceCodeLocation(value *string)() {
    m.source_code_location = value
}
// SetSummary sets the summary property value. A short summary of the advisory.
func (m *GlobalAdvisory) SetSummary(value *string)() {
    m.summary = value
}
// SetTypeEscaped sets the type property value. The type of advisory.
func (m *GlobalAdvisory) SetTypeEscaped(value *GlobalAdvisory_type)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updated_at property value. The date and time of when the advisory was last updated, in ISO 8601 format.
func (m *GlobalAdvisory) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The API URL for the advisory.
func (m *GlobalAdvisory) SetUrl(value *string)() {
    m.url = value
}
// SetVulnerabilities sets the vulnerabilities property value. The products and respective version ranges affected by the advisory.
func (m *GlobalAdvisory) SetVulnerabilities(value []Vulnerabilityable)() {
    m.vulnerabilities = value
}
// SetWithdrawnAt sets the withdrawn_at property value. The date and time of when the advisory was withdrawn, in ISO 8601 format.
func (m *GlobalAdvisory) SetWithdrawnAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.withdrawn_at = value
}
type GlobalAdvisoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCredits()([]GlobalAdvisory_creditsable)
    GetCveId()(*string)
    GetCvss()(GlobalAdvisory_cvssable)
    GetCwes()([]GlobalAdvisory_cwesable)
    GetDescription()(*string)
    GetGhsaId()(*string)
    GetGithubReviewedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHtmlUrl()(*string)
    GetIdentifiers()([]GlobalAdvisory_identifiersable)
    GetNvdPublishedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPublishedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReferences()([]string)
    GetRepositoryAdvisoryUrl()(*string)
    GetSeverity()(*GlobalAdvisory_severity)
    GetSourceCodeLocation()(*string)
    GetSummary()(*string)
    GetTypeEscaped()(*GlobalAdvisory_type)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetVulnerabilities()([]Vulnerabilityable)
    GetWithdrawnAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCredits(value []GlobalAdvisory_creditsable)()
    SetCveId(value *string)()
    SetCvss(value GlobalAdvisory_cvssable)()
    SetCwes(value []GlobalAdvisory_cwesable)()
    SetDescription(value *string)()
    SetGhsaId(value *string)()
    SetGithubReviewedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHtmlUrl(value *string)()
    SetIdentifiers(value []GlobalAdvisory_identifiersable)()
    SetNvdPublishedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPublishedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReferences(value []string)()
    SetRepositoryAdvisoryUrl(value *string)()
    SetSeverity(value *GlobalAdvisory_severity)()
    SetSourceCodeLocation(value *string)()
    SetSummary(value *string)()
    SetTypeEscaped(value *GlobalAdvisory_type)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetVulnerabilities(value []Vulnerabilityable)()
    SetWithdrawnAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
