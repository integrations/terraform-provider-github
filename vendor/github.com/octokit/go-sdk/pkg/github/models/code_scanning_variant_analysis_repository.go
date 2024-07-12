package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningVariantAnalysisRepository repository Identifier
type CodeScanningVariantAnalysisRepository struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The full, globally unique, name of the repository.
    full_name *string
    // A unique identifier of the repository.
    id *int32
    // The name of the repository.
    name *string
    // Whether the repository is private.
    private *bool
    // The stargazers_count property
    stargazers_count *int32
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewCodeScanningVariantAnalysisRepository instantiates a new CodeScanningVariantAnalysisRepository and sets the default values.
func NewCodeScanningVariantAnalysisRepository()(*CodeScanningVariantAnalysisRepository) {
    m := &CodeScanningVariantAnalysisRepository{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningVariantAnalysisRepositoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningVariantAnalysisRepositoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningVariantAnalysisRepository(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningVariantAnalysisRepository) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningVariantAnalysisRepository) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["full_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullName(val)
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
    res["private"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivate(val)
        }
        return nil
    }
    res["stargazers_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStargazersCount(val)
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
// GetFullName gets the full_name property value. The full, globally unique, name of the repository.
// returns a *string when successful
func (m *CodeScanningVariantAnalysisRepository) GetFullName()(*string) {
    return m.full_name
}
// GetId gets the id property value. A unique identifier of the repository.
// returns a *int32 when successful
func (m *CodeScanningVariantAnalysisRepository) GetId()(*int32) {
    return m.id
}
// GetName gets the name property value. The name of the repository.
// returns a *string when successful
func (m *CodeScanningVariantAnalysisRepository) GetName()(*string) {
    return m.name
}
// GetPrivate gets the private property value. Whether the repository is private.
// returns a *bool when successful
func (m *CodeScanningVariantAnalysisRepository) GetPrivate()(*bool) {
    return m.private
}
// GetStargazersCount gets the stargazers_count property value. The stargazers_count property
// returns a *int32 when successful
func (m *CodeScanningVariantAnalysisRepository) GetStargazersCount()(*int32) {
    return m.stargazers_count
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *CodeScanningVariantAnalysisRepository) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *CodeScanningVariantAnalysisRepository) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("full_name", m.GetFullName())
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
    {
        err := writer.WriteBoolValue("private", m.GetPrivate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("stargazers_count", m.GetStargazersCount())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodeScanningVariantAnalysisRepository) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFullName sets the full_name property value. The full, globally unique, name of the repository.
func (m *CodeScanningVariantAnalysisRepository) SetFullName(value *string)() {
    m.full_name = value
}
// SetId sets the id property value. A unique identifier of the repository.
func (m *CodeScanningVariantAnalysisRepository) SetId(value *int32)() {
    m.id = value
}
// SetName sets the name property value. The name of the repository.
func (m *CodeScanningVariantAnalysisRepository) SetName(value *string)() {
    m.name = value
}
// SetPrivate sets the private property value. Whether the repository is private.
func (m *CodeScanningVariantAnalysisRepository) SetPrivate(value *bool)() {
    m.private = value
}
// SetStargazersCount sets the stargazers_count property value. The stargazers_count property
func (m *CodeScanningVariantAnalysisRepository) SetStargazersCount(value *int32)() {
    m.stargazers_count = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *CodeScanningVariantAnalysisRepository) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
type CodeScanningVariantAnalysisRepositoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFullName()(*string)
    GetId()(*int32)
    GetName()(*string)
    GetPrivate()(*bool)
    GetStargazersCount()(*int32)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetFullName(value *string)()
    SetId(value *int32)()
    SetName(value *string)()
    SetPrivate(value *bool)()
    SetStargazersCount(value *int32)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
