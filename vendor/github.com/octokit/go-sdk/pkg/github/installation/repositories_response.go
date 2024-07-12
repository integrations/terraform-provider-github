package installation

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RepositoriesResponse 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type RepositoriesResponse struct {
    RepositoriesGetResponse
}
// NewRepositoriesResponse instantiates a new RepositoriesResponse and sets the default values.
func NewRepositoriesResponse()(*RepositoriesResponse) {
    m := &RepositoriesResponse{
        RepositoriesGetResponse: *NewRepositoriesGetResponse(),
    }
    return m
}
// CreateRepositoriesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRepositoriesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRepositoriesResponse(), nil
}
// RepositoriesResponseable 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type RepositoriesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RepositoriesGetResponseable
}
