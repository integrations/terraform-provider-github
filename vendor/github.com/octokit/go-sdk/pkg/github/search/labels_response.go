package search

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LabelsResponse 
// Deprecated: This class is obsolete. Use labelsGetResponse instead.
type LabelsResponse struct {
    LabelsGetResponse
}
// NewLabelsResponse instantiates a new LabelsResponse and sets the default values.
func NewLabelsResponse()(*LabelsResponse) {
    m := &LabelsResponse{
        LabelsGetResponse: *NewLabelsGetResponse(),
    }
    return m
}
// CreateLabelsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLabelsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelsResponse(), nil
}
// LabelsResponseable 
// Deprecated: This class is obsolete. Use labelsGetResponse instead.
type LabelsResponseable interface {
    LabelsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
