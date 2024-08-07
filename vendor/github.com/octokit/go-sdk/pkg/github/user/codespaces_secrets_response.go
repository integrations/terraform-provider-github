package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespacesSecretsResponse 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type CodespacesSecretsResponse struct {
    CodespacesSecretsGetResponse
}
// NewCodespacesSecretsResponse instantiates a new CodespacesSecretsResponse and sets the default values.
func NewCodespacesSecretsResponse()(*CodespacesSecretsResponse) {
    m := &CodespacesSecretsResponse{
        CodespacesSecretsGetResponse: *NewCodespacesSecretsGetResponse(),
    }
    return m
}
// CreateCodespacesSecretsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCodespacesSecretsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesSecretsResponse(), nil
}
// CodespacesSecretsResponseable 
// Deprecated: This class is obsolete. Use secretsGetResponse instead.
type CodespacesSecretsResponseable interface {
    CodespacesSecretsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
