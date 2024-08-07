package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespacesResponse 
// Deprecated: This class is obsolete. Use codespacesGetResponse instead.
type CodespacesResponse struct {
    CodespacesGetResponse
}
// NewCodespacesResponse instantiates a new CodespacesResponse and sets the default values.
func NewCodespacesResponse()(*CodespacesResponse) {
    m := &CodespacesResponse{
        CodespacesGetResponse: *NewCodespacesGetResponse(),
    }
    return m
}
// CreateCodespacesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCodespacesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesResponse(), nil
}
// CodespacesResponseable 
// Deprecated: This class is obsolete. Use codespacesGetResponse instead.
type CodespacesResponseable interface {
    CodespacesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
