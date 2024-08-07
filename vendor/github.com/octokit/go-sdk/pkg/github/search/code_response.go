package search

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeResponse 
// Deprecated: This class is obsolete. Use codeGetResponse instead.
type CodeResponse struct {
    CodeGetResponse
}
// NewCodeResponse instantiates a new CodeResponse and sets the default values.
func NewCodeResponse()(*CodeResponse) {
    m := &CodeResponse{
        CodeGetResponse: *NewCodeGetResponse(),
    }
    return m
}
// CreateCodeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCodeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeResponse(), nil
}
// CodeResponseable 
// Deprecated: This class is obsolete. Use codeGetResponse instead.
type CodeResponseable interface {
    CodeGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
