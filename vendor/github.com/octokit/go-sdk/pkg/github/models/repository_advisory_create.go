package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryAdvisoryCreate struct {
    // A list of users receiving credit for their participation in the security advisory.
    credits []RepositoryAdvisoryCreate_creditsable
    // The Common Vulnerabilities and Exposures (CVE) ID.
    cve_id *string
    // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
    cvss_vector_string *string
    // A list of Common Weakness Enumeration (CWE) IDs.
    cwe_ids []string
    // A detailed description of what the advisory impacts.
    description *string
    // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
    severity *RepositoryAdvisoryCreate_severity
    // Whether to create a temporary private fork of the repository to collaborate on a fix.
    start_private_fork *bool
    // A short summary of the advisory.
    summary *string
    // A product affected by the vulnerability detailed in a repository security advisory.
    vulnerabilities []RepositoryAdvisoryCreate_vulnerabilitiesable
}
// NewRepositoryAdvisoryCreate instantiates a new RepositoryAdvisoryCreate and sets the default values.
func NewRepositoryAdvisoryCreate()(*RepositoryAdvisoryCreate) {
    m := &RepositoryAdvisoryCreate{
    }
    return m
}
// CreateRepositoryAdvisoryCreateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryAdvisoryCreateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryAdvisoryCreate(), nil
}
// GetCredits gets the credits property value. A list of users receiving credit for their participation in the security advisory.
// returns a []RepositoryAdvisoryCreate_creditsable when successful
func (m *RepositoryAdvisoryCreate) GetCredits()([]RepositoryAdvisoryCreate_creditsable) {
    return m.credits
}
// GetCveId gets the cve_id property value. The Common Vulnerabilities and Exposures (CVE) ID.
// returns a *string when successful
func (m *RepositoryAdvisoryCreate) GetCveId()(*string) {
    return m.cve_id
}
// GetCvssVectorString gets the cvss_vector_string property value. The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
// returns a *string when successful
func (m *RepositoryAdvisoryCreate) GetCvssVectorString()(*string) {
    return m.cvss_vector_string
}
// GetCweIds gets the cwe_ids property value. A list of Common Weakness Enumeration (CWE) IDs.
// returns a []string when successful
func (m *RepositoryAdvisoryCreate) GetCweIds()([]string) {
    return m.cwe_ids
}
// GetDescription gets the description property value. A detailed description of what the advisory impacts.
// returns a *string when successful
func (m *RepositoryAdvisoryCreate) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryAdvisoryCreate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["credits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryAdvisoryCreate_creditsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RepositoryAdvisoryCreate_creditsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RepositoryAdvisoryCreate_creditsable)
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
        val, err := n.GetEnumValue(ParseRepositoryAdvisoryCreate_severity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*RepositoryAdvisoryCreate_severity))
        }
        return nil
    }
    res["start_private_fork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartPrivateFork(val)
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
        val, err := n.GetCollectionOfObjectValues(CreateRepositoryAdvisoryCreate_vulnerabilitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RepositoryAdvisoryCreate_vulnerabilitiesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RepositoryAdvisoryCreate_vulnerabilitiesable)
                }
            }
            m.SetVulnerabilities(res)
        }
        return nil
    }
    return res
}
// GetSeverity gets the severity property value. The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
// returns a *RepositoryAdvisoryCreate_severity when successful
func (m *RepositoryAdvisoryCreate) GetSeverity()(*RepositoryAdvisoryCreate_severity) {
    return m.severity
}
// GetStartPrivateFork gets the start_private_fork property value. Whether to create a temporary private fork of the repository to collaborate on a fix.
// returns a *bool when successful
func (m *RepositoryAdvisoryCreate) GetStartPrivateFork()(*bool) {
    return m.start_private_fork
}
// GetSummary gets the summary property value. A short summary of the advisory.
// returns a *string when successful
func (m *RepositoryAdvisoryCreate) GetSummary()(*string) {
    return m.summary
}
// GetVulnerabilities gets the vulnerabilities property value. A product affected by the vulnerability detailed in a repository security advisory.
// returns a []RepositoryAdvisoryCreate_vulnerabilitiesable when successful
func (m *RepositoryAdvisoryCreate) GetVulnerabilities()([]RepositoryAdvisoryCreate_vulnerabilitiesable) {
    return m.vulnerabilities
}
// Serialize serializes information the current object
func (m *RepositoryAdvisoryCreate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteBoolValue("start_private_fork", m.GetStartPrivateFork())
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
// SetCredits sets the credits property value. A list of users receiving credit for their participation in the security advisory.
func (m *RepositoryAdvisoryCreate) SetCredits(value []RepositoryAdvisoryCreate_creditsable)() {
    m.credits = value
}
// SetCveId sets the cve_id property value. The Common Vulnerabilities and Exposures (CVE) ID.
func (m *RepositoryAdvisoryCreate) SetCveId(value *string)() {
    m.cve_id = value
}
// SetCvssVectorString sets the cvss_vector_string property value. The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
func (m *RepositoryAdvisoryCreate) SetCvssVectorString(value *string)() {
    m.cvss_vector_string = value
}
// SetCweIds sets the cwe_ids property value. A list of Common Weakness Enumeration (CWE) IDs.
func (m *RepositoryAdvisoryCreate) SetCweIds(value []string)() {
    m.cwe_ids = value
}
// SetDescription sets the description property value. A detailed description of what the advisory impacts.
func (m *RepositoryAdvisoryCreate) SetDescription(value *string)() {
    m.description = value
}
// SetSeverity sets the severity property value. The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
func (m *RepositoryAdvisoryCreate) SetSeverity(value *RepositoryAdvisoryCreate_severity)() {
    m.severity = value
}
// SetStartPrivateFork sets the start_private_fork property value. Whether to create a temporary private fork of the repository to collaborate on a fix.
func (m *RepositoryAdvisoryCreate) SetStartPrivateFork(value *bool)() {
    m.start_private_fork = value
}
// SetSummary sets the summary property value. A short summary of the advisory.
func (m *RepositoryAdvisoryCreate) SetSummary(value *string)() {
    m.summary = value
}
// SetVulnerabilities sets the vulnerabilities property value. A product affected by the vulnerability detailed in a repository security advisory.
func (m *RepositoryAdvisoryCreate) SetVulnerabilities(value []RepositoryAdvisoryCreate_vulnerabilitiesable)() {
    m.vulnerabilities = value
}
type RepositoryAdvisoryCreateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCredits()([]RepositoryAdvisoryCreate_creditsable)
    GetCveId()(*string)
    GetCvssVectorString()(*string)
    GetCweIds()([]string)
    GetDescription()(*string)
    GetSeverity()(*RepositoryAdvisoryCreate_severity)
    GetStartPrivateFork()(*bool)
    GetSummary()(*string)
    GetVulnerabilities()([]RepositoryAdvisoryCreate_vulnerabilitiesable)
    SetCredits(value []RepositoryAdvisoryCreate_creditsable)()
    SetCveId(value *string)()
    SetCvssVectorString(value *string)()
    SetCweIds(value []string)()
    SetDescription(value *string)()
    SetSeverity(value *RepositoryAdvisoryCreate_severity)()
    SetStartPrivateFork(value *bool)()
    SetSummary(value *string)()
    SetVulnerabilities(value []RepositoryAdvisoryCreate_vulnerabilitiesable)()
}
