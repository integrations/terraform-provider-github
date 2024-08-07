package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeSecurityConfiguration a code security configuration
type CodeSecurityConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The enablement status of GitHub Advanced Security
    advanced_security *CodeSecurityConfiguration_advanced_security
    // The enablement status of code scanning default setup
    code_scanning_default_setup *CodeSecurityConfiguration_code_scanning_default_setup
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The enablement status of Dependabot alerts
    dependabot_alerts *CodeSecurityConfiguration_dependabot_alerts
    // The enablement status of Dependabot security updates
    dependabot_security_updates *CodeSecurityConfiguration_dependabot_security_updates
    // The enablement status of Dependency Graph
    dependency_graph *CodeSecurityConfiguration_dependency_graph
    // A description of the code security configuration
    description *string
    // The URL of the configuration
    html_url *string
    // The ID of the code security configuration
    id *int32
    // The name of the code security configuration. Must be unique within the organization.
    name *string
    // The enablement status of private vulnerability reporting
    private_vulnerability_reporting *CodeSecurityConfiguration_private_vulnerability_reporting
    // The enablement status of secret scanning
    secret_scanning *CodeSecurityConfiguration_secret_scanning
    // The enablement status of secret scanning push protection
    secret_scanning_push_protection *CodeSecurityConfiguration_secret_scanning_push_protection
    // The type of the code security configuration.
    target_type *CodeSecurityConfiguration_target_type
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The URL of the configuration
    url *string
}
// NewCodeSecurityConfiguration instantiates a new CodeSecurityConfiguration and sets the default values.
func NewCodeSecurityConfiguration()(*CodeSecurityConfiguration) {
    m := &CodeSecurityConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeSecurityConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeSecurityConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeSecurityConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeSecurityConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdvancedSecurity gets the advanced_security property value. The enablement status of GitHub Advanced Security
// returns a *CodeSecurityConfiguration_advanced_security when successful
func (m *CodeSecurityConfiguration) GetAdvancedSecurity()(*CodeSecurityConfiguration_advanced_security) {
    return m.advanced_security
}
// GetCodeScanningDefaultSetup gets the code_scanning_default_setup property value. The enablement status of code scanning default setup
// returns a *CodeSecurityConfiguration_code_scanning_default_setup when successful
func (m *CodeSecurityConfiguration) GetCodeScanningDefaultSetup()(*CodeSecurityConfiguration_code_scanning_default_setup) {
    return m.code_scanning_default_setup
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *CodeSecurityConfiguration) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetDependabotAlerts gets the dependabot_alerts property value. The enablement status of Dependabot alerts
// returns a *CodeSecurityConfiguration_dependabot_alerts when successful
func (m *CodeSecurityConfiguration) GetDependabotAlerts()(*CodeSecurityConfiguration_dependabot_alerts) {
    return m.dependabot_alerts
}
// GetDependabotSecurityUpdates gets the dependabot_security_updates property value. The enablement status of Dependabot security updates
// returns a *CodeSecurityConfiguration_dependabot_security_updates when successful
func (m *CodeSecurityConfiguration) GetDependabotSecurityUpdates()(*CodeSecurityConfiguration_dependabot_security_updates) {
    return m.dependabot_security_updates
}
// GetDependencyGraph gets the dependency_graph property value. The enablement status of Dependency Graph
// returns a *CodeSecurityConfiguration_dependency_graph when successful
func (m *CodeSecurityConfiguration) GetDependencyGraph()(*CodeSecurityConfiguration_dependency_graph) {
    return m.dependency_graph
}
// GetDescription gets the description property value. A description of the code security configuration
// returns a *string when successful
func (m *CodeSecurityConfiguration) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeSecurityConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["advanced_security"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_advanced_security)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedSecurity(val.(*CodeSecurityConfiguration_advanced_security))
        }
        return nil
    }
    res["code_scanning_default_setup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_code_scanning_default_setup)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeScanningDefaultSetup(val.(*CodeSecurityConfiguration_code_scanning_default_setup))
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
    res["dependabot_alerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_dependabot_alerts)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependabotAlerts(val.(*CodeSecurityConfiguration_dependabot_alerts))
        }
        return nil
    }
    res["dependabot_security_updates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_dependabot_security_updates)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependabotSecurityUpdates(val.(*CodeSecurityConfiguration_dependabot_security_updates))
        }
        return nil
    }
    res["dependency_graph"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_dependency_graph)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependencyGraph(val.(*CodeSecurityConfiguration_dependency_graph))
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["private_vulnerability_reporting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_private_vulnerability_reporting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateVulnerabilityReporting(val.(*CodeSecurityConfiguration_private_vulnerability_reporting))
        }
        return nil
    }
    res["secret_scanning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_secret_scanning)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanning(val.(*CodeSecurityConfiguration_secret_scanning))
        }
        return nil
    }
    res["secret_scanning_push_protection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_secret_scanning_push_protection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningPushProtection(val.(*CodeSecurityConfiguration_secret_scanning_push_protection))
        }
        return nil
    }
    res["target_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeSecurityConfiguration_target_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetType(val.(*CodeSecurityConfiguration_target_type))
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
    return res
}
// GetHtmlUrl gets the html_url property value. The URL of the configuration
// returns a *string when successful
func (m *CodeSecurityConfiguration) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The ID of the code security configuration
// returns a *int32 when successful
func (m *CodeSecurityConfiguration) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name of the code security configuration. Must be unique within the organization.
// returns a *string when successful
func (m *CodeSecurityConfiguration) GetName()(*string) {
    return m.name
}
// GetPrivateVulnerabilityReporting gets the private_vulnerability_reporting property value. The enablement status of private vulnerability reporting
// returns a *CodeSecurityConfiguration_private_vulnerability_reporting when successful
func (m *CodeSecurityConfiguration) GetPrivateVulnerabilityReporting()(*CodeSecurityConfiguration_private_vulnerability_reporting) {
    return m.private_vulnerability_reporting
}
// GetSecretScanning gets the secret_scanning property value. The enablement status of secret scanning
// returns a *CodeSecurityConfiguration_secret_scanning when successful
func (m *CodeSecurityConfiguration) GetSecretScanning()(*CodeSecurityConfiguration_secret_scanning) {
    return m.secret_scanning
}
// GetSecretScanningPushProtection gets the secret_scanning_push_protection property value. The enablement status of secret scanning push protection
// returns a *CodeSecurityConfiguration_secret_scanning_push_protection when successful
func (m *CodeSecurityConfiguration) GetSecretScanningPushProtection()(*CodeSecurityConfiguration_secret_scanning_push_protection) {
    return m.secret_scanning_push_protection
}
// GetTargetType gets the target_type property value. The type of the code security configuration.
// returns a *CodeSecurityConfiguration_target_type when successful
func (m *CodeSecurityConfiguration) GetTargetType()(*CodeSecurityConfiguration_target_type) {
    return m.target_type
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *CodeSecurityConfiguration) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// GetUrl gets the url property value. The URL of the configuration
// returns a *string when successful
func (m *CodeSecurityConfiguration) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *CodeSecurityConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAdvancedSecurity() != nil {
        cast := (*m.GetAdvancedSecurity()).String()
        err := writer.WriteStringValue("advanced_security", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCodeScanningDefaultSetup() != nil {
        cast := (*m.GetCodeScanningDefaultSetup()).String()
        err := writer.WriteStringValue("code_scanning_default_setup", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    if m.GetDependabotAlerts() != nil {
        cast := (*m.GetDependabotAlerts()).String()
        err := writer.WriteStringValue("dependabot_alerts", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDependabotSecurityUpdates() != nil {
        cast := (*m.GetDependabotSecurityUpdates()).String()
        err := writer.WriteStringValue("dependabot_security_updates", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDependencyGraph() != nil {
        cast := (*m.GetDependencyGraph()).String()
        err := writer.WriteStringValue("dependency_graph", &cast)
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
    {
        err := writer.WriteStringValue("html_url", m.GetHtmlUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
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
    if m.GetPrivateVulnerabilityReporting() != nil {
        cast := (*m.GetPrivateVulnerabilityReporting()).String()
        err := writer.WriteStringValue("private_vulnerability_reporting", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecretScanning() != nil {
        cast := (*m.GetSecretScanning()).String()
        err := writer.WriteStringValue("secret_scanning", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecretScanningPushProtection() != nil {
        cast := (*m.GetSecretScanningPushProtection()).String()
        err := writer.WriteStringValue("secret_scanning_push_protection", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargetType() != nil {
        cast := (*m.GetTargetType()).String()
        err := writer.WriteStringValue("target_type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
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
func (m *CodeSecurityConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdvancedSecurity sets the advanced_security property value. The enablement status of GitHub Advanced Security
func (m *CodeSecurityConfiguration) SetAdvancedSecurity(value *CodeSecurityConfiguration_advanced_security)() {
    m.advanced_security = value
}
// SetCodeScanningDefaultSetup sets the code_scanning_default_setup property value. The enablement status of code scanning default setup
func (m *CodeSecurityConfiguration) SetCodeScanningDefaultSetup(value *CodeSecurityConfiguration_code_scanning_default_setup)() {
    m.code_scanning_default_setup = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *CodeSecurityConfiguration) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetDependabotAlerts sets the dependabot_alerts property value. The enablement status of Dependabot alerts
func (m *CodeSecurityConfiguration) SetDependabotAlerts(value *CodeSecurityConfiguration_dependabot_alerts)() {
    m.dependabot_alerts = value
}
// SetDependabotSecurityUpdates sets the dependabot_security_updates property value. The enablement status of Dependabot security updates
func (m *CodeSecurityConfiguration) SetDependabotSecurityUpdates(value *CodeSecurityConfiguration_dependabot_security_updates)() {
    m.dependabot_security_updates = value
}
// SetDependencyGraph sets the dependency_graph property value. The enablement status of Dependency Graph
func (m *CodeSecurityConfiguration) SetDependencyGraph(value *CodeSecurityConfiguration_dependency_graph)() {
    m.dependency_graph = value
}
// SetDescription sets the description property value. A description of the code security configuration
func (m *CodeSecurityConfiguration) SetDescription(value *string)() {
    m.description = value
}
// SetHtmlUrl sets the html_url property value. The URL of the configuration
func (m *CodeSecurityConfiguration) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The ID of the code security configuration
func (m *CodeSecurityConfiguration) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name of the code security configuration. Must be unique within the organization.
func (m *CodeSecurityConfiguration) SetName(value *string)() {
    m.name = value
}
// SetPrivateVulnerabilityReporting sets the private_vulnerability_reporting property value. The enablement status of private vulnerability reporting
func (m *CodeSecurityConfiguration) SetPrivateVulnerabilityReporting(value *CodeSecurityConfiguration_private_vulnerability_reporting)() {
    m.private_vulnerability_reporting = value
}
// SetSecretScanning sets the secret_scanning property value. The enablement status of secret scanning
func (m *CodeSecurityConfiguration) SetSecretScanning(value *CodeSecurityConfiguration_secret_scanning)() {
    m.secret_scanning = value
}
// SetSecretScanningPushProtection sets the secret_scanning_push_protection property value. The enablement status of secret scanning push protection
func (m *CodeSecurityConfiguration) SetSecretScanningPushProtection(value *CodeSecurityConfiguration_secret_scanning_push_protection)() {
    m.secret_scanning_push_protection = value
}
// SetTargetType sets the target_type property value. The type of the code security configuration.
func (m *CodeSecurityConfiguration) SetTargetType(value *CodeSecurityConfiguration_target_type)() {
    m.target_type = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *CodeSecurityConfiguration) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
// SetUrl sets the url property value. The URL of the configuration
func (m *CodeSecurityConfiguration) SetUrl(value *string)() {
    m.url = value
}
type CodeSecurityConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedSecurity()(*CodeSecurityConfiguration_advanced_security)
    GetCodeScanningDefaultSetup()(*CodeSecurityConfiguration_code_scanning_default_setup)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDependabotAlerts()(*CodeSecurityConfiguration_dependabot_alerts)
    GetDependabotSecurityUpdates()(*CodeSecurityConfiguration_dependabot_security_updates)
    GetDependencyGraph()(*CodeSecurityConfiguration_dependency_graph)
    GetDescription()(*string)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetPrivateVulnerabilityReporting()(*CodeSecurityConfiguration_private_vulnerability_reporting)
    GetSecretScanning()(*CodeSecurityConfiguration_secret_scanning)
    GetSecretScanningPushProtection()(*CodeSecurityConfiguration_secret_scanning_push_protection)
    GetTargetType()(*CodeSecurityConfiguration_target_type)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUrl()(*string)
    SetAdvancedSecurity(value *CodeSecurityConfiguration_advanced_security)()
    SetCodeScanningDefaultSetup(value *CodeSecurityConfiguration_code_scanning_default_setup)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDependabotAlerts(value *CodeSecurityConfiguration_dependabot_alerts)()
    SetDependabotSecurityUpdates(value *CodeSecurityConfiguration_dependabot_security_updates)()
    SetDependencyGraph(value *CodeSecurityConfiguration_dependency_graph)()
    SetDescription(value *string)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetPrivateVulnerabilityReporting(value *CodeSecurityConfiguration_private_vulnerability_reporting)()
    SetSecretScanning(value *CodeSecurityConfiguration_secret_scanning)()
    SetSecretScanningPushProtection(value *CodeSecurityConfiguration_secret_scanning_push_protection)()
    SetTargetType(value *CodeSecurityConfiguration_target_type)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUrl(value *string)()
}
