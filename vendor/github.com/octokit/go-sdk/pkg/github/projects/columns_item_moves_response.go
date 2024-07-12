package projects

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ColumnsItemMovesResponse 
// Deprecated: This class is obsolete. Use movesPostResponse instead.
type ColumnsItemMovesResponse struct {
    ColumnsItemMovesPostResponse
}
// NewColumnsItemMovesResponse instantiates a new ColumnsItemMovesResponse and sets the default values.
func NewColumnsItemMovesResponse()(*ColumnsItemMovesResponse) {
    m := &ColumnsItemMovesResponse{
        ColumnsItemMovesPostResponse: *NewColumnsItemMovesPostResponse(),
    }
    return m
}
// CreateColumnsItemMovesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateColumnsItemMovesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewColumnsItemMovesResponse(), nil
}
// ColumnsItemMovesResponseable 
// Deprecated: This class is obsolete. Use movesPostResponse instead.
type ColumnsItemMovesResponseable interface {
    ColumnsItemMovesPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
