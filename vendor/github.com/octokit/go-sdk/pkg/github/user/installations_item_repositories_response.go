package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InstallationsItemRepositoriesResponse 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type InstallationsItemRepositoriesResponse struct {
    InstallationsItemRepositoriesGetResponse
}
// NewInstallationsItemRepositoriesResponse instantiates a new InstallationsItemRepositoriesResponse and sets the default values.
func NewInstallationsItemRepositoriesResponse()(*InstallationsItemRepositoriesResponse) {
    m := &InstallationsItemRepositoriesResponse{
        InstallationsItemRepositoriesGetResponse: *NewInstallationsItemRepositoriesGetResponse(),
    }
    return m
}
// CreateInstallationsItemRepositoriesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInstallationsItemRepositoriesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInstallationsItemRepositoriesResponse(), nil
}
// InstallationsItemRepositoriesResponseable 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type InstallationsItemRepositoriesResponseable interface {
    InstallationsItemRepositoriesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
