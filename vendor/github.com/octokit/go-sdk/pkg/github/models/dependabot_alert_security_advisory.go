package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependabotAlertSecurityAdvisory details for the GitHub Security Advisory.
type DependabotAlertSecurityAdvisory struct {
    // The unique CVE ID assigned to the advisory.
    cve_id *string
    // Details for the advisory pertaining to the Common Vulnerability Scoring System.
    cvss DependabotAlertSecurityAdvisory_cvssable
    // Details for the advisory pertaining to Common Weakness Enumeration.
    cwes []DependabotAlertSecurityAdvisory_cwesable
    // A long-form Markdown-supported description of the advisory.
    description *string
    // The unique GitHub Security Advisory ID assigned to the advisory.
    ghsa_id *string
    // Values that identify this advisory among security information sources.
    identifiers []DependabotAlertSecurityAdvisory_identifiersable
    // The time that the advisory was published in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    published_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Links to additional advisory information.
    references []DependabotAlertSecurityAdvisory_referencesable
    // The severity of the advisory.
    severity *DependabotAlertSecurityAdvisory_severity
    // A short, plain text summary of the advisory.
    summary *string
    // The time that the advisory was last modified in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Vulnerable version range information for the advisory.
    vulnerabilities []DependabotAlertSecurityVulnerabilityable
    // The time that the advisory was withdrawn in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
    withdrawn_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewDependabotAlertSecurityAdvisory instantiates a new DependabotAlertSecurityAdvisory and sets the default values.
func NewDependabotAlertSecurityAdvisory()(*DependabotAlertSecurityAdvisory) {
    m := &DependabotAlertSecurityAdvisory{
    }
    return m
}
// CreateDependabotAlertSecurityAdvisoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependabotAlertSecurityAdvisoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependabotAlertSecurityAdvisory(), nil
}
// GetCveId gets the cve_id property value. The unique CVE ID assigned to the advisory.
// returns a *string when successful
func (m *DependabotAlertSecurityAdvisory) GetCveId()(*string) {
    return m.cve_id
}
// GetCvss gets the cvss property value. Details for the advisory pertaining to the Common Vulnerability Scoring System.
// returns a DependabotAlertSecurityAdvisory_cvssable when successful
func (m *DependabotAlertSecurityAdvisory) GetCvss()(DependabotAlertSecurityAdvisory_cvssable) {
    return m.cvss
}
// GetCwes gets the cwes property value. Details for the advisory pertaining to Common Weakness Enumeration.
// returns a []DependabotAlertSecurityAdvisory_cwesable when successful
func (m *DependabotAlertSecurityAdvisory) GetCwes()([]DependabotAlertSecurityAdvisory_cwesable) {
    return m.cwes
}
// GetDescription gets the description property value. A long-form Markdown-supported description of the advisory.
// returns a *string when successful
func (m *DependabotAlertSecurityAdvisory) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependabotAlertSecurityAdvisory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
        val, err := n.GetObjectValue(CreateDependabotAlertSecurityAdvisory_cvssFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCvss(val.(DependabotAlertSecurityAdvisory_cvssable))
        }
        return nil
    }
    res["cwes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDependabotAlertSecurityAdvisory_cwesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DependabotAlertSecurityAdvisory_cwesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DependabotAlertSecurityAdvisory_cwesable)
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
    res["identifiers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDependabotAlertSecurityAdvisory_identifiersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DependabotAlertSecurityAdvisory_identifiersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DependabotAlertSecurityAdvisory_identifiersable)
                }
            }
            m.SetIdentifiers(res)
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
        val, err := n.GetCollectionOfObjectValues(CreateDependabotAlertSecurityAdvisory_referencesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DependabotAlertSecurityAdvisory_referencesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DependabotAlertSecurityAdvisory_referencesable)
                }
            }
            m.SetReferences(res)
        }
        return nil
    }
    res["severity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDependabotAlertSecurityAdvisory_severity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeverity(val.(*DependabotAlertSecurityAdvisory_severity))
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
    res["vulnerabilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDependabotAlertSecurityVulnerabilityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DependabotAlertSecurityVulnerabilityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DependabotAlertSecurityVulnerabilityable)
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
// GetGhsaId gets the ghsa_id property value. The unique GitHub Security Advisory ID assigned to the advisory.
// returns a *string when successful
func (m *DependabotAlertSecurityAdvisory) GetGhsaId()(*string) {
    return m.ghsa_id
}
// GetIdentifiers gets the identifiers property value. Values that identify this advisory among security information sources.
// returns a []DependabotAlertSecurityAdvisory_identifiersable when successful
func (m *DependabotAlertSecurityAdvisory) GetIdentifiers()([]DependabotAlertSecurityAdvisory_identifiersable) {
    return m.identifiers
}
// GetPublishedAt gets the published_at property value. The time that the advisory was published in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *DependabotAlertSecurityAdvisory) GetPublishedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.published_at
}
// GetReferences gets the references property value. Links to additional advisory information.
// returns a []DependabotAlertSecurityAdvisory_referencesable when successful
func (m *DependabotAlertSecurityAdvisory) GetReferences()([]DependabotAlertSecurityAdvisory_referencesable) {
    return m.references
}
// GetSeverity gets the severity property value. The severity of the advisory.
// returns a *DependabotAlertSecurityAdvisory_severity when successful
func (m *DependabotAlertSecurityAdvisory) GetSeverity()(*DependabotAlertSecurityAdvisory_severity) {
    return m.severity
}
// GetSummary gets the summary property value. A short, plain text summary of the advisory.
// returns a *string when successful
func (m *DependabotAlertSecurityAdvisory) GetSummary()(*string) {
    return m.summary
}
// GetUpdatedAt gets the updated_at property value. The time that the advisory was last modified in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *DependabotAlertSecurityAdvisory) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetVulnerabilities gets the vulnerabilities property value. Vulnerable version range information for the advisory.
// returns a []DependabotAlertSecurityVulnerabilityable when successful
func (m *DependabotAlertSecurityAdvisory) GetVulnerabilities()([]DependabotAlertSecurityVulnerabilityable) {
    return m.vulnerabilities
}
// GetWithdrawnAt gets the withdrawn_at property value. The time that the advisory was withdrawn in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *DependabotAlertSecurityAdvisory) GetWithdrawnAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.withdrawn_at
}
// Serialize serializes information the current object
func (m *DependabotAlertSecurityAdvisory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// SetCveId sets the cve_id property value. The unique CVE ID assigned to the advisory.
func (m *DependabotAlertSecurityAdvisory) SetCveId(value *string)() {
    m.cve_id = value
}
// SetCvss sets the cvss property value. Details for the advisory pertaining to the Common Vulnerability Scoring System.
func (m *DependabotAlertSecurityAdvisory) SetCvss(value DependabotAlertSecurityAdvisory_cvssable)() {
    m.cvss = value
}
// SetCwes sets the cwes property value. Details for the advisory pertaining to Common Weakness Enumeration.
func (m *DependabotAlertSecurityAdvisory) SetCwes(value []DependabotAlertSecurityAdvisory_cwesable)() {
    m.cwes = value
}
// SetDescription sets the description property value. A long-form Markdown-supported description of the advisory.
func (m *DependabotAlertSecurityAdvisory) SetDescription(value *string)() {
    m.description = value
}
// SetGhsaId sets the ghsa_id property value. The unique GitHub Security Advisory ID assigned to the advisory.
func (m *DependabotAlertSecurityAdvisory) SetGhsaId(value *string)() {
    m.ghsa_id = value
}
// SetIdentifiers sets the identifiers property value. Values that identify this advisory among security information sources.
func (m *DependabotAlertSecurityAdvisory) SetIdentifiers(value []DependabotAlertSecurityAdvisory_identifiersable)() {
    m.identifiers = value
}
// SetPublishedAt sets the published_at property value. The time that the advisory was published in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *DependabotAlertSecurityAdvisory) SetPublishedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.published_at = value
}
// SetReferences sets the references property value. Links to additional advisory information.
func (m *DependabotAlertSecurityAdvisory) SetReferences(value []DependabotAlertSecurityAdvisory_referencesable)() {
    m.references = value
}
// SetSeverity sets the severity property value. The severity of the advisory.
func (m *DependabotAlertSecurityAdvisory) SetSeverity(value *DependabotAlertSecurityAdvisory_severity)() {
    m.severity = value
}
// SetSummary sets the summary property value. A short, plain text summary of the advisory.
func (m *DependabotAlertSecurityAdvisory) SetSummary(value *string)() {
    m.summary = value
}
// SetUpdatedAt sets the updated_at property value. The time that the advisory was last modified in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *DependabotAlertSecurityAdvisory) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetVulnerabilities sets the vulnerabilities property value. Vulnerable version range information for the advisory.
func (m *DependabotAlertSecurityAdvisory) SetVulnerabilities(value []DependabotAlertSecurityVulnerabilityable)() {
    m.vulnerabilities = value
}
// SetWithdrawnAt sets the withdrawn_at property value. The time that the advisory was withdrawn in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *DependabotAlertSecurityAdvisory) SetWithdrawnAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.withdrawn_at = value
}
type DependabotAlertSecurityAdvisoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCveId()(*string)
    GetCvss()(DependabotAlertSecurityAdvisory_cvssable)
    GetCwes()([]DependabotAlertSecurityAdvisory_cwesable)
    GetDescription()(*string)
    GetGhsaId()(*string)
    GetIdentifiers()([]DependabotAlertSecurityAdvisory_identifiersable)
    GetPublishedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReferences()([]DependabotAlertSecurityAdvisory_referencesable)
    GetSeverity()(*DependabotAlertSecurityAdvisory_severity)
    GetSummary()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVulnerabilities()([]DependabotAlertSecurityVulnerabilityable)
    GetWithdrawnAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCveId(value *string)()
    SetCvss(value DependabotAlertSecurityAdvisory_cvssable)()
    SetCwes(value []DependabotAlertSecurityAdvisory_cwesable)()
    SetDescription(value *string)()
    SetGhsaId(value *string)()
    SetIdentifiers(value []DependabotAlertSecurityAdvisory_identifiersable)()
    SetPublishedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReferences(value []DependabotAlertSecurityAdvisory_referencesable)()
    SetSeverity(value *DependabotAlertSecurityAdvisory_severity)()
    SetSummary(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVulnerabilities(value []DependabotAlertSecurityVulnerabilityable)()
    SetWithdrawnAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
