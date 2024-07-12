package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PagesHealthCheck_domain struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The caa_error property
    caa_error *string
    // The dns_resolves property
    dns_resolves *bool
    // The enforces_https property
    enforces_https *bool
    // The has_cname_record property
    has_cname_record *bool
    // The has_mx_records_present property
    has_mx_records_present *bool
    // The host property
    host *string
    // The https_error property
    https_error *string
    // The is_a_record property
    is_a_record *bool
    // The is_apex_domain property
    is_apex_domain *bool
    // The is_cloudflare_ip property
    is_cloudflare_ip *bool
    // The is_cname_to_fastly property
    is_cname_to_fastly *bool
    // The is_cname_to_github_user_domain property
    is_cname_to_github_user_domain *bool
    // The is_cname_to_pages_dot_github_dot_com property
    is_cname_to_pages_dot_github_dot_com *bool
    // The is_fastly_ip property
    is_fastly_ip *bool
    // The is_https_eligible property
    is_https_eligible *bool
    // The is_non_github_pages_ip_present property
    is_non_github_pages_ip_present *bool
    // The is_old_ip_address property
    is_old_ip_address *bool
    // The is_pages_domain property
    is_pages_domain *bool
    // The is_pointed_to_github_pages_ip property
    is_pointed_to_github_pages_ip *bool
    // The is_proxied property
    is_proxied *bool
    // The is_served_by_pages property
    is_served_by_pages *bool
    // The is_valid property
    is_valid *bool
    // The is_valid_domain property
    is_valid_domain *bool
    // The nameservers property
    nameservers *string
    // The reason property
    reason *string
    // The responds_to_https property
    responds_to_https *bool
    // The should_be_a_record property
    should_be_a_record *bool
    // The uri property
    uri *string
}
// NewPagesHealthCheck_domain instantiates a new PagesHealthCheck_domain and sets the default values.
func NewPagesHealthCheck_domain()(*PagesHealthCheck_domain) {
    m := &PagesHealthCheck_domain{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePagesHealthCheck_domainFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePagesHealthCheck_domainFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPagesHealthCheck_domain(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PagesHealthCheck_domain) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCaaError gets the caa_error property value. The caa_error property
// returns a *string when successful
func (m *PagesHealthCheck_domain) GetCaaError()(*string) {
    return m.caa_error
}
// GetDnsResolves gets the dns_resolves property value. The dns_resolves property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetDnsResolves()(*bool) {
    return m.dns_resolves
}
// GetEnforcesHttps gets the enforces_https property value. The enforces_https property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetEnforcesHttps()(*bool) {
    return m.enforces_https
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PagesHealthCheck_domain) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["caa_error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaaError(val)
        }
        return nil
    }
    res["dns_resolves"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDnsResolves(val)
        }
        return nil
    }
    res["enforces_https"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforcesHttps(val)
        }
        return nil
    }
    res["has_cname_record"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasCnameRecord(val)
        }
        return nil
    }
    res["has_mx_records_present"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasMxRecordsPresent(val)
        }
        return nil
    }
    res["host"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHost(val)
        }
        return nil
    }
    res["https_error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHttpsError(val)
        }
        return nil
    }
    res["is_a_record"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsARecord(val)
        }
        return nil
    }
    res["is_apex_domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApexDomain(val)
        }
        return nil
    }
    res["is_cloudflare_ip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCloudflareIp(val)
        }
        return nil
    }
    res["is_cname_to_fastly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCnameToFastly(val)
        }
        return nil
    }
    res["is_cname_to_github_user_domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCnameToGithubUserDomain(val)
        }
        return nil
    }
    res["is_cname_to_pages_dot_github_dot_com"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCnameToPagesDotGithubDotCom(val)
        }
        return nil
    }
    res["is_fastly_ip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFastlyIp(val)
        }
        return nil
    }
    res["is_https_eligible"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHttpsEligible(val)
        }
        return nil
    }
    res["is_non_github_pages_ip_present"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsNonGithubPagesIpPresent(val)
        }
        return nil
    }
    res["is_old_ip_address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOldIpAddress(val)
        }
        return nil
    }
    res["is_pages_domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPagesDomain(val)
        }
        return nil
    }
    res["is_pointed_to_github_pages_ip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPointedToGithubPagesIp(val)
        }
        return nil
    }
    res["is_proxied"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsProxied(val)
        }
        return nil
    }
    res["is_served_by_pages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsServedByPages(val)
        }
        return nil
    }
    res["is_valid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValid(val)
        }
        return nil
    }
    res["is_valid_domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValidDomain(val)
        }
        return nil
    }
    res["nameservers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNameservers(val)
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    res["responds_to_https"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRespondsToHttps(val)
        }
        return nil
    }
    res["should_be_a_record"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShouldBeARecord(val)
        }
        return nil
    }
    res["uri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUri(val)
        }
        return nil
    }
    return res
}
// GetHasCnameRecord gets the has_cname_record property value. The has_cname_record property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetHasCnameRecord()(*bool) {
    return m.has_cname_record
}
// GetHasMxRecordsPresent gets the has_mx_records_present property value. The has_mx_records_present property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetHasMxRecordsPresent()(*bool) {
    return m.has_mx_records_present
}
// GetHost gets the host property value. The host property
// returns a *string when successful
func (m *PagesHealthCheck_domain) GetHost()(*string) {
    return m.host
}
// GetHttpsError gets the https_error property value. The https_error property
// returns a *string when successful
func (m *PagesHealthCheck_domain) GetHttpsError()(*string) {
    return m.https_error
}
// GetIsApexDomain gets the is_apex_domain property value. The is_apex_domain property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsApexDomain()(*bool) {
    return m.is_apex_domain
}
// GetIsARecord gets the is_a_record property value. The is_a_record property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsARecord()(*bool) {
    return m.is_a_record
}
// GetIsCloudflareIp gets the is_cloudflare_ip property value. The is_cloudflare_ip property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsCloudflareIp()(*bool) {
    return m.is_cloudflare_ip
}
// GetIsCnameToFastly gets the is_cname_to_fastly property value. The is_cname_to_fastly property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsCnameToFastly()(*bool) {
    return m.is_cname_to_fastly
}
// GetIsCnameToGithubUserDomain gets the is_cname_to_github_user_domain property value. The is_cname_to_github_user_domain property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsCnameToGithubUserDomain()(*bool) {
    return m.is_cname_to_github_user_domain
}
// GetIsCnameToPagesDotGithubDotCom gets the is_cname_to_pages_dot_github_dot_com property value. The is_cname_to_pages_dot_github_dot_com property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsCnameToPagesDotGithubDotCom()(*bool) {
    return m.is_cname_to_pages_dot_github_dot_com
}
// GetIsFastlyIp gets the is_fastly_ip property value. The is_fastly_ip property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsFastlyIp()(*bool) {
    return m.is_fastly_ip
}
// GetIsHttpsEligible gets the is_https_eligible property value. The is_https_eligible property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsHttpsEligible()(*bool) {
    return m.is_https_eligible
}
// GetIsNonGithubPagesIpPresent gets the is_non_github_pages_ip_present property value. The is_non_github_pages_ip_present property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsNonGithubPagesIpPresent()(*bool) {
    return m.is_non_github_pages_ip_present
}
// GetIsOldIpAddress gets the is_old_ip_address property value. The is_old_ip_address property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsOldIpAddress()(*bool) {
    return m.is_old_ip_address
}
// GetIsPagesDomain gets the is_pages_domain property value. The is_pages_domain property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsPagesDomain()(*bool) {
    return m.is_pages_domain
}
// GetIsPointedToGithubPagesIp gets the is_pointed_to_github_pages_ip property value. The is_pointed_to_github_pages_ip property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsPointedToGithubPagesIp()(*bool) {
    return m.is_pointed_to_github_pages_ip
}
// GetIsProxied gets the is_proxied property value. The is_proxied property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsProxied()(*bool) {
    return m.is_proxied
}
// GetIsServedByPages gets the is_served_by_pages property value. The is_served_by_pages property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsServedByPages()(*bool) {
    return m.is_served_by_pages
}
// GetIsValid gets the is_valid property value. The is_valid property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsValid()(*bool) {
    return m.is_valid
}
// GetIsValidDomain gets the is_valid_domain property value. The is_valid_domain property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetIsValidDomain()(*bool) {
    return m.is_valid_domain
}
// GetNameservers gets the nameservers property value. The nameservers property
// returns a *string when successful
func (m *PagesHealthCheck_domain) GetNameservers()(*string) {
    return m.nameservers
}
// GetReason gets the reason property value. The reason property
// returns a *string when successful
func (m *PagesHealthCheck_domain) GetReason()(*string) {
    return m.reason
}
// GetRespondsToHttps gets the responds_to_https property value. The responds_to_https property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetRespondsToHttps()(*bool) {
    return m.responds_to_https
}
// GetShouldBeARecord gets the should_be_a_record property value. The should_be_a_record property
// returns a *bool when successful
func (m *PagesHealthCheck_domain) GetShouldBeARecord()(*bool) {
    return m.should_be_a_record
}
// GetUri gets the uri property value. The uri property
// returns a *string when successful
func (m *PagesHealthCheck_domain) GetUri()(*string) {
    return m.uri
}
// Serialize serializes information the current object
func (m *PagesHealthCheck_domain) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("caa_error", m.GetCaaError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("dns_resolves", m.GetDnsResolves())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enforces_https", m.GetEnforcesHttps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_cname_record", m.GetHasCnameRecord())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_mx_records_present", m.GetHasMxRecordsPresent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("host", m.GetHost())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("https_error", m.GetHttpsError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_apex_domain", m.GetIsApexDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_a_record", m.GetIsARecord())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_cloudflare_ip", m.GetIsCloudflareIp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_cname_to_fastly", m.GetIsCnameToFastly())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_cname_to_github_user_domain", m.GetIsCnameToGithubUserDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_cname_to_pages_dot_github_dot_com", m.GetIsCnameToPagesDotGithubDotCom())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_fastly_ip", m.GetIsFastlyIp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_https_eligible", m.GetIsHttpsEligible())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_non_github_pages_ip_present", m.GetIsNonGithubPagesIpPresent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_old_ip_address", m.GetIsOldIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_pages_domain", m.GetIsPagesDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_pointed_to_github_pages_ip", m.GetIsPointedToGithubPagesIp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_proxied", m.GetIsProxied())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_served_by_pages", m.GetIsServedByPages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_valid", m.GetIsValid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("is_valid_domain", m.GetIsValidDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("nameservers", m.GetNameservers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("responds_to_https", m.GetRespondsToHttps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("should_be_a_record", m.GetShouldBeARecord())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uri", m.GetUri())
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
func (m *PagesHealthCheck_domain) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCaaError sets the caa_error property value. The caa_error property
func (m *PagesHealthCheck_domain) SetCaaError(value *string)() {
    m.caa_error = value
}
// SetDnsResolves sets the dns_resolves property value. The dns_resolves property
func (m *PagesHealthCheck_domain) SetDnsResolves(value *bool)() {
    m.dns_resolves = value
}
// SetEnforcesHttps sets the enforces_https property value. The enforces_https property
func (m *PagesHealthCheck_domain) SetEnforcesHttps(value *bool)() {
    m.enforces_https = value
}
// SetHasCnameRecord sets the has_cname_record property value. The has_cname_record property
func (m *PagesHealthCheck_domain) SetHasCnameRecord(value *bool)() {
    m.has_cname_record = value
}
// SetHasMxRecordsPresent sets the has_mx_records_present property value. The has_mx_records_present property
func (m *PagesHealthCheck_domain) SetHasMxRecordsPresent(value *bool)() {
    m.has_mx_records_present = value
}
// SetHost sets the host property value. The host property
func (m *PagesHealthCheck_domain) SetHost(value *string)() {
    m.host = value
}
// SetHttpsError sets the https_error property value. The https_error property
func (m *PagesHealthCheck_domain) SetHttpsError(value *string)() {
    m.https_error = value
}
// SetIsApexDomain sets the is_apex_domain property value. The is_apex_domain property
func (m *PagesHealthCheck_domain) SetIsApexDomain(value *bool)() {
    m.is_apex_domain = value
}
// SetIsARecord sets the is_a_record property value. The is_a_record property
func (m *PagesHealthCheck_domain) SetIsARecord(value *bool)() {
    m.is_a_record = value
}
// SetIsCloudflareIp sets the is_cloudflare_ip property value. The is_cloudflare_ip property
func (m *PagesHealthCheck_domain) SetIsCloudflareIp(value *bool)() {
    m.is_cloudflare_ip = value
}
// SetIsCnameToFastly sets the is_cname_to_fastly property value. The is_cname_to_fastly property
func (m *PagesHealthCheck_domain) SetIsCnameToFastly(value *bool)() {
    m.is_cname_to_fastly = value
}
// SetIsCnameToGithubUserDomain sets the is_cname_to_github_user_domain property value. The is_cname_to_github_user_domain property
func (m *PagesHealthCheck_domain) SetIsCnameToGithubUserDomain(value *bool)() {
    m.is_cname_to_github_user_domain = value
}
// SetIsCnameToPagesDotGithubDotCom sets the is_cname_to_pages_dot_github_dot_com property value. The is_cname_to_pages_dot_github_dot_com property
func (m *PagesHealthCheck_domain) SetIsCnameToPagesDotGithubDotCom(value *bool)() {
    m.is_cname_to_pages_dot_github_dot_com = value
}
// SetIsFastlyIp sets the is_fastly_ip property value. The is_fastly_ip property
func (m *PagesHealthCheck_domain) SetIsFastlyIp(value *bool)() {
    m.is_fastly_ip = value
}
// SetIsHttpsEligible sets the is_https_eligible property value. The is_https_eligible property
func (m *PagesHealthCheck_domain) SetIsHttpsEligible(value *bool)() {
    m.is_https_eligible = value
}
// SetIsNonGithubPagesIpPresent sets the is_non_github_pages_ip_present property value. The is_non_github_pages_ip_present property
func (m *PagesHealthCheck_domain) SetIsNonGithubPagesIpPresent(value *bool)() {
    m.is_non_github_pages_ip_present = value
}
// SetIsOldIpAddress sets the is_old_ip_address property value. The is_old_ip_address property
func (m *PagesHealthCheck_domain) SetIsOldIpAddress(value *bool)() {
    m.is_old_ip_address = value
}
// SetIsPagesDomain sets the is_pages_domain property value. The is_pages_domain property
func (m *PagesHealthCheck_domain) SetIsPagesDomain(value *bool)() {
    m.is_pages_domain = value
}
// SetIsPointedToGithubPagesIp sets the is_pointed_to_github_pages_ip property value. The is_pointed_to_github_pages_ip property
func (m *PagesHealthCheck_domain) SetIsPointedToGithubPagesIp(value *bool)() {
    m.is_pointed_to_github_pages_ip = value
}
// SetIsProxied sets the is_proxied property value. The is_proxied property
func (m *PagesHealthCheck_domain) SetIsProxied(value *bool)() {
    m.is_proxied = value
}
// SetIsServedByPages sets the is_served_by_pages property value. The is_served_by_pages property
func (m *PagesHealthCheck_domain) SetIsServedByPages(value *bool)() {
    m.is_served_by_pages = value
}
// SetIsValid sets the is_valid property value. The is_valid property
func (m *PagesHealthCheck_domain) SetIsValid(value *bool)() {
    m.is_valid = value
}
// SetIsValidDomain sets the is_valid_domain property value. The is_valid_domain property
func (m *PagesHealthCheck_domain) SetIsValidDomain(value *bool)() {
    m.is_valid_domain = value
}
// SetNameservers sets the nameservers property value. The nameservers property
func (m *PagesHealthCheck_domain) SetNameservers(value *string)() {
    m.nameservers = value
}
// SetReason sets the reason property value. The reason property
func (m *PagesHealthCheck_domain) SetReason(value *string)() {
    m.reason = value
}
// SetRespondsToHttps sets the responds_to_https property value. The responds_to_https property
func (m *PagesHealthCheck_domain) SetRespondsToHttps(value *bool)() {
    m.responds_to_https = value
}
// SetShouldBeARecord sets the should_be_a_record property value. The should_be_a_record property
func (m *PagesHealthCheck_domain) SetShouldBeARecord(value *bool)() {
    m.should_be_a_record = value
}
// SetUri sets the uri property value. The uri property
func (m *PagesHealthCheck_domain) SetUri(value *string)() {
    m.uri = value
}
type PagesHealthCheck_domainable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCaaError()(*string)
    GetDnsResolves()(*bool)
    GetEnforcesHttps()(*bool)
    GetHasCnameRecord()(*bool)
    GetHasMxRecordsPresent()(*bool)
    GetHost()(*string)
    GetHttpsError()(*string)
    GetIsApexDomain()(*bool)
    GetIsARecord()(*bool)
    GetIsCloudflareIp()(*bool)
    GetIsCnameToFastly()(*bool)
    GetIsCnameToGithubUserDomain()(*bool)
    GetIsCnameToPagesDotGithubDotCom()(*bool)
    GetIsFastlyIp()(*bool)
    GetIsHttpsEligible()(*bool)
    GetIsNonGithubPagesIpPresent()(*bool)
    GetIsOldIpAddress()(*bool)
    GetIsPagesDomain()(*bool)
    GetIsPointedToGithubPagesIp()(*bool)
    GetIsProxied()(*bool)
    GetIsServedByPages()(*bool)
    GetIsValid()(*bool)
    GetIsValidDomain()(*bool)
    GetNameservers()(*string)
    GetReason()(*string)
    GetRespondsToHttps()(*bool)
    GetShouldBeARecord()(*bool)
    GetUri()(*string)
    SetCaaError(value *string)()
    SetDnsResolves(value *bool)()
    SetEnforcesHttps(value *bool)()
    SetHasCnameRecord(value *bool)()
    SetHasMxRecordsPresent(value *bool)()
    SetHost(value *string)()
    SetHttpsError(value *string)()
    SetIsApexDomain(value *bool)()
    SetIsARecord(value *bool)()
    SetIsCloudflareIp(value *bool)()
    SetIsCnameToFastly(value *bool)()
    SetIsCnameToGithubUserDomain(value *bool)()
    SetIsCnameToPagesDotGithubDotCom(value *bool)()
    SetIsFastlyIp(value *bool)()
    SetIsHttpsEligible(value *bool)()
    SetIsNonGithubPagesIpPresent(value *bool)()
    SetIsOldIpAddress(value *bool)()
    SetIsPagesDomain(value *bool)()
    SetIsPointedToGithubPagesIp(value *bool)()
    SetIsProxied(value *bool)()
    SetIsServedByPages(value *bool)()
    SetIsValid(value *bool)()
    SetIsValidDomain(value *bool)()
    SetNameservers(value *string)()
    SetReason(value *string)()
    SetRespondsToHttps(value *bool)()
    SetShouldBeARecord(value *bool)()
    SetUri(value *string)()
}
