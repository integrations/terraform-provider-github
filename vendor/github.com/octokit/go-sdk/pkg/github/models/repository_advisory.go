package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoryAdvisory a repository security advisory.
type RepositoryAdvisory struct {
    // The author of the advisory.
    author SimpleUserable
    // The date and time of when the advisory was closed, in ISO 8601 format.
    closed_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A list of teams that collaborate on the advisory.
    collaborating_teams []Teamable
    // A list of users that collaborate on the advisory.
    collaborating_users []SimpleUserable
    // The date and time of when the advisory was created, in ISO 8601 format.
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The credits property
    credits []RepositoryAdvisory_creditsable
    // The credits_detailed property
    credits_detailed []RepositoryAdvisoryCreditable
    // The Common Vulnerabilities and Exposures (CVE) ID.
    cve_id *string
    // The cvss property
    cvss RepositoryAdvisory_cvssable
    // A list of only the CWE IDs.
    cwe_ids []string
    // The cwes property
    cwes []RepositoryAdvisory_cwesable
    // A detailed description of what the advisory entails.
    description *string
    // The GitHub Security Advisory ID.
    ghsa_id *string
    // The URL for the advisory.
    html_url *string
    // The identifiers property
    identifiers []RepositoryAdvisory_identifiersable
    // A temporary private fork of the advisory's repository for collaborating on a fix.
    private_fork SimpleRepositoryable
    // The date and time of when the advisory was published, in ISO 8601 format.
    published_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The publisher of the advisory.
    publisher SimpleUserable
    // The severity of the advisory.
    severity *RepositoryAdvisory_severity
    // The state of the advisory.
    state *RepositoryAdvisory_state
    // The submission property
    submission RepositoryAdvisory_submissionable
    // A short summary of the advisory.
    summary *string
    // The date and time of when the advisory was last updated, in ISO 8601 format.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The API URL for the advisory.
    url *string
    // The vulnerabilities property
    vulnerabilities []RepositoryAdvisoryVulnerabilityable
    // The date and time of when the advisory was withdrawn, in ISO 8601 format.
    withdrawn_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewRepositoryAdvisory instantiates a new RepositoryAdvisory and sets the default values.
func NewRepositoryAdvisory()(*RepositoryAdvisory) {
    m := &RepositoryAdvisory{
    }
    return m
}
// CreateRepositoryAdvisoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryAdvisoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryAdvisory(), nil
}
// GetAuthor gets the author property value. The author of the advisory.
// returns a SimpleUserable when successful
func (m *RepositoryAdvisory) GetAuthor()(SimpleUserable) {
    return m.author
}
// GetClosedAt gets the closed_at property value. The date and time of when the advisory was closed, in ISO 8601 format.
// returns a *Time when successful
func (m *RepositoryAdvisory) GetClosedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.closed_at
}
// GetCollaboratingTeams gets the collaborating_teams property value. A list of teams that collaborate on the advisory.
// returns a []Teamable when successful
func (m *RepositoryAdvisory) GetCollaboratingTeams()([]Teamable) {
    return m.collaborating_teams
}
// GetCollaboratingUsers gets the collaborating_users property value. A list of users that collaborate on the advisory.
// returns a []SimpleUserable when successful
func (m *RepositoryAdvisory) GetCollaboratingUsers()([]SimpleUserable) {
    return m.collaborating_users
}
// GetCreatedAt gets the created_at property value. The date and time of when the advisory was created, in ISO 8601 format.
// returns a *Time when successful
func (m *RepositoryAdvisory) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetCredits gets the credits property value. The credits property
// returns a []RepositoryAdvisory_creditsable when successful
func (m *RepositoryAdvisory) GetCredits()([]RepositoryAdvisory_creditsable) {
    return m.credits
}
// GetCreditsDetailed gets the credits_detailed property value. The credits_detailed property
// returns a []RepositoryAdvisoryCreditable when successful
func (m *RepositoryAdvisory) GetCreditsDetailed()([]RepositoryAdvisoryCreditable) {
    return m.credits_detailed
}
// GetCveId gets the cve_id property value. The Common Vulnerabilities and Exposures (CVE) ID.
// returns a *string when successful
func (m *RepositoryAdvisory) GetCveId()(*string) {
    return m.cve_id
}
// GetCvss gets the cvss property value. The cvss property
// returns a RepositoryAdvisory_cvssable when successful
func (m *RepositoryAdvisory) GetCvss()(RepositoryAdvisory_cvssable) {
    return m.cvss
}
// GetCweIds gets the cwe_ids property value. A list of only the CWE IDs.
// returns a []string when successful
func (m *RepositoryAdvisory) GetCweIds()([]string) {
    return m.cwe_ids
}
// GetCwes gets the cwes property value. The cwes property
// returns a []RepositoryAdvisory_cwesable when successful
func (m *RepositoryAdvisory) GetCwes()([]RepositoryAdvisory_cwesable) {
    return m.cwes
}
// GetDescription gets the description property value. A detailed description of what the advisory entails.
// returns a *string when successful
func (m *RepositoryAdvisory) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryAdvisory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val.(SimpleUserable))
        }
        return nil
    }
    res["closed_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClosedAt(val)
        }
        return nil
    }
    res["collaborating_teams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Teamable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Teamable)
                }
            }
            m.SetCollaboratingTeams(res)
        }
        return nil
    }
    res["collaborating_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SimpleUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SimpleUserable)
                }
            }
            m.SetCollaboratingUsers(res)
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["credits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryAdvisory_creditsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RepositoryAdvisory_creditsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RepositoryAdvisory_creditsable)
                }
            }
            m.SetCredits(res)
        }
        return nil
    }
    res["credits_detailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryAdvisoryCreditFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RepositoryAdvisoryCreditable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RepositoryAdvisoryCreditable)
                }
            }
            m.SetCreditsDetailed(res)
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
        val, err := n.GetObjectValue(CreateRepositoryAdvisory_cvssFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCvss(val.(RepositoryAdvisory_cvssable))
        }
        return nil
    }
    res["cwe_ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetCweIds(res)
        }
        return nil
    }
    res["cwes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryAdvisory_cwesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RepositoryAdvisory_cwesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RepositoryAdvisory_cwesable)
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
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryAdvisory_identifiersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RepositoryAdvisory_identifiersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RepositoryAdvisory_identifiersable)
                }
            }
            m.SetIdentifiers(res)
        }
        return nil
    }
    res["private_fork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateFork(val.(SimpleRepositoryable))
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
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val.(SimpleUserable))
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryAdvisory_severity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*RepositoryAdvisory_severity))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryAdvisory_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*RepositoryAdvisory_state))
        }
        return nil
    }
    res["submission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRepositoryAdvisory_submissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmission(val.(RepositoryAdvisory_submissionable))
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
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryAdvisoryVulnerabilityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RepositoryAdvisoryVulnerabilityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RepositoryAdvisoryVulnerabilityable)
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
func (m *RepositoryAdvisory) GetGhsaId()(*string) {
    return m.ghsa_id
}
// GetHtmlUrl gets the html_url property value. The URL for the advisory.
// returns a *string when successful
func (m *RepositoryAdvisory) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetIdentifiers gets the identifiers property value. The identifiers property
// returns a []RepositoryAdvisory_identifiersable when successful
func (m *RepositoryAdvisory) GetIdentifiers()([]RepositoryAdvisory_identifiersable) {
    return m.identifiers
}
// GetPrivateFork gets the private_fork property value. A temporary private fork of the advisory's repository for collaborating on a fix.
// returns a SimpleRepositoryable when successful
func (m *RepositoryAdvisory) GetPrivateFork()(SimpleRepositoryable) {
    return m.private_fork
}
// GetPublishedAt gets the published_at property value. The date and time of when the advisory was published, in ISO 8601 format.
// returns a *Time when successful
func (m *RepositoryAdvisory) GetPublishedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.published_at
}
// GetPublisher gets the publisher property value. The publisher of the advisory.
// returns a SimpleUserable when successful
func (m *RepositoryAdvisory) GetPublisher()(SimpleUserable) {
    return m.publisher
}
// GetSeverity gets the severity property value. The severity of the advisory.
// returns a *RepositoryAdvisory_severity when successful
func (m *RepositoryAdvisory) GetSeverity()(*RepositoryAdvisory_severity) {
    return m.severity
}
// GetState gets the state property value. The state of the advisory.
// returns a *RepositoryAdvisory_state when successful
func (m *RepositoryAdvisory) GetState()(*RepositoryAdvisory_state) {
    return m.state
}
// GetSubmission gets the submission property value. The submission property
// returns a RepositoryAdvisory_submissionable when successful
func (m *RepositoryAdvisory) GetSubmission()(RepositoryAdvisory_submissionable) {
    return m.submission
}
// GetSummary gets the summary property value. A short summary of the advisory.
// returns a *string when successful
func (m *RepositoryAdvisory) GetSummary()(*string) {
    return m.summary
}
// GetUpdatedAt gets the updated_at property value. The date and time of when the advisory was last updated, in ISO 8601 format.
// returns a *Time when successful
func (m *RepositoryAdvisory) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The API URL for the advisory.
// returns a *string when successful
func (m *RepositoryAdvisory) GetUrl()(*string) {
    return m.url
}
// GetVulnerabilities gets the vulnerabilities property value. The vulnerabilities property
// returns a []RepositoryAdvisoryVulnerabilityable when successful
func (m *RepositoryAdvisory) GetVulnerabilities()([]RepositoryAdvisoryVulnerabilityable) {
    return m.vulnerabilities
}
// GetWithdrawnAt gets the withdrawn_at property value. The date and time of when the advisory was withdrawn, in ISO 8601 format.
// returns a *Time when successful
func (m *RepositoryAdvisory) GetWithdrawnAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.withdrawn_at
}
// Serialize serializes information the current object
func (m *RepositoryAdvisory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCollaboratingTeams() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCollaboratingTeams()))
        for i, v := range m.GetCollaboratingTeams() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("collaborating_teams", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCollaboratingUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCollaboratingUsers()))
        for i, v := range m.GetCollaboratingUsers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("collaborating_users", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCredits() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCredits()))
        for i, v := range m.GetCredits() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("credits", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cve_id", m.GetCveId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("cvss", m.GetCvss())
        if err != nil {
            return err
        }
    }
    if m.GetCweIds() != nil {
        err := writer.WriteCollectionOfStringValues("cwe_ids", m.GetCweIds())
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
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err := writer.WriteStringValue("severity", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
// SetAuthor sets the author property value. The author of the advisory.
func (m *RepositoryAdvisory) SetAuthor(value SimpleUserable)() {
    m.author = value
}
// SetClosedAt sets the closed_at property value. The date and time of when the advisory was closed, in ISO 8601 format.
func (m *RepositoryAdvisory) SetClosedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closed_at = value
}
// SetCollaboratingTeams sets the collaborating_teams property value. A list of teams that collaborate on the advisory.
func (m *RepositoryAdvisory) SetCollaboratingTeams(value []Teamable)() {
    m.collaborating_teams = value
}
// SetCollaboratingUsers sets the collaborating_users property value. A list of users that collaborate on the advisory.
func (m *RepositoryAdvisory) SetCollaboratingUsers(value []SimpleUserable)() {
    m.collaborating_users = value
}
// SetCreatedAt sets the created_at property value. The date and time of when the advisory was created, in ISO 8601 format.
func (m *RepositoryAdvisory) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetCredits sets the credits property value. The credits property
func (m *RepositoryAdvisory) SetCredits(value []RepositoryAdvisory_creditsable)() {
    m.credits = value
}
// SetCreditsDetailed sets the credits_detailed property value. The credits_detailed property
func (m *RepositoryAdvisory) SetCreditsDetailed(value []RepositoryAdvisoryCreditable)() {
    m.credits_detailed = value
}
// SetCveId sets the cve_id property value. The Common Vulnerabilities and Exposures (CVE) ID.
func (m *RepositoryAdvisory) SetCveId(value *string)() {
    m.cve_id = value
}
// SetCvss sets the cvss property value. The cvss property
func (m *RepositoryAdvisory) SetCvss(value RepositoryAdvisory_cvssable)() {
    m.cvss = value
}
// SetCweIds sets the cwe_ids property value. A list of only the CWE IDs.
func (m *RepositoryAdvisory) SetCweIds(value []string)() {
    m.cwe_ids = value
}
// SetCwes sets the cwes property value. The cwes property
func (m *RepositoryAdvisory) SetCwes(value []RepositoryAdvisory_cwesable)() {
    m.cwes = value
}
// SetDescription sets the description property value. A detailed description of what the advisory entails.
func (m *RepositoryAdvisory) SetDescription(value *string)() {
    m.description = value
}
// SetGhsaId sets the ghsa_id property value. The GitHub Security Advisory ID.
func (m *RepositoryAdvisory) SetGhsaId(value *string)() {
    m.ghsa_id = value
}
// SetHtmlUrl sets the html_url property value. The URL for the advisory.
func (m *RepositoryAdvisory) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetIdentifiers sets the identifiers property value. The identifiers property
func (m *RepositoryAdvisory) SetIdentifiers(value []RepositoryAdvisory_identifiersable)() {
    m.identifiers = value
}
// SetPrivateFork sets the private_fork property value. A temporary private fork of the advisory's repository for collaborating on a fix.
func (m *RepositoryAdvisory) SetPrivateFork(value SimpleRepositoryable)() {
    m.private_fork = value
}
// SetPublishedAt sets the published_at property value. The date and time of when the advisory was published, in ISO 8601 format.
func (m *RepositoryAdvisory) SetPublishedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.published_at = value
}
// SetPublisher sets the publisher property value. The publisher of the advisory.
func (m *RepositoryAdvisory) SetPublisher(value SimpleUserable)() {
    m.publisher = value
}
// SetSeverity sets the severity property value. The severity of the advisory.
func (m *RepositoryAdvisory) SetSeverity(value *RepositoryAdvisory_severity)() {
    m.severity = value
}
// SetState sets the state property value. The state of the advisory.
func (m *RepositoryAdvisory) SetState(value *RepositoryAdvisory_state)() {
    m.state = value
}
// SetSubmission sets the submission property value. The submission property
func (m *RepositoryAdvisory) SetSubmission(value RepositoryAdvisory_submissionable)() {
    m.submission = value
}
// SetSummary sets the summary property value. A short summary of the advisory.
func (m *RepositoryAdvisory) SetSummary(value *string)() {
    m.summary = value
}
// SetUpdatedAt sets the updated_at property value. The date and time of when the advisory was last updated, in ISO 8601 format.
func (m *RepositoryAdvisory) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The API URL for the advisory.
func (m *RepositoryAdvisory) SetUrl(value *string)() {
    m.url = value
}
// SetVulnerabilities sets the vulnerabilities property value. The vulnerabilities property
func (m *RepositoryAdvisory) SetVulnerabilities(value []RepositoryAdvisoryVulnerabilityable)() {
    m.vulnerabilities = value
}
// SetWithdrawnAt sets the withdrawn_at property value. The date and time of when the advisory was withdrawn, in ISO 8601 format.
func (m *RepositoryAdvisory) SetWithdrawnAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.withdrawn_at = value
}
type RepositoryAdvisoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(SimpleUserable)
    GetClosedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCollaboratingTeams()([]Teamable)
    GetCollaboratingUsers()([]SimpleUserable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCredits()([]RepositoryAdvisory_creditsable)
    GetCreditsDetailed()([]RepositoryAdvisoryCreditable)
    GetCveId()(*string)
    GetCvss()(RepositoryAdvisory_cvssable)
    GetCweIds()([]string)
    GetCwes()([]RepositoryAdvisory_cwesable)
    GetDescription()(*string)
    GetGhsaId()(*string)
    GetHtmlUrl()(*string)
    GetIdentifiers()([]RepositoryAdvisory_identifiersable)
    GetPrivateFork()(SimpleRepositoryable)
    GetPublishedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPublisher()(SimpleUserable)
    GetSeverity()(*RepositoryAdvisory_severity)
    GetState()(*RepositoryAdvisory_state)
    GetSubmission()(RepositoryAdvisory_submissionable)
    GetSummary()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    GetVulnerabilities()([]RepositoryAdvisoryVulnerabilityable)
    GetWithdrawnAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAuthor(value SimpleUserable)()
    SetClosedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCollaboratingTeams(value []Teamable)()
    SetCollaboratingUsers(value []SimpleUserable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCredits(value []RepositoryAdvisory_creditsable)()
    SetCreditsDetailed(value []RepositoryAdvisoryCreditable)()
    SetCveId(value *string)()
    SetCvss(value RepositoryAdvisory_cvssable)()
    SetCweIds(value []string)()
    SetCwes(value []RepositoryAdvisory_cwesable)()
    SetDescription(value *string)()
    SetGhsaId(value *string)()
    SetHtmlUrl(value *string)()
    SetIdentifiers(value []RepositoryAdvisory_identifiersable)()
    SetPrivateFork(value SimpleRepositoryable)()
    SetPublishedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPublisher(value SimpleUserable)()
    SetSeverity(value *RepositoryAdvisory_severity)()
    SetState(value *RepositoryAdvisory_state)()
    SetSubmission(value RepositoryAdvisory_submissionable)()
    SetSummary(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
    SetVulnerabilities(value []RepositoryAdvisoryVulnerabilityable)()
    SetWithdrawnAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
