package search

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TopicsResponse 
// Deprecated: This class is obsolete. Use topicsGetResponse instead.
type TopicsResponse struct {
    TopicsGetResponse
}
// NewTopicsResponse instantiates a new TopicsResponse and sets the default values.
func NewTopicsResponse()(*TopicsResponse) {
    m := &TopicsResponse{
        TopicsGetResponse: *NewTopicsGetResponse(),
    }
    return m
}
// CreateTopicsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTopicsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTopicsResponse(), nil
}
// TopicsResponseable 
// Deprecated: This class is obsolete. Use topicsGetResponse instead.
type TopicsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TopicsGetResponseable
}
