package user

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodespacesItemMachinesResponse 
// Deprecated: This class is obsolete. Use machinesGetResponse instead.
type CodespacesItemMachinesResponse struct {
    CodespacesItemMachinesGetResponse
}
// NewCodespacesItemMachinesResponse instantiates a new CodespacesItemMachinesResponse and sets the default values.
func NewCodespacesItemMachinesResponse()(*CodespacesItemMachinesResponse) {
    m := &CodespacesItemMachinesResponse{
        CodespacesItemMachinesGetResponse: *NewCodespacesItemMachinesGetResponse(),
    }
    return m
}
// CreateCodespacesItemMachinesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCodespacesItemMachinesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodespacesItemMachinesResponse(), nil
}
// CodespacesItemMachinesResponseable 
// Deprecated: This class is obsolete. Use machinesGetResponse instead.
type CodespacesItemMachinesResponseable interface {
    CodespacesItemMachinesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
