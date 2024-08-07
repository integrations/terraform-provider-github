package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemOwnerPatchRequestBody_security_and_analysis specify which security and analysis features to enable or disable for the repository.To use this parameter, you must have admin permissions for the repository or be an owner or security manager for the organization that owns the repository. For more information, see "[Managing security managers in your organization](https://docs.github.com/organizations/managing-peoples-access-to-your-organization-with-roles/managing-security-managers-in-your-organization)."For example, to enable GitHub Advanced Security, use this data in the body of the `PATCH` request:`{ "security_and_analysis": {"advanced_security": { "status": "enabled" } } }`.You can check which security and analysis features are currently enabled by using a `GET /repos/{owner}/{repo}` request.
type ItemItemOwnerPatchRequestBody_security_and_analysis struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Use the `status` property to enable or disable GitHub Advanced Security for this repository. For more information, see "[About GitHub Advanced Security](/github/getting-started-with-github/learning-about-github/about-github-advanced-security)."
    advanced_security ItemItemOwnerPatchRequestBody_security_and_analysis_advanced_securityable
    // Use the `status` property to enable or disable secret scanning for this repository. For more information, see "[About secret scanning](/code-security/secret-security/about-secret-scanning)."
    secret_scanning ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanningable
    // Use the `status` property to enable or disable secret scanning push protection for this repository. For more information, see "[Protecting pushes with secret scanning](/code-security/secret-scanning/protecting-pushes-with-secret-scanning)."
    secret_scanning_push_protection ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning_push_protectionable
}
// NewItemItemOwnerPatchRequestBody_security_and_analysis instantiates a new ItemItemOwnerPatchRequestBody_security_and_analysis and sets the default values.
func NewItemItemOwnerPatchRequestBody_security_and_analysis()(*ItemItemOwnerPatchRequestBody_security_and_analysis) {
    m := &ItemItemOwnerPatchRequestBody_security_and_analysis{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemOwnerPatchRequestBody_security_and_analysisFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemOwnerPatchRequestBody_security_and_analysisFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemOwnerPatchRequestBody_security_and_analysis(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdvancedSecurity gets the advanced_security property value. Use the `status` property to enable or disable GitHub Advanced Security for this repository. For more information, see "[About GitHub Advanced Security](/github/getting-started-with-github/learning-about-github/about-github-advanced-security)."
// returns a ItemItemOwnerPatchRequestBody_security_and_analysis_advanced_securityable when successful
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis) GetAdvancedSecurity()(ItemItemOwnerPatchRequestBody_security_and_analysis_advanced_securityable) {
    return m.advanced_security
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["advanced_security"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemOwnerPatchRequestBody_security_and_analysis_advanced_securityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedSecurity(val.(ItemItemOwnerPatchRequestBody_security_and_analysis_advanced_securityable))
        }
        return nil
    }
    res["secret_scanning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanningFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanning(val.(ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanningable))
        }
        return nil
    }
    res["secret_scanning_push_protection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning_push_protectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretScanningPushProtection(val.(ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning_push_protectionable))
        }
        return nil
    }
    return res
}
// GetSecretScanning gets the secret_scanning property value. Use the `status` property to enable or disable secret scanning for this repository. For more information, see "[About secret scanning](/code-security/secret-security/about-secret-scanning)."
// returns a ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanningable when successful
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis) GetSecretScanning()(ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanningable) {
    return m.secret_scanning
}
// GetSecretScanningPushProtection gets the secret_scanning_push_protection property value. Use the `status` property to enable or disable secret scanning push protection for this repository. For more information, see "[Protecting pushes with secret scanning](/code-security/secret-scanning/protecting-pushes-with-secret-scanning)."
// returns a ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning_push_protectionable when successful
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis) GetSecretScanningPushProtection()(ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning_push_protectionable) {
    return m.secret_scanning_push_protection
}
// Serialize serializes information the current object
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("advanced_security", m.GetAdvancedSecurity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("secret_scanning", m.GetSecretScanning())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("secret_scanning_push_protection", m.GetSecretScanningPushProtection())
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
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdvancedSecurity sets the advanced_security property value. Use the `status` property to enable or disable GitHub Advanced Security for this repository. For more information, see "[About GitHub Advanced Security](/github/getting-started-with-github/learning-about-github/about-github-advanced-security)."
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis) SetAdvancedSecurity(value ItemItemOwnerPatchRequestBody_security_and_analysis_advanced_securityable)() {
    m.advanced_security = value
}
// SetSecretScanning sets the secret_scanning property value. Use the `status` property to enable or disable secret scanning for this repository. For more information, see "[About secret scanning](/code-security/secret-security/about-secret-scanning)."
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis) SetSecretScanning(value ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanningable)() {
    m.secret_scanning = value
}
// SetSecretScanningPushProtection sets the secret_scanning_push_protection property value. Use the `status` property to enable or disable secret scanning push protection for this repository. For more information, see "[Protecting pushes with secret scanning](/code-security/secret-scanning/protecting-pushes-with-secret-scanning)."
func (m *ItemItemOwnerPatchRequestBody_security_and_analysis) SetSecretScanningPushProtection(value ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning_push_protectionable)() {
    m.secret_scanning_push_protection = value
}
type ItemItemOwnerPatchRequestBody_security_and_analysisable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedSecurity()(ItemItemOwnerPatchRequestBody_security_and_analysis_advanced_securityable)
    GetSecretScanning()(ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanningable)
    GetSecretScanningPushProtection()(ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning_push_protectionable)
    SetAdvancedSecurity(value ItemItemOwnerPatchRequestBody_security_and_analysis_advanced_securityable)()
    SetSecretScanning(value ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanningable)()
    SetSecretScanningPushProtection(value ItemItemOwnerPatchRequestBody_security_and_analysis_secret_scanning_push_protectionable)()
}
