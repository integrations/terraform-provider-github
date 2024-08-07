package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageVersion_metadata_docker struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The tag property
    tag []string
}
// NewPackageVersion_metadata_docker instantiates a new PackageVersion_metadata_docker and sets the default values.
func NewPackageVersion_metadata_docker()(*PackageVersion_metadata_docker) {
    m := &PackageVersion_metadata_docker{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePackageVersion_metadata_dockerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageVersion_metadata_dockerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPackageVersion_metadata_docker(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PackageVersion_metadata_docker) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageVersion_metadata_docker) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["tag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTag(res)
        }
        return nil
    }
    return res
}
// GetTag gets the tag property value. The tag property
// returns a []string when successful
func (m *PackageVersion_metadata_docker) GetTag()([]string) {
    return m.tag
}
// Serialize serializes information the current object
func (m *PackageVersion_metadata_docker) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetTag() != nil {
        err := writer.WriteCollectionOfStringValues("tag", m.GetTag())
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
func (m *PackageVersion_metadata_docker) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTag sets the tag property value. The tag property
func (m *PackageVersion_metadata_docker) SetTag(value []string)() {
    m.tag = value
}
type PackageVersion_metadata_dockerable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTag()([]string)
    SetTag(value []string)()
}
