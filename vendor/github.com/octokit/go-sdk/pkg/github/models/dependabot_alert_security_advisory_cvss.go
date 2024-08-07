package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DependabotAlertSecurityAdvisory_cvss details for the advisory pertaining to the Common Vulnerability Scoring System.
type DependabotAlertSecurityAdvisory_cvss struct {
    // The overall CVSS score of the advisory.
    score *float64
    // The full CVSS vector string for the advisory.
    vector_string *string
}
// NewDependabotAlertSecurityAdvisory_cvss instantiates a new DependabotAlertSecurityAdvisory_cvss and sets the default values.
func NewDependabotAlertSecurityAdvisory_cvss()(*DependabotAlertSecurityAdvisory_cvss) {
    m := &DependabotAlertSecurityAdvisory_cvss{
    }
    return m
}
// CreateDependabotAlertSecurityAdvisory_cvssFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDependabotAlertSecurityAdvisory_cvssFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDependabotAlertSecurityAdvisory_cvss(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DependabotAlertSecurityAdvisory_cvss) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetScore gets the score property value. The overall CVSS score of the advisory.
// returns a *float64 when successful
func (m *DependabotAlertSecurityAdvisory_cvss) GetScore()(*float64) {
    return m.score
}
// GetVectorString gets the vector_string property value. The full CVSS vector string for the advisory.
// returns a *string when successful
func (m *DependabotAlertSecurityAdvisory_cvss) GetVectorString()(*string) {
    return m.vector_string
}
// Serialize serializes information the current object
func (m *DependabotAlertSecurityAdvisory_cvss) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    return nil
}
// SetScore sets the score property value. The overall CVSS score of the advisory.
func (m *DependabotAlertSecurityAdvisory_cvss) SetScore(value *float64)() {
    m.score = value
}
// SetVectorString sets the vector_string property value. The full CVSS vector string for the advisory.
func (m *DependabotAlertSecurityAdvisory_cvss) SetVectorString(value *string)() {
    m.vector_string = value
}
type DependabotAlertSecurityAdvisory_cvssable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetScore()(*float64)
    GetVectorString()(*string)
    SetScore(value *float64)()
    SetVectorString(value *string)()
}
