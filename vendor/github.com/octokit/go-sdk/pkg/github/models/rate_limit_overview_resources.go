package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RateLimitOverview_resources struct {
    // The actions_runner_registration property
    actions_runner_registration RateLimitable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The code_scanning_upload property
    code_scanning_upload RateLimitable
    // The code_search property
    code_search RateLimitable
    // The core property
    core RateLimitable
    // The dependency_snapshots property
    dependency_snapshots RateLimitable
    // The graphql property
    graphql RateLimitable
    // The integration_manifest property
    integration_manifest RateLimitable
    // The scim property
    scim RateLimitable
    // The search property
    search RateLimitable
    // The source_import property
    source_import RateLimitable
}
// NewRateLimitOverview_resources instantiates a new RateLimitOverview_resources and sets the default values.
func NewRateLimitOverview_resources()(*RateLimitOverview_resources) {
    m := &RateLimitOverview_resources{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRateLimitOverview_resourcesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRateLimitOverview_resourcesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRateLimitOverview_resources(), nil
}
// GetActionsRunnerRegistration gets the actions_runner_registration property value. The actions_runner_registration property
// returns a RateLimitable when successful
func (m *RateLimitOverview_resources) GetActionsRunnerRegistration()(RateLimitable) {
    return m.actions_runner_registration
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RateLimitOverview_resources) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCodeScanningUpload gets the code_scanning_upload property value. The code_scanning_upload property
// returns a RateLimitable when successful
func (m *RateLimitOverview_resources) GetCodeScanningUpload()(RateLimitable) {
    return m.code_scanning_upload
}
// GetCodeSearch gets the code_search property value. The code_search property
// returns a RateLimitable when successful
func (m *RateLimitOverview_resources) GetCodeSearch()(RateLimitable) {
    return m.code_search
}
// GetCore gets the core property value. The core property
// returns a RateLimitable when successful
func (m *RateLimitOverview_resources) GetCore()(RateLimitable) {
    return m.core
}
// GetDependencySnapshots gets the dependency_snapshots property value. The dependency_snapshots property
// returns a RateLimitable when successful
func (m *RateLimitOverview_resources) GetDependencySnapshots()(RateLimitable) {
    return m.dependency_snapshots
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RateLimitOverview_resources) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actions_runner_registration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRateLimitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionsRunnerRegistration(val.(RateLimitable))
        }
        return nil
    }
    res["code_scanning_upload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRateLimitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeScanningUpload(val.(RateLimitable))
        }
        return nil
    }
    res["code_search"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRateLimitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeSearch(val.(RateLimitable))
        }
        return nil
    }
    res["core"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRateLimitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCore(val.(RateLimitable))
        }
        return nil
    }
    res["dependency_snapshots"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRateLimitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependencySnapshots(val.(RateLimitable))
        }
        return nil
    }
    res["graphql"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRateLimitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGraphql(val.(RateLimitable))
        }
        return nil
    }
    res["integration_manifest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRateLimitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegrationManifest(val.(RateLimitable))
        }
        return nil
    }
    res["scim"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRateLimitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScim(val.(RateLimitable))
        }
        return nil
    }
    res["search"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRateLimitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearch(val.(RateLimitable))
        }
        return nil
    }
    res["source_import"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRateLimitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceImport(val.(RateLimitable))
        }
        return nil
    }
    return res
}
// GetGraphql gets the graphql property value. The graphql property
// returns a RateLimitable when successful
func (m *RateLimitOverview_resources) GetGraphql()(RateLimitable) {
    return m.graphql
}
// GetIntegrationManifest gets the integration_manifest property value. The integration_manifest property
// returns a RateLimitable when successful
func (m *RateLimitOverview_resources) GetIntegrationManifest()(RateLimitable) {
    return m.integration_manifest
}
// GetScim gets the scim property value. The scim property
// returns a RateLimitable when successful
func (m *RateLimitOverview_resources) GetScim()(RateLimitable) {
    return m.scim
}
// GetSearch gets the search property value. The search property
// returns a RateLimitable when successful
func (m *RateLimitOverview_resources) GetSearch()(RateLimitable) {
    return m.search
}
// GetSourceImport gets the source_import property value. The source_import property
// returns a RateLimitable when successful
func (m *RateLimitOverview_resources) GetSourceImport()(RateLimitable) {
    return m.source_import
}
// Serialize serializes information the current object
func (m *RateLimitOverview_resources) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("actions_runner_registration", m.GetActionsRunnerRegistration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("code_scanning_upload", m.GetCodeScanningUpload())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("code_search", m.GetCodeSearch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("core", m.GetCore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("dependency_snapshots", m.GetDependencySnapshots())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("graphql", m.GetGraphql())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("integration_manifest", m.GetIntegrationManifest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("scim", m.GetScim())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("search", m.GetSearch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("source_import", m.GetSourceImport())
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
// SetActionsRunnerRegistration sets the actions_runner_registration property value. The actions_runner_registration property
func (m *RateLimitOverview_resources) SetActionsRunnerRegistration(value RateLimitable)() {
    m.actions_runner_registration = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RateLimitOverview_resources) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCodeScanningUpload sets the code_scanning_upload property value. The code_scanning_upload property
func (m *RateLimitOverview_resources) SetCodeScanningUpload(value RateLimitable)() {
    m.code_scanning_upload = value
}
// SetCodeSearch sets the code_search property value. The code_search property
func (m *RateLimitOverview_resources) SetCodeSearch(value RateLimitable)() {
    m.code_search = value
}
// SetCore sets the core property value. The core property
func (m *RateLimitOverview_resources) SetCore(value RateLimitable)() {
    m.core = value
}
// SetDependencySnapshots sets the dependency_snapshots property value. The dependency_snapshots property
func (m *RateLimitOverview_resources) SetDependencySnapshots(value RateLimitable)() {
    m.dependency_snapshots = value
}
// SetGraphql sets the graphql property value. The graphql property
func (m *RateLimitOverview_resources) SetGraphql(value RateLimitable)() {
    m.graphql = value
}
// SetIntegrationManifest sets the integration_manifest property value. The integration_manifest property
func (m *RateLimitOverview_resources) SetIntegrationManifest(value RateLimitable)() {
    m.integration_manifest = value
}
// SetScim sets the scim property value. The scim property
func (m *RateLimitOverview_resources) SetScim(value RateLimitable)() {
    m.scim = value
}
// SetSearch sets the search property value. The search property
func (m *RateLimitOverview_resources) SetSearch(value RateLimitable)() {
    m.search = value
}
// SetSourceImport sets the source_import property value. The source_import property
func (m *RateLimitOverview_resources) SetSourceImport(value RateLimitable)() {
    m.source_import = value
}
type RateLimitOverview_resourcesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionsRunnerRegistration()(RateLimitable)
    GetCodeScanningUpload()(RateLimitable)
    GetCodeSearch()(RateLimitable)
    GetCore()(RateLimitable)
    GetDependencySnapshots()(RateLimitable)
    GetGraphql()(RateLimitable)
    GetIntegrationManifest()(RateLimitable)
    GetScim()(RateLimitable)
    GetSearch()(RateLimitable)
    GetSourceImport()(RateLimitable)
    SetActionsRunnerRegistration(value RateLimitable)()
    SetCodeScanningUpload(value RateLimitable)()
    SetCodeSearch(value RateLimitable)()
    SetCore(value RateLimitable)()
    SetDependencySnapshots(value RateLimitable)()
    SetGraphql(value RateLimitable)()
    SetIntegrationManifest(value RateLimitable)()
    SetScim(value RateLimitable)()
    SetSearch(value RateLimitable)()
    SetSourceImport(value RateLimitable)()
}
