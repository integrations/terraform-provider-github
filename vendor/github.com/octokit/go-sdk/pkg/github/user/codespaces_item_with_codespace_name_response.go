package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespacesItemWithCodespace_nameResponse 
// Deprecated: This class is obsolete. Use WithCodespace_nameDeleteResponse instead.
type CodespacesItemWithCodespace_nameResponse struct {
    CodespacesItemWithCodespace_nameDeleteResponse
}
// NewCodespacesItemWithCodespace_nameResponse instantiates a new CodespacesItemWithCodespace_nameResponse and sets the default values.
func NewCodespacesItemWithCodespace_nameResponse()(*CodespacesItemWithCodespace_nameResponse) {
    m := &CodespacesItemWithCodespace_nameResponse{
        CodespacesItemWithCodespace_nameDeleteResponse: *NewCodespacesItemWithCodespace_nameDeleteResponse(),
    }
    return m
}
// CreateCodespacesItemWithCodespace_nameResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCodespacesItemWithCodespace_nameResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesItemWithCodespace_nameResponse(), nil
}
// CodespacesItemWithCodespace_nameResponseable 
// Deprecated: This class is obsolete. Use WithCodespace_nameDeleteResponse instead.
type CodespacesItemWithCodespace_nameResponseable interface {
    CodespacesItemWithCodespace_nameDeleteResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
