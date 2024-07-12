package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemPagesDeploymentPostRequestBody the object used to create GitHub Pages deployment
type ItemItemPagesDeploymentPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The URL of an artifact that contains the .zip or .tar of static assets to deploy. The artifact belongs to the repository.
    artifact_url *string
    // The target environment for this GitHub Pages deployment.
    environment *string
    // The OIDC token issued by GitHub Actions certifying the origin of the deployment.
    oidc_token *string
    // A unique string that represents the version of the build for this deployment.
    pages_build_version *string
}
// NewItemItemPagesDeploymentPostRequestBody instantiates a new ItemItemPagesDeploymentPostRequestBody and sets the default values.
func NewItemItemPagesDeploymentPostRequestBody()(*ItemItemPagesDeploymentPostRequestBody) {
    m := &ItemItemPagesDeploymentPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    environmentValue := "github-pages"
    m.SetEnvironment(&environmentValue)
    pages_build_versionValue := "GITHUB_SHA"
    m.SetPagesBuildVersion(&pages_build_versionValue)
    return m
}
// CreateItemItemPagesDeploymentPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemPagesDeploymentPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemPagesDeploymentPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemPagesDeploymentPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArtifactUrl gets the artifact_url property value. The URL of an artifact that contains the .zip or .tar of static assets to deploy. The artifact belongs to the repository.
func (m *ItemItemPagesDeploymentPostRequestBody) GetArtifactUrl()(*string) {
    return m.artifact_url
}
// GetEnvironment gets the environment property value. The target environment for this GitHub Pages deployment.
func (m *ItemItemPagesDeploymentPostRequestBody) GetEnvironment()(*string) {
    return m.environment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemPagesDeploymentPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *ItemItemPagesDeploymentPostRequestBody) GetOidcToken()(*string) {
    return m.oidc_token
}
// GetPagesBuildVersion gets the pages_build_version property value. A unique string that represents the version of the build for this deployment.
func (m *ItemItemPagesDeploymentPostRequestBody) GetPagesBuildVersion()(*string) {
    return m.pages_build_version
}
// Serialize serializes information the current object
func (m *ItemItemPagesDeploymentPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemItemPagesDeploymentPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArtifactUrl sets the artifact_url property value. The URL of an artifact that contains the .zip or .tar of static assets to deploy. The artifact belongs to the repository.
func (m *ItemItemPagesDeploymentPostRequestBody) SetArtifactUrl(value *string)() {
    m.artifact_url = value
}
// SetEnvironment sets the environment property value. The target environment for this GitHub Pages deployment.
func (m *ItemItemPagesDeploymentPostRequestBody) SetEnvironment(value *string)() {
    m.environment = value
}
// SetOidcToken sets the oidc_token property value. The OIDC token issued by GitHub Actions certifying the origin of the deployment.
func (m *ItemItemPagesDeploymentPostRequestBody) SetOidcToken(value *string)() {
    m.oidc_token = value
}
// SetPagesBuildVersion sets the pages_build_version property value. A unique string that represents the version of the build for this deployment.
func (m *ItemItemPagesDeploymentPostRequestBody) SetPagesBuildVersion(value *string)() {
    m.pages_build_version = value
}
// ItemItemPagesDeploymentPostRequestBodyable 
type ItemItemPagesDeploymentPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArtifactUrl()(*string)
    GetEnvironment()(*string)
    GetOidcToken()(*string)
    GetPagesBuildVersion()(*string)
    SetArtifactUrl(value *string)()
    SetEnvironment(value *string)()
    SetOidcToken(value *string)()
    SetPagesBuildVersion(value *string)()
}
