package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryAdvisoryUpdate struct {
    // A list of team slugs which have been granted write access to the advisory.
    collaborating_teams []string
    // A list of usernames who have been granted write access to the advisory.
    collaborating_users []string
    // A list of users receiving credit for their participation in the security advisory.
    credits []RepositoryAdvisoryUpdate_creditsable
    // The Common Vulnerabilities and Exposures (CVE) ID.
    cve_id *string
    // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
    cvss_vector_string *string
    // A list of Common Weakness Enumeration (CWE) IDs.
    cwe_ids []string
    // A detailed description of what the advisory impacts.
    description *string
    // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
    severity *RepositoryAdvisoryUpdate_severity
    // The state of the advisory.
    state *RepositoryAdvisoryUpdate_state
    // A short summary of the advisory.
    summary *string
    // A product affected by the vulnerability detailed in a repository security advisory.
    vulnerabilities []RepositoryAdvisoryUpdate_vulnerabilitiesable
}
// NewRepositoryAdvisoryUpdate instantiates a new RepositoryAdvisoryUpdate and sets the default values.
func NewRepositoryAdvisoryUpdate()(*RepositoryAdvisoryUpdate) {
    m := &RepositoryAdvisoryUpdate{
    }
    return m
}
// CreateRepositoryAdvisoryUpdateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryAdvisoryUpdateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryAdvisoryUpdate(), nil
}
// GetCollaboratingTeams gets the collaborating_teams property value. A list of team slugs which have been granted write access to the advisory.
// returns a []string when successful
func (m *RepositoryAdvisoryUpdate) GetCollaboratingTeams()([]string) {
    return m.collaborating_teams
}
// GetCollaboratingUsers gets the collaborating_users property value. A list of usernames who have been granted write access to the advisory.
// returns a []string when successful
func (m *RepositoryAdvisoryUpdate) GetCollaboratingUsers()([]string) {
    return m.collaborating_users
}
// GetCredits gets the credits property value. A list of users receiving credit for their participation in the security advisory.
// returns a []RepositoryAdvisoryUpdate_creditsable when successful
func (m *RepositoryAdvisoryUpdate) GetCredits()([]RepositoryAdvisoryUpdate_creditsable) {
    return m.credits
}
// GetCveId gets the cve_id property value. The Common Vulnerabilities and Exposures (CVE) ID.
// returns a *string when successful
func (m *RepositoryAdvisoryUpdate) GetCveId()(*string) {
    return m.cve_id
}
// GetCvssVectorString gets the cvss_vector_string property value. The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
// returns a *string when successful
func (m *RepositoryAdvisoryUpdate) GetCvssVectorString()(*string) {
    return m.cvss_vector_string
}
// GetCweIds gets the cwe_ids property value. A list of Common Weakness Enumeration (CWE) IDs.
// returns a []string when successful
func (m *RepositoryAdvisoryUpdate) GetCweIds()([]string) {
    return m.cwe_ids
}
// GetDescription gets the description property value. A detailed description of what the advisory impacts.
// returns a *string when successful
func (m *RepositoryAdvisoryUpdate) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryAdvisoryUpdate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["collaborating_teams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetCollaboratingTeams(res)
        }
        return nil
    }
    res["collaborating_users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetCollaboratingUsers(res)
        }
        return nil
    }
    res["credits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryAdvisoryUpdate_creditsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RepositoryAdvisoryUpdate_creditsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RepositoryAdvisoryUpdate_creditsable)
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
    res["cvss_vector_string"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCvssVectorString(val)
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
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryAdvisoryUpdate_severity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*RepositoryAdvisoryUpdate_severity))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRepositoryAdvisoryUpdate_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*RepositoryAdvisoryUpdate_state))
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
    res["vulnerabilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryAdvisoryUpdate_vulnerabilitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RepositoryAdvisoryUpdate_vulnerabilitiesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RepositoryAdvisoryUpdate_vulnerabilitiesable)
                }
            }
            m.SetVulnerabilities(res)
        }
        return nil
    }
    return res
}
// GetSeverity gets the severity property value. The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
// returns a *RepositoryAdvisoryUpdate_severity when successful
func (m *RepositoryAdvisoryUpdate) GetSeverity()(*RepositoryAdvisoryUpdate_severity) {
    return m.severity
}
// GetState gets the state property value. The state of the advisory.
// returns a *RepositoryAdvisoryUpdate_state when successful
func (m *RepositoryAdvisoryUpdate) GetState()(*RepositoryAdvisoryUpdate_state) {
    return m.state
}
// GetSummary gets the summary property value. A short summary of the advisory.
// returns a *string when successful
func (m *RepositoryAdvisoryUpdate) GetSummary()(*string) {
    return m.summary
}
// GetVulnerabilities gets the vulnerabilities property value. A product affected by the vulnerability detailed in a repository security advisory.
// returns a []RepositoryAdvisoryUpdate_vulnerabilitiesable when successful
func (m *RepositoryAdvisoryUpdate) GetVulnerabilities()([]RepositoryAdvisoryUpdate_vulnerabilitiesable) {
    return m.vulnerabilities
}
// Serialize serializes information the current object
func (m *RepositoryAdvisoryUpdate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCollaboratingTeams() != nil {
        err := writer.WriteCollectionOfStringValues("collaborating_teams", m.GetCollaboratingTeams())
        if err != nil {
            return err
        }
    }
    if m.GetCollaboratingUsers() != nil {
        err := writer.WriteCollectionOfStringValues("collaborating_users", m.GetCollaboratingUsers())
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
        err := writer.WriteStringValue("cvss_vector_string", m.GetCvssVectorString())
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
// SetCollaboratingTeams sets the collaborating_teams property value. A list of team slugs which have been granted write access to the advisory.
func (m *RepositoryAdvisoryUpdate) SetCollaboratingTeams(value []string)() {
    m.collaborating_teams = value
}
// SetCollaboratingUsers sets the collaborating_users property value. A list of usernames who have been granted write access to the advisory.
func (m *RepositoryAdvisoryUpdate) SetCollaboratingUsers(value []string)() {
    m.collaborating_users = value
}
// SetCredits sets the credits property value. A list of users receiving credit for their participation in the security advisory.
func (m *RepositoryAdvisoryUpdate) SetCredits(value []RepositoryAdvisoryUpdate_creditsable)() {
    m.credits = value
}
// SetCveId sets the cve_id property value. The Common Vulnerabilities and Exposures (CVE) ID.
func (m *RepositoryAdvisoryUpdate) SetCveId(value *string)() {
    m.cve_id = value
}
// SetCvssVectorString sets the cvss_vector_string property value. The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
func (m *RepositoryAdvisoryUpdate) SetCvssVectorString(value *string)() {
    m.cvss_vector_string = value
}
// SetCweIds sets the cwe_ids property value. A list of Common Weakness Enumeration (CWE) IDs.
func (m *RepositoryAdvisoryUpdate) SetCweIds(value []string)() {
    m.cwe_ids = value
}
// SetDescription sets the description property value. A detailed description of what the advisory impacts.
func (m *RepositoryAdvisoryUpdate) SetDescription(value *string)() {
    m.description = value
}
// SetSeverity sets the severity property value. The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
func (m *RepositoryAdvisoryUpdate) SetSeverity(value *RepositoryAdvisoryUpdate_severity)() {
    m.severity = value
}
// SetState sets the state property value. The state of the advisory.
func (m *RepositoryAdvisoryUpdate) SetState(value *RepositoryAdvisoryUpdate_state)() {
    m.state = value
}
// SetSummary sets the summary property value. A short summary of the advisory.
func (m *RepositoryAdvisoryUpdate) SetSummary(value *string)() {
    m.summary = value
}
// SetVulnerabilities sets the vulnerabilities property value. A product affected by the vulnerability detailed in a repository security advisory.
func (m *RepositoryAdvisoryUpdate) SetVulnerabilities(value []RepositoryAdvisoryUpdate_vulnerabilitiesable)() {
    m.vulnerabilities = value
}
type RepositoryAdvisoryUpdateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCollaboratingTeams()([]string)
    GetCollaboratingUsers()([]string)
    GetCredits()([]RepositoryAdvisoryUpdate_creditsable)
    GetCveId()(*string)
    GetCvssVectorString()(*string)
    GetCweIds()([]string)
    GetDescription()(*string)
    GetSeverity()(*RepositoryAdvisoryUpdate_severity)
    GetState()(*RepositoryAdvisoryUpdate_state)
    GetSummary()(*string)
    GetVulnerabilities()([]RepositoryAdvisoryUpdate_vulnerabilitiesable)
    SetCollaboratingTeams(value []string)()
    SetCollaboratingUsers(value []string)()
    SetCredits(value []RepositoryAdvisoryUpdate_creditsable)()
    SetCveId(value *string)()
    SetCvssVectorString(value *string)()
    SetCweIds(value []string)()
    SetDescription(value *string)()
    SetSeverity(value *RepositoryAdvisoryUpdate_severity)()
    SetState(value *RepositoryAdvisoryUpdate_state)()
    SetSummary(value *string)()
    SetVulnerabilities(value []RepositoryAdvisoryUpdate_vulnerabilitiesable)()
}
