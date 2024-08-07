package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CheckSuitePreference check suite configuration preferences for a repository.
type CheckSuitePreference struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The preferences property
    preferences CheckSuitePreference_preferencesable
    // Minimal Repository
    repository MinimalRepositoryable
}
// NewCheckSuitePreference instantiates a new CheckSuitePreference and sets the default values.
func NewCheckSuitePreference()(*CheckSuitePreference) {
    m := &CheckSuitePreference{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCheckSuitePreferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCheckSuitePreferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCheckSuitePreference(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CheckSuitePreference) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CheckSuitePreference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["preferences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCheckSuitePreference_preferencesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferences(val.(CheckSuitePreference_preferencesable))
        }
        return nil
    }
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMinimalRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(MinimalRepositoryable))
        }
        return nil
    }
    return res
}
// GetPreferences gets the preferences property value. The preferences property
// returns a CheckSuitePreference_preferencesable when successful
func (m *CheckSuitePreference) GetPreferences()(CheckSuitePreference_preferencesable) {
    return m.preferences
}
// GetRepository gets the repository property value. Minimal Repository
// returns a MinimalRepositoryable when successful
func (m *CheckSuitePreference) GetRepository()(MinimalRepositoryable) {
    return m.repository
}
// Serialize serializes information the current object
func (m *CheckSuitePreference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("preferences", m.GetPreferences())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repository", m.GetRepository())
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
func (m *CheckSuitePreference) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPreferences sets the preferences property value. The preferences property
func (m *CheckSuitePreference) SetPreferences(value CheckSuitePreference_preferencesable)() {
    m.preferences = value
}
// SetRepository sets the repository property value. Minimal Repository
func (m *CheckSuitePreference) SetRepository(value MinimalRepositoryable)() {
    m.repository = value
}
type CheckSuitePreferenceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPreferences()(CheckSuitePreference_preferencesable)
    GetRepository()(MinimalRepositoryable)
    SetPreferences(value CheckSuitePreference_preferencesable)()
    SetRepository(value MinimalRepositoryable)()
}
