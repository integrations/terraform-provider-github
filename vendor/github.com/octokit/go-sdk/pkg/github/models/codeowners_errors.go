package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeownersErrors a list of errors found in a repo's CODEOWNERS file
type CodeownersErrors struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The errors property
    errors []CodeownersErrors_errorsable
}
// NewCodeownersErrors instantiates a new CodeownersErrors and sets the default values.
func NewCodeownersErrors()(*CodeownersErrors) {
    m := &CodeownersErrors{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeownersErrorsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeownersErrorsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeownersErrors(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeownersErrors) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetErrors gets the errors property value. The errors property
// returns a []CodeownersErrors_errorsable when successful
func (m *CodeownersErrors) GetErrors()([]CodeownersErrors_errorsable) {
    return m.errors
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeownersErrors) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["errors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCodeownersErrors_errorsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CodeownersErrors_errorsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CodeownersErrors_errorsable)
                }
            }
            m.SetErrors(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CodeownersErrors) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetErrors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetErrors()))
        for i, v := range m.GetErrors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("errors", cast)
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
func (m *CodeownersErrors) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetErrors sets the errors property value. The errors property
func (m *CodeownersErrors) SetErrors(value []CodeownersErrors_errorsable)() {
    m.errors = value
}
type CodeownersErrorsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetErrors()([]CodeownersErrors_errorsable)
    SetErrors(value []CodeownersErrors_errorsable)()
}
