package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PackageVersion_metadata struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The container property
    container PackageVersion_metadata_containerable
    // The docker property
    docker PackageVersion_metadata_dockerable
    // The package_type property
    package_type *PackageVersion_metadata_package_type
}
// NewPackageVersion_metadata instantiates a new PackageVersion_metadata and sets the default values.
func NewPackageVersion_metadata()(*PackageVersion_metadata) {
    m := &PackageVersion_metadata{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePackageVersion_metadataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePackageVersion_metadataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPackageVersion_metadata(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PackageVersion_metadata) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContainer gets the container property value. The container property
// returns a PackageVersion_metadata_containerable when successful
func (m *PackageVersion_metadata) GetContainer()(PackageVersion_metadata_containerable) {
    return m.container
}
// GetDocker gets the docker property value. The docker property
// returns a PackageVersion_metadata_dockerable when successful
func (m *PackageVersion_metadata) GetDocker()(PackageVersion_metadata_dockerable) {
    return m.docker
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PackageVersion_metadata) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["container"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePackageVersion_metadata_containerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainer(val.(PackageVersion_metadata_containerable))
        }
        return nil
    }
    res["docker"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePackageVersion_metadata_dockerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocker(val.(PackageVersion_metadata_dockerable))
        }
        return nil
    }
    res["package_type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePackageVersion_metadata_package_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageType(val.(*PackageVersion_metadata_package_type))
        }
        return nil
    }
    return res
}
// GetPackageType gets the package_type property value. The package_type property
// returns a *PackageVersion_metadata_package_type when successful
func (m *PackageVersion_metadata) GetPackageType()(*PackageVersion_metadata_package_type) {
    return m.package_type
}
// Serialize serializes information the current object
func (m *PackageVersion_metadata) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("container", m.GetContainer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("docker", m.GetDocker())
        if err != nil {
            return err
        }
    }
    if m.GetPackageType() != nil {
        cast := (*m.GetPackageType()).String()
        err := writer.WriteStringValue("package_type", &cast)
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
func (m *PackageVersion_metadata) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContainer sets the container property value. The container property
func (m *PackageVersion_metadata) SetContainer(value PackageVersion_metadata_containerable)() {
    m.container = value
}
// SetDocker sets the docker property value. The docker property
func (m *PackageVersion_metadata) SetDocker(value PackageVersion_metadata_dockerable)() {
    m.docker = value
}
// SetPackageType sets the package_type property value. The package_type property
func (m *PackageVersion_metadata) SetPackageType(value *PackageVersion_metadata_package_type)() {
    m.package_type = value
}
type PackageVersion_metadataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContainer()(PackageVersion_metadata_containerable)
    GetDocker()(PackageVersion_metadata_dockerable)
    GetPackageType()(*PackageVersion_metadata_package_type)
    SetContainer(value PackageVersion_metadata_containerable)()
    SetDocker(value PackageVersion_metadata_dockerable)()
    SetPackageType(value *PackageVersion_metadata_package_type)()
}
