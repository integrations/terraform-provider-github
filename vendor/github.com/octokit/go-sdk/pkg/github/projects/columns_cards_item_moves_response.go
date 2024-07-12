package projects

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ColumnsCardsItemMovesResponse 
// Deprecated: This class is obsolete. Use movesPostResponse instead.
type ColumnsCardsItemMovesResponse struct {
    ColumnsCardsItemMovesPostResponse
}
// NewColumnsCardsItemMovesResponse instantiates a new ColumnsCardsItemMovesResponse and sets the default values.
func NewColumnsCardsItemMovesResponse()(*ColumnsCardsItemMovesResponse) {
    m := &ColumnsCardsItemMovesResponse{
        ColumnsCardsItemMovesPostResponse: *NewColumnsCardsItemMovesPostResponse(),
    }
    return m
}
// CreateColumnsCardsItemMovesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateColumnsCardsItemMovesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewColumnsCardsItemMovesResponse(), nil
}
// ColumnsCardsItemMovesResponseable 
// Deprecated: This class is obsolete. Use movesPostResponse instead.
type ColumnsCardsItemMovesResponseable interface {
    ColumnsCardsItemMovesPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
