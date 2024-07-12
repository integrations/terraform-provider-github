package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrgRepoCustomPropertyValues list of custom property values for a repository
type OrgRepoCustomPropertyValues struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of custom property names and associated values
    properties []CustomPropertyValueable
    // The repository_full_name property
    repository_full_name *string
    // The repository_id property
    repository_id *int32
    // The repository_name property
    repository_name *string
}
// NewOrgRepoCustomPropertyValues instantiates a new OrgRepoCustomPropertyValues and sets the default values.
func NewOrgRepoCustomPropertyValues()(*OrgRepoCustomPropertyValues) {
    m := &OrgRepoCustomPropertyValues{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrgRepoCustomPropertyValuesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrgRepoCustomPropertyValuesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrgRepoCustomPropertyValues(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrgRepoCustomPropertyValues) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrgRepoCustomPropertyValues) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomPropertyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomPropertyValueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomPropertyValueable)
                }
            }
            m.SetProperties(res)
        }
        return nil
    }
    res["repository_full_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryFullName(val)
        }
        return nil
    }
    res["repository_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryId(val)
        }
        return nil
    }
    res["repository_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryName(val)
        }
        return nil
    }
    return res
}
// GetProperties gets the properties property value. List of custom property names and associated values
// returns a []CustomPropertyValueable when successful
func (m *OrgRepoCustomPropertyValues) GetProperties()([]CustomPropertyValueable) {
    return m.properties
}
// GetRepositoryFullName gets the repository_full_name property value. The repository_full_name property
// returns a *string when successful
func (m *OrgRepoCustomPropertyValues) GetRepositoryFullName()(*string) {
    return m.repository_full_name
}
// GetRepositoryId gets the repository_id property value. The repository_id property
// returns a *int32 when successful
func (m *OrgRepoCustomPropertyValues) GetRepositoryId()(*int32) {
    return m.repository_id
}
// GetRepositoryName gets the repository_name property value. The repository_name property
// returns a *string when successful
func (m *OrgRepoCustomPropertyValues) GetRepositoryName()(*string) {
    return m.repository_name
}
// Serialize serializes information the current object
func (m *OrgRepoCustomPropertyValues) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("properties", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repository_full_name", m.GetRepositoryFullName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("repository_id", m.GetRepositoryId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("repository_name", m.GetRepositoryName())
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
func (m *OrgRepoCustomPropertyValues) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetProperties sets the properties property value. List of custom property names and associated values
func (m *OrgRepoCustomPropertyValues) SetProperties(value []CustomPropertyValueable)() {
    m.properties = value
}
// SetRepositoryFullName sets the repository_full_name property value. The repository_full_name property
func (m *OrgRepoCustomPropertyValues) SetRepositoryFullName(value *string)() {
    m.repository_full_name = value
}
// SetRepositoryId sets the repository_id property value. The repository_id property
func (m *OrgRepoCustomPropertyValues) SetRepositoryId(value *int32)() {
    m.repository_id = value
}
// SetRepositoryName sets the repository_name property value. The repository_name property
func (m *OrgRepoCustomPropertyValues) SetRepositoryName(value *string)() {
    m.repository_name = value
}
type OrgRepoCustomPropertyValuesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetProperties()([]CustomPropertyValueable)
    GetRepositoryFullName()(*string)
    GetRepositoryId()(*int32)
    GetRepositoryName()(*string)
    SetProperties(value []CustomPropertyValueable)()
    SetRepositoryFullName(value *string)()
    SetRepositoryId(value *int32)()
    SetRepositoryName(value *string)()
}
