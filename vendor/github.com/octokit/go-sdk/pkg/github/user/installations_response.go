package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InstallationsResponse 
// Deprecated: This class is obsolete. Use installationsGetResponse instead.
type InstallationsResponse struct {
    InstallationsGetResponse
}
// NewInstallationsResponse instantiates a new InstallationsResponse and sets the default values.
func NewInstallationsResponse()(*InstallationsResponse) {
    m := &InstallationsResponse{
        InstallationsGetResponse: *NewInstallationsGetResponse(),
    }
    return m
}
// CreateInstallationsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInstallationsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInstallationsResponse(), nil
}
// InstallationsResponseable 
// Deprecated: This class is obsolete. Use installationsGetResponse instead.
type InstallationsResponseable interface {
    InstallationsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
