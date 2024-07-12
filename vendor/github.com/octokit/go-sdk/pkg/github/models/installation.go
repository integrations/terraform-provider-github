package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Installation installation
type Installation struct {
    // The access_tokens_url property
    access_tokens_url *string
    // The account property
    account Installation_Installation_accountable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The app_id property
    app_id *int32
    // The app_slug property
    app_slug *string
    // The contact_email property
    contact_email *string
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The events property
    events []string
    // The has_multiple_single_files property
    has_multiple_single_files *bool
    // The html_url property
    html_url *string
    // The ID of the installation.
    id *int32
    // The permissions granted to the user access token.
    permissions AppPermissionsable
    // The repositories_url property
    repositories_url *string
    // Describe whether all repositories have been selected or there's a selection involved
    repository_selection *Installation_repository_selection
    // The single_file_name property
    single_file_name *string
    // The single_file_paths property
    single_file_paths []string
    // The suspended_at property
    suspended_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A GitHub user.
    suspended_by NullableSimpleUserable
    // The ID of the user or organization this token is being scoped to.
    target_id *int32
    // The target_type property
    target_type *string
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// Installation_Installation_account composed type wrapper for classes Enterpriseable, SimpleUserable
type Installation_Installation_account struct {
    // Composed type representation for type Enterpriseable
    enterprise Enterpriseable
    // Composed type representation for type SimpleUserable
    simpleUser SimpleUserable
}
// NewInstallation_Installation_account instantiates a new Installation_Installation_account and sets the default values.
func NewInstallation_Installation_account()(*Installation_Installation_account) {
    m := &Installation_Installation_account{
    }
    return m
}
// CreateInstallation_Installation_accountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInstallation_Installation_accountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewInstallation_Installation_account()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateEnterpriseFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(Enterpriseable); ok {
                result.SetEnterprise(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateSimpleUserFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(SimpleUserable); ok {
                result.SetSimpleUser(cast)
            }
        }
    }
    return result, nil
}
// GetEnterprise gets the enterprise property value. Composed type representation for type Enterpriseable
// returns a Enterpriseable when successful
func (m *Installation_Installation_account) GetEnterprise()(Enterpriseable) {
    return m.enterprise
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Installation_Installation_account) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *Installation_Installation_account) GetIsComposedType()(bool) {
    return true
}
// GetSimpleUser gets the simpleUser property value. Composed type representation for type SimpleUserable
// returns a SimpleUserable when successful
func (m *Installation_Installation_account) GetSimpleUser()(SimpleUserable) {
    return m.simpleUser
}
// Serialize serializes information the current object
func (m *Installation_Installation_account) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEnterprise() != nil {
        err := writer.WriteObjectValue("", m.GetEnterprise())
        if err != nil {
            return err
        }
    } else if m.GetSimpleUser() != nil {
        err := writer.WriteObjectValue("", m.GetSimpleUser())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnterprise sets the enterprise property value. Composed type representation for type Enterpriseable
func (m *Installation_Installation_account) SetEnterprise(value Enterpriseable)() {
    m.enterprise = value
}
// SetSimpleUser sets the simpleUser property value. Composed type representation for type SimpleUserable
func (m *Installation_Installation_account) SetSimpleUser(value SimpleUserable)() {
    m.simpleUser = value
}
type Installation_Installation_accountable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnterprise()(Enterpriseable)
    GetSimpleUser()(SimpleUserable)
    SetEnterprise(value Enterpriseable)()
    SetSimpleUser(value SimpleUserable)()
}
// NewInstallation instantiates a new Installation and sets the default values.
func NewInstallation()(*Installation) {
    m := &Installation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInstallationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInstallationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInstallation(), nil
}
// GetAccessTokensUrl gets the access_tokens_url property value. The access_tokens_url property
// returns a *string when successful
func (m *Installation) GetAccessTokensUrl()(*string) {
    return m.access_tokens_url
}
// GetAccount gets the account property value. The account property
// returns a Installation_Installation_accountable when successful
func (m *Installation) GetAccount()(Installation_Installation_accountable) {
    return m.account
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Installation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppId gets the app_id property value. The app_id property
// returns a *int32 when successful
func (m *Installation) GetAppId()(*int32) {
    return m.app_id
}
// GetAppSlug gets the app_slug property value. The app_slug property
// returns a *string when successful
func (m *Installation) GetAppSlug()(*string) {
    return m.app_slug
}
// GetContactEmail gets the contact_email property value. The contact_email property
// returns a *string when successful
func (m *Installation) GetContactEmail()(*string) {
    return m.contact_email
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *Installation) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetEvents gets the events property value. The events property
// returns a []string when successful
func (m *Installation) GetEvents()([]string) {
    return m.events
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Installation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["access_tokens_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessTokensUrl(val)
        }
        return nil
    }
    res["account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInstallation_Installation_accountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccount(val.(Installation_Installation_accountable))
        }
        return nil
    }
    res["app_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["app_slug"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppSlug(val)
        }
        return nil
    }
    res["contact_email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactEmail(val)
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
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetEvents(res)
        }
        return nil
    }
    res["has_multiple_single_files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasMultipleSingleFiles(val)
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
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppPermissionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissions(val.(AppPermissionsable))
        }
        return nil
    }
    res["repositories_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoriesUrl(val)
        }
        return nil
    }
    res["repository_selection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInstallation_repository_selection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositorySelection(val.(*Installation_repository_selection))
        }
        return nil
    }
    res["single_file_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleFileName(val)
        }
        return nil
    }
    res["single_file_paths"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSingleFilePaths(res)
        }
        return nil
    }
    res["suspended_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuspendedAt(val)
        }
        return nil
    }
    res["suspended_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuspendedBy(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["target_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetId(val)
        }
        return nil
    }
    res["target_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetType(val)
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
    return res
}
// GetHasMultipleSingleFiles gets the has_multiple_single_files property value. The has_multiple_single_files property
// returns a *bool when successful
func (m *Installation) GetHasMultipleSingleFiles()(*bool) {
    return m.has_multiple_single_files
}
// GetHtmlUrl gets the html_url property value. The html_url property
// returns a *string when successful
func (m *Installation) GetHtmlUrl()(*string) {
    return m.html_url
}
// GetId gets the id property value. The ID of the installation.
// returns a *int32 when successful
func (m *Installation) GetId()(*int32) {
    return m.id
}
// GetPermissions gets the permissions property value. The permissions granted to the user access token.
// returns a AppPermissionsable when successful
func (m *Installation) GetPermissions()(AppPermissionsable) {
    return m.permissions
}
// GetRepositoriesUrl gets the repositories_url property value. The repositories_url property
// returns a *string when successful
func (m *Installation) GetRepositoriesUrl()(*string) {
    return m.repositories_url
}
// GetRepositorySelection gets the repository_selection property value. Describe whether all repositories have been selected or there's a selection involved
// returns a *Installation_repository_selection when successful
func (m *Installation) GetRepositorySelection()(*Installation_repository_selection) {
    return m.repository_selection
}
// GetSingleFileName gets the single_file_name property value. The single_file_name property
// returns a *string when successful
func (m *Installation) GetSingleFileName()(*string) {
    return m.single_file_name
}
// GetSingleFilePaths gets the single_file_paths property value. The single_file_paths property
// returns a []string when successful
func (m *Installation) GetSingleFilePaths()([]string) {
    return m.single_file_paths
}
// GetSuspendedAt gets the suspended_at property value. The suspended_at property
// returns a *Time when successful
func (m *Installation) GetSuspendedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.suspended_at
}
// GetSuspendedBy gets the suspended_by property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *Installation) GetSuspendedBy()(NullableSimpleUserable) {
    return m.suspended_by
}
// GetTargetId gets the target_id property value. The ID of the user or organization this token is being scoped to.
// returns a *int32 when successful
func (m *Installation) GetTargetId()(*int32) {
    return m.target_id
}
// GetTargetType gets the target_type property value. The target_type property
// returns a *string when successful
func (m *Installation) GetTargetType()(*string) {
    return m.target_type
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *Installation) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *Installation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("access_tokens_url", m.GetAccessTokensUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("account", m.GetAccount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("app_id", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("app_slug", m.GetAppSlug())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contact_email", m.GetContactEmail())
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
    if m.GetEvents() != nil {
        err := writer.WriteCollectionOfStringValues("events", m.GetEvents())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("has_multiple_single_files", m.GetHasMultipleSingleFiles())
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
        err := writer.WriteObjectValue("permissions", m.GetPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repositories_url", m.GetRepositoriesUrl())
        if err != nil {
            return err
        }
    }
    if m.GetRepositorySelection() != nil {
        cast := (*m.GetRepositorySelection()).String()
        err := writer.WriteStringValue("repository_selection", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("single_file_name", m.GetSingleFileName())
        if err != nil {
            return err
        }
    }
    if m.GetSingleFilePaths() != nil {
        err := writer.WriteCollectionOfStringValues("single_file_paths", m.GetSingleFilePaths())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("suspended_at", m.GetSuspendedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("suspended_by", m.GetSuspendedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("target_id", m.GetTargetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target_type", m.GetTargetType())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessTokensUrl sets the access_tokens_url property value. The access_tokens_url property
func (m *Installation) SetAccessTokensUrl(value *string)() {
    m.access_tokens_url = value
}
// SetAccount sets the account property value. The account property
func (m *Installation) SetAccount(value Installation_Installation_accountable)() {
    m.account = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Installation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppId sets the app_id property value. The app_id property
func (m *Installation) SetAppId(value *int32)() {
    m.app_id = value
}
// SetAppSlug sets the app_slug property value. The app_slug property
func (m *Installation) SetAppSlug(value *string)() {
    m.app_slug = value
}
// SetContactEmail sets the contact_email property value. The contact_email property
func (m *Installation) SetContactEmail(value *string)() {
    m.contact_email = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *Installation) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetEvents sets the events property value. The events property
func (m *Installation) SetEvents(value []string)() {
    m.events = value
}
// SetHasMultipleSingleFiles sets the has_multiple_single_files property value. The has_multiple_single_files property
func (m *Installation) SetHasMultipleSingleFiles(value *bool)() {
    m.has_multiple_single_files = value
}
// SetHtmlUrl sets the html_url property value. The html_url property
func (m *Installation) SetHtmlUrl(value *string)() {
    m.html_url = value
}
// SetId sets the id property value. The ID of the installation.
func (m *Installation) SetId(value *int32)() {
    m.id = value
}
// SetPermissions sets the permissions property value. The permissions granted to the user access token.
func (m *Installation) SetPermissions(value AppPermissionsable)() {
    m.permissions = value
}
// SetRepositoriesUrl sets the repositories_url property value. The repositories_url property
func (m *Installation) SetRepositoriesUrl(value *string)() {
    m.repositories_url = value
}
// SetRepositorySelection sets the repository_selection property value. Describe whether all repositories have been selected or there's a selection involved
func (m *Installation) SetRepositorySelection(value *Installation_repository_selection)() {
    m.repository_selection = value
}
// SetSingleFileName sets the single_file_name property value. The single_file_name property
func (m *Installation) SetSingleFileName(value *string)() {
    m.single_file_name = value
}
// SetSingleFilePaths sets the single_file_paths property value. The single_file_paths property
func (m *Installation) SetSingleFilePaths(value []string)() {
    m.single_file_paths = value
}
// SetSuspendedAt sets the suspended_at property value. The suspended_at property
func (m *Installation) SetSuspendedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.suspended_at = value
}
// SetSuspendedBy sets the suspended_by property value. A GitHub user.
func (m *Installation) SetSuspendedBy(value NullableSimpleUserable)() {
    m.suspended_by = value
}
// SetTargetId sets the target_id property value. The ID of the user or organization this token is being scoped to.
func (m *Installation) SetTargetId(value *int32)() {
    m.target_id = value
}
// SetTargetType sets the target_type property value. The target_type property
func (m *Installation) SetTargetType(value *string)() {
    m.target_type = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *Installation) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
type Installationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessTokensUrl()(*string)
    GetAccount()(Installation_Installation_accountable)
    GetAppId()(*int32)
    GetAppSlug()(*string)
    GetContactEmail()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEvents()([]string)
    GetHasMultipleSingleFiles()(*bool)
    GetHtmlUrl()(*string)
    GetId()(*int32)
    GetPermissions()(AppPermissionsable)
    GetRepositoriesUrl()(*string)
    GetRepositorySelection()(*Installation_repository_selection)
    GetSingleFileName()(*string)
    GetSingleFilePaths()([]string)
    GetSuspendedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSuspendedBy()(NullableSimpleUserable)
    GetTargetId()(*int32)
    GetTargetType()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccessTokensUrl(value *string)()
    SetAccount(value Installation_Installation_accountable)()
    SetAppId(value *int32)()
    SetAppSlug(value *string)()
    SetContactEmail(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEvents(value []string)()
    SetHasMultipleSingleFiles(value *bool)()
    SetHtmlUrl(value *string)()
    SetId(value *int32)()
    SetPermissions(value AppPermissionsable)()
    SetRepositoriesUrl(value *string)()
    SetRepositorySelection(value *Installation_repository_selection)()
    SetSingleFileName(value *string)()
    SetSingleFilePaths(value []string)()
    SetSuspendedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSuspendedBy(value NullableSimpleUserable)()
    SetTargetId(value *int32)()
    SetTargetType(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
