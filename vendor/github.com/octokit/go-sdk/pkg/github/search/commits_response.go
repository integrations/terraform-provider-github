package search

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommitsResponse 
// Deprecated: This class is obsolete. Use commitsGetResponse instead.
type CommitsResponse struct {
    CommitsGetResponse
}
// NewCommitsResponse instantiates a new CommitsResponse and sets the default values.
func NewCommitsResponse()(*CommitsResponse) {
    m := &CommitsResponse{
        CommitsGetResponse: *NewCommitsGetResponse(),
    }
    return m
}
// CreateCommitsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommitsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommitsResponse(), nil
}
// CommitsResponseable 
// Deprecated: This class is obsolete. Use commitsGetResponse instead.
type CommitsResponseable interface {
    CommitsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
