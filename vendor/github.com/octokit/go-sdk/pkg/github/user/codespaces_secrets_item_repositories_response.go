package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespacesSecretsItemRepositoriesResponse 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type CodespacesSecretsItemRepositoriesResponse struct {
    CodespacesSecretsItemRepositoriesGetResponse
}
// NewCodespacesSecretsItemRepositoriesResponse instantiates a new CodespacesSecretsItemRepositoriesResponse and sets the default values.
func NewCodespacesSecretsItemRepositoriesResponse()(*CodespacesSecretsItemRepositoriesResponse) {
    m := &CodespacesSecretsItemRepositoriesResponse{
        CodespacesSecretsItemRepositoriesGetResponse: *NewCodespacesSecretsItemRepositoriesGetResponse(),
    }
    return m
}
// CreateCodespacesSecretsItemRepositoriesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCodespacesSecretsItemRepositoriesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesSecretsItemRepositoriesResponse(), nil
}
// CodespacesSecretsItemRepositoriesResponseable 
// Deprecated: This class is obsolete. Use repositoriesGetResponse instead.
type CodespacesSecretsItemRepositoriesResponseable interface {
    CodespacesSecretsItemRepositoriesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
