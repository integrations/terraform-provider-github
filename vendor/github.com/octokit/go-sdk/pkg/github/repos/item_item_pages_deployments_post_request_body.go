package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemPagesDeploymentsPostRequestBody the object used to create GitHub Pages deployment
type ItemItemPagesDeploymentsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ID of an artifact that contains the .zip or .tar of static assets to deploy. The artifact belongs to the repository. Either `artifact_id` or `artifact_url` are required.
    artifact_id *float64
    // The URL of an artifact that contains the .zip or .tar of static assets to deploy. The artifact belongs to the repository. Either `artifact_id` or `artifact_url` are required.
    artifact_url *string
    // The target environment for this GitHub Pages deployment.
    environment *string
    // The OIDC token issued by GitHub Actions certifying the origin of the deployment.
    oidc_token *string
    // A unique string that represents the version of the build for this deployment.
    pages_build_version *string
}
// NewItemItemPagesDeploymentsPostRequestBody instantiates a new ItemItemPagesDeploymentsPostRequestBody and sets the default values.
func NewItemItemPagesDeploymentsPostRequestBody()(*ItemItemPagesDeploymentsPostRequestBody) {
    m := &ItemItemPagesDeploymentsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    environmentValue := "github-pages"
    m.SetEnvironment(&environmentValue)
    pages_build_versionValue := "GITHUB_SHA"
    m.SetPagesBuildVersion(&pages_build_versionValue)
    return m
}
// CreateItemItemPagesDeploymentsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemPagesDeploymentsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPagesDeploymentsPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemPagesDeploymentsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArtifactId gets the artifact_id property value. The ID of an artifact that contains the .zip or .tar of static assets to deploy. The artifact belongs to the repository. Either `artifact_id` or `artifact_url` are required.
// returns a *float64 when successful
func (m *ItemItemPagesDeploymentsPostRequestBody) GetArtifactId()(*float64) {
    return m.artifact_id
}
// GetArtifactUrl gets the artifact_url property value. The URL of an artifact that contains the .zip or .tar of static assets to deploy. The artifact belongs to the repository. Either `artifact_id` or `artifact_url` are required.
// returns a *string when successful
func (m *ItemItemPagesDeploymentsPostRequestBody) GetArtifactUrl()(*string) {
    return m.artifact_url
}
// GetEnvironment gets the environment property value. The target environment for this GitHub Pages deployment.
// returns a *string when successful
func (m *ItemItemPagesDeploymentsPostRequestBody) GetEnvironment()(*string) {
    return m.environment
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemPagesDeploymentsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["artifact_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArtifactId(val)
        }
        return nil
    }
    res["artifact_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArtifactUrl(val)
        }
        return nil
    }
    res["environment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnvironment(val)
        }
        return nil
    }
    res["oidc_token"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOidcToken(val)
        }
        return nil
    }
    res["pages_build_version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPagesBuildVersion(val)
        }
        return nil
    }
    return res
}
// GetOidcToken gets the oidc_token property value. The OIDC token issued by GitHub Actions certifying the origin of the deployment.
// returns a *string when successful
func (m *ItemItemPagesDeploymentsPostRequestBody) GetOidcToken()(*string) {
    return m.oidc_token
}
// GetPagesBuildVersion gets the pages_build_version property value. A unique string that represents the version of the build for this deployment.
// returns a *string when successful
func (m *ItemItemPagesDeploymentsPostRequestBody) GetPagesBuildVersion()(*string) {
    return m.pages_build_version
}
// Serialize serializes information the current object
func (m *ItemItemPagesDeploymentsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("artifact_id", m.GetArtifactId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("artifact_url", m.GetArtifactUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("environment", m.GetEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("oidc_token", m.GetOidcToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pages_build_version", m.GetPagesBuildVersion())
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
func (m *ItemItemPagesDeploymentsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArtifactId sets the artifact_id property value. The ID of an artifact that contains the .zip or .tar of static assets to deploy. The artifact belongs to the repository. Either `artifact_id` or `artifact_url` are required.
func (m *ItemItemPagesDeploymentsPostRequestBody) SetArtifactId(value *float64)() {
    m.artifact_id = value
}
// SetArtifactUrl sets the artifact_url property value. The URL of an artifact that contains the .zip or .tar of static assets to deploy. The artifact belongs to the repository. Either `artifact_id` or `artifact_url` are required.
func (m *ItemItemPagesDeploymentsPostRequestBody) SetArtifactUrl(value *string)() {
    m.artifact_url = value
}
// SetEnvironment sets the environment property value. The target environment for this GitHub Pages deployment.
func (m *ItemItemPagesDeploymentsPostRequestBody) SetEnvironment(value *string)() {
    m.environment = value
}
// SetOidcToken sets the oidc_token property value. The OIDC token issued by GitHub Actions certifying the origin of the deployment.
func (m *ItemItemPagesDeploymentsPostRequestBody) SetOidcToken(value *string)() {
    m.oidc_token = value
}
// SetPagesBuildVersion sets the pages_build_version property value. A unique string that represents the version of the build for this deployment.
func (m *ItemItemPagesDeploymentsPostRequestBody) SetPagesBuildVersion(value *string)() {
    m.pages_build_version = value
}
type ItemItemPagesDeploymentsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArtifactId()(*float64)
    GetArtifactUrl()(*string)
    GetEnvironment()(*string)
    GetOidcToken()(*string)
    GetPagesBuildVersion()(*string)
    SetArtifactId(value *float64)()
    SetArtifactUrl(value *string)()
    SetEnvironment(value *string)()
    SetOidcToken(value *string)()
    SetPagesBuildVersion(value *string)()
}
