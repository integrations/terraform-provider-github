package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RepositoryAdvisory_cvss struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The CVSS score.
    score *float64
    // The CVSS vector.
    vector_string *string
}
// NewRepositoryAdvisory_cvss instantiates a new RepositoryAdvisory_cvss and sets the default values.
func NewRepositoryAdvisory_cvss()(*RepositoryAdvisory_cvss) {
    m := &RepositoryAdvisory_cvss{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRepositoryAdvisory_cvssFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRepositoryAdvisory_cvssFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoryAdvisory_cvss(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *RepositoryAdvisory_cvss) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RepositoryAdvisory_cvss) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["score"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScore(val)
        }
        return nil
    }
    res["vector_string"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVectorString(val)
        }
        return nil
    }
    return res
}
// GetScore gets the score property value. The CVSS score.
// returns a *float64 when successful
func (m *RepositoryAdvisory_cvss) GetScore()(*float64) {
    return m.score
}
// GetVectorString gets the vector_string property value. The CVSS vector.
// returns a *string when successful
func (m *RepositoryAdvisory_cvss) GetVectorString()(*string) {
    return m.vector_string
}
// Serialize serializes information the current object
func (m *RepositoryAdvisory_cvss) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("vector_string", m.GetVectorString())
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
func (m *RepositoryAdvisory_cvss) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetScore sets the score property value. The CVSS score.
func (m *RepositoryAdvisory_cvss) SetScore(value *float64)() {
    m.score = value
}
// SetVectorString sets the vector_string property value. The CVSS vector.
func (m *RepositoryAdvisory_cvss) SetVectorString(value *string)() {
    m.vector_string = value
}
type RepositoryAdvisory_cvssable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetScore()(*float64)
    GetVectorString()(*string)
    SetScore(value *float64)()
    SetVectorString(value *string)()
}
