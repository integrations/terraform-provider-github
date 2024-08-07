package search

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IssuesResponse 
// Deprecated: This class is obsolete. Use issuesGetResponse instead.
type IssuesResponse struct {
    IssuesGetResponse
}
// NewIssuesResponse instantiates a new IssuesResponse and sets the default values.
func NewIssuesResponse()(*IssuesResponse) {
    m := &IssuesResponse{
        IssuesGetResponse: *NewIssuesGetResponse(),
    }
    return m
}
// CreateIssuesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIssuesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIssuesResponse(), nil
}
// IssuesResponseable 
// Deprecated: This class is obsolete. Use issuesGetResponse instead.
type IssuesResponseable interface {
    IssuesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
