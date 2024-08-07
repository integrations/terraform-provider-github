package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemActionsRunsItemArtifactsResponse 
// Deprecated: This class is obsolete. Use artifactsGetResponse instead.
type ItemItemActionsRunsItemArtifactsResponse struct {
    ItemItemActionsRunsItemArtifactsGetResponse
}
// NewItemItemActionsRunsItemArtifactsResponse instantiates a new ItemItemActionsRunsItemArtifactsResponse and sets the default values.
func NewItemItemActionsRunsItemArtifactsResponse()(*ItemItemActionsRunsItemArtifactsResponse) {
    m := &ItemItemActionsRunsItemArtifactsResponse{
        ItemItemActionsRunsItemArtifactsGetResponse: *NewItemItemActionsRunsItemArtifactsGetResponse(),
    }
    return m
}
// CreateItemItemActionsRunsItemArtifactsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemActionsRunsItemArtifactsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsRunsItemArtifactsResponse(), nil
}
// ItemItemActionsRunsItemArtifactsResponseable 
// Deprecated: This class is obsolete. Use artifactsGetResponse instead.
type ItemItemActionsRunsItemArtifactsResponseable interface {
    ItemItemActionsRunsItemArtifactsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
